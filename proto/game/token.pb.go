// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game/token.proto

package game

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

////////////////////////////////////////////////
// Token
type Token struct {
	Type                 int32    `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	MaxHold              int64    `protobuf:"varint,3,opt,name=MaxHold,proto3" json:"MaxHold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_036e6e92ff2ee4d0, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Token) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Token) GetMaxHold() int64 {
	if m != nil {
		return m.MaxHold
	}
	return 0
}

type C2M_AddToken struct {
	Type                 int32    `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_AddToken) Reset()         { *m = C2M_AddToken{} }
func (m *C2M_AddToken) String() string { return proto.CompactTextString(m) }
func (*C2M_AddToken) ProtoMessage()    {}
func (*C2M_AddToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_036e6e92ff2ee4d0, []int{1}
}

func (m *C2M_AddToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_AddToken.Unmarshal(m, b)
}
func (m *C2M_AddToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_AddToken.Marshal(b, m, deterministic)
}
func (m *C2M_AddToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_AddToken.Merge(m, src)
}
func (m *C2M_AddToken) XXX_Size() int {
	return xxx_messageInfo_C2M_AddToken.Size(m)
}
func (m *C2M_AddToken) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_AddToken.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_AddToken proto.InternalMessageInfo

func (m *C2M_AddToken) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *C2M_AddToken) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type C2M_QueryTokens struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_QueryTokens) Reset()         { *m = C2M_QueryTokens{} }
func (m *C2M_QueryTokens) String() string { return proto.CompactTextString(m) }
func (*C2M_QueryTokens) ProtoMessage()    {}
func (*C2M_QueryTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_036e6e92ff2ee4d0, []int{2}
}

func (m *C2M_QueryTokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_QueryTokens.Unmarshal(m, b)
}
func (m *C2M_QueryTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_QueryTokens.Marshal(b, m, deterministic)
}
func (m *C2M_QueryTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_QueryTokens.Merge(m, src)
}
func (m *C2M_QueryTokens) XXX_Size() int {
	return xxx_messageInfo_C2M_QueryTokens.Size(m)
}
func (m *C2M_QueryTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_QueryTokens.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_QueryTokens proto.InternalMessageInfo

type M2C_TokenList struct {
	Tokens               []*Token `protobuf:"bytes,1,rep,name=Tokens,proto3" json:"Tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_TokenList) Reset()         { *m = M2C_TokenList{} }
func (m *M2C_TokenList) String() string { return proto.CompactTextString(m) }
func (*M2C_TokenList) ProtoMessage()    {}
func (*M2C_TokenList) Descriptor() ([]byte, []int) {
	return fileDescriptor_036e6e92ff2ee4d0, []int{3}
}

func (m *M2C_TokenList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_TokenList.Unmarshal(m, b)
}
func (m *M2C_TokenList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_TokenList.Marshal(b, m, deterministic)
}
func (m *M2C_TokenList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_TokenList.Merge(m, src)
}
func (m *M2C_TokenList) XXX_Size() int {
	return xxx_messageInfo_M2C_TokenList.Size(m)
}
func (m *M2C_TokenList) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_TokenList.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_TokenList proto.InternalMessageInfo

func (m *M2C_TokenList) GetTokens() []*Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type M2C_TokenUpdate struct {
	Info                 *Token   `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_TokenUpdate) Reset()         { *m = M2C_TokenUpdate{} }
func (m *M2C_TokenUpdate) String() string { return proto.CompactTextString(m) }
func (*M2C_TokenUpdate) ProtoMessage()    {}
func (*M2C_TokenUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_036e6e92ff2ee4d0, []int{4}
}

func (m *M2C_TokenUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_TokenUpdate.Unmarshal(m, b)
}
func (m *M2C_TokenUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_TokenUpdate.Marshal(b, m, deterministic)
}
func (m *M2C_TokenUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_TokenUpdate.Merge(m, src)
}
func (m *M2C_TokenUpdate) XXX_Size() int {
	return xxx_messageInfo_M2C_TokenUpdate.Size(m)
}
func (m *M2C_TokenUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_TokenUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_TokenUpdate proto.InternalMessageInfo

func (m *M2C_TokenUpdate) GetInfo() *Token {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "game.Token")
	proto.RegisterType((*C2M_AddToken)(nil), "game.C2M_AddToken")
	proto.RegisterType((*C2M_QueryTokens)(nil), "game.C2M_QueryTokens")
	proto.RegisterType((*M2C_TokenList)(nil), "game.M2C_TokenList")
	proto.RegisterType((*M2C_TokenUpdate)(nil), "game.M2C_TokenUpdate")
}

func init() { proto.RegisterFile("game/token.proto", fileDescriptor_036e6e92ff2ee4d0) }

var fileDescriptor_036e6e92ff2ee4d0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89,
	0x28, 0x79, 0x73, 0xb1, 0x86, 0x80, 0x04, 0x85, 0x84, 0xb8, 0x58, 0x42, 0x2a, 0x0b, 0x52, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xb0, 0xc4, 0x9c, 0xd2,
	0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0xdd, 0x37, 0xb1,
	0xc2, 0x23, 0x3f, 0x27, 0x45, 0x82, 0x19, 0x2c, 0x0e, 0xe3, 0x2a, 0x59, 0x70, 0xf1, 0x38, 0x1b,
	0xf9, 0xc6, 0x3b, 0xa6, 0xa4, 0x90, 0x68, 0xa6, 0x92, 0x20, 0x17, 0x3f, 0x48, 0x67, 0x60, 0x69,
	0x6a, 0x51, 0x25, 0x58, 0x6f, 0xb1, 0x92, 0x09, 0x17, 0xaf, 0xaf, 0x91, 0x73, 0x3c, 0x98, 0xe7,
	0x93, 0x59, 0x5c, 0x22, 0xa4, 0xcc, 0xc5, 0x06, 0x91, 0x92, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36,
	0xe2, 0xd6, 0x03, 0xf9, 0x40, 0x0f, 0x2c, 0x16, 0x04, 0x95, 0x52, 0x32, 0xe2, 0xe2, 0x87, 0xeb,
	0x0a, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x15, 0x92, 0xe7, 0x62, 0xf1, 0xcc, 0x4b, 0xcb, 0x07, 0xbb,
	0x02, 0x4d, 0x17, 0x58, 0xc2, 0x49, 0x23, 0x4a, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0x3f, 0x35, 0xb1, 0xb8, 0x44, 0x37, 0x35, 0x25, 0x35, 0x4f, 0xbf, 0x38, 0xb5,
	0xa8, 0x2c, 0xb5, 0x48, 0x1f, 0x1c, 0x5c, 0xfa, 0x20, 0x5d, 0x49, 0x6c, 0x60, 0xb6, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x15, 0x0b, 0xe1, 0x94, 0x4e, 0x01, 0x00, 0x00,
}
