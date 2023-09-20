package problem

import (
	"context"
	rpcProblem "main/api/problem"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

	GetProblemResponse struct {
		ctl.Response
		Problem *model.Problem `json:"problem"`
	}

	GetProblemListResponse struct {
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
		c.JSON(http.StatusOK, res.CodeOf(common.CodeInvalidParams))
		return
	}

	// 获取题目
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result, err := rpc.ProblemCli.GetProblem(ctx, &rpcProblem.GetProblemRequest{ProblemId: req.ID})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	// 获取失败
	if result.StatusCode != common.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(common.Code(result.StatusCode)))
		return
	}

	// 将结果转换为题目对象
	problem, err := build.UnBuildProblem(result.Problem)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}
	res.Problem = problem

	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetProblemList(c *gin.Context) {
	res := new(GetProblemListResponse)

	// 解析参数，不使用BindQuery是因为要添加默认值
	// 设置了默认值，此处出现错误的概率是极低极低，忽略错误
	page, _ := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	count, _ := strconv.Atoi(c.DefaultQuery("count", strconv.Itoa(defaultCount)))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 获取题目列表
	result, err := rpc.ProblemCli.GetProblemList(ctx, &rpcProblem.GetProblemListRequest{
		Page:  int64(page),
		Count: int64(count),
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}
	if result.StatusCode != common.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(common.Code(result.StatusCode)))
		return
	}

	// 转换为模型对象
	res.ProblemList, err = build.UnBuildProblems(result.GetProblemList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(common.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
