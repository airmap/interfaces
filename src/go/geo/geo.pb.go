// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geo/geo.proto

package geo

import (
	fmt "fmt"
	_ "github.com/airmap/interfaces/src/go"
	measurements "github.com/airmap/interfaces/src/go/measurements"
	units "github.com/airmap/interfaces/src/go/units"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Coordinate2D models a WGS84 latitude and longitude.
type Coordinate2D struct {
	Latitude             *units.Degrees      `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            *units.Degrees      `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Error                *Coordinate2D_Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Coordinate2D) Reset()         { *m = Coordinate2D{} }
func (m *Coordinate2D) String() string { return proto.CompactTextString(m) }
func (*Coordinate2D) ProtoMessage()    {}
func (*Coordinate2D) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{0}
}

func (m *Coordinate2D) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate2D.Unmarshal(m, b)
}
func (m *Coordinate2D) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate2D.Marshal(b, m, deterministic)
}
func (m *Coordinate2D) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate2D.Merge(m, src)
}
func (m *Coordinate2D) XXX_Size() int {
	return xxx_messageInfo_Coordinate2D.Size(m)
}
func (m *Coordinate2D) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate2D.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate2D proto.InternalMessageInfo

func (m *Coordinate2D) GetLatitude() *units.Degrees {
	if m != nil {
		return m.Latitude
	}
	return nil
}

func (m *Coordinate2D) GetLongitude() *units.Degrees {
	if m != nil {
		return m.Longitude
	}
	return nil
}

func (m *Coordinate2D) GetError() *Coordinate2D_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Error models the inherent error in a coordinate measurement.
type Coordinate2D_Error struct {
	Latitude             *units.Degrees `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            *units.Degrees `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Coordinate2D_Error) Reset()         { *m = Coordinate2D_Error{} }
func (m *Coordinate2D_Error) String() string { return proto.CompactTextString(m) }
func (*Coordinate2D_Error) ProtoMessage()    {}
func (*Coordinate2D_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{0, 0}
}

func (m *Coordinate2D_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate2D_Error.Unmarshal(m, b)
}
func (m *Coordinate2D_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate2D_Error.Marshal(b, m, deterministic)
}
func (m *Coordinate2D_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate2D_Error.Merge(m, src)
}
func (m *Coordinate2D_Error) XXX_Size() int {
	return xxx_messageInfo_Coordinate2D_Error.Size(m)
}
func (m *Coordinate2D_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate2D_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate2D_Error proto.InternalMessageInfo

func (m *Coordinate2D_Error) GetLatitude() *units.Degrees {
	if m != nil {
		return m.Latitude
	}
	return nil
}

func (m *Coordinate2D_Error) GetLongitude() *units.Degrees {
	if m != nil {
		return m.Longitude
	}
	return nil
}

// Position models a postion in time
type Position4D struct {
	Timestamp            *timestamp.Timestamp            `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Position             *measurements.Position_Absolute `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Position4D) Reset()         { *m = Position4D{} }
func (m *Position4D) String() string { return proto.CompactTextString(m) }
func (*Position4D) ProtoMessage()    {}
func (*Position4D) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{1}
}

func (m *Position4D) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position4D.Unmarshal(m, b)
}
func (m *Position4D) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position4D.Marshal(b, m, deterministic)
}
func (m *Position4D) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position4D.Merge(m, src)
}
func (m *Position4D) XXX_Size() int {
	return xxx_messageInfo_Position4D.Size(m)
}
func (m *Position4D) XXX_DiscardUnknown() {
	xxx_messageInfo_Position4D.DiscardUnknown(m)
}

var xxx_messageInfo_Position4D proto.InternalMessageInfo

func (m *Position4D) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Position4D) GetPosition() *measurements.Position_Absolute {
	if m != nil {
		return m.Position
	}
	return nil
}

// BoundingBox models a geographic bounds
type BoundingBox struct {
	Southwest            *Coordinate2D `protobuf:"bytes,1,opt,name=southwest,proto3" json:"southwest,omitempty"`
	Northeast            *Coordinate2D `protobuf:"bytes,2,opt,name=northeast,proto3" json:"northeast,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BoundingBox) Reset()         { *m = BoundingBox{} }
func (m *BoundingBox) String() string { return proto.CompactTextString(m) }
func (*BoundingBox) ProtoMessage()    {}
func (*BoundingBox) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{2}
}

func (m *BoundingBox) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingBox.Unmarshal(m, b)
}
func (m *BoundingBox) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingBox.Marshal(b, m, deterministic)
}
func (m *BoundingBox) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingBox.Merge(m, src)
}
func (m *BoundingBox) XXX_Size() int {
	return xxx_messageInfo_BoundingBox.Size(m)
}
func (m *BoundingBox) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingBox.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingBox proto.InternalMessageInfo

func (m *BoundingBox) GetSouthwest() *Coordinate2D {
	if m != nil {
		return m.Southwest
	}
	return nil
}

func (m *BoundingBox) GetNortheast() *Coordinate2D {
	if m != nil {
		return m.Northeast
	}
	return nil
}

// LinearRing models an array of point coordinates where the first and last point as equal.
type LinearRing struct {
	Coordinates          []*Coordinate2D `protobuf:"bytes,1,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LinearRing) Reset()         { *m = LinearRing{} }
func (m *LinearRing) String() string { return proto.CompactTextString(m) }
func (*LinearRing) ProtoMessage()    {}
func (*LinearRing) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{3}
}

