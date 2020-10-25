// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/test-run-service/test-run-service.proto

package go_micro_service_testrunservice

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

// Api Endpoints for TestRunService service

func NewTestRunServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TestRunService service

type TestRunService interface {
	Create(ctx context.Context, in *CreateTestRunRequest, opts ...client.CallOption) (*TestRunDetails, error)
	Get(ctx context.Context, in *TestRun, opts ...client.CallOption) (*TestRunDetails, error)
	List(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*ListResponse, error)
	Delete(ctx context.Context, in *TestRun, opts ...client.CallOption) (*EmptyResponse, error)
	AssignFile(ctx context.Context, in *AssignRequest, opts ...client.CallOption) (*EmptyResponse, error)
}

type testRunService struct {
	c    client.Client
	name string
}

func NewTestRunService(name string, c client.Client) TestRunService {
	return &testRunService{
		c:    c,
		name: name,
	}
}

func (c *testRunService) Create(ctx context.Context, in *CreateTestRunRequest, opts ...client.CallOption) (*TestRunDetails, error) {
	req := c.c.NewRequest(c.name, "TestRunService.Create", in)
	out := new(TestRunDetails)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testRunService) Get(ctx context.Context, in *TestRun, opts ...client.CallOption) (*TestRunDetails, error) {
	req := c.c.NewRequest(c.name, "TestRunService.Get", in)
	out := new(TestRunDetails)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testRunService) List(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "TestRunService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testRunService) Delete(ctx context.Context, in *TestRun, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "TestRunService.Delete", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testRunService) AssignFile(ctx context.Context, in *AssignRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "TestRunService.AssignFile", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestRunService service

type TestRunServiceHandler interface {
	Create(context.Context, *CreateTestRunRequest, *TestRunDetails) error
	Get(context.Context, *TestRun, *TestRunDetails) error
	List(context.Context, *EmptyRequest, *ListResponse) error
	Delete(context.Context, *TestRun, *EmptyResponse) error
	AssignFile(context.Context, *AssignRequest, *EmptyResponse) error
}

func RegisterTestRunServiceHandler(s server.Server, hdlr TestRunServiceHandler, opts ...server.HandlerOption) error {
	type testRunService interface {
		Create(ctx context.Context, in *CreateTestRunRequest, out *TestRunDetails) error
		Get(ctx context.Context, in *TestRun, out *TestRunDetails) error
		List(ctx context.Context, in *EmptyRequest, out *ListResponse) error
		Delete(ctx context.Context, in *TestRun, out *EmptyResponse) error
		AssignFile(ctx context.Context, in *AssignRequest, out *EmptyResponse) error
	}
	type TestRunService struct {
		testRunService
	}
	h := &testRunServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TestRunService{h}, opts...))
}

type testRunServiceHandler struct {
	TestRunServiceHandler
}

func (h *testRunServiceHandler) Create(ctx context.Context, in *CreateTestRunRequest, out *TestRunDetails) error {
	return h.TestRunServiceHandler.Create(ctx, in, out)
}

func (h *testRunServiceHandler) Get(ctx context.Context, in *TestRun, out *TestRunDetails) error {
	return h.TestRunServiceHandler.Get(ctx, in, out)
}

func (h *testRunServiceHandler) List(ctx context.Context, in *EmptyRequest, out *ListResponse) error {
	return h.TestRunServiceHandler.List(ctx, in, out)
}

func (h *testRunServiceHandler) Delete(ctx context.Context, in *TestRun, out *EmptyResponse) error {
	return h.TestRunServiceHandler.Delete(ctx, in, out)
}

func (h *testRunServiceHandler) AssignFile(ctx context.Context, in *AssignRequest, out *EmptyResponse) error {
	return h.TestRunServiceHandler.AssignFile(ctx, in, out)
}
