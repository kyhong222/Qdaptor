// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: Qdaptor_grpc.proto

package Qdaptor_grpc

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

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionClient interface {
	HelloTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*TransactionMessage, error)
	RefCallTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*TransactionMessage, error)
	CallClearTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*TransactionMessage, error)
	GetQueueTrafficTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*TransactionMessage, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) HelloTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*TransactionMessage, error) {
	out := new(TransactionMessage)
	err := c.cc.Invoke(ctx, "/Qdaptor_grpc.Transaction/HelloTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) RefCallTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*TransactionMessage, error) {
	out := new(TransactionMessage)
	err := c.cc.Invoke(ctx, "/Qdaptor_grpc.Transaction/RefCallTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) CallClearTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*TransactionMessage, error) {
	out := new(TransactionMessage)
	err := c.cc.Invoke(ctx, "/Qdaptor_grpc.Transaction/CallClearTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) GetQueueTrafficTransaction(ctx context.Context, in *TransactionMessage, opts ...grpc.CallOption) (*TransactionMessage, error) {
	out := new(TransactionMessage)
	err := c.cc.Invoke(ctx, "/Qdaptor_grpc.Transaction/GetQueueTrafficTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
// All implementations must embed UnimplementedTransactionServer
// for forward compatibility
type TransactionServer interface {
	HelloTransaction(context.Context, *TransactionMessage) (*TransactionMessage, error)
	RefCallTransaction(context.Context, *TransactionMessage) (*TransactionMessage, error)
	CallClearTransaction(context.Context, *TransactionMessage) (*TransactionMessage, error)
	GetQueueTrafficTransaction(context.Context, *TransactionMessage) (*TransactionMessage, error)
	mustEmbedUnimplementedTransactionServer()
}

// UnimplementedTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (UnimplementedTransactionServer) HelloTransaction(context.Context, *TransactionMessage) (*TransactionMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloTransaction not implemented")
}
func (UnimplementedTransactionServer) RefCallTransaction(context.Context, *TransactionMessage) (*TransactionMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefCallTransaction not implemented")
}
func (UnimplementedTransactionServer) CallClearTransaction(context.Context, *TransactionMessage) (*TransactionMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallClearTransaction not implemented")
}
func (UnimplementedTransactionServer) GetQueueTrafficTransaction(context.Context, *TransactionMessage) (*TransactionMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueueTrafficTransaction not implemented")
}
func (UnimplementedTransactionServer) mustEmbedUnimplementedTransactionServer() {}

// UnsafeTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServer will
// result in compilation errors.
type UnsafeTransactionServer interface {
	mustEmbedUnimplementedTransactionServer()
}

func RegisterTransactionServer(s grpc.ServiceRegistrar, srv TransactionServer) {
	s.RegisterService(&Transaction_ServiceDesc, srv)
}

func _Transaction_HelloTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).HelloTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Qdaptor_grpc.Transaction/HelloTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).HelloTransaction(ctx, req.(*TransactionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_RefCallTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).RefCallTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Qdaptor_grpc.Transaction/RefCallTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).RefCallTransaction(ctx, req.(*TransactionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_CallClearTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).CallClearTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Qdaptor_grpc.Transaction/CallClearTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).CallClearTransaction(ctx, req.(*TransactionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_GetQueueTrafficTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetQueueTrafficTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Qdaptor_grpc.Transaction/GetQueueTrafficTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetQueueTrafficTransaction(ctx, req.(*TransactionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Transaction_ServiceDesc is the grpc.ServiceDesc for Transaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Qdaptor_grpc.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloTransaction",
			Handler:    _Transaction_HelloTransaction_Handler,
		},
		{
			MethodName: "RefCallTransaction",
			Handler:    _Transaction_RefCallTransaction_Handler,
		},
		{
			MethodName: "CallClearTransaction",
			Handler:    _Transaction_CallClearTransaction_Handler,
		},
		{
			MethodName: "GetQueueTrafficTransaction",
			Handler:    _Transaction_GetQueueTrafficTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Qdaptor_grpc.proto",
}
