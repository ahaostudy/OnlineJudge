package problem

import (
	rpcProblem "main/api/problem"
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

type ProblemServer struct {
	rpcProblem.UnimplementedProblemServiceServer
}

func Run() error {
	conf := config.ConfProblem

	// 连接contest服务
	conn, err := rpc.InitContestGRPC()
	if err != nil {
		return err
	}
	defer conn.Close()

	// 启动服务
	return run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcProblem.RegisterProblemServiceServer(grpcServ, new(ProblemServer))
	})
}
