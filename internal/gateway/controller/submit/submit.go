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
	SubmitRequest struct {
		ProblemID int64  `json:"problem_id"`
		Code      string `json:"code"`
		LangID    int64  `json:"lang_id"`
	}

	SubmitResponse struct {
		ctl.Response
		SubmitID int64 `json:"submit_id"`
	}
)

func Submit(c *gin.Context) {
	req := new(SubmitRequest)
	res := new(SubmitResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 提交判题
	result, err := rpc.SubmitCli.Submit(ctx, &rpcSubmit.SubmitRequest{
		ProblemID: req.ProblemID,
		UserID:    c.GetInt64("user_id"),
		Code:      []byte(req.Code),
		LangID:    req.LangID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	res.SubmitID = result.GetSubmitID()
	res.Response = res.CodeOf(common.Code(result.StatusCode))
	c.JSON(http.StatusOK, res)
}
