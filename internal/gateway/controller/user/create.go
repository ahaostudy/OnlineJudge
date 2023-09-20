package user

import (
	"main/internal/data/model"
	"main/internal/common/code"
	"main/internal/gateway/controller/ctl"
	"main/internal/gateway/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type (
	CreateUserRequest struct {
		model.User
		Password string `json:"password"`
	}

	CreateUserResponse struct {
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

	// 校验参数是否合法
	req.User.Password = req.Password
	if !user.IsVaildUser(&req.User) {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 创建用户
	err := user.CreateUser(&req.User)
	if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeUserExist))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.CodeSuccess))
}
