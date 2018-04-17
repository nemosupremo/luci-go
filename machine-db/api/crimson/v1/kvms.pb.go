// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/kvms.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "go.chromium.org/luci/machine-db/api/common/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A KVM in the database.
type KVM struct {
	// The name of this KVM on the network. Uniquely identifies this KVM.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The VLAN this KVM belongs to.
	Vlan int64 `protobuf:"varint,2,opt,name=vlan" json:"vlan,omitempty"`
	// The type of platform this KVM is.
	Platform string `protobuf:"bytes,3,opt,name=platform" json:"platform,omitempty"`
	// The datacenter this KVM belongs to.
	Datacenter string `protobuf:"bytes,4,opt,name=datacenter" json:"datacenter,omitempty"`
	// The rack this KVM belongs to.
	Rack string `protobuf:"bytes,5,opt,name=rack" json:"rack,omitempty"`
	// A description of this KVM.
	Description string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	// The MAC address associated with this KVM.
	MacAddress string `protobuf:"bytes,7,opt,name=mac_address,json=macAddress" json:"mac_address,omitempty"`
	// The IPv4 address associated with this KVM.
	Ipv4 string `protobuf:"bytes,8,opt,name=ipv4" json:"ipv4,omitempty"`
	// The state of this KVM.
	State common.State `protobuf:"varint,9,opt,name=state,enum=common.State" json:"state,omitempty"`
}

func (m *KVM) Reset()                    { *m = KVM{} }
func (m *KVM) String() string            { return proto.CompactTextString(m) }
func (*KVM) ProtoMessage()               {}
func (*KVM) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *KVM) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KVM) GetVlan() int64 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *KVM) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *KVM) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

func (m *KVM) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *KVM) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *KVM) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *KVM) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *KVM) GetState() common.State {
	if m != nil {
		return m.State
	}
	return common.State_STATE_UNSPECIFIED
}

// A request to list KVMs in the database.
type ListKVMsRequest struct {
	// The names of KVMs to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	// The VLANs to filter retrieved KVMs on.
	Vlans []int64 `protobuf:"varint,2,rep,packed,name=vlans" json:"vlans,omitempty"`
	// The platforms to filter retrieved KVMs on.
	Platforms []string `protobuf:"bytes,3,rep,name=platforms" json:"platforms,omitempty"`
	// The datacenters to filter retrieved KVMs on.
	Datacenters []string `protobuf:"bytes,4,rep,name=datacenters" json:"datacenters,omitempty"`
	// The racks to filter retrieved KVMs on.
	Racks []string `protobuf:"bytes,5,rep,name=racks" json:"racks,omitempty"`
	// The MAC addresses to filter retrieved KVMs on.
	MacAddresses []string `protobuf:"bytes,6,rep,name=mac_addresses,json=macAddresses" json:"mac_addresses,omitempty"`
	// The IPv4 addresses to filter retrieved KVMs on.
	Ipv4S []string `protobuf:"bytes,7,rep,name=ipv4s" json:"ipv4s,omitempty"`
	// The states to filter retrieved KVMs on.
	States []common.State `protobuf:"varint,8,rep,packed,name=states,enum=common.State" json:"states,omitempty"`
}

func (m *ListKVMsRequest) Reset()                    { *m = ListKVMsRequest{} }
func (m *ListKVMsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListKVMsRequest) ProtoMessage()               {}
func (*ListKVMsRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ListKVMsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListKVMsRequest) GetVlans() []int64 {
	if m != nil {
		return m.Vlans
	}
	return nil
}

