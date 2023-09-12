package user

import (
	"main/server/controller/common"
	"main/server/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	GetCaptchaRequest struct {
		Email string `json:"email" binding:"required"`
	}

	GetCaptchaResponse struct {
		common.Response
	}
)

func GetCaptcha(c *gin.Context) {
	req := new(GetCaptchaRequest)
	res := new(GetCaptchaResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	// 发送验证码到邮箱
	if ok := user.SendCaptcha(req.Email); !ok {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(common.CodeSuccess))
}
