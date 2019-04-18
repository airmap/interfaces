// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracking/collector.proto

package tracking

import (
	context "context"
	fmt "fmt"
	_ "github.com/airmap/interfaces/src/go"
	system "github.com/airmap/interfaces/src/go/system"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
	return fileDescriptor_ca152adb0a29ed43, []int{0}
}

func (m *Update) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update.Unmarshal(m, b)
}
func (m *Update) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update.Marshal(b, m, deterministic)
}
func (m *Update) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update.Merge(m, src)
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
	return fileDescriptor_ca152adb0a29ed43, []int{0, 0}
}

func (m *Update_FromProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_FromProvider.Unmarshal(m, b)
}
func (m *Update_FromProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_FromProvider.Marshal(b, m, deterministic)
}
func (m *Update_FromProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_FromProvider.Merge(m, src)
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

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Update_FromProvider) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Update_FromProvider_Status)(nil),
		(*Update_FromProvider_Batch)(nil),
	}
}

// ToProvider wraps messages being sent from a collector back to a provider.
type Update_ToProvider struct {
	// Types that are valid to be assigned to Details:
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
	return fileDescriptor_ca152adb0a29ed43, []int{0, 1}
}

func (m *Update_ToProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_ToProvider.Unmarshal(m, b)
}
func (m *Update_ToProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_ToProvider.Marshal(b, m, deterministic)
}
func (m *Update_ToProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_ToProvider.Merge(m, src)
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

type Update_ToProvider_Ack struct {
	Ack *system.Ack `protobuf:"bytes,1,opt,name=ack,proto3,oneof"`
}

func (*Update_ToProvider_Ack) isUpdate_ToProvider_Details() {}

func (m *Update_ToProvider) GetDetails() isUpdate_ToProvider_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Update_ToProvider) GetAck() *system.Ack {
	if x, ok := m.GetDetails().(*Update_ToProvider_Ack); ok {
		return x.Ack
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Update_ToProvider) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Update_ToProvider_Ack)(nil),
	}
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
	return fileDescriptor_ca152adb0a29ed43, []int{0, 2}
}

func (m *Update_ToProcessor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_ToProcessor.Unmarshal(m, b)
}
func (m *Update_ToProcessor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_ToProcessor.Marshal(b, m, deterministic)
}
func (m *Update_ToProcessor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_ToProcessor.Merge(m, src)
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

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Update_ToProcessor) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Update_ToProcessor_Status)(nil),
		(*Update_ToProcessor_Batch)(nil),
	}
}

type Update_FromProcessor struct {
	// Types that are valid to be assigned to Details:
	//	*Update_FromProcessor_Ack
	Details              isUpdate_FromProcessor_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Update_FromProcessor) Reset()         { *m = Update_FromProcessor{} }
func (m *Update_FromProcessor) String() string { return proto.CompactTextString(m) }
func (*Update_FromProcessor) ProtoMessage()    {}
func (*Update_FromProcessor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca152adb0a29ed43, []int{0, 3}
}

func (m *Update_FromProcessor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update_FromProcessor.Unmarshal(m, b)
}
func (m *Update_FromProcessor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update_FromProcessor.Marshal(b, m, deterministic)
}
func (m *Update_FromProcessor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update_FromProcessor.Merge(m, src)
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

type Update_FromProcessor_Ack struct {
	Ack *system.Ack `protobuf:"bytes,1,opt,name=ack,proto3,oneof"`
}

func (*Update_FromProcessor_Ack) isUpdate_FromProcessor_Details() {}

func (m *Update_FromProcessor) GetDetails() isUpdate_FromProcessor_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Update_FromProcessor) GetAck() *system.Ack {
	if x, ok := m.GetDetails().(*Update_FromProcessor_Ack); ok {
		return x.Ack
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Update_FromProcessor) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Update_FromProcessor_Ack)(nil),
	}
}

func init() {
	proto.RegisterType((*Update)(nil), "tracking.Update")
	proto.RegisterType((*Update_FromProvider)(nil), "tracking.Update.FromProvider")
	proto.RegisterType((*Update_ToProvider)(nil), "tracking.Update.ToProvider")
	proto.RegisterType((*Update_ToProcessor)(nil), "tracking.Update.ToProcessor")
	proto.RegisterType((*Update_FromProcessor)(nil), "tracking.Update.FromProcessor")
}

func init() { proto.RegisterFile("tracking/collector.proto", fileDescriptor_ca152adb0a29ed43) }

