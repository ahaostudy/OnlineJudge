package judge

import (
	"context"
	"fmt"
	"main/api/judge"
	rpcProblem "main/api/problem"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/middleware/mq"
	"main/internal/service/judge/pkg/compiler"
	"main/rpc"
	"time"

	"github.com/google/uuid"
)

// 判题服务
func (JudgeServer) Judge(ctx context.Context, req *rpcJudge.JudgeRequest) (resp *rpcJudge.JudgeResponse, _ error) {
	resp = new(rpcJudge.JudgeResponse)
	resp.StatusCode = common.CodeServerBusy.Code()
	code, langID, problemID := req.GetCode(), req.GetLangID(), req.GetProblemID()

	// 将代码文件上传
	path, err := compiler.SaveCode(code, int(langID))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	resp.CodePath = path

	// 根据题目ID获取题目信息
	prob, err := rpc.ProblemCli.GetProblem(context.Background(), &rpcProblem.GetProblemRequest{ProblemId: problemID})
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
	mq.ResultChan.Store(resp.JudgeID, make(chan mq.JudgeResponse))
	mq.DoneChan.Store(resp.JudgeID, make(chan struct{}))
	// 如果30s后管道内容仍未被接收将自动清除
	go func() {
		done, ok := mq.GetDoneChan(resp.GetJudgeID())
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
		if ch, ok := mq.GetResultChan(resp.GetJudgeID()); ok {
			close(ch)
		}
		close(done)
		mq.DoneChan.Delete(resp.GetJudgeID())
		mq.ResultChan.Delete(resp.GetJudgeID())
	}()

	// 将请求放入MQ
	msg, err := mq.GenerateJudgeMQMsg(resp.JudgeID, path, langID, problem)
	if err != nil {
		return
	}
	if mq.RMQJudge.Publish(msg) != nil {
		return
	}

	resp.StatusCode = common.CodeSuccess.Code()
	return
}
