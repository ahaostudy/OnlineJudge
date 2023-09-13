package handle

import (
	"context"
	"fmt"
	"main/api/judge"
	rpcProblem "main/api/problem"
	"main/config"
	"main/internal/common"
	"main/internal/data/model"
	"main/internal/middleware/mq"
	"main/rpc"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type JudgeServer struct {
	rpcJudge.UnimplementedJudgeServiceServer
}

// 判题服务
func (JudgeServer) Judge(ctx context.Context, req *rpcJudge.JudgeRequest) (resp *rpcJudge.JudgeResponse, _ error) {
	resp = new(rpcJudge.JudgeResponse)
	resp.StatusCode = common.CodeServerBusy.Code()
	resp.StatusMsg = common.CodeServerBusy.Msg()
	code, langID, _ := req.GetCode(), req.GetLangID(), req.GetProblemID()

	// 将代码文件上传
	suffix := model.GetLangSuffix(int(langID))
	path := filepath.Join(config.ConfJudge.File.TempPath, fmt.Sprintf("%s.%s", uuid.NewString(), suffix))
	os.WriteFile(path, code, 0644)

	// 根据题目ID获取题目信息
	prob, err := rpc.ProblemCli.GetProblem(context.Background(), &rpcProblem.GetProblemRequest{ProblemID: req.GetProblemID()})
	if err != nil {
		return
	}
	reqProblem := prob.GetProblem()
	// 将结果转换为模型对象
	problem := new(model.Problem)
	builder := new(common.Builder).Build(reqProblem, problem)
	for i := range reqProblem.Testcases {
		t := new(model.Testcase)
		builder.Build(reqProblem.Testcases[i], t)
		problem.Testcases = append(problem.Testcases, t)
	}
	if builder.Error() != nil {
		return
	}

	// 将请求放入MQ
	resp.JudgeID = uuid.NewString()
	msg, err := mq.GenerateJudgeMQMsg(resp.JudgeID, path, langID, problem)
	if err != nil {
		return
	}
	if mq.RMQJudge.Publish(msg) != nil {
		return
	}

	// 为当前判题开辟一个管道
	mq.ResultChan[resp.JudgeID] = make(chan mq.JudgeResponse)

	resp.StatusCode = common.CodeSuccess.Code()
	resp.StatusMsg = common.CodeSuccess.Msg()
	return
}

func (JudgeServer) GetResult(ctx context.Context, req *rpcJudge.GetResultRequest) (resp *rpcJudge.GetResultResponse, _ error) {
	resp = new(rpcJudge.GetResultResponse)
	// 读取管道获取结果并关闭管道
	res := <-mq.ResultChan[req.GetJudgeID()]
	close(mq.ResultChan[req.JudgeID])
	delete(mq.ResultChan, req.JudgeID)

	// 判断运行是否错误，并复制Result
	var code common.Code
	resp.Result = new(rpcJudge.JudgeResult)
	if res.Error != nil || new(common.Builder).Build(res.Result, resp.Result).Error() != nil {
		code = common.CodeServerBusy
	} else {
		code = common.CodeSuccess
	}

	resp.StatusCode = code.Code()
	resp.StatusMsg = code.Msg()
	return
}
