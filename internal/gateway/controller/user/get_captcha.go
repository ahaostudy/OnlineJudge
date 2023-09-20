package user

import (
	"main/internal/common/code"
	"main/internal/gateway/controller/ctl"
	"main/internal/gateway/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	GetCaptchaRequest struct {
		Email string `json:"email" binding:"required"`
	}

	GetCaptchaResponse struct {
		ctl.Response
	}
)

func GetCaptcha(c *gin.Context) {
	req := new(GetCaptchaRequest)
	res := new(GetCaptchaResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 发送验证码到邮箱
	if ok := user.SendCaptcha(req.Email); !ok {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.CodeSuccess))
}
