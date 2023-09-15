package private

import (
	"main/api/private"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
	"main/internal/middleware/mq"
	"main/rpc"

	"google.golang.org/grpc"
)

func init() {
	if err := data.InitMySQL(); err != nil {
		panic(err)
	}
}

type PrivateServer struct {
	rpcPrivate.UnimplementedPrivateServiceServer
}

func Run() error {
	conf := config.ConfPrivate

	// 初始化RabbitMQ连接
	mq.InitRabbitMQ()
	defer mq.DestroyRabbitMQ()

	// 初始化Problem服务连接
	conn, err := rpc.InitProblemGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 运行当前服务
	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcPrivate.RegisterPrivateServiceServer(grpcServ, new(PrivateServer))
	})
}
