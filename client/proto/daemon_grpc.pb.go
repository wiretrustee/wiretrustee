// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// DaemonServiceClient is the client API for DaemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonServiceClient interface {
	// Login uses setup key to prepare configuration for the daemon.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// WaitSSOLogin uses the userCode to validate the TokenInfo and
	// waits for the user to continue with the login on a browser
	WaitSSOLogin(ctx context.Context, in *WaitSSOLoginRequest, opts ...grpc.CallOption) (*WaitSSOLoginResponse, error)
	// Up starts engine work in the daemon.
	Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpResponse, error)
	// Status of the service.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// Down engine work in the daemon.
	Down(ctx context.Context, in *DownRequest, opts ...grpc.CallOption) (*DownResponse, error)
	// GetConfig of the daemon.
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
	// List available networks
	ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error)
	// Select specific routes
	SelectNetworks(ctx context.Context, in *SelectNetworksRequest, opts ...grpc.CallOption) (*SelectNetworksResponse, error)
	// Deselect specific routes
	DeselectNetworks(ctx context.Context, in *SelectNetworksRequest, opts ...grpc.CallOption) (*SelectNetworksResponse, error)
	// DebugBundle creates a debug bundle
	DebugBundle(ctx context.Context, in *DebugBundleRequest, opts ...grpc.CallOption) (*DebugBundleResponse, error)
	// GetLogLevel gets the log level of the daemon
	GetLogLevel(ctx context.Context, in *GetLogLevelRequest, opts ...grpc.CallOption) (*GetLogLevelResponse, error)
	// SetLogLevel sets the log level of the daemon
	SetLogLevel(ctx context.Context, in *SetLogLevelRequest, opts ...grpc.CallOption) (*SetLogLevelResponse, error)
	// List all states
	ListStates(ctx context.Context, in *ListStatesRequest, opts ...grpc.CallOption) (*ListStatesResponse, error)
	// Clean specific state or all states
	CleanState(ctx context.Context, in *CleanStateRequest, opts ...grpc.CallOption) (*CleanStateResponse, error)
	// Delete specific state or all states
	DeleteState(ctx context.Context, in *DeleteStateRequest, opts ...grpc.CallOption) (*DeleteStateResponse, error)
	// SetNetworkMapPersistence enables or disables network map persistence
	SetNetworkMapPersistence(ctx context.Context, in *SetNetworkMapPersistenceRequest, opts ...grpc.CallOption) (*SetNetworkMapPersistenceResponse, error)
	TracePacket(ctx context.Context, in *TracePacketRequest, opts ...grpc.CallOption) (*TracePacketResponse, error)
}

type daemonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonServiceClient(cc grpc.ClientConnInterface) DaemonServiceClient {
	return &daemonServiceClient{cc}
}

