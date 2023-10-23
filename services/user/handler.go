package user

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"main/common/code"
	"main/common/jwt"
	"main/common/raw"
	user "main/kitex_gen/user"
	"main/services/user/config"
	"main/services/user/dal/cache"
	"main/services/user/dal/db"
	"main/services/user/dal/model"
	"main/services/user/pack"
	"main/services/user/pkg/email"
	"main/services/user/pkg/sha256"
	"main/services/user/pkg/snowflake"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, _ error) {
	resp = new(user.RegisterResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 校验验证码
	valid, ok := CheckCaptcha(ctx, req.GetEmail(), req.GetCaptcha())
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
	err := db.InsertUser(&user)
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

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, _ error) {
	resp = new(user.LoginResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	fmt.Printf("req: %#v\n", req)

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
		user, err = db.GetUserByUsername(req.GetUsername())
		if len(req.GetEmail()) > 0 && errors.Is(err, gorm.ErrRecordNotFound) {
			user, err = db.GetUserByEmail(req.GetEmail())
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
		user, err = db.GetUserByEmail(req.GetEmail())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.StatusCode = code.CodeUserNotExist.Code()
			return
		}
		if err != nil {
			return
		}

		// 校验验证码是否正确
		vaild, ok := CheckCaptcha(ctx, req.GetEmail(), req.GetCaptcha())
		if !ok || !vaild {
			resp.StatusCode = code.CodeInvalidCaptcha.Code()
			return
		}
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

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, _ error) {
	resp = new(user.CreateUserResponse)
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
	u := model.User{
		ID:        snowflake.Generate().Int64(),
		Nickname:  req.GetNickname(),
		Username:  req.GetUsername(),
		Password:  sha256.Encrypt(req.GetPassword()),
		Email:     req.GetEmail(),
		Avatar:    req.GetAvatar(),
		Signature: req.GetSignature(),
		Role:      int(req.GetRole()),
	}
	err := db.InsertUser(&u)
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

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, _ error) {
	resp = new(user.UpdateUserResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 解析参数
	r := new(raw.Raw)
	if err := r.ReadRawData(req.GetUser()); err != nil {
		return
	}

	// 判断用户是否越权
	res, err := s.IsAdmin(ctx, &user.IsAdminRequest{ID: req.GetLoggedInID()})
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
	err = db.UpdateUser(req.GetID(), r.Map())
	if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
		resp.StatusCode = code.CodeUserExist.Code()
		return
	}
	if err != nil {
		return
	}

	// 防止数据不一致，直接清除Redis
	go cache.Del(cache.GenerateAuthKey(req.GetID()))

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetCaptcha implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetCaptcha(ctx context.Context, req *user.GetCaptchaRequest) (resp *user.GetCaptchaResponse, _ error) {
	resp = new(user.GetCaptchaResponse)

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
		ch <- cache.Rdb.Set(ctx, cache.GenerateCaptchaKey(req.GetEmail()), resp.Captcha, time.Duration(config.Config.Email.Expire)*time.Second).Err()
	}()

	if <-ch == nil && <-ch == nil {
		resp.StatusCode = code.CodeSuccess.Code()
	} else {
		resp.StatusCode = code.CodeServerBusy.Code()
	}

	return
}

// IsAdmin implements the UserServiceImpl interface.
func (s *UserServiceImpl) IsAdmin(ctx context.Context, req *user.IsAdminRequest) (resp *user.IsAdminResponse, _ error) {
	resp = new(user.IsAdminResponse)
	resp.StatusCode = code.CodeServerBusy.Code()
	key := cache.GenerateAuthKey(req.GetID())

	// 先尝试从Redis中用户的权限，获取到直接返回（同时更新过期时间）
	val, err := cache.Rdb.Get(ctx, key).Result()
	if err == nil {
		resp.IsAdmin, err = strconv.ParseBool(val)
		if err == nil {
			resp.StatusCode = code.CodeSuccess.Code()
			go cache.Flush(key)
			return
		}
	}

	// 未命中缓存，回源获取
	user, err := db.GetUser(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.IsAdmin = false
	} else if err != nil {
		return
	} else {
		resp.IsAdmin = user.Role == model.ConstRoleOfAdmin
	}

	// 无异常则并发缓存到数据库
	go cache.Set(key, resp.IsAdmin)
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, _ error) {
	resp = new(user.GetUserResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	var u *model.User
	var err error
	if req.GetID() != 0 {
		u, err = db.GetUser(req.GetID())
	} else {
		u, err = db.GetUserByUsername(req.GetUsername())
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeUserNotExist.Code()
		return
	}
	if err != nil {
		return
	}

	resp.User, err = pack.BuildUser(u)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// CheckCaptcha 校验验证码
func CheckCaptcha(ctx context.Context, email, captche string) (bool, bool) {
	key := cache.GenerateCaptchaKey(email)

	cpt, err := cache.Rdb.Get(ctx, key).Result()
	if err != nil {
		return false, err == cache.Nil
	}

	// 校验完成后将key删除，防止复用
	go cache.Del(key)

	return cpt == captche, true
}

// UploadAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequest) (resp *user.UploadAvatarResponse, _ error) {
	resp = new(user.UploadAvatarResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 将原头像删除（如果存在的话）
	u, err := db.GetUser(req.GetUserID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeUserNotExist.Code()
		return
	}
	if err != nil {
		return
	}
	if u.Avatar != "" {
		_ = os.Remove(filepath.Join(config.Config.Static.Path, u.Avatar))
	}

	// 保存新的头像
	fileName := uuid.New().String() + req.GetExt()
	os.MkdirAll(config.Config.Static.Path, 0755)
	if err := os.WriteFile(filepath.Join(config.Config.Static.Path, fileName), req.GetBody(), 0644); err != nil {
		return
	}

	// 更新数据库
	if err := db.UpdateUser(req.GetUserID(), map[string]any{
		"avatar": fileName,
	}); err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// DownloadAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) DownloadAvatar(ctx context.Context, req *user.DownloadAvatarRequest) (resp *user.DownloadAvatarResponse, _ error) {
	resp = new(user.DownloadAvatarResponse)
	resp.StatusCode = code.CodeServerBusy.Code()
	path := filepath.Join(config.Config.Static.Path, req.GetAvatar())

	if _, err := os.Stat(path); os.IsNotExist(err) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}

	body, err := os.ReadFile(path)
	if err != nil {
		return
	}

	resp.Body = body
	resp.StatusCode = code.CodeSuccess.Code()
	return
}
