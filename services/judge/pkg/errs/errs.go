package errs

import "errors"

var (
	// ErrCodeNotCompiled 代码未编译
	ErrCodeNotCompiled = errors.New("code not compiled")

	// ErrCompilationFailed 编译错误
	ErrCompilationFailed = errors.New("compilation failed")

	// ErrExecutionFailed 执行失败
	ErrExecutionFailed = errors.New("execution failed")

	// ErrJudgeFailed 判题错误
	ErrJudgeFailed = errors.New("judge failed")
)
