package mq

import (
	"encoding/json"
	"main/internal/data/model"
	"main/internal/service/judge/pkg/code"
	"main/internal/service/judge/pkg/judge"

	"github.com/streadway/amqp"
)

type judgeRequest struct {
	JudgeID  string
	Codepath string
	LangID   int
	Problem  *model.Problem
}

type JudgeResponse struct {
	JudgeID string
	Result  code.Result
	Error   error
}

var ResultChan = make(map[string]chan JudgeResponse)

func GenerateJudgeMQMsg(judgeID string, codePath string, langID int64, problem *model.Problem) ([]byte, error) {
	req := judgeRequest{
		JudgeID:  judgeID,
		Codepath: codePath,
		LangID:   int(langID),
		Problem:  problem,
	}
	msg, err := json.Marshal(req)
	return msg, err
}

func Judge(msg *amqp.Delivery) error {
	req := new(judgeRequest)
	if err := json.Unmarshal(msg.Body, req); err != nil {
		return err
	}

	// 定义结果，使用defer确保有结果返回
	res := JudgeResponse{JudgeID: req.JudgeID}
	defer func() {
		if ch, ok := ResultChan[req.JudgeID]; ok && ch != nil {
			ch <- res
		}
	}()

	// 执行判题
	res.Result, res.Error = judge.Judge(req.Problem, req.Codepath, req.LangID)

	return nil
}
