// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: game/player.proto

package game

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

////////////////////////////////////////////////
// player
type LitePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId int64  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Exp       int64  `protobuf:"varint,4,opt,name=exp,proto3" json:"exp,omitempty"`
	Level     int32  `protobuf:"varint,5,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *LitePlayer) Reset() {
	*x = LitePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LitePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LitePlayer) ProtoMessage() {}

func (x *LitePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LitePlayer.ProtoReflect.Descriptor instead.
func (*LitePlayer) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{0}
}

func (x *LitePlayer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LitePlayer) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *LitePlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LitePlayer) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *LitePlayer) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiteInfo *LitePlayer `protobuf:"bytes,1,opt,name=lite_info,json=liteInfo,proto3" json:"lite_info,omitempty"`
	HeroNums int32       `protobuf:"varint,2,opt,name=hero_nums,json=heroNums,proto3" json:"hero_nums,omitempty"`
	ItemNums int32       `protobuf:"varint,3,opt,name=item_nums,json=itemNums,proto3" json:"item_nums,omitempty"`
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerInfo) GetLiteInfo() *LitePlayer {
	if x != nil {
		return x.LiteInfo
	}
	return nil
}

func (x *PlayerInfo) GetHeroNums() int32 {
	if x != nil {
		return x.HeroNums
	}
	return 0
}

func (x *PlayerInfo) GetItemNums() int32 {
	if x != nil {
		return x.ItemNums
	}
	return 0
}

type C2M_CreatePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *C2M_CreatePlayer) Reset() {
	*x = C2M_CreatePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_CreatePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_CreatePlayer) ProtoMessage() {}

func (x *C2M_CreatePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_CreatePlayer.ProtoReflect.Descriptor instead.
func (*C2M_CreatePlayer) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{2}
}

func (x *C2M_CreatePlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type M2C_CreatePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   int32       `protobuf:"varint,91,opt,name=Error,proto3" json:"Error,omitempty"`
	Message string      `protobuf:"bytes,92,opt,name=Message,proto3" json:"Message,omitempty"`
	Info    *PlayerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *M2C_CreatePlayer) Reset() {
	*x = M2C_CreatePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_CreatePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_CreatePlayer) ProtoMessage() {}

func (x *M2C_CreatePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_CreatePlayer.ProtoReflect.Descriptor instead.
func (*M2C_CreatePlayer) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{3}
}

func (x *M2C_CreatePlayer) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *M2C_CreatePlayer) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *M2C_CreatePlayer) GetInfo() *PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type MC_SelectPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MC_SelectPlayer) Reset() {
	*x = MC_SelectPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MC_SelectPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MC_SelectPlayer) ProtoMessage() {}

func (x *MC_SelectPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MC_SelectPlayer.ProtoReflect.Descriptor instead.
func (*MC_SelectPlayer) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{4}
}

func (x *MC_SelectPlayer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MS_SelectPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *PlayerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	ErrorCode int32       `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *MS_SelectPlayer) Reset() {
	*x = MS_SelectPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MS_SelectPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MS_SelectPlayer) ProtoMessage() {}

func (x *MS_SelectPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MS_SelectPlayer.ProtoReflect.Descriptor instead.
func (*MS_SelectPlayer) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{5}
}

func (x *MS_SelectPlayer) GetInfo() *PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *MS_SelectPlayer) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type C2M_QueryPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2M_QueryPlayerInfo) Reset() {
	*x = C2M_QueryPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_QueryPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_QueryPlayerInfo) ProtoMessage() {}

func (x *C2M_QueryPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_QueryPlayerInfo.ProtoReflect.Descriptor instead.
func (*C2M_QueryPlayerInfo) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{6}
}

