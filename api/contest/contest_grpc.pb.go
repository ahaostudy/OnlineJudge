// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: api/contest.proto

package rpcContest

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ContestService_GetContest_FullMethodName    = "/contest.ContestService/GetContest"
	ContestService_CreateContest_FullMethodName = "/contest.ContestService/CreateContest"
	ContestService_DeleteContest_FullMethodName = "/contest.ContestService/DeleteContest"
	ContestService_UpdateContest_FullMethodName = "/contest.ContestService/UpdateContest"
	ContestService_IsRegister_FullMethodName    = "/contest.ContestService/IsRegister"
	ContestService_IsAccessible_FullMethodName  = "/contest.ContestService/IsAccessible"
)

// ContestServiceClient is the client API for ContestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContestServiceClient interface {
	GetContest(ctx context.Context, in *GetContestRequest, opts ...grpc.CallOption) (*GetContestResponse, error)
	CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*CreateContestResponse, error)
	DeleteContest(ctx context.Context, in *DeleteContestRequest, opts ...grpc.CallOption) (*DeleteContestResponse, error)
	UpdateContest(ctx context.Context, in *UpdateContestRequest, opts ...grpc.CallOption) (*UpdateContestResponse, error)
	IsRegister(ctx context.Context, in *IsRegisterRequest, opts ...grpc.CallOption) (*IsRegisterResponse, error)
	IsAccessible(ctx context.Context, in *IsAccessibleRequest, opts ...grpc.CallOption) (*IsAccessibleResponse, error)
}

type contestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContestServiceClient(cc grpc.ClientConnInterface) ContestServiceClient {
	return &contestServiceClient{cc}
}

func (c *contestServiceClient) GetContest(ctx context.Context, in *GetContestRequest, opts ...grpc.CallOption) (*GetContestResponse, error) {
	out := new(GetContestResponse)
	err := c.cc.Invoke(ctx, ContestService_GetContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*CreateContestResponse, error) {
	out := new(CreateContestResponse)
	err := c.cc.Invoke(ctx, ContestService_CreateContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) DeleteContest(ctx context.Context, in *DeleteContestRequest, opts ...grpc.CallOption) (*DeleteContestResponse, error) {
	out := new(DeleteContestResponse)
	err := c.cc.Invoke(ctx, ContestService_DeleteContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) UpdateContest(ctx context.Context, in *UpdateContestRequest, opts ...grpc.CallOption) (*UpdateContestResponse, error) {
	out := new(UpdateContestResponse)
	err := c.cc.Invoke(ctx, ContestService_UpdateContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) IsRegister(ctx context.Context, in *IsRegisterRequest, opts ...grpc.CallOption) (*IsRegisterResponse, error) {
	out := new(IsRegisterResponse)
	err := c.cc.Invoke(ctx, ContestService_IsRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) IsAccessible(ctx context.Context, in *IsAccessibleRequest, opts ...grpc.CallOption) (*IsAccessibleResponse, error) {
	out := new(IsAccessibleResponse)
	err := c.cc.Invoke(ctx, ContestService_IsAccessible_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContestServiceServer is the server API for ContestService service.
// All implementations must embed UnimplementedContestServiceServer
// for forward compatibility
type ContestServiceServer interface {
	GetContest(context.Context, *GetContestRequest) (*GetContestResponse, error)
	CreateContest(context.Context, *CreateContestRequest) (*CreateContestResponse, error)
	DeleteContest(context.Context, *DeleteContestRequest) (*DeleteContestResponse, error)
	UpdateContest(context.Context, *UpdateContestRequest) (*UpdateContestResponse, error)
	IsRegister(context.Context, *IsRegisterRequest) (*IsRegisterResponse, error)
	IsAccessible(context.Context, *IsAccessibleRequest) (*IsAccessibleResponse, error)
	mustEmbedUnimplementedContestServiceServer()
}

// UnimplementedContestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContestServiceServer struct {
}

func (UnimplementedContestServiceServer) GetContest(context.Context, *GetContestRequest) (*GetContestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContest not implemented")
}
func (UnimplementedContestServiceServer) CreateContest(context.Context, *CreateContestRequest) (*CreateContestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContest not implemented")
}
func (UnimplementedContestServiceServer) DeleteContest(context.Context, *DeleteContestRequest) (*DeleteContestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContest not implemented")
}
func (UnimplementedContestServiceServer) UpdateContest(context.Context, *UpdateContestRequest) (*UpdateContestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContest not implemented")
}
func (UnimplementedContestServiceServer) IsRegister(context.Context, *IsRegisterRequest) (*IsRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsRegister not implemented")
}
func (UnimplementedContestServiceServer) IsAccessible(context.Context, *IsAccessibleRequest) (*IsAccessibleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAccessible not implemented")
}
func (UnimplementedContestServiceServer) mustEmbedUnimplementedContestServiceServer() {}

// UnsafeContestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContestServiceServer will
// result in compilation errors.
type UnsafeContestServiceServer interface {
	mustEmbedUnimplementedContestServiceServer()
}

func RegisterContestServiceServer(s grpc.ServiceRegistrar, srv ContestServiceServer) {
	s.RegisterService(&ContestService_ServiceDesc, srv)
}

func _ContestService_GetContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).GetContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_GetContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).GetContest(ctx, req.(*GetContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_CreateContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).CreateContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_CreateContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).CreateContest(ctx, req.(*CreateContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_DeleteContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).DeleteContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_DeleteContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).DeleteContest(ctx, req.(*DeleteContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_UpdateContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).UpdateContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_UpdateContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).UpdateContest(ctx, req.(*UpdateContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_IsRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).IsRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_IsRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).IsRegister(ctx, req.(*IsRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_IsAccessible_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAccessibleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).IsAccessible(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_IsAccessible_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).IsAccessible(ctx, req.(*IsAccessibleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContestService_ServiceDesc is the grpc.ServiceDesc for ContestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contest.ContestService",
	HandlerType: (*ContestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContest",
			Handler:    _ContestService_GetContest_Handler,
		},
		{
			MethodName: "CreateContest",
			Handler:    _ContestService_CreateContest_Handler,
		},
		{
			MethodName: "DeleteContest",
			Handler:    _ContestService_DeleteContest_Handler,
		},
		{
			MethodName: "UpdateContest",
			Handler:    _ContestService_UpdateContest_Handler,
		},
		{
			MethodName: "IsRegister",
			Handler:    _ContestService_IsRegister_Handler,
		},
		{
			MethodName: "IsAccessible",
			Handler:    _ContestService_IsAccessible_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/contest.proto",
}