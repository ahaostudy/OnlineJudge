package submit

import (
	"context"
	"errors"
	rpcSubmit "main/api/submit"
	"main/internal/common/code"
	"main/internal/common/build"
	"main/internal/data/repository"

	"gorm.io/gorm"
)

func (SubmitServer) GetSubmit(_ context.Context, req *rpcSubmit.GetSubmitRequest) (resp *rpcSubmit.GetSubmitResponse, _ error) {
	resp = new(rpcSubmit.GetSubmitResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交数据
	submit, err := repository.GetSubmit(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将模型对象转换为rpc响应
	resp.Submit, err = build.BuildSubmit(submit)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
