// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb_gen

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

// RemoteCodeServiceClient is the client API for RemoteCodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteCodeServiceClient interface {
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
	DownloadRemoteCode(ctx context.Context, in *DownloadRemoteCodeRequest, opts ...grpc.CallOption) (*DownloadRemoteCodeResponse, error)
}

type remoteCodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteCodeServiceClient(cc grpc.ClientConnInterface) RemoteCodeServiceClient {
	return &remoteCodeServiceClient{cc}
}

func (c *remoteCodeServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/pb.RemoteCodeService/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteCodeServiceClient) DownloadRemoteCode(ctx context.Context, in *DownloadRemoteCodeRequest, opts ...grpc.CallOption) (*DownloadRemoteCodeResponse, error) {
	out := new(DownloadRemoteCodeResponse)
	err := c.cc.Invoke(ctx, "/pb.RemoteCodeService/DownloadRemoteCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteCodeServiceServer is the server API for RemoteCodeService service.
// All implementations must embed UnimplementedRemoteCodeServiceServer
// for forward compatibility
type RemoteCodeServiceServer interface {
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	DownloadRemoteCode(context.Context, *DownloadRemoteCodeRequest) (*DownloadRemoteCodeResponse, error)
	mustEmbedUnimplementedRemoteCodeServiceServer()
}

// UnimplementedRemoteCodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteCodeServiceServer struct {
}

func (UnimplementedRemoteCodeServiceServer) HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedRemoteCodeServiceServer) DownloadRemoteCode(context.Context, *DownloadRemoteCodeRequest) (*DownloadRemoteCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadRemoteCode not implemented")
}
func (UnimplementedRemoteCodeServiceServer) mustEmbedUnimplementedRemoteCodeServiceServer() {}

// UnsafeRemoteCodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteCodeServiceServer will
// result in compilation errors.
type UnsafeRemoteCodeServiceServer interface {
	mustEmbedUnimplementedRemoteCodeServiceServer()
}

func RegisterRemoteCodeServiceServer(s grpc.ServiceRegistrar, srv RemoteCodeServiceServer) {
	s.RegisterService(&RemoteCodeService_ServiceDesc, srv)
}

func _RemoteCodeService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteCodeServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RemoteCodeService/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteCodeServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteCodeService_DownloadRemoteCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRemoteCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteCodeServiceServer).DownloadRemoteCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RemoteCodeService/DownloadRemoteCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteCodeServiceServer).DownloadRemoteCode(ctx, req.(*DownloadRemoteCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteCodeService_ServiceDesc is the grpc.ServiceDesc for RemoteCodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteCodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RemoteCodeService",
	HandlerType: (*RemoteCodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _RemoteCodeService_HelloWorld_Handler,
		},
		{
			MethodName: "DownloadRemoteCode",
			Handler:    _RemoteCodeService_DownloadRemoteCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote_code_service.proto",
}
