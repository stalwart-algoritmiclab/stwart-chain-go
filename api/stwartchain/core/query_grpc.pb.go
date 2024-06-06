// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: stwartchain/core/query.proto

package core

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
	Query_Params_FullMethodName           = "/stwartchain.core.Query/Params"
	Query_Stats_FullMethodName            = "/stwartchain.core.Query/Stats"
	Query_StatsAll_FullMethodName         = "/stwartchain.core.Query/StatsAll"
	Query_ModulesAddresses_FullMethodName = "/stwartchain.core.Query/ModulesAddresses"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Stats items.
	Stats(ctx context.Context, in *QueryGetStatsRequest, opts ...grpc.CallOption) (*QueryGetStatsResponse, error)
	StatsAll(ctx context.Context, in *QueryAllStatsRequest, opts ...grpc.CallOption) (*QueryAllStatsResponse, error)
	// Queries a list of ModulesAddresses items.
	ModulesAddresses(ctx context.Context, in *QueryModulesAddressesRequest, opts ...grpc.CallOption) (*QueryModulesAddressesResponse, error)
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

func (c *queryClient) Stats(ctx context.Context, in *QueryGetStatsRequest, opts ...grpc.CallOption) (*QueryGetStatsResponse, error) {
	out := new(QueryGetStatsResponse)
	err := c.cc.Invoke(ctx, Query_Stats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StatsAll(ctx context.Context, in *QueryAllStatsRequest, opts ...grpc.CallOption) (*QueryAllStatsResponse, error) {
	out := new(QueryAllStatsResponse)
	err := c.cc.Invoke(ctx, Query_StatsAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ModulesAddresses(ctx context.Context, in *QueryModulesAddressesRequest, opts ...grpc.CallOption) (*QueryModulesAddressesResponse, error) {
	out := new(QueryModulesAddressesResponse)
	err := c.cc.Invoke(ctx, Query_ModulesAddresses_FullMethodName, in, out, opts...)
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
	// Queries a list of Stats items.
	Stats(context.Context, *QueryGetStatsRequest) (*QueryGetStatsResponse, error)
	StatsAll(context.Context, *QueryAllStatsRequest) (*QueryAllStatsResponse, error)
	// Queries a list of ModulesAddresses items.
	ModulesAddresses(context.Context, *QueryModulesAddressesRequest) (*QueryModulesAddressesResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Stats(context.Context, *QueryGetStatsRequest) (*QueryGetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedQueryServer) StatsAll(context.Context, *QueryAllStatsRequest) (*QueryAllStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatsAll not implemented")
}
func (UnimplementedQueryServer) ModulesAddresses(context.Context, *QueryModulesAddressesRequest) (*QueryModulesAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModulesAddresses not implemented")
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

func _Query_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Stats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Stats(ctx, req.(*QueryGetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StatsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StatsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_StatsAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StatsAll(ctx, req.(*QueryAllStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ModulesAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryModulesAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ModulesAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ModulesAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ModulesAddresses(ctx, req.(*QueryModulesAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stwartchain.core.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _Query_Stats_Handler,
		},
		{
			MethodName: "StatsAll",
			Handler:    _Query_StatsAll_Handler,
		},
		{
			MethodName: "ModulesAddresses",
			Handler:    _Query_ModulesAddresses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stwartchain/core/query.proto",
}
