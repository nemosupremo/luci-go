// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/rpc.proto

package buildbucketpb

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A request message for GetBuild rpc.
type GetBuildRequest struct {
	// Build id.
	// Mutually exclusive with builder and number.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Builder of the build.
	// Requires number. Mutually exclusive with id.
	Builder *Builder_ID `protobuf:"bytes,2,opt,name=builder" json:"builder,omitempty"`
	// Build number.
	// Requires builder. Mutually exclusive with id.
	BuildNumber int32 `protobuf:"varint,3,opt,name=build_number,json=buildNumber" json:"build_number,omitempty"`
	// Fields to include in the response.
	// If not set, the default mask is used, see Build message comments for the
	// list of fields returned by default.
	//
	// Supports advanced semantics, see
	// https://chromium.googlesource.com/infra/luci/luci-py/+/f9ae69a37c4bdd0e08a8b0f7e123f6e403e774eb/appengine/components/components/protoutil/field_masks.py#7
	// In particular, if the client needs only some output properties, they
	// can be requested with paths "output.properties.fields.foo".
	Fields *google_protobuf2.FieldMask `protobuf:"bytes,100,opt,name=fields" json:"fields,omitempty"`
}

func (m *GetBuildRequest) Reset()                    { *m = GetBuildRequest{} }
func (m *GetBuildRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBuildRequest) ProtoMessage()               {}
func (*GetBuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *GetBuildRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetBuildRequest) GetBuilder() *Builder_ID {
	if m != nil {
		return m.Builder
	}
	return nil
}

func (m *GetBuildRequest) GetBuildNumber() int32 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *GetBuildRequest) GetFields() *google_protobuf2.FieldMask {
	if m != nil {
		return m.Fields
	}
	return nil
}

// A request message for SearchBuilds rpc.
type SearchBuildsRequest struct {
	// Returned builds must satisfy this predicate. Required.
	Predicate *BuildPredicate `protobuf:"bytes,1,opt,name=predicate" json:"predicate,omitempty"`
	// Fields to include in the response, see GetBuildRequest.fields.
	// Note that this applies to the response, not each build, so e.g. steps must
	// be requested with a path "builds.*.steps".
	Fields *google_protobuf2.FieldMask `protobuf:"bytes,100,opt,name=fields" json:"fields,omitempty"`
	// Number of builds to return.
	// Any value >100 is interpreted as 100.
	PageSize int32 `protobuf:"varint,101,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Value of SearchBuildsResponse.next_page_token from the previous response.
	// Use it to continue searching.
	PageToken string `protobuf:"bytes,102,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *SearchBuildsRequest) Reset()                    { *m = SearchBuildsRequest{} }
func (m *SearchBuildsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchBuildsRequest) ProtoMessage()               {}
func (*SearchBuildsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *SearchBuildsRequest) GetPredicate() *BuildPredicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func (m *SearchBuildsRequest) GetFields() *google_protobuf2.FieldMask {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchBuildsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchBuildsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// A response message for SearchBuilds rpc.
type SearchBuildsResponse struct {
	// Search results.
	//
	// Ordered by build id, descending. Ids are monotonically decreasing, so in
	// other words the order is newest-to-oldest.
	Builds []*Build `protobuf:"bytes,1,rep,name=builds" json:"builds,omitempty"`
	// Value for SearchBuildsRequest.page_token to continue searching.
	NextPageToken string `protobuf:"bytes,100,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *SearchBuildsResponse) Reset()                    { *m = SearchBuildsResponse{} }
func (m *SearchBuildsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchBuildsResponse) ProtoMessage()               {}
func (*SearchBuildsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *SearchBuildsResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

func (m *SearchBuildsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// A build predicate.
//
// At least one of the following fields is required: builder, gerrit_changes and
// git_commits..
// If a field value is empty, it is ignored, unless stated otherwise.
type BuildPredicate struct {
	// A build must be in this builder.
	Builder *Builder_ID `protobuf:"bytes,1,opt,name=builder" json:"builder,omitempty"`
	// A build must have this status.
	Status Status `protobuf:"varint,2,opt,name=status,enum=buildbucket.v2.Status" json:"status,omitempty"`
	// A build's Build.Input.gerrit_changes must include ALL of these changes.
	GerritChanges []*GerritChange `protobuf:"bytes,3,rep,name=gerrit_changes,json=gerritChanges" json:"gerrit_changes,omitempty"`
	// A build's Build.Input.gitiles_commits must include ALL of these hex sha1s.
	// Not to be confused with blamelist.
	GitCommits []string `protobuf:"bytes,4,rep,name=git_commits,json=gitCommits" json:"git_commits,omitempty"`
	// A build must be created by this identity.
	CreatedBy string `protobuf:"bytes,5,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	// A build must have ALL of these tags.
	// For "ANY of these tags" make separate RPCs.
	Tags []*StringPair `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	// A build must have been created within the specified range.
	// Both boundaries are optional.
	CreateTime *TimeRange `protobuf:"bytes,7,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// If false (default), a build must be non-experimental.
	// Otherwise it may be experimental or non-experimental.
	IncludeExperimental bool `protobuf:"varint,8,opt,name=include_experimental,json=includeExperimental" json:"include_experimental,omitempty"`
}

func (m *BuildPredicate) Reset()                    { *m = BuildPredicate{} }
func (m *BuildPredicate) String() string            { return proto.CompactTextString(m) }
func (*BuildPredicate) ProtoMessage()               {}
func (*BuildPredicate) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *BuildPredicate) GetBuilder() *Builder_ID {
	if m != nil {
		return m.Builder
	}
	return nil
}

func (m *BuildPredicate) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (m *BuildPredicate) GetGerritChanges() []*GerritChange {
	if m != nil {
		return m.GerritChanges
	}
	return nil
}

func (m *BuildPredicate) GetGitCommits() []string {
	if m != nil {
		return m.GitCommits
	}
	return nil
}

func (m *BuildPredicate) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *BuildPredicate) GetTags() []*StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *BuildPredicate) GetCreateTime() *TimeRange {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *BuildPredicate) GetIncludeExperimental() bool {
	if m != nil {
		return m.IncludeExperimental
	}
	return false
}

