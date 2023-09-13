package main

import (
	"context"
	"fmt"
	"main/api/judge"
	"main/config"
	"main/rpc"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
}

func main() {
	conn, err := rpc.InitJudgeGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx := context.Background()

	// 提交代码
	res, err := rpc.JudgeCli.Judge(ctx, &rpcJudge.JudgeRequest{
		ProblemID: 2,
		Code: []byte(`
		#include <iostream>
		using namespace std;
		
		int main() {
			int a, b;
			cin >> a >> b;
			cout << a + b << endl;

			return 0
		}
		`),
		LangID: 2,
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
	fmt.Printf("result.StatusCode: %v\n", result.StatusCode)
	fmt.Printf("result.StatusMsg: %v\n", result.StatusMsg)
	fmt.Printf("result: %+v\n", result.Result)
}
