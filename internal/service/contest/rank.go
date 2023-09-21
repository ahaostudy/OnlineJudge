package contest

import (
	"context"
	"main/api/contest"
	"main/internal/common/code"
	// "main/internal/middleware/redis"
)

var count = 20

func (ContestServer) ContestRank(ctx context.Context, req *rpcContest.ContestRankRequest) (resp *rpcContest.ContestRankResponse, _ error) {
	resp = new(rpcContest.ContestRankResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// TODO
	// 优先从redis获取排名
	// redis.Rdb.ZRandMember(ctx, redis.GenerateRankKey(req.ContestID), count, 0).Result()
	
	// 不存在从mongodb读取放到redis
	// 将通过redis的zset排序后的数据返回

	return
}
