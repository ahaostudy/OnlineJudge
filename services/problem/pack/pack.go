package pack

import (
	"main/common/pack"
	problem "main/kitex_gen/problem"
	"main/services/problem/dal/model"
)

func BuildProblem(p *model.Problem) (*problem.Problem, error) {
	res := new(problem.Problem)
	builder := new(pack.Builder)

	// 转换基础信息
	if builder.Build(p, res).Error() != nil {
		return res, builder.Error()
	}

	// 转换测试样例列表
	if p.Testcases == nil {
		return res, nil
	}
	for i := range p.Testcases {
		t := new(problem.Testcase)
		builder.Build(p.Testcases[i], t)
		res.Testcases = append(res.Testcases, t)
	}

	return res, builder.Error()
}

func UnBuildProblem(p *problem.Problem) (*model.Problem, error) {
	res := new(model.Problem)
	builder := new(pack.Builder)

	// 转换基础信息
	if builder.Build(p, res).Error() != nil {
		return res, builder.Error()
	}

	// 转换测试样例列表
	if p.Testcases == nil {
		return res, nil
	}
	for i := range p.Testcases {
		t := new(model.Testcase)
		builder.Build(p.Testcases[i], t)
		res.Testcases = append(res.Testcases, t)
	}

	return res, builder.Error()
}

func BuildProblems(ps []*model.Problem) ([]*problem.Problem, error) {
	var problems []*problem.Problem
	builder := new(pack.Builder)
	for _, p := range ps {
		t := new(problem.Problem)
		if builder.Build(p, &t).Error() != nil {
			return nil, builder.Error()
		}
		problems = append(problems, t)
	}
	return problems, nil
}

func UnBuildProblems(ps []*problem.Problem) ([]*model.Problem, error) {
	var problems []*model.Problem
	builder := new(pack.Builder)
	for _, p := range ps {
		t := new(model.Problem)
		if builder.Build(p, &t).Error() != nil {
			return nil, builder.Error()
		}
		problems = append(problems, t)
	}
	return problems, builder.Error()
}

func BuildTestcase(t *model.Testcase) (*problem.Testcase, error) {
	testcase := new(problem.Testcase)
	return testcase, new(pack.Builder).Build(t, testcase).Error()
}

func UnBuildTestcase(t *problem.Testcase) (*model.Testcase, error) {
	testcase := new(model.Testcase)
	return testcase, new(pack.Builder).Build(t, testcase).Error()
}
