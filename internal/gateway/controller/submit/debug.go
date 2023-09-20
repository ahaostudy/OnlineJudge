package submit

import (
	"main/api/judge"
	"main/api/submit"
	"main/internal/common/code"
	"main/internal/common/ctxt"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	DebugRequest struct {
		Code   string `json:"code"`
		Input  string `json:"input"`
		LangID int64  `json:"lang_id"`
	}

	DebugResponse struct {
		ctl.Response
		Result *rpcJudge.JudgeResult `json:"result"`
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

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 使用给定的代码、输入和语言ID调用 SubmitCli 的 Debug 方法
	result, err := rpc.SubmitCli.Debug(ctx, &rpcSubmit.DebugReqeust{
		Code:   []byte(req.Code),
		Input:  []byte(req.Input),
		LangID: req.LangID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	// 将响应结果和状态码设置为来自 SubmitCli 响应的值
	res.CodeOf(code.Code(result.StatusCode))
	res.Result = result.Result
	c.JSON(http.StatusOK, res)
}
