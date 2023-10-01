package mq

import (
	"encoding/json"
	"fmt"

	rds "github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"

	"main/internal/common/ctxt"
	"main/internal/data/repository"
	"main/internal/middleware/redis"
	"main/internal/service/judge/pkg/code"
)

type ContestSubmitRequest struct {
	ContestID int64
	UserID    int64
}

func ContestSubmit(msg *amqp.Delivery) error {
	req := new(ContestSubmitRequest)
	if err := json.Unmarshal(msg.Body, req); err != nil {
		return err
	}

	// 获取用户的提交记录
	submits, err := repository.GetContestSubmitsByUser(req.ContestID, req.UserID)
	if err != nil {
		return err
	}

	// 将每题的提交记录分开
	subs := make(map[int64][]int)
	for i, s := range submits {
		subs[s.ProblemID] = append(subs[s.ProblemID], i)
	}

	var score, penalty int
	var acTime int64
	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()
	key := redis.GenerateContestUserKey(req.ContestID, req.UserID)

	// 遍历每题
	for p, ps := range subs {
		s := new(contest.RankStatus)

		// 计算每题的分数
		for _, idx := range ps {
			if submits[idx].Status != int64(code.StatusAccepted) {
				s.Penalty = len(ps)
			} else {
				s.Accepted = true
				s.AcTime = submits[ps[0]].CreatedAt.UnixMilli()
				s.LangID = submits[ps[0]].LangID
				s.Score = 20
			}
		}
		bytes, err := json.Marshal(s)
		if err != nil {
			return err
		}

		// 更新该题的分数
		if err := redis.Rdb.HSet(ctx, key, fmt.Sprintf("problem:%d", p), string(bytes)).Err(); err != nil {
			return err
		}

		// 更新该用户的成绩变量
		if s.Accepted {
			score += s.Score
			penalty += s.Penalty
			if acTime < s.AcTime {
				acTime = s.AcTime
			}
		}
	}

	// 更新用户的总成绩
	bytes, err := json.Marshal(&contest.RankStatus{
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

	// 更新zset维护排行榜
	if err := redis.Rdb.ZRem(ctx, redis.GenerateRankKey(req.ContestID), req.UserID).Err(); err != nil {
		return err
	}
	if err := redis.Rdb.ZAdd(ctx, redis.GenerateRankKey(req.ContestID), &rds.Z{Score: float64(score), Member: req.UserID}).Err(); err != nil {
		return err
	}

	return nil
}
