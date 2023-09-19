package contest

import (
	"context"
	rpcContest "main/api/contest"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/data/repository"
	"time"
)

func (ContestServer) CreateContest(ctx context.Context, req *rpcContest.CreateContestRequest) (resp *rpcContest.CreateContestResponse, _ error) {
	resp = new(rpcContest.CreateContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 插入一条比赛记录
	contest := &model.Contest{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		StartTime:   time.UnixMilli(req.GetStartTime()),
		EndTime:     time.UnixMilli(req.GetEndTime()),
	}
	err := repository.CreateContest(contest)
	if err != nil {
		return
	}

	// 返回响应
	resp.StatusCode = code.CodeSuccess.Code()
	return
}
