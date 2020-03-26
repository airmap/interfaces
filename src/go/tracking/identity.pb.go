// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracking/identity.proto

package tracking

import (
	fmt "fmt"
	_ "github.com/airmap/interfaces/src/go"
	ids "github.com/airmap/interfaces/src/go/ids"
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

// Identity models a unique tracking identity.
type Identity struct {
	// details is a discriminated union of all identity types.
	//
	// Types that are valid to be assigned to Details:
	//	*Identity_ProviderId_
	//	*Identity_TrackId_
	//	*Identity_Callsign_
	//	*Identity_Registration_
	//	*Identity_Operation_
	//	*Identity_Icao_
	//	*Identity_Manufacturer_
	//	*Identity_Network
	//	*Identity_Imei
	//	*Identity_Imsi
	Details              isIdentity_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0}
}

func (m *Identity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity.Unmarshal(m, b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return xxx_messageInfo_Identity.Size(m)
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

type isIdentity_Details interface {
	isIdentity_Details()
}

type Identity_ProviderId_ struct {
	ProviderId *Identity_ProviderId `protobuf:"bytes,1,opt,name=provider_id,json=providerId,proto3,oneof"`
}

type Identity_TrackId_ struct {
	TrackId *Identity_TrackId `protobuf:"bytes,2,opt,name=track_id,json=trackId,proto3,oneof"`
}

type Identity_Callsign_ struct {
	Callsign *Identity_Callsign `protobuf:"bytes,3,opt,name=callsign,proto3,oneof"`
}

type Identity_Registration_ struct {
	Registration *Identity_Registration `protobuf:"bytes,4,opt,name=registration,proto3,oneof"`
}

type Identity_Operation_ struct {
	Operation *Identity_Operation `protobuf:"bytes,5,opt,name=operation,proto3,oneof"`
}

type Identity_Icao_ struct {
	Icao *Identity_Icao `protobuf:"bytes,6,opt,name=icao,proto3,oneof"`
}

type Identity_Manufacturer_ struct {
	Manufacturer *Identity_Manufacturer `protobuf:"bytes,7,opt,name=manufacturer,proto3,oneof"`
}

type Identity_Network struct {
	Network *Identity_NetworkInterface `protobuf:"bytes,8,opt,name=network,proto3,oneof"`
}

type Identity_Imei struct {
	Imei *Identity_IMEI `protobuf:"bytes,9,opt,name=imei,proto3,oneof"`
}

type Identity_Imsi struct {
	Imsi *Identity_IMSI `protobuf:"bytes,10,opt,name=imsi,proto3,oneof"`
}

func (*Identity_ProviderId_) isIdentity_Details() {}

func (*Identity_TrackId_) isIdentity_Details() {}

func (*Identity_Callsign_) isIdentity_Details() {}

func (*Identity_Registration_) isIdentity_Details() {}

func (*Identity_Operation_) isIdentity_Details() {}

func (*Identity_Icao_) isIdentity_Details() {}

func (*Identity_Manufacturer_) isIdentity_Details() {}

func (*Identity_Network) isIdentity_Details() {}

func (*Identity_Imei) isIdentity_Details() {}

func (*Identity_Imsi) isIdentity_Details() {}

func (m *Identity) GetDetails() isIdentity_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Identity) GetProviderId() *Identity_ProviderId {
	if x, ok := m.GetDetails().(*Identity_ProviderId_); ok {
		return x.ProviderId
	}
	return nil
}

func (m *Identity) GetTrackId() *Identity_TrackId {
	if x, ok := m.GetDetails().(*Identity_TrackId_); ok {
		return x.TrackId
	}
	return nil
}

func (m *Identity) GetCallsign() *Identity_Callsign {
	if x, ok := m.GetDetails().(*Identity_Callsign_); ok {
		return x.Callsign
	}
	return nil
}

func (m *Identity) GetRegistration() *Identity_Registration {
	if x, ok := m.GetDetails().(*Identity_Registration_); ok {
		return x.Registration
	}
	return nil
}

func (m *Identity) GetOperation() *Identity_Operation {
	if x, ok := m.GetDetails().(*Identity_Operation_); ok {
		return x.Operation
	}
	return nil
}

func (m *Identity) GetIcao() *Identity_Icao {
	if x, ok := m.GetDetails().(*Identity_Icao_); ok {
		return x.Icao
	}
	return nil
}

func (m *Identity) GetManufacturer() *Identity_Manufacturer {
	if x, ok := m.GetDetails().(*Identity_Manufacturer_); ok {
		return x.Manufacturer
	}
	return nil
}

func (m *Identity) GetNetwork() *Identity_NetworkInterface {
	if x, ok := m.GetDetails().(*Identity_Network); ok {
		return x.Network
	}
	return nil
}

func (m *Identity) GetImei() *Identity_IMEI {
	if x, ok := m.GetDetails().(*Identity_Imei); ok {
		return x.Imei
	}
	return nil
}

func (m *Identity) GetImsi() *Identity_IMSI {
	if x, ok := m.GetDetails().(*Identity_Imsi); ok {
		return x.Imsi
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Identity) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Identity_ProviderId_)(nil),
		(*Identity_TrackId_)(nil),
		(*Identity_Callsign_)(nil),
		(*Identity_Registration_)(nil),
		(*Identity_Operation_)(nil),
		(*Identity_Icao_)(nil),
		(*Identity_Manufacturer_)(nil),
		(*Identity_Network)(nil),
		(*Identity_Imei)(nil),
		(*Identity_Imsi)(nil),
	}
}

