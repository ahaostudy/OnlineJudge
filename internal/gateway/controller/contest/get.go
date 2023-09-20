package contest

import (
	"main/api/contest"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	GetContestResponse struct {
		ctl.Response
		Contest *model.Contest `json:"contest"`
	}

	GetContestListResponse struct {
		ctl.Response
		ContestList []*model.Contest `json:"contest_list"`
	}
)

func GetContest(c *gin.Context) {
	res := new(GetContestResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userID := c.GetInt64("user_id")

	// 获取比赛信息
	result, err := rpc.ContestCli.GetContest(c.Request.Context(), &rpcContest.GetContestRequest{
		Id:     int64(id),
		UserId: userID,
	})
	if err != nil {
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	res.Contest, err = build.UnBuildContest(result.GetContest())
	res.Success()
	return
}

func GetContestList(c *gin.Context) {
	res := new(GetContestListResponse)

	// 解析参数
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	count, err := strconv.ParseInt(c.Query("count"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 获取比赛列表
	result, err := rpc.ContestCli.GetContestList(c.Request.Context(), &rpcContest.GetContestListRequest{
		Page:  int64(page),
		Count: int64(count),
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	// 转换为模型对象
	res.ContestList, err = build.UnBuildContestList(result.GetContestList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
