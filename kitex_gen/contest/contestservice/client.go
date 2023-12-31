// Code generated by Kitex v0.7.2. DO NOT EDIT.

package contestservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	contest "main/kitex_gen/contest"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetContest(ctx context.Context, Req *contest.GetContestRequest, callOptions ...callopt.Option) (r *contest.GetContestResponse, err error)
	GetContestList(ctx context.Context, Req *contest.GetContestListRequest, callOptions ...callopt.Option) (r *contest.GetContestListResponse, err error)
	CreateContest(ctx context.Context, Req *contest.CreateContestRequest, callOptions ...callopt.Option) (r *contest.CreateContestResponse, err error)
	DeleteContest(ctx context.Context, Req *contest.DeleteContestRequest, callOptions ...callopt.Option) (r *contest.DeleteContestResponse, err error)
	UpdateContest(ctx context.Context, Req *contest.UpdateContestRequest, callOptions ...callopt.Option) (r *contest.UpdateContestResponse, err error)
	Register(ctx context.Context, Req *contest.RegisterRequest, callOptions ...callopt.Option) (r *contest.RegisterResponse, err error)
	UnRegister(ctx context.Context, Req *contest.UnRegisterRequest, callOptions ...callopt.Option) (r *contest.UnRegisterResponse, err error)
	IsRegister(ctx context.Context, Req *contest.IsRegisterRequest, callOptions ...callopt.Option) (r *contest.IsRegisterResponse, err error)
	IsAccessible(ctx context.Context, Req *contest.IsAccessibleRequest, callOptions ...callopt.Option) (r *contest.IsAccessibleResponse, err error)
	ContestRank(ctx context.Context, Req *contest.ContestRankRequest, callOptions ...callopt.Option) (r *contest.ContestRankResponse, err error)
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
	return &kContestServiceClient{
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

type kContestServiceClient struct {
	*kClient
}

func (p *kContestServiceClient) GetContest(ctx context.Context, Req *contest.GetContestRequest, callOptions ...callopt.Option) (r *contest.GetContestResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetContest(ctx, Req)
}

func (p *kContestServiceClient) GetContestList(ctx context.Context, Req *contest.GetContestListRequest, callOptions ...callopt.Option) (r *contest.GetContestListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetContestList(ctx, Req)
}

func (p *kContestServiceClient) CreateContest(ctx context.Context, Req *contest.CreateContestRequest, callOptions ...callopt.Option) (r *contest.CreateContestResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateContest(ctx, Req)
}

func (p *kContestServiceClient) DeleteContest(ctx context.Context, Req *contest.DeleteContestRequest, callOptions ...callopt.Option) (r *contest.DeleteContestResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteContest(ctx, Req)
}

func (p *kContestServiceClient) UpdateContest(ctx context.Context, Req *contest.UpdateContestRequest, callOptions ...callopt.Option) (r *contest.UpdateContestResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateContest(ctx, Req)
}

func (p *kContestServiceClient) Register(ctx context.Context, Req *contest.RegisterRequest, callOptions ...callopt.Option) (r *contest.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}

func (p *kContestServiceClient) UnRegister(ctx context.Context, Req *contest.UnRegisterRequest, callOptions ...callopt.Option) (r *contest.UnRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UnRegister(ctx, Req)
}

func (p *kContestServiceClient) IsRegister(ctx context.Context, Req *contest.IsRegisterRequest, callOptions ...callopt.Option) (r *contest.IsRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsRegister(ctx, Req)
}

func (p *kContestServiceClient) IsAccessible(ctx context.Context, Req *contest.IsAccessibleRequest, callOptions ...callopt.Option) (r *contest.IsAccessibleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsAccessible(ctx, Req)
}

func (p *kContestServiceClient) ContestRank(ctx context.Context, Req *contest.ContestRankRequest, callOptions ...callopt.Option) (r *contest.ContestRankResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContestRank(ctx, Req)
}
