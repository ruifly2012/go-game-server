// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pubsub/pubsub.proto

package pubsub

import (
	account "bitbucket.org/east-eden/server/proto/account"
	game "bitbucket.org/east-eden/server/proto/game"
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

/////////////////////////////////////////////////
// pub/sub
/////////////////////////////////////////////////
type PubStartGate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *account.LiteAccount `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PubStartGate) Reset() {
	*x = PubStartGate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubStartGate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubStartGate) ProtoMessage() {}

func (x *PubStartGate) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubStartGate.ProtoReflect.Descriptor instead.
func (*PubStartGate) Descriptor() ([]byte, []int) {
	return file_pubsub_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *PubStartGate) GetInfo() *account.LiteAccount {
	if x != nil {
		return x.Info
	}
	return nil
}

type PubGateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *account.LiteAccount `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Win  bool                 `protobuf:"varint,2,opt,name=win,proto3" json:"win,omitempty"`
}

func (x *PubGateResult) Reset() {
	*x = PubGateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_pubsub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubGateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubGateResult) ProtoMessage() {}

func (x *PubGateResult) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_pubsub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubGateResult.ProtoReflect.Descriptor instead.
func (*PubGateResult) Descriptor() ([]byte, []int) {
	return file_pubsub_pubsub_proto_rawDescGZIP(), []int{1}
}

func (x *PubGateResult) GetInfo() *account.LiteAccount {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *PubGateResult) GetWin() bool {
	if x != nil {
		return x.Win
	}
	return false
}

type PubSyncPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *game.PlayerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PubSyncPlayerInfo) Reset() {
	*x = PubSyncPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_pubsub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubSyncPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubSyncPlayerInfo) ProtoMessage() {}

func (x *PubSyncPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_pubsub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubSyncPlayerInfo.ProtoReflect.Descriptor instead.
func (*PubSyncPlayerInfo) Descriptor() ([]byte, []int) {
	return file_pubsub_pubsub_proto_rawDescGZIP(), []int{2}
}

func (x *PubSyncPlayerInfo) GetInfo() *game.PlayerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_pubsub_pubsub_proto protoreflect.FileDescriptor

var file_pubsub_pubsub_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x1a, 0x15, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x47, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x4c, 0x69, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x4b, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x47, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x77, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x77, 0x69, 0x6e, 0x22, 0x39,
	0x0a, 0x11, 0x50, 0x75, 0x62, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x2d, 0x5a, 0x2b, 0x62, 0x69, 0x74,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d,
	0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pubsub_pubsub_proto_rawDescOnce sync.Once
	file_pubsub_pubsub_proto_rawDescData = file_pubsub_pubsub_proto_rawDesc
)

func file_pubsub_pubsub_proto_rawDescGZIP() []byte {
	file_pubsub_pubsub_proto_rawDescOnce.Do(func() {
		file_pubsub_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_pubsub_pubsub_proto_rawDescData)
	})
	return file_pubsub_pubsub_proto_rawDescData
}

var file_pubsub_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pubsub_pubsub_proto_goTypes = []interface{}{
	(*PubStartGate)(nil),        // 0: pubsub.PubStartGate
	(*PubGateResult)(nil),       // 1: pubsub.PubGateResult
	(*PubSyncPlayerInfo)(nil),   // 2: pubsub.PubSyncPlayerInfo
	(*account.LiteAccount)(nil), // 3: account.LiteAccount
	(*game.PlayerInfo)(nil),     // 4: game.PlayerInfo
}
var file_pubsub_pubsub_proto_depIdxs = []int32{
	3, // 0: pubsub.PubStartGate.info:type_name -> account.LiteAccount
	3, // 1: pubsub.PubGateResult.info:type_name -> account.LiteAccount
	4, // 2: pubsub.PubSyncPlayerInfo.info:type_name -> game.PlayerInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pubsub_pubsub_proto_init() }
func file_pubsub_pubsub_proto_init() {
	if File_pubsub_pubsub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pubsub_pubsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubStartGate); i {
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
		file_pubsub_pubsub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubGateResult); i {
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
		file_pubsub_pubsub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubSyncPlayerInfo); i {
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
			RawDescriptor: file_pubsub_pubsub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pubsub_pubsub_proto_goTypes,
		DependencyIndexes: file_pubsub_pubsub_proto_depIdxs,
		MessageInfos:      file_pubsub_pubsub_proto_msgTypes,
	}.Build()
	File_pubsub_pubsub_proto = out.File
	file_pubsub_pubsub_proto_rawDesc = nil
	file_pubsub_pubsub_proto_goTypes = nil
	file_pubsub_pubsub_proto_depIdxs = nil
}
