package problem

import (
	"context"
	"errors"
	"fmt"
	"main/api/problem"
	"main/internal/common/code"
	"main/internal/common/build"
	"main/internal/data/model"
	"main/internal/data/repository"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (ProblemServer) GetTestcase(ctx context.Context, req *rpcProblem.GetTestcaseRequest) (resp *rpcProblem.GetTestcaseResponse, _ error) {
	resp = new(rpcProblem.GetTestcaseResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取题目信息
	testcase, err := repository.GetTestcase(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将对象转换为rpc响应
	resp.Testcase, err = build.BuildTestcase(testcase)
	if err == nil {
		resp.StatusCode = code.CodeSuccess.Code()
	}

	return
}

func (ProblemServer) CreateTestcase(ctx context.Context, req *rpcProblem.CreateTestcaseRequest) (resp *rpcProblem.CreateTestcaseResponse, _ error) {
	resp = new(rpcProblem.CreateTestcaseResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 将输入输出保存到本地文件
	inputPath := fmt.Sprintf("%d/%s.in", req.GetProblemID(), uuid.New().String())
	outputPath := fmt.Sprintf("%d/%s.out", req.GetProblemID(), uuid.New().String())

	// 创建样例对象
	testcase := &model.Testcase{
		ProblemID:  req.GetProblemID(),
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
	// 写入文件
	if !testcase.UploadInput(req.GetInput()) || !testcase.UploadOutput(req.GetOutput()) {
		return
	}
	// 将对象插入数据库
	if repository.InsertTestcase(testcase) == nil {
		resp.StatusCode = code.CodeSuccess.Code()
	}

	return
}

func (ProblemServer) DeleteTestcase(ctx context.Context, req *rpcProblem.DeleteTestcaseRequest) (resp *rpcProblem.DeleteTestcaseResponse, _ error) {
	resp = new(rpcProblem.DeleteTestcaseResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 判断题目是否存在
	testcase, err := repository.GetTestcase(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
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
		resp.StatusCode = code.CodeSuccess.Code()
	}

	return
}
