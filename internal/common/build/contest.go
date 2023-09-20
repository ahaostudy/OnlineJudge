package build

import (
	rpcContest "main/api/contest"
	"main/internal/data/model"
)

func BuildContest(c *model.Contest) (*rpcContest.Contest, error) {
	contest := new(rpcContest.Contest)
	return contest, new(Builder).Build(c, contest).Error()
}

func UnBuildContest(c *rpcContest.Contest) (*model.Contest, error) {
	contest := new(model.Contest)
	return contest, new(Builder).Build(c, contest).Error()
}

func BuildContestList(contests []*model.Contest) ([]*rpcContest.Contest, error) {
	contestList := make([]*rpcContest.Contest, 0)
	builder := new(Builder)
	for _, c := range contests {
		var contest *rpcContest.Contest
		builder.Build(c, contest)
		if builder.Error() != nil {
			return nil, builder.Error()
		}
		contestList = append(contestList, contest)
	}
	return contestList, builder.Error()
}

func UnBuildContestList(contests []*rpcContest.Contest) ([]*model.Contest, error) {
	contestList := make([]*model.Contest, 0)
	builder := new(Builder)
	for _, c := range contests {
		var contest *model.Contest
		builder.Build(c, contest)
		if builder.Error() != nil {
			return nil, builder.Error()
		}
		contestList = append(contestList, contest)
	}
	return contestList, builder.Error()
}