// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/proto/gitiles/gitiles.proto

package gitiles

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	git "go.chromium.org/luci/common/proto/git"
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

// List copied from
// https://github.com/google/gitiles/blob/65edbe49f2b3882a5979f602383ef0c7b2b8ee0c/java/com/google/gitiles/ArchiveFormat.java
type ArchiveRequest_Format int32

const (
	ArchiveRequest_Invalid ArchiveRequest_Format = 0
	ArchiveRequest_GZIP    ArchiveRequest_Format = 1
	ArchiveRequest_TAR     ArchiveRequest_Format = 2
	ArchiveRequest_BZIP2   ArchiveRequest_Format = 3
	ArchiveRequest_XZ      ArchiveRequest_Format = 4
)

var ArchiveRequest_Format_name = map[int32]string{
	0: "Invalid",
	1: "GZIP",
	2: "TAR",
	3: "BZIP2",
	4: "XZ",
}

var ArchiveRequest_Format_value = map[string]int32{
	"Invalid": 0,
	"GZIP":    1,
	"TAR":     2,
	"BZIP2":   3,
	"XZ":      4,
}

func (x ArchiveRequest_Format) String() string {
	return proto.EnumName(ArchiveRequest_Format_name, int32(x))
}

func (ArchiveRequest_Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cdf9f15dd6bc12a, []int{4, 0}
}

// LogRequest is request message for Gitiles.Log rpc.
type LogRequest struct {
	// Gitiles project, e.g. "chromium/src" part in
	// https://chromium.googlesource.com/chromium/src/+/master
	// Required.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The commit where to start the listing from.
	// The value can be:
	//   - a git revision as 40-char string or its prefix so long as its unique in repo.
	//   - a ref such as "refs/heads/branch" or just "branch"
	//   - a ref defined as n-th parent of R in the form "R~n".
	//     For example, "master~2" or "deadbeef~1".
	// Required.
	Committish string `protobuf:"bytes,3,opt,name=committish,proto3" json:"committish,omitempty"`
	// If specified, only commits not reachable from this commit (inclusive)
	// will be returned.
	//
	// In git's notation, this is
	//   $ git log ^exclude_ancestors_of committish
	//  OR
	//   $ git log exclude_ancestors_of..committish
	// https://git-scm.com/docs/gitrevisions#gitrevisions-Theememtwo-dotRangeNotation
	//
	// For example, given this repo
	//
	//     base -> A -> B -> C == refs/heads/master
	//        \
	//         X -> Y -> Z  == refs/heads/release
	//
	// calling Log(committish='refs/heads/release',
	//             exclude_ancestors_of='refs/heads/master')
	// will return ['Z', Y', 'X'].
	ExcludeAncestorsOf string `protobuf:"bytes,2,opt,name=exclude_ancestors_of,json=excludeAncestorsOf,proto3" json:"exclude_ancestors_of,omitempty"`
	// If true, include tree diff in commits.
	TreeDiff bool `protobuf:"varint,4,opt,name=tree_diff,json=treeDiff,proto3" json:"tree_diff,omitempty"`
	// Value of next_page_token in LogResponse to continue.
	PageToken string `protobuf:"bytes,10,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// If > 0, number of commits to retrieve.
	PageSize             int32    `protobuf:"varint,11,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cdf9f15dd6bc12a, []int{0}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *LogRequest) GetCommittish() string {
	if m != nil {
		return m.Committish
	}
	return ""
}

func (m *LogRequest) GetExcludeAncestorsOf() string {
	if m != nil {
		return m.ExcludeAncestorsOf
	}
	return ""
}

func (m *LogRequest) GetTreeDiff() bool {
	if m != nil {
		return m.TreeDiff
	}
	return false
}

