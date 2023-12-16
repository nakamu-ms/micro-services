// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: rpc/main.proto

package pb

import (
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

type ModelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ModelsRequest) Reset() {
	*x = ModelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelsRequest) ProtoMessage() {}

func (x *ModelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelsRequest.ProtoReflect.Descriptor instead.
func (*ModelsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{0}
}

type ModelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ModelsResponse) Reset() {
	*x = ModelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelsResponse) ProtoMessage() {}

func (x *ModelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelsResponse.ProtoReflect.Descriptor instead.
func (*ModelsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{1}
}

type PromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PromptRequest) Reset() {
	*x = PromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptRequest) ProtoMessage() {}

func (x *PromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptRequest.ProtoReflect.Descriptor instead.
func (*PromptRequest) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{2}
}

type PromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PromptResponse) Reset() {
	*x = PromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptResponse) ProtoMessage() {}

func (x *PromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptResponse.ProtoReflect.Descriptor instead.
func (*PromptResponse) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{3}
}

type IndexesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IndexesRequest) Reset() {
	*x = IndexesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexesRequest) ProtoMessage() {}

func (x *IndexesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexesRequest.ProtoReflect.Descriptor instead.
func (*IndexesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{4}
}

type IndexesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IndexesResponse) Reset() {
	*x = IndexesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexesResponse) ProtoMessage() {}

func (x *IndexesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexesResponse.ProtoReflect.Descriptor instead.
func (*IndexesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{5}
}

type ChainsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChainsRequest) Reset() {
	*x = ChainsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainsRequest) ProtoMessage() {}

func (x *ChainsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainsRequest.ProtoReflect.Descriptor instead.
func (*ChainsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{6}
}

type ChainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChainsResponse) Reset() {
	*x = ChainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainsResponse) ProtoMessage() {}

func (x *ChainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainsResponse.ProtoReflect.Descriptor instead.
func (*ChainsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{7}
}

type AgentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AgentsRequest) Reset() {
	*x = AgentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentsRequest) ProtoMessage() {}

func (x *AgentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentsRequest.ProtoReflect.Descriptor instead.
func (*AgentsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{8}
}

type AgentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AgentsResponse) Reset() {
	*x = AgentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentsResponse) ProtoMessage() {}

func (x *AgentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentsResponse.ProtoReflect.Descriptor instead.
func (*AgentsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{9}
}

type MemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MemoryRequest) Reset() {
	*x = MemoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryRequest) ProtoMessage() {}

func (x *MemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryRequest.ProtoReflect.Descriptor instead.
func (*MemoryRequest) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{10}
}

type MemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MemoryResponse) Reset() {
	*x = MemoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_main_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryResponse) ProtoMessage() {}

func (x *MemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_main_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryResponse.ProtoReflect.Descriptor instead.
func (*MemoryResponse) Descriptor() ([]byte, []int) {
	return file_rpc_main_proto_rawDescGZIP(), []int{11}
}

var File_rpc_main_proto protoreflect.FileDescriptor

var file_rpc_main_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x22, 0x0f, 0x0a, 0x0d, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f,
	0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x10, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x10, 0x0a, 0x0e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a,
	0x0e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x88, 0x03, 0x0a, 0x09, 0x4c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x3d, 0x0a,
	0x06, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68,
	0x61, 0x6e, 0x69, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x18, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x2e, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x06, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68,
	0x61, 0x6e, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x2e, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x2e, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_main_proto_rawDescOnce sync.Once
	file_rpc_main_proto_rawDescData = file_rpc_main_proto_rawDesc
)

func file_rpc_main_proto_rawDescGZIP() []byte {
	file_rpc_main_proto_rawDescOnce.Do(func() {
		file_rpc_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_main_proto_rawDescData)
	})
	return file_rpc_main_proto_rawDescData
}

var file_rpc_main_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_rpc_main_proto_goTypes = []interface{}{
	(*ModelsRequest)(nil),   // 0: langchani.ModelsRequest
	(*ModelsResponse)(nil),  // 1: langchani.ModelsResponse
	(*PromptRequest)(nil),   // 2: langchani.PromptRequest
	(*PromptResponse)(nil),  // 3: langchani.PromptResponse
	(*IndexesRequest)(nil),  // 4: langchani.IndexesRequest
	(*IndexesResponse)(nil), // 5: langchani.IndexesResponse
	(*ChainsRequest)(nil),   // 6: langchani.ChainsRequest
	(*ChainsResponse)(nil),  // 7: langchani.ChainsResponse
	(*AgentsRequest)(nil),   // 8: langchani.AgentsRequest
	(*AgentsResponse)(nil),  // 9: langchani.AgentsResponse
	(*MemoryRequest)(nil),   // 10: langchani.MemoryRequest
	(*MemoryResponse)(nil),  // 11: langchani.MemoryResponse
}
var file_rpc_main_proto_depIdxs = []int32{
	0,  // 0: langchani.Langchain.Models:input_type -> langchani.ModelsRequest
	2,  // 1: langchani.Langchain.Prompt:input_type -> langchani.PromptRequest
	4,  // 2: langchani.Langchain.Indexes:input_type -> langchani.IndexesRequest
	6,  // 3: langchani.Langchain.Chains:input_type -> langchani.ChainsRequest
	8,  // 4: langchani.Langchain.Agents:input_type -> langchani.AgentsRequest
	10, // 5: langchani.Langchain.Memory:input_type -> langchani.MemoryRequest
	1,  // 6: langchani.Langchain.Models:output_type -> langchani.ModelsResponse
	3,  // 7: langchani.Langchain.Prompt:output_type -> langchani.PromptResponse
	5,  // 8: langchani.Langchain.Indexes:output_type -> langchani.IndexesResponse
	7,  // 9: langchani.Langchain.Chains:output_type -> langchani.ChainsResponse
	9,  // 10: langchani.Langchain.Agents:output_type -> langchani.AgentsResponse
	11, // 11: langchani.Langchain.Memory:output_type -> langchani.MemoryResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_main_proto_init() }
func file_rpc_main_proto_init() {
	if File_rpc_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelsRequest); i {
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
		file_rpc_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelsResponse); i {
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
		file_rpc_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptRequest); i {
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
		file_rpc_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptResponse); i {
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
		file_rpc_main_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexesRequest); i {
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
		file_rpc_main_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexesResponse); i {
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
		file_rpc_main_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainsRequest); i {
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
		file_rpc_main_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainsResponse); i {
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
		file_rpc_main_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentsRequest); i {
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
		file_rpc_main_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentsResponse); i {
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
		file_rpc_main_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryRequest); i {
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
		file_rpc_main_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryResponse); i {
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
			RawDescriptor: file_rpc_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_main_proto_goTypes,
		DependencyIndexes: file_rpc_main_proto_depIdxs,
		MessageInfos:      file_rpc_main_proto_msgTypes,
	}.Build()
	File_rpc_main_proto = out.File
	file_rpc_main_proto_rawDesc = nil
	file_rpc_main_proto_goTypes = nil
	file_rpc_main_proto_depIdxs = nil
}
