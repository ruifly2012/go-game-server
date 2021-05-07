// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: player.proto

package global

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

// 创建角色
type C2S_CreatePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *C2S_CreatePlayer) Reset() {
	*x = C2S_CreatePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CreatePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CreatePlayer) ProtoMessage() {}

func (x *C2S_CreatePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_CreatePlayer.ProtoReflect.Descriptor instead.
func (*C2S_CreatePlayer) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_CreatePlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type S2C_CreatePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *PlayerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *S2C_CreatePlayer) Reset() {
	*x = S2C_CreatePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_CreatePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_CreatePlayer) ProtoMessage() {}

func (x *S2C_CreatePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_CreatePlayer.ProtoReflect.Descriptor instead.
func (*S2C_CreatePlayer) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{1}
}

func (x *S2C_CreatePlayer) GetInfo() *PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// 玩家登陆初始信息
type S2C_PlayerInitInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info            *PlayerInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`                       // 玩家基本信息
	Heros           []*Hero     `protobuf:"bytes,2,rep,name=Heros,proto3" json:"Heros,omitempty"`                     // 英雄数据
	Items           []*Item     `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`                     // 物品数据
	Equips          []*Equip    `protobuf:"bytes,4,rep,name=Equips,proto3" json:"Equips,omitempty"`                   // 装备数据
	Crystals        []*Crystal  `protobuf:"bytes,5,rep,name=Crystals,proto3" json:"Crystals,omitempty"`               // 晶石数据
	HeroFrags       []*Fragment `protobuf:"bytes,6,rep,name=HeroFrags,proto3" json:"HeroFrags,omitempty"`             // 英雄碎片数据
	CollectionFrags []*Fragment `protobuf:"bytes,7,rep,name=CollectionFrags,proto3" json:"CollectionFrags,omitempty"` // 收集品碎片数据
	Chapters        []*Chapter  `protobuf:"bytes,8,rep,name=Chapters,proto3" json:"Chapters,omitempty"`               // 章节信息
	Stages          []*Stage    `protobuf:"bytes,9,rep,name=Stages,proto3" json:"Stages,omitempty"`                   // 关卡信息
	GuideInfo       []uint64    `protobuf:"varint,10,rep,packed,name=GuideInfo,proto3" json:"GuideInfo,omitempty"`    // 引导信息
	Quests          []*Quest    `protobuf:"bytes,11,rep,name=Quests,proto3" json:"Quests,omitempty"`                  // 任务信息
}

func (x *S2C_PlayerInitInfo) Reset() {
	*x = S2C_PlayerInitInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_PlayerInitInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_PlayerInitInfo) ProtoMessage() {}

func (x *S2C_PlayerInitInfo) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_PlayerInitInfo.ProtoReflect.Descriptor instead.
func (*S2C_PlayerInitInfo) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{2}
}

func (x *S2C_PlayerInitInfo) GetInfo() *PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetHeros() []*Hero {
	if x != nil {
		return x.Heros
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetEquips() []*Equip {
	if x != nil {
		return x.Equips
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetCrystals() []*Crystal {
	if x != nil {
		return x.Crystals
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetHeroFrags() []*Fragment {
	if x != nil {
		return x.HeroFrags
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetCollectionFrags() []*Fragment {
	if x != nil {
		return x.CollectionFrags
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetChapters() []*Chapter {
	if x != nil {
		return x.Chapters
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetStages() []*Stage {
	if x != nil {
		return x.Stages
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetGuideInfo() []uint64 {
	if x != nil {
		return x.GuideInfo
	}
	return nil
}

func (x *S2C_PlayerInitInfo) GetQuests() []*Quest {
	if x != nil {
		return x.Quests
	}
	return nil
}

type S2C_ExpUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exp   int32 `protobuf:"varint,1,opt,name=Exp,proto3" json:"Exp,omitempty"`
	Level int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *S2C_ExpUpdate) Reset() {
	*x = S2C_ExpUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_ExpUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_ExpUpdate) ProtoMessage() {}

func (x *S2C_ExpUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_ExpUpdate.ProtoReflect.Descriptor instead.
func (*S2C_ExpUpdate) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{3}
}

func (x *S2C_ExpUpdate) GetExp() int32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *S2C_ExpUpdate) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type S2C_VipUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VipExp   int32 `protobuf:"varint,1,opt,name=VipExp,proto3" json:"VipExp,omitempty"`
	VipLevel int32 `protobuf:"varint,2,opt,name=VipLevel,proto3" json:"VipLevel,omitempty"`
}

func (x *S2C_VipUpdate) Reset() {
	*x = S2C_VipUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_VipUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_VipUpdate) ProtoMessage() {}

func (x *S2C_VipUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_VipUpdate.ProtoReflect.Descriptor instead.
func (*S2C_VipUpdate) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{4}
}

func (x *S2C_VipUpdate) GetVipExp() int32 {
	if x != nil {
		return x.VipExp
	}
	return 0
}

func (x *S2C_VipUpdate) GetVipLevel() int32 {
	if x != nil {
		return x.VipLevel
	}
	return 0
}

////////////////////////////////////////////////
// stage
type C2S_StageChallenge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId int32 `protobuf:"varint,1,opt,name=StageId,proto3" json:"StageId,omitempty"`
}

func (x *C2S_StageChallenge) Reset() {
	*x = C2S_StageChallenge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_StageChallenge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_StageChallenge) ProtoMessage() {}

func (x *C2S_StageChallenge) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_StageChallenge.ProtoReflect.Descriptor instead.
func (*C2S_StageChallenge) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{5}
}

func (x *C2S_StageChallenge) GetStageId() int32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

type S2C_StageChallenge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId int32 `protobuf:"varint,1,opt,name=StageId,proto3" json:"StageId,omitempty"`
	Win     bool  `protobuf:"varint,2,opt,name=win,proto3" json:"win,omitempty"`
}

func (x *S2C_StageChallenge) Reset() {
	*x = S2C_StageChallenge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_StageChallenge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_StageChallenge) ProtoMessage() {}

func (x *S2C_StageChallenge) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_StageChallenge.ProtoReflect.Descriptor instead.
func (*S2C_StageChallenge) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{6}
}

func (x *S2C_StageChallenge) GetStageId() int32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *S2C_StageChallenge) GetWin() bool {
	if x != nil {
		return x.Win
	}
	return false
}

// 扫荡关卡
type C2S_StageSweep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId int32 `protobuf:"varint,1,opt,name=StageId,proto3" json:"StageId,omitempty"`
	Times   int32 `protobuf:"varint,2,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (x *C2S_StageSweep) Reset() {
	*x = C2S_StageSweep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_StageSweep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_StageSweep) ProtoMessage() {}