// ProviderId models the unique identifier of a provider of tracks.
type Identity_ProviderId struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_ProviderId) Reset()         { *m = Identity_ProviderId{} }
func (m *Identity_ProviderId) String() string { return proto.CompactTextString(m) }
func (*Identity_ProviderId) ProtoMessage()    {}
func (*Identity_ProviderId) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 0}
}

func (m *Identity_ProviderId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_ProviderId.Unmarshal(m, b)
}
func (m *Identity_ProviderId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_ProviderId.Marshal(b, m, deterministic)
}
func (m *Identity_ProviderId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_ProviderId.Merge(m, src)
}
func (m *Identity_ProviderId) XXX_Size() int {
	return xxx_messageInfo_Identity_ProviderId.Size(m)
}
func (m *Identity_ProviderId) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_ProviderId.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_ProviderId proto.InternalMessageInfo

func (m *Identity_ProviderId) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

// Track models a unique identifier for a track assembled from multiple different observations.
type Identity_TrackId struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_TrackId) Reset()         { *m = Identity_TrackId{} }
func (m *Identity_TrackId) String() string { return proto.CompactTextString(m) }
func (*Identity_TrackId) ProtoMessage()    {}
func (*Identity_TrackId) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 1}
}

func (m *Identity_TrackId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_TrackId.Unmarshal(m, b)
}
func (m *Identity_TrackId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_TrackId.Marshal(b, m, deterministic)
}
func (m *Identity_TrackId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_TrackId.Merge(m, src)
}
func (m *Identity_TrackId) XXX_Size() int {
	return xxx_messageInfo_Identity_TrackId.Size(m)
}
func (m *Identity_TrackId) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_TrackId.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_TrackId proto.InternalMessageInfo

func (m *Identity_TrackId) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

// Callsign models the callsign assigned to a vehicle.
type Identity_Callsign struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_Callsign) Reset()         { *m = Identity_Callsign{} }
func (m *Identity_Callsign) String() string { return proto.CompactTextString(m) }
func (*Identity_Callsign) ProtoMessage()    {}
func (*Identity_Callsign) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 2}
}

func (m *Identity_Callsign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_Callsign.Unmarshal(m, b)
}
func (m *Identity_Callsign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_Callsign.Marshal(b, m, deterministic)
}
func (m *Identity_Callsign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_Callsign.Merge(m, src)
}
func (m *Identity_Callsign) XXX_Size() int {
	return xxx_messageInfo_Identity_Callsign.Size(m)
}
func (m *Identity_Callsign) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_Callsign.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_Callsign proto.InternalMessageInfo

func (m *Identity_Callsign) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

// Registration models the registration of an aircraft.
type Identity_Registration struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_Registration) Reset()         { *m = Identity_Registration{} }
func (m *Identity_Registration) String() string { return proto.CompactTextString(m) }
func (*Identity_Registration) ProtoMessage()    {}
func (*Identity_Registration) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 3}
}

func (m *Identity_Registration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_Registration.Unmarshal(m, b)
}
func (m *Identity_Registration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_Registration.Marshal(b, m, deterministic)
}
func (m *Identity_Registration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_Registration.Merge(m, src)
}
func (m *Identity_Registration) XXX_Size() int {
	return xxx_messageInfo_Identity_Registration.Size(m)
}
func (m *Identity_Registration) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_Registration.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_Registration proto.InternalMessageInfo

