package problem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rpcProblem "main/api/problem"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
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
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 获取样例数据
	result, err := rpc.ProblemCli.GetTestcase(c.Request.Context(), &rpcProblem.GetTestcaseRequest{
		ID: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.StatusCode != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	// 将结果反编译为模型对象
	res.Testcase, err = build.UnBuildTestcase(result.GetTestcase())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
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
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	switch req.ActionType {
	case 1:
		// 创建题目
		result, err := rpc.ProblemCli.CreateTestcase(c.Request.Context(), &rpcProblem.CreateTestcaseRequest{
			ProblemID: req.ProblemID,
			Input:     []byte(req.Input),
			Output:    []byte(req.Output),
		})
		if err != nil {
			c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
			return
		}

		// 响应结果
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
	default:
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
	}
}

func DeleteTestcase(c *gin.Context) {
	res := new(DeleteTestcaseResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 删除样例
	result, err := rpc.ProblemCli.DeleteTestcase(c.Request.Context(), &rpcProblem.DeleteTestcaseRequest{
		ID: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}
