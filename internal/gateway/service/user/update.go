package user

import (
	"fmt"
	"main/internal/gateway/dao"
	"main/internal/middleware/redis"
	"main/internal/gateway/utils/sha256"
)

func UpdateUser(id int64, user map[string]any) error {
	if pwd, ok := user["password"]; ok {
		user["password"] = sha256.Encrypt(pwd.(string))
	}
	delete(user, "id")
	fmt.Printf("user: %#v\n", user)

	// 防止数据不一直，直接清除Redis
	go redis.Del(redis.GenerateAuthKey(id))

	return dao.UpdateUser(id, user)
}
