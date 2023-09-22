package build

import (
	rpcJudge "main/api/judge"
	"main/internal/service/judge/pkg/code"
)

func BuildResult(r *code.Result) (*rpcJudge.JudgeResult, error) {
	result := new(rpcJudge.JudgeResult)
	return result, new(Builder).Build(r, &result).Error()
}

func UnBuildResult(r *rpcJudge.JudgeResult) (*code.Result, error) {
	result := new(code.Result)
	return result, new(Builder).Build(r, result).Error()
}
