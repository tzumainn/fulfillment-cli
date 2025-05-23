//
// Copyright (c) 2025 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/v1/hubs_service.proto

package adminv1

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
	Hubs_List_FullMethodName         = "/admin.v1.Hubs/List"
	Hubs_Get_FullMethodName          = "/admin.v1.Hubs/Get"
	Hubs_Create_FullMethodName       = "/admin.v1.Hubs/Create"
	Hubs_Delete_FullMethodName       = "/admin.v1.Hubs/Delete"
	Hubs_UpdateStatus_FullMethodName = "/admin.v1.Hubs/UpdateStatus"
)

// HubsClient is the client API for Hubs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubsClient interface {
	// Retrieves the list of hubs.
	List(ctx context.Context, in *HubsListRequest, opts ...grpc.CallOption) (*HubsListResponse, error)
	// Retrieves the details of one specific hub.
	Get(ctx context.Context, in *HubsGetRequest, opts ...grpc.CallOption) (*HubsGetResponse, error)
	// Creates a new hub.
	Create(ctx context.Context, in *HubsCreateRequest, opts ...grpc.CallOption) (*HubsCreateResponse, error)
	// Delete a hub.
	Delete(ctx context.Context, in *HubsDeleteRequest, opts ...grpc.CallOption) (*HubsDeleteResponse, error)
	// Updates the status of a hub.
	UpdateStatus(ctx context.Context, in *HubsUpdateStatusRequest, opts ...grpc.CallOption) (*HubsUpdateStatusResponse, error)
}

type hubsClient struct {
	cc grpc.ClientConnInterface
}

func NewHubsClient(cc grpc.ClientConnInterface) HubsClient {
	return &hubsClient{cc}
}

func (c *hubsClient) List(ctx context.Context, in *HubsListRequest, opts ...grpc.CallOption) (*HubsListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HubsListResponse)
	err := c.cc.Invoke(ctx, Hubs_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubsClient) Get(ctx context.Context, in *HubsGetRequest, opts ...grpc.CallOption) (*HubsGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HubsGetResponse)
	err := c.cc.Invoke(ctx, Hubs_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubsClient) Create(ctx context.Context, in *HubsCreateRequest, opts ...grpc.CallOption) (*HubsCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HubsCreateResponse)
	err := c.cc.Invoke(ctx, Hubs_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubsClient) Delete(ctx context.Context, in *HubsDeleteRequest, opts ...grpc.CallOption) (*HubsDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HubsDeleteResponse)
	err := c.cc.Invoke(ctx, Hubs_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubsClient) UpdateStatus(ctx context.Context, in *HubsUpdateStatusRequest, opts ...grpc.CallOption) (*HubsUpdateStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HubsUpdateStatusResponse)
	err := c.cc.Invoke(ctx, Hubs_UpdateStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HubsServer is the server API for Hubs service.
// All implementations must embed UnimplementedHubsServer
// for forward compatibility.
type HubsServer interface {
	// Retrieves the list of hubs.
	List(context.Context, *HubsListRequest) (*HubsListResponse, error)
	// Retrieves the details of one specific hub.
	Get(context.Context, *HubsGetRequest) (*HubsGetResponse, error)
	// Creates a new hub.
	Create(context.Context, *HubsCreateRequest) (*HubsCreateResponse, error)
	// Delete a hub.
	Delete(context.Context, *HubsDeleteRequest) (*HubsDeleteResponse, error)
	// Updates the status of a hub.
	UpdateStatus(context.Context, *HubsUpdateStatusRequest) (*HubsUpdateStatusResponse, error)
	mustEmbedUnimplementedHubsServer()
}

// UnimplementedHubsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHubsServer struct{}

func (UnimplementedHubsServer) List(context.Context, *HubsListRequest) (*HubsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedHubsServer) Get(context.Context, *HubsGetRequest) (*HubsGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHubsServer) Create(context.Context, *HubsCreateRequest) (*HubsCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedHubsServer) Delete(context.Context, *HubsDeleteRequest) (*HubsDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedHubsServer) UpdateStatus(context.Context, *HubsUpdateStatusRequest) (*HubsUpdateStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedHubsServer) mustEmbedUnimplementedHubsServer() {}
func (UnimplementedHubsServer) testEmbeddedByValue()              {}

// UnsafeHubsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubsServer will
// result in compilation errors.
type UnsafeHubsServer interface {
	mustEmbedUnimplementedHubsServer()
}

func RegisterHubsServer(s grpc.ServiceRegistrar, srv HubsServer) {
	// If the following call pancis, it indicates UnimplementedHubsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Hubs_ServiceDesc, srv)
}

func _Hubs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hubs_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubsServer).List(ctx, req.(*HubsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hubs_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hubs_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubsServer).Get(ctx, req.(*HubsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hubs_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubsCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hubs_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubsServer).Create(ctx, req.(*HubsCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hubs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubsDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hubs_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubsServer).Delete(ctx, req.(*HubsDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hubs_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubsUpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubsServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hubs_UpdateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubsServer).UpdateStatus(ctx, req.(*HubsUpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hubs_ServiceDesc is the grpc.ServiceDesc for Hubs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hubs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.v1.Hubs",
	HandlerType: (*HubsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Hubs_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Hubs_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Hubs_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Hubs_Delete_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _Hubs_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/v1/hubs_service.proto",
}
