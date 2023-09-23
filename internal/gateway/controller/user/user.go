package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"main/api/user"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
)

type (
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
)

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
	result, err := rpc.UserCli.CreateUser(c.Request.Context(), &rpcUser.CreateUserRequest{
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
	result, err := rpc.UserCli.UpdateUser(c.Request.Context(), &rpcUser.UpdateUserRequest{
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
