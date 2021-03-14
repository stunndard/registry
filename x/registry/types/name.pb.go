// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registry/name.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Name struct {
	Creator string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string                                   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name    string                                   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Owner   string                                   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Price   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
	Onsale  bool                                     `protobuf:"varint,6,opt,name=onsale,proto3" json:"onsale,omitempty"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c7146e02c63ddfb, []int{0}
}
func (m *Name) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Name.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(m, src)
}
func (m *Name) XXX_Size() int {
	return m.Size()
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Name) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Name) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Name) GetPrice() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Name) GetOnsale() bool {
	if m != nil {
		return m.Onsale
	}
	return false
}

type MsgBuyName struct {
	Creator string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name    string                                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Bid     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=bid,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"bid"`
	Onsale  bool                                     `protobuf:"varint,4,opt,name=onsale,proto3" json:"onsale,omitempty"`
}

func (m *MsgBuyName) Reset()         { *m = MsgBuyName{} }
func (m *MsgBuyName) String() string { return proto.CompactTextString(m) }
func (*MsgBuyName) ProtoMessage()    {}
func (*MsgBuyName) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c7146e02c63ddfb, []int{1}
}
func (m *MsgBuyName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBuyName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBuyName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBuyName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBuyName.Merge(m, src)
}
func (m *MsgBuyName) XXX_Size() int {
	return m.Size()
}
func (m *MsgBuyName) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBuyName.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBuyName proto.InternalMessageInfo

func (m *MsgBuyName) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgBuyName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgBuyName) GetBid() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Bid
	}
	return nil
}

func (m *MsgBuyName) GetOnsale() bool {
	if m != nil {
		return m.Onsale
	}
	return false
}

type MsgUpdateName struct {
	Creator string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name    string                                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner   string                                   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Price   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
	Onsale  bool                                     `protobuf:"varint,5,opt,name=onsale,proto3" json:"onsale,omitempty"`
}

func (m *MsgUpdateName) Reset()         { *m = MsgUpdateName{} }
func (m *MsgUpdateName) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateName) ProtoMessage()    {}
func (*MsgUpdateName) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c7146e02c63ddfb, []int{2}
}
func (m *MsgUpdateName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateName.Merge(m, src)
}
func (m *MsgUpdateName) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateName) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateName.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateName proto.InternalMessageInfo

func (m *MsgUpdateName) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgUpdateName) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgUpdateName) GetPrice() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *MsgUpdateName) GetOnsale() bool {
	if m != nil {
		return m.Onsale
	}
	return false
}

type MsgDeleteName struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *MsgDeleteName) Reset()         { *m = MsgDeleteName{} }
func (m *MsgDeleteName) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteName) ProtoMessage()    {}
func (*MsgDeleteName) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c7146e02c63ddfb, []int{3}
}
func (m *MsgDeleteName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteName.Merge(m, src)
}
func (m *MsgDeleteName) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteName) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteName.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteName proto.InternalMessageInfo

func (m *MsgDeleteName) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Name)(nil), "stunndard.registry.registry.Name")
	proto.RegisterType((*MsgBuyName)(nil), "stunndard.registry.registry.MsgBuyName")
	proto.RegisterType((*MsgUpdateName)(nil), "stunndard.registry.registry.MsgUpdateName")
	proto.RegisterType((*MsgDeleteName)(nil), "stunndard.registry.registry.MsgDeleteName")
}

func init() { proto.RegisterFile("registry/name.proto", fileDescriptor_7c7146e02c63ddfb) }

