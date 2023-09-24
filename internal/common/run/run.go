package run

import (
	"flag"
	"fmt"
	"net"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"

	"main/config"
	"main/discovery"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

}

// Run 运行微服务，传入一个注册函数
func Run(host string, port int, name, version string, register func(grpcServ *grpc.Server)) error {
	// 读取grpc服务启动端口
	flag.IntVar(&port, "p", port, "port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", host, port)

	fmt.Println("address: ", addr)

	// 注册etcd节点
	discovery.RegisterEtcdEndpoint(name, addr, version)

	// 创建并启动GRPC服务
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServ := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpc_opentracing.UnaryServerInterceptor()),
	)

	// 注册服务节点
	register(grpcServ)

	if err := grpcServ.Serve(listen); err != nil {
		return err
	}
	return nil
}
