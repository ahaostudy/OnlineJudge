package code

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"main/pkg/status"
	"main/services/judge/config"
	"main/services/judge/pkg/compiler"
	"main/services/judge/pkg/exec"

)

type Code struct {
	path        string
	langID      int
	maxCpuTime  int
	maxRealTime int
	maxMemory   int
	cpl         compiler.Compiler
	exe         *compiler.Executable
}

func NewCode(path string, langID int) *Code {
	conf := config.Config.Sandbox
	return NewCodeLimit(path, langID, conf.DefaultMaxCpuTime, conf.DefaultMaxMemory)
}

func NewCodeLimit(path string, langID int, maxTime, maxMemory int) *Code {
	conf := config.Config.Sandbox
	return &Code{path: path, langID: langID, maxCpuTime: maxTime, maxRealTime: conf.DefaultMaxRealTime, maxMemory: maxMemory}
}

// Build 编译代码
func (c *Code) Build() (result Result, ok bool) {
	result = Result{}
	result.SetStatus(status.StatusCompileError)

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
		result.SetStatus(status.StatusAccepted)
		ok = true
	}
	return
}

// Run 运行代码
func (c *Code) Run(inputPath string) (Result, error) {
	// 如果未编译则编译代码
	if c.exe == nil {
		if r, ok := c.Build(); !ok {
			return r, nil
		}
	}

	// 初始化结果对象
	result := Result{}
	result.SetStatus(status.StatusServerFailed)

	// 生成临时文件路径
	fileName := uuid.New().String()
	outputPath := filepath.Join(config.Config.File.TempPath, fmt.Sprintf("%s.out", fileName))
	errorPath := filepath.Join(config.Config.File.TempPath, fmt.Sprintf("%s.err", fileName))
	defer func() {
		_ = os.Remove(outputPath)
		_ = os.Remove(errorPath)
	}()

	// 执行
	res, err := exec.NewCommand(*c.exe, inputPath, outputPath, errorPath, c.maxCpuTime, c.maxRealTime, c.maxMemory).Exec()
	if err != nil {
		return result, err
	}

	// 将运行结果赋值到返回结果
	stdout, err := os.ReadFile(outputPath)
	if err != nil {
		return result, err
	}
	result.Time, result.Memory, result.Output = res.CpuTime, res.Memory, string(stdout)

	// 系统错误
	switch res.Result {
	case exec.RESULT_SUCCESS:
		// 正常执行
		result.SetStatus(status.StatusAccepted)
		return result, nil
	case exec.RESULT_CPU_TIME_LIMIT_EXCEEDED:
		// 超出时间限制
		result.SetStatus(status.StatusTimeLimitExceeded)
	case exec.RESULT_REAL_TIME_LIMIT_EXCEEDED:
		// 超出时间限制
		result.SetStatus(status.StatusTimeLimitExceeded)
	case exec.RESULT_MEMORY_LIMIT_EXCEEDED:
		// 超出内存限制
		result.SetStatus(status.StatusMemoryLimitExceeded)
	case exec.RESULT_RUNTIME_ERROR:
		// 运行时错误
		stderr, _ := os.ReadFile(errorPath)
		result.Error = string(stderr)
		result.SetStatus(status.StatusRuntimeError)
	}

	// 超出输出长度限制
	if len(result.Output) >= config.Config.Sandbox.DefaultMaxOutputSize {
		result.SetStatus(status.StatusOutputLimitExceeded)
		result.Output = ""
	}

	return result, nil
}

// Destroy 销毁
func (c *Code) Destroy() {
	_ = c.cpl.Destroy(false)
}
