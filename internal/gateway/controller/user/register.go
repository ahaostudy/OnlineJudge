package user

import (
	"main/internal/common"
	"main/internal/gateway/controller/ctl"
	"main/internal/gateway/middleware/jwt"
	"main/internal/gateway/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type (
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

func Register(c *gin.Context) {
	req := new(RegisterRequest)
	res := new(RegisterResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	// 校验验证码
	valid, ok := user.CheckCaptcha(req.Email, req.Captcha)
	if !ok || !valid {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidCaptcha))
		return
	}

	// 创建用户
	// 判断用户是否存在 (Error 1062: Duplicate entry)
	u, err := user.Register(req.Email, req.Password)
	if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeUserExist))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	// 生成token
	token, err := jwt.GenerateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	// success
	res.Token, res.UserID = token, u.ID
	res.Success()
	c.JSON(http.StatusOK, res)
}
