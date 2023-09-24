package judge

import (
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	rpcJudge "main/api/judge"
	"main/config"
	"main/internal/common/run"
	"main/internal/middleware/tracing"
	"main/internal/middleware/mq"
	"main/rpc"
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
	conf := config.ConfJudge

	// 初始化tracer
	tracer, closer := tracing.InitTracer(conf.Name)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// 连接题目模块RPC服务
	conn, err := rpc.InitProblemGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 启动MQ
	defer mq.RunJudgeMQ().Destroy()

	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcJudge.RegisterJudgeServiceServer(grpcServ, new(JudgeServer))
	})
}