var fileDescriptor_ca152adb0a29ed43 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcd, 0x4e, 0x02, 0x31,
	0x14, 0x85, 0x19, 0x8c, 0x28, 0x17, 0x7f, 0xc8, 0xa8, 0x09, 0x19, 0x7f, 0xe3, 0x8a, 0x85, 0xb6,
	0x06, 0x37, 0x26, 0xae, 0x84, 0xc4, 0xb0, 0x54, 0xc4, 0x8d, 0xbb, 0x52, 0x2a, 0x34, 0xc3, 0xcc,
	0xc5, 0xf6, 0x62, 0xf4, 0xbd, 0x7c, 0x26, 0x9f, 0xc3, 0x30, 0xed, 0x30, 0x24, 0xea, 0xc2, 0x8d,
	0xbb, 0xc9, 0x39, 0xed, 0x77, 0xce, 0xdc, 0x5b, 0x68, 0x90, 0x11, 0x32, 0xd6, 0xe9, 0x88, 0x4b,
	0x9c, 0x4c, 0x94, 0x24, 0x34, 0x6c, 0x6a, 0x90, 0x30, 0x5c, 0xcf, 0x9d, 0xa8, 0xae, 0xde, 0x48,
	0xa5, 0x56, 0x63, 0x6a, 0x9d, 0x17, 0xed, 0xd8, 0x77, 0x4b, 0x2a, 0xe1, 0x96, 0x04, 0xcd, 0x72,
	0xb1, 0xee, 0x45, 0x21, 0x63, 0xaf, 0xec, 0x2e, 0xe0, 0xd9, 0x87, 0x53, 0x4f, 0x3f, 0xcb, 0x50,
	0x79, 0x9c, 0x0e, 0x05, 0xa9, 0xc8, 0xc0, 0xc6, 0xad, 0xc1, 0xe4, 0xce, 0xe0, 0xab, 0x1e, 0x2a,
	0x13, 0x36, 0xa1, 0xe2, 0x90, 0x8d, 0xe0, 0x24, 0x68, 0xd6, 0x5a, 0x5b, 0xcc, 0x31, 0xd9, 0x43,
	0xa6, 0x76, 0x4b, 0x3d, 0xef, 0x87, 0xe7, 0xb0, 0x3a, 0x10, 0x24, 0xc7, 0x8d, 0x72, 0x76, 0x70,
	0x8f, 0xe5, 0x51, 0xac, 0x9f, 0x45, 0xb5, 0xe7, 0x66, 0xb7, 0xd4, 0x73, 0xa7, 0xda, 0x55, 0x58,
	0x1b, 0x2a, 0x12, 0x7a, 0x62, 0xa3, 0x2b, 0x80, 0x3e, 0x2e, 0x12, 0x8f, 0x61, 0x45, 0xc8, 0xd8,
	0xc7, 0xd5, 0xf2, 0xb8, 0x1b, 0x19, 0x77, 0x4b, 0xbd, 0xb9, 0xb3, 0x7c, 0xf3, 0x05, 0x6a, 0xd9,
	0x4d, 0xa9, 0xac, 0xc5, 0xff, 0x29, 0x7b, 0x0d, 0x9b, 0x7e, 0x40, 0x3e, 0xf4, 0x0f, 0x7d, 0x5b,
	0x1f, 0x01, 0x54, 0x3b, 0xf9, 0x56, 0xc3, 0x7b, 0xd8, 0xee, 0x60, 0x9a, 0x2a, 0x49, 0x8b, 0x9f,
	0x3f, 0x2c, 0x8a, 0xb8, 0x85, 0xb0, 0xe5, 0x6d, 0x44, 0xfb, 0xdf, 0xec, 0x62, 0x70, 0xcd, 0xe0,
	0x22, 0x08, 0xfb, 0x50, 0x2f, 0x90, 0xbe, 0xe0, 0xd1, 0x6f, 0x4c, 0xe7, 0x47, 0x07, 0x3f, 0x43,
	0x9d, 0x3b, 0xa7, 0xb6, 0xd9, 0xd3, 0xd9, 0x48, 0xd3, 0x78, 0x36, 0x60, 0x12, 0x13, 0x2e, 0xb4,
	0x49, 0xc4, 0x94, 0xeb, 0x94, 0x94, 0x79, 0x16, 0x52, 0x59, 0x6e, 0x8d, 0xe4, 0x23, 0xe4, 0x39,
	0x66, 0x50, 0xc9, 0x9e, 0xd5, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x54, 0x0e, 0x24,
	0xcb, 0x02, 0x00, 0x00,
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
