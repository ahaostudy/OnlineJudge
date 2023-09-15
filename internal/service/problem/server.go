package problem

import (
	"flag"
	"fmt"
	rpcProblem "main/api/problem"
	"main/config"
	"main/discovery"
	"main/internal/data"
	"net"

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
	// 读取grpc服务启动端口
	var port int
	flag.IntVar(&port, "p", config.ConfProblem.Port, "port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", config.ConfProblem.Host, port)

	fmt.Println("address: ", addr)

	// 注册etcd节点
	discovery.RegisterEtcdEndpoint(config.ConfProblem.Name, addr, config.ConfProblem.Version)

	// 创建并启动GRPC服务
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServ := grpc.NewServer()
	rpcProblem.RegisterProblemServiceServer(grpcServ, new(ProblemServer))
	if err := grpcServ.Serve(listen); err != nil {
		return err
	}
	return nil
}
