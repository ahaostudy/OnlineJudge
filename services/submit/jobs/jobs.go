package jobs

import (
	"context"
	"main/common/code"
	"main/common/status"
	"main/kitex_gen/judge"
	"main/services/submit/client"
	"main/services/submit/config"
	"main/services/submit/dal/cache"
	"main/services/submit/dal/db"
	"main/services/submit/dal/model"
	"main/services/submit/dal/mq"
	"strconv"
	"sync"
	"time"
)

func RunSubmitJobs() {
	go run(saveSubmitResult)
}

func run(f func()) {
	ticker := time.NewTicker(time.Duration(config.Config.Jobs.Time) * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		go f()
	}
}

func saveSubmitResult() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 获取未处理的所有提交
	subs, err := cache.Rdb.SMembers(ctx, cache.GenerateSubmitsKey()).Result()
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
			jid, err := cache.Rdb.Get(ctx, cache.GenerateSubmitKey(sid)).Result()
			if err != nil {
				return
			}

			// 获取运行结果
			res, err := client.JudgeCli.GetResult(ctx, &judge.GetResultRequest{
				JudgeID: jid,
			})
			// TODO: 处理error
			if err != nil || res.StatusCode != code.CodeSuccess.Code() || res.GetResult().GetStatus() == int64(status.StatusRunning) {
				return
			}

			// 将结果更新到数据库
			result := res.GetResult()

			submit := &model.Submit{
				Status: int64(result.Status),
				Time:   result.Time,
				Memory: result.Memory,
			}
			if err := db.UpdateSubmit(sid, submit); err != nil {
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
			cache.Rdb.SRem(ctx, cache.GenerateSubmitsKey(), sid)
		}(s)
	}
	wg.Wait()
}
