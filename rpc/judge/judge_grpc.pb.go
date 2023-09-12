// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: judge.proto

package rpcJudge

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
	JudgeService_Judge_FullMethodName     = "/JudgeService/judge"
	JudgeService_GetResult_FullMethodName = "/JudgeService/getResult"
)

// JudgeServiceClient is the client API for JudgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudgeServiceClient interface {
	Judge(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*JudgeResponse, error)
	GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error)
}

type judgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJudgeServiceClient(cc grpc.ClientConnInterface) JudgeServiceClient {
	return &judgeServiceClient{cc}
}

func (c *judgeServiceClient) Judge(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*JudgeResponse, error) {
	out := new(JudgeResponse)
	err := c.cc.Invoke(ctx, JudgeService_Judge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeServiceClient) GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error) {
	out := new(GetResultResponse)
	err := c.cc.Invoke(ctx, JudgeService_GetResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgeServiceServer is the server API for JudgeService service.
// All implementations must embed UnimplementedJudgeServiceServer
// for forward compatibility
type JudgeServiceServer interface {
	Judge(context.Context, *JudgeRequest) (*JudgeResponse, error)
	GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error)
	mustEmbedUnimplementedJudgeServiceServer()
}

// UnimplementedJudgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJudgeServiceServer struct {
}

func (UnimplementedJudgeServiceServer) Judge(context.Context, *JudgeRequest) (*JudgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Judge not implemented")
}
func (UnimplementedJudgeServiceServer) GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedJudgeServiceServer) mustEmbedUnimplementedJudgeServiceServer() {}

// UnsafeJudgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JudgeServiceServer will
// result in compilation errors.
type UnsafeJudgeServiceServer interface {
	mustEmbedUnimplementedJudgeServiceServer()
}

func RegisterJudgeServiceServer(s grpc.ServiceRegistrar, srv JudgeServiceServer) {
	s.RegisterService(&JudgeService_ServiceDesc, srv)
}

func _JudgeService_Judge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JudgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).Judge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudgeService_Judge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).Judge(ctx, req.(*JudgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudgeService_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudgeService_GetResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).GetResult(ctx, req.(*GetResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JudgeService_ServiceDesc is the grpc.ServiceDesc for JudgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JudgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "JudgeService",
	HandlerType: (*JudgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "judge",
			Handler:    _JudgeService_Judge_Handler,
		},
		{
			MethodName: "getResult",
			Handler:    _JudgeService_GetResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judge.proto",
}
