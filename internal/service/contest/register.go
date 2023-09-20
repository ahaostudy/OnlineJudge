package contest

import (
	"context"
	"main/api/contest"
	"main/internal/common/code"
)

func (ContestServer) Register(ctx context.Context, req *rpcContest.RegisterRequest) (resp *rpcContest.RegisterResponse, _ error) {
	resp = new(rpcContest.RegisterResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	return
}

func (ContestServer) UnRegister(ctx context.Context, req *rpcContest.UnRegisterRequest) (resp *rpcContest.UnRegisterResponse, _ error) {
	resp = new(rpcContest.UnRegisterResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	return
}
