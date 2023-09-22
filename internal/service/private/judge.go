package private

import (
	"context"
	"main/api/private"
	"main/api/problem"
	status "main/internal/common/code"
	"main/internal/common/build"
	"main/internal/middleware/mq"
	"main/internal/service/judge/pkg/compiler"
	"main/rpc"
	"time"

	"github.com/google/uuid"
)

// Judge 判题服务
func (PrivateServer) Judge(ctx context.Context, req *rpcPrivate.JudgeRequest) (resp *rpcPrivate.JudgeResponse, _ error) {
	resp = new(rpcPrivate.JudgeResponse)
	resp.StatusCode = status.CodeServerBusy.Code()
	code, langID, problemID := req.GetCode(), req.GetLangID(), req.GetProblemID()

	// 将代码文件上传
	path, err := compiler.SaveCode(code, int(langID))
	if err != nil {
		return
	}

	// 根据题目ID获取题目信息
	prob, err := rpc.ProblemCli.GetProblem(context.Background(), &rpcProblem.GetProblemRequest{ProblemID: problemID})
	if err != nil {
		return
	}

	// 将结果转换为模型对象
	problem, err := build.UnBuildProblem(prob.GetProblem())
	if err != nil {
		return
	}

	resp.JudgeID = uuid.NewString()

	// 为当前判题开辟一个管道
	mq.PrivateResultChan.Store(resp.JudgeID, make(chan mq.PrivateJudgeResponse))
	mq.PrivateDoneChan.Store(resp.JudgeID, make(chan struct{}))
	// 如果30s后管道内容仍未被接收将自动清除
	go func() {
		done, ok := mq.GetPrivateDoneChan(resp.GetJudgeID())
		if !ok || done == nil {
			return
		}

		timer := time.NewTimer(30 * time.Second)
		select {
		case <-done:
			break
		case <-timer.C:
			break
		}

		timer.Stop()
		if ch, ok := mq.GetPrivateResultChan(resp.GetJudgeID()); ok {
			close(ch)
		}
		close(done)
		mq.PrivateDoneChan.Delete(resp.GetJudgeID())
		mq.PrivateResultChan.Delete(resp.GetJudgeID())
	}()

	// 将请求放入MQ
	msg, err := mq.GeneratePrivateJudgeMQMsg(resp.JudgeID, path, langID, problem)
	if err != nil {
		return
	}
	if mq.RMQPrivate.Publish(msg) != nil {
		return
	}

	resp.StatusCode = status.CodeSuccess.Code()
	return
}

// GetResult 获取运行结果
func (PrivateServer) GetResult(ctx context.Context, req *rpcPrivate.GetResultRequest) (resp *rpcPrivate.GetResultResponse, _ error) {
	resp = new(rpcPrivate.GetResultResponse)
	resp.StatusCode = status.CodeServerBusy.Code()

	// 读取管道获取结果并关闭管道
	ch, ok := mq.GetPrivateResultChan(req.GetJudgeID())
	if !ok || ch == nil {
		resp.StatusCode = status.CodeSubmitNotFound.Code()
		return
	}
	res := <-ch
	if done, ok := mq.GetPrivateDoneChan(req.GetJudgeID()); ok && done != nil {
		done <- struct{}{}
	}

	// 判断运行是否错误，并复制Result
	resp.Result = new(rpcPrivate.JudgeResult)
	if new(build.Builder).Build(res.Result, resp.Result).Error() == nil {
		resp.StatusCode = status.CodeSuccess.Code()
	}

	return
}
