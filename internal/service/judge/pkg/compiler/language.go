package compiler

const (
	LangC = iota + 1
	LangCPP
	LangPython3
	LangGo
	LangJava
)

func GetCompiler(langID int) Compiler {
	switch langID {
	case LangC:
		return new(GCC)
	case LangCPP:
		return new(GPP)
	case LangPython3:
		return new(Python3)
	case LangGo:
		return new(GO)
	case LangJava:
		return new(Java)
	default:
		return nil
	}
}

func SaveCode(code []byte, langID int) (string, error) {
	return GetCompiler(langID).SaveCode(code)
}
