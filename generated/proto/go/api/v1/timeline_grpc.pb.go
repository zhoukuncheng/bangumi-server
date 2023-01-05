// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/timeline.proto

package api

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

// TimeLineServiceClient is the client API for TimeLineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeLineServiceClient interface {
	// Debug function
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SubjectCollect(ctx context.Context, in *SubjectCollectRequest, opts ...grpc.CallOption) (*SubjectCollectResponse, error)
	SubjectProgress(ctx context.Context, in *SubjectProgressRequest, opts ...grpc.CallOption) (*SubjectProgressResponse, error)
	EpisodeCollect(ctx context.Context, in *EpisodeCollectRequest, opts ...grpc.CallOption) (*EpisodeCollectResponse, error)
}

type timeLineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeLineServiceClient(cc grpc.ClientConnInterface) TimeLineServiceClient {
	return &timeLineServiceClient{cc}
}

func (c *timeLineServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/api.v1.TimeLineService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeLineServiceClient) SubjectCollect(ctx context.Context, in *SubjectCollectRequest, opts ...grpc.CallOption) (*SubjectCollectResponse, error) {
	out := new(SubjectCollectResponse)
	err := c.cc.Invoke(ctx, "/api.v1.TimeLineService/SubjectCollect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeLineServiceClient) SubjectProgress(ctx context.Context, in *SubjectProgressRequest, opts ...grpc.CallOption) (*SubjectProgressResponse, error) {
	out := new(SubjectProgressResponse)
	err := c.cc.Invoke(ctx, "/api.v1.TimeLineService/SubjectProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeLineServiceClient) EpisodeCollect(ctx context.Context, in *EpisodeCollectRequest, opts ...grpc.CallOption) (*EpisodeCollectResponse, error) {
	out := new(EpisodeCollectResponse)
	err := c.cc.Invoke(ctx, "/api.v1.TimeLineService/EpisodeCollect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeLineServiceServer is the server API for TimeLineService service.
// All implementations must embed UnimplementedTimeLineServiceServer
// for forward compatibility
type TimeLineServiceServer interface {
	// Debug function
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	SubjectCollect(context.Context, *SubjectCollectRequest) (*SubjectCollectResponse, error)
	SubjectProgress(context.Context, *SubjectProgressRequest) (*SubjectProgressResponse, error)
	EpisodeCollect(context.Context, *EpisodeCollectRequest) (*EpisodeCollectResponse, error)
	mustEmbedUnimplementedTimeLineServiceServer()
}

// UnimplementedTimeLineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeLineServiceServer struct {
}

func (UnimplementedTimeLineServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedTimeLineServiceServer) SubjectCollect(context.Context, *SubjectCollectRequest) (*SubjectCollectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubjectCollect not implemented")
}
func (UnimplementedTimeLineServiceServer) SubjectProgress(context.Context, *SubjectProgressRequest) (*SubjectProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubjectProgress not implemented")
}
func (UnimplementedTimeLineServiceServer) EpisodeCollect(context.Context, *EpisodeCollectRequest) (*EpisodeCollectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpisodeCollect not implemented")
}
func (UnimplementedTimeLineServiceServer) mustEmbedUnimplementedTimeLineServiceServer() {}

// UnsafeTimeLineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeLineServiceServer will
// result in compilation errors.
type UnsafeTimeLineServiceServer interface {
	mustEmbedUnimplementedTimeLineServiceServer()
}

func RegisterTimeLineServiceServer(s grpc.ServiceRegistrar, srv TimeLineServiceServer) {
	s.RegisterService(&TimeLineService_ServiceDesc, srv)
}

func _TimeLineService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeLineServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.TimeLineService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeLineServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeLineService_SubjectCollect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubjectCollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeLineServiceServer).SubjectCollect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.TimeLineService/SubjectCollect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeLineServiceServer).SubjectCollect(ctx, req.(*SubjectCollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeLineService_SubjectProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubjectProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeLineServiceServer).SubjectProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.TimeLineService/SubjectProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeLineServiceServer).SubjectProgress(ctx, req.(*SubjectProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeLineService_EpisodeCollect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EpisodeCollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeLineServiceServer).EpisodeCollect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.TimeLineService/EpisodeCollect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeLineServiceServer).EpisodeCollect(ctx, req.(*EpisodeCollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeLineService_ServiceDesc is the grpc.ServiceDesc for TimeLineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeLineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.TimeLineService",
	HandlerType: (*TimeLineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _TimeLineService_Hello_Handler,
		},
		{
			MethodName: "SubjectCollect",
			Handler:    _TimeLineService_SubjectCollect_Handler,
		},
		{
			MethodName: "SubjectProgress",
			Handler:    _TimeLineService_SubjectProgress_Handler,
		},
		{
			MethodName: "EpisodeCollect",
			Handler:    _TimeLineService_EpisodeCollect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/timeline.proto",
}