func (c *daemonServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) WaitSSOLogin(ctx context.Context, in *WaitSSOLoginRequest, opts ...grpc.CallOption) (*WaitSSOLoginResponse, error) {
	out := new(WaitSSOLoginResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/WaitSSOLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpResponse, error) {
	out := new(UpResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/Up", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) Down(ctx context.Context, in *DownRequest, opts ...grpc.CallOption) (*DownResponse, error) {
	out := new(DownResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/Down", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error) {
	out := new(ListNetworksResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/ListNetworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) SelectNetworks(ctx context.Context, in *SelectNetworksRequest, opts ...grpc.CallOption) (*SelectNetworksResponse, error) {
	out := new(SelectNetworksResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/SelectNetworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) DeselectNetworks(ctx context.Context, in *SelectNetworksRequest, opts ...grpc.CallOption) (*SelectNetworksResponse, error) {
	out := new(SelectNetworksResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/DeselectNetworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) DebugBundle(ctx context.Context, in *DebugBundleRequest, opts ...grpc.CallOption) (*DebugBundleResponse, error) {
	out := new(DebugBundleResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/DebugBundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetLogLevel(ctx context.Context, in *GetLogLevelRequest, opts ...grpc.CallOption) (*GetLogLevelResponse, error) {
	out := new(GetLogLevelResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/GetLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) SetLogLevel(ctx context.Context, in *SetLogLevelRequest, opts ...grpc.CallOption) (*SetLogLevelResponse, error) {
	out := new(SetLogLevelResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/SetLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) ListStates(ctx context.Context, in *ListStatesRequest, opts ...grpc.CallOption) (*ListStatesResponse, error) {
	out := new(ListStatesResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/ListStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) CleanState(ctx context.Context, in *CleanStateRequest, opts ...grpc.CallOption) (*CleanStateResponse, error) {
	out := new(CleanStateResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/CleanState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) DeleteState(ctx context.Context, in *DeleteStateRequest, opts ...grpc.CallOption) (*DeleteStateResponse, error) {
	out := new(DeleteStateResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/DeleteState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) SetNetworkMapPersistence(ctx context.Context, in *SetNetworkMapPersistenceRequest, opts ...grpc.CallOption) (*SetNetworkMapPersistenceResponse, error) {
	out := new(SetNetworkMapPersistenceResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/SetNetworkMapPersistence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) TracePacket(ctx context.Context, in *TracePacketRequest, opts ...grpc.CallOption) (*TracePacketResponse, error) {
	out := new(TracePacketResponse)
	err := c.cc.Invoke(ctx, "/daemon.DaemonService/TracePacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServiceServer is the server API for DaemonService service.
// All implementations must embed UnimplementedDaemonServiceServer
// for forward compatibility
type DaemonServiceServer interface {
	// Login uses setup key to prepare configuration for the daemon.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// WaitSSOLogin uses the userCode to validate the TokenInfo and
	// waits for the user to continue with the login on a browser
	WaitSSOLogin(context.Context, *WaitSSOLoginRequest) (*WaitSSOLoginResponse, error)
	// Up starts engine work in the daemon.
	Up(context.Context, *UpRequest) (*UpResponse, error)
	// Status of the service.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// Down engine work in the daemon.
	Down(context.Context, *DownRequest) (*DownResponse, error)
	// GetConfig of the daemon.
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	// List available networks
	ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error)
	// Select specific routes
	SelectNetworks(context.Context, *SelectNetworksRequest) (*SelectNetworksResponse, error)
	// Deselect specific routes
	DeselectNetworks(context.Context, *SelectNetworksRequest) (*SelectNetworksResponse, error)
	// DebugBundle creates a debug bundle
	DebugBundle(context.Context, *DebugBundleRequest) (*DebugBundleResponse, error)
	// GetLogLevel gets the log level of the daemon
	GetLogLevel(context.Context, *GetLogLevelRequest) (*GetLogLevelResponse, error)
	// SetLogLevel sets the log level of the daemon
	SetLogLevel(context.Context, *SetLogLevelRequest) (*SetLogLevelResponse, error)
	// List all states
	ListStates(context.Context, *ListStatesRequest) (*ListStatesResponse, error)
	// Clean specific state or all states
	CleanState(context.Context, *CleanStateRequest) (*CleanStateResponse, error)
	// Delete specific state or all states
	DeleteState(context.Context, *DeleteStateRequest) (*DeleteStateResponse, error)
	// SetNetworkMapPersistence enables or disables network map persistence
	SetNetworkMapPersistence(context.Context, *SetNetworkMapPersistenceRequest) (*SetNetworkMapPersistenceResponse, error)
	TracePacket(context.Context, *TracePacketRequest) (*TracePacketResponse, error)
	mustEmbedUnimplementedDaemonServiceServer()
}

// UnimplementedDaemonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDaemonServiceServer struct {
}

func (UnimplementedDaemonServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedDaemonServiceServer) WaitSSOLogin(context.Context, *WaitSSOLoginRequest) (*WaitSSOLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitSSOLogin not implemented")
}
func (UnimplementedDaemonServiceServer) Up(context.Context, *UpRequest) (*UpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Up not implemented")
}
func (UnimplementedDaemonServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedDaemonServiceServer) Down(context.Context, *DownRequest) (*DownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Down not implemented")
}
func (UnimplementedDaemonServiceServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedDaemonServiceServer) ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNetworks not implemented")
}
func (UnimplementedDaemonServiceServer) SelectNetworks(context.Context, *SelectNetworksRequest) (*SelectNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectNetworks not implemented")
}
func (UnimplementedDaemonServiceServer) DeselectNetworks(context.Context, *SelectNetworksRequest) (*SelectNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeselectNetworks not implemented")
}
func (UnimplementedDaemonServiceServer) DebugBundle(context.Context, *DebugBundleRequest) (*DebugBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DebugBundle not implemented")
}
func (UnimplementedDaemonServiceServer) GetLogLevel(context.Context, *GetLogLevelRequest) (*GetLogLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogLevel not implemented")
}
func (UnimplementedDaemonServiceServer) SetLogLevel(context.Context, *SetLogLevelRequest) (*SetLogLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogLevel not implemented")
}
func (UnimplementedDaemonServiceServer) ListStates(context.Context, *ListStatesRequest) (*ListStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStates not implemented")
}
func (UnimplementedDaemonServiceServer) CleanState(context.Context, *CleanStateRequest) (*CleanStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanState not implemented")
}
func (UnimplementedDaemonServiceServer) DeleteState(context.Context, *DeleteStateRequest) (*DeleteStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteState not implemented")
}
func (UnimplementedDaemonServiceServer) SetNetworkMapPersistence(context.Context, *SetNetworkMapPersistenceRequest) (*SetNetworkMapPersistenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNetworkMapPersistence not implemented")
}
func (UnimplementedDaemonServiceServer) TracePacket(context.Context, *TracePacketRequest) (*TracePacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TracePacket not implemented")
}
func (UnimplementedDaemonServiceServer) mustEmbedUnimplementedDaemonServiceServer() {}

// UnsafeDaemonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServiceServer will
// result in compilation errors.
type UnsafeDaemonServiceServer interface {
	mustEmbedUnimplementedDaemonServiceServer()
}

func RegisterDaemonServiceServer(s grpc.ServiceRegistrar, srv DaemonServiceServer) {
	s.RegisterService(&DaemonService_ServiceDesc, srv)
}

func _DaemonService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_WaitSSOLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitSSOLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).WaitSSOLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/WaitSSOLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).WaitSSOLogin(ctx, req.(*WaitSSOLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/Up",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).Up(ctx, req.(*UpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_Down_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).Down(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/Down",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).Down(ctx, req.(*DownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_ListNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).ListNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/ListNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).ListNetworks(ctx, req.(*ListNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_SelectNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).SelectNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/SelectNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).SelectNetworks(ctx, req.(*SelectNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_DeselectNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).DeselectNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/DeselectNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).DeselectNetworks(ctx, req.(*SelectNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_DebugBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebugBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).DebugBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/DebugBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).DebugBundle(ctx, req.(*DebugBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/GetLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetLogLevel(ctx, req.(*GetLogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_SetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).SetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/SetLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).SetLogLevel(ctx, req.(*SetLogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_ListStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).ListStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/ListStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).ListStates(ctx, req.(*ListStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_CleanState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).CleanState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/CleanState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).CleanState(ctx, req.(*CleanStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_DeleteState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).DeleteState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/DeleteState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).DeleteState(ctx, req.(*DeleteStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_SetNetworkMapPersistence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNetworkMapPersistenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).SetNetworkMapPersistence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/SetNetworkMapPersistence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).SetNetworkMapPersistence(ctx, req.(*SetNetworkMapPersistenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_TracePacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TracePacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).TracePacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.DaemonService/TracePacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).TracePacket(ctx, req.(*TracePacketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DaemonService_ServiceDesc is the grpc.ServiceDesc for DaemonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DaemonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "daemon.DaemonService",
	HandlerType: (*DaemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _DaemonService_Login_Handler,
		},
		{
			MethodName: "WaitSSOLogin",
			Handler:    _DaemonService_WaitSSOLogin_Handler,
		},
		{
			MethodName: "Up",
			Handler:    _DaemonService_Up_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _DaemonService_Status_Handler,
		},
		{
			MethodName: "Down",
			Handler:    _DaemonService_Down_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _DaemonService_GetConfig_Handler,
		},
		{
			MethodName: "ListNetworks",
			Handler:    _DaemonService_ListNetworks_Handler,
		},
		{
			MethodName: "SelectNetworks",
			Handler:    _DaemonService_SelectNetworks_Handler,
		},
		{
			MethodName: "DeselectNetworks",
			Handler:    _DaemonService_DeselectNetworks_Handler,
		},
		{
			MethodName: "DebugBundle",
			Handler:    _DaemonService_DebugBundle_Handler,
		},
		{
			MethodName: "GetLogLevel",
			Handler:    _DaemonService_GetLogLevel_Handler,
		},
		{
			MethodName: "SetLogLevel",
			Handler:    _DaemonService_SetLogLevel_Handler,
		},
		{
			MethodName: "ListStates",
			Handler:    _DaemonService_ListStates_Handler,
		},
		{
			MethodName: "CleanState",
			Handler:    _DaemonService_CleanState_Handler,
		},
		{
			MethodName: "DeleteState",
			Handler:    _DaemonService_DeleteState_Handler,
		},
		{
			MethodName: "SetNetworkMapPersistence",
			Handler:    _DaemonService_SetNetworkMapPersistence_Handler,
		},
		{
			MethodName: "TracePacket",
			Handler:    _DaemonService_TracePacket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daemon.proto",
}
