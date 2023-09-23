package user

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"main/api/user"
	"main/config"
	"main/internal/common/code"
	"main/internal/middleware/redis"
	"main/internal/service/user/pkg/email"
)

func (UserServer) GetCaptcha(ctx context.Context, req *rpcUser.GetCaptchaRequest) (resp *rpcUser.GetCaptchaResponse, _ error) {
	resp = new(rpcUser.GetCaptchaResponse)

	// 生成随机验证码
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	resp.Captcha = fmt.Sprintf("%d", r.Intn(900000)+100000)

	ch := make(chan error)

	// 发送验证码到邮箱
	go func() {
		ch <- email.SendCaptcha(resp.Captcha, req.GetEmail())
	}()

	// 将验证码缓存到Redis
	go func() {
		ch <- redis.Rdb.Set(ctx, redis.GenerateCaptchaKey(req.GetEmail()), resp.Captcha, time.Duration(config.ConfUser.Email.Expire)*time.Second).Err()
	}()

	if <-ch == nil && <-ch == nil {
		resp.StatusCode = code.CodeSuccess.Code()
	} else {
		resp.StatusCode = code.CodeServerBusy.Code()
	}

	return
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
