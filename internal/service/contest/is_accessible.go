package contest

import (
	"context"
	"main/api/contest"
)

func (ContestServer) IsAccessible(ctx context.Context, req *rpcContest.IsAccessibleRequest) (resp *rpcContest.IsAccessibleResponse, _ error) {
	// TODO: 比赛已结束 或 比赛进行时且用户已报名

	return
}
