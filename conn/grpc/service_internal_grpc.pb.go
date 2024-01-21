// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: service_internal.proto

package grpc

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

// InternalAppServiceClient is the client API for InternalAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalAppServiceClient interface {
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (InternalAppService_ListTransactionsClient, error)
}

type internalAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalAppServiceClient(cc grpc.ClientConnInterface) InternalAppServiceClient {
	return &internalAppServiceClient{cc}
}

func (c *internalAppServiceClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (InternalAppService_ListTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &InternalAppService_ServiceDesc.Streams[0], "/grpc.InternalAppService/ListTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &internalAppServiceListTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InternalAppService_ListTransactionsClient interface {
	Recv() (*ListTransactionsStreamResponse, error)
	grpc.ClientStream
}

type internalAppServiceListTransactionsClient struct {
	grpc.ClientStream
}

func (x *internalAppServiceListTransactionsClient) Recv() (*ListTransactionsStreamResponse, error) {
	m := new(ListTransactionsStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InternalAppServiceServer is the server API for InternalAppService service.
// All implementations must embed UnimplementedInternalAppServiceServer
// for forward compatibility
type InternalAppServiceServer interface {
	ListTransactions(*ListTransactionsRequest, InternalAppService_ListTransactionsServer) error
	mustEmbedUnimplementedInternalAppServiceServer()
}

// UnimplementedInternalAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInternalAppServiceServer struct {
}

func (UnimplementedInternalAppServiceServer) ListTransactions(*ListTransactionsRequest, InternalAppService_ListTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTransactions not implemented")
}
func (UnimplementedInternalAppServiceServer) mustEmbedUnimplementedInternalAppServiceServer() {}

// UnsafeInternalAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalAppServiceServer will
// result in compilation errors.
type UnsafeInternalAppServiceServer interface {
	mustEmbedUnimplementedInternalAppServiceServer()
}

func RegisterInternalAppServiceServer(s grpc.ServiceRegistrar, srv InternalAppServiceServer) {
	s.RegisterService(&InternalAppService_ServiceDesc, srv)
}

func _InternalAppService_ListTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTransactionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InternalAppServiceServer).ListTransactions(m, &internalAppServiceListTransactionsServer{stream})
}

type InternalAppService_ListTransactionsServer interface {
	Send(*ListTransactionsStreamResponse) error
	grpc.ServerStream
}

type internalAppServiceListTransactionsServer struct {
	grpc.ServerStream
}

func (x *internalAppServiceListTransactionsServer) Send(m *ListTransactionsStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// InternalAppService_ServiceDesc is the grpc.ServiceDesc for InternalAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.InternalAppService",
	HandlerType: (*InternalAppServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTransactions",
			Handler:       _InternalAppService_ListTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service_internal.proto",
}
