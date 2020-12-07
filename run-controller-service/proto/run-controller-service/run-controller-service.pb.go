// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.0
// source: proto/run-controller-service/run-controller-service.proto

package go_micro_api_runControllerService

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

type EventMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User          []byte `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Authorization []byte `protobuf:"bytes,2,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (x *EventMetadata) Reset() {
	*x = EventMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetadata) ProtoMessage() {}

func (x *EventMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetadata.ProtoReflect.Descriptor instead.
func (*EventMetadata) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetadata) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *EventMetadata) GetAuthorization() []byte {
	if x != nil {
		return x.Authorization
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string         `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Meta *EventMetadata `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	Data []byte         `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event) GetMeta() *EventMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Event) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type TestRunSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestRunSpec) Reset() {
	*x = TestRunSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRunSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRunSpec) ProtoMessage() {}

func (x *TestRunSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRunSpec.ProtoReflect.Descriptor instead.
func (*TestRunSpec) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{2}
}

func (x *TestRunSpec) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestRunSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FileSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TestRunId    uint32 `protobuf:"varint,2,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Size         int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	MaxChunkSize uint32 `protobuf:"varint,5,opt,name=maxChunkSize,proto3" json:"maxChunkSize,omitempty"`
}

func (x *FileSpec) Reset() {
	*x = FileSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSpec) ProtoMessage() {}

func (x *FileSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[3]
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
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{3}
}

func (x *FileSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FileSpec) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
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

type TestRunInitiatedEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunSpec *TestRunSpec `protobuf:"bytes,1,opt,name=testRunSpec,proto3" json:"testRunSpec,omitempty"`
}

func (x *TestRunInitiatedEventData) Reset() {
	*x = TestRunInitiatedEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRunInitiatedEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRunInitiatedEventData) ProtoMessage() {}

func (x *TestRunInitiatedEventData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRunInitiatedEventData.ProtoReflect.Descriptor instead.
func (*TestRunInitiatedEventData) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{4}
}

func (x *TestRunInitiatedEventData) GetTestRunSpec() *TestRunSpec {
	if x != nil {
		return x.TestRunSpec
	}
	return nil
}

type FileChunksUploadedEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSpec *FileSpec `protobuf:"bytes,2,opt,name=fileSpec,proto3" json:"fileSpec,omitempty"`
}

func (x *FileChunksUploadedEventData) Reset() {
	*x = FileChunksUploadedEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunksUploadedEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunksUploadedEventData) ProtoMessage() {}

func (x *FileChunksUploadedEventData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunksUploadedEventData.ProtoReflect.Descriptor instead.
func (*FileChunksUploadedEventData) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{5}
}

func (x *FileChunksUploadedEventData) GetFileSpec() *FileSpec {
	if x != nil {
		return x.FileSpec
	}
	return nil
}

// Events sent to WSS
type WssEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Target []byte `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WssEvent) Reset() {
	*x = WssEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WssEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WssEvent) ProtoMessage() {}

func (x *WssEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WssEvent.ProtoReflect.Descriptor instead.
func (*WssEvent) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{6}
}

func (x *WssEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *WssEvent) GetTarget() []byte {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *WssEvent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type FileSystemProvisionedEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
}

func (x *FileSystemProvisionedEventData) Reset() {
	*x = FileSystemProvisionedEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSystemProvisionedEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSystemProvisionedEventData) ProtoMessage() {}

func (x *FileSystemProvisionedEventData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSystemProvisionedEventData.ProtoReflect.Descriptor instead.
func (*FileSystemProvisionedEventData) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{7}
}

func (x *FileSystemProvisionedEventData) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

type ExecutorInstanceProvisionedEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
}

func (x *ExecutorInstanceProvisionedEventData) Reset() {
	*x = ExecutorInstanceProvisionedEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutorInstanceProvisionedEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutorInstanceProvisionedEventData) ProtoMessage() {}

func (x *ExecutorInstanceProvisionedEventData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutorInstanceProvisionedEventData.ProtoReflect.Descriptor instead.
func (*ExecutorInstanceProvisionedEventData) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{8}
}

func (x *ExecutorInstanceProvisionedEventData) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

type FileAssemblySucceededEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
}

func (x *FileAssemblySucceededEventData) Reset() {
	*x = FileAssemblySucceededEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileAssemblySucceededEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileAssemblySucceededEventData) ProtoMessage() {}

func (x *FileAssemblySucceededEventData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileAssemblySucceededEventData.ProtoReflect.Descriptor instead.
func (*FileAssemblySucceededEventData) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{9}
}

func (x *FileAssemblySucceededEventData) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

type FileEvaluationFinishedEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
}

func (x *FileEvaluationFinishedEventData) Reset() {
	*x = FileEvaluationFinishedEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileEvaluationFinishedEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileEvaluationFinishedEventData) ProtoMessage() {}

func (x *FileEvaluationFinishedEventData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileEvaluationFinishedEventData.ProtoReflect.Descriptor instead.
func (*FileEvaluationFinishedEventData) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{10}
}

func (x *FileEvaluationFinishedEventData) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

type ServiceErrorEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
	Source    string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Error     []byte `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServiceErrorEventData) Reset() {
	*x = ServiceErrorEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceErrorEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceErrorEventData) ProtoMessage() {}

