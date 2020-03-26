// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry/report.proto

package telemetry

import (
	fmt "fmt"
	_ "github.com/airmap/interfaces/src/go"
	measurements "github.com/airmap/interfaces/src/go/measurements"
	tracking "github.com/airmap/interfaces/src/go/tracking"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Report models a measurement at a given point in time.
type Report struct {
	Observed   *timestamp.Timestamp `protobuf:"bytes,1,opt,name=observed,proto3" json:"observed,omitempty"`
	Identities []*tracking.Identity `protobuf:"bytes,2,rep,name=identities,proto3" json:"identities,omitempty"`
	// details is a discriminated union of all measurement types.
	//
	// Types that are valid to be assigned to Details:
	//	*Report_Spatial_
	//	*Report_Atmospheric_
	//	*Report_System_
	Details              isReport_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_822b117a32fda6fc, []int{0}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
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

func (m *Report) GetIdentities() []*tracking.Identity {
	if m != nil {
		return m.Identities
	}
	return nil
}

type isReport_Details interface {
	isReport_Details()
}

type Report_Spatial_ struct {
	Spatial *Report_Spatial `protobuf:"bytes,3,opt,name=spatial,proto3,oneof"`
}

type Report_Atmospheric_ struct {
	Atmospheric *Report_Atmospheric `protobuf:"bytes,4,opt,name=atmospheric,proto3,oneof"`
}

type Report_System_ struct {
	System *Report_System `protobuf:"bytes,5,opt,name=system,proto3,oneof"`
}

func (*Report_Spatial_) isReport_Details() {}

func (*Report_Atmospheric_) isReport_Details() {}

func (*Report_System_) isReport_Details() {}

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

func (m *Report) GetSystem() *Report_System {
	if x, ok := m.GetDetails().(*Report_System_); ok {
		return x.System
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Report) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Report_Spatial_)(nil),
		(*Report_Atmospheric_)(nil),
		(*Report_System_)(nil),
	}
}

