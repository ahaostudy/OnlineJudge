package problem

import (
	"main/internal/common"
	"main/internal/gateway/controller/ctl"
	"main/internal/gateway/service/problem"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	DeleteProblemRequest struct {
		ID int64
	}

	DeleteProblemResponse struct {
		ctl.Response
	}
)

func DeleteProblem(c *gin.Context) {
	req := new(DeleteProblemRequest)
	res := new(DeleteProblemResponse)

	// 解析参数
	req.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	if req.ID == 0 {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	// 删除题目
	if err := problem.DeleteProblem(req.ID); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(common.CodeSuccess))
}
