package test

import (
	"context"
	"fmt"
	rpcTestcase "main/api/testcase"
	"main/rpc"
	"testing"
)

func TestCreateTestcase(t *testing.T) {
	_ = rpc.InitGRPCClients()
	defer rpc.CloseGPRCClients()

	res, err := rpc.TestcaseCli.CreateTestcase(context.Background(), &rpcTestcase.CreateTestcaseRequest{
		ProblemID: 4,
		Input:     []byte("23 25"),
		Output:    []byte("48   "),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
}
