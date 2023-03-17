// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user.proto

package user

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

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	ListNftCreators(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UsersResponse, error)
	ListAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UsersResponse, error)
	GetUserByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*User, error)
	GetUserByWalletAddress(ctx context.Context, in *WalletAddressRequest, opts ...grpc.CallOption) (*User, error)
	GetUserByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*User, error)
	ListUsersByWalletAddresses(ctx context.Context, in *WalletAddressesRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	ListWalletsByUserID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*WalletsResponse, error)
	GetOrCreateUserAsCreator(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*GetOrCreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*Wallet, error)
	AddWallet(ctx context.Context, in *AddWalletRequest, opts ...grpc.CallOption) (*Wallet, error)
	DeleteWallet(ctx context.Context, in *DeleteWalletRequest, opts ...grpc.CallOption) (*Empty, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) ListNftCreators(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/user.Users/ListNftCreators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ListAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/user.Users/ListAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.Users/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserByWalletAddress(ctx context.Context, in *WalletAddressRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.Users/GetUserByWalletAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.Users/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ListUsersByWalletAddresses(ctx context.Context, in *WalletAddressesRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/user.Users/ListUsersByWalletAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ListWalletsByUserID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*WalletsResponse, error) {
	out := new(WalletsResponse)
	err := c.cc.Invoke(ctx, "/user.Users/ListWalletsByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetOrCreateUserAsCreator(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*GetOrCreateUserResponse, error) {
	out := new(GetOrCreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.Users/GetOrCreateUserAsCreator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.Users/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateWallet(ctx context.Context, in *UpdateWalletRequest, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/user.Users/UpdateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) AddWallet(ctx context.Context, in *AddWalletRequest, opts ...grpc.CallOption) (*Wallet, error) {
	out := new(Wallet)
	err := c.cc.Invoke(ctx, "/user.Users/AddWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteWallet(ctx context.Context, in *DeleteWalletRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.Users/DeleteWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	ListNftCreators(context.Context, *Empty) (*UsersResponse, error)
	ListAll(context.Context, *Empty) (*UsersResponse, error)
	GetUserByEmail(context.Context, *EmailRequest) (*User, error)
	GetUserByWalletAddress(context.Context, *WalletAddressRequest) (*User, error)
	GetUserByID(context.Context, *IDRequest) (*User, error)
	ListUsersByWalletAddresses(context.Context, *WalletAddressesRequest) (*UsersResponse, error)
	ListWalletsByUserID(context.Context, *IDRequest) (*WalletsResponse, error)
	GetOrCreateUserAsCreator(context.Context, *EmailRequest) (*GetOrCreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	UpdateWallet(context.Context, *UpdateWalletRequest) (*Wallet, error)
	AddWallet(context.Context, *AddWalletRequest) (*Wallet, error)
	DeleteWallet(context.Context, *DeleteWalletRequest) (*Empty, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) ListNftCreators(context.Context, *Empty) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNftCreators not implemented")
}
func (UnimplementedUsersServer) ListAll(context.Context, *Empty) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAll not implemented")
}
func (UnimplementedUsersServer) GetUserByEmail(context.Context, *EmailRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUsersServer) GetUserByWalletAddress(context.Context, *WalletAddressRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByWalletAddress not implemented")
}
func (UnimplementedUsersServer) GetUserByID(context.Context, *IDRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedUsersServer) ListUsersByWalletAddresses(context.Context, *WalletAddressesRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersByWalletAddresses not implemented")
}
func (UnimplementedUsersServer) ListWalletsByUserID(context.Context, *IDRequest) (*WalletsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWalletsByUserID not implemented")
}
func (UnimplementedUsersServer) GetOrCreateUserAsCreator(context.Context, *EmailRequest) (*GetOrCreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateUserAsCreator not implemented")
}
func (UnimplementedUsersServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUsersServer) UpdateWallet(context.Context, *UpdateWalletRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWallet not implemented")
}
func (UnimplementedUsersServer) AddWallet(context.Context, *AddWalletRequest) (*Wallet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWallet not implemented")
}
func (UnimplementedUsersServer) DeleteWallet(context.Context, *DeleteWalletRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWallet not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_ListNftCreators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListNftCreators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/ListNftCreators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListNftCreators(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserByEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserByWalletAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserByWalletAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetUserByWalletAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserByWalletAddress(ctx, req.(*WalletAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserByID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ListUsersByWalletAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListUsersByWalletAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/ListUsersByWalletAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListUsersByWalletAddresses(ctx, req.(*WalletAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ListWalletsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListWalletsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/ListWalletsByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListWalletsByUserID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetOrCreateUserAsCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetOrCreateUserAsCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/GetOrCreateUserAsCreator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetOrCreateUserAsCreator(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/UpdateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateWallet(ctx, req.(*UpdateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_AddWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).AddWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/AddWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).AddWallet(ctx, req.(*AddWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Users/DeleteWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteWallet(ctx, req.(*DeleteWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNftCreators",
			Handler:    _Users_ListNftCreators_Handler,
		},
		{
			MethodName: "ListAll",
			Handler:    _Users_ListAll_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _Users_GetUserByEmail_Handler,
		},
		{
			MethodName: "GetUserByWalletAddress",
			Handler:    _Users_GetUserByWalletAddress_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _Users_GetUserByID_Handler,
		},
		{
			MethodName: "ListUsersByWalletAddresses",
			Handler:    _Users_ListUsersByWalletAddresses_Handler,
		},
		{
			MethodName: "ListWalletsByUserID",
			Handler:    _Users_ListWalletsByUserID_Handler,
		},
		{
			MethodName: "GetOrCreateUserAsCreator",
			Handler:    _Users_GetOrCreateUserAsCreator_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Users_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateWallet",
			Handler:    _Users_UpdateWallet_Handler,
		},
		{
			MethodName: "AddWallet",
			Handler:    _Users_AddWallet_Handler,
		},
		{
			MethodName: "DeleteWallet",
			Handler:    _Users_DeleteWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
