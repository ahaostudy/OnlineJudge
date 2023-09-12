package user

import (
	"fmt"
	"main/config"
	"main/server/middleware/redis"
	"main/server/utils/email"
	"math/rand"
	"time"
)

// SendCaptcha 发送验证码
func SendCaptcha(em string) bool {
	// 生成随机验证码
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	captcha := fmt.Sprintf("%d", r.Intn(900000)+100000)

	ch := make(chan error)

	// 发送验证码到邮箱
	go func() {
		ch <- email.SendCaptcha(captcha, em)
	}()

	// 将验证码缓存到Redis
	go func() {
		ctx, cancel := redis.WithTimeoutContext(2)
		defer cancel()

		ch <- redis.Rdb.Set(ctx, redis.GenerateCaptchaKey(em), captcha, time.Duration(config.ConfServer.Email.Expire)*time.Second).Err()
	}()

	return <-ch == nil && <-ch == nil
}

// CheckCaptcha 校验验证码
func CheckCaptcha(email, captche string) (bool, bool) {
	ctx, cancel := redis.WithTimeoutContext(2)
	defer cancel()
	key := redis.GenerateCaptchaKey(email)

	cpt, err := redis.Rdb.Get(ctx, key).Result()
	if err != nil {
		return false, err == redis.Nil
	}

	// 校验完成后将key删除，防止复用
	go redis.Del(key)

	return cpt == captche, true
}
