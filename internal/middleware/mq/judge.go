package mq

import (
	"encoding/json"
	"main/config"
	"main/internal/common/ctxt"
	"main/internal/data/model"
	"main/internal/middleware/redis"
	"main/internal/service/judge/pkg/code"
	"main/internal/service/judge/pkg/judge"
	"time"

	"github.com/streadway/amqp"
)

type (
	JudgeRequest struct {
		JudgeID  string
		Codepath string
		LangID   int
		Problem  *model.Problem
	}
	JudgeResponse struct {
		JudgeID string
		Result  code.Result
		Error   error
	}
)

func Judge(msg *amqp.Delivery) error {
	// TODO: 判题失败应该重新打入MQ重试
	req := new(JudgeRequest)
	if err := json.Unmarshal(msg.Body, req); err != nil {
		return err
	}

	res := JudgeResponse{JudgeID: req.JudgeID}

	// 执行判题
	res.Result, res.Error = judge.Judge(req.Problem, req.Codepath, req.LangID)
	bytes, err := json.Marshal(res)
	if err != nil {
		return err
	}

	// 将执行结果缓存到redis
	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()
	key, ttl := redis.GenerateJudgeKey(req.JudgeID), time.Duration(config.ConfRedis.ShortTtl)*time.Second
	if err := redis.Rdb.Set(ctx, key, bytes, ttl).Err(); err != nil {
		return err
	}

	return nil
}
