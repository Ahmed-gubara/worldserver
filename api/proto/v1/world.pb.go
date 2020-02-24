// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/v1/world.proto

package v1

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

type Aircraft struct {
	CallSign             string   `protobuf:"bytes,1,opt,name=CallSign,proto3" json:"CallSign,omitempty"`
	Heading              float32  `protobuf:"fixed32,2,opt,name=Heading,proto3" json:"Heading,omitempty"`
	Latitude             float32  `protobuf:"fixed32,3,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,4,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Speed                float32  `protobuf:"fixed32,5,opt,name=Speed,proto3" json:"Speed,omitempty"`
	Frequency            float32  `protobuf:"fixed32,6,opt,name=Frequency,proto3" json:"Frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Aircraft) Reset()         { *m = Aircraft{} }
func (m *Aircraft) String() string { return proto.CompactTextString(m) }
func (*Aircraft) ProtoMessage()    {}
func (*Aircraft) Descriptor() ([]byte, []int) {
	return fileDescriptor_99cdefe973785457, []int{0}
}

func (m *Aircraft) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Aircraft.Unmarshal(m, b)
}
func (m *Aircraft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Aircraft.Marshal(b, m, deterministic)
}
func (m *Aircraft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Aircraft.Merge(m, src)
}
func (m *Aircraft) XXX_Size() int {
	return xxx_messageInfo_Aircraft.Size(m)
}
func (m *Aircraft) XXX_DiscardUnknown() {
	xxx_messageInfo_Aircraft.DiscardUnknown(m)
}

var xxx_messageInfo_Aircraft proto.InternalMessageInfo

func (m *Aircraft) GetCallSign() string {
	if m != nil {
		return m.CallSign
	}
	return ""
}

func (m *Aircraft) GetHeading() float32 {
	if m != nil {
		return m.Heading
	}
	return 0
}

func (m *Aircraft) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Aircraft) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Aircraft) GetSpeed() float32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *Aircraft) GetFrequency() float32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

type UpdateMessage struct {
	CallSign             string    `protobuf:"bytes,1,opt,name=CallSign,proto3" json:"CallSign,omitempty"`
	AircraftInfo         *Aircraft `protobuf:"bytes,2,opt,name=AircraftInfo,proto3" json:"AircraftInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateMessage) Reset()         { *m = UpdateMessage{} }
func (m *UpdateMessage) String() string { return proto.CompactTextString(m) }
func (*UpdateMessage) ProtoMessage()    {}
func (*UpdateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_99cdefe973785457, []int{1}
}

func (m *UpdateMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMessage.Unmarshal(m, b)
}
func (m *UpdateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMessage.Marshal(b, m, deterministic)
}
func (m *UpdateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMessage.Merge(m, src)
}
func (m *UpdateMessage) XXX_Size() int {
	return xxx_messageInfo_UpdateMessage.Size(m)
}
func (m *UpdateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMessage proto.InternalMessageInfo

func (m *UpdateMessage) GetCallSign() string {
	if m != nil {
		return m.CallSign
	}
	return ""
}

func (m *UpdateMessage) GetAircraftInfo() *Aircraft {
	if m != nil {
		return m.AircraftInfo
	}
	return nil
}

type RemoveMessage struct {
	CallSign             string   `protobuf:"bytes,1,opt,name=CallSign,proto3" json:"CallSign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveMessage) Reset()         { *m = RemoveMessage{} }
func (m *RemoveMessage) String() string { return proto.CompactTextString(m) }
func (*RemoveMessage) ProtoMessage()    {}
func (*RemoveMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_99cdefe973785457, []int{2}
}

func (m *RemoveMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveMessage.Unmarshal(m, b)
}
func (m *RemoveMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveMessage.Marshal(b, m, deterministic)
}
func (m *RemoveMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveMessage.Merge(m, src)
}
func (m *RemoveMessage) XXX_Size() int {
	return xxx_messageInfo_RemoveMessage.Size(m)
}
func (m *RemoveMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveMessage proto.InternalMessageInfo

func (m *RemoveMessage) GetCallSign() string {
	if m != nil {
		return m.CallSign
	}
	return ""
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_99cdefe973785457, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type MonitorMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MonitorMessage) Reset()         { *m = MonitorMessage{} }
func (m *MonitorMessage) String() string { return proto.CompactTextString(m) }
func (*MonitorMessage) ProtoMessage()    {}
func (*MonitorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_99cdefe973785457, []int{4}
}

func (m *MonitorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorMessage.Unmarshal(m, b)
}
func (m *MonitorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorMessage.Marshal(b, m, deterministic)
}
func (m *MonitorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorMessage.Merge(m, src)
}
func (m *MonitorMessage) XXX_Size() int {
	return xxx_messageInfo_MonitorMessage.Size(m)
}
func (m *MonitorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorMessage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Aircraft)(nil), "v1.Aircraft")
	proto.RegisterType((*UpdateMessage)(nil), "v1.UpdateMessage")
	proto.RegisterType((*RemoveMessage)(nil), "v1.RemoveMessage")
	proto.RegisterType((*Response)(nil), "v1.Response")
	proto.RegisterType((*MonitorMessage)(nil), "v1.MonitorMessage")
}

