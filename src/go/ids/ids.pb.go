// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ids/ids.proto

package ids

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

// Flight models the unique id of a flight in the context of AirMap.
type Flight struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Flight) Reset()         { *m = Flight{} }
func (m *Flight) String() string { return proto.CompactTextString(m) }
func (*Flight) ProtoMessage()    {}
func (*Flight) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb53c9ac7a7c1b5, []int{0}
}

func (m *Flight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Flight.Unmarshal(m, b)
}
func (m *Flight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Flight.Marshal(b, m, deterministic)
}
func (m *Flight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flight.Merge(m, src)
}
func (m *Flight) XXX_Size() int {
	return xxx_messageInfo_Flight.Size(m)
}
func (m *Flight) XXX_DiscardUnknown() {
	xxx_messageInfo_Flight.DiscardUnknown(m)
}

var xxx_messageInfo_Flight proto.InternalMessageInfo

func (m *Flight) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

// USP models the unique id of a UTM Service Provider (USP)
type USP struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *USP) Reset()         { *m = USP{} }
func (m *USP) String() string { return proto.CompactTextString(m) }
func (*USP) ProtoMessage()    {}
func (*USP) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bb53c9ac7a7c1b5, []int{1}
}

func (m *USP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_USP.Unmarshal(m, b)
}
func (m *USP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_USP.Marshal(b, m, deterministic)
}
func (m *USP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_USP.Merge(m, src)
}
func (m *USP) XXX_Size() int {
	return xxx_messageInfo_USP.Size(m)
}
func (m *USP) XXX_DiscardUnknown() {
	xxx_messageInfo_USP.DiscardUnknown(m)
}

var xxx_messageInfo_USP proto.InternalMessageInfo

func (m *USP) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

func init() {
	proto.RegisterType((*Flight)(nil), "ids.Flight")
	proto.RegisterType((*USP)(nil), "ids.USP")
}

func init() { proto.RegisterFile("ids/ids.proto", fileDescriptor_9bb53c9ac7a7c1b5) }

var fileDescriptor_9bb53c9ac7a7c1b5 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4c, 0x29, 0xd6,
	0xcf, 0x4c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4c, 0x29, 0x56, 0x52,
	0xe5, 0x62, 0x73, 0xcb, 0xc9, 0x4c, 0xcf, 0x28, 0x11, 0x92, 0xe6, 0xe2, 0x4c, 0x2c, 0x8e, 0x2f,
	0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x48, 0x2c, 0x0e,
	0x06, 0xf3, 0x95, 0x94, 0xb8, 0x98, 0x43, 0x83, 0x03, 0xf0, 0xaa, 0x71, 0xd2, 0x8c, 0x52, 0x4f,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcc, 0x2c, 0xca, 0x4d, 0x2c,
	0xd0, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0x2d, 0xd6, 0x2f, 0x2e, 0x4a, 0xd6,
	0x4f, 0xcf, 0x07, 0x39, 0x20, 0x89, 0x0d, 0xec, 0x02, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x63, 0x4c, 0x8b, 0x97, 0x92, 0x00, 0x00, 0x00,
}
