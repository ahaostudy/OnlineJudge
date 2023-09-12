package rpc

import (
	"main/api/judge"
	"main/config"

	"google.golang.org/grpc"
)

var JudgeCli rpcJudge.JudgeServiceClient

func InitJudgeGRPC() (*grpc.ClientConn, error) {
	conn, err := NewGRPCClient(config.ConfEtcd.Addr, config.ConfJudge.Name)
	if err != nil {
		return conn, err
	}
	JudgeCli = rpcJudge.NewJudgeServiceClient(conn)

	return conn, nil
}
