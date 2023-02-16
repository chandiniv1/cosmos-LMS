// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: proto/cosmos/lms/v1beta1/tx.proto

package types

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	RegisterAdmin(ctx context.Context, in *RegisterAdminRequest, opts ...grpc.CallOption) (*RegisterAdminResponse, error)
	AddStudent(ctx context.Context, in *AddStudentRequest, opts ...grpc.CallOption) (*AddStudentResponse, error)
	ApplyLeave(ctx context.Context, in *ApplyLeaveRequest, opts ...grpc.CallOption) (*ApplyLeaveResponse, error)
	AcceptLeave(ctx context.Context, in *AcceptLeaveRequest, opts ...grpc.CallOption) (*AcceptLeaveResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterAdmin(ctx context.Context, in *RegisterAdminRequest, opts ...grpc.CallOption) (*RegisterAdminResponse, error) {
	out := new(RegisterAdminResponse)
	err := c.cc.Invoke(ctx, "/Msg/RegisterAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddStudent(ctx context.Context, in *AddStudentRequest, opts ...grpc.CallOption) (*AddStudentResponse, error) {
	out := new(AddStudentResponse)
	err := c.cc.Invoke(ctx, "/Msg/AddStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApplyLeave(ctx context.Context, in *ApplyLeaveRequest, opts ...grpc.CallOption) (*ApplyLeaveResponse, error) {
	out := new(ApplyLeaveResponse)
	err := c.cc.Invoke(ctx, "/Msg/ApplyLeave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AcceptLeave(ctx context.Context, in *AcceptLeaveRequest, opts ...grpc.CallOption) (*AcceptLeaveResponse, error) {
	out := new(AcceptLeaveResponse)
	err := c.cc.Invoke(ctx, "/Msg/AcceptLeave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	RegisterAdmin(context.Context, *RegisterAdminRequest) (*RegisterAdminResponse, error)
	AddStudent(context.Context, *AddStudentRequest) (*AddStudentResponse, error)
	ApplyLeave(context.Context, *ApplyLeaveRequest) (*ApplyLeaveResponse, error)
	AcceptLeave(context.Context, *AcceptLeaveRequest) (*AcceptLeaveResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) RegisterAdmin(context.Context, *RegisterAdminRequest) (*RegisterAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAdmin not implemented")
}
func (UnimplementedMsgServer) AddStudent(context.Context, *AddStudentRequest) (*AddStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudent not implemented")
}
func (UnimplementedMsgServer) ApplyLeave(context.Context, *ApplyLeaveRequest) (*ApplyLeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyLeave not implemented")
}
func (UnimplementedMsgServer) AcceptLeave(context.Context, *AcceptLeaveRequest) (*AcceptLeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptLeave not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_RegisterAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Msg/RegisterAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterAdmin(ctx, req.(*RegisterAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Msg/AddStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddStudent(ctx, req.(*AddStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApplyLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApplyLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Msg/ApplyLeave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApplyLeave(ctx, req.(*ApplyLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AcceptLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AcceptLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Msg/AcceptLeave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptLeave(ctx, req.(*AcceptLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAdmin",
			Handler:    _Msg_RegisterAdmin_Handler,
		},
		{
			MethodName: "AddStudent",
			Handler:    _Msg_AddStudent_Handler,
		},
		{
			MethodName: "ApplyLeave",
			Handler:    _Msg_ApplyLeave_Handler,
		},
		{
			MethodName: "AcceptLeave",
			Handler:    _Msg_AcceptLeave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cosmos/lms/v1beta1/tx.proto",
}
