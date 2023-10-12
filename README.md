# Online Judge

## 项目文档📑

#### [后端开发日志](https://www.yuque.com/ahaostudy/ib4xuq/bug9zzziqc299wu7)

#### [API Document](https://apifox.com/apidoc/shared-cf30a21c-df5d-4034-92fc-b01f89189f50)

#### [QingdaoU/Judger](https://opensource.qduoj.com/#/judger/api)

## Judger 判题器

判题器基于 [QingdaoU沙箱](https://github.com/QingdaoU/Judger) 开发，仅支持在Linux环境下运行。


## 项目环境

- Golang
- MySQL
- Redis
- RabbitMQ
- ETCD
- Seccomp
- GCC、G++、JDK ...

#### Seccomp

```shell
sudo apt-get install libseccomp-dev
```

#### 各语言环境

例（GCC、G++、JDK8）：
```shell
sudo apt-get install gcc
sudo apt-get install g++
sudo apt-get install openjdk-8-jdk
...
```


## 项目启动

### 容器启动

先拉取项目到本地，使用docker-compose启动即可。

```shell
docker-compose up -d
```

### 本地启动

#### 更新配置文件

```shell
cp config/config.yaml.bak config/config.yaml
vim config/config.yaml
```
将配置文件中的路径更改为本地路径。

判题模块的exe处填写的是各编译器路径，如果不清楚路径可以使用 `which` 命令查找，如：`which gcc` 。

此外沙箱执行需要获取用户权限，如果在普通用户环境中请填写sudo的密码，在root用户环境中不需要填写。


#### 初始化项目依赖
```shell
go mod init main
go mod tidy
```

#### 启动项目
进入`cmd`目录，根据需要启动的模块分别启动。
```shell
go run cmd/judge/main.go
go run cmd/problem/main.go
go run cmd/submit/main.go
go run cmd/contest/main.go
go run cmd/user/main.go
go run cmd/chatgpt/main.go
go run cmd/gateway/main.go
```
