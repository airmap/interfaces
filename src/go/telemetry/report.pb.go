// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry/report.proto

package telemetry // import "github.com/airmap/interfaces/src/go/telemetry"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/airmap/interfaces/src/go"
import measurements "github.com/airmap/interfaces/src/go/measurements"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Report models a measurement at a given point in time.
type Report struct {
	Observed *timestamp.Timestamp `protobuf:"bytes,1,opt,name=observed,proto3" json:"observed,omitempty"`
	// details is a discriminated union of all measurement types.
	//
	// Types that are valid to be assigned to Details:
	//	*Report_Spatial_
	//	*Report_Atmospheric_
	Details              isReport_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_report_5ee1b8320ec01c2a, []int{0}
}
func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (dst *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(dst, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetObserved() *timestamp.Timestamp {
	if m != nil {
		return m.Observed
	}
	return nil
}

type isReport_Details interface {
	isReport_Details()
}

type Report_Spatial_ struct {
	Spatial *Report_Spatial `protobuf:"bytes,2,opt,name=spatial,proto3,oneof"`
}

type Report_Atmospheric_ struct {
	Atmospheric *Report_Atmospheric `protobuf:"bytes,3,opt,name=atmospheric,proto3,oneof"`
}

func (*Report_Spatial_) isReport_Details() {}

func (*Report_Atmospheric_) isReport_Details() {}

func (m *Report) GetDetails() isReport_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Report) GetSpatial() *Report_Spatial {
	if x, ok := m.GetDetails().(*Report_Spatial_); ok {
		return x.Spatial
	}
	return nil
}

func (m *Report) GetAtmospheric() *Report_Atmospheric {
	if x, ok := m.GetDetails().(*Report_Atmospheric_); ok {
		return x.Atmospheric
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Report) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Report_OneofMarshaler, _Report_OneofUnmarshaler, _Report_OneofSizer, []interface{}{
		(*Report_Spatial_)(nil),
		(*Report_Atmospheric_)(nil),
	}
}

func _Report_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Report)
	// details
	switch x := m.Details.(type) {
	case *Report_Spatial_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Spatial); err != nil {
			return err
		}
	case *Report_Atmospheric_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Atmospheric); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Report.Details has unexpected type %T", x)
	}
	return nil
}

