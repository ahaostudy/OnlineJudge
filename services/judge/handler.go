package main

import (
	"context"
	judge "main/kitex_gen/judge"
)

// JudgeServiceImpl implements the last service interface defined in the IDL.
type JudgeServiceImpl struct{}

// Judge implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) Judge(ctx context.Context, req *judge.JudgeRequest) (resp *judge.JudgeResponse, err error) {
	// TODO: Your code here...
	return
}

// GetResult implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) GetResult(ctx context.Context, req *judge.GetResultRequest) (resp *judge.GetResultResponse, err error) {
	// TODO: Your code here...
	return
}

// Debug implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) Debug(ctx context.Context, req *judge.DebugRequest) (resp *judge.DebugResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCode implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) GetCode(ctx context.Context, req *judge.GetCodeRequest) (resp *judge.GetCodeResponse, err error) {
	// TODO: Your code here...
	return
}
