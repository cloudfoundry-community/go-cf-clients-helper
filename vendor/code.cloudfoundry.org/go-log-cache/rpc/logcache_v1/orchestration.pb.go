// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: logcache_v1/orchestration.proto

package logcache_v1

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

type Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start is the first hash within the given range. [start..end]
	Start uint64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end is the last hash within the given range. [start..end]
	End uint64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Range) Reset() {
	*x = Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Range) ProtoMessage() {}

func (x *Range) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Range.ProtoReflect.Descriptor instead.
func (*Range) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{0}
}

func (x *Range) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Range) GetEnd() uint64 {
	if x != nil {
		return x.End
	}
	return 0
}

type Ranges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranges []*Range `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *Ranges) Reset() {
	*x = Ranges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ranges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ranges) ProtoMessage() {}

func (x *Ranges) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ranges.ProtoReflect.Descriptor instead.
func (*Ranges) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{1}
}

func (x *Ranges) GetRanges() []*Range {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type AddRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range *Range `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *AddRangeRequest) Reset() {
	*x = AddRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRangeRequest) ProtoMessage() {}

func (x *AddRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRangeRequest.ProtoReflect.Descriptor instead.
func (*AddRangeRequest) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{2}
}

func (x *AddRangeRequest) GetRange() *Range {
	if x != nil {
		return x.Range
	}
	return nil
}

type AddRangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddRangeResponse) Reset() {
	*x = AddRangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRangeResponse) ProtoMessage() {}

func (x *AddRangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRangeResponse.ProtoReflect.Descriptor instead.
func (*AddRangeResponse) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{3}
}

type RemoveRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range *Range `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *RemoveRangeRequest) Reset() {
	*x = RemoveRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRangeRequest) ProtoMessage() {}

func (x *RemoveRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRangeRequest.ProtoReflect.Descriptor instead.
func (*RemoveRangeRequest) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveRangeRequest) GetRange() *Range {
	if x != nil {
		return x.Range
	}
	return nil
}

type RemoveRangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveRangeResponse) Reset() {
	*x = RemoveRangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRangeResponse) ProtoMessage() {}

func (x *RemoveRangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRangeResponse.ProtoReflect.Descriptor instead.
func (*RemoveRangeResponse) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{5}
}

type ListRangesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRangesRequest) Reset() {
	*x = ListRangesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRangesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRangesRequest) ProtoMessage() {}

func (x *ListRangesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRangesRequest.ProtoReflect.Descriptor instead.
func (*ListRangesRequest) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{6}
}

type ListRangesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranges []*Range `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *ListRangesResponse) Reset() {
	*x = ListRangesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRangesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRangesResponse) ProtoMessage() {}

func (x *ListRangesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRangesResponse.ProtoReflect.Descriptor instead.
func (*ListRangesResponse) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{7}
}

