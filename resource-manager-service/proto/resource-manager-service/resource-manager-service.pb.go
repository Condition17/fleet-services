// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.0
// source: proto/resource-manager-service/resource-manager-service.proto

package go_micro_api_ResourceManagerService

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
		mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[0]
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
	return file_proto_resource_manager_service_resource_manager_service_proto_rawDescGZIP(), []int{0}
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

type TestRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FileId string `protobuf:"bytes,3,opt,name=fileId,proto3" json:"fileId,omitempty"`
	UserId uint32 `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
	User   *User  `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *TestRun) Reset() {
	*x = TestRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRun) ProtoMessage() {}

func (x *TestRun) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[1]
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
	return file_proto_resource_manager_service_resource_manager_service_proto_rawDescGZIP(), []int{1}
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

func (x *TestRun) GetFileId() string {
	if x != nil {
		return x.FileId
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

type FileSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IP                  string   `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Name                string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	FileShareCapacityGb int64    `protobuf:"varint,4,opt,name=fileShareCapacityGb,proto3" json:"fileShareCapacityGb,omitempty"`
	FileShareName       string   `protobuf:"bytes,5,opt,name=fileShareName,proto3" json:"fileShareName,omitempty"`
	TestRunId           uint32   `protobuf:"varint,6,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
	TestRun             *TestRun `protobuf:"bytes,7,opt,name=testRun,proto3" json:"testRun,omitempty"`
}

func (x *FileSystem) Reset() {
	*x = FileSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSystem) ProtoMessage() {}

func (x *FileSystem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSystem.ProtoReflect.Descriptor instead.
func (*FileSystem) Descriptor() ([]byte, []int) {
	return file_proto_resource_manager_service_resource_manager_service_proto_rawDescGZIP(), []int{2}
}

func (x *FileSystem) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileSystem) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *FileSystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileSystem) GetFileShareCapacityGb() int64 {
	if x != nil {
		return x.FileShareCapacityGb
	}
	return 0
}

func (x *FileSystem) GetFileShareName() string {
	if x != nil {
		return x.FileShareName
	}
	return ""
}

func (x *FileSystem) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

func (x *FileSystem) GetTestRun() *TestRun {
	if x != nil {
		return x.TestRun
	}
	return nil
}

type FileSystemDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSystem *FileSystem `protobuf:"bytes,1,opt,name=fileSystem,proto3" json:"fileSystem,omitempty"`
}

func (x *FileSystemDetails) Reset() {
	*x = FileSystemDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSystemDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSystemDetails) ProtoMessage() {}

func (x *FileSystemDetails) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSystemDetails.ProtoReflect.Descriptor instead.
func (*FileSystemDetails) Descriptor() ([]byte, []int) {
	return file_proto_resource_manager_service_resource_manager_service_proto_rawDescGZIP(), []int{3}
}

func (x *FileSystemDetails) GetFileSystem() *FileSystem {
	if x != nil {
		return x.FileSystem
	}
	return nil
}

type FileSystemSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId   uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
	SizeInBytes int64  `protobuf:"varint,2,opt,name=sizeInBytes,proto3" json:"sizeInBytes,omitempty"`
}

func (x *FileSystemSpec) Reset() {
	*x = FileSystemSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSystemSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSystemSpec) ProtoMessage() {}

func (x *FileSystemSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSystemSpec.ProtoReflect.Descriptor instead.
func (*FileSystemSpec) Descriptor() ([]byte, []int) {
	return file_proto_resource_manager_service_resource_manager_service_proto_rawDescGZIP(), []int{4}
}

func (x *FileSystemSpec) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

func (x *FileSystemSpec) GetSizeInBytes() int64 {
	if x != nil {
		return x.SizeInBytes
	}
	return 0
}

type ExecutorInstanceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestRunId uint32 `protobuf:"varint,1,opt,name=testRunId,proto3" json:"testRunId,omitempty"`
}

func (x *ExecutorInstanceSpec) Reset() {
	*x = ExecutorInstanceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutorInstanceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutorInstanceSpec) ProtoMessage() {}

