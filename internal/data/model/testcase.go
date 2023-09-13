package model

import (
	"main/config"
	"os"
	"path/filepath"
)

type Testcase struct {
	ID         int64
	ProblemID  int64
	InputPath  string
	OutputPath string
}

// TODO: 完成下面的函数

// GetLocalInput 获取本地的输入文件，实现从从网络获取到本地再返回
func (t *Testcase) GetLocalInput() (string, bool) {
	return filepath.Join(config.ConfJudge.File.DemoPath, t.InputPath), true
}

// GetLocalOutput 获取本地的输出文件，实现从从网络获取到本地再返回
func (t *Testcase) GetLocalOutput() (string, bool) {
	return filepath.Join(config.ConfJudge.File.DemoPath, t.OutputPath), true
}

// GetOutput 获取输出内容
func (t *Testcase) GetOutput() (string, bool) {
	outPath := filepath.Join(config.ConfJudge.File.DemoPath, t.OutputPath)
	bytes, err := os.ReadFile(outPath)
	if err != nil {
		return "", false
	}
	return string(bytes), true
}