func init() {
	proto.RegisterType((*GetBuildRequest)(nil), "buildbucket.v2.GetBuildRequest")
	proto.RegisterType((*SearchBuildsRequest)(nil), "buildbucket.v2.SearchBuildsRequest")
	proto.RegisterType((*SearchBuildsResponse)(nil), "buildbucket.v2.SearchBuildsResponse")
	proto.RegisterType((*BuildPredicate)(nil), "buildbucket.v2.BuildPredicate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Builds service

type BuildsClient interface {
	// Gets a build.
	//
	// By default the returned build does not include all fields.
	// See GetBuildRequest.fields.
	GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*Build, error)
	// Searches for builds.
	SearchBuilds(ctx context.Context, in *SearchBuildsRequest, opts ...grpc.CallOption) (*SearchBuildsResponse, error)
}
type buildsPRPCClient struct {
	client *prpc.Client
}

func NewBuildsPRPCClient(client *prpc.Client) BuildsClient {
	return &buildsPRPCClient{client}
}

func (c *buildsPRPCClient) GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*Build, error) {
	out := new(Build)
	err := c.client.Call(ctx, "buildbucket.v2.Builds", "GetBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsPRPCClient) SearchBuilds(ctx context.Context, in *SearchBuildsRequest, opts ...grpc.CallOption) (*SearchBuildsResponse, error) {
	out := new(SearchBuildsResponse)
	err := c.client.Call(ctx, "buildbucket.v2.Builds", "SearchBuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type buildsClient struct {
	cc *grpc.ClientConn
}

func NewBuildsClient(cc *grpc.ClientConn) BuildsClient {
	return &buildsClient{cc}
}

func (c *buildsClient) GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*Build, error) {
	out := new(Build)
	err := grpc.Invoke(ctx, "/buildbucket.v2.Builds/GetBuild", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) SearchBuilds(ctx context.Context, in *SearchBuildsRequest, opts ...grpc.CallOption) (*SearchBuildsResponse, error) {
	out := new(SearchBuildsResponse)
	err := grpc.Invoke(ctx, "/buildbucket.v2.Builds/SearchBuilds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Builds service

type BuildsServer interface {
	// Gets a build.
	//
	// By default the returned build does not include all fields.
	// See GetBuildRequest.fields.
	GetBuild(context.Context, *GetBuildRequest) (*Build, error)
	// Searches for builds.
	SearchBuilds(context.Context, *SearchBuildsRequest) (*SearchBuildsResponse, error)
}

func RegisterBuildsServer(s prpc.Registrar, srv BuildsServer) {
	s.RegisterService(&_Builds_serviceDesc, srv)
}

func _Builds_GetBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).GetBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbucket.v2.Builds/GetBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).GetBuild(ctx, req.(*GetBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_SearchBuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).SearchBuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbucket.v2.Builds/SearchBuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).SearchBuilds(ctx, req.(*SearchBuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildbucket.v2.Builds",
	HandlerType: (*BuildsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuild",
			Handler:    _Builds_GetBuild_Handler,
		},
		{
			MethodName: "SearchBuilds",
			Handler:    _Builds_SearchBuilds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/buildbucket/proto/rpc.proto",
}

func init() { proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/rpc.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x9b, 0x36, 0x4d, 0x26, 0x6d, 0x2a, 0x6d, 0x0b, 0x5a, 0xc2, 0x47, 0x4d, 0x40, 0xc8,
	0x17, 0x1c, 0x70, 0x2b, 0x0e, 0xc0, 0x29, 0x2d, 0x54, 0x1c, 0x40, 0xd5, 0xb6, 0x27, 0x38, 0x58,
	0xfe, 0x98, 0xba, 0xab, 0xc4, 0x5e, 0xb3, 0xbb, 0x46, 0x6d, 0x7f, 0x0a, 0x7f, 0x01, 0x89, 0x3f,
	0xc1, 0x1f, 0x43, 0x5e, 0x3b, 0xe0, 0x34, 0x15, 0x8a, 0xb8, 0x39, 0x6f, 0xde, 0xdb, 0x79, 0x6f,
	0x76, 0x36, 0xf0, 0x22, 0x11, 0x6e, 0x74, 0x21, 0x45, 0xca, 0x8b, 0xd4, 0x15, 0x32, 0x19, 0x4d,
	0x8b, 0x88, 0x8f, 0xc2, 0x82, 0x4f, 0xe3, 0xb0, 0x88, 0x26, 0xa8, 0x47, 0xb9, 0x14, 0x5a, 0x8c,
	0x64, 0x1e, 0xb9, 0xe6, 0x8b, 0xf4, 0x1b, 0x45, 0xf7, 0x9b, 0x37, 0xb0, 0x13, 0x21, 0x92, 0x29,
	0x56, 0xbc, 0xb0, 0x38, 0x1f, 0x9d, 0x73, 0x9c, 0xc6, 0x7e, 0x1a, 0xa8, 0x49, 0xa5, 0x18, 0xec,
	0x2f, 0xd9, 0x23, 0x12, 0x69, 0x2a, 0xb2, 0x5a, 0xe4, 0x2d, 0x29, 0x32, 0x48, 0xa5, 0x19, 0xfe,
	0xb4, 0x60, 0xfb, 0x18, 0xf5, 0xb8, 0x84, 0x18, 0x7e, 0x2d, 0x50, 0x69, 0xd2, 0x87, 0x55, 0x1e,
	0x53, 0xcb, 0xb6, 0x9c, 0x16, 0x5b, 0xe5, 0x31, 0x39, 0x80, 0x0d, 0x23, 0x41, 0x49, 0x57, 0x6d,
	0xcb, 0xe9, 0x79, 0x03, 0x77, 0x3e, 0x90, 0x3b, 0xae, 0xca, 0xee, 0x87, 0x23, 0x36, 0xa3, 0x92,
	0xc7, 0xb0, 0x69, 0x3e, 0xfd, 0xac, 0x48, 0x43, 0x94, 0xb4, 0x65, 0x5b, 0xce, 0x3a, 0xeb, 0x19,
	0xec, 0x93, 0x81, 0x88, 0x07, 0x6d, 0x93, 0x5c, 0xd1, 0xb8, 0x3e, 0xb7, 0x1a, 0x8c, 0x3b, 0x1b,
	0x8c, 0xfb, 0xbe, 0x2c, 0x7f, 0x0c, 0xd4, 0x84, 0xd5, 0xcc, 0xe1, 0x2f, 0x0b, 0x76, 0x4e, 0x31,
	0x90, 0xd1, 0x85, 0x69, 0xaa, 0x66, 0xa6, 0xdf, 0x42, 0x37, 0x97, 0x18, 0xf3, 0x28, 0xd0, 0x68,
	0xbc, 0xf7, 0xbc, 0x47, 0xb7, 0xda, 0x3c, 0x99, 0xb1, 0xd8, 0x5f, 0xc1, 0xff, 0x38, 0x21, 0xf7,
	0xa1, 0x9b, 0x07, 0x09, 0xfa, 0x8a, 0x5f, 0x23, 0x45, 0x93, 0xae, 0x53, 0x02, 0xa7, 0xfc, 0x1a,
	0xc9, 0x43, 0x00, 0x53, 0xd4, 0x62, 0x82, 0x19, 0x3d, 0xb7, 0x2d, 0xa7, 0xcb, 0x0c, 0xfd, 0xac,
	0x04, 0x86, 0x29, 0xec, 0xce, 0x87, 0x50, 0xb9, 0xc8, 0x14, 0x92, 0xe7, 0xd0, 0x36, 0x9e, 0x15,
	0xb5, 0xec, 0x96, 0xd3, 0xf3, 0xee, 0xdc, 0x1a, 0x81, 0xd5, 0x24, 0xf2, 0x0c, 0xb6, 0x33, 0xbc,
	0xd4, 0x7e, 0xa3, 0x55, 0x6c, 0x5a, 0x6d, 0x95, 0xf0, 0xc9, 0x9f, 0x76, 0xdf, 0x5b, 0xd0, 0x9f,
	0x0f, 0xdf, 0xbc, 0x54, 0x6b, 0xf9, 0x4b, 0x75, 0xa1, 0xad, 0x74, 0xa0, 0x0b, 0x65, 0x36, 0xa1,
	0xef, 0xdd, 0xbd, 0x29, 0x3a, 0x35, 0x55, 0x56, 0xb3, 0xc8, 0x21, 0xf4, 0x13, 0x94, 0x92, 0x6b,
	0x3f, 0xba, 0x08, 0xb2, 0x04, 0x15, 0x6d, 0x99, 0x5c, 0x0f, 0x6e, 0xea, 0x8e, 0x0d, 0xeb, 0xd0,
	0x90, 0xd8, 0x56, 0xd2, 0xf8, 0xa5, 0xc8, 0x1e, 0xf4, 0x92, 0xf2, 0x04, 0x91, 0xa6, 0x5c, 0x2b,
	0xba, 0x66, 0xb7, 0x9c, 0x2e, 0x83, 0x84, 0xeb, 0xc3, 0x0a, 0x29, 0x87, 0x1d, 0x49, 0x0c, 0x34,
	0xc6, 0x7e, 0x78, 0x45, 0xd7, 0xab, 0x61, 0xd7, 0xc8, 0xf8, 0x8a, 0xb8, 0xb0, 0xa6, 0x83, 0x44,
	0xd1, 0xb6, 0x69, 0x3d, 0x58, 0xb4, 0x2c, 0x79, 0x96, 0x9c, 0x04, 0x5c, 0x32, 0xc3, 0x23, 0xaf,
	0xa1, 0x57, 0x89, 0x7d, 0xcd, 0x53, 0xa4, 0x1b, 0x66, 0x3c, 0xf7, 0x6e, 0xca, 0xce, 0x78, 0x8a,
	0xcc, 0xd8, 0xad, 0x9b, 0x97, 0x00, 0x79, 0x09, 0xbb, 0x3c, 0x8b, 0xa6, 0x45, 0x8c, 0x3e, 0x5e,
	0xe6, 0x28, 0x79, 0x8a, 0x99, 0x0e, 0xa6, 0xb4, 0x63, 0x5b, 0x4e, 0x87, 0xed, 0xd4, 0xb5, 0x77,
	0x8d, 0x92, 0xf7, 0xc3, 0x82, 0x76, 0xb5, 0x06, 0xe4, 0x08, 0x3a, 0xb3, 0xc7, 0x48, 0xf6, 0x16,
	0x47, 0x34, 0xf7, 0x4c, 0x07, 0xb7, 0xef, 0xc6, 0x70, 0x85, 0x7c, 0x81, 0xcd, 0xe6, 0x72, 0x91,
	0x27, 0x0b, 0x89, 0x17, 0xdf, 0xcf, 0xe0, 0xe9, 0xbf, 0x49, 0xd5, 0x7e, 0x0e, 0x57, 0xc6, 0xaf,
	0x3e, 0x1f, 0x2c, 0xf7, 0x37, 0xf3, 0xa6, 0x81, 0xe4, 0x61, 0xd8, 0x36, 0xe0, 0xfe, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0xf1, 0x18, 0xdd, 0x3e, 0x05, 0x00, 0x00,
}