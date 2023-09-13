package private

import (
	"main/api/private"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
	"main/internal/service/private/handle"

	"google.golang.org/grpc"
)

func init() {
	if err := data.InitMySQL(); err != nil {
		panic(err)
	}
}

func Run() error {
	conf := config.ConfPrivate

	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcPrivate.RegisterPrivateServiceServer(grpcServ, new(handle.PrivateServer))
	})
}
