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
	"main/rpc"
)

func (SubmitServer) GetSubmit(ctx context.Context, req *rpcSubmit.GetSubmitRequest) (resp *rpcSubmit.GetSubmitResponse, _ error) {
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

	// 获取提交的代码内容
	res, err := rpc.JudgeCli.GetCode(ctx, &rpcJudge.GetCodeRequest{CodePath: submit.Code})
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
	// 运行中，未获取到结果
	if err != nil || len(res) == 0 {
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

func (SubmitServer) GetSubmitStatus(ctx context.Context, req *rpcSubmit.GetSubmitStatusRequest) (resp *rpcSubmit.GetSubmitStatusResponse, _ error) {
	resp = new(rpcSubmit.GetSubmitStatusResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交记录信息
	status, err := repository.GetSubmitStatus()
	if err != nil {
		return
	}

	resp.SubmitStatus = make(map[int64]*rpcSubmit.SubmitStatus)
	for _, s := range status {
		resp.SubmitStatus[s.ProblemID] = &rpcSubmit.SubmitStatus{
			Count:         s.Count,
			AcceptedCount: s.AcceptedCount,
		}
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (SubmitServer) IsAccepted(ctx context.Context, req *rpcSubmit.IsAcceptedRequest) (resp *rpcSubmit.IsAcceptedResponse, _ error) {
	resp = new(rpcSubmit.IsAcceptedResponse)
	resp.StatusCode = code.CodeServerBusy.Code()
	return
}

func (SubmitServer) GetAcceptedStatus(ctx context.Context, req *rpcSubmit.GetAcceptedStatusRequest) (resp *rpcSubmit.GetAcceptedStatusResponse, _ error) {
	resp = new(rpcSubmit.GetAcceptedStatusResponse)

	if req.GetUserID() == 0 {
		resp.StatusCode = code.CodeSuccess.Code()
		return
	}

	status, err := repository.GetAcceptedStatus(req.GetUserID())
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

func (SubmitServer) GetLatestSubmits(ctx context.Context, req *rpcSubmit.GetLatestSubmitsRequest) (resp *rpcSubmit.GetLatestSubmitsResponse, _ error) {
	resp = new(rpcSubmit.GetLatestSubmitsResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取提交记录信息
	submits, err := repository.GetUserLastSubmits(req.GetUserID(), int(req.GetCount()))
	if err != nil {
		return
	}

	resp.SubmitList, err = build.BuildSubmitList(submits)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
