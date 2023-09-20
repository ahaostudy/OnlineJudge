package contest

import (
	"main/api/contest"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
	"main/rpc"

	"google.golang.org/grpc"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := data.InitMySQL(); err != nil {
		panic(err)
	}
}

type ContestServer struct {
	rpcContest.UnimplementedContestServiceServer
}

func Run() error {
	conf := config.ConfContest

	// 连接problem服务
	conn, err := rpc.InitProblemGRPC()
	if err != nil {
		return err
	}
	defer conn.Close()

	// 运行当前服务
	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcContest.RegisterContestServiceServer(grpcServ, new(ContestServer))
	})
}
