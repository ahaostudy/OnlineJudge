package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rpcUser "main/api/user"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
)

type (
	GetUserResponse struct {
		ctl.Response
		User *model.User `json:"user"`
	}
)

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
	result, err := rpc.UserCli.GetUser(c.Request.Context(), &rpcUser.GetUserRequest{
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
	res.User, err = build.UnBuildUser(result.GetUser())
	if err != nil {
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
