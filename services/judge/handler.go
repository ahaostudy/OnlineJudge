package judge

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	judge "main/kitex_gen/judge"
	"main/kitex_gen/problem"
	"main/common/code"
	"main/common/pack"
	"main/common/status"
	"main/services/judge/client"
	"main/services/judge/config"
	"main/services/judge/dal/cache"
	"main/services/judge/dal/model"
	"main/services/judge/dal/mq"
	coder "main/services/judge/pkg/code"
	"main/services/judge/pkg/compiler"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// JudgeServiceImpl implements the last service interface defined in the IDL.
type JudgeServiceImpl struct{}

// Judge implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) Judge(ctx context.Context, req *judge.JudgeRequest) (resp *judge.JudgeResponse, _ error) {
	resp = new(judge.JudgeResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 将代码文件上传
	path, err := compiler.SaveCode(req.GetCode(), int(req.GetLangID()))
	if err != nil {
		return
	}
	resp.CodePath = path

	// 根据题目ID获取题目信息
	prob, err := client.ProblemCli.GetProblem(context.Background(), &problem.GetProblemRequest{ProblemID: req.GetProblemID()})
	if err != nil {
		return
	}

	// 将结果转换为模型对象
	problem := new(model.Problem)
	if new(pack.Builder).Build(prob.GetProblem(), problem).Error() != nil {
		return
	}

	// 将判题请求打入MQ
	resp.JudgeID = uuid.NewString()
	msg, err := mq.GenerateJudgeMQMsg(resp.JudgeID, path, req.GetLangID(), problem)
	if err != nil {
		return
	}
	if mq.RMQJudge.Publish(msg) != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetResult implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) GetResult(ctx context.Context, req *judge.GetResultRequest) (resp *judge.GetResultResponse, _ error) {
	resp = new(judge.GetResultResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取判题结果
	res, err := cache.Rdb.Get(ctx, cache.GenerateJudgeKey(req.GetJudgeID())).Bytes()
	// 运行中，未获取到结果
	if err != nil || len(res) == 0 {
		resp.Result = &judge.JudgeResult{Status: int64(status.StatusRunning)}
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	// 将结果反序列化
	result := new(mq.JudgeResponse)
	if err := json.Unmarshal(res, result); err != nil {
		return
	}
	if result.Error != nil {
		resp.Result = &judge.JudgeResult{Status: int64(status.StatusServerFailed)}
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	// 将结果转换为rpc响应
	resp.Result = new(judge.JudgeResult)
	if new(pack.Builder).Build(result.Result, resp.Result).Error() != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// Debug implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) Debug(ctx context.Context, req *judge.DebugRequest) (resp *judge.DebugResponse, _ error) {
	resp = new(judge.DebugResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	codeName, err := compiler.SaveCode(req.GetCode(), int(req.GetLangID()))
	if err != nil {
		return
	}
	codePath := filepath.Join(config.Config.File.CodePath, codeName)
	defer os.Remove(codePath)

	inputPath := filepath.Join(config.Config.File.TempPath, fmt.Sprintf("%s.in", uuid.New().String()))
	err = os.WriteFile(inputPath, req.GetInput(), 0644)
	if err != nil {
		return
	}
	defer os.Remove(inputPath)

	c := coder.NewCode(codePath, int(req.GetLangID()))

	result, err := c.Run(inputPath)
	defer c.Destroy()
	if err != nil {
		return
	}

	resp.Result = new(judge.JudgeResult)
	if new(pack.Builder).Build(&result, resp.Result).Error() != nil {
		return
	}
	if resp.Result.GetStatus() == int64(status.StatusAccepted) {
		resp.Result.Status = int64(status.StatusFinished)
	}
	fmt.Printf("resp: %v\n", resp.Result)

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetCode implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) GetCode(ctx context.Context, req *judge.GetCodeRequest) (resp *judge.GetCodeResponse, _ error) {
	resp = new(judge.GetCodeResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取代码内容
	body, err := os.ReadFile(filepath.Join(config.Config.File.CodePath, req.GetCodePath()))
	if errors.Is(err, os.ErrNotExist) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}

	// 将模型对象转换为rpc响应
	resp.Code = body

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
