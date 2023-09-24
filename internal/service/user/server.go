package user

import (
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	rpcUser "main/api/user"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
	"main/internal/middleware/tracing"
	"main/internal/middleware/redis"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := data.InitMySQL(); err != nil {
		panic(err)
	}

	if err := redis.InitRedis(); err != nil {
		panic(err)
	}
}

type UserServer struct {
	rpcUser.UnimplementedUserServiceServer
}

func Run() error {
	conf := config.ConfUser

	// 初始化tracer
	tracer, closer := tracing.InitTracer(conf.Name)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcUser.RegisterUserServiceServer(grpcServ, new(UserServer))
	})
}
