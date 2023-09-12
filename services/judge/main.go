package main

import (
	"flag"
	"fmt"
	"main/config"
	"main/discovery"
	rpcJudge "main/rpc/judge"
	"main/services/judge/handle"
	"net"

	"google.golang.org/grpc"
)

func RunGRPCServer() error {
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

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
}

func main() {
	// 运行一个判题器
	go handle.Judger()

	if err := RunGRPCServer(); err != nil {
		panic(err)
	}
}
