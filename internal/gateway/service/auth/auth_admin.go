package auth

import (
	"errors"
	"main/internal/data/model"
	"main/internal/gateway/dao"
	"main/internal/gateway/middleware/redis"
	"strconv"

	"gorm.io/gorm"
)

// IsAdmin 判断用户是否为管理员
func IsAdmin(id int64) (isAdmin bool, ok bool) {
	ctx, cancel := redis.WithTimeoutContext(2)
	defer cancel()
	key := redis.GenerateAuthKey(id)

	// 先尝试从Redis中用户的权限，获取到直接返回（同时更新过期时间）
	val, err := redis.Rdb.Get(ctx, key).Result()
	if err == nil {
		isAdmin, err = strconv.ParseBool(val)
		if err == nil {
			go redis.Flush(key)
			return isAdmin, true
		}
	}

	// 未命中缓存，回源获取
	user, err := dao.GetUser(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		isAdmin, ok = false, true
	} else {
		isAdmin, ok = user.Role == model.ConstRoleOfAdmin, err == nil
	}

	// 无异常则并发缓存到数据库
	if ok {
		go redis.Set(key, isAdmin)
	}

	return isAdmin, ok
}
