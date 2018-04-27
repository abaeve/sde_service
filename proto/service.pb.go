// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package sde is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	TypeIdRequest
	TypeNameRequest
	TypeResponse
	TypeNameAndIdResponse
	Type
*/
package sde

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Locale int32

const (
	Locale_EN Locale = 0
	Locale_DE Locale = 1
	Locale_FR Locale = 2
	Locale_JA Locale = 3
	Locale_RU Locale = 4
	Locale_ZH Locale = 5
)

var Locale_name = map[int32]string{
	0: "EN",
	1: "DE",
	2: "FR",
	3: "JA",
	4: "RU",
	5: "ZH",
}
var Locale_value = map[string]int32{
	"EN": 0,
	"DE": 1,
	"FR": 2,
	"JA": 3,
	"RU": 4,
	"ZH": 5,
}

func (x Locale) String() string {
	return proto.EnumName(Locale_name, int32(x))
}
func (Locale) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TypeIdRequest struct {
	TypeId []int32 `protobuf:"varint,1,rep,packed,name=TypeId" json:"TypeId,omitempty"`
	Locale Locale  `protobuf:"varint,2,opt,name=Locale,enum=sde.Locale" json:"Locale,omitempty"`
}

func (m *TypeIdRequest) Reset()                    { *m = TypeIdRequest{} }
func (m *TypeIdRequest) String() string            { return proto.CompactTextString(m) }
func (*TypeIdRequest) ProtoMessage()               {}
func (*TypeIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TypeIdRequest) GetTypeId() []int32 {
	if m != nil {
		return m.TypeId
	}
	return nil
}

func (m *TypeIdRequest) GetLocale() Locale {
	if m != nil {
		return m.Locale
	}
	return Locale_EN
}

type TypeNameRequest struct {
	TypeName []string `protobuf:"bytes,1,rep,name=TypeName" json:"TypeName,omitempty"`
	Locale   Locale   `protobuf:"varint,2,opt,name=Locale,enum=sde.Locale" json:"Locale,omitempty"`
}

func (m *TypeNameRequest) Reset()                    { *m = TypeNameRequest{} }
func (m *TypeNameRequest) String() string            { return proto.CompactTextString(m) }
func (*TypeNameRequest) ProtoMessage()               {}
func (*TypeNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TypeNameRequest) GetTypeName() []string {
	if m != nil {
		return m.TypeName
	}
	return nil
}

func (m *TypeNameRequest) GetLocale() Locale {
	if m != nil {
		return m.Locale
	}
	return Locale_EN
}

type TypeResponse struct {
	Type []*Type `protobuf:"bytes,1,rep,name=Type" json:"Type,omitempty"`
}

func (m *TypeResponse) Reset()                    { *m = TypeResponse{} }
func (m *TypeResponse) String() string            { return proto.CompactTextString(m) }
func (*TypeResponse) ProtoMessage()               {}
func (*TypeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TypeResponse) GetType() []*Type {
	if m != nil {
		return m.Type
	}
	return nil
}

type TypeNameAndIdResponse struct {
	TypeId []int32  `protobuf:"varint,1,rep,packed,name=TypeId" json:"TypeId,omitempty"`
	Name   []string `protobuf:"bytes,2,rep,name=Name" json:"Name,omitempty"`
}

func (m *TypeNameAndIdResponse) Reset()                    { *m = TypeNameAndIdResponse{} }
func (m *TypeNameAndIdResponse) String() string            { return proto.CompactTextString(m) }
func (*TypeNameAndIdResponse) ProtoMessage()               {}
func (*TypeNameAndIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TypeNameAndIdResponse) GetTypeId() []int32 {
	if m != nil {
		return m.TypeId
	}
	return nil
}

func (m *TypeNameAndIdResponse) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

