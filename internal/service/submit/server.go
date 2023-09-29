package submit

import (
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	rpcSubmit "main/api/submit"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
	"main/internal/middleware/mq"
	"main/internal/middleware/redis"
	"main/internal/middleware/tracing"
	"main/internal/service/submit/jobs"
	"main/rpc"
)

func init() {
	if err := data.InitMySQL(); err != nil {
		panic(err)
	}

	if err := redis.InitRedis(); err != nil {
		panic(err)
	}
}

type SubmitServer struct {
	rpcSubmit.UnimplementedSubmitServiceServer
}

func Run() error {
	conf := config.ConfSubmit

	// 初始化tracer
	tracer, closer := tracing.InitTracer(conf.Name)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// 连接需要的rpc服务
	connJudge, err := rpc.InitJudgeGRPC()
	if err != nil {
		panic(err)
	}
	defer connJudge.Close()
	connProblem, err := rpc.InitProblemGRPC()
	if err != nil {
		panic(err)
	}
	defer connProblem.Close()

	// 启动MQ
	defer mq.RunContestSubmitMQ().Destroy()

	// 启动定时任务
	jobs.RunSubmitJobs()

	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcSubmit.RegisterSubmitServiceServer(grpcServ, new(SubmitServer))
	})
}
