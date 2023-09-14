package testcase

import (
	rpcTestcase "main/api/testcase"
	"main/internal/common"
	"main/internal/common/ctxt"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	DeleteTestcaseResponse struct {
		ctl.Response
	}
)

func DeleteTestcase(c *gin.Context) {
	res := new(DeleteTestcaseResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 删除样例
	result, err := rpc.TestcaseCli.DeleteTestcase(ctx, &rpcTestcase.DeleteTestcaseRequest{
		ID: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(common.Code(result.StatusCode)))
}
