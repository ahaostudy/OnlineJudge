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

// GetLocalInput 获取本地的输入文件
func (t *Testcase) GetLocalInput() string {
	return filepath.Join(config.Config.File.TestcasePath, t.InputPath)
}

// GetLocalOutput 获取本地的输出文件
func (t *Testcase) GetLocalOutput() string {
	return filepath.Join(config.Config.File.TestcasePath, t.OutputPath)
}

// GetInput 获取输入内容
func (t *Testcase) GetInput() ([]byte, bool) {
	bytes, err := os.ReadFile(t.GetLocalInput())
	if err != nil {
		return nil, false
	}
	return bytes, true
}

// GetOutput 获取输出内容
func (t *Testcase) GetOutput() ([]byte, bool) {
	bytes, err := os.ReadFile(t.GetLocalOutput())
	if err != nil {
		return nil, false
	}
	return bytes, true
}

// SaveInput 将输入内容保存到本地
func (t *Testcase) SaveInput(input []byte) bool {
	return save(t.GetLocalInput(), input)
}

// SaveOutput 将输出内容保存到本地
func (t *Testcase) SaveOutput(output []byte) bool {
	return save(t.GetLocalOutput(), output)
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