type M2C_QueryPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *PlayerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Error int32       `protobuf:"varint,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *M2C_QueryPlayerInfo) Reset() {
	*x = M2C_QueryPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_QueryPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_QueryPlayerInfo) ProtoMessage() {}

func (x *M2C_QueryPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_QueryPlayerInfo.ProtoReflect.Descriptor instead.
func (*M2C_QueryPlayerInfo) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{7}
}

func (x *M2C_QueryPlayerInfo) GetInfo() *PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *M2C_QueryPlayerInfo) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

type C2M_ChangeExp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddExp int64 `protobuf:"varint,1,opt,name=AddExp,proto3" json:"AddExp,omitempty"`
}

func (x *C2M_ChangeExp) Reset() {
	*x = C2M_ChangeExp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_ChangeExp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_ChangeExp) ProtoMessage() {}

func (x *C2M_ChangeExp) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_ChangeExp.ProtoReflect.Descriptor instead.
func (*C2M_ChangeExp) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{8}
}

func (x *C2M_ChangeExp) GetAddExp() int64 {
	if x != nil {
		return x.AddExp
	}
	return 0
}

type M2C_ExpUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exp   int64 `protobuf:"varint,1,opt,name=Exp,proto3" json:"Exp,omitempty"`
	Level int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *M2C_ExpUpdate) Reset() {
	*x = M2C_ExpUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_ExpUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_ExpUpdate) ProtoMessage() {}

func (x *M2C_ExpUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_ExpUpdate.ProtoReflect.Descriptor instead.
func (*M2C_ExpUpdate) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{9}
}

func (x *M2C_ExpUpdate) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *M2C_ExpUpdate) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type C2M_ChangeLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddLevel int32 `protobuf:"varint,1,opt,name=AddLevel,proto3" json:"AddLevel,omitempty"`
}

func (x *C2M_ChangeLevel) Reset() {
	*x = C2M_ChangeLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_ChangeLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_ChangeLevel) ProtoMessage() {}

func (x *C2M_ChangeLevel) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_ChangeLevel.ProtoReflect.Descriptor instead.
func (*C2M_ChangeLevel) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{10}
}

func (x *C2M_ChangeLevel) GetAddLevel() int32 {
	if x != nil {
		return x.AddLevel
	}
	return 0
}

type C2M_SyncPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2M_SyncPlayerInfo) Reset() {
	*x = C2M_SyncPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_SyncPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_SyncPlayerInfo) ProtoMessage() {}

func (x *C2M_SyncPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_SyncPlayerInfo.ProtoReflect.Descriptor instead.
func (*C2M_SyncPlayerInfo) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{11}
}

type M2C_SyncPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *M2C_SyncPlayerInfo) Reset() {
	*x = M2C_SyncPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_SyncPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_SyncPlayerInfo) ProtoMessage() {}

func (x *M2C_SyncPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_SyncPlayerInfo.ProtoReflect.Descriptor instead.
func (*M2C_SyncPlayerInfo) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{12}
}

type C2M_PublicSyncPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2M_PublicSyncPlayerInfo) Reset() {
	*x = C2M_PublicSyncPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_PublicSyncPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_PublicSyncPlayerInfo) ProtoMessage() {}

func (x *C2M_PublicSyncPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_PublicSyncPlayerInfo.ProtoReflect.Descriptor instead.
func (*C2M_PublicSyncPlayerInfo) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{13}
}

type M2C_PublicSyncPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *M2C_PublicSyncPlayerInfo) Reset() {
	*x = M2C_PublicSyncPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_PublicSyncPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_PublicSyncPlayerInfo) ProtoMessage() {}

func (x *M2C_PublicSyncPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_PublicSyncPlayerInfo.ProtoReflect.Descriptor instead.
func (*M2C_PublicSyncPlayerInfo) Descriptor() ([]byte, []int) {
	return file_game_player_proto_rawDescGZIP(), []int{14}
}

var File_game_player_proto protoreflect.FileDescriptor

var file_game_player_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x77, 0x0a, 0x0a, 0x4c, 0x69, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0x75, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2d, 0x0a, 0x09, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x08, 0x6c, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x68, 0x65, 0x72, 0x6f, 0x5f, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x68, 0x65, 0x72, 0x6f, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x73, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x32, 0x4d,
	0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x68, 0x0a, 0x10, 0x4d, 0x32, 0x43, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x5b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x5c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x4d,
	0x43, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x56,
	0x0a, 0x0f, 0x4d, 0x53, 0x5f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x32, 0x4d, 0x5f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x51, 0x0a,
	0x13, 0x4d, 0x32, 0x43, 0x5f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x27, 0x0a, 0x0d, 0x43, 0x32, 0x4d, 0x5f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x78,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x45, 0x78, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x41, 0x64, 0x64, 0x45, 0x78, 0x70, 0x22, 0x37, 0x0a, 0x0d, 0x4d, 0x32, 0x43,
	0x5f, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x78,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x45, 0x78, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0x2d, 0x0a, 0x0f, 0x43, 0x32, 0x4d, 0x5f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x41, 0x64, 0x64, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x32, 0x4d, 0x5f, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x4d, 0x32, 0x43, 0x5f, 0x53,
	0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1a, 0x0a,
	0x18, 0x43, 0x32, 0x4d, 0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x79, 0x6e, 0x63, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x4d, 0x32, 0x43,
	0x5f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_player_proto_rawDescOnce sync.Once
	file_game_player_proto_rawDescData = file_game_player_proto_rawDesc
)

func file_game_player_proto_rawDescGZIP() []byte {
	file_game_player_proto_rawDescOnce.Do(func() {
		file_game_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_player_proto_rawDescData)
	})
	return file_game_player_proto_rawDescData
}

var file_game_player_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_game_player_proto_goTypes = []interface{}{
	(*LitePlayer)(nil),               // 0: game.LitePlayer
	(*PlayerInfo)(nil),               // 1: game.PlayerInfo
	(*C2M_CreatePlayer)(nil),         // 2: game.C2M_CreatePlayer
	(*M2C_CreatePlayer)(nil),         // 3: game.M2C_CreatePlayer
	(*MC_SelectPlayer)(nil),          // 4: game.MC_SelectPlayer
	(*MS_SelectPlayer)(nil),          // 5: game.MS_SelectPlayer
	(*C2M_QueryPlayerInfo)(nil),      // 6: game.C2M_QueryPlayerInfo
	(*M2C_QueryPlayerInfo)(nil),      // 7: game.M2C_QueryPlayerInfo
	(*C2M_ChangeExp)(nil),            // 8: game.C2M_ChangeExp
	(*M2C_ExpUpdate)(nil),            // 9: game.M2C_ExpUpdate
	(*C2M_ChangeLevel)(nil),          // 10: game.C2M_ChangeLevel
	(*C2M_SyncPlayerInfo)(nil),       // 11: game.C2M_SyncPlayerInfo
	(*M2C_SyncPlayerInfo)(nil),       // 12: game.M2C_SyncPlayerInfo
	(*C2M_PublicSyncPlayerInfo)(nil), // 13: game.C2M_PublicSyncPlayerInfo
	(*M2C_PublicSyncPlayerInfo)(nil), // 14: game.M2C_PublicSyncPlayerInfo
}
var file_game_player_proto_depIdxs = []int32{
	0, // 0: game.PlayerInfo.lite_info:type_name -> game.LitePlayer
	1, // 1: game.M2C_CreatePlayer.info:type_name -> game.PlayerInfo
	1, // 2: game.MS_SelectPlayer.info:type_name -> game.PlayerInfo
	1, // 3: game.M2C_QueryPlayerInfo.info:type_name -> game.PlayerInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_game_player_proto_init() }
func file_game_player_proto_init() {
	if File_game_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LitePlayer); i {
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
		file_game_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
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
		file_game_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_CreatePlayer); i {
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
		file_game_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_CreatePlayer); i {
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
		file_game_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MC_SelectPlayer); i {
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
		file_game_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MS_SelectPlayer); i {
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
		file_game_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_QueryPlayerInfo); i {
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
		file_game_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_QueryPlayerInfo); i {
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
		file_game_player_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_ChangeExp); i {
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
		file_game_player_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_ExpUpdate); i {
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
		file_game_player_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_ChangeLevel); i {
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
		file_game_player_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_SyncPlayerInfo); i {
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
		file_game_player_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_SyncPlayerInfo); i {
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
		file_game_player_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_PublicSyncPlayerInfo); i {
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
		file_game_player_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_PublicSyncPlayerInfo); i {
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
			RawDescriptor: file_game_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_player_proto_goTypes,
		DependencyIndexes: file_game_player_proto_depIdxs,
		MessageInfos:      file_game_player_proto_msgTypes,
	}.Build()
	File_game_player_proto = out.File
	file_game_player_proto_rawDesc = nil
	file_game_player_proto_goTypes = nil
	file_game_player_proto_depIdxs = nil
}
