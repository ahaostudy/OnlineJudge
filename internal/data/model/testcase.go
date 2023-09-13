package model

import (
	"fmt"
	"main/config"
	"os"
	"path/filepath"
)

type Testcase struct {
	ID         int64  `json:"id"`
	ProblemID  int64  `json:"problem_id,omitempty"`
	InputPath  string `json:"input_path"`
	OutputPath string `json:"output_path"`
}

// TODO: 完成下面的函数

// GetLocalInput 获取本地的输入文件，实现从从网络获取到本地再返回
func (t *Testcase) GetLocalInput() (string, bool) {
	return filepath.Join(config.ConfTestcase.File.Path, t.InputPath), true
}

// GetLocalOutput 获取本地的输出文件，实现从从网络获取到本地再返回
func (t *Testcase) GetLocalOutput() (string, bool) {
	return filepath.Join(config.ConfTestcase.File.Path, t.OutputPath), true
}

// GetOutput 获取输出内容
func (t *Testcase) GetOutput() (string, bool) {
	outPath := filepath.Join(config.ConfTestcase.File.Path, t.OutputPath)
	bytes, err := os.ReadFile(outPath)
	if err != nil {
		return "", false
	}
	return string(bytes), true
}

func (t *Testcase) UploadInput(input []byte) bool {
	inputPath := filepath.Join(config.ConfTestcase.File.Path, t.InputPath)
	fmt.Printf("inputPath: %v\n", inputPath)
	err := os.WriteFile(inputPath, input, 0644)
	fmt.Printf("err: %v\n", err)
	return err == nil
	// return os.WriteFile(inputPath, input, 0644) == nil
}

func (t *Testcase) UploadOutput(output []byte) bool {
	outputPath := filepath.Join(config.ConfTestcase.File.Path, t.OutputPath)
	return os.WriteFile(outputPath, output, 0644) == nil
}
