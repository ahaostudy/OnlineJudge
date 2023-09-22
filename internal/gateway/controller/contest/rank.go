package contest

import (
	rpcContest "main/api/contest"
	"main/internal/common/code"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ContestRankReqest struct {
		ContestID int64 `json:"contest_id"`
		Page      int64 `json:"page"`
		Count     int64 `json:"count"`
	}

	ContestRankResponse struct {
		ctl.Response
		Rank []*rpcContest.UserData `json:"rank"`
	}
)

func ContestRank(c *gin.Context) {
	req := new(ContestRankReqest)
	res := new(ContestRankResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	if req.Page == 0 {
		req.Page = int64(defaultPage)
	}
	if req.Count == 0 {
		req.Count = int64(defaultCount)
	}

	// 获取比赛排名
	result, err := rpc.ContestCli.ContestRank(c.Request.Context(), &rpcContest.ContestRankRequest{
		ContestID: req.ContestID,
		Page:      req.Page,
		Count:     req.Count,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.CodeOf(code.Code(result.StatusCode))
	res.Rank = result.Rank
	c.JSON(http.StatusOK, res)
}
