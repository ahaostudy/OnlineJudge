package rpc

import (
	"main/api/chatgpt"
	"main/config"

	"google.golang.org/grpc"
)

var ChatGPTCli rpcChatGPT.ChatGPTServiceClient

func InitChatGPTGRPC() (*grpc.ClientConn, error) {
	conn, err := NewGRPCClient(config.ConfEtcd.Addr, config.ConfChatGPT.Name)
	if err != nil {
		return conn, err
	}
	ChatGPTCli = rpcChatGPT.NewChatGPTServiceClient(conn)

	return conn, nil
}