func (m *LinearRing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinearRing.Unmarshal(m, b)
}
func (m *LinearRing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinearRing.Marshal(b, m, deterministic)
}
func (m *LinearRing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinearRing.Merge(m, src)
}
func (m *LinearRing) XXX_Size() int {
	return xxx_messageInfo_LinearRing.Size(m)
}
func (m *LinearRing) XXX_DiscardUnknown() {
	xxx_messageInfo_LinearRing.DiscardUnknown(m)
}

var xxx_messageInfo_LinearRing proto.InternalMessageInfo

func (m *LinearRing) GetCoordinates() []*Coordinate2D {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

// Polygon models an array of LinearRings where the first ring is the exterior ring and subsequent rings
// represent the interior rings/cutouts.
type Polygon struct {
	Rings                []*LinearRing `protobuf:"bytes,1,rep,name=rings,proto3" json:"rings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Polygon) Reset()         { *m = Polygon{} }
func (m *Polygon) String() string { return proto.CompactTextString(m) }
func (*Polygon) ProtoMessage()    {}
func (*Polygon) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1967117b7b016d8, []int{4}
}

func (m *Polygon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Polygon.Unmarshal(m, b)
}
func (m *Polygon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Polygon.Marshal(b, m, deterministic)
}
func (m *Polygon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Polygon.Merge(m, src)
}
func (m *Polygon) XXX_Size() int {
	return xxx_messageInfo_Polygon.Size(m)
}
func (m *Polygon) XXX_DiscardUnknown() {
	xxx_messageInfo_Polygon.DiscardUnknown(m)
}

var xxx_messageInfo_Polygon proto.InternalMessageInfo

func (m *Polygon) GetRings() []*LinearRing {
	if m != nil {
		return m.Rings
	}
	return nil
}

func init() {
	proto.RegisterType((*Coordinate2D)(nil), "geo.Coordinate2D")
	proto.RegisterType((*Coordinate2D_Error)(nil), "geo.Coordinate2D.Error")
	proto.RegisterType((*Position4D)(nil), "geo.Position4D")
	proto.RegisterType((*BoundingBox)(nil), "geo.BoundingBox")
	proto.RegisterType((*LinearRing)(nil), "geo.LinearRing")
	proto.RegisterType((*Polygon)(nil), "geo.Polygon")
}

func init() { proto.RegisterFile("geo/geo.proto", fileDescriptor_d1967117b7b016d8) }

var fileDescriptor_d1967117b7b016d8 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x8b, 0x13, 0x41,
	0x10, 0x4d, 0x36, 0x46, 0x77, 0x3a, 0xae, 0xba, 0x0d, 0x62, 0x18, 0x58, 0x76, 0x19, 0x10, 0xf5,
	0xd2, 0x2d, 0x1b, 0x0f, 0xca, 0x5e, 0x76, 0xc7, 0x28, 0x1e, 0x3c, 0x2c, 0x83, 0x78, 0xf0, 0x22,
	0x9d, 0xa4, 0xd2, 0x69, 0x98, 0xe9, 0x0a, 0xdd, 0x35, 0x18, 0x4f, 0x2e, 0xf8, 0x2f, 0x3c, 0x7a,
	0xd8, 0xdf, 0xe1, 0xaf, 0xf0, 0xf7, 0x48, 0xe6, 0x5b, 0x5d, 0xc8, 0x1c, 0x06, 0xea, 0xf5, 0x7b,
	0xaf, 0x5f, 0x55, 0xd1, 0xec, 0x40, 0x03, 0x4a, 0x0d, 0x28, 0xd6, 0x0e, 0x09, 0xf9, 0x40, 0x03,
	0x86, 0x0f, 0x60, 0x43, 0x60, 0xbd, 0x41, 0xeb, 0x4b, 0x38, 0x3c, 0xcc, 0xad, 0x21, 0x2f, 0x8b,
	0x7f, 0x05, 0x3d, 0xec, 0x40, 0x9f, 0x61, 0x43, 0x15, 0x7c, 0x9c, 0x81, 0xf2, 0xb9, 0x83, 0x0c,
	0x2c, 0x79, 0xd9, 0x2d, 0x6a, 0x82, 0x46, 0xd4, 0x29, 0xc8, 0xa2, 0x9a, 0xe5, 0x4b, 0x49, 0x26,
	0x03, 0x4f, 0x2a, 0x5b, 0x97, 0x84, 0xe8, 0xf7, 0x1e, 0xbb, 0xfb, 0x1a, 0xd1, 0x2d, 0x8c, 0x55,
	0x04, 0xa7, 0x53, 0xfe, 0x96, 0xed, 0xa7, 0x8a, 0x0c, 0xe5, 0x0b, 0x18, 0xf7, 0x4f, 0xfa, 0x4f,
	0x47, 0xa7, 0xf7, 0x44, 0x99, 0x64, 0x0a, 0xda, 0x01, 0xf8, 0x38, 0xfc, 0x71, 0x7d, 0x14, 0x04,
	0xbd, 0xed, 0x77, 0xf5, 0xf1, 0xd7, 0xcf, 0x4e, 0x71, 0x9e, 0x34, 0x5a, 0xfe, 0x8e, 0x05, 0x29,
	0x5a, 0x5d, 0x1a, 0xed, 0xed, 0x34, 0x5a, 0x76, 0x8d, 0x96, 0xe7, 0x49, 0x2b, 0xe6, 0x13, 0x36,
	0x04, 0xe7, 0xd0, 0x8d, 0x07, 0x85, 0xcb, 0x23, 0xb1, 0x1d, 0x60, 0x37, 0xb3, 0x78, 0xb3, 0x3d,
	0x8e, 0x6f, 0x5d, 0x5d, 0x1f, 0xf5, 0x93, 0x92, 0x1b, 0x7e, 0x63, 0xc3, 0x02, 0xe5, 0xaf, 0x76,
	0xf6, 0x73, 0xd0, 0xc6, 0xe8, 0xf5, 0x7a, 0x9d, 0x16, 0xce, 0x76, 0xb7, 0xf0, 0x8f, 0xb6, 0xe5,
	0x47, 0xdf, 0xfb, 0x8c, 0x5d, 0xa2, 0x37, 0x64, 0xd0, 0xbe, 0x98, 0xf2, 0x97, 0x2c, 0x68, 0x46,
	0x5f, 0xe5, 0x08, 0x45, 0xb9, 0x1c, 0x51, 0x2f, 0x47, 0x7c, 0xa8, 0x19, 0x49, 0x4b, 0xe6, 0x67,
	0x6c, 0x7f, 0x5d, 0xf9, 0x54, 0x21, 0x8e, 0xc5, 0x5f, 0x9b, 0xae, 0x6f, 0x11, 0x17, 0x33, 0x8f,
	0x69, 0x4e, 0x90, 0x34, 0x82, 0x08, 0xd9, 0x28, 0xc6, 0xdc, 0x2e, 0x8c, 0xd5, 0x31, 0x6e, 0xb8,
	0x64, 0x81, 0xc7, 0x9c, 0x56, 0x5f, 0xc0, 0x53, 0x95, 0xe2, 0xf0, 0xbf, 0x71, 0x26, 0x2d, 0x67,
	0x2b, 0xb0, 0xe8, 0x68, 0x05, 0xca, 0x53, 0x75, 0xfb, 0x4d, 0x82, 0x86, 0x13, 0x5d, 0x30, 0xf6,
	0xde, 0x58, 0x50, 0x2e, 0x31, 0x56, 0xf3, 0x09, 0x1b, 0xcd, 0x1b, 0xa2, 0x1f, 0xf7, 0x4f, 0x06,
	0x37, 0x1b, 0x74, 0x59, 0xd1, 0x73, 0x76, 0xe7, 0x12, 0xd3, 0xaf, 0x1a, 0x2d, 0x7f, 0xcc, 0x86,
	0xce, 0x58, 0x5d, 0x2b, 0xef, 0x17, 0xca, 0xd6, 0x3f, 0x29, 0x4f, 0xe3, 0x67, 0x9f, 0x9e, 0x68,
	0x43, 0xab, 0x7c, 0x26, 0xe6, 0x98, 0x49, 0x65, 0x5c, 0xa6, 0xd6, 0xd2, 0x58, 0x02, 0xb7, 0x54,
	0x73, 0xf0, 0xd2, 0xbb, 0xb9, 0xd4, 0xc5, 0xc3, 0x9b, 0xdd, 0x2e, 0x86, 0x3d, 0xf9, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xa6, 0xb9, 0x0a, 0x10, 0x8a, 0x03, 0x00, 0x00,
}
