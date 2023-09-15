package rpc

import (
	rpcSubmit "main/api/submit"
	"main/config"

	"google.golang.org/grpc"
)

var SubmitCli rpcSubmit.SubmitServiceClient

func InitSubmitGRPC() (*grpc.ClientConn, error) {
	conn, err := NewGRPCClient(config.ConfEtcd.Addr, config.ConfSubmit.Name)
	if err != nil {
		return conn, err
	}
	SubmitCli = rpcSubmit.NewSubmitServiceClient(conn)

	return conn, nil
}
