// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/luci_notify/api/config/settings.proto

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Settings is the top-level configuration message.
type Settings struct {
	// MiloHost is the hostname of the Milo instance luci-notify queries for
	// additional build information.
	//
	// Required.
	MiloHost string `protobuf:"bytes,1,opt,name=milo_host,json=miloHost" json:"milo_host,omitempty"`
}

func (m *Settings) Reset()                    { *m = Settings{} }
func (m *Settings) String() string            { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()               {}
func (*Settings) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Settings) GetMiloHost() string {
	if m != nil {
		return m.MiloHost
	}
	return ""
}

func init() {
	proto.RegisterType((*Settings)(nil), "notify.Settings")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/luci_notify/api/config/settings.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0x04, 0x13, 0xf1, 0x79, 0xf9, 0x25, 0x99, 0x69, 0x95, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0xc9, 0xf9,
	0x79, 0x69, 0x99, 0xe9, 0xfa, 0xc5, 0xa9, 0x25, 0x25, 0x99, 0x79, 0xe9, 0xc5, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x15, 0x4a, 0xea, 0x5c, 0x1c, 0xc1, 0x50, 0x19, 0x21, 0x69,
	0x2e, 0xce, 0xdc, 0xcc, 0x9c, 0xfc, 0xf8, 0x8c, 0xfc, 0xe2, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x0e, 0x90, 0x80, 0x47, 0x7e, 0x71, 0x89, 0x13, 0x47, 0x14, 0x1b, 0xc4, 0xa4, 0x24,
	0x36, 0xb0, 0x09, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xb9, 0xd4, 0x6c, 0x7f, 0x00,
	0x00, 0x00,
}
