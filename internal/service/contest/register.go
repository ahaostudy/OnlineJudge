package contest

import (
	"context"
	"errors"
	"main/api/contest"
	"main/internal/common/code"
	"main/internal/data/repository"
	"time"

	"gorm.io/gorm"
)

func (cs ContestServer) Register(ctx context.Context, req *rpcContest.RegisterRequest) (resp *rpcContest.RegisterResponse, _ error) {
	resp = new(rpcContest.RegisterResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 判断是否在比赛开始前
	contest, err := repository.GetContest(req.GetContestID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeContestNotExist.Code()
	}
	if err != nil {
		return
	}
	if contest.StartTime.UnixMilli() > time.Now().UnixMilli() {
		resp.StatusCode = code.CodeContestHasStarted.Code()
		return
	}

	// 判断用户是否已经报名
	isRegister := true
	_, err = repository.GetRegister(req.GetContestID(), req.GetUserID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		isRegister = false
	} else if err != nil {
		return
	}

	// 用户已经报名
	if isRegister {
		resp.StatusCode = code.CodeAlreadyRegistered.Code()
		return
	}

	// 插入一条报名记录
	err = repository.InsertRegister(req.GetContestID(), req.GetUserID())
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ContestServer) UnRegister(ctx context.Context, req *rpcContest.UnRegisterRequest) (resp *rpcContest.UnRegisterResponse, _ error) {
	resp = new(rpcContest.UnRegisterResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 判断是否在比赛开始前
	contest, err := repository.GetContest(req.GetContestID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeContestNotExist.Code()
	}
	if err != nil {
		return
	}
	if contest.StartTime.UnixMilli() > time.Now().UnixMilli() {
		resp.StatusCode = code.CodeContestHasStarted.Code()
		return
	}

	// 判断用户是否已经报名
	if err = repository.DeleteRegister(req.GetContestID(), req.GetUserID()); err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
