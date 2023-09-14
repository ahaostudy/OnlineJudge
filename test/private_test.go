package test

import (
	"context"
	"fmt"
	rpcPrivate "main/api/private"
	"main/rpc"
	"testing"
	"time"
)

func TestPrivate(t *testing.T) {
	// 初始化
	err := rpc.InitGRPCClients()
	if err != nil {
		panic(err)
	}
	defer rpc.CloseGPRCClients()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 获取题目信息
	problem, err := rpc.PrivateCli.GetProblem(ctx, &rpcPrivate.GetProblemRequest{
		ProblemID: 2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("res: %+v\n", problem)

	// 提交代码
	judgeRes, err := rpc.PrivateCli.Judge(ctx, &rpcPrivate.JudgeRequest{
		ProblemID: problem.Problem.ID,
		Code:      []byte("a, b = map(int, input().split())\nprint(a + b)\n"),
		LangID:    3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("提交成功: %v\n", judgeRes.JudgeID)

	// 获取结果
	result, err := rpc.PrivateCli.GetResult(ctx, &rpcPrivate.GetResultRequest{
		JudgeID: judgeRes.JudgeID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("result.StatusCode: %v\n", result.StatusCode)
	fmt.Printf("result.Result: %+v\n", result.Result)
}
