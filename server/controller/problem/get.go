package problem

import (
	"errors"
	"main/model"
	"main/server/controller/common"
	"main/server/service/problem"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	GetProblemRequest struct {
		ID int64
	}

	GetProblemResponse struct {
		common.Response
		Problem *model.Problem `json:"problem"`
	}
)

func GetProblem(c *gin.Context) {
	req := new(GetProblemRequest)
	res := new(GetProblemResponse)

	// 解析参数
	req.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	if req.ID == 0 {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	// 获取题目
	problem, err := problem.GetProblem(req.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeRecordNotFound))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	res.Problem = problem
	res.Success()
	c.JSON(http.StatusOK, res)
}