func (m *LogRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *LogRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// LogRequest is response message for Gitiles.Log rpc.
type LogResponse struct {
	// Retrieved commits.
	Log []*git.Commit `protobuf:"bytes,1,rep,name=log,proto3" json:"log,omitempty"`
	// A page token for next LogRequest to fetch next page of commits.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cdf9f15dd6bc12a, []int{1}
}

func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (m *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(m, src)
}
func (m *LogResponse) XXX_Size() int {
	return xxx_messageInfo_LogResponse.Size(m)
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetLog() []*git.Commit {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *LogResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// RefsRequest is a request message of Gitiles.Refs RPC.
type RefsRequest struct {
	// Gitiles project, e.g. "chromium/src" part in
	// https://chromium.googlesource.com/chromium/src/+/master
	// Required.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Limits which refs to resolve to only those matching {refsPath}/*.
	//
	// Must be "refs" or start with "refs/".
	// Must not include glob '*'.
	// Use "refs/heads" to retrieve all branches.
	//
	// To fetch **all** refs in a repo, specify just "refs" but beware of two
	// caveats:
	//  * refs returned include a ref for each patchset for each Gerrit change
	//    associated with the repo.
	//  * returned map will contain special "HEAD" ref whose value in resulting map
	//    will be name of the actual ref to which "HEAD" points, which is typically
	//    "refs/heads/master".
	//
	// Thus, if you are looking for all tags and all branches of repo, it's
	// recommended to issue two Refs calls limited to "refs/tags" and "refs/heads"
	// instead of one call for "refs".
	//
	// Since Gerrit allows per-ref ACLs, it is possible that some refs matching
	// refPrefix would not be present in results because current user isn't granted
	// read permission on them.
	RefsPath             string   `protobuf:"bytes,2,opt,name=refs_path,json=refsPath,proto3" json:"refs_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefsRequest) Reset()         { *m = RefsRequest{} }
func (m *RefsRequest) String() string { return proto.CompactTextString(m) }
func (*RefsRequest) ProtoMessage()    {}
func (*RefsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cdf9f15dd6bc12a, []int{2}
}

func (m *RefsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefsRequest.Unmarshal(m, b)
}
func (m *RefsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefsRequest.Marshal(b, m, deterministic)
}
func (m *RefsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefsRequest.Merge(m, src)
}
func (m *RefsRequest) XXX_Size() int {
	return xxx_messageInfo_RefsRequest.Size(m)
}
func (m *RefsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefsRequest proto.InternalMessageInfo

func (m *RefsRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *RefsRequest) GetRefsPath() string {
	if m != nil {
		return m.RefsPath
	}
	return ""
}

// RefsResponse is a response message of Gitiles.Refs RPC.
type RefsResponse struct {
	// revisions maps a ref to a revision.
	// Git branches have keys start with "refs/heads/".
	Revisions            map[string]string `protobuf:"bytes,2,rep,name=revisions,proto3" json:"revisions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RefsResponse) Reset()         { *m = RefsResponse{} }
func (m *RefsResponse) String() string { return proto.CompactTextString(m) }
func (*RefsResponse) ProtoMessage()    {}
func (*RefsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cdf9f15dd6bc12a, []int{3}
}

func (m *RefsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefsResponse.Unmarshal(m, b)
}
func (m *RefsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefsResponse.Marshal(b, m, deterministic)
}
func (m *RefsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefsResponse.Merge(m, src)
}
func (m *RefsResponse) XXX_Size() int {
	return xxx_messageInfo_RefsResponse.Size(m)
}
func (m *RefsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefsResponse proto.InternalMessageInfo

func (m *RefsResponse) GetRevisions() map[string]string {
	if m != nil {
		return m.Revisions
	}
	return nil
}

// ArchiveRequest is a request message of the Gitiles.Archive RPC.
type ArchiveRequest struct {
	// Gitiles project, e.g. "chromium/src" part in
	// https://chromium.googlesource.com/chromium/src/+/master
	// Required.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The ref at which to generate the project archive for.
	//
	// viz refs/for/branch or just branch
	Ref string `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	// Format of the returned project archive.
	Format               ArchiveRequest_Format `protobuf:"varint,3,opt,name=format,proto3,enum=gitiles.ArchiveRequest_Format" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ArchiveRequest) Reset()         { *m = ArchiveRequest{} }
func (m *ArchiveRequest) String() string { return proto.CompactTextString(m) }
func (*ArchiveRequest) ProtoMessage()    {}
func (*ArchiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cdf9f15dd6bc12a, []int{4}
}

func (m *ArchiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveRequest.Unmarshal(m, b)
}
func (m *ArchiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveRequest.Marshal(b, m, deterministic)
}
func (m *ArchiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveRequest.Merge(m, src)
}
func (m *ArchiveRequest) XXX_Size() int {
	return xxx_messageInfo_ArchiveRequest.Size(m)
}
func (m *ArchiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveRequest proto.InternalMessageInfo

func (m *ArchiveRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ArchiveRequest) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *ArchiveRequest) GetFormat() ArchiveRequest_Format {
	if m != nil {
		return m.Format
	}
	return ArchiveRequest_Invalid
}

type ArchiveResponse struct {
	// Suggested name of the returned archive.
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	// Contents of the archive streamed from gitiles.
	//
	// The underlying server RPC streams back the contents. This API simplifies
	// the RPC to a non-streaming response.
	Contents             []byte   `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArchiveResponse) Reset()         { *m = ArchiveResponse{} }
func (m *ArchiveResponse) String() string { return proto.CompactTextString(m) }
func (*ArchiveResponse) ProtoMessage()    {}
func (*ArchiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cdf9f15dd6bc12a, []int{5}
}

func (m *ArchiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveResponse.Unmarshal(m, b)
}
func (m *ArchiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveResponse.Marshal(b, m, deterministic)
}
func (m *ArchiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveResponse.Merge(m, src)
}
func (m *ArchiveResponse) XXX_Size() int {
	return xxx_messageInfo_ArchiveResponse.Size(m)
}
func (m *ArchiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveResponse proto.InternalMessageInfo

func (m *ArchiveResponse) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *ArchiveResponse) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterEnum("gitiles.ArchiveRequest_Format", ArchiveRequest_Format_name, ArchiveRequest_Format_value)
	proto.RegisterType((*LogRequest)(nil), "gitiles.LogRequest")
	proto.RegisterType((*LogResponse)(nil), "gitiles.LogResponse")
	proto.RegisterType((*RefsRequest)(nil), "gitiles.RefsRequest")
	proto.RegisterType((*RefsResponse)(nil), "gitiles.RefsResponse")
	proto.RegisterMapType((map[string]string)(nil), "gitiles.RefsResponse.RevisionsEntry")
	proto.RegisterType((*ArchiveRequest)(nil), "gitiles.ArchiveRequest")
	proto.RegisterType((*ArchiveResponse)(nil), "gitiles.ArchiveResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/proto/gitiles/gitiles.proto", fileDescriptor_4cdf9f15dd6bc12a)
}

var fileDescriptor_4cdf9f15dd6bc12a = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0x8d, 0x2c, 0xc7, 0x92, 0x47, 0xf9, 0x12, 0xb1, 0x9f, 0x4b, 0x85, 0x43, 0x82, 0x11, 0xa5,
	0xf8, 0x4a, 0x2e, 0x2a, 0xfd, 0xa1, 0x2d, 0x85, 0xa4, 0x69, 0x83, 0x21, 0x50, 0xa3, 0xe6, 0xa2,
	0xe4, 0x46, 0xa8, 0xf2, 0x48, 0xde, 0x46, 0xd2, 0xba, 0xda, 0xb5, 0x49, 0xf2, 0x14, 0x7d, 0x95,
	0x42, 0x5f, 0xa4, 0x6f, 0x54, 0x76, 0x25, 0x39, 0x76, 0x7f, 0x48, 0xaf, 0xb4, 0x73, 0xce, 0xcc,
	0xd9, 0x33, 0x33, 0x5a, 0x78, 0x96, 0x32, 0x2f, 0x9e, 0x95, 0x2c, 0xa7, 0x8b, 0xdc, 0x63, 0x65,
	0x3a, 0xca, 0x16, 0x31, 0x1d, 0xc5, 0x2c, 0xcf, 0x59, 0x31, 0x9a, 0x97, 0x4c, 0xb0, 0x51, 0x4a,
	0x05, 0xcd, 0x90, 0x37, 0x5f, 0x4f, 0xa1, 0xc4, 0xa8, 0xc3, 0xbe, 0xff, 0x4f, 0x0a, 0x0a, 0xa0,
	0xa2, 0x2a, 0x76, 0x7f, 0x68, 0x00, 0x67, 0x2c, 0x0d, 0xf0, 0xcb, 0x02, 0xb9, 0x20, 0x0e, 0x18,
	0xf3, 0x92, 0x7d, 0xc6, 0x58, 0x38, 0xda, 0x40, 0x1b, 0x76, 0x83, 0x26, 0x24, 0x87, 0x00, 0x55,
	0xa1, 0xa0, 0x7c, 0xe6, 0xe8, 0x8a, 0x5c, 0x43, 0xc8, 0x23, 0xe8, 0xe1, 0x55, 0x9c, 0x2d, 0xa6,
	0x18, 0x46, 0x45, 0x8c, 0x5c, 0xb0, 0x92, 0x87, 0x2c, 0x71, 0x5a, 0x2a, 0x93, 0xd4, 0xdc, 0x51,
	0x43, 0xbd, 0x4f, 0xc8, 0x3e, 0x74, 0x45, 0x89, 0x18, 0x4e, 0x69, 0x92, 0x38, 0xed, 0x81, 0x36,
	0x34, 0x03, 0x53, 0x02, 0x27, 0x34, 0x49, 0xc8, 0x01, 0xc0, 0x3c, 0x4a, 0x31, 0x14, 0xec, 0x12,
	0x0b, 0x07, 0x94, 0x48, 0x57, 0x22, 0xe7, 0x12, 0x90, 0xb5, 0x8a, 0xe6, 0xf4, 0x06, 0x1d, 0x6b,
	0xa0, 0x0d, 0xb7, 0x03, 0x53, 0x02, 0x1f, 0xe8, 0x0d, 0xba, 0xe7, 0x60, 0xa9, 0x96, 0xf8, 0x9c,
	0x15, 0x1c, 0xc9, 0x01, 0xe8, 0x19, 0x4b, 0x1d, 0x6d, 0xa0, 0x0f, 0x2d, 0xdf, 0xf2, 0x52, 0x2a,
	0xbc, 0x37, 0xca, 0x77, 0x20, 0x71, 0xf2, 0x10, 0xf6, 0x0a, 0xbc, 0x12, 0xe1, 0xda, 0x75, 0x95,
	0xe7, 0xff, 0x24, 0x3c, 0x69, 0xae, 0x74, 0x4f, 0xc0, 0x0a, 0x30, 0xe1, 0x77, 0x4f, 0x6a, 0x1f,
	0xba, 0x25, 0x26, 0x3c, 0x9c, 0x47, 0x62, 0x56, 0x4b, 0x99, 0x12, 0x98, 0x44, 0x62, 0xe6, 0x7e,
	0xd5, 0x60, 0xa7, 0x92, 0xa9, 0xdd, 0x1d, 0xcb, 0xec, 0x25, 0xe5, 0x94, 0x15, 0xdc, 0x69, 0x29,
	0x8f, 0x0f, 0xbc, 0x66, 0xc1, 0xeb, 0x99, 0x5e, 0xd0, 0xa4, 0xbd, 0x2d, 0x44, 0x79, 0x1d, 0xdc,
	0x96, 0xf5, 0x5f, 0xc1, 0xee, 0x26, 0x49, 0x6c, 0xd0, 0x2f, 0xf1, 0xba, 0x76, 0x26, 0x8f, 0xa4,
	0x07, 0xdb, 0xcb, 0x28, 0x5b, 0x60, 0xed, 0xa8, 0x0a, 0x5e, 0xb4, 0x9e, 0x6b, 0xee, 0x37, 0x0d,
	0x76, 0x8f, 0xca, 0x78, 0x46, 0x97, 0x78, 0x77, 0x73, 0x36, 0xe8, 0x25, 0x36, 0x5b, 0x95, 0x47,
	0xf2, 0x14, 0x3a, 0x09, 0x2b, 0xf3, 0x48, 0xa8, 0x9f, 0x62, 0xd7, 0x3f, 0x5c, 0xb9, 0xdf, 0x14,
	0xf5, 0xde, 0xa9, 0xac, 0xa0, 0xce, 0x76, 0x5f, 0x42, 0xa7, 0x42, 0x88, 0x05, 0xc6, 0xb8, 0x58,
	0x46, 0x19, 0x9d, 0xda, 0x5b, 0xc4, 0x84, 0xf6, 0xe9, 0xc5, 0x78, 0x62, 0x6b, 0xc4, 0x00, 0xfd,
	0xfc, 0x28, 0xb0, 0x5b, 0xa4, 0x0b, 0xdb, 0xc7, 0x17, 0xe3, 0x89, 0x6f, 0xeb, 0xa4, 0x03, 0xad,
	0x8f, 0x17, 0x76, 0xdb, 0x1d, 0xc3, 0xde, 0x4a, 0xbd, 0x1e, 0x64, 0x1f, 0xcc, 0x84, 0x66, 0x58,
	0x44, 0x39, 0xd6, 0xa6, 0x57, 0xb1, 0xe4, 0x62, 0x56, 0x08, 0x2c, 0x04, 0x57, 0xd6, 0x77, 0x82,
	0x55, 0xec, 0x7f, 0xd7, 0xc0, 0x38, 0xad, 0x1c, 0x13, 0x1f, 0xf4, 0x33, 0x96, 0x92, 0xff, 0x57,
	0x2d, 0xdc, 0x3e, 0x8d, 0x7e, 0x6f, 0x13, 0xac, 0x6e, 0x75, 0xb7, 0xc8, 0x13, 0x68, 0xcb, 0x35,
	0x91, 0xde, 0x2f, 0x5b, 0xab, 0xaa, 0xee, 0xfd, 0x71, 0x97, 0xee, 0x16, 0x79, 0x0d, 0x46, 0xdd,
	0x01, 0xb9, 0xff, 0x97, 0x89, 0xf5, 0x9d, 0xdf, 0x89, 0xa6, 0xfe, 0x53, 0x47, 0xbd, 0xdf, 0xc7,
	0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xee, 0xbb, 0x26, 0x37, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GitilesClient is the client API for Gitiles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GitilesClient interface {
	// Log retrieves commit log.
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	// Refs retrieves repo refs.
	Refs(ctx context.Context, in *RefsRequest, opts ...grpc.CallOption) (*RefsResponse, error)
	// Archive retrieves archived contents of the project.
	//
	// An archive is a shallow bundle of the contents of a repository.
	Archive(ctx context.Context, in *ArchiveRequest, opts ...grpc.CallOption) (*ArchiveResponse, error)
}
type gitilesPRPCClient struct {
	client *prpc.Client
}

func NewGitilesPRPCClient(client *prpc.Client) GitilesClient {
	return &gitilesPRPCClient{client}
}

func (c *gitilesPRPCClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.client.Call(ctx, "gitiles.Gitiles", "Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitilesPRPCClient) Refs(ctx context.Context, in *RefsRequest, opts ...grpc.CallOption) (*RefsResponse, error) {
	out := new(RefsResponse)
	err := c.client.Call(ctx, "gitiles.Gitiles", "Refs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitilesPRPCClient) Archive(ctx context.Context, in *ArchiveRequest, opts ...grpc.CallOption) (*ArchiveResponse, error) {
	out := new(ArchiveResponse)
	err := c.client.Call(ctx, "gitiles.Gitiles", "Archive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type gitilesClient struct {
	cc *grpc.ClientConn
}

func NewGitilesClient(cc *grpc.ClientConn) GitilesClient {
	return &gitilesClient{cc}
}

func (c *gitilesClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/gitiles.Gitiles/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitilesClient) Refs(ctx context.Context, in *RefsRequest, opts ...grpc.CallOption) (*RefsResponse, error) {
	out := new(RefsResponse)
	err := c.cc.Invoke(ctx, "/gitiles.Gitiles/Refs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitilesClient) Archive(ctx context.Context, in *ArchiveRequest, opts ...grpc.CallOption) (*ArchiveResponse, error) {
	out := new(ArchiveResponse)
	err := c.cc.Invoke(ctx, "/gitiles.Gitiles/Archive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitilesServer is the server API for Gitiles service.
type GitilesServer interface {
	// Log retrieves commit log.
	Log(context.Context, *LogRequest) (*LogResponse, error)
	// Refs retrieves repo refs.
	Refs(context.Context, *RefsRequest) (*RefsResponse, error)
	// Archive retrieves archived contents of the project.
	//
	// An archive is a shallow bundle of the contents of a repository.
	Archive(context.Context, *ArchiveRequest) (*ArchiveResponse, error)
}

func RegisterGitilesServer(s prpc.Registrar, srv GitilesServer) {
	s.RegisterService(&_Gitiles_serviceDesc, srv)
}

func _Gitiles_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitilesServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitiles.Gitiles/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitilesServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gitiles_Refs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitilesServer).Refs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitiles.Gitiles/Refs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitilesServer).Refs(ctx, req.(*RefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gitiles_Archive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitilesServer).Archive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitiles.Gitiles/Archive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitilesServer).Archive(ctx, req.(*ArchiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gitiles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitiles.Gitiles",
	HandlerType: (*GitilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _Gitiles_Log_Handler,
		},
		{
			MethodName: "Refs",
			Handler:    _Gitiles_Refs_Handler,
		},
		{
			MethodName: "Archive",
			Handler:    _Gitiles_Archive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/common/proto/gitiles/gitiles.proto",
}
