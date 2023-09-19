package build

import (
	rpcProblem "main/api/problem"
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
		t := new(rpcProblem.Testcase)
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

func BuildProblems(ps []*model.Problem) ([]*rpcProblem.Problem, error) {
	var problems []*rpcProblem.Problem
	builder := new(Builder)
	for _, p := range ps {
		var problem *rpcProblem.Problem
		if builder.Build(p, &problem).Error() != nil {
			return nil, builder.Error()
		}
		problems = append(problems, problem)
	}
	return problems, nil
}

func UnBuildProblems(ps []*rpcProblem.Problem) ([]*model.Problem, error) {
	var problems []*model.Problem
	builder := new(Builder)
	for _, p := range ps {
		var problem *model.Problem
		if builder.Build(p, &problem).Error() != nil {
			return nil, builder.Error()
		}
		problems = append(problems, problem)
	}
	return problems, builder.Error()
}
