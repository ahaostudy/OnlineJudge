package submit

import (
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"

	rpcJudge "main/api/judge"
	rpcSubmit "main/api/submit"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/data/repository"
	"main/internal/middleware/mq"
	"main/internal/middleware/redis"
	status "main/internal/service/judge/pkg/code"
)

func (SubmitServer) GetSubmit(_ context.Context, req *rpcSubmit.GetSubmitRequest) (resp *rpcSubmit.GetSubmitResponse, _ error) {
	resp = new(rpcSubmit.GetSubmitResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交数据
	submit, err := repository.GetSubmit(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将模型对象转换为rpc响应
	resp.Submit, err = build.BuildSubmit(submit)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (SubmitServer) GetSubmitList(ctx context.Context, req *rpcSubmit.GetSubmitListRequest) (resp *rpcSubmit.GetSubmitListResponse, _ error) {
	resp = new(rpcSubmit.GetSubmitListResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交数据
	var submitList []*model.Submit
	var err error
	if req.GetUserID() == 0 && req.ProblemID == 0 {
		return
	} else if req.GetUserID() == 0 {
		submitList, err = repository.GetSubmitListByProblem(req.GetProblemID())
	} else if req.GetProblemID() == 0 {
		submitList, err = repository.GetSubmitListByUser(req.GetUserID())
	} else {
		submitList, err = repository.GetSubmitList(req.GetUserID(), req.GetProblemID())
	}
	if err != nil {
		return
	}

	// 将对象转换为rpc响应
	resp.SubmitList, err = build.BuildSubmitList(submitList)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (SubmitServer) GetSubmitResult(ctx context.Context, req *rpcSubmit.GetSubmitResultRequest) (resp *rpcSubmit.GetSubmitResultResponse, _ error) {
	resp = new(rpcSubmit.GetSubmitResultResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取判题ID
	judgeID, err := redis.Rdb.Get(ctx, redis.GenerateSubmitKey(req.GetSubmitID())).Result()
	if err != nil {
		return
	}
	// 获取不到，说明提交的key不存在或已过期
	if len(judgeID) == 0 {
		resp.StatusCode = code.CodeSubmitNotFound.Code()
		return
	}

	// 获取判题结果
	res, err := redis.Rdb.Get(ctx, redis.GenerateJudgeKey(judgeID)).Bytes()
	if err != nil {
		return
	}
	// 运行中，未获取到结果
	if len(res) == 0 {
		resp.Result = &rpcJudge.JudgeResult{Status: int64(status.StatusRunning)}
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	// 将结果反序列化
	result := new(mq.JudgeResponse)
	if err := json.Unmarshal(res, result); err != nil {
		return
	}
	if result.Error != nil {
		resp.Result = &rpcJudge.JudgeResult{Status: int64(status.StatusServerFailed)}
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	// 将结果转换为rpc响应
	resp.Result = new(rpcJudge.JudgeResult)
	if new(build.Builder).Build(result.Result, resp.Result).Error() != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
