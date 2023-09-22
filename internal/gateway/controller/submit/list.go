package submit

import (
	"main/api/submit"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/common/ctxt"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	GetSubmitListRequest struct {
		UserID    int64 `json:"user_id"`
		ProblemID int64 `json:"problem_id"`
	}

	GetSubmitListResponse struct {
		ctl.Response
		SubmitList []*model.Submit `json:"submit_list"`
	}
)

func GetSubmitList(c *gin.Context) {
	req := new(GetSubmitListRequest)
	res := new(GetSubmitListResponse)

	// 解析参数
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 获取提交数据
	result, err := rpc.SubmitCli.GetSubmitList(ctx, &rpcSubmit.GetSubmitListRequest{
		UserID:    req.UserID,
		ProblemID: req.ProblemID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.GetStatusCode())))
		return
	}

	// 将结果反编译为模型对象
	submitList, err := build.UnBuildSubmitList(result.GetSubmitList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.SubmitList = submitList
	res.Success()
	c.JSON(http.StatusOK, res)
}
