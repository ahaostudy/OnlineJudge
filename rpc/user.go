package rpc

import (
	"google.golang.org/grpc"

	rpcUser "main/api/user"
	"main/config"
)

var UserCli rpcUser.UserServiceClient

func InitUserGRPC() (*grpc.ClientConn, error) {
	conn, err := NewGRPCClient(config.ConfEtcd.Addr, config.ConfUser.Name)
	if err != nil {
		return conn, err
	}
	UserCli = rpcUser.NewUserServiceClient(conn)

	return conn, nil
}
