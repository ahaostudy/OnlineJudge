package submit

import (
	"context"
	rpcSubmit "main/api/submit"
	"main/internal/common"
	"main/internal/data/repository"

	"gorm.io/gorm"
)

func (SubmitServer) DeleteSubmit(_ context.Context, req *rpcSubmit.DeleteSubmitRequest) (resp *rpcSubmit.DeleteSubmitResponse, _ error) {
	resp = new(rpcSubmit.DeleteSubmitResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 获取提交数据
	submit, err := repository.GetSubmit(req.GetID())
	if err == gorm.ErrRecordNotFound || (err == nil && submit.UserID != req.GetUserID()) {
		resp.StatusCode = common.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 删除一条记录
	err = repository.DeleteSubmit(req.GetID())
	if err != nil {
		return
	}

	resp.StatusCode = common.CodeSuccess.Code()
	return
}
