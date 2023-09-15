package mq

import (
	"encoding/json"
	"main/internal/data/model"
	"main/internal/service/judge/pkg/code"
	"main/internal/service/judge/pkg/judge"
	"sync"

	"github.com/streadway/amqp"
)

type (
	judgeRequest struct {
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

var (
	// ResultChan = make(map[string]chan JudgeResponse)
	ResultChan = sync.Map{}
	// DoneChan   = make(map[string]chan struct{})
	DoneChan = sync.Map{}
)

func GetResultChan(id string) (chan JudgeResponse, bool) {
	v, ok := ResultChan.Load(id)
	if !ok {
		return nil, ok
	}
	ch, ok := v.(chan JudgeResponse)
	return ch, ok
}

func GetDoneChan(id string) (chan struct{}, bool) {
	v, ok := DoneChan.Load(id)
	if !ok {
		return nil, ok
	}
	ch, ok := v.(chan struct{})
	return ch, ok
}

func Judge(msg *amqp.Delivery) error {
	req := new(judgeRequest)
	if err := json.Unmarshal(msg.Body, req); err != nil {
		return err
	}

	// 定义结果，使用defer确保有结果返回
	res := JudgeResponse{JudgeID: req.JudgeID}
	defer func() {
		if ch, ok := GetResultChan(req.JudgeID); ok && ch != nil {
			ch <- res
		}
	}()

	// 执行判题
	res.Result, res.Error = judge.Judge(req.Problem, req.Codepath, req.LangID)

	return nil
}