func (x *C2S_StageSweep) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_StageSweep.ProtoReflect.Descriptor instead.
func (*C2S_StageSweep) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{7}
}

func (x *C2S_StageSweep) GetStageId() int32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *C2S_StageSweep) GetTimes() int32 {
	if x != nil {
		return x.Times
	}
	return 0
}

// 章节信息更新
type S2C_ChapterUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chapter *Chapter `protobuf:"bytes,1,opt,name=Chapter,proto3" json:"Chapter,omitempty"`
}

func (x *S2C_ChapterUpdate) Reset() {
	*x = S2C_ChapterUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_ChapterUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_ChapterUpdate) ProtoMessage() {}

func (x *S2C_ChapterUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_ChapterUpdate.ProtoReflect.Descriptor instead.
func (*S2C_ChapterUpdate) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{8}
}

func (x *S2C_ChapterUpdate) GetChapter() *Chapter {
	if x != nil {
		return x.Chapter
	}
	return nil
}

// 关卡信息更新
type S2C_StageUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stage *Stage `protobuf:"bytes,1,opt,name=Stage,proto3" json:"Stage,omitempty"`
}

func (x *S2C_StageUpdate) Reset() {
	*x = S2C_StageUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_StageUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_StageUpdate) ProtoMessage() {}

func (x *S2C_StageUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_StageUpdate.ProtoReflect.Descriptor instead.
func (*S2C_StageUpdate) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{9}
}

func (x *S2C_StageUpdate) GetStage() *Stage {
	if x != nil {
		return x.Stage
	}
	return nil
}

////////////////////////////////////////////////
// strengthen
// 取回体力
type C2S_WithdrawStrengthen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *C2S_WithdrawStrengthen) Reset() {
	*x = C2S_WithdrawStrengthen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_WithdrawStrengthen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_WithdrawStrengthen) ProtoMessage() {}

func (x *C2S_WithdrawStrengthen) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_WithdrawStrengthen.ProtoReflect.Descriptor instead.
func (*C2S_WithdrawStrengthen) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{10}
}

