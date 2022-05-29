// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pb/person/person.proto

package person

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (*PersonRes, error)
	SearchIn(ctx context.Context, opts ...grpc.CallOption) (SearchService_SearchInClient, error)
	SearchOut(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (SearchService_SearchOutClient, error)
	SearchIO(ctx context.Context, opts ...grpc.CallOption) (SearchService_SearchIOClient, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (*PersonRes, error) {
	out := new(PersonRes)
	err := c.cc.Invoke(ctx, "/person.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchIn(ctx context.Context, opts ...grpc.CallOption) (SearchService_SearchInClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[0], "/person.SearchService/SearchIn", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchServiceSearchInClient{stream}
	return x, nil
}

type SearchService_SearchInClient interface {
	Send(*PersonReq) error
	CloseAndRecv() (*PersonRes, error)
	grpc.ClientStream
}

type searchServiceSearchInClient struct {
	grpc.ClientStream
}

func (x *searchServiceSearchInClient) Send(m *PersonReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *searchServiceSearchInClient) CloseAndRecv() (*PersonRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PersonRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *searchServiceClient) SearchOut(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (SearchService_SearchOutClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[1], "/person.SearchService/SearchOut", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchServiceSearchOutClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SearchService_SearchOutClient interface {
	Recv() (*PersonRes, error)
	grpc.ClientStream
}

type searchServiceSearchOutClient struct {
	grpc.ClientStream
}

func (x *searchServiceSearchOutClient) Recv() (*PersonRes, error) {
	m := new(PersonRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *searchServiceClient) SearchIO(ctx context.Context, opts ...grpc.CallOption) (SearchService_SearchIOClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearchService_ServiceDesc.Streams[2], "/person.SearchService/SearchIO", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchServiceSearchIOClient{stream}
	return x, nil
}

type SearchService_SearchIOClient interface {
	Send(*PersonReq) error
	Recv() (*PersonRes, error)
	grpc.ClientStream
}

type searchServiceSearchIOClient struct {
	grpc.ClientStream
}

func (x *searchServiceSearchIOClient) Send(m *PersonReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *searchServiceSearchIOClient) Recv() (*PersonRes, error) {
	m := new(PersonRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	Search(context.Context, *PersonReq) (*PersonRes, error)
	SearchIn(SearchService_SearchInServer) error
	SearchOut(*PersonReq, SearchService_SearchOutServer) error
	SearchIO(SearchService_SearchIOServer) error
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) Search(context.Context, *PersonReq) (*PersonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSearchServiceServer) SearchIn(SearchService_SearchInServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchIn not implemented")
}
func (UnimplementedSearchServiceServer) SearchOut(*PersonReq, SearchService_SearchOutServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchOut not implemented")
}
func (UnimplementedSearchServiceServer) SearchIO(SearchService_SearchIOServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchIO not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*PersonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchIn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SearchServiceServer).SearchIn(&searchServiceSearchInServer{stream})
}

type SearchService_SearchInServer interface {
	SendAndClose(*PersonRes) error
	Recv() (*PersonReq, error)
	grpc.ServerStream
}

type searchServiceSearchInServer struct {
	grpc.ServerStream
}

func (x *searchServiceSearchInServer) SendAndClose(m *PersonRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *searchServiceSearchInServer) Recv() (*PersonReq, error) {
	m := new(PersonReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SearchService_SearchOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PersonReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchServiceServer).SearchOut(m, &searchServiceSearchOutServer{stream})
}

type SearchService_SearchOutServer interface {
	Send(*PersonRes) error
	grpc.ServerStream
}

type searchServiceSearchOutServer struct {
	grpc.ServerStream
}

func (x *searchServiceSearchOutServer) Send(m *PersonRes) error {
	return x.ServerStream.SendMsg(m)
}

func _SearchService_SearchIO_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SearchServiceServer).SearchIO(&searchServiceSearchIOServer{stream})
}

type SearchService_SearchIOServer interface {
	Send(*PersonRes) error
	Recv() (*PersonReq, error)
	grpc.ServerStream
}

type searchServiceSearchIOServer struct {
	grpc.ServerStream
}

func (x *searchServiceSearchIOServer) Send(m *PersonRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *searchServiceSearchIOServer) Recv() (*PersonReq, error) {
	m := new(PersonReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "person.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchIn",
			Handler:       _SearchService_SearchIn_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SearchOut",
			Handler:       _SearchService_SearchOut_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchIO",
			Handler:       _SearchService_SearchIO_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/person/person.proto",
}
