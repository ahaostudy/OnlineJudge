package submit

import (
	"context"
	"errors"
	rpcJudge "main/api/judge"
	rpcProblem "main/api/problem"
	"main/api/submit"
	"main/config"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/data/repository"
	"main/internal/middleware/mq"
	"main/internal/middleware/redis"
	status "main/internal/service/judge/pkg/code"
	"main/rpc"
	"time"

	"gorm.io/gorm"
)

func (SubmitServer) Submit(ctx context.Context, req *rpcSubmit.SubmitRequest) (resp *rpcSubmit.SubmitResponse, _ error) {
	resp = new(rpcSubmit.SubmitResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 提交判题
	res, err := rpc.JudgeCli.Judge(ctx, &rpcJudge.JudgeRequest{
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
	if err := repository.InsertSubmit(submit); err != nil {
		return
	}

	// 将提交写入缓存
	err = redis.Rdb.Set(ctx, redis.GenerateSubmitKey(submit.ID), res.GetJudgeID(), time.Duration(config.ConfRedis.ShortTtl)*time.Second).Err()
	if err != nil {
		return
	}

	resp.SubmitID = submit.ID
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (SubmitServer) SubmitContest(ctx context.Context, req *rpcSubmit.SubmitContestRequest) (resp *rpcSubmit.SubmitContestResponse, _ error) {
	resp = new(rpcSubmit.SubmitContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 1. 必须已报名且在比赛过程中且该比赛中存在该赛题
	contest, err := repository.GetContestAndIsRegister(req.GetContestID(), req.GetUserID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeContestNotExist.Code()
		return
	}
	if err != nil {
		return
	}
	problem, err := rpc.ProblemCli.GetProblem(ctx, &rpcProblem.GetProblemRequest{
		ProblemID: req.GetProblemID(),
	})
	if err != nil {
		return
	}
	if problem.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = problem.StatusCode
		return
	}
	if problem.Problem.GetContestID() != req.GetContestID() {
		resp.StatusCode = code.CodeContestNotExist.Code()
		return
	}

	// 未报名比赛
	if !contest.IsRegister {
		resp.StatusCode = code.CodeNotRegistred.Code()
		return
	}
	// 判断是否在比赛过程中
	now := time.Now()
	if now.Before(contest.StartTime) || now.After(contest.EndTime) {
		resp.StatusCode = code.CodeContestNotOngoing.Code()
		return
	}

	// 2. 提交判题
	res, err := rpc.JudgeCli.Judge(ctx, &rpcJudge.JudgeRequest{
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
	submit := &model.Submit{
		UserID:    req.GetUserID(),
		ProblemID: req.GetProblemID(),
		LangID:    req.GetLangID(),
		Code:      res.GetCodePath(),
		Status:    int64(status.StatusRunning),
		ContestID: req.GetContestID(),
	}
	if err := repository.InsertSubmit(submit); err != nil {
		return
	}

	// 将提交写入缓存
	err = redis.Rdb.Set(ctx, redis.GenerateSubmitKey(submit.ID), res.GetJudgeID(), time.Duration(config.ConfRedis.ShortTtl)*time.Second).Err()
	if err != nil {
		return
	}

	resp.SubmitID = submit.ID
	resp.StatusCode = code.CodeSuccess.Code()

	// 4. 将提交打入MQ，进行异步计分、排名等
	// BUG: 应等待判题完成后再进行计分
	msg, err := mq.GenerateContestSubmitMQMsg(req.GetContestID(), req.GetProblemID(), req.GetUserID())
	if err != nil {
		return
	}
	if mq.RMQContestSubmit.Publish(msg); err != nil {
		return
	}

	return
}
