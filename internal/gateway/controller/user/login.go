package user

import (
	"errors"
	"main/internal/data/model"
	"main/internal/gateway/controller/common"
	"main/internal/gateway/middleware/jwt"
	"main/internal/gateway/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	LoginRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Captcha  string `json:"captcha"`
	}

	LoginResponse struct {
		common.Response
		Token  string `json:"token"`
		UserID int64  `json:"user_id"`
	}
)

func Login(c *gin.Context) {
	req := new(LoginRequest)
	res := new(LoginResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}
	// 校验参数是否合理
	if len(req.Username) == 0 && len(req.Email) == 0 {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}
	if len(req.Password) == 0 && len(req.Captcha) == 0 {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	var u *model.User
	if len(req.Password) > 0 {
		// 密码登录
		var err error
		u, err = user.LoginByPassword(req.Username, req.Email, req.Password)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, res.CodeOf(common.CodeUserNotExist))
			return
		}
		if err != nil {
			c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidPassword))
			return
		}
	} else {
		// 验证码登录
		var ok bool
		u, ok = user.LoginByCaptcha(req.Email, req.Captcha)
		if !ok {
			c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidCaptcha))
			return
		}
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
