package model

import (
	"os"
	"path/filepath"

	"main/services/problem/config"
)

type Testcase struct {
	ID         int64  `json:"id"`
	ProblemID  int64  `json:"problem_id,omitempty"`
	InputPath  string `json:"input_path"`
	OutputPath string `json:"output_path"`
}

// GetLocalInput 获取本地的输入文件，实现从从网络获取到本地再返回
func (t *Testcase) GetLocalInput() (string, bool) {
	return filepath.Join(config.Config.File.TestcasePath, t.InputPath), true
}

// GetLocalOutput 获取本地的输出文件，实现从从网络获取到本地再返回
func (t *Testcase) GetLocalOutput() (string, bool) {
	return filepath.Join(config.Config.File.TestcasePath, t.OutputPath), true
}

// GetInput 获取输入内容
func (t *Testcase) GetInput() (string, bool) {
	inputPath := filepath.Join(config.Config.File.TestcasePath, t.InputPath)
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		return "", false
	}
	return string(bytes), true
}

// GetOutput 获取输出内容
func (t *Testcase) GetOutput() (string, bool) {
	outPath := filepath.Join(config.Config.File.TestcasePath, t.OutputPath)
	bytes, err := os.ReadFile(outPath)
	if err != nil {
		return "", false
	}
	return string(bytes), true
}

// 将输入内容上传到本地
func (t *Testcase) UploadInput(input []byte) bool {
	inputPath := filepath.Join(config.Config.File.TestcasePath, t.InputPath)
	return save(inputPath, input)
}

// 将输出内容上传到本地
func (t *Testcase) UploadOutput(output []byte) bool {
	outputPath := filepath.Join(config.Config.File.TestcasePath, t.OutputPath)
	return save(outputPath, output)
}

func save(path string, body []byte) bool {
	dirPath := filepath.Dir(path)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return false
	}

	file, err := os.Create(path)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.Write(body)
	return err == nil
}
