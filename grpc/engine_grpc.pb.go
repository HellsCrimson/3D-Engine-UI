// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: grpc/engine.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Engine_GetObjects_FullMethodName   = "/grpc.Engine/GetObjects"
	Engine_AddObject_FullMethodName    = "/grpc.Engine/AddObject"
	Engine_RemoveObject_FullMethodName = "/grpc.Engine/RemoveObject"
	Engine_MoveObject_FullMethodName   = "/grpc.Engine/MoveObject"
	Engine_RotateObject_FullMethodName = "/grpc.Engine/RotateObject"
	Engine_ScaleObject_FullMethodName  = "/grpc.Engine/ScaleObject"
	Engine_UpdateObject_FullMethodName = "/grpc.Engine/UpdateObject"
)

// EngineClient is the client API for Engine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EngineClient interface {
	GetObjects(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Objects, error)
	AddObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MoveObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RotateObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ScaleObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type engineClient struct {
	cc grpc.ClientConnInterface
}

func NewEngineClient(cc grpc.ClientConnInterface) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) GetObjects(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Objects, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Objects)
	err := c.cc.Invoke(ctx, Engine_GetObjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) AddObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_AddObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) RemoveObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_RemoveObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) MoveObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_MoveObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) RotateObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_RotateObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) ScaleObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_ScaleObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) UpdateObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Engine_UpdateObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineServer is the server API for Engine service.
// All implementations must embed UnimplementedEngineServer
// for forward compatibility.
type EngineServer interface {
	GetObjects(context.Context, *emptypb.Empty) (*Objects, error)
	AddObject(context.Context, *Object) (*emptypb.Empty, error)
	RemoveObject(context.Context, *Object) (*emptypb.Empty, error)
	MoveObject(context.Context, *Object) (*emptypb.Empty, error)
	RotateObject(context.Context, *Object) (*emptypb.Empty, error)
	ScaleObject(context.Context, *Object) (*emptypb.Empty, error)
	UpdateObject(context.Context, *Object) (*emptypb.Empty, error)
	mustEmbedUnimplementedEngineServer()
}

// UnimplementedEngineServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEngineServer struct{}

func (UnimplementedEngineServer) GetObjects(context.Context, *emptypb.Empty) (*Objects, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjects not implemented")
}
func (UnimplementedEngineServer) AddObject(context.Context, *Object) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddObject not implemented")
}
func (UnimplementedEngineServer) RemoveObject(context.Context, *Object) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveObject not implemented")
}
func (UnimplementedEngineServer) MoveObject(context.Context, *Object) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveObject not implemented")
}
func (UnimplementedEngineServer) RotateObject(context.Context, *Object) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateObject not implemented")
}
func (UnimplementedEngineServer) ScaleObject(context.Context, *Object) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScaleObject not implemented")
}
func (UnimplementedEngineServer) UpdateObject(context.Context, *Object) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObject not implemented")
}
func (UnimplementedEngineServer) mustEmbedUnimplementedEngineServer() {}
func (UnimplementedEngineServer) testEmbeddedByValue()                {}

// UnsafeEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EngineServer will
// result in compilation errors.
type UnsafeEngineServer interface {
	mustEmbedUnimplementedEngineServer()
}

func RegisterEngineServer(s grpc.ServiceRegistrar, srv EngineServer) {
	// If the following call pancis, it indicates UnimplementedEngineServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Engine_ServiceDesc, srv)
}

func _Engine_GetObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).GetObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_GetObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).GetObjects(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_AddObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).AddObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_AddObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).AddObject(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_RemoveObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).RemoveObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_RemoveObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).RemoveObject(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_MoveObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).MoveObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_MoveObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).MoveObject(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_RotateObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).RotateObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_RotateObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).RotateObject(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_ScaleObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).ScaleObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_ScaleObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).ScaleObject(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_UpdateObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).UpdateObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_UpdateObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).UpdateObject(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

// Engine_ServiceDesc is the grpc.ServiceDesc for Engine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Engine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetObjects",
			Handler:    _Engine_GetObjects_Handler,
		},
		{
			MethodName: "AddObject",
			Handler:    _Engine_AddObject_Handler,
		},
		{
			MethodName: "RemoveObject",
			Handler:    _Engine_RemoveObject_Handler,
		},
		{
			MethodName: "MoveObject",
			Handler:    _Engine_MoveObject_Handler,
		},
		{
			MethodName: "RotateObject",
			Handler:    _Engine_RotateObject_Handler,
		},
		{
			MethodName: "ScaleObject",
			Handler:    _Engine_ScaleObject_Handler,
		},
		{
			MethodName: "UpdateObject",
			Handler:    _Engine_UpdateObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/engine.proto",
}
