package submit

import (
	rpcSubmit "main/api/submit"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/common/ctxt"
	"main/internal/data/model"
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
		Result *model.Submit `json:"result"`
	}
)

func GetResult(c *gin.Context) {
	req := new(GetResultRequest)
	res := new(GetResultResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 获取提交结果
	result, err := rpc.SubmitCli.GetSubmitResult(ctx, &rpcSubmit.GetSubmitResultRequest{SubmitID: req.SubmitID})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != common.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(common.Code(result.GetStatusCode())))
		return
	}

	// 将响应转换为模型对象
	res.Result, err = build.UnBuildSubmit(result.GetResult())
	if err != nil {
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
