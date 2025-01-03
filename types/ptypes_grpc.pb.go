// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: types/ptypes.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DistanceAggregator_AggregateDistance_FullMethodName = "/DistanceAggregator/AggregateDistance"
	DistanceAggregator_GetInvoice_FullMethodName        = "/DistanceAggregator/GetInvoice"
)

// DistanceAggregatorClient is the client API for DistanceAggregator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DistanceAggregatorClient interface {
	AggregateDistance(ctx context.Context, in *AggregatorDistanceRequest, opts ...grpc.CallOption) (*None, error)
	GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*GetInvoiceResponse, error)
}

type distanceAggregatorClient struct {
	cc grpc.ClientConnInterface
}

func NewDistanceAggregatorClient(cc grpc.ClientConnInterface) DistanceAggregatorClient {
	return &distanceAggregatorClient{cc}
}

func (c *distanceAggregatorClient) AggregateDistance(ctx context.Context, in *AggregatorDistanceRequest, opts ...grpc.CallOption) (*None, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(None)
	err := c.cc.Invoke(ctx, DistanceAggregator_AggregateDistance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distanceAggregatorClient) GetInvoice(ctx context.Context, in *GetInvoiceRequest, opts ...grpc.CallOption) (*GetInvoiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInvoiceResponse)
	err := c.cc.Invoke(ctx, DistanceAggregator_GetInvoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistanceAggregatorServer is the server API for DistanceAggregator service.
// All implementations must embed UnimplementedDistanceAggregatorServer
// for forward compatibility.
type DistanceAggregatorServer interface {
	AggregateDistance(context.Context, *AggregatorDistanceRequest) (*None, error)
	GetInvoice(context.Context, *GetInvoiceRequest) (*GetInvoiceResponse, error)
	mustEmbedUnimplementedDistanceAggregatorServer()
}

// UnimplementedDistanceAggregatorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDistanceAggregatorServer struct{}

func (UnimplementedDistanceAggregatorServer) AggregateDistance(context.Context, *AggregatorDistanceRequest) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AggregateDistance not implemented")
}
func (UnimplementedDistanceAggregatorServer) GetInvoice(context.Context, *GetInvoiceRequest) (*GetInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoice not implemented")
}
func (UnimplementedDistanceAggregatorServer) mustEmbedUnimplementedDistanceAggregatorServer() {}
func (UnimplementedDistanceAggregatorServer) testEmbeddedByValue()                            {}

// UnsafeDistanceAggregatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DistanceAggregatorServer will
// result in compilation errors.
type UnsafeDistanceAggregatorServer interface {
	mustEmbedUnimplementedDistanceAggregatorServer()
}

func RegisterDistanceAggregatorServer(s grpc.ServiceRegistrar, srv DistanceAggregatorServer) {
	// If the following call pancis, it indicates UnimplementedDistanceAggregatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DistanceAggregator_ServiceDesc, srv)
}

func _DistanceAggregator_AggregateDistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregatorDistanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistanceAggregatorServer).AggregateDistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DistanceAggregator_AggregateDistance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistanceAggregatorServer).AggregateDistance(ctx, req.(*AggregatorDistanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistanceAggregator_GetInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistanceAggregatorServer).GetInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DistanceAggregator_GetInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistanceAggregatorServer).GetInvoice(ctx, req.(*GetInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DistanceAggregator_ServiceDesc is the grpc.ServiceDesc for DistanceAggregator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DistanceAggregator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DistanceAggregator",
	HandlerType: (*DistanceAggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AggregateDistance",
			Handler:    _DistanceAggregator_AggregateDistance_Handler,
		},
		{
			MethodName: "GetInvoice",
			Handler:    _DistanceAggregator_GetInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types/ptypes.proto",
}