func (x *C2S_WithdrawStrengthen) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// 购买体力
type C2S_BuyStrengthen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2S_BuyStrengthen) Reset() {
	*x = C2S_BuyStrengthen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_BuyStrengthen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_BuyStrengthen) ProtoMessage() {}

func (x *C2S_BuyStrengthen) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_BuyStrengthen.ProtoReflect.Descriptor instead.
func (*C2S_BuyStrengthen) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{11}
}

////////////////////////////////////////////////
// 引导
type C2S_GuidePass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *C2S_GuidePass) Reset() {
	*x = C2S_GuidePass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_GuidePass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_GuidePass) ProtoMessage() {}

func (x *C2S_GuidePass) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_GuidePass.ProtoReflect.Descriptor instead.
func (*C2S_GuidePass) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{12}
}

func (x *C2S_GuidePass) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

// gm 命令:
// gm player level 10
// gm hero add 1
// gm item add 1 10
type C2S_GmCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *C2S_GmCmd) Reset() {
	*x = C2S_GmCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_GmCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_GmCmd) ProtoMessage() {}

func (x *C2S_GmCmd) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_GmCmd.ProtoReflect.Descriptor instead.
func (*C2S_GmCmd) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{13}
}

func (x *C2S_GmCmd) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

var File_player_proto protoreflect.FileDescriptor

var file_player_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x32, 0x53, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x10, 0x53,
	0x32, 0x43, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x25, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xd3, 0x03, 0x0a, 0x12, 0x53, 0x32, 0x43, 0x5f, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x05, 0x48, 0x65, 0x72, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x72, 0x6f,
	0x52, 0x05, 0x48, 0x65, 0x72, 0x6f, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x06, 0x45, 0x71, 0x75, 0x69, 0x70, 0x73,
	0x12, 0x2a, 0x0a, 0x08, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x73, 0x74,
	0x61, 0x6c, 0x52, 0x08, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x73, 0x12, 0x2d, 0x0a, 0x09,
	0x48, 0x65, 0x72, 0x6f, 0x46, 0x72, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x09, 0x48, 0x65, 0x72, 0x6f, 0x46, 0x72, 0x61, 0x67, 0x73, 0x12, 0x39, 0x0a, 0x0f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x61, 0x67, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x61,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x72, 0x61, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x08, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x75, 0x69, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x47, 0x75, 0x69,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x51, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x37, 0x0a, 0x0d,
	0x53, 0x32, 0x43, 0x5f, 0x45, 0x78, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x45, 0x78, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x45, 0x78, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x43, 0x0a, 0x0d, 0x53, 0x32, 0x43, 0x5f, 0x56, 0x69, 0x70,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x69, 0x70, 0x45, 0x78, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x56, 0x69, 0x70, 0x45, 0x78, 0x70, 0x12, 0x1a,
	0x0a, 0x08, 0x56, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x56, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x32,
	0x53, 0x5f, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x53, 0x32,
	0x43, 0x5f, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x22, 0x40, 0x0a, 0x0e,
	0x43, 0x32, 0x53, 0x5f, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x77, 0x65, 0x65, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x3d,
	0x0a, 0x11, 0x53, 0x32, 0x43, 0x5f, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x52, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x22, 0x35, 0x0a,
	0x0f, 0x53, 0x32, 0x43, 0x5f, 0x53, 0x74, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x05, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x16, 0x43, 0x32, 0x53, 0x5f, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x32, 0x53, 0x5f, 0x42, 0x75, 0x79, 0x53,
	0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x22, 0x25, 0x0a, 0x0d, 0x43, 0x32, 0x53,
	0x5f, 0x47, 0x75, 0x69, 0x64, 0x65, 0x50, 0x61, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x1d, 0x0a, 0x09, 0x43, 0x32, 0x53, 0x5f, 0x47, 0x6d, 0x43, 0x6d, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x42,
	0x33, 0x5a, 0x29, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x66, 0x75, 0x6e, 0x70, 0x6c, 0x75, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0xaa, 0x02, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_proto_rawDescOnce sync.Once
	file_player_proto_rawDescData = file_player_proto_rawDesc
)

func file_player_proto_rawDescGZIP() []byte {
	file_player_proto_rawDescOnce.Do(func() {
		file_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_proto_rawDescData)
	})
	return file_player_proto_rawDescData
}

