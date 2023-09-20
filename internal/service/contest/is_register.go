package contest

import (
	"context"
	"errors"
	"main/api/contest"
	"main/internal/common/code"
	"main/internal/data/repository"

	"gorm.io/gorm"
)

func (ContestServer) IsRegister(ctx context.Context, req *rpcContest.IsRegisterRequest) (resp *rpcContest.IsRegisterResponse, _ error) {
	resp = new(rpcContest.IsRegisterResponse)
	resp.StatusCode = code.CodeSuccess.Code()

	// 用户未登录则直接返回
	if req.GetUserId() <= 0 {
		resp.IsRegister = false
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	// 从数据库获取报名信息
	_, err := repository.GetRegister(req.GetContestId(), req.GetUserId())
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
