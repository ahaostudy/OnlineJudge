package submit

import (
	"context"
	"main/api/submit"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/data/model"
	"main/internal/data/repository"
)

func (SubmitServer) GetSubmitList(ctx context.Context, req *rpcSubmit.GetSubmitListRequest) (resp *rpcSubmit.GetSubmitListResponse, _ error) {
	resp = new(rpcSubmit.GetSubmitListResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 获取提交数据
	var submitList []*model.Submit
	var err error
	if req.GetUserID() == 0 && req.ProblemID == 0 {
		return
	} else if req.GetUserID() == 0 {
		submitList, err = repository.GetSubmitListByProblem(req.GetProblemID())
	} else if req.GetProblemID() == 0 {
		submitList, err = repository.GetSubmitListByUser(req.GetUserID())
	} else {
		submitList, err = repository.GetSubmitList(req.GetUserID(), req.GetProblemID())
	}
	if err != nil {
		return
	}

	// 将对象转换为rpc响应
	resp.SubmitList, err = build.BuildSubmitList(submitList)
	if err != nil {
		return
	}

	resp.StatusCode = common.CodeSuccess.Code()
	return
}
