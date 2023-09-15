package testcase

import (
	"context"
	"errors"
	"main/api/testcase"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/data/repository"

	"gorm.io/gorm"
)

func (TestcaseServer) GetTestcase(ctx context.Context, req *rpcTestcase.GetTestcaseRequest) (resp *rpcTestcase.GetTestcaseResponse, _ error) {
	resp = new(rpcTestcase.GetTestcaseResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 获取题目信息
	testcase, err := repository.GetTestcase(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = common.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将对象转换为rpc响应
	resp.Testcase, err = build.BuildTestcase(testcase)
	if err == nil {
		resp.StatusCode = common.CodeSuccess.Code()
	}

	return
}
