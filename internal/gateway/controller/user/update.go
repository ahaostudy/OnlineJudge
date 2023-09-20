package user

import (
	"main/internal/common/code"
	"main/internal/gateway/controller/ctl"
	"main/internal/gateway/service/auth"
	"main/internal/gateway/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	UpdateUserRequest struct {
		ctl.Request
		ID int64
	}

	UpdateUserResponse struct {
		ctl.Response
	}
)

func UpdateUser(c *gin.Context) {
	req := new(UpdateUserRequest)
	res := new(UpdateUserResponse)
	userID := c.GetInt64("user_id")

	// 解析参数 id为必须参数
	req.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	if err := req.ReadRawData(c); req.ID == 0 || err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 获取用户权限，判断用户是否越权
	isAdmin, ok := auth.IsAdmin(userID)
	if !ok {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if !isAdmin && (!req.Exists("role") || userID != req.ID) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	// 更新用户信息
	if err := user.UpdateUser(req.ID, req.Map()); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.CodeSuccess))
}
