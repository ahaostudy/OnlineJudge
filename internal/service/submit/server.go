package submit

import (
	"context"
	rpcSubmit "main/api/submit"
	"main/config"
	"main/internal/common/run"
	"main/internal/data"
	"main/internal/middleware/redis"
	"main/rpc"

	"google.golang.org/grpc"
)

func init() {
	if err := data.InitMySQL(); err != nil {
		panic(err)
	}

	if err := redis.InitRedis(); err != nil {
		panic(err)
	}
}

type SubmitServer struct {
	rpcSubmit.UnimplementedSubmitServiceServer
}

func Run() error {
	conn, err := rpc.InitJudgeGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conf := config.ConfSubmit

	run.Run(conf.Host, conf.Port, conf.Name, conf.Version, func(grpcServ *grpc.Server) {
		rpcSubmit.RegisterSubmitServiceServer(grpcServ, new(SubmitServer))
	})

	return nil
}

func (SubmitServer) Debug(_ context.Context, req *rpcSubmit.DebugReqeust) (resp *rpcSubmit.DebugResponse, _ error) {
	return
}

func (SubmitServer) GetSubmit(_ context.Context, req *rpcSubmit.GetSubmitRequest) (resp *rpcSubmit.GetSubmitResponse, _ error) {
	return
}

func (SubmitServer) DeleteSubmit(_ context.Context, req *rpcSubmit.DeleteSubmitRequest) (resp *rpcSubmit.DeleteSubmitResponse, _ error) {
	return
}
