package judge

import (
	"flag"
	"fmt"
	"main/api/judge"
	"main/config"
	"main/discovery"
	"main/internal/middleware/mq"
	"main/internal/service/judge/handle"
	"main/rpc"
	"net"

	"google.golang.org/grpc"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
}

func Run() error {
	// 连接题目模块RPC服务
	conn, err := rpc.InitProblemGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 启动MQ
	mq.InitRabbitMQ()
	defer mq.DestroyRabbitMQ()

	// 读取grpc服务启动端口
	var port int
	flag.IntVar(&port, "p", config.ConfJudge.Port, "port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", config.ConfJudge.Host, port)

	fmt.Println("address: ", addr)

	// 注册etcd节点
	discovery.RegisterEtcdEndpoint(config.ConfJudge.Name, addr, config.ConfJudge.Version)

	// 创建并启动GRPC服务
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServ := grpc.NewServer()
	rpcJudge.RegisterJudgeServiceServer(grpcServ, new(handle.JudgeServer))
	if err := grpcServ.Serve(listen); err != nil {
		return err
	}
	return nil
}
