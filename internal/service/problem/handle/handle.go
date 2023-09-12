package handle

import (
	"context"
	"fmt"
	rpcProblem "main/api/problem"
	"main/internal/common"
	"main/internal/data/repository"
)

type ProblemServer struct {
	rpcProblem.UnimplementedProblemServiceServer
}

func (ProblemServer) GetProblem(ctx context.Context, req *rpcProblem.GetProblemRequest) (resp *rpcProblem.GetProblemResponse, _ error) {
	resp = new(rpcProblem.GetProblemResponse)

	problem, err := repository.GetProblem(req.GetProblemID())
	if err != nil {
		fmt.Println(err.Error())
		code := common.CodeServerBusy
		resp.StatusCode = code.Code()
		resp.StatusMsg = code.Msg()
		return
	}

	resp.StatusCode = common.CodeSuccess.Code()
	resp.StatusMsg = common.CodeSuccess.Msg()
	resp.Problem = new(rpcProblem.Problem)
	if err := common.Build(problem, resp.Problem); err != nil {
		print(err.Error())
		resp.StatusCode = common.CodeServerBusy.Code()
		resp.StatusMsg = common.CodeServerBusy.Msg()
		return
	}
	fmt.Println(problem)
	fmt.Println(resp.Problem)
	return
}