func (x *ListRangesResponse) GetRanges() []*Range {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type SetRangesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key is the address of the Log Cache node.
	Ranges map[string]*Ranges `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SetRangesRequest) Reset() {
	*x = SetRangesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRangesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRangesRequest) ProtoMessage() {}

func (x *SetRangesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRangesRequest.ProtoReflect.Descriptor instead.
func (*SetRangesRequest) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{8}
}

func (x *SetRangesRequest) GetRanges() map[string]*Ranges {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type SetRangesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetRangesResponse) Reset() {
	*x = SetRangesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logcache_v1_orchestration_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRangesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRangesResponse) ProtoMessage() {}

func (x *SetRangesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logcache_v1_orchestration_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRangesResponse.ProtoReflect.Descriptor instead.
func (*SetRangesResponse) Descriptor() ([]byte, []int) {
	return file_logcache_v1_orchestration_proto_rawDescGZIP(), []int{9}
}

var File_logcache_v1_orchestration_proto protoreflect.FileDescriptor

var file_logcache_v1_orchestration_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x6f, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x2f,
	0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22,
	0x34, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x67, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f,
	0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x40, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6c, 0x6f, 0x67, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0x4e, 0x0a, 0x0b,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c,
	0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x13, 0x0a, 0x11,
	0x53, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xcd, 0x02, 0x0a, 0x0d, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f, 0x2e,
	0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x12, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x12, 0x1d, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logcache_v1_orchestration_proto_rawDescOnce sync.Once
	file_logcache_v1_orchestration_proto_rawDescData = file_logcache_v1_orchestration_proto_rawDesc
)

func file_logcache_v1_orchestration_proto_rawDescGZIP() []byte {
	file_logcache_v1_orchestration_proto_rawDescOnce.Do(func() {
		file_logcache_v1_orchestration_proto_rawDescData = protoimpl.X.CompressGZIP(file_logcache_v1_orchestration_proto_rawDescData)
	})
	return file_logcache_v1_orchestration_proto_rawDescData
}

var file_logcache_v1_orchestration_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_logcache_v1_orchestration_proto_goTypes = []interface{}{
	(*Range)(nil),               // 0: logcache.v1.Range
	(*Ranges)(nil),              // 1: logcache.v1.Ranges
	(*AddRangeRequest)(nil),     // 2: logcache.v1.AddRangeRequest
	(*AddRangeResponse)(nil),    // 3: logcache.v1.AddRangeResponse
	(*RemoveRangeRequest)(nil),  // 4: logcache.v1.RemoveRangeRequest
	(*RemoveRangeResponse)(nil), // 5: logcache.v1.RemoveRangeResponse
	(*ListRangesRequest)(nil),   // 6: logcache.v1.ListRangesRequest
	(*ListRangesResponse)(nil),  // 7: logcache.v1.ListRangesResponse
	(*SetRangesRequest)(nil),    // 8: logcache.v1.SetRangesRequest
	(*SetRangesResponse)(nil),   // 9: logcache.v1.SetRangesResponse
	nil,                         // 10: logcache.v1.SetRangesRequest.RangesEntry
}
var file_logcache_v1_orchestration_proto_depIdxs = []int32{
	0,  // 0: logcache.v1.Ranges.ranges:type_name -> logcache.v1.Range
	0,  // 1: logcache.v1.AddRangeRequest.range:type_name -> logcache.v1.Range
	0,  // 2: logcache.v1.RemoveRangeRequest.range:type_name -> logcache.v1.Range
	0,  // 3: logcache.v1.ListRangesResponse.ranges:type_name -> logcache.v1.Range
	10, // 4: logcache.v1.SetRangesRequest.ranges:type_name -> logcache.v1.SetRangesRequest.RangesEntry
	1,  // 5: logcache.v1.SetRangesRequest.RangesEntry.value:type_name -> logcache.v1.Ranges
	2,  // 6: logcache.v1.Orchestration.AddRange:input_type -> logcache.v1.AddRangeRequest
	4,  // 7: logcache.v1.Orchestration.RemoveRange:input_type -> logcache.v1.RemoveRangeRequest
	6,  // 8: logcache.v1.Orchestration.ListRanges:input_type -> logcache.v1.ListRangesRequest
	8,  // 9: logcache.v1.Orchestration.SetRanges:input_type -> logcache.v1.SetRangesRequest
	3,  // 10: logcache.v1.Orchestration.AddRange:output_type -> logcache.v1.AddRangeResponse
	5,  // 11: logcache.v1.Orchestration.RemoveRange:output_type -> logcache.v1.RemoveRangeResponse
	7,  // 12: logcache.v1.Orchestration.ListRanges:output_type -> logcache.v1.ListRangesResponse
	9,  // 13: logcache.v1.Orchestration.SetRanges:output_type -> logcache.v1.SetRangesResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_logcache_v1_orchestration_proto_init() }
func file_logcache_v1_orchestration_proto_init() {
	if File_logcache_v1_orchestration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logcache_v1_orchestration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Range); i {
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
		file_logcache_v1_orchestration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ranges); i {
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
		file_logcache_v1_orchestration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRangeRequest); i {
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
		file_logcache_v1_orchestration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRangeResponse); i {
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
		file_logcache_v1_orchestration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRangeRequest); i {
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
		file_logcache_v1_orchestration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRangeResponse); i {
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
		file_logcache_v1_orchestration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRangesRequest); i {
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
		file_logcache_v1_orchestration_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRangesResponse); i {
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
		file_logcache_v1_orchestration_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRangesRequest); i {
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
		file_logcache_v1_orchestration_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRangesResponse); i {
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
			RawDescriptor: file_logcache_v1_orchestration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logcache_v1_orchestration_proto_goTypes,
		DependencyIndexes: file_logcache_v1_orchestration_proto_depIdxs,
		MessageInfos:      file_logcache_v1_orchestration_proto_msgTypes,
	}.Build()
	File_logcache_v1_orchestration_proto = out.File
	file_logcache_v1_orchestration_proto_rawDesc = nil
	file_logcache_v1_orchestration_proto_goTypes = nil
	file_logcache_v1_orchestration_proto_depIdxs = nil
}