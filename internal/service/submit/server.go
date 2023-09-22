package submit

import (
	rpcSubmit "main/api/submit"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
	"main/internal/middleware/mq"
	"main/internal/middleware/redis"
	"main/rpc"

	"google.golang.org/grpc"
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

	conf := config.ConfSubmit

	run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcSubmit.RegisterSubmitServiceServer(grpcServ, new(SubmitServer))
	})

	return nil
}
