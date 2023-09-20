package contest

import (
	"main/api/contest"
	"main/internal/common/code"
	"main/internal/common/ctxt"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	RegisterContestRequest struct {
		ActionType int64 `json:"action_type"`
		ContestID  int64 `json:"contest_id"`
	}

	RegisterContestResponse struct {
		ctl.Response
	}
)

func RegisterContest(c *gin.Context) {
	req := new(RegisterContestRequest)
	res := new(RegisterContestResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userID := c.GetInt64("user_id")

	ctx, cancel := ctxt.WithTimeoutContext(2)
	defer cancel()

	// 判断操作类型
	switch req.ActionType {
	case 1:
		// 报名比赛
		result, err := rpc.ContestCli.Register(ctx, &rpcContest.RegisterRequest{
			ContestID: req.ContestID,
			UserID:    userID,
		})
		if err != nil {
			c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
			return
		}
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
	case 2:
		// 取消报名比赛
		result, err := rpc.ContestCli.UnRegister(ctx, &rpcContest.UnRegisterRequest{
			ContestID: req.ContestID,
			UserID:    userID,
		})
		if err != nil {
			c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
			return
		}
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
	default:
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
}
