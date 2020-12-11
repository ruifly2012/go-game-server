// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game/item.proto

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
// Item
type Item struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	TypeId               int32    `protobuf:"varint,2,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
	Num                  int32    `protobuf:"varint,3,opt,name=Num,proto3" json:"Num,omitempty"`
	EquipObjId           int64    `protobuf:"varint,4,opt,name=EquipObjId,proto3" json:"EquipObjId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_20b3be662c2ee0bd, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Item) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Item) GetEquipObjId() int64 {
	if m != nil {
		return m.EquipObjId
	}
	return 0
}

type C2M_AddItem struct {
	TypeId               int32    `protobuf:"varint,1,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_AddItem) Reset()         { *m = C2M_AddItem{} }
func (m *C2M_AddItem) String() string { return proto.CompactTextString(m) }
func (*C2M_AddItem) ProtoMessage()    {}
func (*C2M_AddItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_20b3be662c2ee0bd, []int{1}
}

func (m *C2M_AddItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_AddItem.Unmarshal(m, b)
}
func (m *C2M_AddItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_AddItem.Marshal(b, m, deterministic)
}
func (m *C2M_AddItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_AddItem.Merge(m, src)
}
func (m *C2M_AddItem) XXX_Size() int {
	return xxx_messageInfo_C2M_AddItem.Size(m)
}
func (m *C2M_AddItem) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_AddItem.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_AddItem proto.InternalMessageInfo

func (m *C2M_AddItem) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type C2M_DelItem struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_DelItem) Reset()         { *m = C2M_DelItem{} }
func (m *C2M_DelItem) String() string { return proto.CompactTextString(m) }
func (*C2M_DelItem) ProtoMessage()    {}
func (*C2M_DelItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_20b3be662c2ee0bd, []int{2}
}

func (m *C2M_DelItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_DelItem.Unmarshal(m, b)
}
func (m *C2M_DelItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_DelItem.Marshal(b, m, deterministic)
}
func (m *C2M_DelItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_DelItem.Merge(m, src)
}
func (m *C2M_DelItem) XXX_Size() int {
	return xxx_messageInfo_C2M_DelItem.Size(m)
}
func (m *C2M_DelItem) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_DelItem.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_DelItem proto.InternalMessageInfo

func (m *C2M_DelItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type C2M_UseItem struct {
	ItemId               int64    `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_UseItem) Reset()         { *m = C2M_UseItem{} }
func (m *C2M_UseItem) String() string { return proto.CompactTextString(m) }
func (*C2M_UseItem) ProtoMessage()    {}
func (*C2M_UseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_20b3be662c2ee0bd, []int{3}
}

func (m *C2M_UseItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_UseItem.Unmarshal(m, b)
}
func (m *C2M_UseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_UseItem.Marshal(b, m, deterministic)
}
func (m *C2M_UseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_UseItem.Merge(m, src)
}
func (m *C2M_UseItem) XXX_Size() int {
	return xxx_messageInfo_C2M_UseItem.Size(m)
}
func (m *C2M_UseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_UseItem.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_UseItem proto.InternalMessageInfo

func (m *C2M_UseItem) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

type C2M_QueryItems struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_QueryItems) Reset()         { *m = C2M_QueryItems{} }
func (m *C2M_QueryItems) String() string { return proto.CompactTextString(m) }
func (*C2M_QueryItems) ProtoMessage()    {}
func (*C2M_QueryItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_20b3be662c2ee0bd, []int{4}
}

func (m *C2M_QueryItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_QueryItems.Unmarshal(m, b)
}
func (m *C2M_QueryItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_QueryItems.Marshal(b, m, deterministic)
}
func (m *C2M_QueryItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_QueryItems.Merge(m, src)
}
func (m *C2M_QueryItems) XXX_Size() int {
	return xxx_messageInfo_C2M_QueryItems.Size(m)
}
func (m *C2M_QueryItems) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_QueryItems.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_QueryItems proto.InternalMessageInfo

