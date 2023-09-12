package compiler

var (
	LangC       = 1
	LangCPP     = 2
	LangPython3 = 3
	LangGo      = 4
	LangJava    = 5
)

func GetCompilerByLang(langID int) Compiler {
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
