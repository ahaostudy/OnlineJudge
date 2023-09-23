package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rpcUser "main/api/user"
	"main/internal/common/code"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
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
	result, err := rpc.UserCli.GetCaptcha(c.Request.Context(), &rpcUser.GetCaptchaRequest{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.GetStatusCode())))
}
