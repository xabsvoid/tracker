// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: geo/v1/service.proto

package geov1

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
	GeoService_LocateByUUID_FullMethodName = "/geo.v1.GeoService/LocateByUUID"
	GeoService_Track_FullMethodName        = "/geo.v1.GeoService/Track"
)

// GeoServiceClient is the client API for GeoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeoServiceClient interface {
	LocateByUUID(ctx context.Context, in *LocateByUUIDRequest, opts ...grpc.CallOption) (*LocateByUUIDResponse, error)
	Track(ctx context.Context, in *TrackRequest, opts ...grpc.CallOption) (*TrackResponse, error)
}

type geoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoServiceClient(cc grpc.ClientConnInterface) GeoServiceClient {
	return &geoServiceClient{cc}
}

func (c *geoServiceClient) LocateByUUID(ctx context.Context, in *LocateByUUIDRequest, opts ...grpc.CallOption) (*LocateByUUIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LocateByUUIDResponse)
	err := c.cc.Invoke(ctx, GeoService_LocateByUUID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoServiceClient) Track(ctx context.Context, in *TrackRequest, opts ...grpc.CallOption) (*TrackResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TrackResponse)
	err := c.cc.Invoke(ctx, GeoService_Track_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoServiceServer is the server API for GeoService service.
// All implementations must embed UnimplementedGeoServiceServer
// for forward compatibility.
type GeoServiceServer interface {
	LocateByUUID(context.Context, *LocateByUUIDRequest) (*LocateByUUIDResponse, error)
	Track(context.Context, *TrackRequest) (*TrackResponse, error)
	mustEmbedUnimplementedGeoServiceServer()
}

// UnimplementedGeoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGeoServiceServer struct{}

func (UnimplementedGeoServiceServer) LocateByUUID(context.Context, *LocateByUUIDRequest) (*LocateByUUIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocateByUUID not implemented")
}
func (UnimplementedGeoServiceServer) Track(context.Context, *TrackRequest) (*TrackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Track not implemented")
}
func (UnimplementedGeoServiceServer) mustEmbedUnimplementedGeoServiceServer() {}
func (UnimplementedGeoServiceServer) testEmbeddedByValue()                    {}

// UnsafeGeoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeoServiceServer will
// result in compilation errors.
type UnsafeGeoServiceServer interface {
	mustEmbedUnimplementedGeoServiceServer()
}

func RegisterGeoServiceServer(s grpc.ServiceRegistrar, srv GeoServiceServer) {
	// If the following call pancis, it indicates UnimplementedGeoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GeoService_ServiceDesc, srv)
}

func _GeoService_LocateByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocateByUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).LocateByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeoService_LocateByUUID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).LocateByUUID(ctx, req.(*LocateByUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoService_Track_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).Track(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeoService_Track_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).Track(ctx, req.(*TrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeoService_ServiceDesc is the grpc.ServiceDesc for GeoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geo.v1.GeoService",
	HandlerType: (*GeoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LocateByUUID",
			Handler:    _GeoService_LocateByUUID_Handler,
		},
		{
			MethodName: "Track",
			Handler:    _GeoService_Track_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geo/v1/service.proto",
}
