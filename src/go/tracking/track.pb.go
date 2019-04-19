// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracking/track.proto

package tracking

import (
	fmt "fmt"
	_ "github.com/airmap/interfaces/src/go"
	measurements "github.com/airmap/interfaces/src/go/measurements"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// Track models a single track observation.
type Track struct {
	Observed             *timestamp.Timestamp      `protobuf:"bytes,1,opt,name=observed,proto3" json:"observed,omitempty"`
	Position             *measurements.Position    `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Velocity             *measurements.Velocity    `protobuf:"bytes,3,opt,name=velocity,proto3" json:"velocity,omitempty"`
	Orientation          *measurements.Orientation `protobuf:"bytes,4,opt,name=orientation,proto3" json:"orientation,omitempty"`
	Emitter              Emitter                   `protobuf:"varint,5,opt,name=emitter,proto3,enum=tracking.Emitter" json:"emitter,omitempty"`
	Sensor               *Sensor                   `protobuf:"bytes,6,opt,name=sensor,proto3" json:"sensor,omitempty"`
	Identities           []*Identity               `protobuf:"bytes,7,rep,name=identities,proto3" json:"identities,omitempty"`
	Details              []*any.Any                `protobuf:"bytes,1024,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Track) Reset()         { *m = Track{} }
func (m *Track) String() string { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()    {}
func (*Track) Descriptor() ([]byte, []int) {
	return fileDescriptor_560f96cb68e45a5e, []int{0}
}

func (m *Track) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Track.Unmarshal(m, b)
}
func (m *Track) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Track.Marshal(b, m, deterministic)
}
func (m *Track) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Track.Merge(m, src)
}
func (m *Track) XXX_Size() int {
	return xxx_messageInfo_Track.Size(m)
}
func (m *Track) XXX_DiscardUnknown() {
	xxx_messageInfo_Track.DiscardUnknown(m)
}

var xxx_messageInfo_Track proto.InternalMessageInfo

func (m *Track) GetObserved() *timestamp.Timestamp {
	if m != nil {
		return m.Observed
	}
	return nil
}

func (m *Track) GetPosition() *measurements.Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Track) GetVelocity() *measurements.Velocity {
	if m != nil {
		return m.Velocity
	}
	return nil
}

func (m *Track) GetOrientation() *measurements.Orientation {
	if m != nil {
		return m.Orientation
	}
	return nil
}

func (m *Track) GetEmitter() Emitter {
	if m != nil {
		return m.Emitter
	}
	return Emitter_UNKNOWN_EMITTER
}

func (m *Track) GetSensor() *Sensor {
	if m != nil {
		return m.Sensor
	}
	return nil
}

func (m *Track) GetIdentities() []*Identity {
	if m != nil {
		return m.Identities
	}
	return nil
}

func (m *Track) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

