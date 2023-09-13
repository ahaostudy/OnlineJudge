package rpc

import (
	"main/api/testcase"
	"main/config"

	"google.golang.org/grpc"
)

var TestcaseCli rpcTestcase.TestcaseServiceClient

func InitTestcaseGRPC() (*grpc.ClientConn, error) {
	conn, err := NewGRPCClient(config.ConfEtcd.Addr, config.ConfTestcase.Name)
	if err != nil {
		return conn, err
	}
	TestcaseCli = rpcTestcase.NewTestcaseServiceClient(conn)

	return conn, nil
}
