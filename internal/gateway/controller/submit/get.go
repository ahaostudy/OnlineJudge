package submit

import (
	"main/api/submit"
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
	GetSubmitRequest struct {
		ID int64 `json:"id"`
	}

	GetSubmitResponse struct {
		ctl.Response
		StatusCode int           `json:"status_code"`
		StatusMsg  string        `json:"status_msg"`
		Submit     *model.Submit `json:"submit"`
	}
)

func GetSubmit(c *gin.Context) {
	req := new(GetSubmitRequest)
	res := new(GetSubmitResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 获取提交
	result, err := rpc.SubmitCli.GetSubmit(ctx, &rpcSubmit.GetSubmitRequest{
		ID: req.ID,
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
	res.Submit, err = build.UnBuildSubmit(result.GetSubmit())
	if err != nil {
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
