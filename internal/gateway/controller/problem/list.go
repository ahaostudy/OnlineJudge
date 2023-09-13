package problem

import (
	"main/internal/common"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/internal/gateway/service/problem"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 默认每页的数量
const (
	defaultPage  = 1
	defaultCount = 20
)

type (
	GetProblemListResponse struct {
		ctl.Response
		ProblemList []*model.Problem `json:"problem_list"`
	}
)

func GetProblemList(c *gin.Context) {
	res := new(GetProblemListResponse)

	// 解析参数，不适用BindQuery是因为要添加默认值
	page, _ := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	count, _ := strconv.Atoi(c.DefaultQuery("count", strconv.Itoa(defaultCount)))

	// 获取题目列表
	problems, err := problem.GetProblemList(page, count)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	// 响应结果
	res.ProblemList = problems
	res.Success()
	c.JSON(http.StatusOK, res)
}
