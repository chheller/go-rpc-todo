// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: modules/todo/todo.proto

package todo

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
	Greeter_FindTodo_FullMethodName   = "/todo.Greeter/FindTodo"
	Greeter_ListTodos_FullMethodName  = "/todo.Greeter/ListTodos"
	Greeter_CreateTodo_FullMethodName = "/todo.Greeter/CreateTodo"
	Greeter_UpdateTodo_FullMethodName = "/todo.Greeter/UpdateTodo"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	FindTodo(ctx context.Context, in *FindTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	ListTodos(ctx context.Context, in *ListTodosRequest, opts ...grpc.CallOption) (*TodoResponsePage, error)
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) FindTodo(ctx context.Context, in *FindTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, Greeter_FindTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ListTodos(ctx context.Context, in *ListTodosRequest, opts ...grpc.CallOption) (*TodoResponsePage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoResponsePage)
	err := c.cc.Invoke(ctx, Greeter_ListTodos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, Greeter_CreateTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, Greeter_UpdateTodo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility.
type GreeterServer interface {
	FindTodo(context.Context, *FindTodoRequest) (*TodoResponse, error)
	ListTodos(context.Context, *ListTodosRequest) (*TodoResponsePage, error)
	CreateTodo(context.Context, *CreateTodoRequest) (*TodoResponse, error)
	UpdateTodo(context.Context, *UpdateTodoRequest) (*TodoResponse, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGreeterServer struct{}

func (UnimplementedGreeterServer) FindTodo(context.Context, *FindTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTodo not implemented")
}
func (UnimplementedGreeterServer) ListTodos(context.Context, *ListTodosRequest) (*TodoResponsePage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodos not implemented")
}
func (UnimplementedGreeterServer) CreateTodo(context.Context, *CreateTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedGreeterServer) UpdateTodo(context.Context, *UpdateTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}
func (UnimplementedGreeterServer) testEmbeddedByValue()                 {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	// If the following call pancis, it indicates UnimplementedGreeterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_FindTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).FindTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_FindTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).FindTodo(ctx, req.(*FindTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ListTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ListTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_ListTodos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ListTodos(ctx, req.(*ListTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_CreateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_UpdateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateTodo(ctx, req.(*UpdateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindTodo",
			Handler:    _Greeter_FindTodo_Handler,
		},
		{
			MethodName: "ListTodos",
			Handler:    _Greeter_ListTodos_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _Greeter_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _Greeter_UpdateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/todo/todo.proto",
}
