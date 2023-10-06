package model

import (
	"context"
	"main/common/code"
	"main/kitex_gen/problem"
	"main/services/judge/client"
	"main/services/judge/config"
	"os"
	"path/filepath"
	"time"
)

type Testcase struct {
	ID         int64  `json:"id"`
	ProblemID  int64  `json:"problem_id,omitempty"`
	InputPath  string `json:"input_path"`
	OutputPath string `json:"output_path"`
	input      []byte
	output     []byte
}

// GetLocalInput 获取本地的输入文件，从网络获取到本地
func (t *Testcase) GetLocalInput() (string, bool) {
	path := filepath.Join(config.Config.File.TestcasePath, t.InputPath)
	if len(t.input) > 0 {
		return path, true
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	res, err := client.ProblemCli.GetTestcase(ctx, &problem.GetTestcaseRequest{ID: t.ID})
	if err != nil || res.StatusCode != code.CodeSuccess.Code() {
		return path, false
	}

	t.input = res.GetTestcase().GetInput()
	t.output = res.GetTestcase().GetOutput()

	err = os.WriteFile(path, res.Testcase.Input, 0644)

	return path, err == nil
}

// GetLocalOutput 获取本地的输出文件，从网络获取到本地
func (t *Testcase) GetLocalOutput() (string, bool) {
	path := filepath.Join(config.Config.File.TestcasePath, t.OutputPath)
	if len(t.output) > 0 {
		return path, true
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	res, err := client.ProblemCli.GetTestcase(ctx, &problem.GetTestcaseRequest{ID: t.ID})
	if err != nil || res.StatusCode != code.CodeSuccess.Code() {
		return path, false
	}

	t.input = res.GetTestcase().GetInput()
	t.output = res.GetTestcase().GetOutput()

	err = os.WriteFile(path, res.Testcase.Output, 0644)

	return path, err == nil
}

// GetInput 获取输入内容
func (t *Testcase) GetInput() ([]byte, bool) {
	if len(t.input) > 0 {
		return t.input, true
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := client.ProblemCli.GetTestcase(ctx, &problem.GetTestcaseRequest{ID: t.ID})
	if err != nil || res.StatusCode != code.CodeSuccess.Code() {
		return nil, false
	}

	t.input = res.GetTestcase().GetInput()
	t.output = res.GetTestcase().GetOutput()

	return t.input, true
}

// GetOutput 获取输出内容
func (t *Testcase) GetOutput() ([]byte, bool) {
	if len(t.output) > 0 {
		return t.output, true
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := client.ProblemCli.GetTestcase(ctx, &problem.GetTestcaseRequest{ID: t.ID})
	if err != nil || res.StatusCode != code.CodeSuccess.Code() {
		return nil, false
	}

	t.input = res.GetTestcase().GetInput()
	t.output = res.GetTestcase().GetOutput()

	return t.output, true
}
