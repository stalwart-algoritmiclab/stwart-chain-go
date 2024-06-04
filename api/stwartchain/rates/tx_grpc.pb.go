// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: stwartchain/rates/tx.proto

package rates

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
	Msg_UpdateParams_FullMethodName    = "/stwartchain.rates.Msg/UpdateParams"
	Msg_CreateAddresses_FullMethodName = "/stwartchain.rates.Msg/CreateAddresses"
	Msg_UpdateAddresses_FullMethodName = "/stwartchain.rates.Msg/UpdateAddresses"
	Msg_DeleteAddresses_FullMethodName = "/stwartchain.rates.Msg/DeleteAddresses"
	Msg_CreateRates_FullMethodName     = "/stwartchain.rates.Msg/CreateRates"
	Msg_UpdateRates_FullMethodName     = "/stwartchain.rates.Msg/UpdateRates"
	Msg_DeleteRates_FullMethodName     = "/stwartchain.rates.Msg/DeleteRates"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateAddresses(ctx context.Context, in *MsgCreateAddresses, opts ...grpc.CallOption) (*MsgCreateAddressesResponse, error)
	UpdateAddresses(ctx context.Context, in *MsgUpdateAddresses, opts ...grpc.CallOption) (*MsgUpdateAddressesResponse, error)
	DeleteAddresses(ctx context.Context, in *MsgDeleteAddresses, opts ...grpc.CallOption) (*MsgDeleteAddressesResponse, error)
	CreateRates(ctx context.Context, in *MsgCreateRates, opts ...grpc.CallOption) (*MsgCreateRatesResponse, error)
	UpdateRates(ctx context.Context, in *MsgUpdateRates, opts ...grpc.CallOption) (*MsgUpdateRatesResponse, error)
	DeleteRates(ctx context.Context, in *MsgDeleteRates, opts ...grpc.CallOption) (*MsgDeleteRatesResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateAddresses(ctx context.Context, in *MsgCreateAddresses, opts ...grpc.CallOption) (*MsgCreateAddressesResponse, error) {
	out := new(MsgCreateAddressesResponse)
	err := c.cc.Invoke(ctx, Msg_CreateAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateAddresses(ctx context.Context, in *MsgUpdateAddresses, opts ...grpc.CallOption) (*MsgUpdateAddressesResponse, error) {
	out := new(MsgUpdateAddressesResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteAddresses(ctx context.Context, in *MsgDeleteAddresses, opts ...grpc.CallOption) (*MsgDeleteAddressesResponse, error) {
	out := new(MsgDeleteAddressesResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateRates(ctx context.Context, in *MsgCreateRates, opts ...grpc.CallOption) (*MsgCreateRatesResponse, error) {
	out := new(MsgCreateRatesResponse)
	err := c.cc.Invoke(ctx, Msg_CreateRates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRates(ctx context.Context, in *MsgUpdateRates, opts ...grpc.CallOption) (*MsgUpdateRatesResponse, error) {
	out := new(MsgUpdateRatesResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateRates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteRates(ctx context.Context, in *MsgDeleteRates, opts ...grpc.CallOption) (*MsgDeleteRatesResponse, error) {
	out := new(MsgDeleteRatesResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteRates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateAddresses(context.Context, *MsgCreateAddresses) (*MsgCreateAddressesResponse, error)
	UpdateAddresses(context.Context, *MsgUpdateAddresses) (*MsgUpdateAddressesResponse, error)
	DeleteAddresses(context.Context, *MsgDeleteAddresses) (*MsgDeleteAddressesResponse, error)
	CreateRates(context.Context, *MsgCreateRates) (*MsgCreateRatesResponse, error)
	UpdateRates(context.Context, *MsgUpdateRates) (*MsgUpdateRatesResponse, error)
	DeleteRates(context.Context, *MsgDeleteRates) (*MsgDeleteRatesResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateAddresses(context.Context, *MsgCreateAddresses) (*MsgCreateAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddresses not implemented")
}
func (UnimplementedMsgServer) UpdateAddresses(context.Context, *MsgUpdateAddresses) (*MsgUpdateAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddresses not implemented")
}
func (UnimplementedMsgServer) DeleteAddresses(context.Context, *MsgDeleteAddresses) (*MsgDeleteAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddresses not implemented")
}
func (UnimplementedMsgServer) CreateRates(context.Context, *MsgCreateRates) (*MsgCreateRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRates not implemented")
}
func (UnimplementedMsgServer) UpdateRates(context.Context, *MsgUpdateRates) (*MsgUpdateRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRates not implemented")
}
func (UnimplementedMsgServer) DeleteRates(context.Context, *MsgDeleteRates) (*MsgDeleteRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRates not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAddresses)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAddresses(ctx, req.(*MsgCreateAddresses))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateAddresses)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateAddresses(ctx, req.(*MsgUpdateAddresses))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteAddresses)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteAddresses(ctx, req.(*MsgDeleteAddresses))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateRates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRates(ctx, req.(*MsgCreateRates))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateRates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRates(ctx, req.(*MsgUpdateRates))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteRates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteRates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteRates(ctx, req.(*MsgDeleteRates))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stwartchain.rates.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateAddresses",
			Handler:    _Msg_CreateAddresses_Handler,
		},
		{
			MethodName: "UpdateAddresses",
			Handler:    _Msg_UpdateAddresses_Handler,
		},
		{
			MethodName: "DeleteAddresses",
			Handler:    _Msg_DeleteAddresses_Handler,
		},
		{
			MethodName: "CreateRates",
			Handler:    _Msg_CreateRates_Handler,
		},
		{
			MethodName: "UpdateRates",
			Handler:    _Msg_UpdateRates_Handler,
		},
		{
			MethodName: "DeleteRates",
			Handler:    _Msg_DeleteRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stwartchain/rates/tx.proto",
}
