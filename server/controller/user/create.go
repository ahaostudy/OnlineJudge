package user

import (
	"main/model"
	"main/server/controller/common"
	"main/server/service/user"
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
		common.Response
	}
)

// CreateUser 创建账号（管理员操作，暂不进行邮箱验证）
func CreateUser(c *gin.Context) {
	req := new(CreateUserRequest)
	res := new(CreateUserResponse)

	// 解析参数
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	// 校验参数是否合法
	req.User.Password = req.Password
	if !user.IsVaildUser(&req.User) {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	// 创建用户
	err := user.CreateUser(&req.User)
	if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeUserExist))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(common.CodeSuccess))
}
