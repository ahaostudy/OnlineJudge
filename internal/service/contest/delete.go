package contest

import (
	"context"
	"main/api/contest"
	"main/internal/common/code"
	"main/internal/data/repository"
)

func (ContestServer) DeleteContest(ctx context.Context, req *rpcContest.DeleteContestRequest) (resp *rpcContest.DeleteContestResponse, _ error) {
	resp = new(rpcContest.DeleteContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 删除比赛记录
	err := repository.DeleteContest(req.GetId())
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
