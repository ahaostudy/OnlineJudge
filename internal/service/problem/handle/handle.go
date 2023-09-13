package handle

import (
	"context"
	rpcProblem "main/api/problem"
	"main/internal/common"
	"main/internal/data/repository"
)

type ProblemServer struct {
	rpcProblem.UnimplementedProblemServiceServer
}

func (ProblemServer) GetProblem(ctx context.Context, req *rpcProblem.GetProblemRequest) (resp *rpcProblem.GetProblemResponse, _ error) {
	resp = new(rpcProblem.GetProblemResponse)
	resp.StatusCode = common.CodeServerBusy.Code()
	resp.StatusMsg = common.CodeServerBusy.Msg()

	// 访问数据库获取题目信息
	problem, err := repository.GetProblem_(req.GetProblemID())
	if err != nil {
		return
	}

	// 响应结果
	resp.StatusCode = common.CodeSuccess.Code()
	resp.StatusMsg = common.CodeSuccess.Msg()

	resp.Problem = new(rpcProblem.Problem)
	// 将模型对象转换为响应结果
	builder := new(common.Builder).Build(problem, resp.Problem)
	for i := range problem.Testcases {
		t := new(rpcProblem.Testcase)
		builder.Build(problem.Testcases[i], t)
		resp.Problem.Testcases = append(resp.Problem.Testcases, t)
	}
	if builder.Error() != nil {
		return
	}

	resp.StatusCode = common.CodeSuccess.Code()
	resp.StatusMsg = common.CodeSuccess.Msg()
	return
}
