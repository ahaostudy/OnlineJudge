package submit

import (
	rpcSubmit "main/api/submit"
	"main/internal/common"
	"main/internal/common/ctxt"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	DeleteSubmitRequest struct {
		ID int64 `json:"id"`
	}

	DeleteSubmitResponse struct {
		ctl.Response
	}
)

func DeleteSubmit(c *gin.Context) {
	req := new(DeleteSubmitRequest)
	res := new(DeleteSubmitResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 删除一条记录
	result, err := rpc.SubmitCli.DeleteSubmit(ctx, &rpcSubmit.DeleteSubmitRequest{ID: req.ID})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}
	if result.StatusCode != common.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(common.Code(result.StatusCode)))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
