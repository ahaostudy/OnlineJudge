package build

import (
	rpcTestcase "main/api/testcase"
	"main/internal/data/model"
)

func BuildTestcase(t *model.Testcase) (*rpcTestcase.Testcase, error) {
	testcase := new(rpcTestcase.Testcase)
	return testcase, new(Builder).Build(t, testcase).Error()
}

func UnBuildTestcase(t *rpcTestcase.Testcase) (*model.Testcase, error) {
	testcase := new(model.Testcase)
	return testcase, new(Builder).Build(t, testcase).Error()
}
