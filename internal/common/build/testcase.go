package build

import (
	"main/api/problem"
	"main/internal/data/model"
)

func BuildTestcase(t *model.Testcase) (*rpcProblem.Testcase, error) {
	testcase := new(rpcProblem.Testcase)
	return testcase, new(Builder).Build(t, testcase).Error()
}

func UnBuildTestcase(t *rpcProblem.Testcase) (*model.Testcase, error) {
	testcase := new(model.Testcase)
	return testcase, new(Builder).Build(t, testcase).Error()
}
