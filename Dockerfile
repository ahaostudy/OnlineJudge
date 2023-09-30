FROM golang:1.20

ENV GOPROXY https://goproxy.cn
ENV SERVER gateway

WORKDIR $GOPATH/src/OnlineJudge

COPY . .

RUN go build -o judge ./cmd/judge/main.go
RUN go build -o problem ./cmd/problem/main.go
RUN go build -o submit ./cmd/submit/main.go
RUN go build -o contest ./cmd/contest/main.go
RUN go build -o user ./cmd/user/main.go
RUN go build -o chatgpt ./cmd/chatgpt/main.go
RUN go build -o gateway ./cmd/gateway/main.go

EXPOSE 8080 9991 9992 9993 9994 9995 9996

ENTRYPOINT ./$SERVER
