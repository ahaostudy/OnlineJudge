package contest

import (
	"context"
	"errors"
	"main/api/contest"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/repository"

	"gorm.io/gorm"
)

func (ContestServer) GetContest(ctx context.Context, req *rpcContest.GetContestRequest) (resp *rpcContest.GetContestResponse, _ error) {
	resp = new(rpcContest.GetContestResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 从数据库获取比赛信息
	contest, err := repository.GetContent(req.GetId())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}
	// TODO: 返回比赛的题目基础信息
	// 包括ID、标题等

	// 将比赛信息转换为rpc响应
	resp.Contest, err = build.BuildContest(contest)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
