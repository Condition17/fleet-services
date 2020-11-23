// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_micro_api_fileservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	CreateChunk(ctx context.Context, in *ChunkSpec, opts ...grpc.CallOption) (*EmptyResponse, error)
	CreateFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*Response, error)
	ReadFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*Response, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) CreateChunk(ctx context.Context, in *ChunkSpec, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/go.micro.api.fileservice.FileService/CreateChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) CreateFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.api.fileservice.FileService/CreateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) ReadFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.api.fileservice.FileService/ReadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	CreateChunk(context.Context, *ChunkSpec) (*EmptyResponse, error)
	CreateFile(context.Context, *File) (*Response, error)
	ReadFile(context.Context, *File) (*Response, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) CreateChunk(context.Context, *ChunkSpec) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChunk not implemented")
}
func (UnimplementedFileServiceServer) CreateFile(context.Context, *File) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedFileServiceServer) ReadFile(context.Context, *File) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_CreateChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.api.fileservice.FileService/CreateChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateChunk(ctx, req.(*ChunkSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.api.fileservice.FileService/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_ReadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).ReadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.api.fileservice.FileService/ReadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).ReadFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.api.fileservice.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChunk",
			Handler:    _FileService_CreateChunk_Handler,
		},
		{
			MethodName: "CreateFile",
			Handler:    _FileService_CreateFile_Handler,
		},
		{
			MethodName: "ReadFile",
			Handler:    _FileService_ReadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file-service.proto",
}
