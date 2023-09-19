package problem

import (
	"context"
	"errors"
	"main/api/contest"
	"main/api/problem"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/rpc"

	"main/internal/data/repository"

	"gorm.io/gorm"
)

// 返回值中的测试样例应返回前面两个作为示例
func (ProblemServer) GetProblem(ctx context.Context, req *rpcProblem.GetProblemRequest) (resp *rpcProblem.GetProblemResponse, _ error) {
	resp = new(rpcProblem.GetProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 访问数据库获取题目信息
	problem, err := repository.GetProblem_(req.GetProblemId())
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

func (ProblemServer) GetContestProblem(ctx context.Context, req *rpcProblem.GetContestProblemRequest) (resp *rpcProblem.GetContestProblemResponse, _ error) {
	resp = new(rpcProblem.GetContestProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取题目信息
	problem, err := repository.GetContestProblem(req.GetUserId(), req.GetProblemId())
	if err != nil {
		return nil, err
	}

	// 判断用户是否有访问权限
	res, err := rpc.ContestCli.IsAccessible(ctx, &rpcContest.IsAccessibleRequest{
		UserId:    req.GetUserId(),
		ContestId: problem.ContestID,
	})
	if err != nil {
		return
	}
	if res.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = res.StatusCode
		return
	}

	// 无访问权限
	if !res.GetIsAccessible() {
		resp.StatusCode = code.CodeNoRegistration.Code()
		return
	}

	// 将模型对象转换为响应结果
	resp.Problem, err = build.BuildProblem(problem)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ProblemServer) GetContestProblemList(ctx context.Context, req *rpcProblem.GetContestProblemListRequest) (resp *rpcProblem.GetContestProblemListResponse, _ error) {
	resp = new(rpcProblem.GetContestProblemListResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 判断用户是否有访问权限
	res, err := rpc.ContestCli.IsAccessible(ctx, &rpcContest.IsAccessibleRequest{
		UserId:    req.GetUserId(),
		ContestId: req.GetContestId(),
	})
	if err != nil {
		return
	}
	if res.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = res.StatusCode
		return
	}

	// 无访问权限
	if !res.GetIsAccessible() {
		resp.StatusCode = code.CodeNoRegistration.Code()
		return
	}

	// 获取题目列表
	problems, err := repository.GetContestProblemList(req.GetContestId())
	if err != nil {
		return
	}
	resp.Problems, err = build.BuildProblems(problems)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
