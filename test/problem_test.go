package test

import (
	"context"
	"fmt"
	rpcProblem "main/api/problem"
	"main/rpc"
	"testing"
)

func TestProblem(t *testing.T) {
	// 初始化
	err := rpc.InitGRPCClients()
	if err != nil {
		panic(err)
	}
	defer rpc.CloseGPRCClients()

	// 获取题目信息
	res, err := rpc.ProblemCli.GetProblem(context.Background(), &rpcProblem.GetProblemRequest{
		ProblemID: 2,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("res: %+v\n", res)
}
