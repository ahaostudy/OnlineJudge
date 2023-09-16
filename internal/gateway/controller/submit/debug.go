package submit

import (
	"main/api/judge"
	"main/api/submit"
	"main/internal/common"
	"main/internal/common/ctxt"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	DebugRequest struct {
		ID     int64  `json:"id"`
		Code   []byte `json:"code"`
		Input  []byte `json:"input"`
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
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 使用给定的代码、输入和语言ID调用 JudgeCli 的 Debug 方法
	result, err := rpc.SubmitCli.Debug(ctx, &rpcSubmit.DebugReqeust{
		Code:   req.Code,
		Input:  req.Input,
		LangID: req.LangID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	// 将响应结果和状态码设置为来自 JudgeCli 响应的值
	res.CodeOf(common.Code(result.StatusCode))
	res.Result = result.Result
	c.JSON(http.StatusOK, res)
}
