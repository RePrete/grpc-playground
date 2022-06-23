// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/iot.proto

package iot

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IotServerClient is the client API for IotServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IotServerClient interface {
	GetEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (IotServer_GetEventsClient, error)
}

type iotServerClient struct {
	cc grpc.ClientConnInterface
}

func NewIotServerClient(cc grpc.ClientConnInterface) IotServerClient {
	return &iotServerClient{cc}
}

func (c *iotServerClient) GetEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (IotServer_GetEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &IotServer_ServiceDesc.Streams[0], "/iot.IotServer/GetEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &iotServerGetEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IotServer_GetEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type iotServerGetEventsClient struct {
	grpc.ClientStream
}

func (x *iotServerGetEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IotServerServer is the server API for IotServer service.
// All implementations must embed UnimplementedIotServerServer
// for forward compatibility
type IotServerServer interface {
	GetEvents(*emptypb.Empty, IotServer_GetEventsServer) error
	mustEmbedUnimplementedIotServerServer()
}

// UnimplementedIotServerServer must be embedded to have forward compatible implementations.
type UnimplementedIotServerServer struct {
}

func (UnimplementedIotServerServer) GetEvents(*emptypb.Empty, IotServer_GetEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedIotServerServer) mustEmbedUnimplementedIotServerServer() {}

// UnsafeIotServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IotServerServer will
// result in compilation errors.
type UnsafeIotServerServer interface {
	mustEmbedUnimplementedIotServerServer()
}

func RegisterIotServerServer(s grpc.ServiceRegistrar, srv IotServerServer) {
	s.RegisterService(&IotServer_ServiceDesc, srv)
}

func _IotServer_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IotServerServer).GetEvents(m, &iotServerGetEventsServer{stream})
}

type IotServer_GetEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type iotServerGetEventsServer struct {
	grpc.ServerStream
}

func (x *iotServerGetEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// IotServer_ServiceDesc is the grpc.ServiceDesc for IotServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IotServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iot.IotServer",
	HandlerType: (*IotServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _IotServer_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/iot.proto",
}
