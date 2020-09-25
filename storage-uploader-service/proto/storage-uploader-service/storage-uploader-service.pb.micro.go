// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/storage-uploader-service/storage-uploader-service.proto

package go_micro_service_storageuploaderservice

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for StorageUploaderService service

func NewStorageUploaderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for StorageUploaderService service

type StorageUploaderService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (StorageUploaderService_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (StorageUploaderService_PingPongService, error)
}

type storageUploaderService struct {
	c    client.Client
	name string
}

func NewStorageUploaderService(name string, c client.Client) StorageUploaderService {
	return &storageUploaderService{
		c:    c,
		name: name,
	}
}

func (c *storageUploaderService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "StorageUploaderService.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageUploaderService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (StorageUploaderService_StreamService, error) {
	req := c.c.NewRequest(c.name, "StorageUploaderService.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &storageUploaderServiceStream{stream}, nil
}

type StorageUploaderService_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type storageUploaderServiceStream struct {
	stream client.Stream
}

func (x *storageUploaderServiceStream) Close() error {
	return x.stream.Close()
}

func (x *storageUploaderServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *storageUploaderServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageUploaderServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageUploaderServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageUploaderService) PingPong(ctx context.Context, opts ...client.CallOption) (StorageUploaderService_PingPongService, error) {
	req := c.c.NewRequest(c.name, "StorageUploaderService.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &storageUploaderServicePingPong{stream}, nil
}

type StorageUploaderService_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type storageUploaderServicePingPong struct {
	stream client.Stream
}

func (x *storageUploaderServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *storageUploaderServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *storageUploaderServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageUploaderServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageUploaderServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *storageUploaderServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for StorageUploaderService service

type StorageUploaderServiceHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, StorageUploaderService_StreamStream) error
	PingPong(context.Context, StorageUploaderService_PingPongStream) error
}

func RegisterStorageUploaderServiceHandler(s server.Server, hdlr StorageUploaderServiceHandler, opts ...server.HandlerOption) error {
	type storageUploaderService interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type StorageUploaderService struct {
		storageUploaderService
	}
	h := &storageUploaderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StorageUploaderService{h}, opts...))
}

type storageUploaderServiceHandler struct {
	StorageUploaderServiceHandler
}

func (h *storageUploaderServiceHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.StorageUploaderServiceHandler.Call(ctx, in, out)
}

func (h *storageUploaderServiceHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.StorageUploaderServiceHandler.Stream(ctx, m, &storageUploaderServiceStreamStream{stream})
}

type StorageUploaderService_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type storageUploaderServiceStreamStream struct {
	stream server.Stream
}

func (x *storageUploaderServiceStreamStream) Close() error {
	return x.stream.Close()
}

func (x *storageUploaderServiceStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *storageUploaderServiceStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageUploaderServiceStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageUploaderServiceStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *storageUploaderServiceHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.StorageUploaderServiceHandler.PingPong(ctx, &storageUploaderServicePingPongStream{stream})
}

type StorageUploaderService_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type storageUploaderServicePingPongStream struct {
	stream server.Stream
}

func (x *storageUploaderServicePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *storageUploaderServicePingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *storageUploaderServicePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *storageUploaderServicePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *storageUploaderServicePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *storageUploaderServicePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
