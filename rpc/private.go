package rpc

import (
	rpcPrivate "main/api/private"
	"main/config"

	"google.golang.org/grpc"
)

var PrivateCli rpcPrivate.PrivateServiceClient

func InitPrivateGRPC() (*grpc.ClientConn, error) {
	conn, err := NewGRPCClient(config.ConfEtcd.Addr, config.ConfPrivate.Name)
	if err != nil {
		return conn, err
	}
	PrivateCli = rpcPrivate.NewPrivateServiceClient(conn)

	return conn, nil
}