func (x *ExecutorInstanceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutorInstanceSpec.ProtoReflect.Descriptor instead.
func (*ExecutorInstanceSpec) Descriptor() ([]byte, []int) {
	return file_proto_resource_manager_service_resource_manager_service_proto_rawDescGZIP(), []int{5}
}

func (x *ExecutorInstanceSpec) GetTestRunId() uint32 {
	if x != nil {
		return x.TestRunId
	}
	return 0
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[6]
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
	return file_proto_resource_manager_service_resource_manager_service_proto_rawDescGZIP(), []int{6}
}

var File_proto_resource_manager_service_resource_manager_service_proto protoreflect.FileDescriptor

var file_proto_resource_manager_service_resource_manager_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x23, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x5a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x9c, 0x01, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0xfe, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x47, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x13, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x47, 0x62, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74,
	0x52, 0x75, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e,
	0x22, 0x64, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x4f, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x50, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74,
	0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x6e,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x69, 0x7a,
	0x65, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x14, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x22, 0x0f,
	0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xaa, 0x03, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x13, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x33, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8c, 0x01,
	0x0a, 0x19, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7e, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x33, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x70,
	0x65, 0x63, 0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_resource_manager_service_resource_manager_service_proto_rawDescOnce sync.Once
	file_proto_resource_manager_service_resource_manager_service_proto_rawDescData = file_proto_resource_manager_service_resource_manager_service_proto_rawDesc
)

func file_proto_resource_manager_service_resource_manager_service_proto_rawDescGZIP() []byte {
	file_proto_resource_manager_service_resource_manager_service_proto_rawDescOnce.Do(func() {
		file_proto_resource_manager_service_resource_manager_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_resource_manager_service_resource_manager_service_proto_rawDescData)
	})
	return file_proto_resource_manager_service_resource_manager_service_proto_rawDescData
}

var file_proto_resource_manager_service_resource_manager_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_resource_manager_service_resource_manager_service_proto_goTypes = []interface{}{
	(*User)(nil),                 // 0: go.micro.api.ResourceManagerService.User
	(*TestRun)(nil),              // 1: go.micro.api.ResourceManagerService.TestRun
	(*FileSystem)(nil),           // 2: go.micro.api.ResourceManagerService.FileSystem
	(*FileSystemDetails)(nil),    // 3: go.micro.api.ResourceManagerService.FileSystemDetails
	(*FileSystemSpec)(nil),       // 4: go.micro.api.ResourceManagerService.FileSystemSpec
	(*ExecutorInstanceSpec)(nil), // 5: go.micro.api.ResourceManagerService.ExecutorInstanceSpec
	(*EmptyResponse)(nil),        // 6: go.micro.api.ResourceManagerService.EmptyResponse
}
var file_proto_resource_manager_service_resource_manager_service_proto_depIdxs = []int32{
	0, // 0: go.micro.api.ResourceManagerService.TestRun.user:type_name -> go.micro.api.ResourceManagerService.User
	1, // 1: go.micro.api.ResourceManagerService.FileSystem.testRun:type_name -> go.micro.api.ResourceManagerService.TestRun
	2, // 2: go.micro.api.ResourceManagerService.FileSystemDetails.fileSystem:type_name -> go.micro.api.ResourceManagerService.FileSystem
	4, // 3: go.micro.api.ResourceManagerService.ResourceManagerService.ProvisionFileSystem:input_type -> go.micro.api.ResourceManagerService.FileSystemSpec
	5, // 4: go.micro.api.ResourceManagerService.ResourceManagerService.ProvisionExecutorInstance:input_type -> go.micro.api.ResourceManagerService.ExecutorInstanceSpec
	4, // 5: go.micro.api.ResourceManagerService.ResourceManagerService.GetFileSystem:input_type -> go.micro.api.ResourceManagerService.FileSystemSpec
	6, // 6: go.micro.api.ResourceManagerService.ResourceManagerService.ProvisionFileSystem:output_type -> go.micro.api.ResourceManagerService.EmptyResponse
	6, // 7: go.micro.api.ResourceManagerService.ResourceManagerService.ProvisionExecutorInstance:output_type -> go.micro.api.ResourceManagerService.EmptyResponse
	3, // 8: go.micro.api.ResourceManagerService.ResourceManagerService.GetFileSystem:output_type -> go.micro.api.ResourceManagerService.FileSystemDetails
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_resource_manager_service_resource_manager_service_proto_init() }
func file_proto_resource_manager_service_resource_manager_service_proto_init() {
	if File_proto_resource_manager_service_resource_manager_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSystem); i {
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
		file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSystemDetails); i {
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
		file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSystemSpec); i {
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
		file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutorInstanceSpec); i {
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
		file_proto_resource_manager_service_resource_manager_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_resource_manager_service_resource_manager_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_resource_manager_service_resource_manager_service_proto_goTypes,
		DependencyIndexes: file_proto_resource_manager_service_resource_manager_service_proto_depIdxs,
		MessageInfos:      file_proto_resource_manager_service_resource_manager_service_proto_msgTypes,
	}.Build()
	File_proto_resource_manager_service_resource_manager_service_proto = out.File
	file_proto_resource_manager_service_resource_manager_service_proto_rawDesc = nil
	file_proto_resource_manager_service_resource_manager_service_proto_goTypes = nil
	file_proto_resource_manager_service_resource_manager_service_proto_depIdxs = nil
}
