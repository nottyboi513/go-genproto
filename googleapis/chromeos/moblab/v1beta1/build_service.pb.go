// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/chromeos/moblab/v1beta1/build_service.proto

package moblab

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for listing builds.
// NEXT_TAG: 6
type ListBuildsRequest struct {
	// Required. The full resource name of the model. The model id is the same as
	// the build target id for non-unified builds.
	// For example,
	// 'buildTargets/octopus/models/bobba'.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The number of builds to return in a page.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A page token, received from a previous `ListBuilds` call. Provide this to
	// retrieve the subsequent page.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Filter that specifies value constraints of fields. For example, the
	// filter can be set as "filter='milestone=milestones/80'" to only select
	// builds in milestone 80.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Read mask that specifies which Build fields to return. If empty, all Build
	// fields will be returned.
	// Valid fields: name, milestone, build_version.
	// For example, if the read_mask is set as "read_mask='milestone'", the
	// ListBuilds will return a list of Builds object with only the milestone
	// field.
	ReadMask             *field_mask.FieldMask `protobuf:"bytes,5,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListBuildsRequest) Reset()         { *m = ListBuildsRequest{} }
func (m *ListBuildsRequest) String() string { return proto.CompactTextString(m) }
func (*ListBuildsRequest) ProtoMessage()    {}
func (*ListBuildsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f08d110278057f, []int{0}
}

func (m *ListBuildsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBuildsRequest.Unmarshal(m, b)
}
func (m *ListBuildsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBuildsRequest.Marshal(b, m, deterministic)
}
func (m *ListBuildsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBuildsRequest.Merge(m, src)
}
func (m *ListBuildsRequest) XXX_Size() int {
	return xxx_messageInfo_ListBuildsRequest.Size(m)
}
func (m *ListBuildsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBuildsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBuildsRequest proto.InternalMessageInfo

func (m *ListBuildsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListBuildsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListBuildsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListBuildsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListBuildsRequest) GetReadMask() *field_mask.FieldMask {
	if m != nil {
		return m.ReadMask
	}
	return nil
}

// Response message for listing builds.
// NEXT_TAG: 4
type ListBuildsResponse struct {
	// The list of builds.
	Builds []*Build `protobuf:"bytes,1,rep,name=builds,proto3" json:"builds,omitempty"`
	// Token to retrieve the next page of builds. If this field is omitted, there
	// are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Total number of builds.
	TotalSize            int32    `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBuildsResponse) Reset()         { *m = ListBuildsResponse{} }
func (m *ListBuildsResponse) String() string { return proto.CompactTextString(m) }
func (*ListBuildsResponse) ProtoMessage()    {}
func (*ListBuildsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f08d110278057f, []int{1}
}

