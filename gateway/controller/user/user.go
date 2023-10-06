package user

import (
	"main/common/code"
	"main/gateway/client"
	"main/gateway/controller/ctl"
	"main/gateway/pkg/model"
	"main/gateway/pkg/pack"
	"main/kitex_gen/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	LoginRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Captcha  string `json:"captcha"`
	}

	LoginResponse struct {
		ctl.Response
		Token  string `json:"token"`
		UserID int64  `json:"user_id"`
	}
	
	RegisterRequest struct {
		Email    string `json:"email" binding:"required"`
		Captcha  string `json:"captcha"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		ctl.Response
		Token  string `json:"token"`
		UserID int64  `json:"user_id"`
	}

	GetUserResponse struct {
		ctl.Response
		User *model.User `json:"user"`
	}

	CreateUserRequest struct {
		model.User
		Password string `json:"password"`
	}

	CreateUserResponse struct {
		ctl.Response
	}

	UpdateUserResponse struct {
		ctl.Response
	}

	GetCaptchaRequest struct {
		Email string `json:"email" binding:"required"`
	}

	GetCaptchaResponse struct {
		ctl.Response
	}
)

func Login(c *gin.Context) {
	req := new(LoginRequest)
	res := new(LoginResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 用户登录
	result, err := client.UserCli.Login(c.Request.Context(), &user.LoginRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Captcha:  req.Captcha,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.GetStatusCode())))
		return
	}

	// success
	res.Token, res.UserID = result.GetToken(), result.GetUserID()
	res.Success()
	c.JSON(http.StatusOK, res)
}

func Register(c *gin.Context) {
	req := new(RegisterRequest)
	res := new(RegisterResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 用户注册
	result, err := client.UserCli.Register(c.Request.Context(), &user.RegisterRequest{
		Email:    req.Email,
		Captcha:  req.Captcha,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.GetStatusCode())))
		return
	}

	// success
	res.Token, res.UserID = result.GetToken(), result.GetUserID()
	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetUser(c *gin.Context) {
	res := new(GetUserResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	if id == 0 {
		id = c.GetInt64("user_id")
	}

	// 获取提交
	result, err := client.UserCli.GetUser(c.Request.Context(), &user.GetUserRequest{
		ID: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.StatusCode != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	// 将结果反编译为模型对象
	res.User, err = pack.UnBuildUser(result.GetUser())
	if err != nil {
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}


// CreateUser 创建账号（管理员操作，暂不进行邮箱验证）
func CreateUser(c *gin.Context) {
	req := new(CreateUserRequest)
	res := new(CreateUserResponse)

	// 解析参数
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 创建用户
	result, err := client.UserCli.CreateUser(c.Request.Context(), &user.CreateUserRequest{
		Nickname:  req.Nickname,
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		Avatar:    req.Avatar,
		Signature: req.Signature,
		Role:      int64(req.Role),
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func UpdateUser(c *gin.Context) {
	res := new(UpdateUserResponse)

	// 解析参数 id为必须参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	rawData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userID := c.GetInt64("user_id")

	// 更新用户
	result, err := client.UserCli.UpdateUser(c.Request.Context(), &user.UpdateUserRequest{
		ID:         id,
		User:       rawData,
		LoggedInID: userID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func GetCaptcha(c *gin.Context) {
	req := new(GetCaptchaRequest)
	res := new(GetCaptchaResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 发送验证码到邮箱
	result, err := client.UserCli.GetCaptcha(c.Request.Context(), &user.GetCaptchaRequest{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.GetStatusCode())))
}
