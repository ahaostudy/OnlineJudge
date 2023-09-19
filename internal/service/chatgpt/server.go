package chatgpt

import (
	"main/api/chatgpt"
	"main/config"
	"main/internal/common/run"

	"google.golang.org/grpc"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
}

type ChatGPTServer struct {
	rpcChatGPT.UnimplementedChatGPTServiceServer
}

func Run() error {
	conf := config.ConfChatGPT
	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcChatGPT.RegisterChatGPTServiceServer(grpcServ, new(ChatGPTServer))
	})
}
