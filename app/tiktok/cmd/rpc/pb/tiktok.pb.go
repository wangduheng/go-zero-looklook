// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: tiktok.proto

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

type LiveData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenUrl    string `protobuf:"bytes,1,opt,name=openUrl,proto3" json:"openUrl"`
	Info       string `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Id         int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	CreateTime int64  `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime"`
	OpenTime   int64  `protobuf:"varint,6,opt,name=openTime,proto3" json:"openTime"`
	IsNeedFork int64  `protobuf:"varint,7,opt,name=isNeedFork,proto3" json:"isNeedFork"`
}

func (x *LiveData) Reset() {
	*x = LiveData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveData) ProtoMessage() {}

func (x *LiveData) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveData.ProtoReflect.Descriptor instead.
func (*LiveData) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{0}
}

func (x *LiveData) GetOpenUrl() string {
	if x != nil {
		return x.OpenUrl
	}
	return ""
}

func (x *LiveData) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *LiveData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LiveData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LiveData) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *LiveData) GetOpenTime() int64 {
	if x != nil {
		return x.OpenTime
	}
	return 0
}

func (x *LiveData) GetIsNeedFork() int64 {
	if x != nil {
		return x.IsNeedFork
	}
	return 0
}

type SaveLiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenUrl    string `protobuf:"bytes,1,opt,name=openUrl,proto3" json:"openUrl"`
	Info       string `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	OpenTime   int64  `protobuf:"varint,4,opt,name=openTime,proto3" json:"openTime"`
	IsNeedFork int64  `protobuf:"varint,5,opt,name=isNeedFork,proto3" json:"isNeedFork"`
}

func (x *SaveLiveReq) Reset() {
	*x = SaveLiveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveLiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveLiveReq) ProtoMessage() {}

func (x *SaveLiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveLiveReq.ProtoReflect.Descriptor instead.
func (*SaveLiveReq) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{1}
}

func (x *SaveLiveReq) GetOpenUrl() string {
	if x != nil {
		return x.OpenUrl
	}
	return ""
}

func (x *SaveLiveReq) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *SaveLiveReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SaveLiveReq) GetOpenTime() int64 {
	if x != nil {
		return x.OpenTime
	}
	return 0
}

func (x *SaveLiveReq) GetIsNeedFork() int64 {
	if x != nil {
		return x.IsNeedFork
	}
	return 0
}

type SaveLiveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id"`
}

func (x *SaveLiveResp) Reset() {
	*x = SaveLiveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveLiveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveLiveResp) ProtoMessage() {}

func (x *SaveLiveResp) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveLiveResp.ProtoReflect.Descriptor instead.
func (*SaveLiveResp) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{2}
}

func (x *SaveLiveResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FetchLiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastSecond int64 `protobuf:"varint,1,opt,name=lastSecond,proto3" json:"lastSecond"`
	UserId     int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId"`
}

func (x *FetchLiveReq) Reset() {
	*x = FetchLiveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchLiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLiveReq) ProtoMessage() {}

func (x *FetchLiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLiveReq.ProtoReflect.Descriptor instead.
func (*FetchLiveReq) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{3}
}

func (x *FetchLiveReq) GetLastSecond() int64 {
	if x != nil {
		return x.LastSecond
	}
	return 0
}

