package test

import (
	"context"
	"fmt"
	rpcProblem "main/api/problem"
	"main/rpc"
	"testing"
)

func TestCreateTestcase(t *testing.T) {
	_ = rpc.InitGRPCClients()
	defer rpc.CloseGPRCClients()

	res, err := rpc.ProblemCli.CreateTestcase(context.Background(), &rpcProblem.CreateTestcaseRequest{
		ProblemId: 4,
		Input:     []byte("23 25"),
		Output:    []byte("48"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
}