// Spatial bundles geospatial measurements.
type Report_Spatial struct {
	Position             *measurements.Position     `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Velocity             *measurements.Velocity     `protobuf:"bytes,2,opt,name=velocity,proto3" json:"velocity,omitempty"`
	Orientation          *measurements.Orientation  `protobuf:"bytes,3,opt,name=orientation,proto3" json:"orientation,omitempty"`
	Acceleration         *measurements.Acceleration `protobuf:"bytes,4,opt,name=acceleration,proto3" json:"acceleration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Report_Spatial) Reset()         { *m = Report_Spatial{} }
func (m *Report_Spatial) String() string { return proto.CompactTextString(m) }
func (*Report_Spatial) ProtoMessage()    {}
func (*Report_Spatial) Descriptor() ([]byte, []int) {
	return fileDescriptor_822b117a32fda6fc, []int{0, 0}
}

func (m *Report_Spatial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report_Spatial.Unmarshal(m, b)
}
func (m *Report_Spatial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report_Spatial.Marshal(b, m, deterministic)
}
func (m *Report_Spatial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report_Spatial.Merge(m, src)
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

func (m *Report_Spatial) GetAcceleration() *measurements.Acceleration {
	if m != nil {
		return m.Acceleration
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
	return fileDescriptor_822b117a32fda6fc, []int{0, 1}
}

func (m *Report_Atmospheric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report_Atmospheric.Unmarshal(m, b)
}
func (m *Report_Atmospheric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report_Atmospheric.Marshal(b, m, deterministic)
}
func (m *Report_Atmospheric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report_Atmospheric.Merge(m, src)
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

type Report_System struct {
	Battery              *Report_System_Battery `protobuf:"bytes,1,opt,name=battery,proto3" json:"battery,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Report_System) Reset()         { *m = Report_System{} }
func (m *Report_System) String() string { return proto.CompactTextString(m) }
func (*Report_System) ProtoMessage()    {}
func (*Report_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_822b117a32fda6fc, []int{0, 2}
}

func (m *Report_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report_System.Unmarshal(m, b)
}
func (m *Report_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report_System.Marshal(b, m, deterministic)
}
func (m *Report_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report_System.Merge(m, src)
}
func (m *Report_System) XXX_Size() int {
	return xxx_messageInfo_Report_System.Size(m)
}
func (m *Report_System) XXX_DiscardUnknown() {
	xxx_messageInfo_Report_System.DiscardUnknown(m)
}

var xxx_messageInfo_Report_System proto.InternalMessageInfo

func (m *Report_System) GetBattery() *Report_System_Battery {
	if m != nil {
		return m.Battery
	}
	return nil
}

type Report_System_Battery struct {
	Endurance            *duration.Duration `protobuf:"bytes,1,opt,name=endurance,proto3" json:"endurance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Report_System_Battery) Reset()         { *m = Report_System_Battery{} }
func (m *Report_System_Battery) String() string { return proto.CompactTextString(m) }
func (*Report_System_Battery) ProtoMessage()    {}
func (*Report_System_Battery) Descriptor() ([]byte, []int) {
	return fileDescriptor_822b117a32fda6fc, []int{0, 2, 0}
}

func (m *Report_System_Battery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report_System_Battery.Unmarshal(m, b)
}
func (m *Report_System_Battery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report_System_Battery.Marshal(b, m, deterministic)
}
func (m *Report_System_Battery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report_System_Battery.Merge(m, src)
}
func (m *Report_System_Battery) XXX_Size() int {
	return xxx_messageInfo_Report_System_Battery.Size(m)
}
func (m *Report_System_Battery) XXX_DiscardUnknown() {
	xxx_messageInfo_Report_System_Battery.DiscardUnknown(m)
}

var xxx_messageInfo_Report_System_Battery proto.InternalMessageInfo

func (m *Report_System_Battery) GetEndurance() *duration.Duration {
	if m != nil {
		return m.Endurance
	}
	return nil
}

func init() {
	proto.RegisterType((*Report)(nil), "telemetry.Report")
	proto.RegisterType((*Report_Spatial)(nil), "telemetry.Report.Spatial")
	proto.RegisterType((*Report_Atmospheric)(nil), "telemetry.Report.Atmospheric")
	proto.RegisterType((*Report_System)(nil), "telemetry.Report.System")
	proto.RegisterType((*Report_System_Battery)(nil), "telemetry.Report.System.Battery")
}

func init() { proto.RegisterFile("telemetry/report.proto", fileDescriptor_822b117a32fda6fc) }

var fileDescriptor_822b117a32fda6fc = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xdb, 0xae, 0x34, 0xad, 0xcb, 0x01, 0xf9, 0x30, 0xb2, 0x48, 0x63, 0x15, 0x07, 0xb4,
	0x0b, 0xb6, 0x54, 0xc4, 0x1f, 0xc1, 0xa9, 0xd5, 0x0e, 0xe5, 0xc4, 0x64, 0x26, 0x90, 0xb8, 0xb9,
	0xc9, 0xbb, 0xce, 0x22, 0xb1, 0x23, 0xfb, 0xed, 0xa0, 0xb7, 0x1d, 0xf9, 0x08, 0x7c, 0x02, 0x3e,
	0x14, 0x9f, 0x06, 0x25, 0x71, 0xd2, 0x84, 0x95, 0x03, 0xc7, 0xf8, 0xf9, 0x3d, 0xaf, 0x1f, 0x3f,
	0x8e, 0xc9, 0x31, 0x42, 0x0a, 0x19, 0xa0, 0xdd, 0x71, 0x0b, 0xb9, 0xb1, 0xc8, 0x72, 0x6b, 0xd0,
	0xd0, 0x49, 0xb3, 0x1e, 0x9d, 0x6d, 0x8c, 0xd9, 0xa4, 0xc0, 0x4b, 0x61, 0xbd, 0xbd, 0xe6, 0xa8,
	0x32, 0x70, 0x28, 0xb3, 0xbc, 0x62, 0xa3, 0x27, 0x7f, 0x03, 0xc9, 0xd6, 0x4a, 0x54, 0x46, 0x7b,
	0xfd, 0x11, 0x7c, 0x47, 0xd0, 0x4e, 0x19, 0xed, 0xfc, 0xca, 0x59, 0x06, 0xd2, 0x6d, 0x2d, 0x64,
	0xa0, 0xd1, 0xf1, 0xf6, 0x87, 0x07, 0x1e, 0xa3, 0x95, 0xf1, 0x57, 0xa5, 0x37, 0x5c, 0x25, 0xa0,
	0x51, 0xe1, 0xae, 0x12, 0x9e, 0xfe, 0x0c, 0xc8, 0x48, 0x94, 0x41, 0xe9, 0x2b, 0x32, 0x36, 0x6b,
	0x07, 0xf6, 0x16, 0x92, 0xb0, 0x3f, 0xeb, 0x9f, 0x4f, 0xe7, 0x11, 0xab, 0x92, 0xb0, 0x3a, 0x09,
	0xbb, 0xaa, 0xa3, 0x8a, 0x86, 0xa5, 0x73, 0x42, 0xfc, 0x50, 0x05, 0x2e, 0x1c, 0xcc, 0x8e, 0xce,
	0xa7, 0x73, 0xca, 0xea, 0x0d, 0xd9, 0x7b, 0xbf, 0xa1, 0x68, 0x51, 0xf4, 0x25, 0x09, 0x5c, 0x2e,
	0x51, 0xc9, 0x34, 0x3c, 0x2a, 0xb7, 0x3a, 0x61, 0x4d, 0x41, 0xac, 0xca, 0xc3, 0x3e, 0x56, 0xc0,
	0xaa, 0x27, 0x6a, 0x96, 0x2e, 0xc8, 0x54, 0x62, 0x66, 0x5c, 0x7e, 0x03, 0x56, 0xc5, 0xe1, 0xb0,
	0xb4, 0x9e, 0xde, 0xb7, 0x2e, 0xf6, 0xd0, 0xaa, 0x27, 0xda, 0x1e, 0x3a, 0x27, 0x23, 0xb7, 0x73,
	0x08, 0x59, 0xf8, 0xa0, 0x74, 0x87, 0x07, 0x36, 0x2e, 0xf5, 0x55, 0x4f, 0x78, 0x32, 0xba, 0x1b,
	0x90, 0xc0, 0xa7, 0xa1, 0x73, 0x32, 0xce, 0x8d, 0x53, 0xc5, 0x75, 0xf8, 0x96, 0x8e, 0x59, 0xa7,
	0xf0, 0x4b, 0xaf, 0x8a, 0x86, 0xa3, 0x6f, 0xc8, 0xf8, 0x16, 0x52, 0x13, 0x2b, 0xdc, 0x85, 0x83,
	0x43, 0x9e, 0x4f, 0x5e, 0x5d, 0x0e, 0xef, 0x7e, 0x9d, 0xf6, 0x45, 0x43, 0x17, 0x07, 0x36, 0x56,
	0x81, 0xc6, 0xf2, 0xfe, 0x9b, 0xae, 0x3a, 0xe6, 0x0f, 0x7b, 0xc0, 0xfb, 0xdb, 0x1e, 0x7a, 0x41,
	0x1e, 0xca, 0x38, 0x86, 0x14, 0xaa, 0x7f, 0xc8, 0x97, 0x16, 0x75, 0x67, 0x2c, 0x5a, 0x84, 0x1f,
	0xd2, 0x71, 0x45, 0xbf, 0xfb, 0x64, 0xba, 0xe8, 0xd4, 0xf8, 0xff, 0x35, 0x14, 0x1e, 0x0b, 0xae,
	0x60, 0x0e, 0xd7, 0x70, 0xe9, 0x55, 0xd1, 0x70, 0xf4, 0x1d, 0x99, 0x22, 0x64, 0x79, 0x11, 0xa3,
	0xb0, 0x1d, 0x2c, 0xe0, 0x6a, 0x0f, 0x88, 0x36, 0x4d, 0x9f, 0x91, 0xe1, 0x37, 0xa5, 0x13, 0x7f,
	0x64, 0xda, 0x75, 0x7d, 0x56, 0x3a, 0x11, 0xa5, 0x1e, 0xfd, 0xe8, 0x93, 0x51, 0x75, 0xe9, 0xf4,
	0x2d, 0x09, 0xd6, 0x12, 0x11, 0xec, 0xce, 0x1f, 0x6b, 0xf6, 0xaf, 0xff, 0x83, 0x2d, 0x2b, 0x4e,
	0xd4, 0x86, 0x68, 0x49, 0x02, 0xbf, 0x46, 0x5f, 0x93, 0x09, 0xe8, 0xe2, 0xd9, 0xea, 0x18, 0xfc,
	0xa0, 0x93, 0x7b, 0x8f, 0xe9, 0xc2, 0x3f, 0x6b, 0xb1, 0x67, 0x97, 0x13, 0x12, 0x24, 0x80, 0x52,
	0xa5, 0x6e, 0xc9, 0xbf, 0x3c, 0xdf, 0x28, 0xbc, 0xd9, 0xae, 0x59, 0x6c, 0x32, 0x2e, 0x95, 0xcd,
	0x64, 0xce, 0x95, 0x46, 0xb0, 0xd7, 0x32, 0x06, 0xc7, 0x9d, 0x8d, 0xf9, 0xc6, 0xf0, 0x26, 0xde,
	0x7a, 0x54, 0x4e, 0x7e, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x09, 0xbd, 0x87, 0xba, 0x84, 0x04,
	0x00, 0x00,
}
