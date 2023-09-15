package judge

import (
	"context"
	"fmt"
	"main/api/judge"
	rpcProblem "main/api/problem"
	"main/config"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/data/model"
	"main/internal/middleware/mq"
	"main/rpc"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// 判题服务
func (JudgeServer) Judge(ctx context.Context, req *rpcJudge.JudgeRequest) (resp *rpcJudge.JudgeResponse, _ error) {
	resp = new(rpcJudge.JudgeResponse)
	resp.StatusCode = common.CodeServerBusy.Code()
	code, langID, _ := req.GetCode(), req.GetLangID(), req.GetProblemID()

	// 将代码文件上传
	suffix := model.GetLangSuffix(int(langID))
	resp.CodePath = fmt.Sprintf("%s.%s", uuid.NewString(), suffix)
	codePath := filepath.Join(config.ConfJudge.File.CodePath, resp.CodePath)
	os.WriteFile(codePath, code, 0644)

	// 根据题目ID获取题目信息
	prob, err := rpc.ProblemCli.GetProblem(context.Background(), &rpcProblem.GetProblemRequest{ProblemID: req.GetProblemID()})
	if err != nil {
		return
	}

	// 将结果转换为模型对象
	problem, err := build.UnBuildProblem(prob.GetProblem())
	if err != nil {
		return
	}

	// 将请求放入MQ
	resp.JudgeID = uuid.NewString()
	msg, err := mq.GenerateJudgeMQMsg(resp.JudgeID, codePath, langID, problem)
	if err != nil {
		return
	}
	if mq.RMQJudge.Publish(msg) != nil {
		return
	}

	// 为当前判题开辟一个管道
	mq.ResultChan[resp.JudgeID] = make(chan mq.JudgeResponse)
	// 延迟30s，如果管道未使用将自动清空
	// TODO: 待优化，可以让执行结果被读取后自动关闭该协程
	// go func() {
	// 	time.Sleep(30 * time.Second)
	// 	if ch, ok := mq.ResultChan[resp.JudgeID]; ok {
	// 		if ch != nil {
	// 			close(ch)
	// 		}
	// 		delete(mq.ResultChan, resp.JudgeID)
	// 	}
	// }()

	resp.StatusCode = common.CodeSuccess.Code()
	return
}
