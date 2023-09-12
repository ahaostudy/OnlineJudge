package discovery

import (
	"context"
	"errors"
	"fmt"
	"log"
	"main/config"
	"os"
	"os/signal"
	"syscall"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type Etcd struct {
	addr          string                                  // etcd地址
	ttl           int64                                   // 每次租约的时长
	srvInfo       *Server                                 // grpc服务信息
	cli           *clientv3.Client                        // etcd cli
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse // 保活管道（租满时，租约客户端将继续向 etcd 服务器发送保持活动请求）
	closeChan     chan struct{}                           // 服务关闭的管道
}

// RegisterEtcdEndpoint
// 使用config中的信息将GRPC服务注册成ETCD节点
func RegisterEtcdEndpoint(name, addr, version string) {
	// 创建etcd对象
	etcd, err := NewEtcd(config.ConfEtcd.Addr, int64(config.ConfEtcd.Ttl), &Server{
		Name:    name,
		Addr:    addr,
		Version: version,
	})
	if err != nil {
		panic(err)
	}

	// 注册etcd
	closeChan, err := etcd.Register()
	if err != nil {
		panic(err)
	}

	// 监听关闭信号注销etcd
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		select {
		case <-ch:
			closeChan <- struct{}{}
			return
		}
	}()
}

// NewEtcd 创建etcd对象
func NewEtcd(addr string, ttl int64, srvInfo *Server) (*Etcd, error) {
	cli, err := clientv3.NewFromURL(addr)
	if err != nil {
		return nil, err
	}
	return &Etcd{addr: addr, ttl: ttl, srvInfo: srvInfo, cli: cli}, nil
}

// Register 注册服务，返回关闭管道，外部往管道传入值时注销服务
func (e *Etcd) Register() (chan<- struct{}, error) {
	if e.cli == nil {
		return nil, errors.New("the cli does not exist")
	}
	if err := e.register(); err != nil {
		return nil, err
	}
	e.closeChan = make(chan struct{})
	go e.keepAlive()
	return e.closeChan, nil
}

// etcd注册具体逻辑
func (e *Etcd) register() (err error) {
	// 申请租约
	lease, err := e.cli.Grant(context.TODO(), e.ttl)
	if err != nil {
		return
	}

	// 创建一个节点管理对象em
	em, err := endpoints.NewManager(e.cli, e.srvInfo.target())
	if err != nil {
		return
	}

	// 为当前服务创建节点
	err = em.AddEndpoint(context.TODO(), e.srvInfo.key(), endpoints.Endpoint{Addr: e.srvInfo.Addr}, clientv3.WithLease(lease.ID))
	if err != nil {
		return
	}

	// 自动续约
	e.keepAliveChan, err = e.cli.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return
	}

	return
}

// 注销服务
func (e *Etcd) unRegister() (err error) {
	if e.cli == nil {
		return errors.New("the cli does not exist")
	}

	// 创建一个节点管理对象em
	em, err := endpoints.NewManager(e.cli, e.srvInfo.target())
	if err != nil {
		return
	}

	// 删除当前服务节点
	err = em.DeleteEndpoint(context.TODO(), e.srvInfo.key())

	return

}

// 保持服务连接
func (e *Etcd) keepAlive() {
	for {
		select {
		// 接收到 close channel 的信号时注销服务
		case <-e.closeChan:
			if err := e.unRegister(); err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		case res := <-e.keepAliveChan:
			if res == nil {
				fmt.Println("keep clive log ... ")
				err := e.register()
				if err != nil {
					log.Println(err.Error())
					return
				}
			}
		}
	}
}
