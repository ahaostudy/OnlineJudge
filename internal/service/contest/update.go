package contest

import (
	"context"
	"main/api/contest"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/data/repository"
	"time"
)

func (ContestServer) UpdateContest(ctx context.Context, req *rpcContest.UpdateContestRequest) (resp *rpcContest.UpdateContestResponse, _ error) {
	resp = new(rpcContest.UpdateContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 更新比赛记录
	contest := &model.Contest{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		StartTime:   time.UnixMilli(req.GetStartTime()),
		EndTime:     time.UnixMilli(req.GetEndTime()),
	}
	err := repository.UpdateContest(req.GetId(), contest)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
