package judge

import (
	"context"
	"main/api/judge"
	"main/internal/common/code"
	"main/internal/common/build"
	"main/internal/middleware/mq"
)

func (JudgeServer) GetResult(ctx context.Context, req *rpcJudge.GetResultRequest) (resp *rpcJudge.GetResultResponse, _ error) {
	resp = new(rpcJudge.GetResultResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 读取管道获取结果并关闭管道
	ch, ok := mq.GetResultChan(req.GetJudgeID())
	if !ok || ch == nil {
		resp.StatusCode = code.CodeSubmitNotFound.Code()
		return
	}
	res := <-ch
	if done, ok := mq.GetDoneChan(req.GetJudgeID()); ok && done != nil {
		done <- struct{}{}
	}

	// 判断运行是否错误，并复制Result
	result, err := build.BuildResult(&res.Result)
	if err == nil {
		resp.Result = result
		resp.StatusCode = code.CodeSuccess.Code()
	}

	return
}
