// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracking/collector.proto

package tracking // import "github.com/airmap/interfaces/src/go/tracking"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/airmap/interfaces/src/go"
import system "github.com/airmap/interfaces/src/go/system"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ConnectProcessorParameters models configuration parameters for processor streams
type ConnectProcessorParameters struct {
	EnableProjection     bool     `protobuf:"varint,1,opt,name=enable_projection,json=enableProjection,proto3" json:"enable_projection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectProcessorParameters) Reset()         { *m = ConnectProcessorParameters{} }
func (m *ConnectProcessorParameters) String() string { return proto.CompactTextString(m) }
func (*ConnectProcessorParameters) ProtoMessage()    {}
func (*ConnectProcessorParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_collector_f50ae4f6c1836da5, []int{0}
}
func (m *ConnectProcessorParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectProcessorParameters.Unmarshal(m, b)
}
func (m *ConnectProcessorParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectProcessorParameters.Marshal(b, m, deterministic)
}
func (dst *ConnectProcessorParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectProcessorParameters.Merge(dst, src)
}
func (m *ConnectProcessorParameters) XXX_Size() int {
	return xxx_messageInfo_ConnectProcessorParameters.Size(m)
}
func (m *ConnectProcessorParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectProcessorParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectProcessorParameters proto.InternalMessageInfo

func (m *ConnectProcessorParameters) GetEnableProjection() bool {
	if m != nil {
		return m.EnableProjection
	}
	return false
}

// Update bundles types used in the exchange of tracks.
type Update struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Update) Reset()         { *m = Update{} }
func (m *Update) String() string { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()    {}
func (*Update) Descriptor() ([]byte, []int) {
	return fileDescriptor_collector_f50ae4f6c1836da5, []int{1}
}
func (m *Update) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update.Unmarshal(m, b)
}
func (m *Update) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update.Marshal(b, m, deterministic)
}
func (dst *Update) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update.Merge(dst, src)
}
func (m *Update) XXX_Size() int {
	return xxx_messageInfo_Update.Size(m)
}
func (m *Update) XXX_DiscardUnknown() {
	xxx_messageInfo_Update.DiscardUnknown(m)
}

var xxx_messageInfo_Update proto.InternalMessageInfo

// FromProvider wraps messages being sent by a provider to a traffic collector.
type Update_FromProvider struct {
	// Types that are valid to be assigned to Details:
	//	*Update_FromProvider_Status
	//	*Update_FromProvider_Batch
	Details              isUpdate_FromProvider_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Update_FromProvider) Reset()         { *m = Update_FromProvider{} }
func (m *Update_FromProvider) String() string { return proto.CompactTextString(m) }
func (*Update_FromProvider) ProtoMessage()    {}
func (*Update_FromProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_collector_f50ae4f6c1836da5, []int{1, 0}
}
func (m *Update_FromProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_FromProvider.Unmarshal(m, b)
}
func (m *Update_FromProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_FromProvider.Marshal(b, m, deterministic)
}
func (dst *Update_FromProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_FromProvider.Merge(dst, src)
}
func (m *Update_FromProvider) XXX_Size() int {
	return xxx_messageInfo_Update_FromProvider.Size(m)
}
func (m *Update_FromProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_Update_FromProvider.DiscardUnknown(m)
}

var xxx_messageInfo_Update_FromProvider proto.InternalMessageInfo

type isUpdate_FromProvider_Details interface {
	isUpdate_FromProvider_Details()
}

type Update_FromProvider_Status struct {
	Status *system.Status `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type Update_FromProvider_Batch struct {
	Batch *Track_Batch `protobuf:"bytes,2,opt,name=batch,proto3,oneof"`
}

func (*Update_FromProvider_Status) isUpdate_FromProvider_Details() {}

func (*Update_FromProvider_Batch) isUpdate_FromProvider_Details() {}

func (m *Update_FromProvider) GetDetails() isUpdate_FromProvider_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Update_FromProvider) GetStatus() *system.Status {
	if x, ok := m.GetDetails().(*Update_FromProvider_Status); ok {
		return x.Status
	}
	return nil
}

