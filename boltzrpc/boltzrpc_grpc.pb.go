// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package boltzrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BoltzClient is the client API for Boltz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoltzClient interface {
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
	GetServiceInfo(ctx context.Context, in *GetServiceInfoRequest, opts ...grpc.CallOption) (*GetServiceInfoResponse, error)
	ListSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error)
	GetSwapInfo(ctx context.Context, in *GetSwapInfoRequest, opts ...grpc.CallOption) (*GetSwapInfoResponse, error)
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
	CreateSwap(ctx context.Context, in *CreateSwapRequest, opts ...grpc.CallOption) (*CreateSwapResponse, error)
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateSwapResponse, error)
	CreateReverseSwap(ctx context.Context, in *CreateReverseSwapRequest, opts ...grpc.CallOption) (*CreateReverseSwapResponse, error)
}

type boltzClient struct {
	cc grpc.ClientConnInterface
}

func NewBoltzClient(cc grpc.ClientConnInterface) BoltzClient {
	return &boltzClient{cc}
}

func (c *boltzClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/boltzrpc.Boltz/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boltzClient) GetServiceInfo(ctx context.Context, in *GetServiceInfoRequest, opts ...grpc.CallOption) (*GetServiceInfoResponse, error) {
	out := new(GetServiceInfoResponse)
	err := c.cc.Invoke(ctx, "/boltzrpc.Boltz/GetServiceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boltzClient) ListSwaps(ctx context.Context, in *ListSwapsRequest, opts ...grpc.CallOption) (*ListSwapsResponse, error) {
	out := new(ListSwapsResponse)
	err := c.cc.Invoke(ctx, "/boltzrpc.Boltz/ListSwaps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boltzClient) GetSwapInfo(ctx context.Context, in *GetSwapInfoRequest, opts ...grpc.CallOption) (*GetSwapInfoResponse, error) {
	out := new(GetSwapInfoResponse)
	err := c.cc.Invoke(ctx, "/boltzrpc.Boltz/GetSwapInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boltzClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, "/boltzrpc.Boltz/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boltzClient) CreateSwap(ctx context.Context, in *CreateSwapRequest, opts ...grpc.CallOption) (*CreateSwapResponse, error) {
	out := new(CreateSwapResponse)
	err := c.cc.Invoke(ctx, "/boltzrpc.Boltz/CreateSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boltzClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateSwapResponse, error) {
	out := new(CreateSwapResponse)
	err := c.cc.Invoke(ctx, "/boltzrpc.Boltz/CreateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boltzClient) CreateReverseSwap(ctx context.Context, in *CreateReverseSwapRequest, opts ...grpc.CallOption) (*CreateReverseSwapResponse, error) {
	out := new(CreateReverseSwapResponse)
	err := c.cc.Invoke(ctx, "/boltzrpc.Boltz/CreateReverseSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoltzServer is the server API for Boltz service.
// All implementations must embed UnimplementedBoltzServer
// for forward compatibility
type BoltzServer interface {
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	GetServiceInfo(context.Context, *GetServiceInfoRequest) (*GetServiceInfoResponse, error)
	ListSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error)
	GetSwapInfo(context.Context, *GetSwapInfoRequest) (*GetSwapInfoResponse, error)
	Deposit(context.Context, *DepositRequest) (*DepositResponse, error)
	CreateSwap(context.Context, *CreateSwapRequest) (*CreateSwapResponse, error)
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateSwapResponse, error)
	CreateReverseSwap(context.Context, *CreateReverseSwapRequest) (*CreateReverseSwapResponse, error)
	mustEmbedUnimplementedBoltzServer()
}

// UnimplementedBoltzServer must be embedded to have forward compatible implementations.
type UnimplementedBoltzServer struct {
}

func (UnimplementedBoltzServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedBoltzServer) GetServiceInfo(context.Context, *GetServiceInfoRequest) (*GetServiceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceInfo not implemented")
}
func (UnimplementedBoltzServer) ListSwaps(context.Context, *ListSwapsRequest) (*ListSwapsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSwaps not implemented")
}
func (UnimplementedBoltzServer) GetSwapInfo(context.Context, *GetSwapInfoRequest) (*GetSwapInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSwapInfo not implemented")
}
func (UnimplementedBoltzServer) Deposit(context.Context, *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedBoltzServer) CreateSwap(context.Context, *CreateSwapRequest) (*CreateSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSwap not implemented")
}
func (UnimplementedBoltzServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedBoltzServer) CreateReverseSwap(context.Context, *CreateReverseSwapRequest) (*CreateReverseSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReverseSwap not implemented")
}
func (UnimplementedBoltzServer) mustEmbedUnimplementedBoltzServer() {}

// UnsafeBoltzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoltzServer will
// result in compilation errors.
type UnsafeBoltzServer interface {
	mustEmbedUnimplementedBoltzServer()
}

func RegisterBoltzServer(s grpc.ServiceRegistrar, srv BoltzServer) {
	s.RegisterService(&_Boltz_serviceDesc, srv)
}

func _Boltz_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoltzServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boltzrpc.Boltz/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoltzServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boltz_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoltzServer).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boltzrpc.Boltz/GetServiceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoltzServer).GetServiceInfo(ctx, req.(*GetServiceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boltz_ListSwaps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSwapsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoltzServer).ListSwaps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boltzrpc.Boltz/ListSwaps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoltzServer).ListSwaps(ctx, req.(*ListSwapsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boltz_GetSwapInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSwapInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoltzServer).GetSwapInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boltzrpc.Boltz/GetSwapInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoltzServer).GetSwapInfo(ctx, req.(*GetSwapInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boltz_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoltzServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boltzrpc.Boltz/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoltzServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boltz_CreateSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoltzServer).CreateSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boltzrpc.Boltz/CreateSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoltzServer).CreateSwap(ctx, req.(*CreateSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boltz_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoltzServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boltzrpc.Boltz/CreateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoltzServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Boltz_CreateReverseSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReverseSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoltzServer).CreateReverseSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/boltzrpc.Boltz/CreateReverseSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoltzServer).CreateReverseSwap(ctx, req.(*CreateReverseSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Boltz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "boltzrpc.Boltz",
	HandlerType: (*BoltzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Boltz_GetInfo_Handler,
		},
		{
			MethodName: "GetServiceInfo",
			Handler:    _Boltz_GetServiceInfo_Handler,
		},
		{
			MethodName: "ListSwaps",
			Handler:    _Boltz_ListSwaps_Handler,
		},
		{
			MethodName: "GetSwapInfo",
			Handler:    _Boltz_GetSwapInfo_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Boltz_Deposit_Handler,
		},
		{
			MethodName: "CreateSwap",
			Handler:    _Boltz_CreateSwap_Handler,
		},
		{
			MethodName: "CreateChannel",
			Handler:    _Boltz_CreateChannel_Handler,
		},
		{
			MethodName: "CreateReverseSwap",
			Handler:    _Boltz_CreateReverseSwap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "boltzrpc.proto",
}
