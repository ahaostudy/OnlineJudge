package problem

import (
	"main/internal/common"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/internal/gateway/service/problem"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	CreateProblemRequest struct {
		model.Problem
	}

	CreateProblemResponse struct {
		ctl.Response
	}
)

func CreateProblem(c *gin.Context) {
	req := new(CreateProblemRequest)
	res := new(CreateProblemResponse)

	// 解析参数
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}
	req.AuthorID = c.GetInt64("user_id")

	// 创建问题
	if err := problem.CreateProlem(&req.Problem); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(common.CodeSuccess))
}
