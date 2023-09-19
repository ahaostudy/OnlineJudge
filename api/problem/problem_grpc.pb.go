// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: api/problem.proto

package rpcProblem

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
	ProblemService_GetProblem_FullMethodName            = "/problem.ProblemService/GetProblem"
	ProblemService_GetContestProblem_FullMethodName     = "/problem.ProblemService/GetContestProblem"
	ProblemService_GetContestProblemList_FullMethodName = "/problem.ProblemService/GetContestProblemList"
	ProblemService_CreateTestcase_FullMethodName        = "/problem.ProblemService/CreateTestcase"
	ProblemService_GetTestcase_FullMethodName           = "/problem.ProblemService/GetTestcase"
	ProblemService_DeleteTestcase_FullMethodName        = "/problem.ProblemService/DeleteTestcase"
)

// ProblemServiceClient is the client API for ProblemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemServiceClient interface {
	GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemResponse, error)
	GetContestProblem(ctx context.Context, in *GetContestProblemRequest, opts ...grpc.CallOption) (*GetContestProblemResponse, error)
	GetContestProblemList(ctx context.Context, in *GetContestProblemListRequest, opts ...grpc.CallOption) (*GetContestProblemListResponse, error)
	CreateTestcase(ctx context.Context, in *CreateTestcaseRequest, opts ...grpc.CallOption) (*CreateTestcaseResponse, error)
	GetTestcase(ctx context.Context, in *GetTestcaseRequest, opts ...grpc.CallOption) (*GetTestcaseResponse, error)
	DeleteTestcase(ctx context.Context, in *DeleteTestcaseRequest, opts ...grpc.CallOption) (*DeleteTestcaseResponse, error)
}

type problemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemServiceClient(cc grpc.ClientConnInterface) ProblemServiceClient {
	return &problemServiceClient{cc}
}

func (c *problemServiceClient) GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemResponse, error) {
	out := new(GetProblemResponse)
	err := c.cc.Invoke(ctx, ProblemService_GetProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) GetContestProblem(ctx context.Context, in *GetContestProblemRequest, opts ...grpc.CallOption) (*GetContestProblemResponse, error) {
	out := new(GetContestProblemResponse)
	err := c.cc.Invoke(ctx, ProblemService_GetContestProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) GetContestProblemList(ctx context.Context, in *GetContestProblemListRequest, opts ...grpc.CallOption) (*GetContestProblemListResponse, error) {
	out := new(GetContestProblemListResponse)
	err := c.cc.Invoke(ctx, ProblemService_GetContestProblemList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) CreateTestcase(ctx context.Context, in *CreateTestcaseRequest, opts ...grpc.CallOption) (*CreateTestcaseResponse, error) {
	out := new(CreateTestcaseResponse)
	err := c.cc.Invoke(ctx, ProblemService_CreateTestcase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) GetTestcase(ctx context.Context, in *GetTestcaseRequest, opts ...grpc.CallOption) (*GetTestcaseResponse, error) {
	out := new(GetTestcaseResponse)
	err := c.cc.Invoke(ctx, ProblemService_GetTestcase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DeleteTestcase(ctx context.Context, in *DeleteTestcaseRequest, opts ...grpc.CallOption) (*DeleteTestcaseResponse, error) {
	out := new(DeleteTestcaseResponse)
	err := c.cc.Invoke(ctx, ProblemService_DeleteTestcase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemServiceServer is the server API for ProblemService service.
// All implementations must embed UnimplementedProblemServiceServer
// for forward compatibility
type ProblemServiceServer interface {
	GetProblem(context.Context, *GetProblemRequest) (*GetProblemResponse, error)
	GetContestProblem(context.Context, *GetContestProblemRequest) (*GetContestProblemResponse, error)
	GetContestProblemList(context.Context, *GetContestProblemListRequest) (*GetContestProblemListResponse, error)
	CreateTestcase(context.Context, *CreateTestcaseRequest) (*CreateTestcaseResponse, error)
	GetTestcase(context.Context, *GetTestcaseRequest) (*GetTestcaseResponse, error)
	DeleteTestcase(context.Context, *DeleteTestcaseRequest) (*DeleteTestcaseResponse, error)
	mustEmbedUnimplementedProblemServiceServer()
}

// UnimplementedProblemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProblemServiceServer struct {
}

func (UnimplementedProblemServiceServer) GetProblem(context.Context, *GetProblemRequest) (*GetProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblem not implemented")
}
func (UnimplementedProblemServiceServer) GetContestProblem(context.Context, *GetContestProblemRequest) (*GetContestProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContestProblem not implemented")
}
func (UnimplementedProblemServiceServer) GetContestProblemList(context.Context, *GetContestProblemListRequest) (*GetContestProblemListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContestProblemList not implemented")
}
func (UnimplementedProblemServiceServer) CreateTestcase(context.Context, *CreateTestcaseRequest) (*CreateTestcaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestcase not implemented")
}
func (UnimplementedProblemServiceServer) GetTestcase(context.Context, *GetTestcaseRequest) (*GetTestcaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestcase not implemented")
}
func (UnimplementedProblemServiceServer) DeleteTestcase(context.Context, *DeleteTestcaseRequest) (*DeleteTestcaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTestcase not implemented")
}
func (UnimplementedProblemServiceServer) mustEmbedUnimplementedProblemServiceServer() {}

// UnsafeProblemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemServiceServer will
// result in compilation errors.
type UnsafeProblemServiceServer interface {
	mustEmbedUnimplementedProblemServiceServer()
}

func RegisterProblemServiceServer(s grpc.ServiceRegistrar, srv ProblemServiceServer) {
	s.RegisterService(&ProblemService_ServiceDesc, srv)
}

func _ProblemService_GetProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_GetProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetProblem(ctx, req.(*GetProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_GetContestProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContestProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetContestProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_GetContestProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetContestProblem(ctx, req.(*GetContestProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_GetContestProblemList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContestProblemListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetContestProblemList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_GetContestProblemList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetContestProblemList(ctx, req.(*GetContestProblemListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_CreateTestcase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestcaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).CreateTestcase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_CreateTestcase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).CreateTestcase(ctx, req.(*CreateTestcaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_GetTestcase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestcaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetTestcase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_GetTestcase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetTestcase(ctx, req.(*GetTestcaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DeleteTestcase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTestcaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DeleteTestcase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DeleteTestcase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DeleteTestcase(ctx, req.(*DeleteTestcaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProblemService_ServiceDesc is the grpc.ServiceDesc for ProblemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProblemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "problem.ProblemService",
	HandlerType: (*ProblemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProblem",
			Handler:    _ProblemService_GetProblem_Handler,
		},
		{
			MethodName: "GetContestProblem",
			Handler:    _ProblemService_GetContestProblem_Handler,
		},
		{
			MethodName: "GetContestProblemList",
			Handler:    _ProblemService_GetContestProblemList_Handler,
		},
		{
			MethodName: "CreateTestcase",
			Handler:    _ProblemService_CreateTestcase_Handler,
		},
		{
			MethodName: "GetTestcase",
			Handler:    _ProblemService_GetTestcase_Handler,
		},
		{
			MethodName: "DeleteTestcase",
			Handler:    _ProblemService_DeleteTestcase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/problem.proto",
}
