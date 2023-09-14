package testcase

import (
	rpcTestcase "main/api/testcase"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/common/ctxt"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	GetTestcaseResponse struct {
		ctl.Response
		Testcase *model.Testcase `json:"testcase"`
	}
)

func GetTestcase(c *gin.Context) {
	res := new(GetTestcaseResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 获取样例数据
	result, err := rpc.TestcaseCli.GetTestcase(ctx, &rpcTestcase.GetTestcaseRequest{
		ID: id,
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
	res.Testcase, err = build.UnBuildTestcase(result.GetTestcase())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
