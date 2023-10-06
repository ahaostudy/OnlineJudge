package contest

import (
	"net/http"
	"strconv"

	"main/common/code"
	"main/gateway/client"
	"main/gateway/controller/ctl"
	"main/gateway/pkg/model"
	"main/gateway/pkg/pack"
	"main/kitex_gen/contest"

	"github.com/gin-gonic/gin"
)

var (
	defaultPage  = 1
	defaultCount = 20
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

	GetContestResponse struct {
		ctl.Response
		Contest *model.Contest `json:"contest"`
	}

	GetContestListResponse struct {
		ctl.Response
		ContestList []*model.Contest `json:"contest_list"`
	}

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

	RegisterContestRequest struct {
		ActionType int64 `json:"action_type"`
		ContestID  int64 `json:"contest_id"`
	}

	RegisterContestResponse struct {
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
	result, err := client.ContestCli.CreateContest(c.Request.Context(), &contest.CreateContestRequest{
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
	result, err := client.ContestCli.DeleteContest(c.Request.Context(), &contest.DeleteContestRequest{
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
	result, err := client.ContestCli.UpdateContest(c.Request.Context(), &contest.UpdateContestRequest{
		ID:      int64(id),
		Contest: rawData,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

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
	result, err := client.ContestCli.GetContest(c.Request.Context(), &contest.GetContestRequest{
		ID:     int64(id),
		UserID: userID,
	})
	if err != nil {
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	// 转换为模型对象
	res.Contest, err = pack.UnBuildContest(result.GetContest())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetContestList(c *gin.Context) {
	res := new(GetContestListResponse)

	// 解析参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	count, _ := strconv.Atoi(c.DefaultQuery("count", strconv.Itoa(defaultCount)))

	// 获取比赛列表
	result, err := client.ContestCli.GetContestList(c.Request.Context(), &contest.GetContestListRequest{
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
	res.ContestList, err = pack.UnBuildContestList(result.GetContestList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

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
	result, err := client.ContestCli.ContestRank(c.Request.Context(), &contest.ContestRankRequest{
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
func RegisterContest(c *gin.Context) {
	req := new(RegisterContestRequest)
	res := new(RegisterContestResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userID := c.GetInt64("user_id")

	// 判断操作类型
	switch req.ActionType {
	case 1:
		// 报名比赛
		result, err := client.ContestCli.Register(c.Request.Context(), &contest.RegisterRequest{
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
		result, err := client.ContestCli.UnRegister(c.Request.Context(), &contest.UnRegisterRequest{
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
