package submit

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rpcSubmit "main/api/submit"
	"main/internal/common/code"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
)

type (
	SubmitRequest struct {
		ProblemID int64  `json:"problem_id"`
		Code      string `json:"code"`
		LangID    int64  `json:"lang_id"`
		ContestID int64  `json:"contest_id"`
	}

	SubmitResponse struct {
		ctl.Response
		SubmitID int64 `json:"submit_id"`
	}

	DeleteSubmitResponse struct {
		ctl.Response
	}
)

func Submit(c *gin.Context) {
	req := new(SubmitRequest)
	res := new(SubmitResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	var submitID, statusCode int64
	var err error
	if req.ContestID == 0 {
		// 提交判题
		result, e := rpc.SubmitCli.Submit(c.Request.Context(), &rpcSubmit.SubmitRequest{
			ProblemID: req.ProblemID,
			UserID:    c.GetInt64("user_id"),
			Code:      []byte(req.Code),
			LangID:    req.LangID,
		})
		submitID, statusCode, err = result.GetSubmitID(), result.GetStatusCode(), e
	} else {
		result, e := rpc.SubmitCli.SubmitContest(c.Request.Context(), &rpcSubmit.SubmitContestRequest{
			ProblemID: req.ProblemID,
			UserID:    c.GetInt64("user_id"),
			Code:      []byte(req.Code),
			LangID:    req.LangID,
			ContestID: req.ContestID,
		})
		submitID, statusCode, err = result.GetSubmitID(), result.GetStatusCode(), e
	}
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.SubmitID = submitID
	res.Response = res.CodeOf(code.Code(statusCode))
	c.JSON(http.StatusOK, res)
}

type ()

func DeleteSubmit(c *gin.Context) {
	res := new(DeleteSubmitResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userID := c.GetInt64("user_id")

	// 删除一条记录
	result, err := rpc.SubmitCli.DeleteSubmit(c.Request.Context(), &rpcSubmit.DeleteSubmitRequest{
		ID:     id,
		UserID: userID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}
