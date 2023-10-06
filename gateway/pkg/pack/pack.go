package pack

import (
	"main/common/pack"
	"main/gateway/pkg/model"
	"main/kitex_gen/contest"
	"main/kitex_gen/problem"
	"main/kitex_gen/submit"
	"main/kitex_gen/user"
	"time"
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

func UnBuildResult(r *submit.JudgeResult) (*model.JudgeResult, error) {
	result := new(model.JudgeResult)
	return result, new(pack.Builder).Build(r, &result).Error()
}

func BuildSubmit(s *model.Submit) (*submit.Submit, error) {
	t := new(submit.Submit)
	if err := new(pack.Builder).Build(s, t).Error(); err != nil {
		return nil, err
	}
	t.CreatedAt = s.CreatedAt.UnixMilli()
	return t, nil
}

func UnBuildSubmit(s *submit.Submit) (*model.Submit, error) {
	t := new(model.Submit)
	if err := new(pack.Builder).Build(s, t).Error(); err != nil {
		return nil, err
	}
	t.CreatedAt = time.UnixMilli(s.GetCreatedAt())
	return t, nil
}

func BuildSubmitList(submits []*model.Submit) ([]*submit.Submit, error) {
	var submitList []*submit.Submit
	for _, s := range submits {
		t, err := BuildSubmit(s)
		if err != nil {
			return nil, err
		}
		submitList = append(submitList, t)
	}
	return submitList, nil
}

func UnBuildSubmitList(submits []*submit.Submit) ([]*model.Submit, error) {
	var submitList []*model.Submit
	for _, s := range submits {
		t, err := UnBuildSubmit(s)
		if err != nil {
			return nil, err
		}
		submitList = append(submitList, t)
	}
	return submitList, nil
}

func BuildUser(u *model.User) (*user.User, error) {
	t := new(user.User)
	if err := new(pack.Builder).Build(u, t).Error(); err != nil {
		return nil, err
	}
	return t, nil
}

func UnBuildUser(u *user.User) (*model.User, error) {
	t := new(model.User)
	if err := new(pack.Builder).Build(u, t).Error(); err != nil {
		return nil, err
	}
	return t, nil
}

func BuildContest(c *model.Contest) (*contest.Contest, error) {
	t := new(contest.Contest)
	builder := new(pack.Builder)
	if builder.Build(c, t).Error() != nil {
		return nil, builder.Error()
	}
	t.StartTime = c.StartTime.UnixMilli()
	t.EndTime = c.EndTime.UnixMilli()
	return t, nil
}

func UnBuildContest(c *contest.Contest) (*model.Contest, error) {
	contest := new(model.Contest)
	builder := new(pack.Builder)
	if builder.Build(c, contest).Error() != nil {
		return nil, builder.Error()
	}

	// problemList, err := UnBuildProblems(c.ProblemList)
	// if err != nil {
	// 	return nil, err
	// }
	// contest.ProblemList = problemList

	contest.StartTime = time.UnixMilli(c.StartTime)
	contest.EndTime = time.UnixMilli(c.EndTime)
	return contest, nil
}

func BuildContestList(contests []*model.Contest) ([]*contest.Contest, error) {
	var contestList []*contest.Contest
	for _, c := range contests {
		contest, err := BuildContest(c)
		if err != nil {
			return nil, err
		}
		contestList = append(contestList, contest)
	}
	return contestList, nil
}

func UnBuildContestList(contests []*contest.Contest) ([]*model.Contest, error) {
	var contestList []*model.Contest
	for _, c := range contests {
		contest, err := UnBuildContest(c)
		if err != nil {
			return nil, err
		}
		contestList = append(contestList, contest)
	}
	return contestList, nil
}