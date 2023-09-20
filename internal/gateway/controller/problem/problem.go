package problem

import (
	rpcProblem "main/api/problem"
	"main/internal/common/code"
	"main/internal/common/build"
	"main/internal/common/ctxt"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	CreateProblemRequest struct {
		model.Problem
	}

	CreateProblemResponse struct {
		ctl.Response
	}

	DeleteProblemResponse struct {
		ctl.Response
	}

	UpdateProblemResponse struct {
		ctl.Response
	}
)

func CreateProblem(c *gin.Context) {
	req := new(CreateProblemRequest)
	res := new(CreateProblemResponse)

	// 解析参数
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	req.AuthorID = c.GetInt64("user_id")

	// 将参数中的问题信息转换为rpc参数
	problem, err := build.BuildProblem(&req.Problem)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 创建问题
	result, err := rpc.ProblemCli.CreateProblem(ctx, &rpcProblem.CreateProblemRequest{Problem: problem})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func DeleteProblem(c *gin.Context) {
	res := new(DeleteProblemResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 删除题目
	result, err := rpc.ProblemCli.DeleteProblem(ctx, &rpcProblem.DeleteProblemRequest{
		ProblemId: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func UpdateProblem(c *gin.Context) {
	res := new(UpdateProblemResponse)

	// 解析参数 id为必须参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	rawData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 更新题目信息
	result, err := rpc.ProblemCli.UpdateProblem(ctx, &rpcProblem.UpdateProblemRequest{
		ProblemId: id,
		Problem:   rawData,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}
