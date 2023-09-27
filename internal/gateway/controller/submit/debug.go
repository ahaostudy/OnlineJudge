package submit

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rpcSubmit "main/api/submit"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/gateway/controller/ctl"
	cr "main/internal/service/judge/pkg/code"
	"main/rpc"
)

type (
	DebugRequest struct {
		Code   string `json:"code"`
		Input  string `json:"input"`
		LangID int64  `json:"lang_id"`
	}

	DebugResponse struct {
		ctl.Response
		Result *cr.Result `json:"result"`
	}
)

func Debug(c *gin.Context) {
	req := new(DebugRequest)
	res := new(DebugResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 使用给定的代码、输入和语言ID调用 SubmitCli 的 Debug 方法
	result, err := rpc.SubmitCli.Debug(c.Request.Context(), &rpcSubmit.DebugReqeust{
		Code:   []byte(req.Code),
		Input:  []byte(req.Input),
		LangID: req.LangID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		res.CodeOf(code.Code(result.StatusCode))
		return
	}

	// 将响应结果和状态码设置为来自 SubmitCli 响应的值
	res.Result, err = build.UnBuildResult(result.Result)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