func (m *ListBuildsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBuildsResponse.Unmarshal(m, b)
}
func (m *ListBuildsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBuildsResponse.Marshal(b, m, deterministic)
}
func (m *ListBuildsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBuildsResponse.Merge(m, src)
}
func (m *ListBuildsResponse) XXX_Size() int {
	return xxx_messageInfo_ListBuildsResponse.Size(m)
}
func (m *ListBuildsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBuildsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBuildsResponse proto.InternalMessageInfo

func (m *ListBuildsResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

func (m *ListBuildsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListBuildsResponse) GetTotalSize() int32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

// Request message for checking if the build artifact is staged.
type CheckBuildStageStatusRequest struct {
	// Required. The full resource name of the build artifact.
	// For example,
	// 'buildTargets/octopus/models/bobba/builds/12607.6.0/artifacts/chromeos-moblab-peng-staging'.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckBuildStageStatusRequest) Reset()         { *m = CheckBuildStageStatusRequest{} }
func (m *CheckBuildStageStatusRequest) String() string { return proto.CompactTextString(m) }
func (*CheckBuildStageStatusRequest) ProtoMessage()    {}
func (*CheckBuildStageStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f08d110278057f, []int{2}
}

func (m *CheckBuildStageStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckBuildStageStatusRequest.Unmarshal(m, b)
}
func (m *CheckBuildStageStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckBuildStageStatusRequest.Marshal(b, m, deterministic)
}
func (m *CheckBuildStageStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckBuildStageStatusRequest.Merge(m, src)
}
func (m *CheckBuildStageStatusRequest) XXX_Size() int {
	return xxx_messageInfo_CheckBuildStageStatusRequest.Size(m)
}
func (m *CheckBuildStageStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckBuildStageStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckBuildStageStatusRequest proto.InternalMessageInfo

func (m *CheckBuildStageStatusRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response message for checking the stage status of a build artifact.
// NEXT_TAG: 4
type CheckBuildStageStatusResponse struct {
	// The status to represent if the build is staged or not.
	IsBuildStaged bool `protobuf:"varint,1,opt,name=is_build_staged,json=isBuildStaged,proto3" json:"is_build_staged,omitempty"`
	// The staged build artifact in the destination bucket.
	StagedBuildArtifact *BuildArtifact `protobuf:"bytes,2,opt,name=staged_build_artifact,json=stagedBuildArtifact,proto3" json:"staged_build_artifact,omitempty"`
	// The source build artifact in the source bucket.
	SourceBuildArtifact  *BuildArtifact `protobuf:"bytes,3,opt,name=source_build_artifact,json=sourceBuildArtifact,proto3" json:"source_build_artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CheckBuildStageStatusResponse) Reset()         { *m = CheckBuildStageStatusResponse{} }
func (m *CheckBuildStageStatusResponse) String() string { return proto.CompactTextString(m) }
func (*CheckBuildStageStatusResponse) ProtoMessage()    {}
func (*CheckBuildStageStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f08d110278057f, []int{3}
}

func (m *CheckBuildStageStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckBuildStageStatusResponse.Unmarshal(m, b)
}
func (m *CheckBuildStageStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckBuildStageStatusResponse.Marshal(b, m, deterministic)
}
func (m *CheckBuildStageStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckBuildStageStatusResponse.Merge(m, src)
}
func (m *CheckBuildStageStatusResponse) XXX_Size() int {
	return xxx_messageInfo_CheckBuildStageStatusResponse.Size(m)
}
func (m *CheckBuildStageStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckBuildStageStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckBuildStageStatusResponse proto.InternalMessageInfo

func (m *CheckBuildStageStatusResponse) GetIsBuildStaged() bool {
	if m != nil {
		return m.IsBuildStaged
	}
	return false
}

func (m *CheckBuildStageStatusResponse) GetStagedBuildArtifact() *BuildArtifact {
	if m != nil {
		return m.StagedBuildArtifact
	}
	return nil
}

func (m *CheckBuildStageStatusResponse) GetSourceBuildArtifact() *BuildArtifact {
	if m != nil {
		return m.SourceBuildArtifact
	}
	return nil
}

// Request message for staging a build artifact.
type StageBuildRequest struct {
	// Required. The full resource name of the build artifact.
	// For example,
	// 'buildTargets/octopus/models/bobba/builds/12607.6.0/artifacts/chromeos-moblab-peng-staging'.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StageBuildRequest) Reset()         { *m = StageBuildRequest{} }
func (m *StageBuildRequest) String() string { return proto.CompactTextString(m) }
func (*StageBuildRequest) ProtoMessage()    {}
func (*StageBuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f08d110278057f, []int{4}
}

func (m *StageBuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StageBuildRequest.Unmarshal(m, b)
}
func (m *StageBuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StageBuildRequest.Marshal(b, m, deterministic)
}
func (m *StageBuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StageBuildRequest.Merge(m, src)
}
func (m *StageBuildRequest) XXX_Size() int {
	return xxx_messageInfo_StageBuildRequest.Size(m)
}
func (m *StageBuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StageBuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StageBuildRequest proto.InternalMessageInfo

func (m *StageBuildRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response message for staging a build artifact.
type StageBuildResponse struct {
	// The staged build in the destination bucket.
	StagedBuildArtifact  *BuildArtifact `protobuf:"bytes,1,opt,name=staged_build_artifact,json=stagedBuildArtifact,proto3" json:"staged_build_artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StageBuildResponse) Reset()         { *m = StageBuildResponse{} }
func (m *StageBuildResponse) String() string { return proto.CompactTextString(m) }
func (*StageBuildResponse) ProtoMessage()    {}
func (*StageBuildResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f08d110278057f, []int{5}
}

func (m *StageBuildResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StageBuildResponse.Unmarshal(m, b)
}
func (m *StageBuildResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StageBuildResponse.Marshal(b, m, deterministic)
}
func (m *StageBuildResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StageBuildResponse.Merge(m, src)
}
func (m *StageBuildResponse) XXX_Size() int {
	return xxx_messageInfo_StageBuildResponse.Size(m)
}
func (m *StageBuildResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StageBuildResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StageBuildResponse proto.InternalMessageInfo

func (m *StageBuildResponse) GetStagedBuildArtifact() *BuildArtifact {
	if m != nil {
		return m.StagedBuildArtifact
	}
	return nil
}

// Metadata message for staging a build artifact.
// NEXT_TAG: 4
type StageBuildMetadata struct {
	// Approximate percentage of progress, e.g. "50" means 50%.
	ProgressPercent float32 `protobuf:"fixed32,1,opt,name=progress_percent,json=progressPercent,proto3" json:"progress_percent,omitempty"`
	// Build stage start time.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Build stage end time.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StageBuildMetadata) Reset()         { *m = StageBuildMetadata{} }
func (m *StageBuildMetadata) String() string { return proto.CompactTextString(m) }
func (*StageBuildMetadata) ProtoMessage()    {}
func (*StageBuildMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_14f08d110278057f, []int{6}
}

func (m *StageBuildMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StageBuildMetadata.Unmarshal(m, b)
}
func (m *StageBuildMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StageBuildMetadata.Marshal(b, m, deterministic)
}
func (m *StageBuildMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StageBuildMetadata.Merge(m, src)
}
func (m *StageBuildMetadata) XXX_Size() int {
	return xxx_messageInfo_StageBuildMetadata.Size(m)
}
func (m *StageBuildMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_StageBuildMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_StageBuildMetadata proto.InternalMessageInfo

func (m *StageBuildMetadata) GetProgressPercent() float32 {
	if m != nil {
		return m.ProgressPercent
	}
	return 0
}

func (m *StageBuildMetadata) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *StageBuildMetadata) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ListBuildsRequest)(nil), "google.chromeos.moblab.v1beta1.ListBuildsRequest")
	proto.RegisterType((*ListBuildsResponse)(nil), "google.chromeos.moblab.v1beta1.ListBuildsResponse")
	proto.RegisterType((*CheckBuildStageStatusRequest)(nil), "google.chromeos.moblab.v1beta1.CheckBuildStageStatusRequest")
	proto.RegisterType((*CheckBuildStageStatusResponse)(nil), "google.chromeos.moblab.v1beta1.CheckBuildStageStatusResponse")
	proto.RegisterType((*StageBuildRequest)(nil), "google.chromeos.moblab.v1beta1.StageBuildRequest")
	proto.RegisterType((*StageBuildResponse)(nil), "google.chromeos.moblab.v1beta1.StageBuildResponse")
	proto.RegisterType((*StageBuildMetadata)(nil), "google.chromeos.moblab.v1beta1.StageBuildMetadata")
}

func init() {
	proto.RegisterFile("google/chromeos/moblab/v1beta1/build_service.proto", fileDescriptor_14f08d110278057f)
}

var fileDescriptor_14f08d110278057f = []byte{
	// 873 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0xd8, 0x4d, 0x88, 0x5f, 0xa8, 0x42, 0x06, 0x55, 0x18, 0xd3, 0x80, 0xb5, 0x55, 0x2b,
	0x93, 0xaa, 0xbb, 0xc4, 0x15, 0x07, 0x4a, 0x73, 0x58, 0xb7, 0xa2, 0x1c, 0x88, 0x08, 0xdb, 0x9c,
	0x10, 0x92, 0x35, 0xde, 0x9d, 0x6c, 0x46, 0xd9, 0x9d, 0x59, 0x66, 0xc6, 0x01, 0x15, 0x21, 0x04,
	0x77, 0x4e, 0xdc, 0xf8, 0x16, 0x1c, 0xe0, 0x43, 0xf4, 0xca, 0xad, 0x17, 0x7a, 0x40, 0xe2, 0x33,
	0xc0, 0x09, 0xcd, 0x9f, 0xb5, 0x37, 0x4d, 0x62, 0x43, 0xdb, 0xe3, 0xbe, 0x79, 0xbf, 0xf7, 0xde,
	0xef, 0xf7, 0xfe, 0xd8, 0x30, 0xcc, 0x85, 0xc8, 0x0b, 0x1a, 0xa5, 0x47, 0x52, 0x94, 0x54, 0xa8,
	0xa8, 0x14, 0x93, 0x82, 0x4c, 0xa2, 0x93, 0x9d, 0x09, 0xd5, 0x64, 0x27, 0x9a, 0x4c, 0x59, 0x91,
	0x8d, 0x15, 0x95, 0x27, 0x2c, 0xa5, 0x61, 0x25, 0x85, 0x16, 0xf8, 0x6d, 0x87, 0x09, 0x6b, 0x4c,
	0xe8, 0x30, 0xa1, 0xc7, 0xf4, 0xae, 0xfa, 0x98, 0xa4, 0x62, 0x11, 0xe1, 0x5c, 0x68, 0xa2, 0x99,
	0xe0, 0xca, 0xa1, 0x7b, 0x6f, 0x34, 0x5e, 0xd3, 0x82, 0x51, 0xae, 0xfd, 0xc3, 0x3b, 0x8d, 0x87,
	0x43, 0x46, 0x8b, 0x6c, 0x3c, 0xa1, 0x47, 0xe4, 0x84, 0x09, 0xe9, 0x1d, 0xde, 0x6c, 0x38, 0x48,
	0xaa, 0xc4, 0x54, 0xd6, 0x25, 0xf5, 0xc2, 0x25, 0x34, 0x6a, 0xf7, 0xba, 0x88, 0x6b, 0xde, 0xbf,
	0x10, 0x3c, 0x97, 0x53, 0xce, 0x19, 0xcf, 0x23, 0x51, 0x51, 0x79, 0xaa, 0xd2, 0xbe, 0x77, 0xb2,
	0x5f, 0x93, 0xe9, 0xa1, 0xaf, 0xaa, 0x24, 0xea, 0xf8, 0x99, 0x92, 0x67, 0x1e, 0x9a, 0x95, 0x54,
	0x69, 0x52, 0x56, 0xce, 0x21, 0xf8, 0x1b, 0xc1, 0xe6, 0x27, 0x4c, 0xe9, 0x91, 0x91, 0x51, 0x25,
	0xf4, 0xcb, 0x29, 0x55, 0x1a, 0xdf, 0x83, 0xd5, 0x8a, 0x48, 0xca, 0x75, 0x17, 0xf5, 0xd1, 0xa0,
	0x33, 0xba, 0xf9, 0x34, 0x6e, 0xfd, 0x13, 0x5f, 0x87, 0x6b, 0x75, 0xf9, 0x5e, 0x50, 0x17, 0x9d,
	0x54, 0x4c, 0x85, 0xa9, 0x28, 0xa3, 0x3d, 0x91, 0xd1, 0x22, 0xf1, 0x50, 0xdc, 0x87, 0x4e, 0x45,
	0x72, 0x3a, 0x56, 0xec, 0x11, 0xed, 0xb6, 0xfa, 0x68, 0xb0, 0x32, 0x6a, 0x3f, 0x8d, 0x51, 0xb2,
	0x66, 0xac, 0x0f, 0xd9, 0x23, 0x8a, 0x03, 0x00, 0xeb, 0xa1, 0xc5, 0x31, 0xe5, 0xdd, 0xb6, 0x4d,
	0x65, 0x5d, 0x2c, 0xf0, 0xc0, 0x58, 0xf1, 0x5b, 0xb0, 0x7a, 0xc8, 0x0a, 0x4d, 0x65, 0xf7, 0xd2,
	0xfc, 0xdd, 0x9b, 0xf0, 0x5d, 0xe8, 0x48, 0x4a, 0x1c, 0xe3, 0xee, 0x4a, 0x1f, 0x0d, 0xd6, 0x87,
	0xb5, 0xd2, 0x61, 0x4d, 0x39, 0xfc, 0xc8, 0x88, 0xb2, 0x47, 0xd4, 0xb1, 0x4f, 0x6f, 0x10, 0xe6,
	0x33, 0xf8, 0x19, 0x01, 0x6e, 0x72, 0x57, 0x95, 0xe0, 0x8a, 0xe2, 0x5d, 0x58, 0xb5, 0x43, 0xa5,
	0xba, 0xa8, 0xdf, 0x1e, 0xac, 0x0f, 0xaf, 0x87, 0x8b, 0xc7, 0x29, 0xb4, 0xf8, 0xc4, 0x83, 0xf0,
	0x0d, 0xd8, 0xe0, 0xf4, 0x6b, 0x3d, 0x6e, 0x30, 0x33, 0xe4, 0x3b, 0xc9, 0x65, 0x63, 0xde, 0x9f,
	0x11, 0xdb, 0x02, 0xd0, 0x42, 0x93, 0xc2, 0xe9, 0x63, 0xc8, 0xaf, 0x24, 0x1d, 0x6b, 0x31, 0xda,
	0x04, 0x39, 0x5c, 0xbd, 0x77, 0x44, 0xd3, 0x63, 0x1b, 0xfc, 0xa1, 0x36, 0x92, 0x69, 0xa2, 0xa7,
	0xb3, 0x16, 0x3d, 0x80, 0x4b, 0x9c, 0x94, 0xd4, 0x37, 0xe8, 0xb6, 0x6d, 0xd0, 0x2d, 0xb8, 0xb9,
	0xb8, 0x41, 0x36, 0x58, 0x2c, 0x35, 0x3b, 0x24, 0xa9, 0x4e, 0x6c, 0x80, 0xe0, 0xc7, 0x16, 0x6c,
	0x5d, 0x90, 0xc9, 0x0b, 0x72, 0x03, 0x36, 0x98, 0x1a, 0xfb, 0x45, 0x33, 0xef, 0x99, 0xcd, 0xba,
	0x96, 0x5c, 0x66, 0x6a, 0x0e, 0xca, 0x30, 0x81, 0x2b, 0xee, 0xd9, 0xfb, 0x12, 0x9f, 0xc8, 0xf2,
	0x5f, 0x1f, 0xde, 0xfa, 0x4f, 0x3a, 0xce, 0xaa, 0x7b, 0xdd, 0xc5, 0x3a, 0x65, 0xb4, 0x29, 0xec,
	0x9e, 0x3c, 0x9b, 0xa2, 0xfd, 0x7c, 0x29, 0x6c, 0xac, 0x53, 0xc6, 0xe0, 0x0b, 0xd8, 0xb4, 0x7c,
	0x5c, 0x57, 0x5f, 0xb6, 0xda, 0x5f, 0x01, 0x6e, 0x46, 0xf7, 0x0a, 0x5f, 0xa8, 0x1c, 0x7a, 0x59,
	0xca, 0x05, 0xbf, 0xa0, 0x66, 0xe6, 0x3d, 0xaa, 0x49, 0x46, 0x34, 0xc1, 0xef, 0xc2, 0x6b, 0x95,
	0x14, 0xb9, 0xa4, 0x4a, 0x8d, 0x2b, 0x2a, 0xd3, 0x7a, 0xe7, 0x5b, 0xc9, 0x46, 0x6d, 0xdf, 0x77,
	0x66, 0xfc, 0x01, 0x80, 0xd2, 0x44, 0xea, 0xb1, 0xb9, 0x21, 0xbe, 0xa7, 0x67, 0xb7, 0xed, 0xa0,
	0x3e, 0x30, 0x49, 0xc7, 0x7a, 0x9b, 0x6f, 0xfc, 0x3e, 0xac, 0x51, 0x9e, 0x39, 0x60, 0x7b, 0x29,
	0xf0, 0x15, 0xca, 0x33, 0xf3, 0x35, 0xfc, 0x75, 0x05, 0x5e, 0x75, 0x03, 0xe6, 0xce, 0x3b, 0xfe,
	0x0d, 0x01, 0xcc, 0x37, 0x16, 0xef, 0x2c, 0xd3, 0xe5, 0xcc, 0x65, 0xeb, 0x0d, 0xff, 0x0f, 0xc4,
	0x75, 0x27, 0x88, 0x9f, 0xc4, 0xfe, 0xa6, 0xfd, 0xf0, 0xfb, 0x9f, 0x3f, 0xb5, 0x86, 0xf8, 0xbd,
	0xd9, 0xd9, 0xfe, 0xc6, 0xd9, 0x77, 0x6d, 0xdb, 0x0e, 0x88, 0xcc, 0xa9, 0x56, 0xd1, 0x76, 0x54,
	0x9a, 0x43, 0xa8, 0xa2, 0xed, 0x6f, 0x23, 0x7f, 0x14, 0xfe, 0x42, 0x70, 0xe5, 0xdc, 0x25, 0xc3,
	0x77, 0x97, 0x15, 0xb4, 0xe8, 0x0a, 0xf4, 0x76, 0x9f, 0x13, 0xed, 0x99, 0x7d, 0xf6, 0x24, 0xb6,
	0x63, 0x69, 0x79, 0xdd, 0xc7, 0xa3, 0x39, 0x2f, 0x63, 0xbd, 0x88, 0x95, 0x27, 0x15, 0x6d, 0x47,
	0xf5, 0x9c, 0x1a, 0xaa, 0x77, 0x52, 0x93, 0x0a, 0xff, 0x81, 0x00, 0xe6, 0x73, 0xb6, 0xbc, 0x45,
	0x67, 0x76, 0xad, 0xb7, 0x55, 0x43, 0x1a, 0xbf, 0x7d, 0xe1, 0xa7, 0xf5, 0x6f, 0x5f, 0xf0, 0xdd,
	0xe3, 0x78, 0x70, 0xee, 0x12, 0x9d, 0x33, 0xde, 0x4d, 0x7a, 0x0f, 0x82, 0x17, 0xa3, 0x67, 0x97,
	0xea, 0x0e, 0xda, 0xee, 0xf5, 0x1f, 0xc7, 0x5b, 0x0b, 0x17, 0x7f, 0xf4, 0x3d, 0x82, 0x20, 0x15,
	0xe5, 0x12, 0xea, 0xa3, 0xcd, 0xe6, 0x6c, 0xef, 0x9b, 0x3d, 0xf8, 0x18, 0xed, 0xa3, 0xcf, 0xef,
	0x7b, 0x58, 0x2e, 0x0a, 0xc2, 0xf3, 0x50, 0xc8, 0x3c, 0xca, 0x29, 0xb7, 0x7b, 0x12, 0xcd, 0x13,
	0x5d, 0xf4, 0x4f, 0xe2, 0x43, 0xf7, 0x39, 0x59, 0xb5, 0x80, 0xdb, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0xd0, 0xd8, 0x40, 0x40, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BuildServiceClient is the client API for BuildService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildServiceClient interface {
	// Lists all builds for the given build target and model in descending order
	// for the milestones and build versions.
	ListBuilds(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsResponse, error)
	// Checks the stage status for a given build artifact in a partner Google
	// Cloud Storage bucket.
	CheckBuildStageStatus(ctx context.Context, in *CheckBuildStageStatusRequest, opts ...grpc.CallOption) (*CheckBuildStageStatusResponse, error)
	// Stages a given build artifact from a internal Google Cloud Storage bucket
	// to a partner Google Cloud Storage bucket. If any of objects has already
	// been copied, it will overwrite the previous objects. Operation <response:
	// [StageBuildResponse][google.chromeos.moblab.v1beta1.StageBuildResponse],
	//            metadata: [StageBuildMetadata][google.chromeos.moblab.v1beta1.StageBuildMetadata]>
	StageBuild(ctx context.Context, in *StageBuildRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type buildServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildServiceClient(cc grpc.ClientConnInterface) BuildServiceClient {
	return &buildServiceClient{cc}
}

func (c *buildServiceClient) ListBuilds(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsResponse, error) {
	out := new(ListBuildsResponse)
	err := c.cc.Invoke(ctx, "/google.chromeos.moblab.v1beta1.BuildService/ListBuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) CheckBuildStageStatus(ctx context.Context, in *CheckBuildStageStatusRequest, opts ...grpc.CallOption) (*CheckBuildStageStatusResponse, error) {
	out := new(CheckBuildStageStatusResponse)
	err := c.cc.Invoke(ctx, "/google.chromeos.moblab.v1beta1.BuildService/CheckBuildStageStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) StageBuild(ctx context.Context, in *StageBuildRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.chromeos.moblab.v1beta1.BuildService/StageBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildServiceServer is the server API for BuildService service.
type BuildServiceServer interface {
	// Lists all builds for the given build target and model in descending order
	// for the milestones and build versions.
	ListBuilds(context.Context, *ListBuildsRequest) (*ListBuildsResponse, error)
	// Checks the stage status for a given build artifact in a partner Google
	// Cloud Storage bucket.
	CheckBuildStageStatus(context.Context, *CheckBuildStageStatusRequest) (*CheckBuildStageStatusResponse, error)
	// Stages a given build artifact from a internal Google Cloud Storage bucket
	// to a partner Google Cloud Storage bucket. If any of objects has already
	// been copied, it will overwrite the previous objects. Operation <response:
	// [StageBuildResponse][google.chromeos.moblab.v1beta1.StageBuildResponse],
	//            metadata: [StageBuildMetadata][google.chromeos.moblab.v1beta1.StageBuildMetadata]>
	StageBuild(context.Context, *StageBuildRequest) (*longrunning.Operation, error)
}

// UnimplementedBuildServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBuildServiceServer struct {
}

func (*UnimplementedBuildServiceServer) ListBuilds(ctx context.Context, req *ListBuildsRequest) (*ListBuildsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuilds not implemented")
}
func (*UnimplementedBuildServiceServer) CheckBuildStageStatus(ctx context.Context, req *CheckBuildStageStatusRequest) (*CheckBuildStageStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBuildStageStatus not implemented")
}
func (*UnimplementedBuildServiceServer) StageBuild(ctx context.Context, req *StageBuildRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StageBuild not implemented")
}

func RegisterBuildServiceServer(s *grpc.Server, srv BuildServiceServer) {
	s.RegisterService(&_BuildService_serviceDesc, srv)
}

func _BuildService_ListBuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).ListBuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.chromeos.moblab.v1beta1.BuildService/ListBuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).ListBuilds(ctx, req.(*ListBuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_CheckBuildStageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckBuildStageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).CheckBuildStageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.chromeos.moblab.v1beta1.BuildService/CheckBuildStageStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).CheckBuildStageStatus(ctx, req.(*CheckBuildStageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuildService_StageBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StageBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).StageBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.chromeos.moblab.v1beta1.BuildService/StageBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).StageBuild(ctx, req.(*StageBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuildService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.chromeos.moblab.v1beta1.BuildService",
	HandlerType: (*BuildServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBuilds",
			Handler:    _BuildService_ListBuilds_Handler,
		},
		{
			MethodName: "CheckBuildStageStatus",
			Handler:    _BuildService_CheckBuildStageStatus_Handler,
		},
		{
			MethodName: "StageBuild",
			Handler:    _BuildService_StageBuild_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/chromeos/moblab/v1beta1/build_service.proto",
}
