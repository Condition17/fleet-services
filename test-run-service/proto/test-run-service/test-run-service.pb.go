// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.0
// source: proto/test-run-service/test-run-service.proto

package go_micro_api_testRunService

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateTestRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRun  *TestRun  `protobuf:"bytes,1,opt,name=testRun,proto3" json:"testRun,omitempty"`
	FileSpec *FileSpec `protobuf:"bytes,2,opt,name=fileSpec,proto3" json:"fileSpec,omitempty"`
}

func (x *CreateTestRunRequest) Reset() {
	*x = CreateTestRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestRunRequest) ProtoMessage() {}

func (x *CreateTestRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestRunRequest.ProtoReflect.Descriptor instead.
func (*CreateTestRunRequest) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTestRunRequest) GetTestRun() *TestRun {
	if x != nil {
		return x.TestRun
	}
	return nil
}

func (x *CreateTestRunRequest) GetFileSpec() *FileSpec {
	if x != nil {
		return x.FileSpec
	}
	return nil
}

type FileSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size         int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	MaxChunkSize uint32 `protobuf:"varint,4,opt,name=maxChunkSize,proto3" json:"maxChunkSize,omitempty"`
}

func (x *FileSpec) Reset() {
	*x = FileSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSpec) ProtoMessage() {}

func (x *FileSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSpec.ProtoReflect.Descriptor instead.
func (*FileSpec) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{1}
}

func (x *FileSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FileSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileSpec) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileSpec) GetMaxChunkSize() uint32 {
	if x != nil {
		return x.MaxChunkSize
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company string `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type RunIssue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId         uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
	BinaryPath        string `protobuf:"bytes,2,opt,name=binaryPath,proto3" json:"binaryPath,omitempty"`
	Issue             string `protobuf:"bytes,3,opt,name=issue,proto3" json:"issue,omitempty"`
	InputBytes        []byte `protobuf:"bytes,4,opt,name=inputBytes,proto3" json:"inputBytes,omitempty"`
	InputBytesPreview []byte `protobuf:"bytes,5,opt,name=inputBytesPreview,proto3" json:"inputBytesPreview,omitempty"`
	InputBinUrl       string `protobuf:"bytes,6,opt,name=inputBinUrl,proto3" json:"inputBinUrl,omitempty"`
}

func (x *RunIssue) Reset() {
	*x = RunIssue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunIssue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunIssue) ProtoMessage() {}

func (x *RunIssue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunIssue.ProtoReflect.Descriptor instead.
func (*RunIssue) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{3}
}

func (x *RunIssue) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

func (x *RunIssue) GetBinaryPath() string {
	if x != nil {
		return x.BinaryPath
	}
	return ""
}

func (x *RunIssue) GetIssue() string {
	if x != nil {
		return x.Issue
	}
	return ""
}

func (x *RunIssue) GetInputBytes() []byte {
	if x != nil {
		return x.InputBytes
	}
	return nil
}

func (x *RunIssue) GetInputBytesPreview() []byte {
	if x != nil {
		return x.InputBytesPreview
	}
	return nil
}

func (x *RunIssue) GetInputBinUrl() string {
	if x != nil {
		return x.InputBinUrl
	}
	return ""
}

type TestRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId        uint32      `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
	User          *User       `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	State         string      `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	StateMetadata string      `protobuf:"bytes,7,opt,name=stateMetadata,proto3" json:"stateMetadata,omitempty"`
	RunIssues     []*RunIssue `protobuf:"bytes,8,rep,name=runIssues,proto3" json:"runIssues,omitempty"`
}

func (x *TestRun) Reset() {
	*x = TestRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRun) ProtoMessage() {}

func (x *TestRun) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRun.ProtoReflect.Descriptor instead.
func (*TestRun) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{4}
}

func (x *TestRun) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestRun) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestRun) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TestRun) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TestRun) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *TestRun) GetStateMetadata() string {
	if x != nil {
		return x.StateMetadata
	}
	return ""
}

