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

func (cs ContestServer) IsAccessible(ctx context.Context, req *rpcContest.IsAccessibleRequest) (resp *rpcContest.IsAccessibleResponse, _ error) {
	resp = new(rpcContest.IsAccessibleResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 用户未登录则直接返回
	if req.GetUserId() <= 0 {
		resp.IsAccessible = false
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	// 获取比赛信息
	contest, err := repository.GetContest(req.GetContestId())
	if err == gorm.ErrRecordNotFound {
		resp.IsAccessible = false
		return
	}
	if err != nil {
		resp.StatusCode = code.CodeServerBusy.Code()
		return
	}

	now := time.Now()
	// 比赛开始前
	if now.Before(contest.StartTime) {
		resp.IsAccessible = false
		return
	}
	// 比赛结束后
	if now.After(contest.EndTime) {
		resp.IsAccessible = true
		return
	}

	// 比赛过程中
	// 判断用户是否已经报名了比赛
	_, err = repository.GetRegister(req.GetContestId(), req.GetUserId())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.IsAccessible = false
		return
	}
	if err != nil {
		resp.StatusCode = code.CodeServerBusy.Code()
		return
	}

	resp.IsAccessible = true
	return
}
