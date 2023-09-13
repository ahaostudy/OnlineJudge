package handle

import (
	"context"
	"main/api/testcase"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/data/repository"
)

func (TestcaseServer) GetTestcase(ctx context.Context, req *rpcTestcase.GetTestcaseRequest) (resp *rpcTestcase.GetTestcaseResponse, _ error) {
	resp = new(rpcTestcase.GetTestcaseResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 获取题目信息
	testcase, err := repository.GetTestcase(req.GetID())
	if err != nil {
		return
	}
	resp.Testcase, err = build.BuildTestcase(testcase)
	if err == nil {
		resp.StatusCode = common.CodeSuccess.Code()
	}

	return
}
