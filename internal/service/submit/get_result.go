package submit

import (
	"context"
	"encoding/json"
	rpcJudge "main/api/judge"
	rpcSubmit "main/api/submit"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/middleware/mq"
	"main/internal/middleware/redis"
	status "main/internal/service/judge/pkg/code"
)

func (SubmitServer) GetSubmitResult(ctx context.Context, req *rpcSubmit.GetSubmitResultRequest) (resp *rpcSubmit.GetSubmitResultResponse, _ error) {
	resp = new(rpcSubmit.GetSubmitResultResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取判题ID
	judgeID, err := redis.Rdb.Get(ctx, redis.GenerateSubmitKey(req.GetSubmitID())).Result()
	if err != nil {
		return
	}
	// 获取不到，说明提交的key不存在或已过期
	if len(judgeID) == 0 {
		resp.StatusCode = code.CodeSubmitNotFound.Code()
		return
	}

	// 获取判题结果
	res, err := redis.Rdb.Get(ctx, redis.GenerateJudgeKey(judgeID)).Bytes()
	if err != nil {
		return
	}
	// 运行中，未获取到结果
	if len(res) == 0 {
		resp.Result = &rpcJudge.JudgeResult{Status: int64(status.StatusRunning)}
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	// 将结果反序列化
	result := new(mq.JudgeResponse)
	if err := json.Unmarshal(res, result); err != nil {
		return
	}
	if result.Error != nil {
		resp.Result = &rpcJudge.JudgeResult{Status: int64(status.StatusServerFailed)}
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	// 将结果转换为rpc响应
	resp.Result = new(rpcJudge.JudgeResult)
	if new(build.Builder).Build(result.Result, resp.Result).Error() != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
