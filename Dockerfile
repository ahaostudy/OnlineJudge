FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG C.UTF-8

WORKDIR $GOPATH/app
COPY . .

VOLUME ["/etc/oj/data", "/etc/oj/config"]

RUN apt-get clean && \
    apt-get update && \
    apt-get install -y gcc g++ openjdk-8-jdk libseccomp-dev wget && \
    wget https://dl.google.com/go/go1.20.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz

# RUN apt-get install -y python3.8

ENV PATH="/usr/local/go/bin:${PATH}"

RUN go mod tidy && \
    mkdir /app && \
    go build -o /app/service-judge cmd/judge/main.go && \
    go build -o /app/service-problem cmd/problem/main.go && \
    go build -o /app/service-submit cmd/submit/main.go && \
    go build -o /app/service-user cmd/user/main.go && \
    go build -o /app/service-chatgpt cmd/chatgpt/main.go && \
    go build -o /app/service-gateway cmd/gateway/main.go && \
    mkdir -p /usr/lib/judger && cp lib/libjudger.so /usr/lib/judger/libjudger.so && \
    rm ./ -rf

EXPOSE 8080 9991 9992 9993 9994 9995 9996

ENTRYPOINT /app/$SERVICE --cp=/etc/oj/config/config.yaml
