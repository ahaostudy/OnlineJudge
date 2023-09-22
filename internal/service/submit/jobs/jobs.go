package jobs

import (
	"encoding/json"
	"main/config"
	"main/internal/common/ctxt"
	"main/internal/data/model"
	"main/internal/data/repository"
	"main/internal/middleware/mq"
	"main/internal/middleware/redis"
	"strconv"
	"sync"
	"time"
)

func RunSubmitJobs() {
	go run(saveSubmitResult)
}

func run(f func()) {
	ticker := time.NewTicker(time.Duration(config.ConfSubmit.Jobs.Time) * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		go f()
	}
}

func saveSubmitResult() {
	ctx, cancel := ctxt.WithTimeoutContext(3)
	defer cancel()

	// 获取未处理的所有提交
	subs, err := redis.Rdb.SMembers(ctx, redis.GenerateSubmitsKey()).Result()
	if err != nil {
		return
	}

	wg := new(sync.WaitGroup)
	for _, s := range subs {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			sid, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return
			}

			// 获取judge_id
			jid, err := redis.Rdb.Get(ctx, redis.GenerateSubmitKey(sid)).Result()
			if err != nil {
				return
			}

			// 获取判题结果
			res, err := redis.Rdb.Get(ctx, redis.GenerateJudgeKey(jid)).Bytes()
			if err != nil {
				return
			}

			// 如果判题未完成则跳过
			if len(res) == 0 {
				return
			}

			// 将结果更新到数据库
			result := new(mq.JudgeResponse)
			if err := json.Unmarshal(res, result); err != nil {
				return
			}
			// TODO: 处理result的error

			submit := &model.Submit{
				Status: int64(result.Result.Status),
				Time:   result.Result.Time,
				Memory: result.Result.Memory,
			}
			if err := repository.UpdateSubmit(sid, submit); err != nil {
				return
			}

			// 如果为比赛时提交，则通过MQ更新排行榜
			if submit.ContestID != 0 {
				msg, err := mq.GenerateContestSubmitMQMsg(submit.ContestID, submit.UserID)
				if err != nil {
					return
				}
				if err := mq.RMQContestSubmit.Publish(msg); err != nil {
					return
				}
			}

			// 成功后将ID缓存删除
			redis.Rdb.SRem(ctx, redis.GenerateSubmitsKey(), sid)
		}(s)
	}
	wg.Wait()
}
