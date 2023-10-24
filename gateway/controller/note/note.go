package note

import (
	"fmt"
	"main/common/code"
	build "main/common/pack"
	"main/gateway/client"
	"main/gateway/controller/ctl"
	"main/gateway/pkg/model"
	"main/gateway/pkg/pack"
	"main/kitex_gen/submit"
	"main/kitex_gen/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	defaultPage  = 1
	defaultCount = 20
)

type (
	GetNoteResopnse struct {
		ctl.Response
		User *model.User `json:"user"`
		Note *model.Note `json:"note"`
	}

	GetNoteListRequest struct {
		Page      int64 `form:"page"`
		Count     int64 `form:"count"`
		UserID    int64 `form:"user_id"`
		ProblemID int64 `form:"problem_id"`
		SubmitID  int64 `form:"submit_id"`
	}

	GetNoteListResponse struct {
		ctl.Response
		NoteList []*model.Note `json:"note_list"`
	}

	CreateNoteRequest struct {
		model.Note
	}

	CreateNoteResponse struct {
		ctl.Response
	}

	UpdateNoteResponse struct {
		ctl.Response
	}

	DeleteNoteResponse struct {
		ctl.Response
	}
)

func GetNote(c *gin.Context) {
	res := new(GetNoteResopnse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	// 获取笔记信息
	result, err := client.SubmitCli.GetNote(c.Request.Context(), &submit.GetNoteRequest{
		ID:     id,
		UserID: c.GetInt64("user_id"),
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result.GetStatusCode())))
		return
	}

	res.Note, err = pack.UnBuildNote(result.Note)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	// 获取作者信息
	result1, err := client.UserCli.GetUser(c.Request.Context(), &user.GetUserRequest{
		ID: res.Note.UserID,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result1.GetStatusCode() != code.CodeSuccess.Code() {
		c.JSON(http.StatusOK, res.CodeOf(code.Code(result1.GetStatusCode())))
		return
	}

	res.User, err = pack.UnBuildUser(result1.User)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func GetNoteList(c *gin.Context) {
	req := new(GetNoteListRequest)
	res := new(GetNoteListResponse)

	// 解析参数
	if err := c.BindQuery(req); err != nil {
		fmt.Println("query")
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	if req.UserID == -1 {
		req.UserID = c.GetInt64("user_id")
	}
	if req.Page == 0 {
		req.Page = defaultPage
	}
	if req.Count == 0 {
		req.Count = defaultCount
	}

	// 获取笔记列表
	result, err := client.SubmitCli.GetNoteList(c.Request.Context(), &submit.GetNoteListRequest{
		UserID:    req.UserID,
		ProblemID: req.ProblemID,
		SubmitID:  req.SubmitID,
		Page:      req.Page,
		Count:     req.Count,
		IsPublic:  true,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}
	if result.GetStatusCode() != code.CodeSuccess.Code() {
		res.CodeOf(code.Code(result.StatusCode))
		return
	}

	res.NoteList, err = pack.UnBuildNoteList(result.NoteList)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	// 转换作者的信息
	builder := new(build.Builder)
	for i := range result.GetNoteList() {
		res.NoteList[i].Author = new(model.User)
		builder.Build(result.GetNoteList()[i].GetAuthor(), res.NoteList[i].Author)

		// 减少响应内容
		res.NoteList[i].Author.Signature = ""
	}
	if builder.Error() != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	c.JSON(http.StatusOK, res)
}

func CreateNote(c *gin.Context) {
	req := new(CreateNoteRequest)
	res := new(CreateNoteResponse)

	// 解析参数
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	note, err := pack.BuildNote(&req.Note)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	// 创建笔记
	result, err := client.SubmitCli.CreateNote(c.Request.Context(), &submit.CreateNoteRequest{
		UserID: c.GetInt64("user_id"),
		Note:   note,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func UpdateNote(c *gin.Context) {
	res := new(UpdateNoteResponse)

	// 解析参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	raw, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	result, err := client.SubmitCli.UpdateNote(c.Request.Context(), &submit.UpdateNoteRequest{
		ID:     id,
		UserID: c.GetInt64("user_id"),
		Note:   raw,
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}

func DeleteNote(c *gin.Context) {
	res := new(DeleteNoteResponse)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}

	result, err := client.SubmitCli.DeleteNote(c.Request.Context(), &submit.DeleteNoteRequest{
		ID:     id,
		UserID: c.GetInt64("user_id"),
	})
	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	c.JSON(http.StatusOK, res.CodeOf(code.Code(result.StatusCode)))
}
