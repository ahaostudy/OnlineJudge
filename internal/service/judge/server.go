package judge

import (
	"main/api/judge"
	"main/config"
	"main/internal/common/run"
	"main/internal/middleware/mq"
	"main/rpc"

	"google.golang.org/grpc"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
}

type JudgeServer struct {
	rpcJudge.UnimplementedJudgeServiceServer
}

func Run() error {
	// 连接题目模块RPC服务
	conn, err := rpc.InitProblemGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 启动MQ
	defer mq.RunJudgeMQ().Destroy()

	conf := config.ConfJudge
	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcJudge.RegisterJudgeServiceServer(grpcServ, new(JudgeServer))
	})
}
