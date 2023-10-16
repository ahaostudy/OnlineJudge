package problem

import (
	"main/common/code"
	"main/gateway/client"
	"main/gateway/controller/ctl"
	"main/gateway/pkg/model"
	"main/gateway/pkg/pack"
	"main/kitex_gen/problem"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


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

	CreateProblemRequest struct {
		model.Problem
	}

	CreateProblemResponse struct {
		ctl.Response
	}

	DeleteProblemResponse struct {
		ctl.Response
	}

	UpdateProblemResponse struct {
		ctl.Response
	}

	GetTestcaseResponse struct {
		ctl.Response
		Testcase *model.Testcase `json:"testcase"`
	}

	CreateTestcaseRequest struct {
		ActionType int    `json:"action_type"` // 操作类型：1. 使用文本 2. 使用文件
		ProblemID  int64  `json:"problem_id"`
		Input      string `json:"input"`  // 输入文本
		Output     string `json:"output"` // 输出文本
	}

	CreateTestcaseResponse struct {
		ctl.Response
	}

	DeleteTestcaseResponse struct {
		ctl.Response
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
	result, err := client.ProblemCli.GetProblem(c.Request.Context(), &problem.GetProblemRequest{ProblemID: req.ID})
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
	p, err := pack.UnBuildProblem(result.Problem)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	res.Problem = p
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
	result, err := client.ProblemCli.GetProblemList(c.Request.Context(), &problem.GetProblemListRequest{
		Page:   int64(page),
		Count:  int64(count),
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
	res.ProblemList, err = pack.UnBuildProblems(result.GetProblemList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func CreateProblemCount(c *gin.Context) {
	res := new(GetProblemCountResponse)

	result, err := client.ProblemCli.GetProblemCount(c.Request.Context(), &problem.GetProblemCountRequest{})
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
	result, err := client.ProblemCli.GetContestProblem(c.Request.Context(), &problem.GetContestProblemRequest{
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
	res.Problem, err = pack.UnBuildProblem(result.GetProblem())
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
	result, err := client.ProblemCli.GetContestProblemList(c.Request.Context(), &problem.GetContestProblemListRequest{
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
	res.ProblemList, err = pack.UnBuildProblems(result.GetProblemList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func CreateProblem(c *gin.Context) {
	req := new(CreateProblemRequest)
	res := new(CreateProblemResponse)

	// 解析参数
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	req.AuthorID = c.GetInt64("user_id")

	// 将参数中的问题信息转换为rpc参数
	p, err := pack.BuildProblem(&req.Problem)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	// 创建问题
	result, err := client.ProblemCli.CreateProblem(c.Request.Context(), &problem.CreateProblemRequest{Problem: p})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func DeleteProblem(c *gin.Context) {
	res := new(DeleteProblemResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 删除题目
	result, err := client.ProblemCli.DeleteProblem(c.Request.Context(), &problem.DeleteProblemRequest{
		ProblemID: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func UpdateProblem(c *gin.Context) {
	res := new(UpdateProblemResponse)

	// 解析参数 id为必须参数
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

	// 更新题目信息
	result, err := client.ProblemCli.UpdateProblem(c.Request.Context(), &problem.UpdateProblemRequest{
		ProblemID: id,
		Problem:   rawData,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func GetTestcase(c *gin.Context) {
	res := new(GetTestcaseResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 获取样例数据
	result, err := client.ProblemCli.GetTestcase(c.Request.Context(), &problem.GetTestcaseRequest{
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
	res.Testcase, err = pack.UnBuildTestcase(result.GetTestcase())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func CreateTestcase(c *gin.Context) {
	req := new(CreateTestcaseRequest)
	res := new(CreateTestcaseResponse)

	// 解析参数
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	switch req.ActionType {
	case 1:
		// 创建题目
		result, err := client.ProblemCli.CreateTestcase(c.Request.Context(), &problem.CreateTestcaseRequest{
			ProblemID: req.ProblemID,
			Input:     []byte(req.Input),
			Output:    []byte(req.Output),
		})
		if err != nil {
			c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
			return
		}

		// 响应结果
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
	default:
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
	}
}

func DeleteTestcase(c *gin.Context) {
	res := new(DeleteTestcaseResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 删除样例
	result, err := client.ProblemCli.DeleteTestcase(c.Request.Context(), &problem.DeleteTestcaseRequest{
		ID: id,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}