func (m *Update_FromProvider) GetBatch() *Track_Batch {
	if x, ok := m.GetDetails().(*Update_FromProvider_Batch); ok {
		return x.Batch
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Update_FromProvider) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Update_FromProvider_OneofMarshaler, _Update_FromProvider_OneofUnmarshaler, _Update_FromProvider_OneofSizer, []interface{}{
		(*Update_FromProvider_Status)(nil),
		(*Update_FromProvider_Batch)(nil),
	}
}

func _Update_FromProvider_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Update_FromProvider)
	// details
	switch x := m.Details.(type) {
	case *Update_FromProvider_Status:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Status); err != nil {
			return err
		}
	case *Update_FromProvider_Batch:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Batch); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Update_FromProvider.Details has unexpected type %T", x)
	}
	return nil
}

func _Update_FromProvider_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Update_FromProvider)
	switch tag {
	case 1: // details.status
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(system.Status)
		err := b.DecodeMessage(msg)
		m.Details = &Update_FromProvider_Status{msg}
		return true, err
	case 2: // details.batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Track_Batch)
		err := b.DecodeMessage(msg)
		m.Details = &Update_FromProvider_Batch{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Update_FromProvider_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Update_FromProvider)
	// details
	switch x := m.Details.(type) {
	case *Update_FromProvider_Status:
		s := proto.Size(x.Status)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_FromProvider_Batch:
		s := proto.Size(x.Batch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// ToProvider wraps messages being sent from a collector back to a provider.
type Update_ToProvider struct {
	// Types that are valid to be assigned to Details:
	//	*Update_ToProvider_Status
	//	*Update_ToProvider_Ack
	Details              isUpdate_ToProvider_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Update_ToProvider) Reset()         { *m = Update_ToProvider{} }
func (m *Update_ToProvider) String() string { return proto.CompactTextString(m) }
func (*Update_ToProvider) ProtoMessage()    {}
func (*Update_ToProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_collector_f50ae4f6c1836da5, []int{1, 1}
}
func (m *Update_ToProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_ToProvider.Unmarshal(m, b)
}
func (m *Update_ToProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_ToProvider.Marshal(b, m, deterministic)
}
func (dst *Update_ToProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_ToProvider.Merge(dst, src)
}
func (m *Update_ToProvider) XXX_Size() int {
	return xxx_messageInfo_Update_ToProvider.Size(m)
}
func (m *Update_ToProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_Update_ToProvider.DiscardUnknown(m)
}

var xxx_messageInfo_Update_ToProvider proto.InternalMessageInfo

type isUpdate_ToProvider_Details interface {
	isUpdate_ToProvider_Details()
}

type Update_ToProvider_Status struct {
	Status *system.Status `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type Update_ToProvider_Ack struct {
	Ack *system.Ack `protobuf:"bytes,2,opt,name=ack,proto3,oneof"`
}

func (*Update_ToProvider_Status) isUpdate_ToProvider_Details() {}

func (*Update_ToProvider_Ack) isUpdate_ToProvider_Details() {}

func (m *Update_ToProvider) GetDetails() isUpdate_ToProvider_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Update_ToProvider) GetStatus() *system.Status {
	if x, ok := m.GetDetails().(*Update_ToProvider_Status); ok {
		return x.Status
	}
	return nil
}

func (m *Update_ToProvider) GetAck() *system.Ack {
	if x, ok := m.GetDetails().(*Update_ToProvider_Ack); ok {
		return x.Ack
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Update_ToProvider) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Update_ToProvider_OneofMarshaler, _Update_ToProvider_OneofUnmarshaler, _Update_ToProvider_OneofSizer, []interface{}{
		(*Update_ToProvider_Status)(nil),
		(*Update_ToProvider_Ack)(nil),
	}
}

func _Update_ToProvider_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Update_ToProvider)
	// details
	switch x := m.Details.(type) {
	case *Update_ToProvider_Status:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Status); err != nil {
			return err
		}
	case *Update_ToProvider_Ack:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ack); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Update_ToProvider.Details has unexpected type %T", x)
	}
	return nil
}

func _Update_ToProvider_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Update_ToProvider)
	switch tag {
	case 1: // details.status
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(system.Status)
		err := b.DecodeMessage(msg)
		m.Details = &Update_ToProvider_Status{msg}
		return true, err
	case 2: // details.ack
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(system.Ack)
		err := b.DecodeMessage(msg)
		m.Details = &Update_ToProvider_Ack{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Update_ToProvider_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Update_ToProvider)
	// details
	switch x := m.Details.(type) {
	case *Update_ToProvider_Status:
		s := proto.Size(x.Status)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_ToProvider_Ack:
		s := proto.Size(x.Ack)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// ToProcessor wraps messages being sent by a collector to a processor.
type Update_ToProcessor struct {
	// Types that are valid to be assigned to Details:
	//	*Update_ToProcessor_Status
	//	*Update_ToProcessor_Batch
	Details              isUpdate_ToProcessor_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Update_ToProcessor) Reset()         { *m = Update_ToProcessor{} }
func (m *Update_ToProcessor) String() string { return proto.CompactTextString(m) }
func (*Update_ToProcessor) ProtoMessage()    {}
func (*Update_ToProcessor) Descriptor() ([]byte, []int) {
	return fileDescriptor_collector_f50ae4f6c1836da5, []int{1, 2}
}
func (m *Update_ToProcessor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_ToProcessor.Unmarshal(m, b)
}
func (m *Update_ToProcessor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_ToProcessor.Marshal(b, m, deterministic)
}
func (dst *Update_ToProcessor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_ToProcessor.Merge(dst, src)
}
func (m *Update_ToProcessor) XXX_Size() int {
	return xxx_messageInfo_Update_ToProcessor.Size(m)
}
func (m *Update_ToProcessor) XXX_DiscardUnknown() {
	xxx_messageInfo_Update_ToProcessor.DiscardUnknown(m)
}

var xxx_messageInfo_Update_ToProcessor proto.InternalMessageInfo

type isUpdate_ToProcessor_Details interface {
	isUpdate_ToProcessor_Details()
}

type Update_ToProcessor_Status struct {
	Status *system.Status `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type Update_ToProcessor_Batch struct {
	Batch *Track_Batch `protobuf:"bytes,2,opt,name=batch,proto3,oneof"`
}

func (*Update_ToProcessor_Status) isUpdate_ToProcessor_Details() {}

func (*Update_ToProcessor_Batch) isUpdate_ToProcessor_Details() {}

func (m *Update_ToProcessor) GetDetails() isUpdate_ToProcessor_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Update_ToProcessor) GetStatus() *system.Status {
	if x, ok := m.GetDetails().(*Update_ToProcessor_Status); ok {
		return x.Status
	}
	return nil
}

func (m *Update_ToProcessor) GetBatch() *Track_Batch {
	if x, ok := m.GetDetails().(*Update_ToProcessor_Batch); ok {
		return x.Batch
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Update_ToProcessor) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Update_ToProcessor_OneofMarshaler, _Update_ToProcessor_OneofUnmarshaler, _Update_ToProcessor_OneofSizer, []interface{}{
		(*Update_ToProcessor_Status)(nil),
		(*Update_ToProcessor_Batch)(nil),
	}
}

func _Update_ToProcessor_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Update_ToProcessor)
	// details
	switch x := m.Details.(type) {
	case *Update_ToProcessor_Status:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Status); err != nil {
			return err
		}
	case *Update_ToProcessor_Batch:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Batch); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Update_ToProcessor.Details has unexpected type %T", x)
	}
	return nil
}

