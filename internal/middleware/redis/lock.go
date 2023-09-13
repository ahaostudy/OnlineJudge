package redis

import (
	"errors"
	"fmt"
	"main/config"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// Lock 获取乐观锁
func Lock(rdb *redis.Client, key string) (string, error) {
	ctx, cancel := WithTimeoutContext(2)
	defer cancel()

	// 生成锁的key和ID
	lockKey := fmt.Sprintf("%s:%s", config.ConfRedis.LockKey, key)
	lockID := uuid.New().String()

	// Redis 主动轮询取锁
	// 每隔50ms尝试一次取锁，最多取锁5次，取锁成功后返回锁ID
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < 5; i++ {
		result, err := rdb.SetNX(ctx, lockKey, lockID, time.Second).Result()
		if err == nil && result {
			return lockID, nil
		}
		<-ticker.C
	}

	// 多次取锁失败
	return lockID, errors.New("lock is acquired by others")
}

// Unlock 解锁
func Unlock(rdb *redis.Client, key string, lockID string) bool {
	ctx, cancel := WithTimeoutContext(2)
	defer cancel()

	lockKey := fmt.Sprintf("%s:%s", config.ConfRedis.LockKey, key)

	// 获取当前锁的ID
	// 检测锁ID是否匹配，不匹配则解锁失败
	lockVal, err := rdb.Get(ctx, lockKey).Result()
	if err != nil || lockVal != lockID {
		return false
	}

	// 解锁
	return rdb.Del(ctx, lockKey).Err() == nil
}
