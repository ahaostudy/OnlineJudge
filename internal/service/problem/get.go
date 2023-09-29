package problem

import (
	"context"
	"errors"

	"gorm.io/gorm"

	rpcContest "main/api/contest"
	rpcProblem "main/api/problem"
	rpcSubmit "main/api/submit"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/repository"
	"main/rpc"
)

// 返回值中的测试样例应返回前面两个作为示例
func (ProblemServer) GetProblem(ctx context.Context, req *rpcProblem.GetProblemRequest) (resp *rpcProblem.GetProblemResponse, _ error) {
	resp = new(rpcProblem.GetProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 访问数据库获取题目信息
	problem, err := repository.GetProblemDetail(req.GetProblemID())
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

	// 获取示例内容
	for i := 0; i < problem.SampleCount && i < len(problem.Testcases); i++ {
		input, ok := problem.Testcases[i].GetInput()
		if !ok {
			return
		}
		output, ok := problem.Testcases[i].GetOutput()
		if !ok {
			return
		}
		resp.Problem.Samples = append(resp.Problem.Samples, &rpcProblem.Sample{
			Input:  input,
			Output: output,
		})
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ProblemServer) GetProblemList(ctx context.Context, req *rpcProblem.GetProblemListRequest) (resp *rpcProblem.GetProblemListResponse, _ error) {
	resp = new(rpcProblem.GetProblemListResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	page, count := int(req.GetPage()), int(req.GetCount())
	problemList, err := repository.GetProblemListLimit((page-1)*count, count)
	if err != nil {
		return
	}

	resp.ProblemList, err = build.BuildProblems(problemList)
	if err != nil {
		return
	}

	// 获取题目的提交情况
	submitStatus, err := rpc.SubmitCli.GetSubmitStatus(ctx, &rpcSubmit.GetSubmitStatusRequest{})
	if err != nil || submitStatus.StatusCode != code.CodeSuccess.Code() {
		return
	}
	// 判断当前用户是否ac
	acceptedStatus, err := rpc.SubmitCli.GetAcceptedStatus(ctx, &rpcSubmit.GetAcceptedStatusRequest{UserID: req.GetUserID()})
	if err != nil || acceptedStatus.StatusCode != code.CodeSuccess.Code() {
		return
	}

	for i := range resp.ProblemList {
		id := resp.ProblemList[i].ID
		if v, ok := submitStatus.SubmitStatus[id]; ok {
			resp.ProblemList[i].SubmitCount = v.Count
			resp.ProblemList[i].AcceptedCount = v.AcceptedCount
		}
		if v, ok := acceptedStatus.AcceptedStatus[id]; ok {
			resp.ProblemList[i].IsAccepted = v
		}
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ProblemServer) GetProblemCount(ctx context.Context, req *rpcProblem.GetProblemCountRequest) (resp *rpcProblem.GetProblemCountResponse, _ error) {
	resp = new(rpcProblem.GetProblemCountResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	count, err := repository.GetProblemCount()
	if err != nil {
		return
	}

	resp.Count = count
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (ProblemServer) GetContestProblem(ctx context.Context, req *rpcProblem.GetContestProblemRequest) (resp *rpcProblem.GetContestProblemResponse, _ error) {
	resp = new(rpcProblem.GetContestProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取题目信息
	problem, err := repository.GetContestProblem(req.GetUserID(), req.GetProblemID())
	if err != nil {
		return nil, err
	}

	// 判断用户是否有访问权限
	res, err := rpc.ContestCli.IsAccessible(ctx, &rpcContest.IsAccessibleRequest{
		UserID:    req.GetUserID(),
		ContestID: problem.ContestID,
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
		resp.StatusCode = code.CodeNotRegistred.Code()
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
		UserID:    req.GetUserID(),
		ContestID: req.GetContestID(),
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
		resp.StatusCode = code.CodeNotRegistred.Code()
		return
	}

	// 获取题目列表
	problems, err := repository.GetContestProblemList(req.GetContestID())
	if err != nil {
		return
	}
	resp.ProblemList, err = build.BuildProblems(problems)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
