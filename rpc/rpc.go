package rpc

import (
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

var conns []*grpc.ClientConn

type run struct {
	err error
}

// 运行代码，不出现错误时才运行
func (r *run) Run(fun func() (*grpc.ClientConn, error)) {
	if r.err != nil {
		return
	}
	conn, err := fun()
	r.err = err
	conns = append(conns, conn)
}

func (r *run) Error() error {
	return r.err
}

// 初始化所有Client
func InitGRPCClients() error {
	r := new(run)
	r.Run(InitJudgeGRPC)
	r.Run(InitProblemGRPC)
	r.Run(InitSubmitGRPC)
	r.Run(InitPrivateGRPC)
	r.Run(InitChatGPTGRPC)
	r.Run(InitContestGRPC)

	return r.Error()
}

// 关闭所有Client
func CloseGPRCClients() {
	for _, conn := range conns {
		_ = conn.Close()
	}
}

// 创建GRPC客户端
func NewGRPCClient(addr, name string) (*grpc.ClientConn, error) {
	// etcd
	etcdCli, err := clientv3.NewFromURL(addr)
	if err != nil {
		panic(err)
	}
	etcdResolver, err := resolver.NewBuilder(etcdCli)
	if err != nil {
		return nil, err
	}

	// dial
	conn, err := grpc.Dial("etcd:///"+name,
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
	)
	if err != nil {
		return conn, err
	}

	return conn, nil
}
