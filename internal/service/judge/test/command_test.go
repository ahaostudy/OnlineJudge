package test

import (
	"fmt"
	"main/config"
	"main/internal/service/judge/pkg/compiler"
	"main/internal/service/judge/pkg/exec"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
)

func TestExec(t *testing.T) {
	// param
	codePath := filepath.Join(config.ConfJudge.File.DemoPath, "c/test.c")
	inputPath := filepath.Join(config.ConfJudge.File.DemoPath, "1.in")

	// 创建编译器对象
	cpl := compiler.NewCompiler(compiler.LangC)

	// 编译代码
	_, err := cpl.Build(codePath)
	defer cpl.Destroy(false)
	if err != nil {
		fmt.Println("编译错误, err:", err.Error())
		return
	}
	exe, _ := cpl.Executable()

	fileName := uuid.New().String()
	outputPath := filepath.Join(config.ConfJudge.File.TempPath, fmt.Sprintf("%s.out", fileName))
	errorPath := filepath.Join(config.ConfJudge.File.TempPath, fmt.Sprintf("%s.err", fileName))

	// 运行代码
	result, err := exec.NewDefaultCommand(*exe, inputPath, outputPath, errorPath).Exec()
	if err != nil {
		fmt.Println("运行错误, err:", err.Error())
		return
	}

	// 读取结果
	fmt.Printf("result: %#v\n", result)

	stdout, _ := os.ReadFile(outputPath)
	fmt.Println("out:", string(stdout))

	stderr, _ := os.ReadFile(errorPath)
	fmt.Println("err:", string(stderr))

	_ = os.Remove(outputPath)
	_ = os.Remove(errorPath)
}
