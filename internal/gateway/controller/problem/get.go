package problem

import (
	"context"
	rpcProblem "main/api/problem"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	GetProblemRequest struct {
		ID int64
	}

	GetProblemResponse struct {
		ctl.Response
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result, err := rpc.ProblemCli.GetProblem(ctx, &rpcProblem.GetProblemRequest{ProblemId: req.ID})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	// 获取失败
	if result.StatusCode != common.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(common.Code(result.StatusCode)))
		return
	}

	// 将结果转换为题目对象
	problem, err := build.UnBuildProblem(result.Problem)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}
	res.Problem = problem

	res.Success()
	c.JSON(http.StatusOK, res)
}
