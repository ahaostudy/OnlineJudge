package model

const (
	LangC = iota + 1
	LangCPP
	LangPython3
	LangGo
	LangJava
)

var suffix = map[int]string{
	LangC:       "c",
	LangCPP:     "cpp",
	LangPython3: "py",
	LangGo:      "go",
	LangJava:    "java",
}

func GetLangSuffix(langID int) string {
	return suffix[langID]
}
