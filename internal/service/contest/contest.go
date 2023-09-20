package contest

import (
	"context"
	rpcContest "main/api/contest"
	"main/internal/common/code"
	"main/internal/common/request"
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
	err := repository.InsertContest(contest)
	if err != nil {
		return
	}

	// 返回响应
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ContestServer) DeleteContest(ctx context.Context, req *rpcContest.DeleteContestRequest) (resp *rpcContest.DeleteContestResponse, _ error) {
	resp = new(rpcContest.DeleteContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 删除比赛记录
	err := repository.DeleteContest(req.GetID())
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ContestServer) UpdateContest(ctx context.Context, req *rpcContest.UpdateContestRequest) (resp *rpcContest.UpdateContestResponse, _ error) {
	resp = new(rpcContest.UpdateContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 更新比赛信息
	// TODO: 比赛结束后不应该再更新
	r := new(request.Request)
	if err := r.ReadRawData(req.GetContest()); err != nil {
		return
	}
	if r.Exists("start_time") {
		r.Set("start_time", time.UnixMilli(r.GetInt64("start_time")))
	}
	if r.Exists("end_time") {
		r.Set("end_time", time.UnixMilli(r.GetInt64("end_time")))
	}
	if err := repository.UpdateContest(req.GetID(), r.Map()); err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
