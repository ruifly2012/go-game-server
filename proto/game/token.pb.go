// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: game/token.proto

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
// Token
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Value   int64 `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	MaxHold int64 `protobuf:"varint,3,opt,name=MaxHold,proto3" json:"MaxHold,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_game_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_game_token_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Token) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Token) GetMaxHold() int64 {
	if x != nil {
		return x.MaxHold
	}
	return 0
}

type C2M_AddToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Value int64 `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *C2M_AddToken) Reset() {
	*x = C2M_AddToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_AddToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_AddToken) ProtoMessage() {}

func (x *C2M_AddToken) ProtoReflect() protoreflect.Message {
	mi := &file_game_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_AddToken.ProtoReflect.Descriptor instead.
func (*C2M_AddToken) Descriptor() ([]byte, []int) {
	return file_game_token_proto_rawDescGZIP(), []int{1}
}

func (x *C2M_AddToken) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *C2M_AddToken) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type C2M_QueryTokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2M_QueryTokens) Reset() {
	*x = C2M_QueryTokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2M_QueryTokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2M_QueryTokens) ProtoMessage() {}

func (x *C2M_QueryTokens) ProtoReflect() protoreflect.Message {
	mi := &file_game_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2M_QueryTokens.ProtoReflect.Descriptor instead.
func (*C2M_QueryTokens) Descriptor() ([]byte, []int) {
	return file_game_token_proto_rawDescGZIP(), []int{2}
}

type M2C_TokenList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*Token `protobuf:"bytes,1,rep,name=Tokens,proto3" json:"Tokens,omitempty"`
}

func (x *M2C_TokenList) Reset() {
	*x = M2C_TokenList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_TokenList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_TokenList) ProtoMessage() {}

func (x *M2C_TokenList) ProtoReflect() protoreflect.Message {
	mi := &file_game_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_TokenList.ProtoReflect.Descriptor instead.
func (*M2C_TokenList) Descriptor() ([]byte, []int) {
	return file_game_token_proto_rawDescGZIP(), []int{3}
}

func (x *M2C_TokenList) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type M2C_TokenUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Token `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *M2C_TokenUpdate) Reset() {
	*x = M2C_TokenUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2C_TokenUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2C_TokenUpdate) ProtoMessage() {}

func (x *M2C_TokenUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_game_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2C_TokenUpdate.ProtoReflect.Descriptor instead.
func (*M2C_TokenUpdate) Descriptor() ([]byte, []int) {
	return file_game_token_proto_rawDescGZIP(), []int{4}
}

func (x *M2C_TokenUpdate) GetInfo() *Token {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_game_token_proto protoreflect.FileDescriptor

var file_game_token_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x61, 0x78, 0x48, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x61,
	0x78, 0x48, 0x6f, 0x6c, 0x64, 0x22, 0x38, 0x0a, 0x0c, 0x43, 0x32, 0x4d, 0x5f, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x11, 0x0a, 0x0f, 0x43, 0x32, 0x4d, 0x5f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x22, 0x34, 0x0a, 0x0d, 0x4d, 0x32, 0x43, 0x5f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x32, 0x0a, 0x0f, 0x4d, 0x32, 0x43, 0x5f,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2b, 0x5a, 0x29,
	0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x65, 0x61,
	0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_game_token_proto_rawDescOnce sync.Once
	file_game_token_proto_rawDescData = file_game_token_proto_rawDesc
)

func file_game_token_proto_rawDescGZIP() []byte {
	file_game_token_proto_rawDescOnce.Do(func() {
		file_game_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_token_proto_rawDescData)
	})
	return file_game_token_proto_rawDescData
}

var file_game_token_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_game_token_proto_goTypes = []interface{}{
	(*Token)(nil),           // 0: game.Token
	(*C2M_AddToken)(nil),    // 1: game.C2M_AddToken
	(*C2M_QueryTokens)(nil), // 2: game.C2M_QueryTokens
	(*M2C_TokenList)(nil),   // 3: game.M2C_TokenList
	(*M2C_TokenUpdate)(nil), // 4: game.M2C_TokenUpdate
}
var file_game_token_proto_depIdxs = []int32{
	0, // 0: game.M2C_TokenList.Tokens:type_name -> game.Token
	0, // 1: game.M2C_TokenUpdate.Info:type_name -> game.Token
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_game_token_proto_init() }
func file_game_token_proto_init() {
	if File_game_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_game_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_AddToken); i {
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
		file_game_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2M_QueryTokens); i {
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
		file_game_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_TokenList); i {
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
		file_game_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2C_TokenUpdate); i {
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
			RawDescriptor: file_game_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_token_proto_goTypes,
		DependencyIndexes: file_game_token_proto_depIdxs,
		MessageInfos:      file_game_token_proto_msgTypes,
	}.Build()
	File_game_token_proto = out.File
	file_game_token_proto_rawDesc = nil
	file_game_token_proto_goTypes = nil
	file_game_token_proto_depIdxs = nil
}
