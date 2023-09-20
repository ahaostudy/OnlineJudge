package problem

import (
	"context"
	rpcProblem "main/api/problem"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/common/ctxt"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	GetTestcaseResponse struct {
		ctl.Response
		Testcase *model.Testcase `json:"testcase"`
	}

	CreateTestcaseRequest struct {
		ActionType int    `json:"action_type"` // 操作类型：1. 使用文本 2. 使用文件
		ProblemID  int64  `json:"problem_id"`
		Input      string `json:"input"`  // 输入文本
		Output     string `json:"output"` // 输出文本
	}

	CreateTestcaseResponse struct {
		ctl.Response
	}

	DeleteTestcaseResponse struct {
		ctl.Response
	}
)

func GetTestcase(c *gin.Context) {
	res := new(GetTestcaseResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 获取样例数据
	result, err := rpc.ProblemCli.GetTestcase(ctx, &rpcProblem.GetTestcaseRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}
	if result.StatusCode != common.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(common.Code(result.StatusCode)))
		return
	}

	// 将结果反编译为模型对象
	res.Testcase, err = build.UnBuildTestcase(result.GetTestcase())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func CreateTestcase(c *gin.Context) {
	req := new(CreateTestcaseRequest)
	res := new(CreateTestcaseResponse)

	// 解析参数
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	switch req.ActionType {
	case 1:
		// 创建题目
		result, err := rpc.ProblemCli.CreateTestcase(ctx, &rpcProblem.CreateTestcaseRequest{
			ProblemId: req.ProblemID,
			Input:     []byte(req.Input),
			Output:    []byte(req.Output),
		})
		if err != nil {
			c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
			return
		}

		// 响应结果
		c.JSON(http.StatusOK, res.CodeOf(common.Code(result.StatusCode)))
	default:
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
	}
}

func DeleteTestcase(c *gin.Context) {
	res := new(DeleteTestcaseResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 删除样例
	result, err := rpc.ProblemCli.DeleteTestcase(ctx, &rpcProblem.DeleteTestcaseRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(common.Code(result.StatusCode)))
}
