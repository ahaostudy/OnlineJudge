package testcase

import (
	rpcTestcase "main/api/testcase"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"

	"google.golang.org/grpc"
)

func init() {
	if err := data.InitMySQL(); err != nil {
		panic(err)
	}
}

type TestcaseServer struct {
	rpcTestcase.UnimplementedTestcaseServiceServer
}

func Run() error {
	conf := config.ConfTestcase

	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcTestcase.RegisterTestcaseServiceServer(grpcServ, new(TestcaseServer))
	})
}