func init() { proto.RegisterFile("api/proto/v1/world.proto", fileDescriptor_99cdefe973785457) }

var fileDescriptor_99cdefe973785457 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x49, 0xb5, 0xdb, 0x76, 0x6c, 0x45, 0x83, 0x48, 0x28, 0x1e, 0x4a, 0x4e, 0x2d, 0x62,
	0xff, 0xf9, 0x04, 0x22, 0x88, 0x42, 0x7b, 0xd9, 0x22, 0x9e, 0x3c, 0xc4, 0xee, 0x74, 0x09, 0xd4,
	0x64, 0x4d, 0xd2, 0x15, 0x9f, 0xc3, 0xf7, 0xf0, 0x19, 0x65, 0x93, 0xed, 0xca, 0x7a, 0x10, 0x8f,
	0xdf, 0x7c, 0xdf, 0xcc, 0xfc, 0x66, 0x37, 0xc0, 0x44, 0x26, 0x27, 0x99, 0xd1, 0x4e, 0x4f, 0xf2,
	0xd9, 0xe4, 0x5d, 0x9b, 0x6d, 0x32, 0xf6, 0x92, 0x36, 0xf2, 0x19, 0xff, 0x22, 0xd0, 0xbe, 0x91,
	0x66, 0x6d, 0xc4, 0xc6, 0xd1, 0x3e, 0xb4, 0x6f, 0xc5, 0x76, 0xbb, 0x92, 0xa9, 0x62, 0x64, 0x40,
	0x86, 0x9d, 0xb8, 0xd2, 0x94, 0x41, 0xeb, 0x1e, 0x45, 0x22, 0x55, 0xca, 0x1a, 0x03, 0x32, 0x6c,
	0xc4, 0x7b, 0x59, 0x74, 0x2d, 0x84, 0x93, 0x6e, 0x97, 0x20, 0x3b, 0xf0, 0x56, 0xa5, 0xe9, 0x05,
	0x74, 0x16, 0x5a, 0xa5, 0xc1, 0x3c, 0xf4, 0xe6, 0x4f, 0x81, 0x9e, 0x41, 0x73, 0x95, 0x21, 0x26,
	0xac, 0xe9, 0x9d, 0x20, 0x8a, 0x9e, 0x3b, 0x83, 0x6f, 0x3b, 0x54, 0xeb, 0x0f, 0x16, 0x85, 0x9e,
	0xaa, 0xc0, 0x9f, 0xa1, 0xf7, 0x98, 0x25, 0xc2, 0xe1, 0x12, 0xad, 0x15, 0x29, 0xfe, 0x09, 0x3d,
	0x85, 0xee, 0xfe, 0xb8, 0x07, 0xb5, 0xd1, 0x9e, 0xfc, 0x68, 0xde, 0x1d, 0xe7, 0xb3, 0xf1, 0xbe,
	0x1e, 0xd7, 0x12, 0xfc, 0x12, 0x7a, 0x31, 0xbe, 0xea, 0xfc, 0x3f, 0xe3, 0x39, 0x87, 0x76, 0x8c,
	0x36, 0xd3, 0xca, 0x22, 0x3d, 0x87, 0xc8, 0x3a, 0xe1, 0x76, 0xb6, 0x4c, 0x95, 0x8a, 0x9f, 0xc0,
	0xf1, 0x52, 0x2b, 0xe9, 0xb4, 0x29, 0x27, 0xce, 0x3f, 0x09, 0x74, 0x9f, 0x8a, 0xdf, 0xb0, 0x42,
	0x93, 0xcb, 0x35, 0xd2, 0x11, 0x44, 0xe1, 0x24, 0x7a, 0x5a, 0x90, 0xd5, 0xce, 0xeb, 0x7b, 0xd8,
	0x6a, 0xcb, 0x08, 0xa2, 0x80, 0x17, 0xa2, 0x35, 0xd4, 0x5f, 0xd1, 0x2b, 0x68, 0x95, 0x8b, 0x29,
	0x2d, 0x8c, 0x3a, 0x45, 0xbf, 0xf6, 0x11, 0xa6, 0xe4, 0x25, 0xf2, 0x6f, 0xe2, 0xfa, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0xf3, 0x8a, 0x60, 0xc4, 0x2f, 0x02, 0x00, 0x00,
}
