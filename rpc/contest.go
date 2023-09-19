package rpc

import (
	"main/api/contest"
	"main/config"

	"google.golang.org/grpc"
)

var ContestCli rpcContest.ContestServiceClient

func InitContestGRPC() (*grpc.ClientConn, error) {
	conn, err := NewGRPCClient(config.ConfEtcd.Addr, config.ConfContest.Name)
	if err != nil {
		return conn, err
	}
	ContestCli = rpcContest.NewContestServiceClient(conn)
	return conn, nil
}