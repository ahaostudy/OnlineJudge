package submit

import (
	"main/api/submit"
	"main/internal/common/code"
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
	GetSubmitResponse struct {
		ctl.Response
		StatusCode int           `json:"status_code"`
		StatusMsg  string        `json:"status_msg"`
		Submit     *model.Submit `json:"submit"`
	}
)

func GetSubmit(c *gin.Context) {
	res := new(GetSubmitResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 获取提交
	result, err := rpc.SubmitCli.GetSubmit(ctx, &rpcSubmit.GetSubmitRequest{
		ID: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.StatusCode != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
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
