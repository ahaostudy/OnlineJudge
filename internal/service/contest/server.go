package contest

import (
	"main/api/contest"
	"main/config"
	"main/internal/common/run"

	"google.golang.org/grpc"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
}

type ContestServer struct {
	rpcContest.UnimplementedContestServiceServer
}

func Run() error {
	conf := config.ConfContest

	// 运行当前服务
	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcContest.RegisterContestServiceServer(grpcServ, new(ContestServer))
	})
}
