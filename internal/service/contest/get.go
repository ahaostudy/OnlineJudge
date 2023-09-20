package contest

import (
	"context"
	"errors"
	"main/api/contest"
	rpcProblem "main/api/problem"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/repository"
	"main/rpc"
	"time"

	"gorm.io/gorm"
)

func (ContestServer) GetContest(ctx context.Context, req *rpcContest.GetContestRequest) (resp *rpcContest.GetContestResponse, _ error) {
	resp = new(rpcContest.GetContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 从数据库获取比赛信息
	contest, err := repository.GetContest(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将比赛信息转换为rpc响应结构
	resp.Contest, err = build.BuildContest(contest)
	if err != nil {
		return
	}

	// 获取比赛题目
	// 里面会自带鉴权内容
	res, err := rpc.ProblemCli.GetContestProblemList(ctx, &rpcProblem.GetContestProblemListRequest{
		UserID:    req.GetUserID(),
		ContestID: req.GetID(),
	})
	if err != nil {
		return
	}
	if res.GetStatusCode() != code.CodeSuccess.Code() {
		resp.StatusCode = res.GetStatusCode()
		return
	}

	resp.Contest.ProblemList = res.GetProblemList()
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ContestServer) GetContestList(ctx context.Context, req *rpcContest.GetContestListRequest) (resp *rpcContest.GetContestListResponse, _ error) {
	resp = new(rpcContest.GetContestListResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取比赛列表
	page, count := int(req.GetPage()), int(req.GetCount())
	contestList, err := repository.GetContestList((page-1)*count, count)
	if err != nil {
		return
	}

	// 将比赛信息转换为rpc响应结构
	resp.ContestList, err = build.BuildContestList(contestList)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (cs ContestServer) IsAccessible(ctx context.Context, req *rpcContest.IsAccessibleRequest) (resp *rpcContest.IsAccessibleResponse, _ error) {
	resp = new(rpcContest.IsAccessibleResponse)
	resp.StatusCode = code.CodeSuccess.Code()

	// 用户未登录则直接返回
	if req.GetUserID() <= 0 {
		resp.IsAccessible = false
		return
	}

	// 获取比赛信息
	contest, err := repository.GetContest(req.GetContestID())
	if err == gorm.ErrRecordNotFound {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		resp.StatusCode = code.CodeServerBusy.Code()
		return
	}

	now := time.Now()
	// 比赛开始前
	if now.Before(contest.StartTime) {
		resp.StatusCode = code.CodeContestNotStarted.Code()
		return
	}
	// 比赛结束后
	if now.After(contest.EndTime) {
		resp.IsAccessible = true
		return
	}

	// 比赛过程中
	// 判断用户是否已经报名了比赛
	_, err = repository.GetRegister(req.GetContestID(), req.GetUserID())
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

func (ContestServer) IsRegister(ctx context.Context, req *rpcContest.IsRegisterRequest) (resp *rpcContest.IsRegisterResponse, _ error) {
	resp = new(rpcContest.IsRegisterResponse)
	resp.StatusCode = code.CodeSuccess.Code()

	// 用户未登录则直接返回
	if req.GetUserID() <= 0 {
		resp.IsRegister = false
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	// 从数据库获取报名信息
	_, err := repository.GetRegister(req.GetContestID(), req.GetUserID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.IsRegister = false
		return
	}
	if err != nil {
		resp.StatusCode = code.CodeServerBusy.Code()
		return
	}

	resp.IsRegister = true
	return
}