func _Report_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Report)
	switch tag {
	case 2: // details.spatial
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Report_Spatial)
		err := b.DecodeMessage(msg)
		m.Details = &Report_Spatial_{msg}
		return true, err
	case 3: // details.atmospheric
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Report_Atmospheric)
		err := b.DecodeMessage(msg)
		m.Details = &Report_Atmospheric_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Report_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Report)
	// details
	switch x := m.Details.(type) {
	case *Report_Spatial_:
		s := proto.Size(x.Spatial)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Report_Atmospheric_:
		s := proto.Size(x.Atmospheric)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Spatial bundles geospatial measurements.
type Report_Spatial struct {
	Position             *measurements.Position    `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Velocity             *measurements.Velocity    `protobuf:"bytes,2,opt,name=velocity,proto3" json:"velocity,omitempty"`
	Orientation          *measurements.Orientation `protobuf:"bytes,3,opt,name=orientation,proto3" json:"orientation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Report_Spatial) Reset()         { *m = Report_Spatial{} }
func (m *Report_Spatial) String() string { return proto.CompactTextString(m) }
func (*Report_Spatial) ProtoMessage()    {}
func (*Report_Spatial) Descriptor() ([]byte, []int) {
	return fileDescriptor_report_5ee1b8320ec01c2a, []int{0, 0}
}
func (m *Report_Spatial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report_Spatial.Unmarshal(m, b)
}
func (m *Report_Spatial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report_Spatial.Marshal(b, m, deterministic)
}
func (dst *Report_Spatial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report_Spatial.Merge(dst, src)
}
func (m *Report_Spatial) XXX_Size() int {
	return xxx_messageInfo_Report_Spatial.Size(m)
}
func (m *Report_Spatial) XXX_DiscardUnknown() {
	xxx_messageInfo_Report_Spatial.DiscardUnknown(m)
}

var xxx_messageInfo_Report_Spatial proto.InternalMessageInfo

func (m *Report_Spatial) GetPosition() *measurements.Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Report_Spatial) GetVelocity() *measurements.Velocity {
	if m != nil {
		return m.Velocity
	}
	return nil
}

func (m *Report_Spatial) GetOrientation() *measurements.Orientation {
	if m != nil {
		return m.Orientation
	}
	return nil
}

// Atmospheric bundles atmospheric measurements at a given position.
type Report_Atmospheric struct {
	Position             *measurements.Position    `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Pressure             *measurements.Pressure    `protobuf:"bytes,2,opt,name=pressure,proto3" json:"pressure,omitempty"`
	Temperature          *measurements.Temperature `protobuf:"bytes,3,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Wind                 *measurements.Wind        `protobuf:"bytes,4,opt,name=wind,proto3" json:"wind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Report_Atmospheric) Reset()         { *m = Report_Atmospheric{} }
func (m *Report_Atmospheric) String() string { return proto.CompactTextString(m) }
func (*Report_Atmospheric) ProtoMessage()    {}
func (*Report_Atmospheric) Descriptor() ([]byte, []int) {
	return fileDescriptor_report_5ee1b8320ec01c2a, []int{0, 1}
}
func (m *Report_Atmospheric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report_Atmospheric.Unmarshal(m, b)
}
func (m *Report_Atmospheric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report_Atmospheric.Marshal(b, m, deterministic)
}
func (dst *Report_Atmospheric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report_Atmospheric.Merge(dst, src)
}
func (m *Report_Atmospheric) XXX_Size() int {
	return xxx_messageInfo_Report_Atmospheric.Size(m)
}
func (m *Report_Atmospheric) XXX_DiscardUnknown() {
	xxx_messageInfo_Report_Atmospheric.DiscardUnknown(m)
}

var xxx_messageInfo_Report_Atmospheric proto.InternalMessageInfo

func (m *Report_Atmospheric) GetPosition() *measurements.Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Report_Atmospheric) GetPressure() *measurements.Pressure {
	if m != nil {
		return m.Pressure
	}
	return nil
}

func (m *Report_Atmospheric) GetTemperature() *measurements.Temperature {
	if m != nil {
		return m.Temperature
	}
	return nil
}

func (m *Report_Atmospheric) GetWind() *measurements.Wind {
	if m != nil {
		return m.Wind
	}
	return nil
}

func init() {
	proto.RegisterType((*Report)(nil), "telemetry.Report")
	proto.RegisterType((*Report_Spatial)(nil), "telemetry.Report.Spatial")
	proto.RegisterType((*Report_Atmospheric)(nil), "telemetry.Report.Atmospheric")
}

func init() { proto.RegisterFile("telemetry/report.proto", fileDescriptor_report_5ee1b8320ec01c2a) }

