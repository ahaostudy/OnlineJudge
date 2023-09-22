package mq

import (
	"encoding/json"
	"fmt"
	"main/internal/common/ctxt"
	"main/internal/data/repository"
	"main/internal/service/contest"
	"main/internal/service/judge/pkg/code"
	"strings"

	"main/internal/middleware/redis"

	rds "github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

type contestSubmitRequest struct {
	ContestID int64
	ProblemID int64
	UserID    int64
}

func ContestSubmit(msg *amqp.Delivery) error {
	// TODO: 写入redis失败后，应考虑重新打入mq重试，或者在定时任务中重新计算

	req := new(contestSubmitRequest)
	if err := json.Unmarshal(msg.Body, req); err != nil {
		return err
	}

	// 1. 获取该题目的提交记录并统计该题的做题情况
	submits, err := repository.GetContestUserProblemSubmits(req.ContestID, req.ProblemID, req.UserID)
	if err != nil {
		return err
	}
	status := new(contest.RankStatus)
	for _, s := range submits {
		if s.Status != int64(code.StatusAccepted) {
			status.Penalty++
		} else {
			status.Accepted = true
			status.AcTime = s.CreatedAt.UnixMilli()
			status.LangID = s.LangID
			status.Score = 20
			break
		}
	}
	bytes, err := json.Marshal(status)
	if err != nil {
		return err
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 2. 更新该赛题的提交情况
	key := redis.GenerateContestUserKey(req.ContestID, req.UserID)
	err = redis.Rdb.HSet(ctx, key, fmt.Sprintf("problem:%d", req.ProblemID), string(bytes)).Err()
	if err != nil {
		return err
	}

	// 3. 计算并更新该场比赛用户的总分
	pss, err := redis.Rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}
	var score, penalty int
	var acTime int64
	for k, v := range pss {
		if !strings.HasPrefix(k, "problem:") {
			continue
		}
		status := new(contest.RankStatus)
		if err := json.Unmarshal([]byte(v), status); err != nil {
			return err
		}
		// TODO: 应遵循不同赛制计分
		// TODO: 应该考虑每题的分数占比不一样的情况
		// 目前实现为：每题20分，忽略罚时
		if status.Accepted {
			score += status.Score
			penalty += status.Penalty
			if acTime < status.AcTime {
				acTime = status.AcTime
			}
		}
	}
	bytes, err = json.Marshal(&contest.RankStatus{
		Penalty: penalty,
		AcTime:  acTime,
		Score:   score,
	})
	if err != nil {
		return err
	}
	if err := redis.Rdb.HSet(ctx, key, "all", string(bytes)).Err(); err != nil {
		return err
	}

	// 4. 更新zset维护排行榜
	if err := redis.Rdb.ZRem(ctx, redis.GenerateRankKey(req.ContestID), req.UserID).Err(); err != nil {
		return err
	}
	err = redis.Rdb.ZAdd(ctx, redis.GenerateRankKey(req.ContestID), &rds.Z{Score: float64(score), Member: req.UserID}).Err()

	return err
}
