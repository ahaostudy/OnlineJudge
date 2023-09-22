package submit

import (
	rpcSubmit "main/api/submit"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/common/ctxt"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	GetResultRequest struct {
		SubmitID int64 `json:"submit_id"`
	}

	GetResultResponse struct {
		ctl.Response
		Result any `json:"result"`
	}
)

func GetResult(c *gin.Context) {
	req := new(GetResultRequest)
	res := new(GetResultResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 获取提交结果
	result, err := rpc.SubmitCli.GetSubmitResult(ctx, &rpcSubmit.GetSubmitResultRequest{SubmitID: req.SubmitID})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.GetStatusCode())))
		return
	}

	// 将响应转换为模型对象
	res.Result, err = build.UnBuildResult(result.GetResult())
	if err != nil {
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
