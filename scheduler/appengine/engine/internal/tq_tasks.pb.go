// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/engine/internal/tq_tasks.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/scheduler/appengine/engine/internal/tq_tasks.proto

It has these top-level messages:
	TickLaterTask
	StartInvocationTask
*/
package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TickLaterTask struct {
	JobId     string `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	TickNonce int64  `protobuf:"varint,2,opt,name=tick_nonce,json=tickNonce" json:"tick_nonce,omitempty"`
}

func (m *TickLaterTask) Reset()                    { *m = TickLaterTask{} }
func (m *TickLaterTask) String() string            { return proto.CompactTextString(m) }
func (*TickLaterTask) ProtoMessage()               {}
func (*TickLaterTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TickLaterTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *TickLaterTask) GetTickNonce() int64 {
	if m != nil {
		return m.TickNonce
	}
	return 0
}

type StartInvocationTask struct {
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *StartInvocationTask) Reset()                    { *m = StartInvocationTask{} }
func (m *StartInvocationTask) String() string            { return proto.CompactTextString(m) }
func (*StartInvocationTask) ProtoMessage()               {}
func (*StartInvocationTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StartInvocationTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func init() {
	proto.RegisterType((*TickLaterTask)(nil), "internal.tq_tasks.TickLaterTask")
	proto.RegisterType((*StartInvocationTask)(nil), "internal.tq_tasks.StartInvocationTask")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/engine/internal/tq_tasks.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0x31, 0x6b, 0x84, 0x40,
	0x10, 0x46, 0x31, 0x21, 0x82, 0x0b, 0x29, 0x62, 0x08, 0xd8, 0x04, 0xc4, 0xca, 0x22, 0xec, 0x16,
	0xf9, 0x0d, 0x21, 0x08, 0x21, 0x85, 0x67, 0x2f, 0xeb, 0xba, 0xe8, 0xb8, 0x3a, 0xe3, 0xed, 0x8e,
	0xf7, 0xfb, 0x0f, 0x8f, 0xb3, 0xbd, 0xea, 0x83, 0xf7, 0xe0, 0xe3, 0x89, 0xdf, 0x81, 0xa4, 0x19,
	0x3d, 0x2d, 0xb0, 0x2d, 0x92, 0xfc, 0xa0, 0xe6, 0xcd, 0x80, 0x0a, 0x66, 0xb4, 0xfd, 0x36, 0x5b,
	0xaf, 0xf4, 0xba, 0x5a, 0x1c, 0x00, 0xad, 0xba, 0x0f, 0x20, 0x5b, 0x8f, 0x7a, 0x56, 0x7c, 0x6e,
	0x59, 0x07, 0x17, 0xe4, 0xea, 0x89, 0x29, 0x7d, 0x3b, 0x84, 0x3c, 0x44, 0xf1, 0x23, 0x5e, 0x1b,
	0x30, 0xee, 0x4f, 0xb3, 0xf5, 0x8d, 0x0e, 0x2e, 0xfd, 0x10, 0xf1, 0x44, 0x5d, 0x0b, 0x7d, 0x16,
	0xe5, 0x51, 0x99, 0xd4, 0x2f, 0x13, 0x75, 0x55, 0x9f, 0x7e, 0x0a, 0xc1, 0x60, 0x5c, 0x8b, 0x84,
	0xc6, 0x66, 0x4f, 0x79, 0x54, 0x3e, 0xd7, 0xc9, 0x4e, 0xfe, 0x77, 0x50, 0x7c, 0x89, 0xf7, 0x13,
	0x6b, 0xcf, 0x15, 0x5e, 0xc8, 0x68, 0x06, 0xc2, 0x07, 0x67, 0x5d, 0x7c, 0xcb, 0xf9, 0xbe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x17, 0x48, 0xc1, 0x92, 0xd9, 0x00, 0x00, 0x00,
}