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

	"gorm.io/gorm"
)

func (ContestServer) GetContest(ctx context.Context, req *rpcContest.GetContestRequest) (resp *rpcContest.GetContestResponse, _ error) {
	resp = new(rpcContest.GetContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 从数据库获取比赛信息
	contest, err := repository.GetContest(req.GetId())
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
		UserId:    req.GetUserId(),
		ContestId: req.GetId(),
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
