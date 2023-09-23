package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"main/api/user"
	"main/internal/common/code"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
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
	result, err := rpc.UserCli.Login(c.Request.Context(), &rpcUser.LoginRequest{
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
	result, err := rpc.UserCli.Register(c.Request.Context(), &rpcUser.RegisterRequest{
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
