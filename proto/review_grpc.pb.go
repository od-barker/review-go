// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: review.proto

package proto

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

// ReviewClient is the client API for Review service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewClient interface {
	ProcessReview(ctx context.Context, in *MyReview, opts ...grpc.CallOption) (*MyResponse, error)
}

type reviewClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewClient(cc grpc.ClientConnInterface) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) ProcessReview(ctx context.Context, in *MyReview, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/Review/ProcessReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServer is the server API for Review service.
// All implementations must embed UnimplementedReviewServer
// for forward compatibility
type ReviewServer interface {
	ProcessReview(context.Context, *MyReview) (*MyResponse, error)
	mustEmbedUnimplementedReviewServer()
}

// UnimplementedReviewServer must be embedded to have forward compatible implementations.
type UnimplementedReviewServer struct {
}

func (UnimplementedReviewServer) ProcessReview(context.Context, *MyReview) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessReview not implemented")
}
func (UnimplementedReviewServer) mustEmbedUnimplementedReviewServer() {}

// UnsafeReviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServer will
// result in compilation errors.
type UnsafeReviewServer interface {
	mustEmbedUnimplementedReviewServer()
}

func RegisterReviewServer(s grpc.ServiceRegistrar, srv ReviewServer) {
	s.RegisterService(&Review_ServiceDesc, srv)
}

func _Review_ProcessReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ProcessReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Review/ProcessReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ProcessReview(ctx, req.(*MyReview))
	}
	return interceptor(ctx, in, info, handler)
}

// Review_ServiceDesc is the grpc.ServiceDesc for Review service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Review_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessReview",
			Handler:    _Review_ProcessReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review.proto",
}
