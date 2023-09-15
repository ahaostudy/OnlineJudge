package test

import (
	"fmt"
	rpcSubmit "main/api/submit"
	"main/internal/common/ctxt"
	"main/rpc"
	"testing"
)

func TestSubmit(t *testing.T) {
	_ = rpc.InitGRPCClients()
	defer rpc.CloseGPRCClients()

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 提交代码
	submitRes, err := rpc.SubmitCli.Submit(ctx, &rpcSubmit.SubmitRequest{
		ProblemID: 2,
		UserID:    1,
		Code: []byte(`
		#include <iostream>
		using namespace std;
		int main() {
			int a, b;
			cin >> a >> b;
			cout << a + b << endl;
			return 0;
		}
		`),
		LangID: 2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("提交成功")
	fmt.Printf("submitRes.StatusCode: %v\n", submitRes.StatusCode)
	fmt.Printf("submitRes.SubmitID: %v\n", submitRes.SubmitID)

	// 获取结果
	result, err := rpc.SubmitCli.GetSubmitResult(ctx, &rpcSubmit.GetSubmitResultRequest{
		SubmitID: submitRes.SubmitID,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("获取结果成功")
	fmt.Printf("result.StatusCode: %v\n", result.StatusCode)
	fmt.Printf("result.Result: %v\n", result.Result)
}
