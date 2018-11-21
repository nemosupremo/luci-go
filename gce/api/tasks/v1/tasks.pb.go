// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/gce/api/tasks/v1/tasks.proto

package tasks

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	v1 "go.chromium.org/luci/gce/api/config/v1"
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

// A task to create or update a particular VM.
type Ensure struct {
	// The ID of the VM to create or update.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The attributes of the VM.
	Attributes           *v1.Block `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Ensure) Reset()         { *m = Ensure{} }
func (m *Ensure) String() string { return proto.CompactTextString(m) }
func (*Ensure) ProtoMessage()    {}
func (*Ensure) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{0}
}

func (m *Ensure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ensure.Unmarshal(m, b)
}
func (m *Ensure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ensure.Marshal(b, m, deterministic)
}
func (m *Ensure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ensure.Merge(m, src)
}
func (m *Ensure) XXX_Size() int {
	return xxx_messageInfo_Ensure.Size(m)
}
func (m *Ensure) XXX_DiscardUnknown() {
	xxx_messageInfo_Ensure.DiscardUnknown(m)
}

var xxx_messageInfo_Ensure proto.InternalMessageInfo

func (m *Ensure) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Ensure) GetAttributes() *v1.Block {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A task to expand a VMs config.
type Expand struct {
	// The ID of the VMs block to expand.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expand) Reset()         { *m = Expand{} }
func (m *Expand) String() string { return proto.CompactTextString(m) }
func (*Expand) ProtoMessage()    {}
func (*Expand) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{1}
}

func (m *Expand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expand.Unmarshal(m, b)
}
func (m *Expand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expand.Marshal(b, m, deterministic)
}
func (m *Expand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expand.Merge(m, src)
}
func (m *Expand) XXX_Size() int {
	return xxx_messageInfo_Expand.Size(m)
}
func (m *Expand) XXX_DiscardUnknown() {
	xxx_messageInfo_Expand.DiscardUnknown(m)
}

var xxx_messageInfo_Expand proto.InternalMessageInfo

func (m *Expand) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Ensure)(nil), "tasks.Ensure")
	proto.RegisterType((*Expand)(nil), "tasks.Expand")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/gce/api/tasks/v1/tasks.proto", fileDescriptor_f63d8744087b0bbc)
}

var fileDescriptor_f63d8744087b0bbc = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0x4f, 0x4e, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x49, 0x2c, 0xce, 0x2e, 0xd6, 0x2f,
	0x33, 0x84, 0x30, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0xc1, 0x1c, 0x29, 0x63, 0xbc,
	0x3a, 0x93, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0x41, 0x5a, 0x21, 0x2c, 0x88, 0x5e, 0x25, 0x77, 0x2e,
	0x36, 0xd7, 0xbc, 0xe2, 0xd2, 0xa2, 0x54, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x5d, 0x2e, 0xae, 0xc4, 0x92, 0x92, 0xa2, 0xcc,
	0xa4, 0xd2, 0x92, 0xd4, 0x62, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x5e, 0x3d, 0xa8, 0x66,
	0xa7, 0x9c, 0xfc, 0xe4, 0xec, 0x20, 0x24, 0x05, 0x4a, 0x12, 0x5c, 0x6c, 0xae, 0x15, 0x05, 0x89,
	0x79, 0x29, 0xe8, 0x06, 0x25, 0xb1, 0x81, 0x6d, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1c,
	0xda, 0x6a, 0x51, 0xda, 0x00, 0x00, 0x00,
}
