package judge

import (
	"context"
	"main/api/judge"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/middleware/mq"
)

func (JudgeServer) GetResult(ctx context.Context, req *rpcJudge.GetResultRequest) (resp *rpcJudge.GetResultResponse, _ error) {
	resp = new(rpcJudge.GetResultResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 读取管道获取结果并关闭管道
	if _, ok := mq.ResultChan[req.GetJudgeID()]; !ok {
		resp.StatusCode = common.CodeSubmitNotFound.Code()
		return
	}
	res := <-mq.ResultChan[req.GetJudgeID()]
	close(mq.ResultChan[req.JudgeID])
	delete(mq.ResultChan, req.JudgeID)

	// 判断运行是否错误，并复制Result
	result, err := build.BuildResult(&res.Result)
	if err == nil {
		resp.Result = result
		resp.StatusCode = common.CodeSuccess.Code()
	}

	return
}
