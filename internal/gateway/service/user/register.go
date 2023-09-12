package user

import (
	"main/internal/data/model"
	"main/internal/gateway/dao"
	"main/internal/gateway/utils/sha256"
	"main/internal/gateway/utils/snowflake"
)

// Register 注册用户
func Register(email, password string) (*model.User, error) {
	// 根据参数创建用户
	// 使用雪花算法生成一个用户ID，角色设置为普通用户
	id := snowflake.Generate().Int64()
	user := &model.User{
		ID:       id,
		Email:    email,
		Nickname: email,
		Username: email,
		Password: sha256.Encrypt(password),
		Role:     model.ConstRoleOfUser,
	}

	// 插入用户到数据库
	err := dao.InsertUser(user)

	return user, err
}
