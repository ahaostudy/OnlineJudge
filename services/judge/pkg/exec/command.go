package exec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"reflect"
	"regexp"
	"strings"

	"main/services/judge/config"
	"main/services/judge/pkg/compiler"
	"main/services/judge/pkg/errs"
)

// Command 沙箱调用命令
type Command struct {
	exe           compiler.Executable // 可执行文件对象
	InputPath     string              // 输入文件路径
	OutputPath    string              // 输出文件路径
	ErrorPath     string              // 错误文件路径
	MaxCpuTime    int                 // 最大CPU运行时间
	MaxRealTime   int                 // 最大实际运行时间
	MaxMemory     int                 // 最大内存
	MaxOutputSize int                 // 最大输出大小
}

// NewCommand 创建命令对象
func NewCommand(exe compiler.Executable, inputPath string, outputPath string, errorPath string, maxCpuTime int, maxRealTime int, maxMemory int) *Command {
	conf := config.Config.Sandbox
	return &Command{exe: exe, InputPath: inputPath, OutputPath: outputPath, ErrorPath: errorPath, MaxCpuTime: maxCpuTime, MaxRealTime: maxRealTime, MaxMemory: maxMemory, MaxOutputSize: conf.DefaultMaxOutputSize}
}

// NewDefaultCommand 创建默认命令
func NewDefaultCommand(exe compiler.Executable, inputPath string, outputPath string, errorPath string) *Command {
	conf := config.Config.Sandbox
	return NewCommand(exe, inputPath, outputPath, errorPath, conf.DefaultMaxCpuTime, conf.DefaultMaxRealTime, conf.DefaultMaxMemory)
}

// Exec 执行命令
func (c *Command) Exec() (*Result, error) {
	// 获取执行命令
	command, err := c.Command()
	if err != nil {
		return nil, errs.ErrCodeNotCompiled
	}
	fmt.Printf("command: %v\n", command)

	// 为命令添加sudo权限
	cmd := exec.Command("sudo", "-S", "sh", "-c", command)
	cmd.Stdin = strings.NewReader(config.Config.System.SudoPwd + "\n")

	// 执行命令
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// 处理执行错误
	if err := cmd.Run(); err != nil {
		return nil, errs.ErrExecutionFailed
	}

	// 解析执行结果
	result := new(Result)
	err = json.Unmarshal(stdout.Bytes(), result)
	if err != nil {
		return nil, errs.ErrExecutionFailed
	}

	return result, nil
}

// GetArgs 通过字段获取参数值
func (c *Command) GetArgs(field string) (value interface{}) {
	// 使用反射获取对象的类型信息
	v := reflect.ValueOf(*c)
	val := v.FieldByName(field)
	if val.IsValid() {
		value = val.Interface()
	}
	return value
}

// Command 命令序列化
func (c *Command) Command() (string, error) {
	// 沙箱参数
	var args []string
	// 命令执行参数
	for _, arg := range c.exe.Args() {
		arg = regexp.MustCompile(`{([^}]+)}`).ReplaceAllStringFunc(arg, func(s string) string {
			v := c.GetArgs(s[1 : len(s)-1])
			return fmt.Sprintf(`%v`, v)
		})
		args = append(args, fmt.Sprintf(`--args="%s"`, arg))
	}
	// 命令执行环境
	for _, e := range c.exe.Env() {
		args = append(args, fmt.Sprintf(`--env="%s"`, e))
	}
	// 添加其它参数
	c.exe.SetKwargs(map[string]interface{}{
		"exe_path":    c.exe.Path(),
		"input_path":  c.InputPath,
		"output_path": c.OutputPath,
		"error_path":  c.ErrorPath,
		"log_path":    config.Config.Sandbox.LogPath,
	})
	c.exe.AddKwargs(map[string]interface{}{
		"max_memory":      c.MaxMemory,
		"max_cpu_time":    c.MaxCpuTime,
		"max_real_time":   c.MaxRealTime,
		"max_output_size": c.MaxOutputSize,
	})
	// 其它参数
	c.exe.Kwargs().Range(func(k, v any) bool {
		if v != nil {
			args = append(args, fmt.Sprintf(`--%s=%v`, k, v))
		}
		return true
	})

	// 返回沙箱执行命令
	return exec.Command(config.Config.Sandbox.ExePath, args...).String(), nil
}
