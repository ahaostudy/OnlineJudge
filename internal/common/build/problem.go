package build

import (
	rpcProblem "main/api/problem"
	rpcTestcase "main/api/testcase"
	"main/internal/data/model"
)

func BuildProblem(p *model.Problem) (*rpcProblem.Problem, error) {
	problem := new(rpcProblem.Problem)
	builder := new(Builder)

	// 转换基础信息
	if builder.Build(p, problem).Error() != nil {
		return problem, builder.Error()
	}

	// 转换测试样例列表
	if p.Testcases == nil {
		return problem, nil
	}
	for i := range p.Testcases {
		t := new(rpcTestcase.Testcase)
		builder.Build(p.Testcases[i], t)
		problem.Testcases = append(problem.Testcases, t)
	}

	return problem, builder.Error()
}

func UnBuildProblem(p *rpcProblem.Problem) (*model.Problem, error) {
	problem := new(model.Problem)
	builder := new(Builder)

	// 转换基础信息
	if builder.Build(p, problem).Error() != nil {
		return problem, builder.Error()
	}

	// 转换测试样例列表
	if p.Testcases == nil {
		return problem, nil
	}
	for i := range p.Testcases {
		t := new(model.Testcase)
		builder.Build(p.Testcases[i], t)
		problem.Testcases = append(problem.Testcases, t)
	}

	return problem, builder.Error()
}
