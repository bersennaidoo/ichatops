// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: jaeger/sampling.proto

package jaeger

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

// SamplingManagerClient is the client API for SamplingManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SamplingManagerClient interface {
	GetSamplingStrategy(ctx context.Context, in *SamplingStrategyParameters, opts ...grpc.CallOption) (*SamplingStrategyResponse, error)
}

type samplingManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewSamplingManagerClient(cc grpc.ClientConnInterface) SamplingManagerClient {
	return &samplingManagerClient{cc}
}

func (c *samplingManagerClient) GetSamplingStrategy(ctx context.Context, in *SamplingStrategyParameters, opts ...grpc.CallOption) (*SamplingStrategyResponse, error) {
	out := new(SamplingStrategyResponse)
	err := c.cc.Invoke(ctx, "/jaeger.api_v2.SamplingManager/GetSamplingStrategy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SamplingManagerServer is the server API for SamplingManager service.
// All implementations must embed UnimplementedSamplingManagerServer
// for forward compatibility
type SamplingManagerServer interface {
	GetSamplingStrategy(context.Context, *SamplingStrategyParameters) (*SamplingStrategyResponse, error)
	mustEmbedUnimplementedSamplingManagerServer()
}

// UnimplementedSamplingManagerServer must be embedded to have forward compatible implementations.
type UnimplementedSamplingManagerServer struct {
}

func (UnimplementedSamplingManagerServer) GetSamplingStrategy(context.Context, *SamplingStrategyParameters) (*SamplingStrategyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSamplingStrategy not implemented")
}
func (UnimplementedSamplingManagerServer) mustEmbedUnimplementedSamplingManagerServer() {}

// UnsafeSamplingManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SamplingManagerServer will
// result in compilation errors.
type UnsafeSamplingManagerServer interface {
	mustEmbedUnimplementedSamplingManagerServer()
}

func RegisterSamplingManagerServer(s grpc.ServiceRegistrar, srv SamplingManagerServer) {
	s.RegisterService(&SamplingManager_ServiceDesc, srv)
}

func _SamplingManager_GetSamplingStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SamplingStrategyParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SamplingManagerServer).GetSamplingStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.api_v2.SamplingManager/GetSamplingStrategy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SamplingManagerServer).GetSamplingStrategy(ctx, req.(*SamplingStrategyParameters))
	}
	return interceptor(ctx, in, info, handler)
}

// SamplingManager_ServiceDesc is the grpc.ServiceDesc for SamplingManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SamplingManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.api_v2.SamplingManager",
	HandlerType: (*SamplingManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSamplingStrategy",
			Handler:    _SamplingManager_GetSamplingStrategy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jaeger/sampling.proto",
}
