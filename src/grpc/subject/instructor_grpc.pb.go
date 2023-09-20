// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: instructor.proto

package subject

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

// InstructorServiceClient is the client API for InstructorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstructorServiceClient interface {
	PaginateInstructor(ctx context.Context, in *PaginateInstructorRequest, opts ...grpc.CallOption) (*PaginateInstructorResponse, error)
	GetInstructorbyId(ctx context.Context, in *GetInstructorbyIdRequest, opts ...grpc.CallOption) (*GetInstructorbyIdResponse, error)
	CreateInstructor(ctx context.Context, in *CreateInstructorRequest, opts ...grpc.CallOption) (*CreateInstructorResponse, error)
	UpdateInstructor(ctx context.Context, in *UpdateInstructorRequest, opts ...grpc.CallOption) (*UpdateInstructorResponse, error)
	DeleteInstructor(ctx context.Context, in *DeleteInstructorRequest, opts ...grpc.CallOption) (*DeleteInstructorResponse, error)
}

type instructorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstructorServiceClient(cc grpc.ClientConnInterface) InstructorServiceClient {
	return &instructorServiceClient{cc}
}

func (c *instructorServiceClient) PaginateInstructor(ctx context.Context, in *PaginateInstructorRequest, opts ...grpc.CallOption) (*PaginateInstructorResponse, error) {
	out := new(PaginateInstructorResponse)
	err := c.cc.Invoke(ctx, "/InstructorService/PaginateInstructor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instructorServiceClient) GetInstructorbyId(ctx context.Context, in *GetInstructorbyIdRequest, opts ...grpc.CallOption) (*GetInstructorbyIdResponse, error) {
	out := new(GetInstructorbyIdResponse)
	err := c.cc.Invoke(ctx, "/InstructorService/GetInstructorbyId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instructorServiceClient) CreateInstructor(ctx context.Context, in *CreateInstructorRequest, opts ...grpc.CallOption) (*CreateInstructorResponse, error) {
	out := new(CreateInstructorResponse)
	err := c.cc.Invoke(ctx, "/InstructorService/CreateInstructor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instructorServiceClient) UpdateInstructor(ctx context.Context, in *UpdateInstructorRequest, opts ...grpc.CallOption) (*UpdateInstructorResponse, error) {
	out := new(UpdateInstructorResponse)
	err := c.cc.Invoke(ctx, "/InstructorService/UpdateInstructor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instructorServiceClient) DeleteInstructor(ctx context.Context, in *DeleteInstructorRequest, opts ...grpc.CallOption) (*DeleteInstructorResponse, error) {
	out := new(DeleteInstructorResponse)
	err := c.cc.Invoke(ctx, "/InstructorService/DeleteInstructor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstructorServiceServer is the server API for InstructorService service.
// All implementations must embed UnimplementedInstructorServiceServer
// for forward compatibility
type InstructorServiceServer interface {
	PaginateInstructor(context.Context, *PaginateInstructorRequest) (*PaginateInstructorResponse, error)
	GetInstructorbyId(context.Context, *GetInstructorbyIdRequest) (*GetInstructorbyIdResponse, error)
	CreateInstructor(context.Context, *CreateInstructorRequest) (*CreateInstructorResponse, error)
	UpdateInstructor(context.Context, *UpdateInstructorRequest) (*UpdateInstructorResponse, error)
	DeleteInstructor(context.Context, *DeleteInstructorRequest) (*DeleteInstructorResponse, error)
	mustEmbedUnimplementedInstructorServiceServer()
}

// UnimplementedInstructorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInstructorServiceServer struct {
}

func (UnimplementedInstructorServiceServer) PaginateInstructor(context.Context, *PaginateInstructorRequest) (*PaginateInstructorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaginateInstructor not implemented")
}
func (UnimplementedInstructorServiceServer) GetInstructorbyId(context.Context, *GetInstructorbyIdRequest) (*GetInstructorbyIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstructorbyId not implemented")
}
func (UnimplementedInstructorServiceServer) CreateInstructor(context.Context, *CreateInstructorRequest) (*CreateInstructorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstructor not implemented")
}
func (UnimplementedInstructorServiceServer) UpdateInstructor(context.Context, *UpdateInstructorRequest) (*UpdateInstructorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstructor not implemented")
}
func (UnimplementedInstructorServiceServer) DeleteInstructor(context.Context, *DeleteInstructorRequest) (*DeleteInstructorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstructor not implemented")
}
func (UnimplementedInstructorServiceServer) mustEmbedUnimplementedInstructorServiceServer() {}

// UnsafeInstructorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstructorServiceServer will
// result in compilation errors.
type UnsafeInstructorServiceServer interface {
	mustEmbedUnimplementedInstructorServiceServer()
}

func RegisterInstructorServiceServer(s grpc.ServiceRegistrar, srv InstructorServiceServer) {
	s.RegisterService(&InstructorService_ServiceDesc, srv)
}

func _InstructorService_PaginateInstructor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginateInstructorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstructorServiceServer).PaginateInstructor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstructorService/PaginateInstructor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstructorServiceServer).PaginateInstructor(ctx, req.(*PaginateInstructorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstructorService_GetInstructorbyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstructorbyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstructorServiceServer).GetInstructorbyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstructorService/GetInstructorbyId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstructorServiceServer).GetInstructorbyId(ctx, req.(*GetInstructorbyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstructorService_CreateInstructor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstructorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstructorServiceServer).CreateInstructor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstructorService/CreateInstructor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstructorServiceServer).CreateInstructor(ctx, req.(*CreateInstructorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstructorService_UpdateInstructor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstructorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstructorServiceServer).UpdateInstructor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstructorService/UpdateInstructor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstructorServiceServer).UpdateInstructor(ctx, req.(*UpdateInstructorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstructorService_DeleteInstructor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstructorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstructorServiceServer).DeleteInstructor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstructorService/DeleteInstructor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstructorServiceServer).DeleteInstructor(ctx, req.(*DeleteInstructorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InstructorService_ServiceDesc is the grpc.ServiceDesc for InstructorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstructorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InstructorService",
	HandlerType: (*InstructorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PaginateInstructor",
			Handler:    _InstructorService_PaginateInstructor_Handler,
		},
		{
			MethodName: "GetInstructorbyId",
			Handler:    _InstructorService_GetInstructorbyId_Handler,
		},
		{
			MethodName: "CreateInstructor",
			Handler:    _InstructorService_CreateInstructor_Handler,
		},
		{
			MethodName: "UpdateInstructor",
			Handler:    _InstructorService_UpdateInstructor_Handler,
		},
		{
			MethodName: "DeleteInstructor",
			Handler:    _InstructorService_DeleteInstructor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "instructor.proto",
}