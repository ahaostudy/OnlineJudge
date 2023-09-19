package problem

import (
	"context"
	rpcProblem "main/api/problem"
	"main/internal/common"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	CreateTestcaseRequest struct {
		ActionType int    `json:"action_type"` // 操作类型：1. 使用文本 2. 使用文件
		ProblemID  int64  `json:"problem_id"`
		Input      string `json:"input"`  // 输入文本
		Output     string `json:"output"` // 输出文本
	}

	CreateTestcaseResponse struct {
		ctl.Response
	}
)

// 添加测试样例
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
