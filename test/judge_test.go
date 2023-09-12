package test

import (
	"context"
	"fmt"
	"main/clients"
	rpcJudge "main/rpc/judge"
	"testing"
)

func TestJudge(t *testing.T) {
	conn, err := clients.InitJudgeGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx := context.Background()

	// 提交代码
	res, err := clients.JudgeCli.Judge(ctx, &rpcJudge.JudgeRequest{
		ProblemID: 1,
		Code:      []byte("a, b = map(int, input().split())\nprint(a + b)"),
		LangID:    3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("提交成功: %v\n", res.JudgeID)

	// 获取结果
	result, err := clients.JudgeCli.GetResult(ctx, &rpcJudge.GetResultRequest{
		JudgeID: res.JudgeID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("result.Result: %+v\n", result.Result)
}
