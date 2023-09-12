package test

import (
	"fmt"
	"main/config"
	"main/services/judge/pkg/code"
	"main/services/judge/pkg/compiler"
	"path/filepath"
	"testing"
)

func TestCode(t *testing.T) {
	// param
	codePath := filepath.Join(config.ConfJudge.File.DemoPath, "c/test.c")
	inputPath := filepath.Join(config.ConfJudge.File.DemoPath, "1.in")

	c := code.NewCode(codePath, compiler.LangC)
	result, err := c.Run(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", result)
}
