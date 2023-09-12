package compiler

// Compiler 编译器接口
type Compiler interface {
	Build(codePath string) (msg string, err error) // 编译代码
	Executable() (*Executable, error)              // 获取编译后的可执行对象
	Destroy(removeCode bool) error                 // 销毁编译对象（删除临时文件等）
}

func NewCompiler(langID int) Compiler {
	return GetCompilerByLang(langID)
}
