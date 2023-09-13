package test

import (
	"context"
	"fmt"
	rpcPrivate "main/api/private"
	"main/rpc"
	"testing"
)

func TestPrivate(t *testing.T) {
	// 初始化
	err := rpc.InitGRPCClients()
	if err != nil {
		panic(err)
	}
	defer rpc.CloseGPRCClients()

	// 获取题目信息
	res, err := rpc.PrivateCli.GetProblem(context.Background(), &rpcPrivate.GetProblemRequest{
		ProblemID: 3,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("res: %+v\n", res)
}
