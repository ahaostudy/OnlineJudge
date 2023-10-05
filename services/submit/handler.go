package submit

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"main/kitex_gen/contest"
	judge "main/kitex_gen/judge"
	"main/kitex_gen/problem"
	submit "main/kitex_gen/submit"
	"main/pkg/code"
	"main/pkg/pack"
	"main/pkg/status"
	"main/rpc"
	"main/services/submit/client"
	"main/services/submit/config"
	"main/services/submit/dal/cache"
	"main/services/submit/dal/db"
	"main/services/submit/dal/model"
)

// SubmitServiceImpl implements the last service interface defined in the IDL.
type SubmitServiceImpl struct{}

// Debug implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) Debug(ctx context.Context, req *submit.DebugReqeust) (resp *submit.DebugResponse, _ error) {
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
		Time:    res.Result.GetTime(),
		Memory:  res.Result.GetMemory(),
		Status:  res.Result.GetStatus(),
		Message: res.Result.GetMessage(),
		Output:  res.Result.GetOutput(),
		Error:   res.Result.GetError(),
	}
	resp.StatusCode = res.StatusCode

	return
}

// Submit implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) Submit(ctx context.Context, req *submit.SubmitRequest) (resp *submit.SubmitResponse, _ error) {
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
func (s *SubmitServiceImpl) SubmitContest(ctx context.Context, req *submit.SubmitContestRequest) (resp *submit.SubmitContestResponse, _ error) {
	resp = new(submit.SubmitContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 1. 必须已报名且在比赛过程中且该比赛中存在该赛题
	c, err := client.ContestCli.GetContest(ctx, &contest.GetContestRequest{
		ID:     req.GetContestID(),
		UserID: req.GetUserID(),
	})
	if err != nil {
		return
	}
	if !c.Contest.IsRegister {
		resp.StatusCode = code.CodeNotRegistred.Code()
		return
	}
	// 判断是否在比赛过程中
	now := time.Now().UnixMilli()
	if now < c.GetContest().GetStartTime() || now > c.GetContest().GetEndTime() {
		resp.StatusCode = code.CodeContestNotOngoing.Code()
		return
	}
	// 判断是否存在该赛题
	p, err := client.ProblemCli.GetProblem(ctx, &problem.GetProblemRequest{
		ProblemID: req.GetProblemID(),
	})
	if err != nil {
		return
	}
	if p.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = p.StatusCode
		return
	}
	if p.Problem.GetContestID() != req.GetContestID() {
		resp.StatusCode = code.CodeContestNotExist.Code()
		return
	}

	// 2. 提交判题
	res, err := rpc.JudgeCli.Judge(ctx, &judge.JudgeRequest{
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

	// 3. 将提交记录入库MySQL
	sub := &model.Submit{
		UserID:    req.GetUserID(),
		ProblemID: req.GetProblemID(),
		LangID:    req.GetLangID(),
		Code:      res.GetCodePath(),
		Status:    int64(status.StatusRunning),
		ContestID: req.GetContestID(),
	}
	if err := db.InsertSubmit(sub); err != nil {
		return
	}

	// 将提交写入缓存
	if cache.Rdb.Set(ctx, cache.GenerateSubmitKey(sub.ID), res.GetJudgeID(), time.Duration(config.Config.Redis.ShortTtl)*time.Second).Err() != nil {
		return
	}
	if cache.Rdb.SAdd(ctx, cache.GenerateSubmitsKey(), sub.ID).Err() != nil {
		return
	}

	resp.SubmitID = sub.ID
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetSubmitResult implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitResult(ctx context.Context, req *submit.GetSubmitResultRequest) (resp *submit.GetSubmitResultResponse, _ error) {
	resp = new(submit.GetSubmitResultResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取判题ID
	judgeID, err := cache.Rdb.Get(ctx, cache.GenerateSubmitKey(req.GetSubmitID())).Result()
	if err != nil {
		return
	}
	// 获取不到，说明提交的key不存在或已过期
	if len(judgeID) == 0 {
		resp.StatusCode = code.CodeSubmitNotFound.Code()
		return
	}

	// 获取判题结果
	res, err := client.JudgeCli.GetResult(ctx, &judge.GetResultRequest{
		JudgeID: judgeID,
	})
	if err != nil {
		return
	}
	if res.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = res.StatusCode
		return
	}

	resp.Result, err = pack.BuildResult(res.Result)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetSubmitList implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitList(ctx context.Context, req *submit.GetSubmitListRequest) (resp *submit.GetSubmitListResponse, _ error) {
	resp = new(submit.GetSubmitListResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交数据
	var submitList []*model.Submit
	var err error
	if req.GetUserID() == 0 && req.ProblemID == 0 {
		return
	} else if req.GetUserID() == 0 {
		submitList, err = db.GetSubmitListByProblem(req.GetProblemID())
	} else if req.GetProblemID() == 0 {
		submitList, err = db.GetSubmitListByUser(req.GetUserID())
	} else {
		submitList, err = db.GetSubmitList(req.GetUserID(), req.GetProblemID())
	}
	if err != nil {
		return
	}

	// 将对象转换为rpc响应
	resp.SubmitList, err = pack.BuildSubmitList(submitList)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetSubmit implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmit(ctx context.Context, req *submit.GetSubmitRequest) (resp *submit.GetSubmitResponse, _ error) {
	resp = new(submit.GetSubmitResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交数据
	sub, err := db.GetSubmit(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将模型对象转换为rpc响应
	resp.Submit, err = pack.BuildSubmit(sub)
	if err != nil {
		return
	}

	// 获取提交的代码内容
	res, err := rpc.JudgeCli.GetCode(ctx, &judge.GetCodeRequest{CodePath: sub.Code})
	if err != nil {
		return
	}
	if res.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = res.StatusCode
		return
	}
	// 复制代码内容
	resp.Submit.Code = string(res.Code)

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetSubmitStatus implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitStatus(ctx context.Context, req *submit.GetSubmitStatusRequest) (resp *submit.GetSubmitStatusResponse, _ error) {
	resp = new(submit.GetSubmitStatusResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交记录信息
	status, err := db.GetSubmitStatus()
	if err != nil {
		return
	}

	resp.SubmitStatus = make(map[int64]*submit.SubmitStatus)
	for _, s := range status {
		resp.SubmitStatus[s.ProblemID] = &submit.SubmitStatus{
			Count:         s.Count,
			AcceptedCount: s.AcceptedCount,
		}
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// IsAccepted implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) IsAccepted(ctx context.Context, req *submit.IsAcceptedRequest) (resp *submit.IsAcceptedResponse, _ error) {
	// TODO: Your code here...
	resp = new(submit.IsAcceptedResponse)
	resp.StatusCode = code.CodeServerBusy.Code()
	return
}

// GetAcceptedStatus implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetAcceptedStatus(ctx context.Context, req *submit.GetAcceptedStatusRequest) (resp *submit.GetAcceptedStatusResponse, _ error) {
	resp = new(submit.GetAcceptedStatusResponse)

	if req.GetUserID() == 0 {
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	status, err := db.GetAcceptedStatus(req.GetUserID())
	if err != nil {
		resp.StatusCode = code.CodeServerBusy.Code()
		return
	}

	resp.AcceptedStatus = make(map[int64]bool)
	for _, s := range status {
		resp.AcceptedStatus[s.ProblemID] = s.IsAccepted
	}
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetLatestSubmits implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetLatestSubmits(ctx context.Context, req *submit.GetLatestSubmitsRequest) (resp *submit.GetLatestSubmitsResponse, _ error) {
	resp = new(submit.GetLatestSubmitsResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交记录信息
	submits, err := db.GetUserLastSubmits(req.GetUserID(), int(req.GetCount()))
	if err != nil {
		return
	}

	resp.SubmitList, err = pack.BuildSubmitList(submits)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// DeleteSubmit implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) DeleteSubmit(ctx context.Context, req *submit.DeleteSubmitRequest) (resp *submit.DeleteSubmitResponse, _ error) {
	resp = new(submit.DeleteSubmitResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交数据
	sub, err := db.GetSubmit(req.GetID())
	if err == gorm.ErrRecordNotFound || (err == nil && sub.UserID != req.GetUserID()) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 删除一条记录
	err = db.DeleteSubmit(req.GetID())
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
