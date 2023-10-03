package submit

import (
	"context"
	judge "main/kitex_gen/judge"
	submit "main/kitex_gen/submit"
	"main/pkg/code"
	"main/pkg/status"
	"main/services/submit/client"
	"main/services/submit/dal/db"
	"main/services/submit/dal/model"
)

// SubmitServiceImpl implements the last service interface defined in the IDL.
type SubmitServiceImpl struct{}

// Debug implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) Debug(ctx context.Context, req *submit.DebugReqeust) (resp *submit.DebugResponse, err error) {
	resp = new(submit.DebugResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 使用给定的代码、输入和语言ID调用 JudgeCli 的 Debug 方法
	res, err := client.JudgeCli.Debug(ctx, &judge.DebugRequest{
		Code:   req.GetCode(),
		Input:  req.GetInput(),
		LangID: req.GetLangID(),
	})
	if err != nil {
		return
	}

	// 将响应结果和状态码设置为来自 JudgeCli 响应的值
	resp.Result = &submit.JudgeResult{
		Time: res.Result.GetTime(),
		Memory: res.Result.GetMemory(),
		Status: res.Result.GetStatus(),
		Message: res.Result.GetMessage(),
		Output: res.Result.GetOutput(),
		Error: res.Result.GetError(),
	}
	resp.StatusCode = res.StatusCode

	return
}

// Submit implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) Submit(ctx context.Context, req *submit.SubmitRequest) (resp *submit.SubmitResponse, err error) {
	resp = new(submit.SubmitResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 提交判题
	res, err := client.JudgeCli.Judge(ctx, &judge.JudgeRequest{
		ProblemID: req.GetProblemID(),
		Code:      req.GetCode(),
		LangID:    req.GetLangID(),
	})
	if err != nil {
		return
	}
	if res.GetStatusCode() != code.CodeSuccess.Code() {
		resp.StatusCode = res.GetStatusCode()
		return
	}

	// 将提交写入数据库
	submit := &model.Submit{
		UserID:    req.GetUserID(),
		ProblemID: req.GetProblemID(),
		LangID:    req.GetLangID(),
		Code:      res.GetCodePath(),
		Status:    int64(status.StatusRunning),
	}
	if err := db.InsertSubmit(submit); err != nil {
		return
	}

	// 将提交写入缓存
	if cache.Rdb.Set(ctx, cache.GenerateSubmitKey(submit.ID), res.GetJudgeID(), time.Duration(config.Config.Redis.ShortTtl)*time.Second).Err() != nil {
		return
	}
	if cache.Rdb.SAdd(ctx, cache.GenerateSubmitsKey(), submit.ID).Err() != nil {
		return
	}

	resp.SubmitID = submit.ID
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// SubmitContest implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) SubmitContest(ctx context.Context, req *submit.SubmitContestRequest) (resp *submit.SubmitContestResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSubmitResult implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitResult(ctx context.Context, req *submit.GetSubmitResultRequest) (resp *submit.GetSubmitResultResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSubmitList implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitList(ctx context.Context, req *submit.GetSubmitListRequest) (resp *submit.GetSubmitListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSubmit implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmit(ctx context.Context, req *submit.GetSubmitRequest) (resp *submit.GetSubmitResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSubmitStatus implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitStatus(ctx context.Context, req *submit.GetSubmitStatusRequest) (resp *submit.GetSubmitStatusResponse, err error) {
	// TODO: Your code here...
	return
}

// IsAccepted implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) IsAccepted(ctx context.Context, req *submit.IsAcceptedRequest) (resp *submit.IsAcceptedResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAcceptedStatus implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetAcceptedStatus(ctx context.Context, req *submit.GetAcceptedStatusRequest) (resp *submit.GetAcceptedStatusResponse, err error) {
	// TODO: Your code here...
	return
}

// GetLatestSubmits implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetLatestSubmits(ctx context.Context, req *submit.GetLatestSubmitsRequest) (resp *submit.GetLatestSubmitsResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteSubmit implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) DeleteSubmit(ctx context.Context, req *submit.DeleteSubmitRequest) (resp *submit.DeleteSubmitResponse, err error) {
	// TODO: Your code here...
	return
}
