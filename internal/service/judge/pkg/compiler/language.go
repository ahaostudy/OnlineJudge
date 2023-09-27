package compiler

const (
	LangCPP = iota + 1
	LangC
	LangPython3
	LangJava
	LangGo
)

func GetCompiler(langID int) Compiler {
	switch langID {
	case LangCPP:
		return new(GPP)
	case LangC:
		return new(GCC)
	case LangPython3:
		return new(Python3)
	case LangJava:
		return new(Java)
	case LangGo:
		return new(GO)
	default:
		return nil
	}
}

func SaveCode(code []byte, langID int) (string, error) {
	return GetCompiler(langID).SaveCode(code)
}