// Batch models a batch of tracks.
type Track_Batch struct {
	Tracks               []*Track `protobuf:"bytes,1,rep,name=tracks,proto3" json:"tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Track_Batch) Reset()         { *m = Track_Batch{} }
func (m *Track_Batch) String() string { return proto.CompactTextString(m) }
func (*Track_Batch) ProtoMessage()    {}
func (*Track_Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_560f96cb68e45a5e, []int{0, 0}
}

func (m *Track_Batch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Track_Batch.Unmarshal(m, b)
}
func (m *Track_Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Track_Batch.Marshal(b, m, deterministic)
}
func (m *Track_Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Track_Batch.Merge(m, src)
}
func (m *Track_Batch) XXX_Size() int {
	return xxx_messageInfo_Track_Batch.Size(m)
}
func (m *Track_Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Track_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Track_Batch proto.InternalMessageInfo

func (m *Track_Batch) GetTracks() []*Track {
	if m != nil {
		return m.Tracks
	}
	return nil
}

func init() {
	proto.RegisterType((*Track)(nil), "tracking.Track")
	proto.RegisterType((*Track_Batch)(nil), "tracking.Track.Batch")
}

func init() { proto.RegisterFile("tracking/track.proto", fileDescriptor_560f96cb68e45a5e) }

var fileDescriptor_560f96cb68e45a5e = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x15, 0xba, 0xa6, 0x91, 0x2b, 0xc1, 0xb0, 0xa6, 0xe1, 0x45, 0x42, 0x54, 0x5c, 0xc8,
	0x01, 0xd9, 0x10, 0x04, 0xe2, 0xba, 0x4a, 0x1c, 0x38, 0x81, 0xc2, 0xc4, 0x81, 0x9b, 0x93, 0xbe,
	0x65, 0x16, 0x8d, 0x1d, 0xd9, 0xaf, 0x13, 0xbb, 0xed, 0xd3, 0xf0, 0xe1, 0xf8, 0x14, 0x08, 0xc7,
	0x4e, 0x5a, 0xba, 0x5b, 0xf2, 0xfe, 0xbf, 0x9f, 0xde, 0xd3, 0xdf, 0xe4, 0x0c, 0xad, 0x6c, 0x7e,
	0x2a, 0xdd, 0x0a, 0xff, 0xc1, 0x7b, 0x6b, 0xd0, 0xd0, 0x2c, 0x4e, 0xf3, 0x8b, 0xd6, 0x98, 0x76,
	0x0b, 0xc2, 0xcf, 0xeb, 0xdd, 0xb5, 0x90, 0xfa, 0x6e, 0x80, 0xf2, 0x17, 0xff, 0x47, 0xa8, 0x3a,
	0x70, 0x28, 0xbb, 0x3e, 0x00, 0xa7, 0xf0, 0x0b, 0x41, 0x3b, 0x65, 0xb4, 0x8b, 0x4a, 0x07, 0xd2,
	0xed, 0x2c, 0x74, 0xa0, 0xd1, 0x89, 0xfd, 0x9f, 0x00, 0x9c, 0x8f, 0xe7, 0x40, 0xa7, 0x10, 0xc1,
	0x1e, 0xcd, 0x1d, 0x68, 0x67, 0x6c, 0xe4, 0x9f, 0x8d, 0x73, 0xb5, 0x01, 0x8d, 0x0a, 0xc3, 0x71,
	0x2f, 0xff, 0xcc, 0xc8, 0xfc, 0xea, 0x5f, 0x46, 0x3f, 0x90, 0xcc, 0xd4, 0x0e, 0xec, 0x2d, 0x6c,
	0x58, 0xb2, 0x4a, 0x8a, 0x65, 0x99, 0xf3, 0xe1, 0x72, 0x1e, 0x2f, 0xe7, 0x57, 0xf1, 0xf2, 0x6a,
	0x64, 0x69, 0x49, 0xb2, 0xde, 0x38, 0x85, 0xca, 0x68, 0xf6, 0xc8, 0x7b, 0xe7, 0xfc, 0xe0, 0xe2,
	0xaf, 0x21, 0xad, 0x46, 0x8e, 0x7e, 0x24, 0xd9, 0x2d, 0x6c, 0x4d, 0xa3, 0xf0, 0x8e, 0xcd, 0x1e,
	0x72, 0xbe, 0x87, 0x74, 0x7d, 0x72, 0xff, 0xfb, 0x79, 0x52, 0x8d, 0x34, 0xbd, 0x24, 0x4b, 0x63,
	0x15, 0x68, 0x94, 0x7e, 0xe1, 0x89, 0x97, 0x2f, 0x0e, 0xe5, 0x2f, 0x13, 0x10, 0xfc, 0x7d, 0x87,
	0xbe, 0x25, 0x8b, 0x50, 0x1a, 0x9b, 0xaf, 0x92, 0xe2, 0x71, 0xf9, 0x94, 0xc7, 0x76, 0xf8, 0xa7,
	0x21, 0x08, 0x5a, 0xe4, 0x68, 0x41, 0xd2, 0xa1, 0x4f, 0x96, 0xfa, 0x85, 0xa7, 0x93, 0xf1, 0xcd,
	0xcf, 0xab, 0x90, 0xd3, 0x92, 0x90, 0xd0, 0xb0, 0x02, 0xc7, 0x16, 0xab, 0x59, 0xb1, 0x2c, 0xe9,
	0x44, 0x7f, 0x0e, 0xed, 0x57, 0x7b, 0x14, 0x7d, 0x4f, 0x16, 0x1b, 0x40, 0xa9, 0xb6, 0x8e, 0xdd,
	0x67, 0xde, 0x38, 0x3b, 0x6a, 0xfe, 0x52, 0xc7, 0x2e, 0x22, 0x9b, 0xbf, 0x21, 0xf3, 0xb5, 0xc4,
	0xe6, 0x86, 0xbe, 0x22, 0xa9, 0x5f, 0xe0, 0x58, 0xe2, 0xed, 0x27, 0xd3, 0x3e, 0xff, 0xb4, 0x55,
	0x88, 0xd7, 0xfc, 0xc7, 0xeb, 0x56, 0xe1, 0xcd, 0xae, 0xe6, 0x8d, 0xe9, 0x84, 0x54, 0xb6, 0x93,
	0xbd, 0x50, 0x1a, 0xc1, 0x5e, 0xcb, 0x06, 0x9c, 0x70, 0xb6, 0x11, 0xad, 0x11, 0xd1, 0xae, 0x53,
	0xbf, 0xfe, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x17, 0x1c, 0x11, 0xfd, 0x02, 0x00,
	0x00,
}
