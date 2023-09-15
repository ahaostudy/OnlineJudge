package private

import (
	"context"
	"fmt"
	"main/api/private"
	"main/api/problem"
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

// Judge 判题服务
func (PrivateServer) Judge(ctx context.Context, req *rpcPrivate.JudgeRequest) (resp *rpcPrivate.JudgeResponse, _ error) {
	resp = new(rpcPrivate.JudgeResponse)
	resp.StatusCode = common.CodeServerBusy.Code()
	code, langID, problemID := req.GetCode(), req.GetLangID(), req.GetProblemID()

	// 将代码文件上传
	suffix := model.GetLangSuffix(int(langID))
	path := filepath.Join(config.ConfJudge.File.CodePath, fmt.Sprintf("private/%s.%s", uuid.NewString(), suffix))
	os.WriteFile(path, code, 0644)

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

	// 将请求放入MQ
	resp.JudgeID = uuid.NewString()
	msg, err := mq.GenerateJudgeMQMsg(resp.JudgeID, path, langID, problem)
	if err != nil {
		return
	}
	if mq.RMQPrivate.Publish(msg) != nil {
		return
	}

	// 为当前判题开辟一个管道
	mq.ResultChan[resp.JudgeID] = make(chan mq.JudgeResponse)

	resp.StatusCode = common.CodeSuccess.Code()
	return
}

// GetResult 获取运行结果
func (PrivateServer) GetResult(ctx context.Context, req *rpcPrivate.GetResultRequest) (resp *rpcPrivate.GetResultResponse, _ error) {
	resp = new(rpcPrivate.GetResultResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 读取管道获取结果并关闭管道
	res := <-mq.ResultChan[req.GetJudgeID()]
	close(mq.ResultChan[req.JudgeID])
	delete(mq.ResultChan, req.JudgeID)

	// 判断运行是否错误，并复制Result
	resp.Result = new(rpcPrivate.JudgeResult)
	if new(build.Builder).Build(res.Result, resp.Result).Error() == nil {
		resp.StatusCode = common.CodeSuccess.Code()
	}

	return
}