var fileDescriptor_report_5ee1b8320ec01c2a = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x6e, 0xd4, 0x30,
	0x14, 0x85, 0x1b, 0x1a, 0x75, 0xa6, 0xce, 0x06, 0x79, 0x51, 0x85, 0x48, 0x55, 0x11, 0x0b, 0xc4,
	0x06, 0x5b, 0x2a, 0x02, 0x21, 0xb1, 0xea, 0xac, 0xba, 0xa3, 0x32, 0x15, 0x48, 0xec, 0x9c, 0xe4,
	0x36, 0xb5, 0x14, 0xff, 0xc8, 0xbe, 0x53, 0x98, 0x1d, 0x4f, 0xc3, 0x43, 0xf0, 0x28, 0x6c, 0x78,
	0x15, 0x14, 0xc7, 0x93, 0xc9, 0xc0, 0xac, 0xba, 0x8c, 0xef, 0xf7, 0xf9, 0x9e, 0xa3, 0x98, 0x9c,
	0x21, 0xf4, 0xa0, 0x01, 0xfd, 0x86, 0x7b, 0x70, 0xd6, 0x23, 0x73, 0xde, 0xa2, 0xa5, 0xa7, 0xd3,
	0x79, 0x75, 0xd1, 0x59, 0xdb, 0xf5, 0xc0, 0xe3, 0xa0, 0x5e, 0xdf, 0x71, 0x54, 0x1a, 0x02, 0x4a,
	0xed, 0x46, 0xb6, 0x7a, 0x0a, 0xdf, 0x11, 0x4c, 0x50, 0xd6, 0x84, 0x74, 0x72, 0xa1, 0x41, 0x86,
	0xb5, 0x07, 0x0d, 0x06, 0x03, 0x9f, 0x7f, 0x8c, 0xc0, 0x8b, 0x3f, 0x39, 0x39, 0x11, 0x71, 0x1f,
	0x7d, 0x47, 0x96, 0xb6, 0x0e, 0xe0, 0x1f, 0xa0, 0x2d, 0xb3, 0xe7, 0xd9, 0xab, 0xe2, 0xb2, 0x62,
	0xe3, 0x46, 0xb6, 0xdd, 0xc8, 0x6e, 0xb7, 0x1b, 0xc5, 0xc4, 0xd2, 0xb7, 0x64, 0x11, 0x9c, 0x44,
	0x25, 0xfb, 0xf2, 0x49, 0xd4, 0x9e, 0xb1, 0x29, 0x33, 0x1b, 0xef, 0x66, 0x9f, 0x46, 0xe0, 0xfa,
	0x48, 0x6c, 0x59, 0x7a, 0x45, 0x0a, 0x89, 0xda, 0x06, 0x77, 0x0f, 0x5e, 0x35, 0xe5, 0x71, 0x54,
	0xcf, 0xff, 0x57, 0xaf, 0x76, 0xd0, 0xf5, 0x91, 0x98, 0x3b, 0xd5, 0xaf, 0x8c, 0x2c, 0xd2, 0xcd,
	0xf4, 0x92, 0x2c, 0x9d, 0x0d, 0x0a, 0x95, 0x35, 0x29, 0xfd, 0x19, 0xdb, 0xeb, 0x7b, 0x93, 0xa6,
	0x62, 0xe2, 0xe8, 0x7b, 0xb2, 0x7c, 0x80, 0xde, 0x36, 0x0a, 0x37, 0x29, 0xfa, 0x3f, 0xce, 0xe7,
	0x34, 0x5d, 0xe5, 0x3f, 0x7e, 0x9e, 0x67, 0x62, 0xa2, 0x87, 0xf0, 0xd6, 0x2b, 0x30, 0x28, 0xe3,
	0xc2, 0xe3, 0xd4, 0x7b, 0x4f, 0xfe, 0xb8, 0x03, 0x92, 0x3f, 0x77, 0xaa, 0xdf, 0x19, 0x29, 0x66,
	0xdd, 0x1e, 0x55, 0x60, 0x70, 0x3c, 0x84, 0x81, 0x39, 0x5c, 0xe0, 0x26, 0x4d, 0xc5, 0xc4, 0xd1,
	0x0f, 0xa4, 0x40, 0xd0, 0x0e, 0xbc, 0xc4, 0x41, 0x3b, 0x18, 0xfd, 0x76, 0x07, 0x88, 0x39, 0x4d,
	0x5f, 0x92, 0xfc, 0x9b, 0x32, 0x6d, 0x99, 0x47, 0x8b, 0xee, 0x5b, 0x5f, 0x94, 0x69, 0x45, 0x9c,
	0xaf, 0x4e, 0xc9, 0xa2, 0x05, 0x94, 0xaa, 0x0f, 0x2b, 0xfe, 0xf5, 0x75, 0xa7, 0xf0, 0x7e, 0x5d,
	0xb3, 0xc6, 0x6a, 0x2e, 0x95, 0xd7, 0xd2, 0x71, 0x65, 0x10, 0xfc, 0x9d, 0x6c, 0x20, 0xf0, 0xe0,
	0x1b, 0xde, 0x59, 0x3e, 0xfd, 0xf7, 0xfa, 0x24, 0xbe, 0xb6, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0xe4, 0x2b, 0x5b, 0x12, 0x03, 0x00, 0x00,
}
