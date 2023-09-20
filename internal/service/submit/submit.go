package submit

import (
	"context"
	rpcJudge "main/api/judge"
	"main/api/submit"
	"main/config"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/data/repository"
	"main/internal/middleware/redis"
	status "main/internal/service/judge/pkg/code"
	"main/rpc"
	"time"
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

	// TODO
	// 1. 必须已报名且在比赛过程中
	// 2. 将提交记录入库MySQL
	// 3. 并发计算分数
	// 4. 将分数存储到mongodb
	// 5. 定期刷到redis

	return
}
