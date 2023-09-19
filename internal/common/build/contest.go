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
