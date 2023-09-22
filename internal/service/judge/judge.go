package judge

import (
	"context"
	"fmt"
	"main/api/judge"
	rpcProblem "main/api/problem"
	"main/internal/common/build"
	status "main/internal/common/code"
	"main/internal/middleware/mq"
	"main/internal/service/judge/pkg/compiler"
	"main/rpc"

	"github.com/google/uuid"
)

// 判题服务
func (JudgeServer) Judge(ctx context.Context, req *rpcJudge.JudgeRequest) (resp *rpcJudge.JudgeResponse, _ error) {
	resp = new(rpcJudge.JudgeResponse)
	resp.StatusCode = status.CodeServerBusy.Code()
	code, langID, problemID := req.GetCode(), req.GetLangID(), req.GetProblemID()

	// 将代码文件上传
	path, err := compiler.SaveCode(code, int(langID))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	resp.CodePath = path

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

	// 将判题请求打入MQ
	resp.JudgeID = uuid.NewString()
	msg, err := mq.GenerateJudgeMQMsg(resp.JudgeID, path, langID, problem)
	if err != nil {
		return
	}
	if mq.RMQJudge.Publish(msg) != nil {
		return
	}

	resp.StatusCode = status.CodeSuccess.Code()
	return
}
