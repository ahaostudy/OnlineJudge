package submit

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	rpcSubmit "main/api/submit"
	"main/internal/common/build"
	"main/internal/common/code"
	"main/internal/data/model"
	"main/internal/gateway/controller/ctl"
	"main/rpc"
)

const defaultLatestCount int64 = 10

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
		Count int64 `form:"count"`
	}

	GetLatestSubmitsResponse struct {
		ctl.Response
		SubmitList []*model.Submit `json:"submit_list"`
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
	result, err := rpc.SubmitCli.GetSubmit(c.Request.Context(), &rpcSubmit.GetSubmitRequest{
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
	res.Submit, err = build.UnBuildSubmit(result.GetSubmit())
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
	result, err := rpc.SubmitCli.GetSubmitList(c.Request.Context(), &rpcSubmit.GetSubmitListRequest{
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
	res.SubmitList, err = build.UnBuildSubmitList(result.GetSubmitList())
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
	if req.Count == 0 {
		req.Count = defaultLatestCount
	}

	// 获取提交数据
	result, err := rpc.SubmitCli.GetLatestSubmits(c.Request.Context(), &rpcSubmit.GetLatestSubmitsRequest{
		UserID: c.GetInt64("user_id"),
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
	res.SubmitList, err = build.UnBuildSubmitList(result.GetSubmitList())
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
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
	result, err := rpc.SubmitCli.GetSubmitResult(c.Request.Context(), &rpcSubmit.GetSubmitResultRequest{SubmitID: req.SubmitID})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.GetStatusCode())))
		return
	}

	// 将响应转换为模型对象
	res.Result, err = build.UnBuildResult(result.GetResult())
	if err != nil {
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}
