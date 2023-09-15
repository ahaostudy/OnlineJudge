package testcase

import (
	"context"
	"errors"
	"main/api/testcase"
	"main/internal/common"
	"main/internal/data/repository"
	"os"

	"gorm.io/gorm"
)

func (TestcaseServer) DeleteTestcase(ctx context.Context, req *rpcTestcase.DeleteTestcaseRequest) (resp *rpcTestcase.DeleteTestcaseResponse, _ error) {
	resp = new(rpcTestcase.DeleteTestcaseResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 判断题目是否存在
	testcase, err := repository.GetTestcase(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = common.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 并发将该样例的输入输出文件删除
	go func() {
		if p, ok := testcase.GetLocalInput(); !ok {
			os.Remove(p)
		}
		if p, ok := testcase.GetLocalOutput(); !ok {
			os.Remove(p)
		}
	}()

	// 删除改样例数据
	if err = repository.DeleteTestcase(req.GetID()); err == nil {
		resp.StatusCode = common.CodeSuccess.Code()
	}

	return
}
