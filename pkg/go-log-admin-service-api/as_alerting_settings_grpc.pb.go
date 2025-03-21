// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: as_alerting_settings.proto

package goLog_admin_service

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

const (
	AlertingSettingsService_CreateAlertingSettings_FullMethodName = "/admin_service.AlertingSettingsService/CreateAlertingSettings"
	AlertingSettingsService_GetAlertingSettings_FullMethodName    = "/admin_service.AlertingSettingsService/GetAlertingSettings"
	AlertingSettingsService_ListAlertingSettings_FullMethodName   = "/admin_service.AlertingSettingsService/ListAlertingSettings"
	AlertingSettingsService_DeleteAlertingSettings_FullMethodName = "/admin_service.AlertingSettingsService/DeleteAlertingSettings"
	AlertingSettingsService_PatchAlertingSettings_FullMethodName  = "/admin_service.AlertingSettingsService/PatchAlertingSettings"
)

// AlertingSettingsServiceClient is the client API for AlertingSettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertingSettingsServiceClient interface {
	CreateAlertingSettings(ctx context.Context, in *CreateAlertingSettingsRequest, opts ...grpc.CallOption) (*CreateAlertingSettingsResponse, error)
	GetAlertingSettings(ctx context.Context, in *GetAlertingSettingsRequest, opts ...grpc.CallOption) (*AlertingSetting, error)
	ListAlertingSettings(ctx context.Context, in *ListAlertingSettingsRequest, opts ...grpc.CallOption) (*ListAlertingSettingsResponse, error)
	DeleteAlertingSettings(ctx context.Context, in *DeleteAlertingSettingsRequest, opts ...grpc.CallOption) (*DeleteAlertingSettingsResponse, error)
	PatchAlertingSettings(ctx context.Context, in *PatchAlertingSettingsRequest, opts ...grpc.CallOption) (*AlertingSetting, error)
}

type alertingSettingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertingSettingsServiceClient(cc grpc.ClientConnInterface) AlertingSettingsServiceClient {
	return &alertingSettingsServiceClient{cc}
}

