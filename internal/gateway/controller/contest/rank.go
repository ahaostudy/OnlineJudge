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

	Status struct {
		ProblemID int64 `json:"problem_id"`
		Penalty   int64 `json:"penalty"`
		Accepted  bool  `json:"accepted"`
		AcTime    int64 `json:"ac_time"`
		LangID    int64 `json:"lang_id"`
		Score     int64 `json:"score"`
	}

	UserData struct {
		UserID int64     `json:"user_id"`
		Status []*Status `json:"status"`
	}

	ContestRankResponse struct {
		ctl.Response
		Rank []*UserData `json:"rank"`
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
	for i := range result.Rank {
		res.Rank = append(res.Rank, &UserData{
			UserID: result.Rank[i].UserID,
		})
		for j := range result.Rank[i].Status {
			res.Rank[i].Status = append(
				res.Rank[i].Status,
				&Status{
					ProblemID: result.Rank[i].Status[j].ProblemID,
					Penalty:   result.Rank[i].Status[j].Penalty,
					Accepted:  result.Rank[i].Status[j].Accepted,
					AcTime:    result.Rank[i].Status[j].AcTime,
					LangID:    result.Rank[i].Status[j].LangID,
					Score:     result.Rank[i].Status[j].Score,
				},
			)
		}
	}

	c.JSON(http.StatusOK, res)
}
