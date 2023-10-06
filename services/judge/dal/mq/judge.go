package mq

import (
	"context"
	"encoding/json"
	"main/services/judge/config"
	"main/services/judge/dal/cache"
	"main/services/judge/dal/model"
	"main/services/judge/pkg/code"
	"main/services/judge/pkg/judge"
	"path/filepath"
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

// GenerateJudgeMQMsg 生成判题服务消息
func GenerateJudgeMQMsg(judgeID string, codePath string, langID int64, problem *model.Problem) ([]byte, error) {
	req := JudgeRequest{
		JudgeID:  judgeID,
		Codepath: filepath.Join(config.Config.File.CodePath, codePath),
		LangID:   int(langID),
		Problem:  problem,
	}
	msg, err := json.Marshal(req)
	return msg, err
}

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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	key, ttl := cache.GenerateJudgeKey(req.JudgeID), time.Duration(config.Config.Redis.ShortTtl)*time.Second
	if err := cache.Rdb.Set(ctx, key, bytes, ttl).Err(); err != nil {
		return err
	}

	return nil
}
