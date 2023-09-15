package problem

import (
	"context"
	"errors"
	rpcProblem "main/api/problem"
	"main/internal/common/code"
	"main/internal/common/build"

	"main/internal/data/repository"

	"gorm.io/gorm"
)


func (ProblemServer) GetProblem(ctx context.Context, req *rpcProblem.GetProblemRequest) (resp *rpcProblem.GetProblemResponse, _ error) {
	resp = new(rpcProblem.GetProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 访问数据库获取题目信息
	problem, err := repository.GetProblem_(req.GetProblemID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将模型对象转换为响应结果
	p, err := build.BuildProblem(problem)
	if err != nil {
		return
	}
	resp.Problem = p

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
