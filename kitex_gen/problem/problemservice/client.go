// Code generated by Kitex v0.7.2. DO NOT EDIT.

package problemservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	problem "main/kitex_gen/problem"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetProblem(ctx context.Context, Req *problem.GetProblemRequest, callOptions ...callopt.Option) (r *problem.GetProblemResponse, err error)
	GetProblemList(ctx context.Context, Req *problem.GetProblemListRequest, callOptions ...callopt.Option) (r *problem.GetProblemListResponse, err error)
	GetProblemCount(ctx context.Context, Req *problem.GetProblemCountRequest, callOptions ...callopt.Option) (r *problem.GetProblemCountResponse, err error)
	GetContestProblem(ctx context.Context, Req *problem.GetContestProblemRequest, callOptions ...callopt.Option) (r *problem.GetContestProblemResponse, err error)
	GetContestProblemList(ctx context.Context, Req *problem.GetContestProblemListRequest, callOptions ...callopt.Option) (r *problem.GetContestProblemListResponse, err error)
	CreateProblem(ctx context.Context, Req *problem.CreateProblemRequest, callOptions ...callopt.Option) (r *problem.CreateProblemResponse, err error)
	DeleteProblem(ctx context.Context, Req *problem.DeleteProblemRequest, callOptions ...callopt.Option) (r *problem.DeleteProblemResponse, err error)
	UpdateProblem(ctx context.Context, Req *problem.UpdateProblemRequest, callOptions ...callopt.Option) (r *problem.UpdateProblemResponse, err error)
	CreateTestcase(ctx context.Context, Req *problem.CreateTestcaseRequest, callOptions ...callopt.Option) (r *problem.CreateTestcaseResponse, err error)
	GetTestcase(ctx context.Context, Req *problem.GetTestcaseRequest, callOptions ...callopt.Option) (r *problem.GetTestcaseResponse, err error)
	DeleteTestcase(ctx context.Context, Req *problem.DeleteTestcaseRequest, callOptions ...callopt.Option) (r *problem.DeleteTestcaseResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProblemServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kProblemServiceClient struct {
	*kClient
}

func (p *kProblemServiceClient) GetProblem(ctx context.Context, Req *problem.GetProblemRequest, callOptions ...callopt.Option) (r *problem.GetProblemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProblem(ctx, Req)
}

func (p *kProblemServiceClient) GetProblemList(ctx context.Context, Req *problem.GetProblemListRequest, callOptions ...callopt.Option) (r *problem.GetProblemListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProblemList(ctx, Req)
}

func (p *kProblemServiceClient) GetProblemCount(ctx context.Context, Req *problem.GetProblemCountRequest, callOptions ...callopt.Option) (r *problem.GetProblemCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProblemCount(ctx, Req)
}

func (p *kProblemServiceClient) GetContestProblem(ctx context.Context, Req *problem.GetContestProblemRequest, callOptions ...callopt.Option) (r *problem.GetContestProblemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetContestProblem(ctx, Req)
}

func (p *kProblemServiceClient) GetContestProblemList(ctx context.Context, Req *problem.GetContestProblemListRequest, callOptions ...callopt.Option) (r *problem.GetContestProblemListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetContestProblemList(ctx, Req)
}

func (p *kProblemServiceClient) CreateProblem(ctx context.Context, Req *problem.CreateProblemRequest, callOptions ...callopt.Option) (r *problem.CreateProblemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateProblem(ctx, Req)
}

func (p *kProblemServiceClient) DeleteProblem(ctx context.Context, Req *problem.DeleteProblemRequest, callOptions ...callopt.Option) (r *problem.DeleteProblemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteProblem(ctx, Req)
}

func (p *kProblemServiceClient) UpdateProblem(ctx context.Context, Req *problem.UpdateProblemRequest, callOptions ...callopt.Option) (r *problem.UpdateProblemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateProblem(ctx, Req)
}

func (p *kProblemServiceClient) CreateTestcase(ctx context.Context, Req *problem.CreateTestcaseRequest, callOptions ...callopt.Option) (r *problem.CreateTestcaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateTestcase(ctx, Req)
}

func (p *kProblemServiceClient) GetTestcase(ctx context.Context, Req *problem.GetTestcaseRequest, callOptions ...callopt.Option) (r *problem.GetTestcaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTestcase(ctx, Req)
}

func (p *kProblemServiceClient) DeleteTestcase(ctx context.Context, Req *problem.DeleteTestcaseRequest, callOptions ...callopt.Option) (r *problem.DeleteTestcaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTestcase(ctx, Req)
}
