// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: global/rune.proto

package global

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

type Rune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64      `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	TypeId     int32      `protobuf:"varint,2,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
	EquipObjId int64      `protobuf:"varint,3,opt,name=EquipObjId,proto3" json:"EquipObjId,omitempty"` // 装备者objid
	Atts       []*RuneAtt `protobuf:"bytes,4,rep,name=Atts,proto3" json:"Atts,omitempty"`              // 1条主属性+5条副属性
}

func (x *Rune) Reset() {
	*x = Rune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rune) ProtoMessage() {}

func (x *Rune) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rune.ProtoReflect.Descriptor instead.
func (*Rune) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{0}
}

func (x *Rune) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rune) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *Rune) GetEquipObjId() int64 {
	if x != nil {
		return x.EquipObjId
	}
	return 0
}

func (x *Rune) GetAtts() []*RuneAtt {
	if x != nil {
		return x.Atts
	}
	return nil
}

type C2S_AddRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId int32 `protobuf:"varint,1,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
}

func (x *C2S_AddRune) Reset() {
	*x = C2S_AddRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_AddRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_AddRune) ProtoMessage() {}

func (x *C2S_AddRune) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_AddRune.ProtoReflect.Descriptor instead.
func (*C2S_AddRune) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{1}
}

func (x *C2S_AddRune) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type C2S_DelRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *C2S_DelRune) Reset() {
	*x = C2S_DelRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_DelRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_DelRune) ProtoMessage() {}

func (x *C2S_DelRune) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_DelRune.ProtoReflect.Descriptor instead.
func (*C2S_DelRune) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{2}
}

func (x *C2S_DelRune) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type C2S_QueryRunes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2S_QueryRunes) Reset() {
	*x = C2S_QueryRunes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_QueryRunes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_QueryRunes) ProtoMessage() {}

func (x *C2S_QueryRunes) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_QueryRunes.ProtoReflect.Descriptor instead.
func (*C2S_QueryRunes) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{3}
}

type S2C_RuneList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runes []*Rune `protobuf:"bytes,1,rep,name=runes,proto3" json:"runes,omitempty"`
}

func (x *S2C_RuneList) Reset() {
	*x = S2C_RuneList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_RuneList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_RuneList) ProtoMessage() {}

func (x *S2C_RuneList) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_RuneList.ProtoReflect.Descriptor instead.
func (*S2C_RuneList) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{4}
}

func (x *S2C_RuneList) GetRunes() []*Rune {
	if x != nil {
		return x.Runes
	}
	return nil
}

type S2C_RuneAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rune *Rune `protobuf:"bytes,1,opt,name=rune,proto3" json:"rune,omitempty"`
}

func (x *S2C_RuneAdd) Reset() {
	*x = S2C_RuneAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_RuneAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_RuneAdd) ProtoMessage() {}

func (x *S2C_RuneAdd) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_RuneAdd.ProtoReflect.Descriptor instead.
func (*S2C_RuneAdd) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{5}
}

func (x *S2C_RuneAdd) GetRune() *Rune {
	if x != nil {
		return x.Rune
	}
	return nil
}

type S2C_RuneUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rune *Rune `protobuf:"bytes,1,opt,name=rune,proto3" json:"rune,omitempty"`
}

func (x *S2C_RuneUpdate) Reset() {
	*x = S2C_RuneUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_RuneUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_RuneUpdate) ProtoMessage() {}

func (x *S2C_RuneUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_RuneUpdate.ProtoReflect.Descriptor instead.
func (*S2C_RuneUpdate) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{6}
}

func (x *S2C_RuneUpdate) GetRune() *Rune {
	if x != nil {
		return x.Rune
	}
	return nil
}

type S2C_DelRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuneId int64 `protobuf:"varint,1,opt,name=RuneId,proto3" json:"RuneId,omitempty"`
}

func (x *S2C_DelRune) Reset() {
	*x = S2C_DelRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_DelRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_DelRune) ProtoMessage() {}

func (x *S2C_DelRune) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_DelRune.ProtoReflect.Descriptor instead.
func (*S2C_DelRune) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{7}
}

func (x *S2C_DelRune) GetRuneId() int64 {
	if x != nil {
		return x.RuneId
	}
	return 0
}

type C2S_PutonRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	RuneId int64 `protobuf:"varint,2,opt,name=RuneId,proto3" json:"RuneId,omitempty"`
}

func (x *C2S_PutonRune) Reset() {
	*x = C2S_PutonRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_PutonRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_PutonRune) ProtoMessage() {}

func (x *C2S_PutonRune) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_PutonRune.ProtoReflect.Descriptor instead.
func (*C2S_PutonRune) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{8}
}

func (x *C2S_PutonRune) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2S_PutonRune) GetRuneId() int64 {
	if x != nil {
		return x.RuneId
	}
	return 0
}

type C2S_TakeoffRune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroId int64 `protobuf:"varint,1,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	Pos    int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (x *C2S_TakeoffRune) Reset() {
	*x = C2S_TakeoffRune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_rune_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_TakeoffRune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_TakeoffRune) ProtoMessage() {}

