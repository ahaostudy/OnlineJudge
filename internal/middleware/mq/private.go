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
	PrivateJudgeRequest struct {
		JudgeID  string
		Codepath string
		LangID   int
		Problem  *model.Problem
	}
	PrivateJudgeResponse struct {
		JudgeID string
		Result  code.Result
		Error   error
	}
)

var (
	// ResultChan = make(map[string]chan JudgeResponse)
	PrivateResultChan = sync.Map{}
	// DoneChan   = make(map[string]chan struct{})
	PrivateDoneChan = sync.Map{}
)

func GetPrivateResultChan(id string) (chan PrivateJudgeResponse, bool) {
	v, ok := PrivateResultChan.Load(id)
	if !ok {
		return nil, ok
	}
	ch, ok := v.(chan PrivateJudgeResponse)
	return ch, ok
}

func GetPrivateDoneChan(id string) (chan struct{}, bool) {
	v, ok := PrivateDoneChan.Load(id)
	if !ok {
		return nil, ok
	}
	ch, ok := v.(chan struct{})
	return ch, ok
}

func PrivateJudge(msg *amqp.Delivery) error {
	req := new(PrivateJudgeRequest)
	if err := json.Unmarshal(msg.Body, req); err != nil {
		return err
	}

	// 定义结果，使用defer确保有结果返回
	res := PrivateJudgeResponse{JudgeID: req.JudgeID}
	defer func() {
		if ch, ok := GetPrivateResultChan(req.JudgeID); ok && ch != nil {
			ch <- res
		}
	}()

	// 执行判题
	res.Result, res.Error = judge.Judge(req.Problem, req.Codepath, req.LangID)

	return nil
}
