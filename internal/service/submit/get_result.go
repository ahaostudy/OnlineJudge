package submit

import (
	"context"
	rpcJudge "main/api/judge"
	rpcSubmit "main/api/submit"
	"main/internal/common/code"
	"main/internal/common/build"
	"main/internal/data/model"
	"main/internal/data/repository"
	"main/internal/middleware/redis"
	"main/rpc"
)

func (SubmitServer) GetSubmitResult(ctx context.Context, req *rpcSubmit.GetSubmitResultRequest) (resp *rpcSubmit.GetSubmitResultResponse, _ error) {
	resp = new(rpcSubmit.GetSubmitResultResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取判题id，并将提交缓存删除
	judgeID, err := redis.Rdb.Get(ctx, redis.GenerateSubmitKey(req.GetSubmitID())).Result()
	if err != nil {
		resp.StatusCode = code.CodeSubmitNotFound.Code()
		return
	}
	go redis.Del(redis.GenerateSubmitKey(req.GetSubmitID()))

	// 获取运行结果
	res, err := rpc.JudgeCli.GetResult(ctx, &rpcJudge.GetResultRequest{
		JudgeID: judgeID,
	})
	if err != nil {
		return
	}
	if res.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = res.GetStatusCode()
		return
	}

	// 将结果更新到数据库
	result := &model.Submit{
		Status: res.GetResult().GetStatus(),
		Time:   res.GetResult().GetTime(),
		Memory: res.GetResult().GetMemory(),
	}
	if err := repository.UpdateSubmit(req.GetSubmitID(), result); err != nil {
		return
	}
	// 将结果转换为rpc响应
	resp.Result, err = build.BuildSubmit(result)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
