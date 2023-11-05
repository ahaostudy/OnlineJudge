package util

import (
	"fmt"
	"regexp"
)

func RemoveDirectoryFromPath(input, directory string) string {
	pattern := fmt.Sprintf(`%s(/|$)`, regexp.QuoteMeta(directory))
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(input, "")

	return result
}
