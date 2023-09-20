package contest

import (
	"context"
	"main/api/contest"
	"main/internal/common/code"
	"main/internal/common/request"
	"main/internal/data/repository"
)

func (ContestServer) UpdateContest(ctx context.Context, req *rpcContest.UpdateContestRequest) (resp *rpcContest.UpdateContestResponse, _ error) {
	resp = new(rpcContest.UpdateContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 更新比赛信息
	// TODO: 比赛结束后不应该再更新
	r := new(request.Request)
	r.ReadRawData(req.GetContest())
	err := repository.UpdateContest(req.GetId(), r.Map())
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
