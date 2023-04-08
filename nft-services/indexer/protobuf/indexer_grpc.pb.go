// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: indexer.proto

package protobuf

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

// IndexerRpcServiceClient is the client API for IndexerRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexerRpcServiceClient interface {
	GetLatestBlock(ctx context.Context, in *LatestBlock, opts ...grpc.CallOption) (*LatestBlockRep, error)
}

type indexerRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexerRpcServiceClient(cc grpc.ClientConnInterface) IndexerRpcServiceClient {
	return &indexerRpcServiceClient{cc}
}

func (c *indexerRpcServiceClient) GetLatestBlock(ctx context.Context, in *LatestBlock, opts ...grpc.CallOption) (*LatestBlockRep, error) {
	out := new(LatestBlockRep)
	err := c.cc.Invoke(ctx, "/indexer.protobuf.IndexerRpcService/GetLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexerRpcServiceServer is the server API for IndexerRpcService service.
// All implementations must embed UnimplementedIndexerRpcServiceServer
// for forward compatibility
type IndexerRpcServiceServer interface {
	GetLatestBlock(context.Context, *LatestBlock) (*LatestBlockRep, error)
}

// UnimplementedIndexerRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIndexerRpcServiceServer struct {
}

func (UnimplementedIndexerRpcServiceServer) GetLatestBlock(context.Context, *LatestBlock) (*LatestBlockRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlock not implemented")
}
func (UnimplementedIndexerRpcServiceServer) mustEmbedUnimplementedIndexerRpcServiceServer() {}

// UnsafeIndexerRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexerRpcServiceServer will
// result in compilation errors.
type UnsafeIndexerRpcServiceServer interface {
	mustEmbedUnimplementedIndexerRpcServiceServer()
}

func RegisterIndexerRpcServiceServer(s grpc.ServiceRegistrar, srv IndexerRpcServiceServer) {
	s.RegisterService(&IndexerRpcService_ServiceDesc, srv)
}

func _IndexerRpcService_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerRpcServiceServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indexer.protobuf.IndexerRpcService/GetLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerRpcServiceServer).GetLatestBlock(ctx, req.(*LatestBlock))
	}
	return interceptor(ctx, in, info, handler)
}

// IndexerRpcService_ServiceDesc is the grpc.ServiceDesc for IndexerRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IndexerRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "indexer.protobuf.IndexerRpcService",
	HandlerType: (*IndexerRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestBlock",
			Handler:    _IndexerRpcService_GetLatestBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indexer.proto",
}