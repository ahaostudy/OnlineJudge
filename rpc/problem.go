package rpc

import (
	rpcProblem "main/api/problem"
	"main/config"

	"google.golang.org/grpc"
)

var ProblemCli rpcProblem.ProblemServiceClient

func InitProblemGRPC() (*grpc.ClientConn, error) {
	conn, err := NewGRPCClient(config.ConfEtcd.Addr, config.ConfProblem.Name)
	if err != nil {
		return conn, err
	}
	ProblemCli = rpcProblem.NewProblemServiceClient(conn)

	return conn, nil
}
