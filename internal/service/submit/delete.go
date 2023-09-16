package submit

import (
	"context"
	rpcSubmit "main/api/submit"
	"main/internal/common"
	"main/internal/data/repository"
)

func (SubmitServer) DeleteSubmit(_ context.Context, req *rpcSubmit.DeleteSubmitRequest) (resp *rpcSubmit.DeleteSubmitResponse, _ error) {
	resp = new(rpcSubmit.DeleteSubmitResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 删除一条记录
	err := repository.DeleteSubmit(req.GetID())
	if err != nil {
		return
	}

	resp.StatusCode = common.CodeSuccess.Code()
	return
}