func (x *TestRun) GetRunIssues() []*RunIssue {
	if x != nil {
		return x.RunIssues
	}
	return nil
}

type TestRunDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRun *TestRun `protobuf:"bytes,1,opt,name=testRun,proto3" json:"testRun,omitempty"`
}

func (x *TestRunDetails) Reset() {
	*x = TestRunDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRunDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRunDetails) ProtoMessage() {}

func (x *TestRunDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRunDetails.ProtoReflect.Descriptor instead.
func (*TestRunDetails) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{5}
}

func (x *TestRunDetails) GetTestRun() *TestRun {
	if x != nil {
		return x.TestRun
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRuns []*TestRun `protobuf:"bytes,1,rep,name=testRuns,proto3" json:"testRuns,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListResponse) GetTestRuns() []*TestRun {
	if x != nil {
		return x.TestRuns
	}
	return nil
}

type TestRunStateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId     uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
	State         string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	StateMetadata string `protobuf:"bytes,3,opt,name=stateMetadata,proto3" json:"stateMetadata,omitempty"`
}

func (x *TestRunStateSpec) Reset() {
	*x = TestRunStateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRunStateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRunStateSpec) ProtoMessage() {}

func (x *TestRunStateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRunStateSpec.ProtoReflect.Descriptor instead.
func (*TestRunStateSpec) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{7}
}

func (x *TestRunStateSpec) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

func (x *TestRunStateSpec) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *TestRunStateSpec) GetStateMetadata() string {
	if x != nil {
		return x.StateMetadata
	}
	return ""
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{8}
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_run_service_test_run_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_proto_test_run_service_test_run_service_proto_rawDescGZIP(), []int{9}
}

var File_proto_test_run_service_test_run_service_proto protoreflect.FileDescriptor

var file_proto_test_run_service_test_run_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x72, 0x75, 0x6e,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x72, 0x75,
	0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x99, 0x01, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x07, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x41, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x22, 0x66, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x6d, 0x61, 0x78, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x5a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xce, 0x01, 0x0a,
	0x08, 0x52, 0x75, 0x6e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x73, 0x73, 0x75, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a,
	0x11, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x42, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0xfd, 0x01,
	0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52,
	0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x22, 0x50, 0x0a,
	0x0e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x3e, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x22,
	0x50, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x73, 0x22, 0x6c, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75,
	0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xc3, 0x05, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12,
	0x5a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x1a, 0x2b, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x75, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x1a, 0x2b, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x75, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x5e, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x29, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x22, 0x00, 0x12, 0x67, 0x0a,
	0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x75, 0x6e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_test_run_service_test_run_service_proto_rawDescOnce sync.Once
	file_proto_test_run_service_test_run_service_proto_rawDescData = file_proto_test_run_service_test_run_service_proto_rawDesc
)

func file_proto_test_run_service_test_run_service_proto_rawDescGZIP() []byte {
	file_proto_test_run_service_test_run_service_proto_rawDescOnce.Do(func() {
		file_proto_test_run_service_test_run_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_test_run_service_test_run_service_proto_rawDescData)
	})
	return file_proto_test_run_service_test_run_service_proto_rawDescData
}

var file_proto_test_run_service_test_run_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_test_run_service_test_run_service_proto_goTypes = []interface{}{
	(*CreateTestRunRequest)(nil), // 0: go.micro.api.testRunService.CreateTestRunRequest
	(*FileSpec)(nil),             // 1: go.micro.api.testRunService.FileSpec
	(*User)(nil),                 // 2: go.micro.api.testRunService.User
	(*RunIssue)(nil),             // 3: go.micro.api.testRunService.RunIssue
	(*TestRun)(nil),              // 4: go.micro.api.testRunService.TestRun
	(*TestRunDetails)(nil),       // 5: go.micro.api.testRunService.TestRunDetails
	(*ListResponse)(nil),         // 6: go.micro.api.testRunService.ListResponse
	(*TestRunStateSpec)(nil),     // 7: go.micro.api.testRunService.TestRunStateSpec
	(*EmptyRequest)(nil),         // 8: go.micro.api.testRunService.EmptyRequest
	(*EmptyResponse)(nil),        // 9: go.micro.api.testRunService.EmptyResponse
}
var file_proto_test_run_service_test_run_service_proto_depIdxs = []int32{
	4,  // 0: go.micro.api.testRunService.CreateTestRunRequest.testRun:type_name -> go.micro.api.testRunService.TestRun
	1,  // 1: go.micro.api.testRunService.CreateTestRunRequest.fileSpec:type_name -> go.micro.api.testRunService.FileSpec
	2,  // 2: go.micro.api.testRunService.TestRun.user:type_name -> go.micro.api.testRunService.User
	3,  // 3: go.micro.api.testRunService.TestRun.runIssues:type_name -> go.micro.api.testRunService.RunIssue
	4,  // 4: go.micro.api.testRunService.TestRunDetails.testRun:type_name -> go.micro.api.testRunService.TestRun
	4,  // 5: go.micro.api.testRunService.ListResponse.testRuns:type_name -> go.micro.api.testRunService.TestRun
	0,  // 6: go.micro.api.testRunService.TestRunService.Create:input_type -> go.micro.api.testRunService.CreateTestRunRequest
	4,  // 7: go.micro.api.testRunService.TestRunService.Get:input_type -> go.micro.api.testRunService.TestRun
	4,  // 8: go.micro.api.testRunService.TestRunService.GetById:input_type -> go.micro.api.testRunService.TestRun
	8,  // 9: go.micro.api.testRunService.TestRunService.List:input_type -> go.micro.api.testRunService.EmptyRequest
	4,  // 10: go.micro.api.testRunService.TestRunService.Delete:input_type -> go.micro.api.testRunService.TestRun
	7,  // 11: go.micro.api.testRunService.TestRunService.ChangeState:input_type -> go.micro.api.testRunService.TestRunStateSpec
	3,  // 12: go.micro.api.testRunService.TestRunService.RegisterRunIssue:input_type -> go.micro.api.testRunService.RunIssue
	5,  // 13: go.micro.api.testRunService.TestRunService.Create:output_type -> go.micro.api.testRunService.TestRunDetails
	5,  // 14: go.micro.api.testRunService.TestRunService.Get:output_type -> go.micro.api.testRunService.TestRunDetails
	5,  // 15: go.micro.api.testRunService.TestRunService.GetById:output_type -> go.micro.api.testRunService.TestRunDetails
	6,  // 16: go.micro.api.testRunService.TestRunService.List:output_type -> go.micro.api.testRunService.ListResponse
	9,  // 17: go.micro.api.testRunService.TestRunService.Delete:output_type -> go.micro.api.testRunService.EmptyResponse
	4,  // 18: go.micro.api.testRunService.TestRunService.ChangeState:output_type -> go.micro.api.testRunService.TestRun
	9,  // 19: go.micro.api.testRunService.TestRunService.RegisterRunIssue:output_type -> go.micro.api.testRunService.EmptyResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_test_run_service_test_run_service_proto_init() }
func file_proto_test_run_service_test_run_service_proto_init() {
	if File_proto_test_run_service_test_run_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_test_run_service_test_run_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestRunRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_run_service_test_run_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_run_service_test_run_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_run_service_test_run_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunIssue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_run_service_test_run_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRun); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_run_service_test_run_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRunDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_run_service_test_run_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_run_service_test_run_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRunStateSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_run_service_test_run_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_test_run_service_test_run_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_test_run_service_test_run_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_test_run_service_test_run_service_proto_goTypes,
		DependencyIndexes: file_proto_test_run_service_test_run_service_proto_depIdxs,
		MessageInfos:      file_proto_test_run_service_test_run_service_proto_msgTypes,
	}.Build()
	File_proto_test_run_service_test_run_service_proto = out.File
	file_proto_test_run_service_test_run_service_proto_rawDesc = nil
	file_proto_test_run_service_test_run_service_proto_goTypes = nil
	file_proto_test_run_service_test_run_service_proto_depIdxs = nil
}