func _Update_ToProcessor_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Update_ToProcessor)
	switch tag {
	case 1: // details.status
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(system.Status)
		err := b.DecodeMessage(msg)
		m.Details = &Update_ToProcessor_Status{msg}
		return true, err
	case 2: // details.batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Track_Batch)
		err := b.DecodeMessage(msg)
		m.Details = &Update_ToProcessor_Batch{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Update_ToProcessor_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Update_ToProcessor)
	// details
	switch x := m.Details.(type) {
	case *Update_ToProcessor_Status:
		s := proto.Size(x.Status)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_ToProcessor_Batch:
		s := proto.Size(x.Batch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Update_FromProcessor struct {
	// Types that are valid to be assigned to Details:
	//	*Update_FromProcessor_Status
	//	*Update_FromProcessor_Ack
	//	*Update_FromProcessor_Params
	Details              isUpdate_FromProcessor_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Update_FromProcessor) Reset()         { *m = Update_FromProcessor{} }
func (m *Update_FromProcessor) String() string { return proto.CompactTextString(m) }
func (*Update_FromProcessor) ProtoMessage()    {}
func (*Update_FromProcessor) Descriptor() ([]byte, []int) {
	return fileDescriptor_collector_f50ae4f6c1836da5, []int{1, 3}
}
func (m *Update_FromProcessor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_FromProcessor.Unmarshal(m, b)
}
func (m *Update_FromProcessor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_FromProcessor.Marshal(b, m, deterministic)
}
func (dst *Update_FromProcessor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_FromProcessor.Merge(dst, src)
}
func (m *Update_FromProcessor) XXX_Size() int {
	return xxx_messageInfo_Update_FromProcessor.Size(m)
}
func (m *Update_FromProcessor) XXX_DiscardUnknown() {
	xxx_messageInfo_Update_FromProcessor.DiscardUnknown(m)
}

var xxx_messageInfo_Update_FromProcessor proto.InternalMessageInfo

type isUpdate_FromProcessor_Details interface {
	isUpdate_FromProcessor_Details()
}

type Update_FromProcessor_Status struct {
	Status *system.Status `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type Update_FromProcessor_Ack struct {
	Ack *system.Ack `protobuf:"bytes,2,opt,name=ack,proto3,oneof"`
}

type Update_FromProcessor_Params struct {
	Params *ConnectProcessorParameters `protobuf:"bytes,3,opt,name=params,proto3,oneof"`
}

func (*Update_FromProcessor_Status) isUpdate_FromProcessor_Details() {}

func (*Update_FromProcessor_Ack) isUpdate_FromProcessor_Details() {}

func (*Update_FromProcessor_Params) isUpdate_FromProcessor_Details() {}

func (m *Update_FromProcessor) GetDetails() isUpdate_FromProcessor_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Update_FromProcessor) GetStatus() *system.Status {
	if x, ok := m.GetDetails().(*Update_FromProcessor_Status); ok {
		return x.Status
	}
	return nil
}

func (m *Update_FromProcessor) GetAck() *system.Ack {
	if x, ok := m.GetDetails().(*Update_FromProcessor_Ack); ok {
		return x.Ack
	}
	return nil
}

func (m *Update_FromProcessor) GetParams() *ConnectProcessorParameters {
	if x, ok := m.GetDetails().(*Update_FromProcessor_Params); ok {
		return x.Params
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Update_FromProcessor) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Update_FromProcessor_OneofMarshaler, _Update_FromProcessor_OneofUnmarshaler, _Update_FromProcessor_OneofSizer, []interface{}{
		(*Update_FromProcessor_Status)(nil),
		(*Update_FromProcessor_Ack)(nil),
		(*Update_FromProcessor_Params)(nil),
	}
}

func _Update_FromProcessor_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Update_FromProcessor)
	// details
	switch x := m.Details.(type) {
	case *Update_FromProcessor_Status:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Status); err != nil {
			return err
		}
	case *Update_FromProcessor_Ack:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ack); err != nil {
			return err
		}
	case *Update_FromProcessor_Params:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Params); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Update_FromProcessor.Details has unexpected type %T", x)
	}
	return nil
}

func _Update_FromProcessor_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Update_FromProcessor)
	switch tag {
	case 1: // details.status
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(system.Status)
		err := b.DecodeMessage(msg)
		m.Details = &Update_FromProcessor_Status{msg}
		return true, err
	case 2: // details.ack
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(system.Ack)
		err := b.DecodeMessage(msg)
		m.Details = &Update_FromProcessor_Ack{msg}
		return true, err
	case 3: // details.params
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConnectProcessorParameters)
		err := b.DecodeMessage(msg)
		m.Details = &Update_FromProcessor_Params{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Update_FromProcessor_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Update_FromProcessor)
	// details
	switch x := m.Details.(type) {
	case *Update_FromProcessor_Status:
		s := proto.Size(x.Status)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_FromProcessor_Ack:
		s := proto.Size(x.Ack)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_FromProcessor_Params:
		s := proto.Size(x.Params)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ConnectProcessorParameters)(nil), "tracking.ConnectProcessorParameters")
	proto.RegisterType((*Update)(nil), "tracking.Update")
	proto.RegisterType((*Update_FromProvider)(nil), "tracking.Update.FromProvider")
	proto.RegisterType((*Update_ToProvider)(nil), "tracking.Update.ToProvider")
	proto.RegisterType((*Update_ToProcessor)(nil), "tracking.Update.ToProcessor")
	proto.RegisterType((*Update_FromProcessor)(nil), "tracking.Update.FromProcessor")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CollectorClient is the client API for Collector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectorClient interface {
	// ConnectProvider connects a stream of updates from a provider to a collector.
	ConnectProvider(ctx context.Context, opts ...grpc.CallOption) (Collector_ConnectProviderClient, error)
	// ConnectProcessor connects a stream of updates from a collector to a processor.
	ConnectProcessor(ctx context.Context, opts ...grpc.CallOption) (Collector_ConnectProcessorClient, error)
}

type collectorClient struct {
	cc *grpc.ClientConn
}

func NewCollectorClient(cc *grpc.ClientConn) CollectorClient {
	return &collectorClient{cc}
}

func (c *collectorClient) ConnectProvider(ctx context.Context, opts ...grpc.CallOption) (Collector_ConnectProviderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Collector_serviceDesc.Streams[0], "/tracking.Collector/ConnectProvider", opts...)
	if err != nil {
		return nil, err
	}
	x := &collectorConnectProviderClient{stream}
	return x, nil
}

type Collector_ConnectProviderClient interface {
	Send(*Update_FromProvider) error
	Recv() (*Update_ToProvider, error)
	grpc.ClientStream
}

type collectorConnectProviderClient struct {
	grpc.ClientStream
}

func (x *collectorConnectProviderClient) Send(m *Update_FromProvider) error {
	return x.ClientStream.SendMsg(m)
}

func (x *collectorConnectProviderClient) Recv() (*Update_ToProvider, error) {
	m := new(Update_ToProvider)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *collectorClient) ConnectProcessor(ctx context.Context, opts ...grpc.CallOption) (Collector_ConnectProcessorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Collector_serviceDesc.Streams[1], "/tracking.Collector/ConnectProcessor", opts...)
	if err != nil {
		return nil, err
	}
	x := &collectorConnectProcessorClient{stream}
	return x, nil
}

type Collector_ConnectProcessorClient interface {
	Send(*Update_FromProcessor) error
	Recv() (*Update_ToProcessor, error)
	grpc.ClientStream
}

type collectorConnectProcessorClient struct {
	grpc.ClientStream
}

func (x *collectorConnectProcessorClient) Send(m *Update_FromProcessor) error {
	return x.ClientStream.SendMsg(m)
}

func (x *collectorConnectProcessorClient) Recv() (*Update_ToProcessor, error) {
	m := new(Update_ToProcessor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CollectorServer is the server API for Collector service.
type CollectorServer interface {
	// ConnectProvider connects a stream of updates from a provider to a collector.
	ConnectProvider(Collector_ConnectProviderServer) error
	// ConnectProcessor connects a stream of updates from a collector to a processor.
	ConnectProcessor(Collector_ConnectProcessorServer) error
}

func RegisterCollectorServer(s *grpc.Server, srv CollectorServer) {
	s.RegisterService(&_Collector_serviceDesc, srv)
}

func _Collector_ConnectProvider_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CollectorServer).ConnectProvider(&collectorConnectProviderServer{stream})
}

type Collector_ConnectProviderServer interface {
	Send(*Update_ToProvider) error
	Recv() (*Update_FromProvider, error)
	grpc.ServerStream
}

type collectorConnectProviderServer struct {
	grpc.ServerStream
}

func (x *collectorConnectProviderServer) Send(m *Update_ToProvider) error {
	return x.ServerStream.SendMsg(m)
}

func (x *collectorConnectProviderServer) Recv() (*Update_FromProvider, error) {
	m := new(Update_FromProvider)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Collector_ConnectProcessor_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CollectorServer).ConnectProcessor(&collectorConnectProcessorServer{stream})
}

type Collector_ConnectProcessorServer interface {
	Send(*Update_ToProcessor) error
	Recv() (*Update_FromProcessor, error)
	grpc.ServerStream
}

type collectorConnectProcessorServer struct {
	grpc.ServerStream
}

func (x *collectorConnectProcessorServer) Send(m *Update_ToProcessor) error {
	return x.ServerStream.SendMsg(m)
}

func (x *collectorConnectProcessorServer) Recv() (*Update_FromProcessor, error) {
	m := new(Update_FromProcessor)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Collector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tracking.Collector",
	HandlerType: (*CollectorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectProvider",
			Handler:       _Collector_ConnectProvider_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ConnectProcessor",
			Handler:       _Collector_ConnectProcessor_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "tracking/collector.proto",
}

func init() { proto.RegisterFile("tracking/collector.proto", fileDescriptor_collector_f50ae4f6c1836da5) }

var fileDescriptor_collector_f50ae4f6c1836da5 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcb, 0xae, 0xd3, 0x30,
	0x14, 0x6c, 0xa8, 0x08, 0xf7, 0x9e, 0xf2, 0x08, 0x06, 0xa4, 0x28, 0x3c, 0x75, 0xc5, 0xa2, 0x12,
	0xe0, 0xa0, 0xb2, 0x47, 0xa2, 0x95, 0x50, 0xd9, 0x85, 0x50, 0x36, 0x6c, 0x90, 0xe3, 0x1e, 0x5a,
	0x93, 0xc4, 0x0e, 0xb6, 0x8b, 0xe0, 0x83, 0xf8, 0x03, 0xfe, 0x8e, 0x0d, 0x8a, 0x9d, 0x34, 0x15,
	0x50, 0x21, 0x58, 0xdc, 0x5d, 0x34, 0x73, 0x3c, 0x33, 0xf6, 0x99, 0x40, 0x6c, 0x35, 0xe3, 0xa5,
	0x90, 0x9b, 0x94, 0xab, 0xaa, 0x42, 0x6e, 0x95, 0xa6, 0x8d, 0x56, 0x56, 0x91, 0x93, 0x9e, 0x49,
	0x22, 0xfc, 0x62, 0x51, 0x1a, 0xa1, 0xa4, 0xf1, 0x5c, 0x72, 0xc3, 0x7c, 0x35, 0x16, 0xeb, 0xd4,
	0x58, 0x66, 0x77, 0x3d, 0x18, 0x75, 0x20, 0xe3, 0x65, 0x87, 0xdc, 0xdc, 0x8b, 0xbb, 0x0f, 0x8f,
	0x9e, 0xbd, 0x82, 0x64, 0xa1, 0xa4, 0x44, 0x6e, 0x33, 0xad, 0x38, 0x1a, 0xa3, 0x74, 0xc6, 0x34,
	0xab, 0xd1, 0xa2, 0x36, 0xe4, 0x11, 0x5c, 0x47, 0xc9, 0x8a, 0x0a, 0xdf, 0x37, 0x5a, 0x7d, 0x44,
	0x6e, 0x85, 0x92, 0x71, 0xf0, 0x20, 0x98, 0x9e, 0xe4, 0x91, 0x27, 0xb2, 0x3d, 0x7e, 0xf6, 0x63,
	0x0c, 0xe1, 0xdb, 0x66, 0xcd, 0x2c, 0x26, 0x1a, 0x2e, 0xbf, 0xd4, 0xaa, 0xce, 0xb4, 0xfa, 0x2c,
	0xd6, 0xa8, 0xc9, 0x14, 0x42, 0x9f, 0xce, 0x1d, 0x9e, 0xcc, 0xae, 0x52, 0x1f, 0x8f, 0xbe, 0x71,
	0xe8, 0x72, 0x94, 0x77, 0x3c, 0x79, 0x02, 0x17, 0x0b, 0x66, 0xf9, 0x36, 0xbe, 0xe0, 0x06, 0x6f,
	0xd1, 0x3e, 0x35, 0x5d, 0xb9, 0xd4, 0xf3, 0x96, 0x5c, 0x8e, 0x72, 0x3f, 0x35, 0x3f, 0x85, 0x4b,
	0x6b, 0xb4, 0x4c, 0x54, 0x26, 0x29, 0x00, 0x56, 0xea, 0x3f, 0x1c, 0xef, 0xc3, 0x98, 0xf1, 0xb2,
	0xf3, 0x9b, 0xf4, 0x63, 0x2f, 0x78, 0xb9, 0x1c, 0xe5, 0x2d, 0x73, 0xe8, 0xf1, 0x09, 0x26, 0xce,
	0xc3, 0x3f, 0xd4, 0xb9, 0x5c, 0xeb, 0x5b, 0x00, 0x57, 0xba, 0xb7, 0xfc, 0x67, 0xd7, 0xbf, 0x5d,
	0x8d, 0x3c, 0x87, 0xb0, 0x69, 0xb7, 0x6d, 0xe2, 0xb1, 0x9b, 0x79, 0x38, 0xe4, 0x3a, 0xde, 0x8a,
	0xd6, 0xc0, 0x9f, 0x3a, 0xc8, 0x39, 0xfb, 0x1e, 0xc0, 0xe9, 0xa2, 0x6f, 0x2d, 0x79, 0x0d, 0xd7,
	0x06, 0x01, 0xbf, 0x91, 0xbb, 0x83, 0xb6, 0x6f, 0x09, 0x3d, 0xac, 0x48, 0x72, 0xfb, 0x37, 0x7a,
	0xd8, 0xe6, 0x34, 0x78, 0x1a, 0x90, 0x15, 0x44, 0xbf, 0x66, 0x22, 0xf7, 0x8e, 0x69, 0x7a, 0x3e,
	0xb9, 0xf3, 0x67, 0x51, 0xcf, 0xb6, 0xaa, 0x73, 0xfa, 0xee, 0xf1, 0x46, 0xd8, 0xed, 0xae, 0xa0,
	0x5c, 0xd5, 0x29, 0x13, 0xba, 0x66, 0x4d, 0x2a, 0xa4, 0x45, 0xfd, 0x81, 0x71, 0x34, 0xa9, 0xd1,
	0x3c, 0xdd, 0xa8, 0xb4, 0x97, 0x29, 0x42, 0xf7, 0xdb, 0x3c, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff,
	0x37, 0x15, 0x7b, 0xd6, 0xab, 0x03, 0x00, 0x00,
}