func (m *ListKVMsRequest) GetPlatforms() []string {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func (m *ListKVMsRequest) GetDatacenters() []string {
	if m != nil {
		return m.Datacenters
	}
	return nil
}

func (m *ListKVMsRequest) GetRacks() []string {
	if m != nil {
		return m.Racks
	}
	return nil
}

func (m *ListKVMsRequest) GetMacAddresses() []string {
	if m != nil {
		return m.MacAddresses
	}
	return nil
}

func (m *ListKVMsRequest) GetIpv4S() []string {
	if m != nil {
		return m.Ipv4S
	}
	return nil
}

func (m *ListKVMsRequest) GetStates() []common.State {
	if m != nil {
		return m.States
	}
	return nil
}

// A response containing a list of KVMs in the database.
type ListKVMsResponse struct {
	// The KVMs matching the request.
	Kvms []*KVM `protobuf:"bytes,1,rep,name=kvms" json:"kvms,omitempty"`
}

func (m *ListKVMsResponse) Reset()                    { *m = ListKVMsResponse{} }
func (m *ListKVMsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListKVMsResponse) ProtoMessage()               {}
func (*ListKVMsResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *ListKVMsResponse) GetKvms() []*KVM {
	if m != nil {
		return m.Kvms
	}
	return nil
}

func init() {
	proto.RegisterType((*KVM)(nil), "crimson.KVM")
	proto.RegisterType((*ListKVMsRequest)(nil), "crimson.ListKVMsRequest")
	proto.RegisterType((*ListKVMsResponse)(nil), "crimson.ListKVMsResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/kvms.proto", fileDescriptor5)
}

var fileDescriptor5 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0x13, 0x31,
	0x10, 0x85, 0xb5, 0xdd, 0x24, 0x4d, 0x26, 0x2d, 0x20, 0x8b, 0xc3, 0xa8, 0x42, 0x60, 0xa5, 0x42,
	0xda, 0x0b, 0xbb, 0xa2, 0xf4, 0x02, 0x37, 0xce, 0x21, 0x97, 0x45, 0xca, 0x15, 0xb9, 0x5e, 0xd3,
	0x5a, 0x8d, 0xd7, 0x8b, 0xc7, 0xc9, 0x3f, 0xe7, 0xc0, 0x0d, 0x79, 0x1c, 0x9a, 0x95, 0xb8, 0xf4,
	0x36, 0xef, 0x3d, 0xe7, 0x39, 0x9f, 0x77, 0xe0, 0xf3, 0xbd, 0xaf, 0xf5, 0x43, 0xf0, 0xce, 0xee,
	0x5d, 0xed, 0xc3, 0x7d, 0xb3, 0xdb, 0x6b, 0xdb, 0x38, 0xa5, 0x1f, 0x6c, 0x6f, 0x3e, 0x74, 0x77,
	0x8d, 0x1a, 0x6c, 0xa3, 0x83, 0x75, 0xe4, 0xfb, 0xe6, 0xf0, 0xb1, 0x79, 0x3c, 0x38, 0xaa, 0x87,
	0xe0, 0xa3, 0x17, 0xe7, 0x47, 0xfb, 0xea, 0xcb, 0xb3, 0x3a, 0xbc, 0x73, 0xb9, 0x82, 0xa2, 0x8a,
	0xe6, 0x58, 0xb2, 0xfa, 0x53, 0x40, 0xb9, 0xde, 0x6e, 0x84, 0x80, 0x49, 0xaf, 0x9c, 0xc1, 0x42,
	0x16, 0xd5, 0xa2, 0xe5, 0x39, 0x79, 0x87, 0x9d, 0xea, 0xf1, 0x4c, 0x16, 0x55, 0xd9, 0xf2, 0x2c,
	0xae, 0x60, 0x3e, 0xec, 0x54, 0xfc, 0xe9, 0x83, 0xc3, 0x92, 0xcf, 0x3e, 0x69, 0xf1, 0x16, 0xa0,
	0x53, 0x51, 0x69, 0xd3, 0x47, 0x13, 0x70, 0xc2, 0xe9, 0xc8, 0x49, 0x7d, 0x41, 0xe9, 0x47, 0x9c,
	0xe6, 0x3b, 0xd2, 0x2c, 0x24, 0x2c, 0x3b, 0x43, 0x3a, 0xd8, 0x21, 0x5a, 0xdf, 0xe3, 0x8c, 0xa3,
	0xb1, 0x25, 0xde, 0xc1, 0xd2, 0x29, 0xfd, 0x43, 0x75, 0x5d, 0x30, 0x44, 0x78, 0x9e, 0x6b, 0x9d,
	0xd2, 0x5f, 0xb3, 0x93, 0x6a, 0xed, 0x70, 0xb8, 0xc5, 0x79, 0xae, 0x4d, 0xb3, 0xb8, 0x86, 0x29,
	0x63, 0xe2, 0x42, 0x16, 0xd5, 0x8b, 0x9b, 0xcb, 0x3a, 0xe3, 0xd7, 0xdf, 0x93, 0xd9, 0xe6, 0x6c,
	0xf5, 0xbb, 0x80, 0x97, 0xdf, 0x2c, 0xc5, 0xf5, 0x76, 0x43, 0xad, 0xf9, 0xb5, 0x37, 0x14, 0xc5,
	0x6b, 0x98, 0x26, 0x76, 0xc2, 0x42, 0x96, 0xd5, 0xa2, 0xcd, 0x22, 0xb9, 0x89, 0x9e, 0xf0, 0x4c,
	0x96, 0x55, 0xd9, 0x66, 0x21, 0xde, 0xc0, 0xe2, 0x1f, 0x3b, 0x61, 0xc9, 0xe7, 0x4f, 0x06, 0x93,
	0x3d, 0xb1, 0x13, 0x4e, 0x38, 0x1f, 0x5b, 0xa9, 0x35, 0xbd, 0x01, 0xe1, 0x34, 0xdf, 0xc5, 0x42,
	0x5c, 0xc3, 0xe5, 0x88, 0xd7, 0x10, 0xce, 0x38, 0xbd, 0x38, 0x11, 0xe7, 0x3f, 0x94, 0x38, 0xd3,
	0x73, 0xf0, 0x4f, 0x59, 0x88, 0xf7, 0x30, 0xcb, 0x1f, 0x17, 0xe7, 0xb2, 0xfc, 0x1f, 0xfb, 0x18,
	0xae, 0x6e, 0xe1, 0xd5, 0x09, 0x9b, 0x06, 0xdf, 0x93, 0x11, 0x12, 0x26, 0x69, 0xb5, 0x18, 0x7b,
	0x79, 0x73, 0x51, 0x1f, 0x77, 0xab, 0x5e, 0x6f, 0x37, 0x2d, 0x27, 0x77, 0x33, 0x5e, 0x98, 0x4f,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xb5, 0xd5, 0x78, 0xb2, 0x02, 0x00, 0x00,
}