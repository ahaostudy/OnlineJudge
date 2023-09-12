package user

import (
	"main/internal/data/model"
	"main/internal/gateway/dao"
	"main/internal/gateway/utils/sha256"
	"main/internal/gateway/utils/snowflake"
)

// IsVaildUser 校验用户必选是否合理
func IsVaildUser(user *model.User) bool {
	return len(user.Username) > 0 &&
		len(user.Email) > 0 &&
		len(user.Password) > 0 &&
		(user.Role == model.ConstRoleOfUser || user.Role == model.ConstRoleOfAdmin)
}

// CraeteUser 创建用户
func CreateUser(user *model.User) error {
	user.ID = snowflake.Generate().Int64()
	user.Role = model.ConstRoleOfAdmin
	user.Password = sha256.Encrypt(user.Password)
	return dao.InsertUser(user)
}
