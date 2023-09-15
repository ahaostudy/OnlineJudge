package submit

import (
	"context"
	rpcJudge "main/api/judge"
	"main/api/submit"
	"main/config"
	"main/internal/common"
	"main/internal/data/model"
	"main/internal/data/repository"
	"main/internal/middleware/redis"
	"main/internal/service/judge/pkg/code"
	"main/rpc"
	"time"
)

func (SubmitServer) Submit(ctx context.Context, req *rpcSubmit.SubmitRequest) (resp *rpcSubmit.SubmitResponse, _ error) {
	resp = new(rpcSubmit.SubmitResponse)
	resp.StatusCode = common.CodeServerBusy.Code()

	// 提交判题
	res, err := rpc.JudgeCli.Judge(ctx, &rpcJudge.JudgeRequest{
		ProblemID: req.GetProblemID(),
		Code:      req.GetCode(),
		LangID:    req.GetLangID(),
	})
	if err != nil {
		return
	}
	if res.GetStatusCode() != common.CodeSuccess.Code() {
		resp.StatusCode = res.GetStatusCode()
		return
	}

	// 将提交写入数据库
	submit := &model.Submit{
		UserID:    req.GetUserID(),
		ProblemID: req.GetProblemID(),
		LangID:    req.GetLangID(),
		Code:      res.GetCodePath(),
		Status:    int64(code.StatusRunning),
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
	resp.StatusCode = common.CodeSuccess.Code()
	return
}
