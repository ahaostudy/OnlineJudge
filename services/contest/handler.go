package main

import (
	"context"
	contest "main/kitex_gen/contest"
)

// ContestServiceImpl implements the last service interface defined in the IDL.
type ContestServiceImpl struct{}

// GetContest implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) GetContest(ctx context.Context, req *contest.GetContestRequest) (resp *contest.GetContestResponse, err error) {
	// TODO: Your code here...
	return
}

// GetContestList implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) GetContestList(ctx context.Context, req *contest.GetContestListRequest) (resp *contest.GetContestListResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateContest implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) CreateContest(ctx context.Context, req *contest.CreateContestRequest) (resp *contest.CreateContestResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteContest implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) DeleteContest(ctx context.Context, req *contest.DeleteContestRequest) (resp *contest.DeleteContestResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateContest implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) UpdateContest(ctx context.Context, req *contest.UpdateContestRequest) (resp *contest.UpdateContestResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) Register(ctx context.Context, req *contest.RegisterRequest) (resp *contest.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UnRegister implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) UnRegister(ctx context.Context, req *contest.UnRegisterRequest) (resp *contest.UnRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// IsRegister implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) IsRegister(ctx context.Context, req *contest.IsRegisterRequest) (resp *contest.IsRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// IsAccessible implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) IsAccessible(ctx context.Context, req *contest.IsAccessibleRequest) (resp *contest.IsAccessibleResponse, err error) {
	// TODO: Your code here...
	return
}

// ContestRank implements the ContestServiceImpl interface.
func (s *ContestServiceImpl) ContestRank(ctx context.Context, req *contest.ContestRankRequest) (resp *contest.ContestRankResponse, err error) {
	// TODO: Your code here...
	return
}
