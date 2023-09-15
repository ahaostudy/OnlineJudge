package testcase

import (
	"context"
	"fmt"
	"main/api/testcase"
	"main/internal/common"
	"main/internal/data/model"
	"main/internal/data/repository"

	"github.com/google/uuid"
)

func (TestcaseServer) CreateTestcase(ctx context.Context, req *rpcTestcase.CreateTestcaseRequest) (resp *rpcTestcase.CreateTestcaseResponse, _ error) {
	resp = new(rpcTestcase.CreateTestcaseResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 将输入输出保存到本地文件
	inputPath := fmt.Sprintf("%d/%s.in", req.ProblemID, uuid.New().String())
	outputPath := fmt.Sprintf("%d/%s.out", req.ProblemID, uuid.New().String())

	// 创建样例对象
	testcase := &model.Testcase{
		ProblemID:  req.ProblemID,
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
	// 写入文件
	if !testcase.UploadInput(req.GetInput()) || !testcase.UploadOutput(req.GetOutput()) {
		return
	}
	// 将对象插入数据库
	if repository.InsertTestcase(testcase) == nil {
		resp.StatusCode = common.CodeSuccess.Code()
	}

	return
}
