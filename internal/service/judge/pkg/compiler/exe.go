package compiler

import "sync"

// Executable 编译后的执行参数
type Executable struct {
	path   string    // 可执行文件的路径
	args   []string  // 执行所需要的参数
	env    []string  // 执行环境配置
	kwargs *sync.Map // 其它沙箱参数
}

func (e *Executable) Kwargs() *sync.Map {
	return e.kwargs
}

// AddKwarg 添加参数，在参数不存在时添加，存在时不添加，不覆盖已有参数
func (e *Executable) AddKwarg(key string, value interface{}) {
	if e.kwargs == nil {
		e.kwargs = new(sync.Map)
	}
	if _, ok := e.kwargs.Load(key); ok {
		return
	}
	e.kwargs.Store(key, value)
}

// AddKwargs 添加一组参数，在参数不存在时添加，存在时不添加，不覆盖已有参数
func (e *Executable) AddKwargs(kwargs map[string]interface{}) {
	for k, v := range kwargs {
		e.AddKwarg(k, v)
	}
}

// SetKwarg 设置参数，会覆盖已有参数
func (e *Executable) SetKwarg(key string, value interface{}) {
	if e.kwargs == nil {
		e.kwargs = new(sync.Map)
	}
	e.kwargs.Store(key, value)
}

// SetKwargs 设置一组参数，会覆盖要设置的参数，但不会影响不需要设置的参数
func (e *Executable) SetKwargs(kwargs map[string]interface{}) {
	for k, v := range kwargs {
		e.SetKwarg(k, v)
	}
}

func (e *Executable) Path() string {
	return e.path
}

func (e *Executable) SetPath(path string) {
	e.path = path
}

func (e *Executable) Args() []string {
	return e.args
}

func (e *Executable) SetArgs(args []string) {
	e.args = args
}

func (e *Executable) Env() []string {
	return e.env
}

func (e *Executable) SetEnv(env []string) {
	e.env = env
}
