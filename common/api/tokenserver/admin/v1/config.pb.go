// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

package admin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// TokenServerConfig is read from tokenserver.cfg in luci-config.
type TokenServerConfig struct {
	// List of CAs we trust.
	CertificateAuthority []*CertificateAuthorityConfig `protobuf:"bytes,1,rep,name=certificate_authority,json=certificateAuthority" json:"certificate_authority,omitempty"`
}

func (m *TokenServerConfig) Reset()                    { *m = TokenServerConfig{} }
func (m *TokenServerConfig) String() string            { return proto.CompactTextString(m) }
func (*TokenServerConfig) ProtoMessage()               {}
func (*TokenServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TokenServerConfig) GetCertificateAuthority() []*CertificateAuthorityConfig {
	if m != nil {
		return m.CertificateAuthority
	}
	return nil
}

// CertificateAuthorityConfig defines a single CA we trust.
//
// Such CA issues certificates for nodes that use The Token Service. Each node
// has a private key and certificate with Common Name set to the FQDN of this
// node, e.g. "CN=slave43-c1.c.chromecompute.google.com.internal".
//
// The Token Server uses this CN to derive an identity string for a machine. It
// splits FQDN into a hostname ("slave43-c1") and a domain name
// ("c.chromecompute.google.com.internal"), searches for a domain name in
// "known_domains" set, and, if it is present, uses parameters described there
// for generating a token with machine_id <hostname>@<token-server-url>.
type CertificateAuthorityConfig struct {
	UniqueId int64  `protobuf:"varint,6,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	Cn       string `protobuf:"bytes,1,opt,name=cn" json:"cn,omitempty"`
	CertPath string `protobuf:"bytes,2,opt,name=cert_path,json=certPath" json:"cert_path,omitempty"`
	CrlUrl   string `protobuf:"bytes,3,opt,name=crl_url,json=crlUrl" json:"crl_url,omitempty"`
	UseOauth bool   `protobuf:"varint,4,opt,name=use_oauth,json=useOauth" json:"use_oauth,omitempty"`
	// KnownDomains describes parameters to use for each particular domain.
	KnownDomains []*DomainConfig `protobuf:"bytes,5,rep,name=known_domains,json=knownDomains" json:"known_domains,omitempty"`
}

func (m *CertificateAuthorityConfig) Reset()                    { *m = CertificateAuthorityConfig{} }
func (m *CertificateAuthorityConfig) String() string            { return proto.CompactTextString(m) }
func (*CertificateAuthorityConfig) ProtoMessage()               {}
func (*CertificateAuthorityConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CertificateAuthorityConfig) GetKnownDomains() []*DomainConfig {
	if m != nil {
		return m.KnownDomains
	}
	return nil
}

// DomainConfig is used inside CertificateAuthorityConfig.
type DomainConfig struct {
	// Domain is domain names of hosts this config applies to.
	Domain []string `protobuf:"bytes,1,rep,name=domain" json:"domain,omitempty"`
	// CloudProjectName is a name of Google Cloud Project to create service
	// accounts in (used for OAuth2 tokens).
	//
	// The Token Server's own service account must have Editor permission in this
	// project.
	CloudProjectName string `protobuf:"bytes,2,opt,name=cloud_project_name,json=cloudProjectName" json:"cloud_project_name,omitempty"`
	// AllowedOauth2Scope is a whitelist of OAuth2 scopes the token server is
	// willing to mint an OAuth2 access token with.
	AllowedOauth2Scope []string `protobuf:"bytes,3,rep,name=allowed_oauth2_scope,json=allowedOauth2Scope" json:"allowed_oauth2_scope,omitempty"`
	// MachineTokenLifetime is how long generated machine tokens live, in seconds.
	//
	// If 0, machine tokens are not allowed.
	MachineTokenLifetime int64 `protobuf:"varint,5,opt,name=machine_token_lifetime,json=machineTokenLifetime" json:"machine_token_lifetime,omitempty"`
	// Location is used as a part of machine_id in the machine token.
	//
	// It is the short alias for a domain machine belongs to. The machine identity
	// will be <hostname>@<location>.
	//
	// For example, if vm123-m4.golo.chromium.org is requesting a token, and there
	// is DomainConfig with domain == 'golo.chromium.org' and location == 'golo',
	// the resulting machine_id will be 'vm123-m4@golo'.
	Location string `protobuf:"bytes,6,opt,name=location" json:"location,omitempty"`
}

func (m *DomainConfig) Reset()                    { *m = DomainConfig{} }
func (m *DomainConfig) String() string            { return proto.CompactTextString(m) }
func (*DomainConfig) ProtoMessage()               {}
func (*DomainConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func init() {
	proto.RegisterType((*TokenServerConfig)(nil), "tokenserver.admin.TokenServerConfig")
	proto.RegisterType((*CertificateAuthorityConfig)(nil), "tokenserver.admin.CertificateAuthorityConfig")
	proto.RegisterType((*DomainConfig)(nil), "tokenserver.admin.DomainConfig")
}

var fileDescriptor1 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xe2, 0x30,
	0x10, 0xc6, 0x15, 0xfe, 0x64, 0x83, 0x97, 0x5d, 0x81, 0xc5, 0xb2, 0x11, 0x7b, 0x58, 0xc4, 0x89,
	0xc3, 0x6e, 0x54, 0xd1, 0xbe, 0x40, 0x05, 0x97, 0x56, 0x55, 0x8b, 0xa0, 0x3d, 0x5b, 0xc6, 0x31,
	0xc5, 0xc5, 0xb1, 0xa9, 0xe3, 0x14, 0xf5, 0x65, 0x7b, 0xe8, 0x93, 0xd4, 0x9e, 0x04, 0xb5, 0x52,
	0xe9, 0x2d, 0xf3, 0xfd, 0x26, 0x33, 0xf3, 0xcd, 0x18, 0xb5, 0x99, 0x56, 0x6b, 0x71, 0x9f, 0xec,
	0x8c, 0xb6, 0x1a, 0x77, 0xad, 0xde, 0x72, 0x95, 0x73, 0xf3, 0xc4, 0x4d, 0x42, 0xd3, 0x4c, 0xa8,
	0xd1, 0x1e, 0x75, 0x6f, 0xbd, 0xb8, 0x04, 0x71, 0x0a, 0xd9, 0x78, 0x85, 0x7e, 0x31, 0x6e, 0xac,
	0x58, 0x0b, 0x46, 0x2d, 0x27, 0xb4, 0xb0, 0x1b, 0x6d, 0x84, 0x7d, 0x8e, 0x83, 0x61, 0x7d, 0xfc,
	0x7d, 0xf2, 0x3f, 0xf9, 0x54, 0x27, 0x99, 0xbe, 0xe7, 0x9f, 0x1f, 0xd2, 0xcb, 0x6a, 0x8b, 0x1e,
	0x3b, 0xc2, 0x46, 0xaf, 0x01, 0x1a, 0x7c, 0xfd, 0x13, 0xfe, 0x83, 0x5a, 0x85, 0x12, 0x8f, 0x05,
	0x27, 0x22, 0x8d, 0xc3, 0x61, 0x30, 0xae, 0x2f, 0xa2, 0x52, 0xb8, 0x48, 0xf1, 0x4f, 0x54, 0x63,
	0xca, 0x0d, 0x13, 0x8c, 0x5b, 0x0b, 0xf7, 0xe5, 0x93, 0x7d, 0x0f, 0xb2, 0xa3, 0x76, 0x13, 0xd7,
	0x40, 0x8e, 0xbc, 0x30, 0x77, 0x31, 0xfe, 0x8d, 0xbe, 0x31, 0x23, 0x49, 0x61, 0x64, 0x5c, 0x07,
	0x14, 0xba, 0xf0, 0xce, 0x48, 0x68, 0x91, 0x73, 0xa2, 0xbd, 0xbd, 0xb8, 0xe1, 0x50, 0xe4, 0x5a,
	0xe4, 0xfc, 0xc6, 0xc7, 0x78, 0x86, 0x7e, 0x6c, 0x95, 0xde, 0x2b, 0x92, 0xea, 0x8c, 0x0a, 0x95,
	0xc7, 0x4d, 0xb0, 0xfe, 0xf7, 0x88, 0xf5, 0x19, 0x64, 0x54, 0x66, 0xdb, 0xf0, 0x57, 0x29, 0xe5,
	0xa3, 0x97, 0x00, 0xb5, 0x3f, 0x62, 0xdc, 0x47, 0x61, 0x59, 0x10, 0x56, 0xe9, 0x66, 0x29, 0x23,
	0xfc, 0x0f, 0x61, 0x26, 0x75, 0x91, 0x12, 0x77, 0xa8, 0x07, 0xce, 0x2c, 0x51, 0x34, 0xe3, 0x95,
	0x95, 0x0e, 0x90, 0x79, 0x09, 0xae, 0x9d, 0x8e, 0x4f, 0x50, 0x8f, 0x4a, 0xa9, 0xf7, 0x3c, 0x2d,
	0xa7, 0x9f, 0x90, 0x9c, 0xe9, 0x1d, 0x77, 0xfe, 0x7c, 0x4d, 0x5c, 0x31, 0x30, 0x32, 0x59, 0x7a,
	0x82, 0xcf, 0x50, 0x3f, 0xa3, 0x6c, 0x23, 0x14, 0x27, 0x60, 0x80, 0x48, 0xb1, 0xe6, 0x56, 0xb8,
	0x1e, 0x4d, 0xd8, 0x6d, 0xaf, 0xa2, 0xf0, 0x16, 0xae, 0x2a, 0x86, 0x07, 0x28, 0x92, 0xda, 0x5d,
	0x47, 0x68, 0x05, 0x37, 0x70, 0x6b, 0x3d, 0xc4, 0x97, 0x8d, 0xa8, 0xd1, 0x69, 0xae, 0x42, 0x78,
	0x58, 0xa7, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x63, 0xf3, 0xf3, 0x68, 0x02, 0x00, 0x00,
}
