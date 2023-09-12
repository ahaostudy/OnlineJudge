package test

import (
	"fmt"
	"main/config"
	"main/internal/data/model"
	"main/internal/service/judge/pkg/compiler"
	"main/internal/service/judge/pkg/judge"
	"path/filepath"
	"testing"
)

func TestJudge(t *testing.T) {
	codePath := filepath.Join(config.ConfJudge.File.DemoPath, "c/test.c")

	problem := &model.Problem{
		MaxTime:   1000,
		MaxMemory: 512 * 1024 * 1024,
		Testcases: []*model.Testcase{
			{ID: 1, InputPath: "1.in", OutputPath: "1.out"},
			{ID: 2, InputPath: "2.in", OutputPath: "2.out"},
			{ID: 3, InputPath: "3.in", OutputPath: "3.out"},
		},
	}

	result, err := judge.Judge(problem, codePath, compiler.LangC)
	fmt.Printf("%#v\n", result)
	if err != nil {
		panic(err)
	}
}
