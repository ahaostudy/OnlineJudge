package testcase

import (
	rpcTestcase "main/api/testcase"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
	"main/internal/service/testcase/handle"

	"google.golang.org/grpc"
)

func init() {
	if err := data.InitMySQL(); err != nil {
		panic(err)
	}
}

func Run() error {
	conf := config.ConfTestcase

	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcTestcase.RegisterTestcaseServiceServer(grpcServ, new(handle.TestcaseServer))
	})
}
