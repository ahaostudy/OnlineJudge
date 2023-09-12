package problem

import (
	"main/internal/gateway/controller/common"
	"main/internal/gateway/service/problem"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	UpdateProblemRequest struct {
		ID int64
		common.Request
	}

	UpdateProblemResponse struct {
		common.Response
	}
)

func UpdateProblem(c *gin.Context) {
	req := new(UpdateProblemRequest)
	res := new(UpdateProblemResponse)

	// 解析参数 id为必须参数
	req.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	if err := req.ReadRawData(c); req.ID == 0 || err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	// 更新题目信息
	if err := problem.UpdateProblem(req.ID, req.Map()); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(common.CodeSuccess))
}
