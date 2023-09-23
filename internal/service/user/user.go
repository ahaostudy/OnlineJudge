package user

import (
	"context"

	"github.com/go-sql-driver/mysql"

	rpcUser "main/api/user"
	"main/internal/common/code"
	"main/internal/common/request"
	"main/internal/data/model"
	"main/internal/data/repository"
	"main/internal/middleware/redis"
	"main/internal/service/user/pkg/sha256"
	"main/internal/service/user/pkg/snowflake"
)

func (UserServer) CreateUser(ctx context.Context, req *rpcUser.CreateUserRequest) (resp *rpcUser.CreateUserResponse, _ error) {
	resp = new(rpcUser.CreateUserResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 校验参数是否合法
	// 用户名、邮箱、密码字段不能为空
	// 用户必须是管理员或者普通用户
	if len(req.GetUsername()) == 0 || len(req.GetEmail()) == 0 || len(req.GetPassword()) == 0 ||
		(req.GetRole() != model.ConstRoleOfUser && req.GetRole() != model.ConstRoleOfAdmin) {
		resp.StatusCode = code.CodeInvalidParams.Code()
		return
	}

	// 创建用户
	user := model.User{
		ID:        snowflake.Generate().Int64(),
		Nickname:  req.GetNickname(),
		Username:  req.GetUsername(),
		Password:  sha256.Encrypt(req.GetPassword()),
		Email:     req.GetEmail(),
		Avatar:    req.GetAvatar(),
		Signature: req.GetSignature(),
		Role:      int(req.GetRole()),
	}
	err := repository.InsertUser(&user)
	if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
		resp.StatusCode = code.CodeUserExist.Code()
		return
	}
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

func (us UserServer) UpdateUser(ctx context.Context, req *rpcUser.UpdateUserRequest) (resp *rpcUser.UpdateUserResponse, _ error) {
	resp = new(rpcUser.UpdateUserResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 解析参数
	r := new(request.Request)
	if err := r.ReadRawData(req.GetUser()); err != nil {
		return
	}

	// 判断用户是否越权
	res, err := us.IsAdmin(ctx, &rpcUser.IsAdminRequest{ID: req.GetLoggedInID()})
	if err != nil {
		return
	}
	if res.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = res.StatusCode
		return
	}
	// 如果不是管理员，则不能修改自己的权限或其他人的数据
	if !res.GetIsAdmin() && (r.Exists("role") || req.GetLoggedInID() != req.GetID()) {
		resp.StatusCode = code.CodeForbidden.Code()
		return
	}

	// 更新用户信息
	if pwd, ok := r.Get("password"); ok {
		r.Set("password", sha256.Encrypt(pwd.(string)))
	}
	delete(r.Map(), "id")
	err = repository.UpdateUser(req.GetID(), r.Map())
	if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
		resp.StatusCode = code.CodeUserExist.Code()
		return
	}
	if err != nil {
		return
	}

	// 防止数据不一致，直接清除Redis
	go redis.Del(redis.GenerateAuthKey(req.GetID()))

	resp.StatusCode = code.CodeSuccess.Code()
	return
}
