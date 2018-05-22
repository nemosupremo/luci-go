// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/skylark/skylarkproto/testprotos/proto2.proto

package testprotos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Proto2Message struct {
	I                *int64  `protobuf:"varint,1,opt,name=i" json:"i,omitempty"`
	RepI             []int64 `protobuf:"varint,2,rep,name=rep_i,json=repI" json:"rep_i,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Proto2Message) Reset()                    { *m = Proto2Message{} }
func (m *Proto2Message) String() string            { return proto.CompactTextString(m) }
func (*Proto2Message) ProtoMessage()               {}
func (*Proto2Message) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Proto2Message) GetI() int64 {
	if m != nil && m.I != nil {
		return *m.I
	}
	return 0
}

func (m *Proto2Message) GetRepI() []int64 {
	if m != nil {
		return m.RepI
	}
	return nil
}

func init() {
	proto.RegisterType((*Proto2Message)(nil), "testprotos.Proto2Message")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/skylark/skylarkproto/testprotos/proto2.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x2f, 0xce, 0xae, 0xcc, 0x49, 0x2c, 0xca, 0x86, 0xd1, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xfa,
	0x25, 0xa9, 0xc5, 0x25, 0x60, 0x56, 0xb1, 0x3e, 0x98, 0x32, 0xd2, 0x03, 0x53, 0x42, 0x5c, 0x08,
	0x09, 0x25, 0x23, 0x2e, 0xde, 0x00, 0xb0, 0x9c, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x10,
	0x0f, 0x17, 0x63, 0xa6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x63, 0xa6, 0x90, 0x30, 0x17,
	0x6b, 0x51, 0x6a, 0x41, 0x7c, 0xa6, 0x04, 0x93, 0x02, 0xb3, 0x06, 0x73, 0x10, 0x4b, 0x51, 0x6a,
	0x81, 0x27, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x1e, 0xb8, 0xa2, 0x83, 0x00, 0x00, 0x00,
}