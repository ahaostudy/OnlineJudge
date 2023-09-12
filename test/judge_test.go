package test

import (
	"context"
	"fmt"
	"main/api/judge"
	"main/rpc"
	"testing"
)

func TestJudge(t *testing.T) {
	err := rpc.InitGRPCClients()
	if err != nil {
		panic(err)
	}
	defer rpc.CloseGPRCClients()

	ctx := context.Background()

	// 提交代码
	res, err := rpc.JudgeCli.Judge(ctx, &rpcJudge.JudgeRequest{
		ProblemID: 1,
		Code:      []byte("a, b = map(int, input().split())\nprint(a + b)"),
		LangID:    3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("提交成功: %v\n", res.JudgeID)

	// 获取结果
	result, err := rpc.JudgeCli.GetResult(ctx, &rpcJudge.GetResultRequest{
		JudgeID: res.JudgeID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("result.Result: %+v\n", result.Result)
}
