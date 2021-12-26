// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: quest.proto

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

// 领取玩家任务奖励
type C2S_PlayerQuestReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"` // 任务id
}

func (x *C2S_PlayerQuestReward) Reset() {
	*x = C2S_PlayerQuestReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_PlayerQuestReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_PlayerQuestReward) ProtoMessage() {}

func (x *C2S_PlayerQuestReward) ProtoReflect() protoreflect.Message {
	mi := &file_quest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_PlayerQuestReward.ProtoReflect.Descriptor instead.
func (*C2S_PlayerQuestReward) Descriptor() ([]byte, []int) {
	return file_quest_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_PlayerQuestReward) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 任务更新
type S2C_QuestUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quest *Quest `protobuf:"bytes,1,opt,name=Quest,proto3" json:"Quest,omitempty"`
}

func (x *S2C_QuestUpdate) Reset() {
	*x = S2C_QuestUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_QuestUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_QuestUpdate) ProtoMessage() {}

func (x *S2C_QuestUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_quest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_QuestUpdate.ProtoReflect.Descriptor instead.
func (*S2C_QuestUpdate) Descriptor() ([]byte, []int) {
	return file_quest_proto_rawDescGZIP(), []int{1}
}

func (x *S2C_QuestUpdate) GetQuest() *Quest {
	if x != nil {
		return x.Quest
	}
	return nil
}

var File_quest_proto protoreflect.FileDescriptor

var file_quest_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x32, 0x53, 0x5f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x53,
	0x32, 0x43, 0x5f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x32, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0xaa, 0x02,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quest_proto_rawDescOnce sync.Once
	file_quest_proto_rawDescData = file_quest_proto_rawDesc
)

func file_quest_proto_rawDescGZIP() []byte {
	file_quest_proto_rawDescOnce.Do(func() {
		file_quest_proto_rawDescData = protoimpl.X.CompressGZIP(file_quest_proto_rawDescData)
	})
	return file_quest_proto_rawDescData
}

var file_quest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_quest_proto_goTypes = []interface{}{
	(*C2S_PlayerQuestReward)(nil), // 0: proto.C2S_PlayerQuestReward
	(*S2C_QuestUpdate)(nil),       // 1: proto.S2C_QuestUpdate
	(*Quest)(nil),                 // 2: proto.Quest
}
var file_quest_proto_depIdxs = []int32{
	2, // 0: proto.S2C_QuestUpdate.Quest:type_name -> proto.Quest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_quest_proto_init() }
func file_quest_proto_init() {
	if File_quest_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_quest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_PlayerQuestReward); i {
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
		file_quest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_QuestUpdate); i {
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
			RawDescriptor: file_quest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_quest_proto_goTypes,
		DependencyIndexes: file_quest_proto_depIdxs,
		MessageInfos:      file_quest_proto_msgTypes,
	}.Build()
	File_quest_proto = out.File
	file_quest_proto_rawDesc = nil
	file_quest_proto_goTypes = nil
	file_quest_proto_depIdxs = nil
}
