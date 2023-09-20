// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: api/submit.proto

package rpcSubmit

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
	SubmitService_Debug_FullMethodName           = "/submit.SubmitService/Debug"
	SubmitService_Submit_FullMethodName          = "/submit.SubmitService/Submit"
	SubmitService_SubmitContest_FullMethodName   = "/submit.SubmitService/SubmitContest"
	SubmitService_GetSubmitResult_FullMethodName = "/submit.SubmitService/GetSubmitResult"
	SubmitService_GetSubmitList_FullMethodName   = "/submit.SubmitService/GetSubmitList"
	SubmitService_GetSubmit_FullMethodName       = "/submit.SubmitService/GetSubmit"
	SubmitService_DeleteSubmit_FullMethodName    = "/submit.SubmitService/DeleteSubmit"
)

// SubmitServiceClient is the client API for SubmitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmitServiceClient interface {
	Debug(ctx context.Context, in *DebugReqeust, opts ...grpc.CallOption) (*DebugResponse, error)
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	SubmitContest(ctx context.Context, in *SubmitContestRequest, opts ...grpc.CallOption) (*SubmitContestResponse, error)
	GetSubmitResult(ctx context.Context, in *GetSubmitResultRequest, opts ...grpc.CallOption) (*GetSubmitResultResponse, error)
	GetSubmitList(ctx context.Context, in *GetSubmitListRequest, opts ...grpc.CallOption) (*GetSubmitListResponse, error)
	GetSubmit(ctx context.Context, in *GetSubmitRequest, opts ...grpc.CallOption) (*GetSubmitResponse, error)
	DeleteSubmit(ctx context.Context, in *DeleteSubmitRequest, opts ...grpc.CallOption) (*DeleteSubmitResponse, error)
}

type submitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmitServiceClient(cc grpc.ClientConnInterface) SubmitServiceClient {
	return &submitServiceClient{cc}
}

