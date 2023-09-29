package problem

import (
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	rpcProblem "main/api/problem"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
	"main/internal/middleware/tracing"
	"main/rpc"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := data.InitMySQL(); err != nil {
		panic(err)
	}
}

type ProblemServer struct {
	rpcProblem.UnimplementedProblemServiceServer
}

func Run() error {
	conf := config.ConfProblem

	// 初始化tracer
	tracer, closer := tracing.InitTracer(conf.Name)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// 连接contest服务
	connContest, err := rpc.InitContestGRPC()
	if err != nil {
		return err
	}
	defer connContest.Close()
	// 连接submit服务
	connSubmit, err := rpc.InitSubmitGRPC()
	if err != nil {
		return err
	}
	defer connSubmit.Close()

	// 启动服务
	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcProblem.RegisterProblemServiceServer(grpcServ, new(ProblemServer))
	})
}
