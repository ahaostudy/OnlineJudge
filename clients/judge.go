package clients

import (
	"fmt"
	"main/config"
	rpcJudge "main/rpc/judge"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

var JudgeCli rpcJudge.JudgeServiceClient

func InitJudgeGRPC() (*grpc.ClientConn, error) {
	// etcd
	etcdCli, err := clientv3.NewFromURL(config.ConfEtcd.Addr)
	if err != nil {
		panic(err)
	}
	etcdResolver, err := resolver.NewBuilder(etcdCli)
	if err != nil {
		return nil, err
	}

	// dial
	conn, err := grpc.Dial("etcd:///"+config.ConfJudge.Name,
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
	)
	if err != nil {
		return conn, err
	}

	// new client
	JudgeCli = rpcJudge.NewJudgeServiceClient(conn)

	return conn, nil
}