func (x *ServiceErrorEventData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_run_controller_service_run_controller_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceErrorEventData.ProtoReflect.Descriptor instead.
func (*ServiceErrorEventData) Descriptor() ([]byte, []int) {
	return file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP(), []int{11}
}

func (x *ServiceErrorEventData) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

func (x *ServiceErrorEventData) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ServiceErrorEventData) GetError() []byte {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_proto_run_controller_service_run_controller_service_proto protoreflect.FileDescriptor

var file_proto_run_controller_service_run_controller_service_proto_rawDesc = []byte{
	0x0a, 0x39, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x75, 0x6e, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72,
	0x75, 0x6e, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x75, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x49,
	0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x31, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61,
	0x78, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6d, 0x0a, 0x19, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x50, 0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x52,
	0x75, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x75, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x22, 0x66, 0x0a, 0x1b, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x75, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x22, 0x4a, 0x0a, 0x08, 0x57, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a,
	0x1e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x22, 0x44, 0x0a,
	0x24, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75,
	0x6e, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x1e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x65, 0x6d,
	0x62, 0x6c, 0x79, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75,
	0x6e, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x1f, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52,
	0x75, 0x6e, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x16, 0x0a, 0x14, 0x52, 0x75, 0x6e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_run_controller_service_run_controller_service_proto_rawDescOnce sync.Once
	file_proto_run_controller_service_run_controller_service_proto_rawDescData = file_proto_run_controller_service_run_controller_service_proto_rawDesc
)

func file_proto_run_controller_service_run_controller_service_proto_rawDescGZIP() []byte {
	file_proto_run_controller_service_run_controller_service_proto_rawDescOnce.Do(func() {
		file_proto_run_controller_service_run_controller_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_run_controller_service_run_controller_service_proto_rawDescData)
	})
	return file_proto_run_controller_service_run_controller_service_proto_rawDescData
}

var file_proto_run_controller_service_run_controller_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_run_controller_service_run_controller_service_proto_goTypes = []interface{}{
	(*EventMetadata)(nil),                        // 0: go.micro.api.runControllerService.EventMetadata
	(*Event)(nil),                                // 1: go.micro.api.runControllerService.Event
	(*TestRunSpec)(nil),                          // 2: go.micro.api.runControllerService.TestRunSpec
	(*FileSpec)(nil),                             // 3: go.micro.api.runControllerService.FileSpec
	(*TestRunInitiatedEventData)(nil),            // 4: go.micro.api.runControllerService.TestRunInitiatedEventData
	(*FileChunksUploadedEventData)(nil),          // 5: go.micro.api.runControllerService.FileChunksUploadedEventData
	(*WssEvent)(nil),                             // 6: go.micro.api.runControllerService.WssEvent
	(*FileSystemProvisionedEventData)(nil),       // 7: go.micro.api.runControllerService.FileSystemProvisionedEventData
	(*ExecutorInstanceProvisionedEventData)(nil), // 8: go.micro.api.runControllerService.ExecutorInstanceProvisionedEventData
	(*FileAssemblySucceededEventData)(nil),       // 9: go.micro.api.runControllerService.FileAssemblySucceededEventData
	(*FileEvaluationFinishedEventData)(nil),      // 10: go.micro.api.runControllerService.FileEvaluationFinishedEventData
	(*ServiceErrorEventData)(nil),                // 11: go.micro.api.runControllerService.ServiceErrorEventData
}
var file_proto_run_controller_service_run_controller_service_proto_depIdxs = []int32{
	0, // 0: go.micro.api.runControllerService.Event.meta:type_name -> go.micro.api.runControllerService.EventMetadata
	2, // 1: go.micro.api.runControllerService.TestRunInitiatedEventData.testRunSpec:type_name -> go.micro.api.runControllerService.TestRunSpec
	3, // 2: go.micro.api.runControllerService.FileChunksUploadedEventData.fileSpec:type_name -> go.micro.api.runControllerService.FileSpec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_run_controller_service_run_controller_service_proto_init() }
func file_proto_run_controller_service_run_controller_service_proto_init() {
	if File_proto_run_controller_service_run_controller_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMetadata); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRunSpec); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRunInitiatedEventData); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunksUploadedEventData); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WssEvent); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSystemProvisionedEventData); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutorInstanceProvisionedEventData); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileAssemblySucceededEventData); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileEvaluationFinishedEventData); i {
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
		file_proto_run_controller_service_run_controller_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceErrorEventData); i {
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
			RawDescriptor: file_proto_run_controller_service_run_controller_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_run_controller_service_run_controller_service_proto_goTypes,
		DependencyIndexes: file_proto_run_controller_service_run_controller_service_proto_depIdxs,
		MessageInfos:      file_proto_run_controller_service_run_controller_service_proto_msgTypes,
	}.Build()
	File_proto_run_controller_service_run_controller_service_proto = out.File
	file_proto_run_controller_service_run_controller_service_proto_rawDesc = nil
	file_proto_run_controller_service_run_controller_service_proto_goTypes = nil
	file_proto_run_controller_service_run_controller_service_proto_depIdxs = nil
}
