package main

import (
	"context"
	submit "main/kitex_gen/submit"
)

// SubmitServiceImpl implements the last service interface defined in the IDL.
type SubmitServiceImpl struct{}

// Debug implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) Debug(ctx context.Context, req *submit.DebugReqeust) (resp *submit.DebugResponse, err error) {
	// TODO: Your code here...
	return
}

// Submit implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) Submit(ctx context.Context, req *submit.SubmitRequest) (resp *submit.SubmitResponse, err error) {
	// TODO: Your code here...
	return
}

// SubmitContest implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) SubmitContest(ctx context.Context, req *submit.SubmitContestRequest) (resp *submit.SubmitContestResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSubmitResult implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitResult(ctx context.Context, req *submit.GetSubmitResultRequest) (resp *submit.GetSubmitResultResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSubmitList implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitList(ctx context.Context, req *submit.GetSubmitListRequest) (resp *submit.GetSubmitListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSubmit implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmit(ctx context.Context, req *submit.GetSubmitRequest) (resp *submit.GetSubmitResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSubmitStatus implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitStatus(ctx context.Context, req *submit.GetSubmitStatusRequest) (resp *submit.GetSubmitStatusResponse, err error) {
	// TODO: Your code here...
	return
}

// IsAccepted implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) IsAccepted(ctx context.Context, req *submit.IsAcceptedRequest) (resp *submit.IsAcceptedResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAcceptedStatus implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetAcceptedStatus(ctx context.Context, req *submit.GetAcceptedStatusRequest) (resp *submit.GetAcceptedStatusResponse, err error) {
	// TODO: Your code here...
	return
}

// GetLatestSubmits implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetLatestSubmits(ctx context.Context, req *submit.GetLatestSubmitsRequest) (resp *submit.GetLatestSubmitsResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteSubmit implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) DeleteSubmit(ctx context.Context, req *submit.DeleteSubmitRequest) (resp *submit.DeleteSubmitResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSubmitCalendar implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetSubmitCalendar(ctx context.Context, req *submit.GetSubmitCalendarRequest) (resp *submit.GetSubmitCalendarResponse, err error) {
	// TODO: Your code here...
	return
}

// GetNote implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetNote(ctx context.Context, req *submit.GetNoteRequest) (resp *submit.GetNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// GetNoteList implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) GetNoteList(ctx context.Context, req *submit.GetNoteListRequest) (resp *submit.GetNoteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateNote implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) CreateNote(ctx context.Context, req *submit.CreateNoteRequest) (resp *submit.CreateNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteNote implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) DeleteNote(ctx context.Context, req *submit.DeleteNoteRequest) (resp *submit.DeleteNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateNote implements the SubmitServiceImpl interface.
func (s *SubmitServiceImpl) UpdateNote(ctx context.Context, req *submit.UpdateNoteRequest) (resp *submit.UpdateNoteResponse, err error) {
	// TODO: Your code here...
	return
}
