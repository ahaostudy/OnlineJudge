package problem

import (
	"context"
	"main/api/problem"
	"main/internal/common/build"
	"main/internal/common/code"
	reqeust "main/internal/common/request"
	"main/internal/data/repository"
)

func (ProblemServer) CreateProblem(ctx context.Context, req *rpcProblem.CreateProblemRequest) (resp *rpcProblem.CreateProblemResponse, _ error) {
	resp = new(rpcProblem.CreateProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 将参数转换为题目对象
	problem, err := build.UnBuildProblem(req.GetProblem())
	if err != nil {
		return
	}
	problem.ID = 0

	// 插入一条题目信息
	if err := repository.InsertProblem(problem); err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ProblemServer) UpdateProblem(ctx context.Context, req *rpcProblem.UpdateProblemRequest) (resp *rpcProblem.UpdateProblemResponse, _ error) {
	resp = new(rpcProblem.UpdateProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 解析参数
	r := new(reqeust.Request)
	if err := r.ReadRawData(req.GetProblem()); err != nil {
		return
	}
	// 忽略id和author_id字段
	delete(r.Map(), "id")
	delete(r.Map(), "author_id")

	// 更新题目
	if err := repository.UpdateProblem(req.GetProblemID(), r.Map()); err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ProblemServer) DeleteProblem(ctx context.Context, req *rpcProblem.DeleteProblemRequest) (resp *rpcProblem.DeleteProblemResponse, _ error) {
	resp = new(rpcProblem.DeleteProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 删除题目
	if err := repository.DeleteProblem(req.GetProblemID()); err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
