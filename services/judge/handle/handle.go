package handle

import (
	"context"
	"fmt"
	"main/config"
	"main/model"
	rpcJudge "main/rpc/judge"
	"main/services/judge/pkg/code"
	"main/services/judge/pkg/judge"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type JudgeServer struct {
	rpcJudge.UnimplementedJudgeServiceServer
}

type JudgeRequest struct {
	JudgeID  string
	Problem  model.Problem
	CodePath string
	LangID   int
}

type JudgeResponse struct {
	JudgeID string
	Result  code.Result
	Err     error
}

var reqMQ = make(chan JudgeRequest)
var resultChan = make(map[string]chan JudgeResponse)

// 判题器
func Judger() {
	for {
		req := <-reqMQ
		result, err := judge.Judge(&req.Problem, req.CodePath, req.LangID)
		resultChan[req.JudgeID] <- JudgeResponse{JudgeID: req.JudgeID, Result: result, Err: err}
	}
}

func (JudgeServer) Judge(ctx context.Context, req *rpcJudge.JudgeRequest) (*rpcJudge.JudgeResponse, error) {
	fmt.Println("judge ... ")

	code, langID, _ := req.GetCode(), req.GetLangID(), req.GetProblemID()

	// TODO: 将代码文件上传
	path := filepath.Join(config.ConfJudge.File.TempPath, fmt.Sprintf("%s.py", uuid.NewString()))
	os.WriteFile(path, code, 0644)

	// TODO: 根据题目ID获取题目信息
	problem := &model.Problem{
		MaxTime:   1000,
		MaxMemory: 512 * 1024 * 1024,
		Testcases: []*model.Testcase{
			{ID: 1, InputPath: "1.in", OutputPath: "1.out"},
			{ID: 2, InputPath: "2.in", OutputPath: "2.out"},
			{ID: 3, InputPath: "3.in", OutputPath: "3.out"},
		},
	}

	// TODO: 将请求放入MQ
	judgeID := uuid.NewString()
	reqMQ <- JudgeRequest{
		JudgeID:  judgeID,
		Problem:  *problem,
		CodePath: path,
		LangID:   int(langID),
	}
	resultChan[judgeID] = make(chan JudgeResponse)

	return &rpcJudge.JudgeResponse{
		JudgeID: judgeID,
		Ok:      true,
	}, nil
}

func (JudgeServer) GetResult(ctx context.Context, req *rpcJudge.GetResultRequest) (*rpcJudge.GetResultResponse, error) {
	fmt.Println("get result ... ")
	res := <-resultChan[req.GetJudgeID()]

	close(resultChan[req.JudgeID])
	resultChan[req.JudgeID] = nil

	return &rpcJudge.GetResultResponse{
		Result: &rpcJudge.JudgeResult{
			Time:    res.Result.Time,
			Memory:  res.Result.Memory,
			Status:  int64(res.Result.Status),
			Message: res.Result.Message,
			Output:  res.Result.Output,
			Error:   res.Result.Error,
		},
		Ok: res.Err == nil,
	}, nil
}
