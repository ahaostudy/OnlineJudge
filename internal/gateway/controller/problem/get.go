package problem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rpcProblem "main/api/problem"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
)

// 默认每页的数量
const (
	defaultPage  = 1
	defaultCount = 20
)

type (
	GetProblemRequest struct {
		ID int64
	}

	Sample struct {
		Input  string `json:"input"`
		Output string `json:"output"`
	}

	GetProblemResponse struct {
		ctl.Response
		Problem *model.Problem `json:"problem"`
		Samples []*Sample      `json:"samples"`
	}

	GetProblemListResponse struct {
		ctl.Response
		ProblemList []*model.Problem `json:"problem_list"`
	}

	GetProblemCountResponse struct {
		ctl.Response
		Count int64 `json:"count"`
	}

	GetContestProblemResponse struct {
		ctl.Response
		Problem *model.Problem `json:"problem"`
	}

	GetContestProblemListResponse struct {
		ctl.Response
		ProblemList []*model.Problem `json:"problem_list"`
	}
)

func GetProblem(c *gin.Context) {
	req := new(GetProblemRequest)
	res := new(GetProblemResponse)

	// 解析参数
	req.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	if req.ID == 0 {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 获取题目
	result, err := rpc.ProblemCli.GetProblem(c.Request.Context(), &rpcProblem.GetProblemRequest{ProblemID: req.ID})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	// 获取失败
	if result.StatusCode != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	// 将结果转换为题目对象
	problem, err := build.UnBuildProblem(result.Problem)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	res.Problem = problem
	res.Problem.Testcases = nil

	// 复制示例
	for _, sample := range result.GetProblem().GetSamples() {
		res.Samples = append(res.Samples, &Sample{
			Input:  sample.Input,
			Output: sample.Output,
		})
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetProblemList(c *gin.Context) {
	res := new(GetProblemListResponse)

	// 解析参数，不使用BindQuery是因为要添加默认值
	// 设置了默认值，此处出现错误的概率是极低极低，忽略错误
	page, _ := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	count, _ := strconv.Atoi(c.DefaultQuery("count", strconv.Itoa(defaultCount)))

	// 获取题目列表
	result, err := rpc.ProblemCli.GetProblemList(c.Request.Context(), &rpcProblem.GetProblemListRequest{
		Page:  int64(page),
		Count: int64(count),
		UserID: c.GetInt64("user_id"),
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.StatusCode != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	// 转换为模型对象
	res.ProblemList, err = build.UnBuildProblems(result.GetProblemList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func CreateProblemCount(c *gin.Context) {
	res := new(GetProblemCountResponse)

	result, err := rpc.ProblemCli.GetProblemCount(c.Request.Context(), &rpcProblem.GetProblemCountRequest{})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Count = result.GetCount()
	res.CodeOf(code.Code(result.GetStatusCode()))
	c.JSON(http.StatusOK, res)
}

func GetContestProblem(c *gin.Context) {
	res := new(GetContestProblemResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userID := c.GetInt64("user_id")

	// 获取比赛题目
	result, err := rpc.ProblemCli.GetContestProblem(c.Request.Context(), &rpcProblem.GetContestProblemRequest{
		UserID:    userID,
		ProblemID: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.StatusCode != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	// 转换为模型对象
	res.Problem, err = build.UnBuildProblem(result.GetProblem())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetContestProblemList(c *gin.Context) {
	res := new(GetProblemListResponse)

	// 解析参数
	contestID, err := strconv.ParseInt(c.Query("contest_id"), 10, 64)
	if err != nil || contestID == 0 {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userID := c.GetInt64("user_id")

	// 获取比赛题目列表
	result, err := rpc.ProblemCli.GetContestProblemList(c.Request.Context(), &rpcProblem.GetContestProblemListRequest{
		ContestID: contestID,
		UserID:    userID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.StatusCode != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
		return
	}

	// 转换为模型对象
	res.ProblemList, err = build.UnBuildProblems(result.GetProblemList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
