package submit

import (
	"context"

	rpcJudge "main/api/judge"
	rpcSubmit "main/api/submit"
	"main/internal/common/code"
	"main/rpc"
)

// Debug 是一个处理调试请求的函数。
// 它接收一个上下文和一个 DebugReqeust 作为输入，并返回一个 DebugResponse 和一个错误。
func (SubmitServer) Debug(ctx context.Context, req *rpcSubmit.DebugReqeust) (resp *rpcSubmit.DebugResponse, _ error) {
	// 初始化响应对象
	resp = new(rpcSubmit.DebugResponse)

	// 将默认状态码设置为 CodeServerBusy
	resp.StatusCode = code.CodeServerBusy.Code()

	// 使用给定的代码、输入和语言ID调用 JudgeCli 的 Debug 方法
	res, err := rpc.JudgeCli.Debug(ctx, &rpcJudge.DebugRequest{
		Code:   req.GetCode(),
		Input:  req.GetInput(),
		LangID: req.GetLangID(),
	})
	if err != nil {
		return
	}

	// 将响应结果和状态码设置为来自 JudgeCli 响应的值
	resp.Result = res.Result
	resp.StatusCode = res.StatusCode

	return
}
