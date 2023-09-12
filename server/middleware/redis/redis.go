package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"main/config"
	"time"
)

var (
	ctx = context.Background()
	Rdb *redis.Client
	Nil = redis.Nil
)

// InitRedis 初始化Redis
func InitRedis() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.ConfRedis.Addr,
		Password: config.ConfRedis.Password,
		DB:       0,
	})

	// 测试连接
	ctx, cancel := WithTimeoutContext(2)
	defer cancel()
	if err := Rdb.Ping(ctx).Err(); err != nil {
		return err
	}

	return nil
}

// WithTimeoutContext 超时上下文
func WithTimeoutContext(second time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, second*time.Second)
}

// Flush 刷新key的有效时间
func Flush(key string) {
	ctx, cancel := WithTimeoutContext(2)
	defer cancel()

	Rdb.Expire(ctx, key, time.Duration(time.Duration(config.ConfRedis.Ttl)*time.Second))
}

// Set 设置一个key
// 该函数主要提供给需要并发且忽略错误的情况
func Set(key string, val any) {
	ctx, cancel := WithTimeoutContext(2)
	defer cancel()

	Rdb.Set(ctx, key, val, time.Duration(time.Duration(config.ConfRedis.Ttl)*time.Second))
}

// Del 删除一个key
// 该函数主要提供给需要并发且忽略错误的情况
func Del(key string) {
	ctx, cancel := WithTimeoutContext(2)
	defer cancel()

	Rdb.Del(ctx, key)
}
