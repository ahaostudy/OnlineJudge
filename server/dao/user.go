package dao

import (
	"main/model"
)

func GetUser(id int64) (*model.User, error) {
	user := new(model.User)
	err := DB.Where("id = ?", id).First(user).Error
	return user, err
}

// GetUserByUsername 根据用户名获取用户对象
func GetUserByUsername(username string) (*model.User, error) {
	user := new(model.User)
	err := DB.Where("username = ?", username).First(user).Error
	return user, err
}

// GetUserByEmail 根据邮箱获取用户对象
func GetUserByEmail(email string) (*model.User, error) {
	user := new(model.User)
	err := DB.Where("email = ?", email).First(user).Error
	return user, err
}

// InsertUser 插入用户信息
func InsertUser(user *model.User) error {
	return DB.Create(user).Error
}

// UpdateUser 更新用户信息，必须指定用户ID，防止修改错误
func UpdateUser(id int64, user map[string]any) error {
	return DB.Model(new(model.User)).Where("id = ?", id).Updates(user).Error
}