type Type struct {
	TypeId        int32   `protobuf:"varint,1,opt,name=TypeId" json:"TypeId,omitempty"`
	BasePrice     float32 `protobuf:"fixed32,2,opt,name=BasePrice" json:"BasePrice,omitempty"`
	Capacity      float32 `protobuf:"fixed32,3,opt,name=Capacity" json:"Capacity,omitempty"`
	Description   string  `protobuf:"bytes,4,opt,name=Description" json:"Description,omitempty"`
	GraphicId     int32   `protobuf:"varint,5,opt,name=GraphicId" json:"GraphicId,omitempty"`
	GroupId       int32   `protobuf:"varint,6,opt,name=GroupId" json:"GroupId,omitempty"`
	IconId        int32   `protobuf:"varint,7,opt,name=IconId" json:"IconId,omitempty"`
	MarketGroupId int32   `protobuf:"varint,8,opt,name=MarketGroupId" json:"MarketGroupId,omitempty"`
	Mass          float32 `protobuf:"fixed32,9,opt,name=Mass" json:"Mass,omitempty"`
	Name          string  `protobuf:"bytes,10,opt,name=Name" json:"Name,omitempty"`
	PortionSize   int32   `protobuf:"varint,11,opt,name=PortionSize" json:"PortionSize,omitempty"`
	Published     bool    `protobuf:"varint,12,opt,name=Published" json:"Published,omitempty"`
	Radius        float32 `protobuf:"fixed32,13,opt,name=Radius" json:"Radius,omitempty"`
	FactionName   string  `protobuf:"bytes,14,opt,name=FactionName" json:"FactionName,omitempty"`
	Volume        float32 `protobuf:"fixed32,15,opt,name=Volume" json:"Volume,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Type) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Type) GetBasePrice() float32 {
	if m != nil {
		return m.BasePrice
	}
	return 0
}

func (m *Type) GetCapacity() float32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Type) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Type) GetGraphicId() int32 {
	if m != nil {
		return m.GraphicId
	}
	return 0
}

func (m *Type) GetGroupId() int32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *Type) GetIconId() int32 {
	if m != nil {
		return m.IconId
	}
	return 0
}

func (m *Type) GetMarketGroupId() int32 {
	if m != nil {
		return m.MarketGroupId
	}
	return 0
}

func (m *Type) GetMass() float32 {
	if m != nil {
		return m.Mass
	}
	return 0
}

func (m *Type) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type) GetPortionSize() int32 {
	if m != nil {
		return m.PortionSize
	}
	return 0
}

func (m *Type) GetPublished() bool {
	if m != nil {
		return m.Published
	}
	return false
}

func (m *Type) GetRadius() float32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *Type) GetFactionName() string {
	if m != nil {
		return m.FactionName
	}
	return ""
}

func (m *Type) GetVolume() float32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func init() {
	proto.RegisterType((*TypeIdRequest)(nil), "sde.TypeIdRequest")
	proto.RegisterType((*TypeNameRequest)(nil), "sde.TypeNameRequest")
	proto.RegisterType((*TypeResponse)(nil), "sde.TypeResponse")
	proto.RegisterType((*TypeNameAndIdResponse)(nil), "sde.TypeNameAndIdResponse")
	proto.RegisterType((*Type)(nil), "sde.Type")
	proto.RegisterEnum("sde.Locale", Locale_name, Locale_value)
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xcd, 0xe6, 0xab, 0xd9, 0x49, 0x93, 0x2e, 0x56, 0x41, 0x56, 0x04, 0x52, 0xb4, 0x70, 0x88,
	0x90, 0xc8, 0x21, 0x5c, 0xb8, 0x70, 0x48, 0x9b, 0xa6, 0x04, 0xb5, 0x55, 0x70, 0x81, 0x03, 0x37,
	0x77, 0xd7, 0x52, 0x2c, 0xd2, 0x78, 0xb1, 0xb3, 0x48, 0xe1, 0x27, 0xf1, 0x77, 0xf8, 0x43, 0x68,
	0xc6, 0xd9, 0x4d, 0x82, 0xa8, 0xd4, 0xd3, 0xf8, 0x3d, 0xfb, 0x3d, 0x3f, 0xdb, 0x63, 0xe8, 0x38,
	0x65, 0x7f, 0xea, 0x44, 0x0d, 0x33, 0x6b, 0xd6, 0x86, 0xd5, 0x5c, 0xaa, 0xe2, 0x2b, 0xe8, 0x7c,
	0xde, 0x64, 0x6a, 0x96, 0x0a, 0xf5, 0x23, 0x57, 0x6e, 0xcd, 0x9e, 0x41, 0xd3, 0x13, 0x3c, 0xe8,
	0xd7, 0x06, 0x0d, 0xb1, 0x45, 0xec, 0x25, 0x34, 0xaf, 0x4c, 0x22, 0x97, 0x8a, 0x57, 0xfb, 0xc1,
	0xa0, 0x3b, 0x6a, 0x0f, 0x5d, 0xaa, 0x86, 0x9e, 0x12, 0xdb, 0xa9, 0x58, 0xc0, 0x09, 0x2e, 0xbf,
	0x91, 0xf7, 0xaa, 0xf0, 0xeb, 0x41, 0xab, 0xa0, 0xc8, 0x31, 0x14, 0x25, 0x7e, 0x9c, 0xe7, 0x1b,
	0x38, 0x46, 0x81, 0x50, 0x2e, 0x33, 0x2b, 0xa7, 0xd8, 0x0b, 0xa8, 0x23, 0x26, 0xb3, 0xf6, 0x28,
	0x24, 0x09, 0x2d, 0x20, 0x3a, 0x3e, 0x87, 0xa7, 0x85, 0xff, 0x78, 0x95, 0xe2, 0xb9, 0xb6, 0xba,
	0x87, 0x0e, 0xc6, 0xa0, 0x4e, 0xe1, 0xaa, 0x14, 0x8e, 0xc6, 0xf1, 0xef, 0x9a, 0xdf, 0xe4, 0x40,
	0x14, 0xec, 0x89, 0x9e, 0x43, 0x78, 0x26, 0x9d, 0x9a, 0x5b, 0x9d, 0xf8, 0xf0, 0x55, 0xb1, 0x23,
	0xf0, 0xcc, 0xe7, 0x32, 0x93, 0x89, 0x5e, 0x6f, 0x78, 0x8d, 0x26, 0x4b, 0xcc, 0xfa, 0xd0, 0x9e,
	0x28, 0x97, 0x58, 0x9d, 0xad, 0xb5, 0x59, 0xf1, 0x7a, 0x3f, 0x18, 0x84, 0x62, 0x9f, 0x42, 0xef,
	0x4b, 0x2b, 0xb3, 0x85, 0x4e, 0x66, 0x29, 0x6f, 0xd0, 0xb6, 0x3b, 0x82, 0x71, 0x38, 0xba, 0xb4,
	0x26, 0xcf, 0x66, 0x29, 0x6f, 0xd2, 0x5c, 0x01, 0x31, 0xeb, 0x2c, 0x31, 0xab, 0x59, 0xca, 0x8f,
	0x7c, 0x56, 0x8f, 0xd8, 0x2b, 0xe8, 0x5c, 0x4b, 0xfb, 0x5d, 0xad, 0x0b, 0x5d, 0x8b, 0xa6, 0x0f,
	0x49, 0xbc, 0x86, 0x6b, 0xe9, 0x1c, 0x0f, 0x29, 0x2f, 0x8d, 0xcb, 0xab, 0x01, 0x0a, 0x49, 0x63,
	0xcc, 0x3f, 0x37, 0x16, 0x83, 0xde, 0xea, 0x5f, 0x8a, 0xb7, 0xc9, 0x6b, 0x9f, 0xc2, 0xfc, 0xf3,
	0xfc, 0x6e, 0xa9, 0xdd, 0x42, 0xa5, 0xfc, 0xb8, 0x1f, 0x0c, 0x5a, 0x62, 0x47, 0x60, 0x4a, 0x21,
	0x53, 0x9d, 0x3b, 0xde, 0xa1, 0x9d, 0xb6, 0x08, 0x7d, 0xa7, 0x32, 0x41, 0x13, 0xda, 0xb2, 0xeb,
	0xef, 0x65, 0x8f, 0x42, 0xe5, 0x57, 0xb3, 0xcc, 0xef, 0x15, 0x3f, 0xf1, 0x4a, 0x8f, 0x5e, 0xbf,
	0x2b, 0xba, 0x88, 0x35, 0xa1, 0x7a, 0x71, 0x13, 0x55, 0xb0, 0x4e, 0x2e, 0xa2, 0x00, 0xeb, 0x54,
	0x44, 0x55, 0xac, 0x1f, 0xc7, 0x51, 0x0d, 0xab, 0xf8, 0x12, 0xd5, 0xb1, 0x7e, 0xfb, 0x10, 0x35,
	0x46, 0x7f, 0x02, 0x08, 0xf1, 0x41, 0x3f, 0xe5, 0xca, 0x6e, 0xd8, 0x7b, 0x60, 0x53, 0xbd, 0x4a,
	0x91, 0x70, 0x67, 0x1b, 0xff, 0xd0, 0x8e, 0xb1, 0xb2, 0xc1, 0xca, 0x3f, 0xd2, 0x7b, 0xb2, 0x6b,
	0xba, 0x6d, 0x77, 0xc5, 0x15, 0x36, 0x86, 0xd3, 0x7f, 0xe4, 0x98, 0xda, 0xb1, 0xd3, 0x72, 0xf1,
	0xde, 0xb7, 0xf8, 0xbf, 0xc5, 0x04, 0xba, 0xb7, 0x4a, 0xda, 0x64, 0x31, 0x35, 0x96, 0x7c, 0x1e,
	0x10, 0xf7, 0x0e, 0xd8, 0x83, 0x36, 0x8f, 0x2b, 0x77, 0x4d, 0xfa, 0xde, 0x6f, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xbf, 0xd4, 0xf3, 0xdb, 0xef, 0x03, 0x00, 0x00,
}
