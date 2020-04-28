// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game/att.proto

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
// Att
type Att struct {
	AttType              int32    `protobuf:"varint,1,opt,name=AttType,proto3" json:"AttType,omitempty"`
	AttValue             int64    `protobuf:"varint,2,opt,name=AttValue,proto3" json:"AttValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Att) Reset()         { *m = Att{} }
func (m *Att) String() string { return proto.CompactTextString(m) }
func (*Att) ProtoMessage()    {}
func (*Att) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b8149fbf8c0a13, []int{0}
}

func (m *Att) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Att.Unmarshal(m, b)
}
func (m *Att) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Att.Marshal(b, m, deterministic)
}
func (m *Att) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Att.Merge(m, src)
}
func (m *Att) XXX_Size() int {
	return xxx_messageInfo_Att.Size(m)
}
func (m *Att) XXX_DiscardUnknown() {
	xxx_messageInfo_Att.DiscardUnknown(m)
}

var xxx_messageInfo_Att proto.InternalMessageInfo

func (m *Att) GetAttType() int32 {
	if m != nil {
		return m.AttType
	}
	return 0
}

func (m *Att) GetAttValue() int64 {
	if m != nil {
		return m.AttValue
	}
	return 0
}

func init() {
	proto.RegisterType((*Att)(nil), "yokai_game.Att")
}

func init() { proto.RegisterFile("game/att.proto", fileDescriptor_38b8149fbf8c0a13) }

var fileDescriptor_38b8149fbf8c0a13 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4f, 0xcc, 0x4d,
	0xd5, 0x4f, 0x2c, 0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xaa, 0xcc, 0xcf, 0x4e,
	0xcc, 0x8c, 0x07, 0x89, 0x2a, 0x59, 0x73, 0x31, 0x3b, 0x96, 0x94, 0x08, 0x49, 0x70, 0xb1, 0x3b,
	0x96, 0x94, 0x84, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0xc1, 0xb8, 0x42,
	0x52, 0x5c, 0x1c, 0x8e, 0x25, 0x25, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0xcc, 0x41, 0x70, 0xbe, 0x93, 0x4e, 0x94, 0x56, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x3e, 0xd8, 0xd4, 0xcc, 0x7c, 0x08, 0x1d, 0x5f, 0x9c, 0x5a, 0x54, 0x96, 0x5a, 0xa4,
	0x0f, 0xb6, 0x51, 0x1f, 0x64, 0x55, 0x12, 0x1b, 0x98, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x6f, 0x8e, 0x0c, 0xaf, 0x8f, 0x00, 0x00, 0x00,
}
