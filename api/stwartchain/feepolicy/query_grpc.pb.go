// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: stwartchain/feepolicy/query.proto

package feepolicy

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
	Query_Params_FullMethodName       = "/stwartchain.feepolicy.Query/Params"
	Query_Addresses_FullMethodName    = "/stwartchain.feepolicy.Query/Addresses"
	Query_AddressesAll_FullMethodName = "/stwartchain.feepolicy.Query/AddressesAll"
	Query_Tariffs_FullMethodName      = "/stwartchain.feepolicy.Query/Tariffs"
	Query_TariffsAll_FullMethodName   = "/stwartchain.feepolicy.Query/TariffsAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Addresses items.
	Addresses(ctx context.Context, in *QueryGetAddressesRequest, opts ...grpc.CallOption) (*QueryGetAddressesResponse, error)
	AddressesAll(ctx context.Context, in *QueryAllAddressesRequest, opts ...grpc.CallOption) (*QueryAllAddressesResponse, error)
	// Queries a list of Tariffs items.
	Tariffs(ctx context.Context, in *QueryGetTariffsRequest, opts ...grpc.CallOption) (*QueryGetTariffsResponse, error)
	TariffsAll(ctx context.Context, in *QueryAllTariffsRequest, opts ...grpc.CallOption) (*QueryAllTariffsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Addresses(ctx context.Context, in *QueryGetAddressesRequest, opts ...grpc.CallOption) (*QueryGetAddressesResponse, error) {
	out := new(QueryGetAddressesResponse)
	err := c.cc.Invoke(ctx, Query_Addresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AddressesAll(ctx context.Context, in *QueryAllAddressesRequest, opts ...grpc.CallOption) (*QueryAllAddressesResponse, error) {
	out := new(QueryAllAddressesResponse)
	err := c.cc.Invoke(ctx, Query_AddressesAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Tariffs(ctx context.Context, in *QueryGetTariffsRequest, opts ...grpc.CallOption) (*QueryGetTariffsResponse, error) {
	out := new(QueryGetTariffsResponse)
	err := c.cc.Invoke(ctx, Query_Tariffs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TariffsAll(ctx context.Context, in *QueryAllTariffsRequest, opts ...grpc.CallOption) (*QueryAllTariffsResponse, error) {
	out := new(QueryAllTariffsResponse)
	err := c.cc.Invoke(ctx, Query_TariffsAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Addresses items.
	Addresses(context.Context, *QueryGetAddressesRequest) (*QueryGetAddressesResponse, error)
	AddressesAll(context.Context, *QueryAllAddressesRequest) (*QueryAllAddressesResponse, error)
	// Queries a list of Tariffs items.
	Tariffs(context.Context, *QueryGetTariffsRequest) (*QueryGetTariffsResponse, error)
	TariffsAll(context.Context, *QueryAllTariffsRequest) (*QueryAllTariffsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Addresses(context.Context, *QueryGetAddressesRequest) (*QueryGetAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Addresses not implemented")
}
func (UnimplementedQueryServer) AddressesAll(context.Context, *QueryAllAddressesRequest) (*QueryAllAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressesAll not implemented")
}
func (UnimplementedQueryServer) Tariffs(context.Context, *QueryGetTariffsRequest) (*QueryGetTariffsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tariffs not implemented")
}
func (UnimplementedQueryServer) TariffsAll(context.Context, *QueryAllTariffsRequest) (*QueryAllTariffsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TariffsAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Addresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Addresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Addresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Addresses(ctx, req.(*QueryGetAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AddressesAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AddressesAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AddressesAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AddressesAll(ctx, req.(*QueryAllAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Tariffs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTariffsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Tariffs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Tariffs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Tariffs(ctx, req.(*QueryGetTariffsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TariffsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTariffsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TariffsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TariffsAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TariffsAll(ctx, req.(*QueryAllTariffsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stwartchain.feepolicy.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Addresses",
			Handler:    _Query_Addresses_Handler,
		},
		{
			MethodName: "AddressesAll",
			Handler:    _Query_AddressesAll_Handler,
		},
		{
			MethodName: "Tariffs",
			Handler:    _Query_Tariffs_Handler,
		},
		{
			MethodName: "TariffsAll",
			Handler:    _Query_TariffsAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stwartchain/feepolicy/query.proto",
}