type M2C_ItemList struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_ItemList) Reset()         { *m = M2C_ItemList{} }
func (m *M2C_ItemList) String() string { return proto.CompactTextString(m) }
func (*M2C_ItemList) ProtoMessage()    {}
func (*M2C_ItemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_20b3be662c2ee0bd, []int{5}
}

func (m *M2C_ItemList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_ItemList.Unmarshal(m, b)
}
func (m *M2C_ItemList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_ItemList.Marshal(b, m, deterministic)
}
func (m *M2C_ItemList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_ItemList.Merge(m, src)
}
func (m *M2C_ItemList) XXX_Size() int {
	return xxx_messageInfo_M2C_ItemList.Size(m)
}
func (m *M2C_ItemList) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_ItemList.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_ItemList proto.InternalMessageInfo

func (m *M2C_ItemList) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type M2C_ItemAdd struct {
	Item                 *Item    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_ItemAdd) Reset()         { *m = M2C_ItemAdd{} }
func (m *M2C_ItemAdd) String() string { return proto.CompactTextString(m) }
func (*M2C_ItemAdd) ProtoMessage()    {}
func (*M2C_ItemAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_20b3be662c2ee0bd, []int{6}
}

func (m *M2C_ItemAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_ItemAdd.Unmarshal(m, b)
}
func (m *M2C_ItemAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_ItemAdd.Marshal(b, m, deterministic)
}
func (m *M2C_ItemAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_ItemAdd.Merge(m, src)
}
func (m *M2C_ItemAdd) XXX_Size() int {
	return xxx_messageInfo_M2C_ItemAdd.Size(m)
}
func (m *M2C_ItemAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_ItemAdd.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_ItemAdd proto.InternalMessageInfo

func (m *M2C_ItemAdd) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type M2C_ItemUpdate struct {
	Item                 *Item    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_ItemUpdate) Reset()         { *m = M2C_ItemUpdate{} }
func (m *M2C_ItemUpdate) String() string { return proto.CompactTextString(m) }
func (*M2C_ItemUpdate) ProtoMessage()    {}
func (*M2C_ItemUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_20b3be662c2ee0bd, []int{7}
}

func (m *M2C_ItemUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_ItemUpdate.Unmarshal(m, b)
}
func (m *M2C_ItemUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_ItemUpdate.Marshal(b, m, deterministic)
}
func (m *M2C_ItemUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_ItemUpdate.Merge(m, src)
}
func (m *M2C_ItemUpdate) XXX_Size() int {
	return xxx_messageInfo_M2C_ItemUpdate.Size(m)
}
func (m *M2C_ItemUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_ItemUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_ItemUpdate proto.InternalMessageInfo

func (m *M2C_ItemUpdate) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type M2C_DelItem struct {
	ItemId               int64    `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_DelItem) Reset()         { *m = M2C_DelItem{} }
func (m *M2C_DelItem) String() string { return proto.CompactTextString(m) }
func (*M2C_DelItem) ProtoMessage()    {}
func (*M2C_DelItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_20b3be662c2ee0bd, []int{8}
}

func (m *M2C_DelItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_DelItem.Unmarshal(m, b)
}
func (m *M2C_DelItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_DelItem.Marshal(b, m, deterministic)
}
func (m *M2C_DelItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_DelItem.Merge(m, src)
}
func (m *M2C_DelItem) XXX_Size() int {
	return xxx_messageInfo_M2C_DelItem.Size(m)
}
func (m *M2C_DelItem) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_DelItem.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_DelItem proto.InternalMessageInfo

func (m *M2C_DelItem) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "game.Item")
	proto.RegisterType((*C2M_AddItem)(nil), "game.C2M_AddItem")
	proto.RegisterType((*C2M_DelItem)(nil), "game.C2M_DelItem")
	proto.RegisterType((*C2M_UseItem)(nil), "game.C2M_UseItem")
	proto.RegisterType((*C2M_QueryItems)(nil), "game.C2M_QueryItems")
	proto.RegisterType((*M2C_ItemList)(nil), "game.M2C_ItemList")
	proto.RegisterType((*M2C_ItemAdd)(nil), "game.M2C_ItemAdd")
	proto.RegisterType((*M2C_ItemUpdate)(nil), "game.M2C_ItemUpdate")
	proto.RegisterType((*M2C_DelItem)(nil), "game.M2C_DelItem")
}

func init() { proto.RegisterFile("game/item.proto", fileDescriptor_20b3be662c2ee0bd) }

var fileDescriptor_20b3be662c2ee0bd = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x43, 0xa1, 0x3d, 0x0c, 0x06, 0x9b, 0x3d, 0x18, 0x2e, 0x36, 0x64, 0x93, 0x1a, 0x2e,
	0x85, 0x06, 0x9f, 0xa0, 0x56, 0x0f, 0x24, 0x56, 0x23, 0xb1, 0x17, 0x2f, 0x15, 0xdc, 0x49, 0xc5,
	0x88, 0x20, 0xbb, 0x98, 0xf4, 0xed, 0xcd, 0x2c, 0x5d, 0xad, 0x89, 0xc6, 0x13, 0x33, 0xff, 0xfc,
	0xf3, 0x0d, 0x33, 0x0b, 0xc7, 0xdb, 0xbc, 0xc2, 0xb8, 0x54, 0x58, 0x45, 0x4d, 0x5b, 0xab, 0x9a,
	0x39, 0x24, 0xf0, 0x47, 0x70, 0x52, 0x85, 0x15, 0xf3, 0x60, 0x90, 0x0a, 0xdf, 0x0a, 0xac, 0xd0,
	0xce, 0x06, 0xa9, 0x60, 0x27, 0x30, 0xba, 0xdf, 0x35, 0x98, 0x0a, 0x7f, 0x10, 0x58, 0xe1, 0x30,
	0xdb, 0x67, 0x6c, 0x0c, 0xf6, 0x4d, 0x57, 0xf9, 0xb6, 0x16, 0x29, 0x64, 0x13, 0x80, 0xab, 0xf7,
	0xae, 0x6c, 0x6e, 0x8b, 0x97, 0x54, 0xf8, 0x8e, 0x26, 0x1c, 0x28, 0x7c, 0x0a, 0xee, 0x32, 0x59,
	0x6d, 0x16, 0x42, 0xe8, 0x41, 0xdf, 0x60, 0xeb, 0x10, 0xcc, 0x4f, 0x7b, 0xdb, 0x25, 0xbe, 0xfe,
	0xf6, 0x3f, 0x86, 0xb2, 0x96, 0x68, 0x28, 0xf4, 0xfd, 0xb2, 0xec, 0x33, 0x3e, 0x06, 0x8f, 0x6c,
	0x77, 0x1d, 0xb6, 0x3b, 0x92, 0x24, 0x9f, 0xc3, 0xd1, 0x2a, 0x59, 0x6e, 0x28, 0xb9, 0x2e, 0xa5,
	0x62, 0x01, 0x0c, 0xe9, 0x08, 0xd2, 0xb7, 0x02, 0x3b, 0x74, 0x13, 0x88, 0xe8, 0x0c, 0x11, 0x95,
	0xb3, 0xbe, 0xc0, 0x67, 0xe0, 0x9a, 0x8e, 0x85, 0x10, 0x6c, 0x02, 0x0e, 0xe9, 0x7a, 0xd0, 0x4f,
	0xbf, 0xd6, 0xf9, 0x1c, 0x3c, 0x63, 0x5f, 0x37, 0x22, 0x57, 0xf8, 0x6f, 0xc7, 0xb4, 0x1f, 0x60,
	0x56, 0xfd, 0x63, 0x97, 0x8b, 0xf0, 0xe1, 0x6c, 0x5b, 0xaa, 0xe7, 0xae, 0x88, 0x9e, 0xea, 0x2a,
	0xc6, 0x5c, 0xaa, 0x19, 0x0a, 0x7c, 0x8b, 0x25, 0xb6, 0x1f, 0xd8, 0xc6, 0xfa, 0x15, 0x63, 0x62,
	0x17, 0x23, 0x1d, 0x9f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x69, 0x9f, 0x61, 0xe4, 0x01,
	0x00, 0x00,
}
