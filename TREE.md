### 项目结构参考

> 来自ChatGPT

```
project/
├── api/                  # 存放 .proto 文件和生成的 gRPC 代码
│   ├── service1.proto    # 第一个 gRPC 服务定义文件
│   ├── service2.proto    # 第二个 gRPC 服务定义文件
│   └── ...
├── cmd/
│   ├── service1/         # 第一个服务的启动代码
│   │   └── main.go
│   ├── service2/         # 第二个服务的启动代码
│   │   └── main.go
│   └── gateway/          # 网关层的启动代码
│       └── main.go
├── config/
│   ├── service1.yaml     # 第一个服务的配置文件
│   ├── service2.yaml     # 第二个服务的配置文件
│   └── ...
├── discovery/
│   ├── etcd.go           # etcd注册和发现逻辑
│   └── register.go       # 服务注册接口
├── internal/
│   ├── data/             # 数据访问层
│   │   ├── models/       # 数据模型定义
│   │   ├── repository/   # 数据库交互代码
│   │   ├── migrations/   # 数据库迁移脚本
│   │   └── db.go         # 数据库连接初始化
│   ├── server/           # 服务实现
│   │   ├── service1_handler.go   # 第一个服务的业务逻辑实现
│   │   ├── service2_handler.go   # 第二个服务的业务逻辑实现
│   │   └── ...
│   ├── middleware/
│   └── ...
├── rpc/
│   ├── service1/
│   │   └── client.go
│   ├── service2/
│   │   └── client.go
│   └── ...
├── scripts/
├── third_party/
├── Dockerfile
├── Makefile
├── README.md
├── go.mod
└── go.sum
```