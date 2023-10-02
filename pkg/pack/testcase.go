package pack

import (
	problem "main/kitex_gen/problem"
	"main/services/problem/dal/model"
)

func BuildTestcase(t *model.Testcase) (*problem.Testcase, error) {
	testcase := new(problem.Testcase)
	return testcase, new(Builder).Build(t, testcase).Error()
}

func UnBuildTestcase(t *problem.Testcase) (*model.Testcase, error) {
	testcase := new(model.Testcase)
	return testcase, new(Builder).Build(t, testcase).Error()
}