var file_player_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_player_proto_goTypes = []interface{}{
	(*C2S_CreatePlayer)(nil),       // 0: proto.C2S_CreatePlayer
	(*S2C_CreatePlayer)(nil),       // 1: proto.S2C_CreatePlayer
	(*S2C_PlayerInitInfo)(nil),     // 2: proto.S2C_PlayerInitInfo
	(*S2C_ExpUpdate)(nil),          // 3: proto.S2C_ExpUpdate
	(*S2C_VipUpdate)(nil),          // 4: proto.S2C_VipUpdate
	(*C2S_StageChallenge)(nil),     // 5: proto.C2S_StageChallenge
	(*S2C_StageChallenge)(nil),     // 6: proto.S2C_StageChallenge
	(*C2S_StageSweep)(nil),         // 7: proto.C2S_StageSweep
	(*S2C_ChapterUpdate)(nil),      // 8: proto.S2C_ChapterUpdate
	(*S2C_StageUpdate)(nil),        // 9: proto.S2C_StageUpdate
	(*C2S_WithdrawStrengthen)(nil), // 10: proto.C2S_WithdrawStrengthen
	(*C2S_BuyStrengthen)(nil),      // 11: proto.C2S_BuyStrengthen
	(*C2S_GuidePass)(nil),          // 12: proto.C2S_GuidePass
	(*C2S_GmCmd)(nil),              // 13: proto.C2S_GmCmd
	(*PlayerInfo)(nil),             // 14: proto.PlayerInfo
	(*Hero)(nil),                   // 15: proto.Hero
	(*Item)(nil),                   // 16: proto.Item
	(*Equip)(nil),                  // 17: proto.Equip
	(*Crystal)(nil),                // 18: proto.Crystal
	(*Fragment)(nil),               // 19: proto.Fragment
	(*Chapter)(nil),                // 20: proto.Chapter
	(*Stage)(nil),                  // 21: proto.Stage
	(*Quest)(nil),                  // 22: proto.Quest
}
var file_player_proto_depIdxs = []int32{
	14, // 0: proto.S2C_CreatePlayer.info:type_name -> proto.PlayerInfo
	14, // 1: proto.S2C_PlayerInitInfo.Info:type_name -> proto.PlayerInfo
	15, // 2: proto.S2C_PlayerInitInfo.Heros:type_name -> proto.Hero
	16, // 3: proto.S2C_PlayerInitInfo.Items:type_name -> proto.Item
	17, // 4: proto.S2C_PlayerInitInfo.Equips:type_name -> proto.Equip
	18, // 5: proto.S2C_PlayerInitInfo.Crystals:type_name -> proto.Crystal
	19, // 6: proto.S2C_PlayerInitInfo.HeroFrags:type_name -> proto.Fragment
	19, // 7: proto.S2C_PlayerInitInfo.CollectionFrags:type_name -> proto.Fragment
	20, // 8: proto.S2C_PlayerInitInfo.Chapters:type_name -> proto.Chapter
	21, // 9: proto.S2C_PlayerInitInfo.Stages:type_name -> proto.Stage
	22, // 10: proto.S2C_PlayerInitInfo.Quests:type_name -> proto.Quest
	20, // 11: proto.S2C_ChapterUpdate.Chapter:type_name -> proto.Chapter
	21, // 12: proto.S2C_StageUpdate.Stage:type_name -> proto.Stage
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_player_proto_init() }
func file_player_proto_init() {
	if File_player_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_CreatePlayer); i {
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
		file_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_CreatePlayer); i {
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
		file_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_PlayerInitInfo); i {
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
		file_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_ExpUpdate); i {
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
		file_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_VipUpdate); i {
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
		file_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_StageChallenge); i {
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
		file_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_StageChallenge); i {
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
		file_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_StageSweep); i {
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
		file_player_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_ChapterUpdate); i {
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
		file_player_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_StageUpdate); i {
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
		file_player_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_WithdrawStrengthen); i {
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
		file_player_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_BuyStrengthen); i {
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
		file_player_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_GuidePass); i {
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
		file_player_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_GmCmd); i {
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
			RawDescriptor: file_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_player_proto_goTypes,
		DependencyIndexes: file_player_proto_depIdxs,
		MessageInfos:      file_player_proto_msgTypes,
	}.Build()
	File_player_proto = out.File
	file_player_proto_rawDesc = nil
	file_player_proto_goTypes = nil
	file_player_proto_depIdxs = nil
}
