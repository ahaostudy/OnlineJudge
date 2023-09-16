package submit

import (
	rpcSubmit "main/api/submit"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
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
	conn, err := rpc.InitJudgeGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conf := config.ConfSubmit

	run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcSubmit.RegisterSubmitServiceServer(grpcServ, new(SubmitServer))
	})

	return nil
}
