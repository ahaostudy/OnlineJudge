package build

import (
	rpcContest "main/api/contest"
)

func BuildRankStatus(s any) (*rpcContest.Status, error) {
	status := new(rpcContest.Status)
	return status, new(Builder).Build(s, status).Error()
}