var fileDescriptor_7c7146e02c63ddfb = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xbf, 0x6e, 0xdb, 0x30,
	0x10, 0xc6, 0x45, 0xfd, 0x71, 0x5b, 0x16, 0xed, 0xc0, 0x1a, 0x05, 0xeb, 0x02, 0xb2, 0xe1, 0x49,
	0x4b, 0xc9, 0xba, 0x9d, 0xbb, 0xb8, 0x1d, 0xba, 0xb8, 0x83, 0x81, 0x2e, 0x01, 0x32, 0x50, 0x12,
	0xa1, 0x10, 0xb1, 0x44, 0x81, 0xa4, 0x93, 0xf8, 0x2d, 0xf2, 0x1c, 0xd9, 0xf2, 0x16, 0x9e, 0x02,
	0x2f, 0x01, 0x32, 0x25, 0x81, 0xfd, 0x22, 0x81, 0x28, 0x5b, 0xb1, 0x81, 0x00, 0x01, 0x8c, 0x20,
	0x93, 0xbe, 0xbb, 0xd3, 0x1d, 0x7e, 0x1f, 0x8e, 0x07, 0x3f, 0x29, 0x9e, 0x09, 0x6d, 0xd4, 0x8c,
	0x16, 0x2c, 0xe7, 0xa4, 0x54, 0xd2, 0x48, 0xf4, 0x55, 0x9b, 0x69, 0x51, 0xa4, 0x4c, 0xa5, 0x64,
	0x53, 0x6e, 0x44, 0xa7, 0x9d, 0xc9, 0x4c, 0xda, 0xff, 0x68, 0xa5, 0xea, 0x96, 0x4e, 0x98, 0x48,
	0x9d, 0x4b, 0x4d, 0x63, 0xa6, 0x39, 0x3d, 0x19, 0xc4, 0xdc, 0xb0, 0x01, 0x4d, 0xa4, 0x28, 0xea,
	0x7a, 0xff, 0x1a, 0x40, 0xff, 0x1f, 0xcb, 0x39, 0xc2, 0xf0, 0x4d, 0xa2, 0x38, 0x33, 0x52, 0x61,
	0xd0, 0x03, 0xd1, 0xbb, 0xf1, 0x26, 0x44, 0x1f, 0xa1, 0x2b, 0x52, 0xec, 0xda, 0xa4, 0x2b, 0x52,
	0x84, 0xa0, 0x5f, 0x31, 0x61, 0xcf, 0x66, 0xac, 0x46, 0x6d, 0x18, 0xc8, 0xd3, 0x82, 0x2b, 0xec,
	0xdb, 0x64, 0x1d, 0x20, 0x06, 0x83, 0x52, 0x89, 0x84, 0xe3, 0xa0, 0xe7, 0x45, 0xef, 0x7f, 0x7c,
	0x21, 0x35, 0x0c, 0xa9, 0x60, 0xc8, 0x1a, 0x86, 0xfc, 0x96, 0xa2, 0x18, 0x7e, 0x9f, 0xdf, 0x76,
	0x9d, 0x8b, 0xbb, 0x6e, 0x94, 0x09, 0x73, 0x34, 0x8d, 0x49, 0x22, 0x73, 0xba, 0x26, 0xaf, 0x3f,
	0xdf, 0x74, 0x7a, 0x4c, 0xcd, 0xac, 0xe4, 0xda, 0x36, 0xe8, 0x71, 0x3d, 0x19, 0x7d, 0x86, 0x2d,
	0x59, 0x68, 0x36, 0xe1, 0xb8, 0xd5, 0x03, 0xd1, 0xdb, 0xf1, 0x3a, 0xea, 0x5f, 0x02, 0x08, 0x47,
	0x3a, 0x1b, 0x4e, 0x67, 0xcf, 0xb8, 0xdb, 0xb8, 0x71, 0xb7, 0xdc, 0x1c, 0x42, 0x2f, 0x16, 0x29,
	0xf6, 0x5e, 0x9e, 0xba, 0x9a, 0xbb, 0xc5, 0xec, 0xef, 0x30, 0x5f, 0x01, 0xf8, 0x61, 0xa4, 0xb3,
	0xff, 0x65, 0xca, 0x0c, 0xdf, 0x03, 0xbb, 0x59, 0x82, 0xf7, 0xe4, 0x12, 0xfc, 0x57, 0x58, 0x42,
	0xb0, 0x63, 0xe8, 0x97, 0xf5, 0xf3, 0x87, 0x4f, 0xf8, 0x3e, 0x7e, 0x86, 0x7f, 0xe7, 0xcb, 0x10,
	0x2c, 0x96, 0x21, 0xb8, 0x5f, 0x86, 0xe0, 0x7c, 0x15, 0x3a, 0x8b, 0x55, 0xe8, 0xdc, 0xac, 0x42,
	0xe7, 0x80, 0x6c, 0x11, 0x36, 0x37, 0x41, 0x9b, 0x93, 0x39, 0x7b, 0x94, 0x96, 0x36, 0x6e, 0xd9,
	0xc7, 0xfe, 0xf3, 0x21, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x39, 0x7d, 0x81, 0x56, 0x03, 0x00, 0x00,
}

func (m *Name) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Name) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Name) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Onsale {
		i--
		if m.Onsale {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintName(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintName(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintName(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintName(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintName(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBuyName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBuyName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBuyName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Onsale {
		i--
		if m.Onsale {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Bid) > 0 {
		for iNdEx := len(m.Bid) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bid[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintName(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintName(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintName(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Onsale {
		i--
		if m.Onsale {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintName(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintName(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintName(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintName(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintName(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintName(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintName(dAtA []byte, offset int, v uint64) int {
	offset -= sovName(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Name) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovName(uint64(l))
		}
	}
	if m.Onsale {
		n += 2
	}
	return n
}

func (m *MsgBuyName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	if len(m.Bid) > 0 {
		for _, e := range m.Bid {
			l = e.Size()
			n += 1 + l + sovName(uint64(l))
		}
	}
	if m.Onsale {
		n += 2
	}
	return n
}

func (m *MsgUpdateName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovName(uint64(l))
		}
	}
	if m.Onsale {
		n += 2
	}
	return n
}

func (m *MsgDeleteName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovName(uint64(l))
	}
	return n
}

func sovName(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozName(x uint64) (n int) {
	return sovName(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Name) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowName
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Name: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Name: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Onsale", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Onsale = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthName
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgBuyName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowName
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgBuyName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBuyName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bid = append(m.Bid, types.Coin{})
			if err := m.Bid[len(m.Bid)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Onsale", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Onsale = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthName
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowName
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Onsale", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Onsale = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthName
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowName
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowName
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthName
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthName
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipName(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthName
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipName(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowName
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowName
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowName
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthName
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupName
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthName
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthName        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowName          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupName = fmt.Errorf("proto: unexpected end of group")
)