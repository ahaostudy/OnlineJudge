package contest

import (
	"context"
	"main/api/contest"
	"main/internal/common/code"
	"main/internal/data/repository"
)

func (ContestServer) IsRegister(ctx context.Context, req *rpcContest.IsRegisterRequest) (resp *rpcContest.IsRegisterResponse, _ error) {
	resp = new(rpcContest.IsRegisterResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 从数据库获取比赛信息

	return
}
