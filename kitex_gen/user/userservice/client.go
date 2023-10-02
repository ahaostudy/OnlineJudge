// Code generated by Kitex v0.7.2. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "main/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, Req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	Login(ctx context.Context, Req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error)
	CreateUser(ctx context.Context, Req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error)
	UpdateUser(ctx context.Context, Req *user.UpdateUserRequest, callOptions ...callopt.Option) (r *user.UpdateUserResponse, err error)
	GetCaptcha(ctx context.Context, Req *user.GetCaptchaRequest, callOptions ...callopt.Option) (r *user.GetCaptchaResponse, err error)
	IsAdmin(ctx context.Context, Req *user.IsAdminRequest, callOptions ...callopt.Option) (r *user.IsAdminResponse, err error)
	GetUser(ctx context.Context, Req *user.GetUserRequest, callOptions ...callopt.Option) (r *user.GetUserResponse, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, Req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}

func (p *kUserServiceClient) Login(ctx context.Context, Req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, Req)
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, Req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, Req)
}

func (p *kUserServiceClient) UpdateUser(ctx context.Context, Req *user.UpdateUserRequest, callOptions ...callopt.Option) (r *user.UpdateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, Req)
}

func (p *kUserServiceClient) GetCaptcha(ctx context.Context, Req *user.GetCaptchaRequest, callOptions ...callopt.Option) (r *user.GetCaptchaResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCaptcha(ctx, Req)
}

func (p *kUserServiceClient) IsAdmin(ctx context.Context, Req *user.IsAdminRequest, callOptions ...callopt.Option) (r *user.IsAdminResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsAdmin(ctx, Req)
}

func (p *kUserServiceClient) GetUser(ctx context.Context, Req *user.GetUserRequest, callOptions ...callopt.Option) (r *user.GetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, Req)
}