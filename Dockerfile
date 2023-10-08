FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG C.UTF-8
ENV GOPATH=/go

WORKDIR $GOPATH/app
COPY . .

VOLUME ["/etc/oj/data", "/etc/oj/config"]

RUN apt-get clean && apt-get update && \
    apt-get install -y wget libseccomp-dev gcc g++ openjdk-8-jdk && \
    wget https://dl.google.com/go/go1.20.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz

# RUN apt-get install -y python3.8
ENV PATH="/usr/local/go/bin:${PATH}"

RUN go mod init main && \
    go mod tidy && \
    mkdir build && \
    go build -o build/service-judge cmd/judge/main.go && \
    go build -o build/service-problem cmd/problem/main.go && \
    go build -o build/service-submit cmd/submit/main.go && \
    go build -o build/service-user cmd/user/main.go && \
    go build -o build/service-chatgpt cmd/chatgpt/main.go && \
    go build -o build/service-gateway cmd/gateway/main.go && \
    mkdir -p /usr/lib/judger && \
    cp lib/libjudger.so /usr/lib/judger/libjudger.so && \
    mkdir /app && \
    cp build/* /app && \
    rm ./* -rf

EXPOSE 8080 9991 9992 9993 9994 9995 9996

ENTRYPOINT ./app/$SERVICE --cp=/etc/oj/config/config.yaml
