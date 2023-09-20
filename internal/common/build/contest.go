package build

import (
	rpcContest "main/api/contest"
	"main/internal/data/model"
	"time"
)

func BuildContest(c *model.Contest) (*rpcContest.Contest, error) {
	contest := new(rpcContest.Contest)
	builder := new(Builder)
	if builder.Build(c, contest).Error() != nil {
		return nil, builder.Error()
	}
	contest.StartTime = c.StartTime.UnixMilli()
	contest.EndTime = c.EndTime.UnixMilli()
	return contest, nil
}

func UnBuildContest(c *rpcContest.Contest) (*model.Contest, error) {
	contest := new(model.Contest)
	builder := new(Builder)
	if builder.Build(c, contest).Error() != nil {
		return nil, builder.Error()
	}

	problemList, err := UnBuildProblems(c.ProblemList)
	if err != nil {
		return nil, err
	}
	contest.ProblemList = problemList

	contest.StartTime = time.UnixMilli(c.StartTime)
	contest.EndTime = time.UnixMilli(c.EndTime)
	return contest, nil
}

func BuildContestList(contests []*model.Contest) ([]*rpcContest.Contest, error) {
	var contestList []*rpcContest.Contest
	for _, c := range contests {
		contest, err := BuildContest(c)
		if err != nil {
			return nil, err
		}
		contestList = append(contestList, contest)
	}
	return contestList, nil
}

func UnBuildContestList(contests []*rpcContest.Contest) ([]*model.Contest, error) {
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
