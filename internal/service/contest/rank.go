package contest

import (
	"context"
	"encoding/json"
	"fmt"
	"main/api/contest"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/data/repository"
	"main/internal/middleware/redis"
	"strconv"

	status "main/internal/service/judge/pkg/code"

	rds "github.com/go-redis/redis/v8"
)

type RankStatus struct {
	Penalty  int   `json:"penalty"`
	Accepted bool  `json:"accepted"`
	AcTime   int64 `json:"ac_time"`
	LangID   int64 `json:"lang_id"`
	Score    int   `json:"score"`
}

func (ContestServer) ContestRank(ctx context.Context, req *rpcContest.ContestRankRequest) (resp *rpcContest.ContestRankResponse, _ error) {
	resp = new(rpcContest.ContestRankResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 从redis获取排名列表
	start := (req.GetPage() - 1) * req.GetCount()
	stop := req.GetPage() * req.GetCount()
	rank, err := redis.Rdb.ZRevRange(ctx, redis.GenerateRankKey(req.ContestID), start, stop).Result()
	if err != nil {
		return
	}

	// 判断是否存在缓存
	// TODO: 处理缓存穿透问题
	if len(rank) == 0 {
		// 不存在缓存时从MySQL获取到redis
		submits, err := repository.GetContestSubmits(req.ContestID)
		if err != nil {
			return
		}
		fmt.Printf("submits: %v\n", submits)
		if err := updateRedis(ctx, req.ContestID, submits); err != nil {
			return
		}

		// 重新获取redis的排名列表
		rank, err = redis.Rdb.ZRevRange(ctx, redis.GenerateRankKey(req.ContestID), start, stop).Result()
		if err != nil {
			return
		}
	}

	// 通过排名（id）列表获取用户详细分数信息
	for i, u := range rank {
		uid, err := strconv.ParseInt(u, 10, 64)
		if err != nil {
			return
		}
		resp.Rank = append(resp.Rank, &rpcContest.UserData{
			UserID: uid,
			Status: make(map[string]*rpcContest.Status),
		})
		us, err := redis.Rdb.HGetAll(ctx, redis.GenerateContestUserKey(req.ContestID, uid)).Result()
		if err != nil {
			return
		}
		for k, v := range us {
			s := new(RankStatus)
			if err := json.Unmarshal([]byte(v), s); err != nil {
				return
			}
			t, err := build.BuildRankStatus(s)
			if err != nil {
				return
			}
			resp.Rank[i].Status[k] = t
		}
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// 根据提交记录计算分数并更新到redis
func updateRedis(ctx context.Context, contestID int64, submits []*model.Submit) error {
	// BUG: 添加对redis更新失败的处理

	// 整理每位用户每题的提交记录
	subs := make(map[int64]map[int64][]int)
	for i, submit := range submits {
		if v, ok := subs[submit.UserID]; !ok || v == nil {
			subs[submit.UserID] = make(map[int64][]int)
		}
		subs[submit.UserID][submit.ProblemID] = append(subs[submit.UserID][submit.ProblemID], i)
	}

	// 遍历每位用户
	for u, us := range subs {
		var score, penalty int
		var acTime int64
		key := redis.GenerateContestUserKey(contestID, u)

		// 遍历每题
		for p, ps := range us {
			s := new(RankStatus)

			// 计算每题的分数
			for _, idx := range ps {
				if submits[idx].Status != int64(status.StatusAccepted) {
					s.Penalty = len(ps)
				} else {
					s.Accepted = true
					s.AcTime = submits[ps[0]].CreatedAt.UnixMilli()
					s.LangID = submits[ps[0]].LangID
					s.Score = 20
				}
			}
			bytes, err := json.Marshal(s)
			if err != nil {
				return err
			}

			// 更新该题的分数
			if err := redis.Rdb.HSet(ctx, key, fmt.Sprintf("problem:%d", p), string(bytes)).Err(); err != nil {
				return err
			}

			// 更新该用户的总成绩
			if s.Accepted {
				score += s.Score
				penalty += s.Penalty
				if acTime < s.AcTime {
					acTime = s.AcTime
				}
			}
		}

		// 更新用户的总成绩
		bytes, err := json.Marshal(&RankStatus{
			Penalty: penalty,
			AcTime:  acTime,
			Score:   score,
		})
		if err != nil {
			return err
		}
		if err := redis.Rdb.HSet(ctx, key, "all", string(bytes)).Err(); err != nil {
			return err
		}

		// 更新zset维护排行榜
		if err := redis.Rdb.ZRem(ctx, redis.GenerateRankKey(contestID), u).Err(); err != nil {
			return err
		}
		if err := redis.Rdb.ZAdd(ctx, redis.GenerateRankKey(contestID), &rds.Z{Score: float64(score), Member: u}).Err(); err != nil {
			return err
		}
	}

	return nil
}
