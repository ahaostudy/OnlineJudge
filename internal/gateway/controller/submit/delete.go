package submit

import (
	rpcSubmit "main/api/submit"
	"main/internal/common"
	"main/internal/common/ctxt"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	DeleteSubmitResponse struct {
		ctl.Response
	}
)

func DeleteSubmit(c *gin.Context) {
	res := new(DeleteSubmitResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}
	userID := c.GetInt64("user_id")

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 删除一条记录
	result, err := rpc.SubmitCli.DeleteSubmit(ctx, &rpcSubmit.DeleteSubmitRequest{
		ID:     id,
		UserID: userID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(common.Code(result.StatusCode)))
}