func (m *Identity_Registration) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

// Operation models an operation.
type Identity_Operation struct {
	OperationId          *ids.Operation `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	ServiceProviderId    *ids.USS       `protobuf:"bytes,2,opt,name=service_provider_id,json=serviceProviderId,proto3" json:"service_provider_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Identity_Operation) Reset()         { *m = Identity_Operation{} }
func (m *Identity_Operation) String() string { return proto.CompactTextString(m) }
func (*Identity_Operation) ProtoMessage()    {}
func (*Identity_Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 4}
}

func (m *Identity_Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_Operation.Unmarshal(m, b)
}
func (m *Identity_Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_Operation.Marshal(b, m, deterministic)
}
func (m *Identity_Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_Operation.Merge(m, src)
}
func (m *Identity_Operation) XXX_Size() int {
	return xxx_messageInfo_Identity_Operation.Size(m)
}
func (m *Identity_Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_Operation proto.InternalMessageInfo

func (m *Identity_Operation) GetOperationId() *ids.Operation {
	if m != nil {
		return m.OperationId
	}
	return nil
}

func (m *Identity_Operation) GetServiceProviderId() *ids.USS {
	if m != nil {
		return m.ServiceProviderId
	}
	return nil
}

// Icao bundles up information that describes an aircraft in the ICAO context.
type Identity_Icao struct {
	Address              *Identity_Icao_Address24    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AircraftType         *Identity_Icao_AircraftType `protobuf:"bytes,2,opt,name=aircraft_type,json=aircraftType,proto3" json:"aircraft_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Identity_Icao) Reset()         { *m = Identity_Icao{} }
func (m *Identity_Icao) String() string { return proto.CompactTextString(m) }
func (*Identity_Icao) ProtoMessage()    {}
func (*Identity_Icao) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 5}
}

func (m *Identity_Icao) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_Icao.Unmarshal(m, b)
}
func (m *Identity_Icao) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_Icao.Marshal(b, m, deterministic)
}
func (m *Identity_Icao) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_Icao.Merge(m, src)
}
func (m *Identity_Icao) XXX_Size() int {
	return xxx_messageInfo_Identity_Icao.Size(m)
}
func (m *Identity_Icao) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_Icao.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_Icao proto.InternalMessageInfo

func (m *Identity_Icao) GetAddress() *Identity_Icao_Address24 {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Identity_Icao) GetAircraftType() *Identity_Icao_AircraftType {
	if m != nil {
		return m.AircraftType
	}
	return nil
}

// Address24 models the unique address of an aircraft assigned by ICAO.
type Identity_Icao_Address24 struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_Icao_Address24) Reset()         { *m = Identity_Icao_Address24{} }
func (m *Identity_Icao_Address24) String() string { return proto.CompactTextString(m) }
func (*Identity_Icao_Address24) ProtoMessage()    {}
func (*Identity_Icao_Address24) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 5, 0}
}

func (m *Identity_Icao_Address24) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_Icao_Address24.Unmarshal(m, b)
}
func (m *Identity_Icao_Address24) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_Icao_Address24.Marshal(b, m, deterministic)
}
func (m *Identity_Icao_Address24) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_Icao_Address24.Merge(m, src)
}
func (m *Identity_Icao_Address24) XXX_Size() int {
	return xxx_messageInfo_Identity_Icao_Address24.Size(m)
}
func (m *Identity_Icao_Address24) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_Icao_Address24.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_Icao_Address24 proto.InternalMessageInfo

func (m *Identity_Icao_Address24) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

// AircraftType models an ICAO-registered aircraft type.
type Identity_Icao_AircraftType struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_Icao_AircraftType) Reset()         { *m = Identity_Icao_AircraftType{} }
func (m *Identity_Icao_AircraftType) String() string { return proto.CompactTextString(m) }
func (*Identity_Icao_AircraftType) ProtoMessage()    {}
func (*Identity_Icao_AircraftType) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 5, 1}
}

func (m *Identity_Icao_AircraftType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_Icao_AircraftType.Unmarshal(m, b)
}
func (m *Identity_Icao_AircraftType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_Icao_AircraftType.Marshal(b, m, deterministic)
}
func (m *Identity_Icao_AircraftType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_Icao_AircraftType.Merge(m, src)
}
func (m *Identity_Icao_AircraftType) XXX_Size() int {
	return xxx_messageInfo_Identity_Icao_AircraftType.Size(m)
}
func (m *Identity_Icao_AircraftType) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_Icao_AircraftType.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_Icao_AircraftType proto.InternalMessageInfo

func (m *Identity_Icao_AircraftType) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

// Manufacturer identifies a make, model, and serial number of an aircraft or sensor.
type Identity_Manufacturer struct {
	Make                 string   `protobuf:"bytes,1,opt,name=make,proto3" json:"make,omitempty"`
	Model                string   `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	SerialNumber         string   `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_Manufacturer) Reset()         { *m = Identity_Manufacturer{} }
func (m *Identity_Manufacturer) String() string { return proto.CompactTextString(m) }
func (*Identity_Manufacturer) ProtoMessage()    {}
func (*Identity_Manufacturer) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 6}
}

func (m *Identity_Manufacturer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_Manufacturer.Unmarshal(m, b)
}
func (m *Identity_Manufacturer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_Manufacturer.Marshal(b, m, deterministic)
}
func (m *Identity_Manufacturer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_Manufacturer.Merge(m, src)
}
func (m *Identity_Manufacturer) XXX_Size() int {
	return xxx_messageInfo_Identity_Manufacturer.Size(m)
}
func (m *Identity_Manufacturer) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_Manufacturer.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_Manufacturer proto.InternalMessageInfo

func (m *Identity_Manufacturer) GetMake() string {
	if m != nil {
		return m.Make
	}
	return ""
}

func (m *Identity_Manufacturer) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Identity_Manufacturer) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

// NetworkInterface identifies a vehicle by its network interfaces.
type Identity_NetworkInterface struct {
	MacAddress           *Identity_NetworkInterface_MACAddress `protobuf:"bytes,1,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *Identity_NetworkInterface) Reset()         { *m = Identity_NetworkInterface{} }
func (m *Identity_NetworkInterface) String() string { return proto.CompactTextString(m) }
func (*Identity_NetworkInterface) ProtoMessage()    {}
func (*Identity_NetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 7}
}

func (m *Identity_NetworkInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_NetworkInterface.Unmarshal(m, b)
}
func (m *Identity_NetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_NetworkInterface.Marshal(b, m, deterministic)
}
func (m *Identity_NetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_NetworkInterface.Merge(m, src)
}
func (m *Identity_NetworkInterface) XXX_Size() int {
	return xxx_messageInfo_Identity_NetworkInterface.Size(m)
}
func (m *Identity_NetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_NetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_NetworkInterface proto.InternalMessageInfo

func (m *Identity_NetworkInterface) GetMacAddress() *Identity_NetworkInterface_MACAddress {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

// MACAddress models the hardware address of the network interface.
type Identity_NetworkInterface_MACAddress struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_NetworkInterface_MACAddress) Reset()         { *m = Identity_NetworkInterface_MACAddress{} }
func (m *Identity_NetworkInterface_MACAddress) String() string { return proto.CompactTextString(m) }
func (*Identity_NetworkInterface_MACAddress) ProtoMessage()    {}
func (*Identity_NetworkInterface_MACAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 7, 0}
}

func (m *Identity_NetworkInterface_MACAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_NetworkInterface_MACAddress.Unmarshal(m, b)
}
func (m *Identity_NetworkInterface_MACAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_NetworkInterface_MACAddress.Marshal(b, m, deterministic)
}
func (m *Identity_NetworkInterface_MACAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_NetworkInterface_MACAddress.Merge(m, src)
}
func (m *Identity_NetworkInterface_MACAddress) XXX_Size() int {
	return xxx_messageInfo_Identity_NetworkInterface_MACAddress.Size(m)
}
func (m *Identity_NetworkInterface_MACAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_NetworkInterface_MACAddress.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_NetworkInterface_MACAddress proto.InternalMessageInfo

func (m *Identity_NetworkInterface_MACAddress) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

// IMEI models the International Mobile Equipment Identity.
type Identity_IMEI struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_IMEI) Reset()         { *m = Identity_IMEI{} }
func (m *Identity_IMEI) String() string { return proto.CompactTextString(m) }
func (*Identity_IMEI) ProtoMessage()    {}
func (*Identity_IMEI) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 8}
}

