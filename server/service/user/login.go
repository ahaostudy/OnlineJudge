package user

import (
	"errors"
	"main/model"
	"main/server/dao"
	"main/server/utils/sha256"
)

// LoginByPassword 密码登录
func LoginByPassword(username, email, password string) (*model.User, error) {
	// 获取用户信息
	var user *model.User
	var err error
	if len(username) > 0 {
		user, err = dao.GetUserByUsername(username)
	} else {
		user, err = dao.GetUserByEmail(email)
	}
	if err != nil {
		return nil, err
	}

	// 校验密码
	if sha256.Encrypt(password) != user.Password {
		return nil, errors.New("password verification failed")
	}

	return user, nil
}

// LoginByCaptcha 验证码登录
func LoginByCaptcha(email, captcha string) (*model.User, bool) {
	vaild, ok := CheckCaptcha(email, captcha)
	if !ok || !vaild {
		return nil, false
	}

	user, err := dao.GetUserByEmail(email)

	return user, err == nil
}
