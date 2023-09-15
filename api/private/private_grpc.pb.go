// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: private.proto

package rpcPrivate

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
	PrivateService_GetProblem_FullMethodName = "/private.PrivateService/GetProblem"
	PrivateService_Judge_FullMethodName      = "/private.PrivateService/judge"
	PrivateService_GetResult_FullMethodName  = "/private.PrivateService/getResult"
	PrivateService_Debug_FullMethodName      = "/private.PrivateService/debug"
)

// PrivateServiceClient is the client API for PrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivateServiceClient interface {
	GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemResponse, error)
	Judge(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*JudgeResponse, error)
	GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error)
	Debug(ctx context.Context, in *DebugRequest, opts ...grpc.CallOption) (*DebugResponse, error)
}

type privateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivateServiceClient(cc grpc.ClientConnInterface) PrivateServiceClient {
	return &privateServiceClient{cc}
}

func (c *privateServiceClient) GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemResponse, error) {
	out := new(GetProblemResponse)
	err := c.cc.Invoke(ctx, PrivateService_GetProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) Judge(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*JudgeResponse, error) {
	out := new(JudgeResponse)
	err := c.cc.Invoke(ctx, PrivateService_Judge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error) {
	out := new(GetResultResponse)
	err := c.cc.Invoke(ctx, PrivateService_GetResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) Debug(ctx context.Context, in *DebugRequest, opts ...grpc.CallOption) (*DebugResponse, error) {
	out := new(DebugResponse)
	err := c.cc.Invoke(ctx, PrivateService_Debug_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateServiceServer is the server API for PrivateService service.
// All implementations must embed UnimplementedPrivateServiceServer
// for forward compatibility
type PrivateServiceServer interface {
	GetProblem(context.Context, *GetProblemRequest) (*GetProblemResponse, error)
	Judge(context.Context, *JudgeRequest) (*JudgeResponse, error)
	GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error)
	Debug(context.Context, *DebugRequest) (*DebugResponse, error)
	mustEmbedUnimplementedPrivateServiceServer()
}

// UnimplementedPrivateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrivateServiceServer struct {
}

func (UnimplementedPrivateServiceServer) GetProblem(context.Context, *GetProblemRequest) (*GetProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblem not implemented")
}
func (UnimplementedPrivateServiceServer) Judge(context.Context, *JudgeRequest) (*JudgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Judge not implemented")
}
func (UnimplementedPrivateServiceServer) GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedPrivateServiceServer) Debug(context.Context, *DebugRequest) (*DebugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Debug not implemented")
}
func (UnimplementedPrivateServiceServer) mustEmbedUnimplementedPrivateServiceServer() {}

// UnsafePrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivateServiceServer will
// result in compilation errors.
type UnsafePrivateServiceServer interface {
	mustEmbedUnimplementedPrivateServiceServer()
}

func RegisterPrivateServiceServer(s grpc.ServiceRegistrar, srv PrivateServiceServer) {
	s.RegisterService(&PrivateService_ServiceDesc, srv)
}

func _PrivateService_GetProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateService_GetProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetProblem(ctx, req.(*GetProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_Judge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JudgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).Judge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateService_Judge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).Judge(ctx, req.(*JudgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateService_GetResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetResult(ctx, req.(*GetResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_Debug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).Debug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateService_Debug_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).Debug(ctx, req.(*DebugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivateService_ServiceDesc is the grpc.ServiceDesc for PrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "private.PrivateService",
	HandlerType: (*PrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProblem",
			Handler:    _PrivateService_GetProblem_Handler,
		},
		{
			MethodName: "judge",
			Handler:    _PrivateService_Judge_Handler,
		},
		{
			MethodName: "getResult",
			Handler:    _PrivateService_GetResult_Handler,
		},
		{
			MethodName: "debug",
			Handler:    _PrivateService_Debug_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "private.proto",
}
