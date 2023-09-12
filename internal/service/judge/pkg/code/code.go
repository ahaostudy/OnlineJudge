package code

import (
	"fmt"
	"github.com/google/uuid"
	"main/config"
	"main/internal/service/judge/pkg/compiler"
	"main/internal/service/judge/pkg/errs"
	"main/internal/service/judge/pkg/exec"
	"os"
	"path/filepath"
)

type Code struct {
	path      string
	langID    int
	maxTime   int
	maxMemory int
	cpl       compiler.Compiler
	exe       *compiler.Executable
}

func NewCode(path string, langID int) *Code {
	conf := config.ConfJudge.Sandbox
	return NewCodeLimit(path, langID, conf.DefaultMaxTime, conf.DefaultMaxMemory)
}

func NewCodeLimit(path string, langID int, maxTime, maxMemory int) *Code {
	return &Code{path: path, langID: langID, maxTime: maxTime, maxMemory: maxMemory}
}

// Build 编译代码
func (c *Code) Build() (result Result, ok bool) {
	result = Result{}
	result.SetStatus(StatusCompileError)

	// 创建编译器对象
	c.cpl = compiler.NewCompiler(c.langID)

	// 编译代码
	msg, err := c.cpl.Build(c.path)
	if err != nil {
		result.Error = msg
		return
	}

	// 获取可执行文件对象
	c.exe, err = c.cpl.Executable()
	if err == nil {
		result.SetStatus(StatusAccepted)
		ok = true
	}
	return
}

// Run 运行代码
func (c *Code) Run(inputPath string) (Result, error) {
	// 如果未编译则编译代码
	if c.exe == nil {
		if r, ok := c.Build(); !ok {
			return r, errs.ErrCompilationFailed
		}
	}

	// 初始化结果对象
	result := Result{}
	result.SetStatus(StatusServerFailed)

	// 生成临时文件路径
	fileName := uuid.New().String()
	outputPath := filepath.Join(config.ConfJudge.File.TempPath, fmt.Sprintf("%s.out", fileName))
	errorPath := filepath.Join(config.ConfJudge.File.TempPath, fmt.Sprintf("%s.err", fileName))
	defer func() {
		_ = os.Remove(outputPath)
		_ = os.Remove(errorPath)
	}()

	// 执行
	res, err := exec.NewCommand(*c.exe, inputPath, outputPath, errorPath, c.maxTime, c.maxMemory).Exec()
	if err != nil {
		fmt.Println("1", err.Error())
		return result, err
	}

	// 将运行结果赋值到返回结果
	stdout, err := os.ReadFile(outputPath)
	if err != nil {
		return result, err
	}
	result.Time, result.Memory, result.Output = res.RealTime, res.Memory, string(stdout)

	// 系统错误
	switch res.Result {
	case exec.RESULT_SUCCESS:
		// 正常执行
		result.SetStatus(StatusAccepted)
		return result, nil
	case exec.RESULT_CPU_TIME_LIMIT_EXCEEDED:
		// 超出时间限制
		result.SetStatus(StatusTimeLimitExceeded)
	case exec.RESULT_REAL_TIME_LIMIT_EXCEEDED:
		// 超出时间限制
		result.SetStatus(StatusTimeLimitExceeded)
	case exec.RESULT_MEMORY_LIMIT_EXCEEDED:
		// 超出内存限制
		result.SetStatus(StatusMemoryLimitExceeded)
	case exec.RESULT_RUNTIME_ERROR:
		// 运行时错误
		stderr, _ := os.ReadFile(errorPath)
		result.Error = string(stderr)
		result.SetStatus(StatusRuntimeError)
	}

	// 超出输出长度限制
	if len(result.Output) >= config.ConfJudge.Sandbox.DefaultMaxOutputSize {
		result.SetStatus(StatusOutputLimitExceeded)
	}

	return result, nil
}

// Destroy 销毁
func (c *Code) Destroy() {
	_ = c.cpl.Destroy(false)
}
