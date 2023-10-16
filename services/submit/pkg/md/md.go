package md

import (
	"strings"
	"unicode/utf8"

	"github.com/russross/blackfriday/v2"
)

func ExtractTextFromMarkdown(markdown string, count int) string {
	// 去除最后一段的成绩模版
	ls := strings.Split(markdown, "***")
	markdown = strings.Join(ls[:len(ls)-1], "***")

	// 将Markdown转换为HTML
	html := string(blackfriday.Run([]byte(markdown)))

	// 去除HTML标签，只保留文本内容
	text := stripTags(html)

	// 截取前面的字符
	if utf8.RuneCountInString(text) > count {
		text = string([]rune(text)[:count])
	}

	return text
}

func stripTags(html string) string {
	var result strings.Builder
	var insideTag bool

	for _, char := range html {
		if char == '<' {
			insideTag = true
		} else if char == '>' {
			insideTag = false
		} else if !insideTag {
			result.WriteRune(char)
		}
	}

	return result.String()
}
