package contest

import (
	"context"
	"main/api/contest"
	"main/internal/common/code"
)

func (ContestServer) ContestRank(ctx context.Context, req *rpcContest.ContestRankRequest) (resp *rpcContest.ContestRankResponse, _ error) {
	resp = new(rpcContest.ContestRankResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	return
}