func (x *C2S_TakeoffRune) ProtoReflect() protoreflect.Message {
	mi := &file_global_rune_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_TakeoffRune.ProtoReflect.Descriptor instead.
func (*C2S_TakeoffRune) Descriptor() ([]byte, []int) {
	return file_global_rune_proto_rawDescGZIP(), []int{9}
}

func (x *C2S_TakeoffRune) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

func (x *C2S_TakeoffRune) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

var File_global_rune_proto protoreflect.FileDescriptor

var file_global_rune_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x72, 0x75, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x1a, 0x13, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x73, 0x0a, 0x04, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4f, 0x62, 0x6a, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x04, 0x41, 0x74, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x52, 0x75, 0x6e, 0x65, 0x41, 0x74, 0x74, 0x52,
	0x04, 0x41, 0x74, 0x74, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x43, 0x32, 0x53, 0x5f, 0x41, 0x64, 0x64,
	0x52, 0x75, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x0b,
	0x43, 0x32, 0x53, 0x5f, 0x44, 0x65, 0x6c, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x43,
	0x32, 0x53, 0x5f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x65, 0x73, 0x22, 0x32, 0x0a,
	0x0c, 0x53, 0x32, 0x43, 0x5f, 0x52, 0x75, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x05, 0x72, 0x75, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x65,
	0x73, 0x22, 0x2f, 0x0a, 0x0b, 0x53, 0x32, 0x43, 0x5f, 0x52, 0x75, 0x6e, 0x65, 0x41, 0x64, 0x64,
	0x12, 0x20, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x04, 0x72, 0x75,
	0x6e, 0x65, 0x22, 0x32, 0x0a, 0x0e, 0x53, 0x32, 0x43, 0x5f, 0x52, 0x75, 0x6e, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x52, 0x75, 0x6e, 0x65,
	0x52, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x32, 0x43, 0x5f, 0x44, 0x65,
	0x6c, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x3f, 0x0a,
	0x0d, 0x43, 0x32, 0x53, 0x5f, 0x50, 0x75, 0x74, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x3b,
	0x0a, 0x0f, 0x43, 0x32, 0x53, 0x5f, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x66, 0x66, 0x52, 0x75, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x6f, 0x73, 0x42, 0x34, 0x5a, 0x29, 0x62,
	0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x75, 0x6e,
	0x70, 0x6c, 0x75, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0xaa, 0x02, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_global_rune_proto_rawDescOnce sync.Once
	file_global_rune_proto_rawDescData = file_global_rune_proto_rawDesc
)

func file_global_rune_proto_rawDescGZIP() []byte {
	file_global_rune_proto_rawDescOnce.Do(func() {
		file_global_rune_proto_rawDescData = protoimpl.X.CompressGZIP(file_global_rune_proto_rawDescData)
	})
	return file_global_rune_proto_rawDescData
}

var file_global_rune_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_global_rune_proto_goTypes = []interface{}{
	(*Rune)(nil),            // 0: global.Rune
	(*C2S_AddRune)(nil),     // 1: global.C2S_AddRune
	(*C2S_DelRune)(nil),     // 2: global.C2S_DelRune
	(*C2S_QueryRunes)(nil),  // 3: global.C2S_QueryRunes
	(*S2C_RuneList)(nil),    // 4: global.S2C_RuneList
	(*S2C_RuneAdd)(nil),     // 5: global.S2C_RuneAdd
	(*S2C_RuneUpdate)(nil),  // 6: global.S2C_RuneUpdate
	(*S2C_DelRune)(nil),     // 7: global.S2C_DelRune
	(*C2S_PutonRune)(nil),   // 8: global.C2S_PutonRune
	(*C2S_TakeoffRune)(nil), // 9: global.C2S_TakeoffRune
	(*RuneAtt)(nil),         // 10: global.RuneAtt
}
var file_global_rune_proto_depIdxs = []int32{
	10, // 0: global.Rune.Atts:type_name -> global.RuneAtt
	0,  // 1: global.S2C_RuneList.runes:type_name -> global.Rune
	0,  // 2: global.S2C_RuneAdd.rune:type_name -> global.Rune
	0,  // 3: global.S2C_RuneUpdate.rune:type_name -> global.Rune
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_global_rune_proto_init() }
func file_global_rune_proto_init() {
	if File_global_rune_proto != nil {
		return
	}
	file_global_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_global_rune_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rune); i {
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
		file_global_rune_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_AddRune); i {
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
		file_global_rune_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_DelRune); i {
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
		file_global_rune_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_QueryRunes); i {
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
		file_global_rune_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_RuneList); i {
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
		file_global_rune_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_RuneAdd); i {
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
		file_global_rune_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_RuneUpdate); i {
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
		file_global_rune_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_DelRune); i {
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
		file_global_rune_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_PutonRune); i {
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
		file_global_rune_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_TakeoffRune); i {
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
			RawDescriptor: file_global_rune_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_global_rune_proto_goTypes,
		DependencyIndexes: file_global_rune_proto_depIdxs,
		MessageInfos:      file_global_rune_proto_msgTypes,
	}.Build()
	File_global_rune_proto = out.File
	file_global_rune_proto_rawDesc = nil
	file_global_rune_proto_goTypes = nil
	file_global_rune_proto_depIdxs = nil
}
