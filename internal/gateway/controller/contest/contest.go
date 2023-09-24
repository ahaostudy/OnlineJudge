package contest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rpcContest "main/api/contest"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
)

type (
	CreateContestRequest struct {
		model.Contest
		StartTime int64 `json:"start_time"`
		EndTime   int64 `json:"end_time"`
	}

	CreateContestResonse struct {
		ctl.Response
	}

	DeleteContestResonse struct {
		ctl.Response
	}

	UpdateContestResonse struct {
		ctl.Response
	}
)

func CreateContest(c *gin.Context) {
	req := new(CreateContestRequest)
	res := new(CreateContestResonse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 创建比赛
	result, err := rpc.ContestCli.CreateContest(c.Request.Context(), &rpcContest.CreateContestRequest{
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func DeleteContest(c *gin.Context) {
	res := new(DeleteContestResonse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 删除比赛
	result, err := rpc.ContestCli.DeleteContest(c.Request.Context(), &rpcContest.DeleteContestRequest{
		ID: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func UpdateContest(c *gin.Context) {
	// span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "UpdateContest")
	// defer span.Finish()
	res := new(UpdateContestResonse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	rawData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 更新比赛
	result, err := rpc.ContestCli.UpdateContest(c.Request.Context(), &rpcContest.UpdateContestRequest{
		ID:      int64(id),
		Contest: rawData,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}