func (c *submitServiceClient) Debug(ctx context.Context, in *DebugReqeust, opts ...grpc.CallOption) (*DebugResponse, error) {
	out := new(DebugResponse)
	err := c.cc.Invoke(ctx, SubmitService_Debug_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitServiceClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, SubmitService_Submit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitServiceClient) SubmitContest(ctx context.Context, in *SubmitContestRequest, opts ...grpc.CallOption) (*SubmitContestResponse, error) {
	out := new(SubmitContestResponse)
	err := c.cc.Invoke(ctx, SubmitService_SubmitContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitServiceClient) GetSubmitResult(ctx context.Context, in *GetSubmitResultRequest, opts ...grpc.CallOption) (*GetSubmitResultResponse, error) {
	out := new(GetSubmitResultResponse)
	err := c.cc.Invoke(ctx, SubmitService_GetSubmitResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitServiceClient) GetSubmitList(ctx context.Context, in *GetSubmitListRequest, opts ...grpc.CallOption) (*GetSubmitListResponse, error) {
	out := new(GetSubmitListResponse)
	err := c.cc.Invoke(ctx, SubmitService_GetSubmitList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitServiceClient) GetSubmit(ctx context.Context, in *GetSubmitRequest, opts ...grpc.CallOption) (*GetSubmitResponse, error) {
	out := new(GetSubmitResponse)
	err := c.cc.Invoke(ctx, SubmitService_GetSubmit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submitServiceClient) DeleteSubmit(ctx context.Context, in *DeleteSubmitRequest, opts ...grpc.CallOption) (*DeleteSubmitResponse, error) {
	out := new(DeleteSubmitResponse)
	err := c.cc.Invoke(ctx, SubmitService_DeleteSubmit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmitServiceServer is the server API for SubmitService service.
// All implementations must embed UnimplementedSubmitServiceServer
// for forward compatibility
type SubmitServiceServer interface {
	Debug(context.Context, *DebugReqeust) (*DebugResponse, error)
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	SubmitContest(context.Context, *SubmitContestRequest) (*SubmitContestResponse, error)
	GetSubmitResult(context.Context, *GetSubmitResultRequest) (*GetSubmitResultResponse, error)
	GetSubmitList(context.Context, *GetSubmitListRequest) (*GetSubmitListResponse, error)
	GetSubmit(context.Context, *GetSubmitRequest) (*GetSubmitResponse, error)
	DeleteSubmit(context.Context, *DeleteSubmitRequest) (*DeleteSubmitResponse, error)
	mustEmbedUnimplementedSubmitServiceServer()
}

// UnimplementedSubmitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubmitServiceServer struct {
}

func (UnimplementedSubmitServiceServer) Debug(context.Context, *DebugReqeust) (*DebugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Debug not implemented")
}
func (UnimplementedSubmitServiceServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedSubmitServiceServer) SubmitContest(context.Context, *SubmitContestRequest) (*SubmitContestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitContest not implemented")
}
func (UnimplementedSubmitServiceServer) GetSubmitResult(context.Context, *GetSubmitResultRequest) (*GetSubmitResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmitResult not implemented")
}
func (UnimplementedSubmitServiceServer) GetSubmitList(context.Context, *GetSubmitListRequest) (*GetSubmitListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmitList not implemented")
}
func (UnimplementedSubmitServiceServer) GetSubmit(context.Context, *GetSubmitRequest) (*GetSubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmit not implemented")
}
func (UnimplementedSubmitServiceServer) DeleteSubmit(context.Context, *DeleteSubmitRequest) (*DeleteSubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubmit not implemented")
}
func (UnimplementedSubmitServiceServer) mustEmbedUnimplementedSubmitServiceServer() {}

// UnsafeSubmitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmitServiceServer will
// result in compilation errors.
type UnsafeSubmitServiceServer interface {
	mustEmbedUnimplementedSubmitServiceServer()
}

func RegisterSubmitServiceServer(s grpc.ServiceRegistrar, srv SubmitServiceServer) {
	s.RegisterService(&SubmitService_ServiceDesc, srv)
}

func _SubmitService_Debug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebugReqeust)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).Debug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmitService_Debug_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).Debug(ctx, req.(*DebugReqeust))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmitService_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmitService_Submit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmitService_SubmitContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).SubmitContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmitService_SubmitContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).SubmitContest(ctx, req.(*SubmitContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmitService_GetSubmitResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmitResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).GetSubmitResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmitService_GetSubmitResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).GetSubmitResult(ctx, req.(*GetSubmitResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmitService_GetSubmitList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmitListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).GetSubmitList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmitService_GetSubmitList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).GetSubmitList(ctx, req.(*GetSubmitListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmitService_GetSubmit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).GetSubmit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmitService_GetSubmit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).GetSubmit(ctx, req.(*GetSubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubmitService_DeleteSubmit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmitServiceServer).DeleteSubmit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubmitService_DeleteSubmit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmitServiceServer).DeleteSubmit(ctx, req.(*DeleteSubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubmitService_ServiceDesc is the grpc.ServiceDesc for SubmitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubmitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "submit.SubmitService",
	HandlerType: (*SubmitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Debug",
			Handler:    _SubmitService_Debug_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _SubmitService_Submit_Handler,
		},
		{
			MethodName: "SubmitContest",
			Handler:    _SubmitService_SubmitContest_Handler,
		},
		{
			MethodName: "GetSubmitResult",
			Handler:    _SubmitService_GetSubmitResult_Handler,
		},
		{
			MethodName: "GetSubmitList",
			Handler:    _SubmitService_GetSubmitList_Handler,
		},
		{
			MethodName: "GetSubmit",
			Handler:    _SubmitService_GetSubmit_Handler,
		},
		{
			MethodName: "DeleteSubmit",
			Handler:    _SubmitService_DeleteSubmit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/submit.proto",
}
