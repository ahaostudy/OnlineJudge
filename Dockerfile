FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG C.UTF-8

ENV GOPATH=/go
ENV PATH="/usr/local/go/bin:${PATH}"
WORKDIR $GOPATH/app

COPY . .

VOLUME ["/etc/oj/data", "/etc/oj/config"]

# 更新软件包列表并安装依赖
RUN apt-get update && \
    apt-get install -y \
    curl \
    wget \
    libseccomp-dev \
    gcc \
    g++ \
    openjdk-8-jdk && \
    wget https://dl.google.com/go/go1.20.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz && \
    rm go1.20.linux-amd64.tar.gz

# Go 依赖和构建
RUN go mod tidy && \
    mkdir build && \
    go build -o build/service-judge cmd/judge/main.go && \
    go build -o build/service-problem cmd/problem/main.go && \
    go build -o build/service-submit cmd/submit/main.go && \
    go build -o build/service-user cmd/user/main.go && \
    go build -o build/service-chatgpt cmd/chatgpt/main.go && \
    go build -o build/service-gateway cmd/gateway/main.go

# 安装其他依赖和复制文件
RUN mkdir -p /app /usr/lib/judger && \
    cp -r script build/* /app && \
    cp lib/libjudger.so /usr/lib/judger/libjudger.so

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf ./*
