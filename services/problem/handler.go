package problem

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"main/common/code"
	"main/common/raw"
	"main/kitex_gen/contest"
	"main/kitex_gen/problem"
	"main/kitex_gen/submit"
	"main/services/problem/client"
	"main/services/problem/dal/db"
	"main/services/problem/dal/model"
	"main/services/problem/pack"
)

// ProblemServiceImpl implements the last service interface defined in the IDL.
type ProblemServiceImpl struct{}

// GetProblem implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) GetProblem(ctx context.Context, req *problem.GetProblemRequest) (resp *problem.GetProblemResponse, _ error) {
	resp = new(problem.GetProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 访问数据库获取题目信息
	prob, err := db.GetProblemDetail(req.GetProblemID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将模型对象转换为响应结果
	p, err := pack.BuildProblem(prob)
	if err != nil {
		return
	}
	resp.Problem = p

	// 获取示例内容
	for i := 0; i < prob.SampleCount && i < len(prob.Testcases); i++ {
		input, ok := prob.Testcases[i].GetInput()
		if !ok {
			return
		}
		output, ok := prob.Testcases[i].GetOutput()
		if !ok {
			return
		}
		resp.Problem.Samples = append(resp.Problem.Samples, &problem.Sample{
			Input:  input,
			Output: output,
		})
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetProblemList implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) GetProblemList(ctx context.Context, req *problem.GetProblemListRequest) (resp *problem.GetProblemListResponse, _ error) {
	resp = new(problem.GetProblemListResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	page, count := int(req.GetPage()), int(req.GetCount())
	problemList, err := db.GetProblemListLimit((page-1)*count, count)
	if err != nil {
		return
	}

	resp.ProblemList, err = pack.BuildProblems(problemList)
	if err != nil {
		return
	}

	// 获取题目的提交情况
	submitStatus, err := client.SubmitCli.GetSubmitStatus(ctx, &submit.GetSubmitStatusRequest{})
	if err != nil || submitStatus.StatusCode != code.CodeSuccess.Code() {
		return
	}
	// 判断当前用户是否ac
	acceptedStatus, err := client.SubmitCli.GetAcceptedStatus(ctx, &submit.GetAcceptedStatusRequest{UserID: req.GetUserID()})
	if err != nil || acceptedStatus.StatusCode != code.CodeSuccess.Code() {
		return
	}

	for i := range resp.ProblemList {
		id := resp.ProblemList[i].ID
		if v, ok := submitStatus.SubmitStatus[id]; ok {
			resp.ProblemList[i].SubmitCount = v.Count
			resp.ProblemList[i].AcceptedCount = v.AcceptedCount
		}
		if v, ok := acceptedStatus.AcceptedStatus[id]; ok {
			resp.ProblemList[i].IsAccepted = v
		}
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetProblemCount implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) GetProblemCount(ctx context.Context, req *problem.GetProblemCountRequest) (resp *problem.GetProblemCountResponse, _ error) {
	resp = new(problem.GetProblemCountResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	count, err := db.GetProblemCount()
	if err != nil {
		return
	}

	resp.Count = count
	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetContestProblem implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) GetContestProblem(ctx context.Context, req *problem.GetContestProblemRequest) (resp *problem.GetContestProblemResponse, _ error) {
	resp = new(problem.GetContestProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取题目信息
	problem, err := db.GetContestProblem(req.GetUserID(), req.GetProblemID())
	if err != nil {
		return nil, err
	}

	// 判断用户是否有访问权限
	res, err := client.ContestCli.IsAccessible(ctx, &contest.IsAccessibleRequest{
		UserID:    req.GetUserID(),
		ContestID: problem.ContestID,
	})
	if err != nil {
		return
	}
	if res.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = res.StatusCode
		return
	}

	// 无访问权限
	if !res.GetIsAccessible() {
		resp.StatusCode = code.CodeNotRegistred.Code()
		return
	}

	// 将模型对象转换为响应结果
	resp.Problem, err = pack.BuildProblem(problem)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// GetContestProblemList implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) GetContestProblemList(ctx context.Context, req *problem.GetContestProblemListRequest) (resp *problem.GetContestProblemListResponse, _ error) {
	resp = new(problem.GetContestProblemListResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 判断用户是否有访问权限
	res, err := client.ContestCli.IsAccessible(ctx, &contest.IsAccessibleRequest{
		UserID:    req.GetUserID(),
		ContestID: req.GetContestID(),
	})
	if err != nil {
		return
	}
	if res.StatusCode != code.CodeSuccess.Code() {
		resp.StatusCode = res.StatusCode
		return
	}

	// 无访问权限
	if !res.GetIsAccessible() {
		resp.StatusCode = code.CodeNotRegistred.Code()
		return
	}

	// 获取题目列表
	problems, err := db.GetContestProblemList(req.GetContestID())
	if err != nil {
		return
	}
	resp.ProblemList, err = pack.BuildProblems(problems)
	if err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// CreateProblem implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) CreateProblem(ctx context.Context, req *problem.CreateProblemRequest) (resp *problem.CreateProblemResponse, _ error) {
	resp = new(problem.CreateProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 将参数转换为题目对象
	problem, err := pack.UnBuildProblem(req.GetProblem())
	if err != nil {
		return
	}
	problem.ID = 0

	// 插入一条题目信息
	if err := db.InsertProblem(problem); err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// DeleteProblem implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) DeleteProblem(ctx context.Context, req *problem.DeleteProblemRequest) (resp *problem.DeleteProblemResponse, _ error) {
	resp = new(problem.DeleteProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 删除题目
	if err := db.DeleteProblem(req.GetProblemID()); err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// UpdateProblem implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) UpdateProblem(ctx context.Context, req *problem.UpdateProblemRequest) (resp *problem.UpdateProblemResponse, _ error) {
	resp = new(problem.UpdateProblemResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 解析参数
	r := new(raw.Raw)
	if err := r.ReadRawData(req.GetProblem()); err != nil {
		return
	}
	// 忽略id和author_id字段
	delete(r.Map(), "id")
	delete(r.Map(), "author_id")

	// 更新题目
	if err := db.UpdateProblem(req.GetProblemID(), r.Map()); err != nil {
		return
	}

	resp.StatusCode = code.CodeSuccess.Code()
	return
}

// CreateTestcase implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) CreateTestcase(ctx context.Context, req *problem.CreateTestcaseRequest) (resp *problem.CreateTestcaseResponse, _ error) {
	resp = new(problem.CreateTestcaseResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 将输入输出保存到本地文件
	inputPath := fmt.Sprintf("%d/%s.in", req.GetProblemID(), uuid.New().String())
	outputPath := fmt.Sprintf("%d/%s.out", req.GetProblemID(), uuid.New().String())

	// 创建样例对象
	testcase := &model.Testcase{
		ProblemID:  req.GetProblemID(),
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
	// 写入文件
	if !testcase.UploadInput(req.GetInput()) || !testcase.UploadOutput(req.GetOutput()) {
		return
	}
	// 将对象插入数据库
	if db.InsertTestcase(testcase) == nil {
		resp.StatusCode = code.CodeSuccess.Code()
	}

	return
}

// GetTestcase implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) GetTestcase(ctx context.Context, req *problem.GetTestcaseRequest) (resp *problem.GetTestcaseResponse, _ error) {
	resp = new(problem.GetTestcaseResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 获取题目信息
	testcase, err := db.GetTestcase(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 将对象转换为rpc响应
	resp.Testcase, err = pack.BuildTestcase(testcase)
	if err == nil {
		resp.StatusCode = code.CodeSuccess.Code()
	}

	return
}

// DeleteTestcase implements the ProblemServiceImpl interface.
func (s *ProblemServiceImpl) DeleteTestcase(ctx context.Context, req *problem.DeleteTestcaseRequest) (resp *problem.DeleteTestcaseResponse, _ error) {
	resp = new(problem.DeleteTestcaseResponse)
	resp.StatusCode = code.CodeServerBusy.Code()

	// 判断题目是否存在
	testcase, err := db.GetTestcase(req.GetID())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.StatusCode = code.CodeRecordNotFound.Code()
		return
	}
	if err != nil {
		return
	}

	// 并发将该样例的输入输出文件删除
	go func() {
		if p, ok := testcase.GetLocalInput(); !ok {
			os.Remove(p)
		}
		if p, ok := testcase.GetLocalOutput(); !ok {
			os.Remove(p)
		}
	}()

	// 删除改样例数据
	if err = db.DeleteTestcase(req.GetID()); err == nil {
		resp.StatusCode = code.CodeSuccess.Code()
	}

	return
}
