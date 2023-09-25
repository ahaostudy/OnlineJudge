package user

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	rpcUser "main/api/user"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/data/repository"
	"main/internal/gateway/middleware/jwt"
	"main/internal/middleware/redis"
	"main/internal/service/user/pkg/sha256"
	"main/internal/service/user/pkg/snowflake"
)

func (UserServer) Login(ctx context.Context, req *rpcUser.LoginRequest) (resp *rpcUser.LoginResponse, _ error) {
	resp = new(rpcUser.LoginResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 校验参数是否合理
	if len(req.GetUsername()) == 0 && len(req.GetEmail()) == 0 {
		resp.StatusCode = code.CodeInvalidParams.Code()
		return
	}
	if len(req.Password) == 0 && len(req.Captcha) == 0 {
		resp.StatusCode = code.CodeInvalidParams.Code()
		return
	}

	var user *model.User
	var err error
	if len(req.Password) > 0 {
		// 获取用户信息
		if len(req.GetUsername()) > 0 {
			user, err = repository.GetUserByUsername(req.GetUsername())
		}
		if len(req.GetEmail()) > 0 && errors.Is(err, gorm.ErrRecordNotFound) {
			user, err = repository.GetUserByEmail(req.GetEmail())
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.StatusCode = code.CodeUserNotExist.Code()
			return
		}
		if err != nil {
			return
		}

		// 校验密码
		if sha256.Encrypt(req.GetPassword()) != user.Password {
			resp.StatusCode = code.CodeInvalidPassword.Code()
			return
		}
	} else {
		// 通过邮箱获取用户，判断用户是否存在
		user, err = repository.GetUserByEmail(req.GetEmail())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.StatusCode = code.CodeUserNotExist.Code()
			return
		}
		if err != nil {
			return
		}

		// 校验验证码是否正确
		vaild, ok := CheckCaptcha(req.GetEmail(), req.GetCaptcha())
		if !ok || !vaild {
			resp.StatusCode = code.CodeInvalidCaptcha.Code()
			return
		}
	}

	if user == nil {
		resp.StatusCode = code.CodeServerBusy.Code()
		return
	}

	// 生成token
	resp.Token, err = jwt.GenerateToken(user.ID)
	if err != nil {
		return
	}

	// success
	resp.UserID = user.ID
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (UserServer) Register(ctx context.Context, req *rpcUser.RegisterRequest) (resp *rpcUser.RegisterResponse, _ error) {
	resp = new(rpcUser.RegisterResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 校验验证码
	valid, ok := CheckCaptcha(req.GetEmail(), req.GetCaptcha())
	if !ok || !valid {
		resp.StatusCode = code.CodeInvalidCaptcha.Code()
		return
	}

	// 创建用户
	id := snowflake.Generate().Int64()
	user := model.User{
		ID:       id,
		Email:    req.GetEmail(),
		Nickname: req.GetEmail(),
		Username: req.GetEmail(),
		Password: sha256.Encrypt(req.Password),
		Role:     model.ConstRoleOfUser,
	}
	err := repository.InsertUser(&user)
	// 判断用户是否已存在 (Error 1062: Duplicate entry)
	if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
		resp.StatusCode = code.CodeUserExist.Code()
		return
	}
	if err != nil {
		return
	}

	// 生成token
	resp.Token, err = jwt.GenerateToken(user.ID)
	if err != nil {
		return
	}

	resp.UserID = user.ID
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (UserServer) IsAdmin(ctx context.Context, req *rpcUser.IsAdminRequest) (resp *rpcUser.IsAdminResponse, _ error) {
	resp = new(rpcUser.IsAdminResponse)
	resp.StatusCode = code.CodeServerBusy.Code()
	key := redis.GenerateAuthKey(req.GetID())

	// 先尝试从Redis中用户的权限，获取到直接返回（同时更新过期时间）
	val, err := redis.Rdb.Get(ctx, key).Result()
	if err == nil {
		resp.IsAdmin, err = strconv.ParseBool(val)
		if err == nil {
			resp.StatusCode = code.CodeSuccess.Code()
			go redis.Flush(key)
			return
		}
	}

	// 未命中缓存，回源获取
	user, err := repository.GetUser(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.IsAdmin = false
	} else if err != nil {
		return
	} else {
		resp.IsAdmin = user.Role == model.ConstRoleOfAdmin
	}

	// 无异常则并发缓存到数据库
	go redis.Set(key, resp.IsAdmin)
	resp.StatusCode = code.CodeSuccess.Code()
	return
}