func (x *FetchLiveReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FetchLiveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*LiveData `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
}

func (x *FetchLiveResp) Reset() {
	*x = FetchLiveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchLiveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLiveResp) ProtoMessage() {}

func (x *FetchLiveResp) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLiveResp.ProtoReflect.Descriptor instead.
func (*FetchLiveResp) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{4}
}

func (x *FetchLiveResp) GetList() []*LiveData {
	if x != nil {
		return x.List
	}
	return nil
}

type DeleteLiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteLiveReq) Reset() {
	*x = DeleteLiveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLiveReq) ProtoMessage() {}

func (x *DeleteLiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLiveReq.ProtoReflect.Descriptor instead.
func (*DeleteLiveReq) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteLiveReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteLiveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteLiveResp) Reset() {
	*x = DeleteLiveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLiveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLiveResp) ProtoMessage() {}

func (x *DeleteLiveResp) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLiveResp.ProtoReflect.Descriptor instead.
func (*DeleteLiveResp) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteLiveResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindLiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *FindLiveReq) Reset() {
	*x = FindLiveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLiveReq) ProtoMessage() {}

func (x *FindLiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLiveReq.ProtoReflect.Descriptor instead.
func (*FindLiveReq) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{7}
}

func (x *FindLiveReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindLiveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiveDate *LiveData `protobuf:"bytes,1,opt,name=liveDate,proto3" json:"liveDate"`
}

func (x *FindLiveResp) Reset() {
	*x = FindLiveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindLiveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindLiveResp) ProtoMessage() {}

func (x *FindLiveResp) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindLiveResp.ProtoReflect.Descriptor instead.
func (*FindLiveResp) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{8}
}

func (x *FindLiveResp) GetLiveDate() *LiveData {
	if x != nil {
		return x.LiveDate
	}
	return nil
}

var File_tiktok_proto protoreflect.FileDescriptor

var file_tiktok_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0xb8, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x70, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6b, 0x22, 0x8b, 0x01,
	0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x70, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x70, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x73, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x69, 0x73, 0x4e, 0x65, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6b, 0x22, 0x1e, 0x0a, 0x0c, 0x53,
	0x61, 0x76, 0x65, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x0c, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x6c,
	0x61, 0x73, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x46, 0x69, 0x6e,
	0x64, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64,
	0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x08, 0x6c, 0x69, 0x76, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6c, 0x69, 0x76, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x32, 0xd4, 0x01, 0x0a, 0x06, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x12, 0x2d, 0x0a,
	0x08, 0x73, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x0a,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x76, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x36, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x6e, 0x65,
	0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4c,
	0x69, 0x76, 0x65, 0x4f, 0x6e, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tiktok_proto_rawDescOnce sync.Once
	file_tiktok_proto_rawDescData = file_tiktok_proto_rawDesc
)

func file_tiktok_proto_rawDescGZIP() []byte {
	file_tiktok_proto_rawDescOnce.Do(func() {
		file_tiktok_proto_rawDescData = protoimpl.X.CompressGZIP(file_tiktok_proto_rawDescData)
	})
	return file_tiktok_proto_rawDescData
}

var file_tiktok_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tiktok_proto_goTypes = []interface{}{
	(*LiveData)(nil),       // 0: pb.LiveData
	(*SaveLiveReq)(nil),    // 1: pb.SaveLiveReq
	(*SaveLiveResp)(nil),   // 2: pb.SaveLiveResp
	(*FetchLiveReq)(nil),   // 3: pb.FetchLiveReq
	(*FetchLiveResp)(nil),  // 4: pb.FetchLiveResp
	(*DeleteLiveReq)(nil),  // 5: pb.DeleteLiveReq
	(*DeleteLiveResp)(nil), // 6: pb.DeleteLiveResp
	(*FindLiveReq)(nil),    // 7: pb.FindLiveReq
	(*FindLiveResp)(nil),   // 8: pb.FindLiveResp
}
var file_tiktok_proto_depIdxs = []int32{
	0, // 0: pb.FetchLiveResp.list:type_name -> pb.LiveData
	0, // 1: pb.FindLiveResp.liveDate:type_name -> pb.LiveData
	1, // 2: pb.tiktok.saveLive:input_type -> pb.SaveLiveReq
	3, // 3: pb.tiktok.fetchLives:input_type -> pb.FetchLiveReq
	5, // 4: pb.tiktok.DeleteLiveOne:input_type -> pb.DeleteLiveReq
	7, // 5: pb.tiktok.FindLiveOne:input_type -> pb.FindLiveReq
	2, // 6: pb.tiktok.saveLive:output_type -> pb.SaveLiveResp
	4, // 7: pb.tiktok.fetchLives:output_type -> pb.FetchLiveResp
	6, // 8: pb.tiktok.DeleteLiveOne:output_type -> pb.DeleteLiveResp
	8, // 9: pb.tiktok.FindLiveOne:output_type -> pb.FindLiveResp
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tiktok_proto_init() }
func file_tiktok_proto_init() {
	if File_tiktok_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tiktok_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveData); i {
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
		file_tiktok_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveLiveReq); i {
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
		file_tiktok_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveLiveResp); i {
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
		file_tiktok_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchLiveReq); i {
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
		file_tiktok_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchLiveResp); i {
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
		file_tiktok_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLiveReq); i {
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
		file_tiktok_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLiveResp); i {
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
		file_tiktok_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLiveReq); i {
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
		file_tiktok_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindLiveResp); i {
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
			RawDescriptor: file_tiktok_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tiktok_proto_goTypes,
		DependencyIndexes: file_tiktok_proto_depIdxs,
		MessageInfos:      file_tiktok_proto_msgTypes,
	}.Build()
	File_tiktok_proto = out.File
	file_tiktok_proto_rawDesc = nil
	file_tiktok_proto_goTypes = nil
	file_tiktok_proto_depIdxs = nil
}