func (c *alertingSettingsServiceClient) CreateAlertingSettings(ctx context.Context, in *CreateAlertingSettingsRequest, opts ...grpc.CallOption) (*CreateAlertingSettingsResponse, error) {
	out := new(CreateAlertingSettingsResponse)
	err := c.cc.Invoke(ctx, AlertingSettingsService_CreateAlertingSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingSettingsServiceClient) GetAlertingSettings(ctx context.Context, in *GetAlertingSettingsRequest, opts ...grpc.CallOption) (*AlertingSetting, error) {
	out := new(AlertingSetting)
	err := c.cc.Invoke(ctx, AlertingSettingsService_GetAlertingSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingSettingsServiceClient) ListAlertingSettings(ctx context.Context, in *ListAlertingSettingsRequest, opts ...grpc.CallOption) (*ListAlertingSettingsResponse, error) {
	out := new(ListAlertingSettingsResponse)
	err := c.cc.Invoke(ctx, AlertingSettingsService_ListAlertingSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingSettingsServiceClient) DeleteAlertingSettings(ctx context.Context, in *DeleteAlertingSettingsRequest, opts ...grpc.CallOption) (*DeleteAlertingSettingsResponse, error) {
	out := new(DeleteAlertingSettingsResponse)
	err := c.cc.Invoke(ctx, AlertingSettingsService_DeleteAlertingSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertingSettingsServiceClient) PatchAlertingSettings(ctx context.Context, in *PatchAlertingSettingsRequest, opts ...grpc.CallOption) (*AlertingSetting, error) {
	out := new(AlertingSetting)
	err := c.cc.Invoke(ctx, AlertingSettingsService_PatchAlertingSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertingSettingsServiceServer is the server API for AlertingSettingsService service.
// All implementations must embed UnimplementedAlertingSettingsServiceServer
// for forward compatibility
type AlertingSettingsServiceServer interface {
	CreateAlertingSettings(context.Context, *CreateAlertingSettingsRequest) (*CreateAlertingSettingsResponse, error)
	GetAlertingSettings(context.Context, *GetAlertingSettingsRequest) (*AlertingSetting, error)
	ListAlertingSettings(context.Context, *ListAlertingSettingsRequest) (*ListAlertingSettingsResponse, error)
	DeleteAlertingSettings(context.Context, *DeleteAlertingSettingsRequest) (*DeleteAlertingSettingsResponse, error)
	PatchAlertingSettings(context.Context, *PatchAlertingSettingsRequest) (*AlertingSetting, error)
	mustEmbedUnimplementedAlertingSettingsServiceServer()
}

// UnimplementedAlertingSettingsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAlertingSettingsServiceServer struct {
}

func (UnimplementedAlertingSettingsServiceServer) CreateAlertingSettings(context.Context, *CreateAlertingSettingsRequest) (*CreateAlertingSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertingSettings not implemented")
}
func (UnimplementedAlertingSettingsServiceServer) GetAlertingSettings(context.Context, *GetAlertingSettingsRequest) (*AlertingSetting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertingSettings not implemented")
}
func (UnimplementedAlertingSettingsServiceServer) ListAlertingSettings(context.Context, *ListAlertingSettingsRequest) (*ListAlertingSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertingSettings not implemented")
}
func (UnimplementedAlertingSettingsServiceServer) DeleteAlertingSettings(context.Context, *DeleteAlertingSettingsRequest) (*DeleteAlertingSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertingSettings not implemented")
}
func (UnimplementedAlertingSettingsServiceServer) PatchAlertingSettings(context.Context, *PatchAlertingSettingsRequest) (*AlertingSetting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchAlertingSettings not implemented")
}
func (UnimplementedAlertingSettingsServiceServer) mustEmbedUnimplementedAlertingSettingsServiceServer() {
}

// UnsafeAlertingSettingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertingSettingsServiceServer will
// result in compilation errors.
type UnsafeAlertingSettingsServiceServer interface {
	mustEmbedUnimplementedAlertingSettingsServiceServer()
}

func RegisterAlertingSettingsServiceServer(s grpc.ServiceRegistrar, srv AlertingSettingsServiceServer) {
	s.RegisterService(&AlertingSettingsService_ServiceDesc, srv)
}

func _AlertingSettingsService_CreateAlertingSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertingSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingSettingsServiceServer).CreateAlertingSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertingSettingsService_CreateAlertingSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingSettingsServiceServer).CreateAlertingSettings(ctx, req.(*CreateAlertingSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertingSettingsService_GetAlertingSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertingSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingSettingsServiceServer).GetAlertingSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertingSettingsService_GetAlertingSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingSettingsServiceServer).GetAlertingSettings(ctx, req.(*GetAlertingSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertingSettingsService_ListAlertingSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertingSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingSettingsServiceServer).ListAlertingSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertingSettingsService_ListAlertingSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingSettingsServiceServer).ListAlertingSettings(ctx, req.(*ListAlertingSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertingSettingsService_DeleteAlertingSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertingSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingSettingsServiceServer).DeleteAlertingSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertingSettingsService_DeleteAlertingSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingSettingsServiceServer).DeleteAlertingSettings(ctx, req.(*DeleteAlertingSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertingSettingsService_PatchAlertingSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchAlertingSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertingSettingsServiceServer).PatchAlertingSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertingSettingsService_PatchAlertingSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertingSettingsServiceServer).PatchAlertingSettings(ctx, req.(*PatchAlertingSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertingSettingsService_ServiceDesc is the grpc.ServiceDesc for AlertingSettingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertingSettingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_service.AlertingSettingsService",
	HandlerType: (*AlertingSettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlertingSettings",
			Handler:    _AlertingSettingsService_CreateAlertingSettings_Handler,
		},
		{
			MethodName: "GetAlertingSettings",
			Handler:    _AlertingSettingsService_GetAlertingSettings_Handler,
		},
		{
			MethodName: "ListAlertingSettings",
			Handler:    _AlertingSettingsService_ListAlertingSettings_Handler,
		},
		{
			MethodName: "DeleteAlertingSettings",
			Handler:    _AlertingSettingsService_DeleteAlertingSettings_Handler,
		},
		{
			MethodName: "PatchAlertingSettings",
			Handler:    _AlertingSettingsService_PatchAlertingSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as_alerting_settings.proto",
}
