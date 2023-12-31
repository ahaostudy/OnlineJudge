package submit

import (
	"main/common/code"
	"main/gateway/client"
	"main/gateway/controller/ctl"
	"main/gateway/pkg/model"
	"main/gateway/pkg/pack"
	"main/kitex_gen/submit"
	"net/http"
	"sort"
	"strconv"

	build "main/common/pack"

	"github.com/gin-gonic/gin"
)

const defaultLatestCount = 10

type (
	GetSubmitResponse struct {
		ctl.Response
		Submit *model.Submit `json:"submit"`
	}

	GetSubmitListRequest struct {
		ProblemID int64 `form:"problem_id"`
	}

	GetSubmitListResponse struct {
		ctl.Response
		SubmitList []*model.Submit `json:"submit_list"`
	}

	GetResultRequest struct {
		SubmitID int64 `json:"submit_id"`
	}

	GetResultResponse struct {
		ctl.Response
		Result any `json:"result"`
	}

	GetLatestSubmitsRequest struct {
		UserID int64 `form:"user_id"`
		Count  int64 `form:"count"`
	}

	GetLatestSubmitsResponse struct {
		ctl.Response
		SubmitList []*model.Submit `json:"submit_list"`
	}

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

	DebugRequest struct {
		Code   string `json:"code"`
		Input  string `json:"input"`
		LangID int64  `json:"lang_id"`
	}

	DebugResponse struct {
		ctl.Response
		Result *model.JudgeResult `json:"result"`
	}

	GetSubmitCalendarRequest struct {
		UserID int64 `form:"user_id"`
	}

	GetSubmitCalendarResponse struct {
		ctl.Response
		Data map[string]int64 `json:"data"`
	}

	GetSubmitStatisticsRequest struct {
		UserID int64 `form:"user_id"`
	}

	GetSubmitStatisticsResponse struct {
		ctl.Response
		SloveCount  int64 `json:"slove_count"`
		SubmitCount int64 `json:"submit_count"`
		EasyCount   int64 `json:"easy_count"`
		MiddleCount int64 `json:"middle_count"`
		HardCount   int64 `json:"hard_count"`
		LangCounts  []struct {
			ID int64 `json:"id"`
			Count  int64 `json:"count"`
		} `json:"lang_counts"`
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

	// 获取提交
	result, err := client.SubmitCli.GetSubmit(c.Request.Context(), &submit.GetSubmitRequest{
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
	res.Submit, err = pack.UnBuildSubmit(result.GetSubmit())
	if err != nil {
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetSubmitList(c *gin.Context) {
	req := new(GetSubmitListRequest)
	res := new(GetSubmitListResponse)

	// 解析参数
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 获取提交数据
	result, err := client.SubmitCli.GetSubmitList(c.Request.Context(), &submit.GetSubmitListRequest{
		UserID:    c.GetInt64("user_id"),
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
	res.SubmitList, err = pack.UnBuildSubmitList(result.GetSubmitList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetLatestSubmits(c *gin.Context) {
	req := new(GetLatestSubmitsRequest)
	res := new(GetLatestSubmitsResponse)

	// 解析参数
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	if req.UserID == 0 {
		req.UserID = c.GetInt64("user_id")
	}
	if req.Count == 0 {
		req.Count = defaultLatestCount
	}

	// 获取提交数据
	result, err := client.SubmitCli.GetLatestSubmits(c.Request.Context(), &submit.GetLatestSubmitsRequest{
		UserID: req.UserID,
		Count:  req.Count,
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
	res.SubmitList, err = pack.UnBuildSubmitList(result.GetSubmitList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	builder := new(build.Builder)
	for i := range res.SubmitList {
		res.SubmitList[i].Problem = new(model.Problem)
		builder.Build(*result.SubmitList[i].GetProblem(), res.SubmitList[i].Problem)
		if builder.Error() != nil {
			c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
			return
		}
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetResult(c *gin.Context) {
	req := new(GetResultRequest)
	res := new(GetResultResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 获取提交结果
	result, err := client.SubmitCli.GetSubmitResult(c.Request.Context(), &submit.GetSubmitResultRequest{SubmitID: req.SubmitID})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.GetStatusCode())))
		return
	}

	// 将响应转换为模型对象
	res.Result, err = pack.UnBuildResult(result.GetResult())
	if err != nil {
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

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
		result, e := client.SubmitCli.Submit(c.Request.Context(), &submit.SubmitRequest{
			ProblemID: req.ProblemID,
			UserID:    c.GetInt64("user_id"),
			Code:      []byte(req.Code),
			LangID:    req.LangID,
		})
		submitID, statusCode, err = result.GetSubmitID(), result.GetStatusCode(), e
	} else {
		result, e := client.SubmitCli.SubmitContest(c.Request.Context(), &submit.SubmitContestRequest{
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
	result, err := client.SubmitCli.DeleteSubmit(c.Request.Context(), &submit.DeleteSubmitRequest{
		ID:     id,
		UserID: userID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func Debug(c *gin.Context) {
	req := new(DebugRequest)
	res := new(DebugResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 使用给定的代码、输入和语言ID调用 SubmitCli 的 Debug 方法
	result, err := client.SubmitCli.Debug(c.Request.Context(), &submit.DebugReqeust{
		Code:   []byte(req.Code),
		Input:  []byte(req.Input),
		LangID: req.LangID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		res.CodeOf(code.Code(result.StatusCode))
		return
	}

	// 将响应结果和状态码设置为来自 SubmitCli 响应的值
	res.Result, err = pack.UnBuildResult(result.Result)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetSubmitCalendar(c *gin.Context) {
	req := new(GetSubmitCalendarRequest)
	res := new(GetSubmitCalendarResponse)

	// 解析参数
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 获取用户提交日历
	result, err := client.SubmitCli.GetSubmitCalendar(c.Request.Context(), &submit.GetSubmitCalendarRequest{
		UserID: req.UserID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Data = result.GetSubmitCalendar()
	res.CodeOf(code.Code(result.StatusCode))
	c.JSON(http.StatusOK, res)
}

func GetSubmitStatistics(c *gin.Context) {
	req := new(GetSubmitStatisticsRequest)
	res := new(GetSubmitStatisticsResponse)

	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	result, err := client.SubmitCli.GetSubmitStatistics(c.Request.Context(), &submit.GetSubmitStatisticsRequest{
		UserID: req.UserID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.StatusCode != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	res.SubmitCount = result.GetSubmitCount()
	res.SloveCount = result.GetSloveCount()
	res.EasyCount = result.GetEasyCount()
	res.MiddleCount = result.GetMiddleCount()
	res.HardCount = result.GetHardCount()
	for k, v := range result.GetLangCounts() {
		res.LangCounts = append(res.LangCounts, struct{ID int64 "json:\"id\""; Count int64 "json:\"count\""}{
			ID: k,
			Count: v,
		})
	}
	sort.Slice(res.LangCounts, func(i, j int) bool {
		return res.LangCounts[i].Count > res.LangCounts[j].Count
	})
	res.CodeOf(code.Code(result.StatusCode))
	c.JSON(http.StatusOK, res)
}
