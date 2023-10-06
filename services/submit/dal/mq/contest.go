package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"

	"main/common/status"
	"main/services/submit/dal/cache"
	"main/services/submit/dal/db"
)

type ContestSubmitRequest struct {
	ContestID int64
	UserID    int64
}

type RankStatus struct {
	Penalty  int   `json:"penalty"`
	Accepted bool  `json:"accepted"`
	AcTime   int64 `json:"ac_time"`
	LangID   int64 `json:"lang_id"`
	Score    int   `json:"score"`
}

// GenerateContestSubmitMQMsg 生成比赛提交服务消息
func GenerateContestSubmitMQMsg(contestID, userID int64) ([]byte, error) {
	req := ContestSubmitRequest{
		ContestID: contestID,
		UserID:    userID,
	}
	msg, err := json.Marshal(req)
	return msg, err
}

func ContestSubmit(msg *amqp.Delivery) error {
	req := new(ContestSubmitRequest)
	if err := json.Unmarshal(msg.Body, req); err != nil {
		return err
	}

	// 获取用户的提交记录
	submits, err := db.GetContestSubmitsByUser(req.ContestID, req.UserID)
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key := cache.GenerateContestUserKey(req.ContestID, req.UserID)

	// 遍历每题
	for p, ps := range subs {
		s := new(RankStatus)

		// 计算每题的分数
		for _, idx := range ps {
			if submits[idx].Status != int64(status.StatusAccepted) {
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
		if err := cache.Rdb.HSet(ctx, key, fmt.Sprintf("problem:%d", p), string(bytes)).Err(); err != nil {
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
	bytes, err := json.Marshal(&RankStatus{
		Penalty: penalty,
		AcTime:  acTime,
		Score:   score,
	})
	if err != nil {
		return err
	}
	if err := cache.Rdb.HSet(ctx, key, "all", string(bytes)).Err(); err != nil {
		return err
	}

	// 更新zset维护排行榜
	if err := cache.Rdb.ZRem(ctx, cache.GenerateRankKey(req.ContestID), req.UserID).Err(); err != nil {
		return err
	}
	if err := cache.Rdb.ZAdd(ctx, cache.GenerateRankKey(req.ContestID), &redis.Z{Score: float64(score), Member: req.UserID}).Err(); err != nil {
		return err
	}

	return nil
}
