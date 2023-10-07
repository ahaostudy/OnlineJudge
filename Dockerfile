FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG C.UTF-8

VOLUME ["/data"]

RUN sed -i 's@http://archive.ubuntu.com/ubuntu/@http://mirrors.aliyun.com/ubuntu/@g' /etc/apt/sources.list
RUN apt-get clean && apt-get update
RUN apt-get install -y wget

RUN apt-get install libseccomp-dev

RUN apt-get install -y gcc
RUN apt-get install -y g++
RUN apt-get install -y openjdk-8-jdk
# RUN apt-get install -y python3.8

RUN wget https://dl.google.com/go/go1.20.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPROXY https://goproxy.cn

WORKDIR $GOPATH/src/OnlineJudge
COPY . .
RUN mkdir -p /usr/lib/judger && cp lib/libjudger.so /usr/lib/judger/libjudger.so

RUN go mod tidy
RUN go build -o service-judge cmd/judge/main.go
RUN go build -o service-problem cmd/problem/main.go
RUN go build -o service-submit cmd/submit/main.go
RUN go build -o service-user cmd/user/main.go
RUN go build -o service-chatgpt cmd/chatgpt/main.go
RUN go build -o service-gateway cmd/gateway/main.go

EXPOSE 8080 9991 9992 9993 9994 9995 9996

ENTRYPOINT ./$SERVICE