func (m *Identity_IMEI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_IMEI.Unmarshal(m, b)
}
func (m *Identity_IMEI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_IMEI.Marshal(b, m, deterministic)
}
func (m *Identity_IMEI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_IMEI.Merge(m, src)
}
func (m *Identity_IMEI) XXX_Size() int {
	return xxx_messageInfo_Identity_IMEI.Size(m)
}
func (m *Identity_IMEI) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_IMEI.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_IMEI proto.InternalMessageInfo

func (m *Identity_IMEI) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

// IMSI models the International Mobile Subscriber Identity.
type Identity_IMSI struct {
	AsString             string   `protobuf:"bytes,1,opt,name=as_string,json=asString,proto3" json:"as_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identity_IMSI) Reset()         { *m = Identity_IMSI{} }
func (m *Identity_IMSI) String() string { return proto.CompactTextString(m) }
func (*Identity_IMSI) ProtoMessage()    {}
func (*Identity_IMSI) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f0748654e63f4b5, []int{0, 9}
}

func (m *Identity_IMSI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identity_IMSI.Unmarshal(m, b)
}
func (m *Identity_IMSI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identity_IMSI.Marshal(b, m, deterministic)
}
func (m *Identity_IMSI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity_IMSI.Merge(m, src)
}
func (m *Identity_IMSI) XXX_Size() int {
	return xxx_messageInfo_Identity_IMSI.Size(m)
}
func (m *Identity_IMSI) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity_IMSI.DiscardUnknown(m)
}

var xxx_messageInfo_Identity_IMSI proto.InternalMessageInfo

func (m *Identity_IMSI) GetAsString() string {
	if m != nil {
		return m.AsString
	}
	return ""
}

func init() {
	proto.RegisterType((*Identity)(nil), "tracking.Identity")
	proto.RegisterType((*Identity_ProviderId)(nil), "tracking.Identity.ProviderId")
	proto.RegisterType((*Identity_TrackId)(nil), "tracking.Identity.TrackId")
	proto.RegisterType((*Identity_Callsign)(nil), "tracking.Identity.Callsign")
	proto.RegisterType((*Identity_Registration)(nil), "tracking.Identity.Registration")
	proto.RegisterType((*Identity_Operation)(nil), "tracking.Identity.Operation")
	proto.RegisterType((*Identity_Icao)(nil), "tracking.Identity.Icao")
	proto.RegisterType((*Identity_Icao_Address24)(nil), "tracking.Identity.Icao.Address24")
	proto.RegisterType((*Identity_Icao_AircraftType)(nil), "tracking.Identity.Icao.AircraftType")
	proto.RegisterType((*Identity_Manufacturer)(nil), "tracking.Identity.Manufacturer")
	proto.RegisterType((*Identity_NetworkInterface)(nil), "tracking.Identity.NetworkInterface")
	proto.RegisterType((*Identity_NetworkInterface_MACAddress)(nil), "tracking.Identity.NetworkInterface.MACAddress")
	proto.RegisterType((*Identity_IMEI)(nil), "tracking.Identity.IMEI")
	proto.RegisterType((*Identity_IMSI)(nil), "tracking.Identity.IMSI")
}

func init() { proto.RegisterFile("tracking/identity.proto", fileDescriptor_0f0748654e63f4b5) }

var fileDescriptor_0f0748654e63f4b5 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x5d, 0x4f, 0x13, 0x4d,
	0x14, 0xc7, 0x5b, 0x9e, 0x42, 0xdb, 0xd3, 0xe5, 0x49, 0x9f, 0x79, 0x4c, 0xd8, 0x2c, 0xae, 0x0a,
	0x8a, 0x91, 0x28, 0xbb, 0x82, 0x26, 0x2a, 0x51, 0x91, 0x12, 0x62, 0x9b, 0x08, 0x98, 0x2d, 0x26,
	0xc6, 0x9b, 0x66, 0xd8, 0x1d, 0xea, 0x84, 0xee, 0x4b, 0x66, 0xa6, 0x08, 0x77, 0x26, 0xde, 0xf5,
	0x83, 0xf4, 0xde, 0xcf, 0xe3, 0x97, 0x31, 0x3b, 0x3b, 0xbb, 0xdd, 0xd2, 0x97, 0x98, 0x98, 0xa6,
	0xcd, 0xce, 0xcc, 0xef, 0xff, 0x3f, 0xa7, 0x67, 0xce, 0x69, 0x61, 0x45, 0x30, 0xec, 0x5e, 0xd0,
	0xa0, 0x6b, 0x53, 0x8f, 0x04, 0x82, 0x8a, 0x6b, 0x2b, 0x62, 0xa1, 0x08, 0x51, 0x25, 0x3d, 0x30,
	0xea, 0xe4, 0x4a, 0x90, 0x80, 0xd3, 0x30, 0xe0, 0xc9, 0x99, 0xb1, 0x4c, 0x3d, 0x6e, 0x53, 0x4f,
	0x2d, 0xd7, 0x7f, 0x69, 0x50, 0x69, 0x29, 0x35, 0x7a, 0x07, 0xb5, 0x88, 0x85, 0x97, 0xd4, 0x23,
	0xac, 0x43, 0x3d, 0xbd, 0x78, 0xaf, 0xf8, 0xa8, 0xb6, 0x63, 0x5a, 0xa9, 0x9b, 0x95, 0x82, 0xd6,
	0x47, 0x45, 0xb5, 0xbc, 0x66, 0xc1, 0x81, 0x28, 0x5b, 0xa1, 0x17, 0x90, 0xc4, 0x8e, 0xe5, 0x0b,
	0x52, 0x6e, 0x4c, 0x91, 0x9f, 0xc6, 0x3b, 0x52, 0x5b, 0x16, 0xc9, 0x23, 0x7a, 0x05, 0x15, 0x17,
	0xf7, 0x7a, 0x9c, 0x76, 0x03, 0xfd, 0x1f, 0x29, 0x5c, 0x9d, 0x22, 0x3c, 0x50, 0x48, 0xb3, 0xe0,
	0x64, 0x38, 0x3a, 0x04, 0x8d, 0x91, 0x2e, 0xe5, 0x82, 0x61, 0x41, 0xc3, 0x40, 0x2f, 0x49, 0xf9,
	0xdd, 0x29, 0x72, 0x27, 0x87, 0x35, 0x0b, 0xce, 0x98, 0x0c, 0xbd, 0x86, 0x6a, 0x18, 0x11, 0xe5,
	0xb1, 0x28, 0x3d, 0x6e, 0x4f, 0xf1, 0x38, 0x49, 0x99, 0x66, 0xc1, 0x19, 0x09, 0xd0, 0x16, 0x94,
	0xa8, 0x8b, 0x43, 0x7d, 0x49, 0x0a, 0x57, 0xa6, 0x08, 0x5b, 0x2e, 0x0e, 0x9b, 0x05, 0x47, 0x62,
	0x71, 0xce, 0x3e, 0x0e, 0xfa, 0xe7, 0xd8, 0x15, 0x7d, 0x46, 0x98, 0x5e, 0x9e, 0x99, 0xf3, 0x51,
	0x0e, 0x8b, 0x73, 0xce, 0xcb, 0xd0, 0x1e, 0x94, 0x03, 0x22, 0xbe, 0x85, 0xec, 0x42, 0xaf, 0x48,
	0x87, 0xfb, 0x53, 0x1c, 0x8e, 0x13, 0xa2, 0x15, 0x08, 0xc2, 0xce, 0xb1, 0x4b, 0xe2, 0xb2, 0x2b,
	0x95, 0x4c, 0xdb, 0x27, 0x54, 0xaf, 0xce, 0x4e, 0xfb, 0xe8, 0xb0, 0x25, 0xd3, 0xf6, 0x09, 0x4d,
	0x70, 0x4e, 0x75, 0x98, 0x83, 0xb7, 0x15, 0xce, 0xa9, 0xb1, 0x09, 0x30, 0xea, 0x14, 0xb4, 0x0a,
	0x55, 0xcc, 0x3b, 0x5c, 0x30, 0x1a, 0x74, 0x65, 0x6f, 0x55, 0x9d, 0x0a, 0xe6, 0x6d, 0xb9, 0x36,
	0x1e, 0x42, 0x59, 0x75, 0xc5, 0x7c, 0xee, 0x0d, 0x54, 0xd2, 0x26, 0x40, 0xdb, 0x13, 0x60, 0xe3,
	0xd6, 0x60, 0x68, 0xd6, 0x37, 0xdc, 0xf1, 0x57, 0x4e, 0xfe, 0x18, 0xb4, 0x7c, 0x13, 0xcc, 0x8f,
	0x75, 0x05, 0xd5, 0xec, 0xb6, 0xd1, 0x36, 0x68, 0xd9, 0x6d, 0x8f, 0x86, 0xe3, 0x5f, 0x2b, 0x1e,
	0xa5, 0x8c, 0x72, 0x6a, 0x19, 0xd3, 0xf2, 0xd0, 0x4b, 0xf8, 0x9f, 0x13, 0x76, 0x49, 0x5d, 0xd2,
	0xc9, 0x8f, 0x55, 0x32, 0x17, 0x15, 0xa9, 0xfc, 0xd4, 0x6e, 0x3b, 0xff, 0x29, 0x68, 0x54, 0x2a,
	0xe3, 0xc7, 0x02, 0x94, 0xe2, 0x7e, 0x41, 0xfb, 0x50, 0xc6, 0x9e, 0xc7, 0x08, 0xe7, 0x2a, 0xe0,
	0xda, 0x8c, 0xce, 0xb2, 0xf6, 0x13, 0x6c, 0xe7, 0x79, 0xa3, 0xf4, 0x7d, 0x68, 0x16, 0x9d, 0x54,
	0x87, 0x4e, 0x60, 0x19, 0x53, 0xe6, 0x32, 0x7c, 0x2e, 0x3a, 0xe2, 0x3a, 0x22, 0x2a, 0xfe, 0x83,
	0x99, 0x46, 0x0a, 0x3e, 0xbd, 0x8e, 0x88, 0xf2, 0xd2, 0x70, 0x6e, 0xcf, 0xd8, 0x85, 0x6a, 0x16,
	0x0c, 0x6d, 0x4d, 0xde, 0x41, 0x7d, 0x30, 0x34, 0xb5, 0x8d, 0xa7, 0x3b, 0x9f, 0xd3, 0xf7, 0x78,
	0xfd, 0xf3, 0xfe, 0xf3, 0xeb, 0x1f, 0x82, 0x96, 0xef, 0x7e, 0xa4, 0x43, 0xc9, 0xc7, 0x17, 0x44,
	0x85, 0x49, 0x52, 0x93, 0x3b, 0xc8, 0x80, 0x45, 0x3f, 0xf4, 0x48, 0x4f, 0x7e, 0xb7, 0xf4, 0x28,
	0xd9, 0x42, 0x9b, 0xb0, 0xcc, 0x09, 0xa3, 0xb8, 0xd7, 0x09, 0xfa, 0xfe, 0x19, 0x61, 0xf2, 0xe7,
	0x25, 0x65, 0xb4, 0xe4, 0xe8, 0x58, 0x9e, 0x18, 0x3f, 0x8b, 0x50, 0xbf, 0x39, 0x2d, 0xe8, 0x04,
	0x6a, 0x3e, 0x76, 0x3b, 0xe3, 0xd7, 0x60, 0xfd, 0xc1, 0x9c, 0x59, 0x47, 0xfb, 0x07, 0xaa, 0x50,
	0x0e, 0xf8, 0xd8, 0x55, 0xcf, 0xc6, 0x07, 0x80, 0xd1, 0x09, 0x7a, 0x3b, 0x59, 0xc0, 0xb5, 0xc1,
	0xd0, 0x34, 0xe3, 0xc2, 0xed, 0xce, 0xfa, 0xc8, 0x15, 0xe9, 0x3d, 0x94, 0xe2, 0x11, 0x45, 0x7b,
	0x93, 0x3e, 0xeb, 0x83, 0xa1, 0x79, 0xe7, 0xe6, 0x30, 0xcc, 0x1c, 0x0d, 0x69, 0xd4, 0xfe, 0x7b,
	0xa3, 0x46, 0x15, 0xca, 0x1e, 0x11, 0x98, 0xf6, 0x78, 0xc3, 0xfa, 0xf2, 0xa4, 0x4b, 0xc5, 0xd7,
	0xfe, 0x99, 0xe5, 0x86, 0xbe, 0x8d, 0x29, 0xf3, 0x71, 0x64, 0xd3, 0xb4, 0x40, 0xdc, 0xe6, 0xcc,
	0xb5, 0xbb, 0xa1, 0x9d, 0xd6, 0xf2, 0x6c, 0x49, 0xfe, 0x29, 0x3d, 0xfb, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0xb4, 0xa6, 0x6c, 0xda, 0x06, 0x00, 0x00,
}
