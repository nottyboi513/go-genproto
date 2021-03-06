// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/search_term_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A search term view with metrics aggregated by search term at the ad group
// level.
type SearchTermView struct {
	// Output only. The resource name of the search term view.
	// Search term view resource names have the form:
	//
	// `customers/{customer_id}/searchTermViews/{campaign_id}~{ad_group_id}~{URL-base64_search_term}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The search term.
	SearchTerm *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	// Output only. The ad group the search term served in.
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Output only. Indicates whether the search term is currently one of your
	// targeted or excluded keywords.
	Status               enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_unrecognized     []byte                                                        `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *SearchTermView) Reset()         { *m = SearchTermView{} }
func (m *SearchTermView) String() string { return proto.CompactTextString(m) }
func (*SearchTermView) ProtoMessage()    {}
func (*SearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_c10e074d81af5129, []int{0}
}

func (m *SearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermView.Unmarshal(m, b)
}
func (m *SearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermView.Marshal(b, m, deterministic)
}
func (m *SearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermView.Merge(m, src)
}
func (m *SearchTermView) XXX_Size() int {
	return xxx_messageInfo_SearchTermView.Size(m)
}
func (m *SearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermView proto.InternalMessageInfo

func (m *SearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func (m *SearchTermView) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *SearchTermView) GetStatus() enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus {
	if m != nil {
		return m.Status
	}
	return enums.SearchTermTargetingStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*SearchTermView)(nil), "google.ads.googleads.v2.resources.SearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/search_term_view.proto", fileDescriptor_c10e074d81af5129)
}

var fileDescriptor_c10e074d81af5129 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0x89, 0x3d, 0xba, 0x4d, 0xdd, 0x7a, 0xf0, 0x2e, 0x59, 0x29, 0x5b, 0x5a, 0x28, 0xcb,
	0x2e, 0x12, 0x78, 0x87, 0x0d, 0xef, 0x32, 0x19, 0x46, 0x60, 0x87, 0x51, 0x92, 0xe0, 0xc3, 0x08,
	0x18, 0xc5, 0x7e, 0x55, 0x0d, 0xb1, 0x64, 0x24, 0x39, 0x39, 0x94, 0xc2, 0x3e, 0xcb, 0x8e, 0xfb,
	0x28, 0xfb, 0x14, 0xbd, 0xae, 0x1f, 0x61, 0xa7, 0x11, 0xcb, 0x72, 0x92, 0x41, 0xd6, 0xde, 0xfe,
	0xe1, 0xfd, 0xdf, 0xef, 0xfd, 0xfd, 0x9e, 0x82, 0x3e, 0x70, 0x29, 0xf9, 0x02, 0x08, 0xcb, 0x35,
	0xb1, 0x72, 0xad, 0x96, 0x21, 0x51, 0xa0, 0x65, 0xad, 0x32, 0xd0, 0x44, 0x03, 0x53, 0xd9, 0x55,
	0x6a, 0x40, 0x95, 0xe9, 0xb2, 0x80, 0x15, 0xae, 0x94, 0x34, 0x32, 0x38, 0xb5, 0x76, 0xcc, 0x72,
	0x8d, 0xbb, 0x4e, 0xbc, 0x0c, 0x71, 0xd7, 0x79, 0xfc, 0x69, 0x1f, 0x1c, 0x44, 0x5d, 0xee, 0x82,
	0x0d, 0x53, 0x1c, 0x4c, 0x21, 0x78, 0xaa, 0x0d, 0x33, 0xb5, 0xb6, 0x43, 0x8e, 0x5f, 0x3b, 0x42,
	0x55, 0x90, 0xcb, 0x02, 0x16, 0x79, 0x3a, 0x87, 0x2b, 0xb6, 0x2c, 0xa4, 0x6a, 0x0d, 0x2f, 0xb7,
	0x0c, 0x6e, 0x70, 0x5b, 0x7a, 0xd5, 0x96, 0x9a, 0x5f, 0xf3, 0xfa, 0x92, 0xac, 0x14, 0xab, 0x2a,
	0x50, 0x8e, 0x7d, 0xb2, 0xd5, 0xca, 0x84, 0x90, 0x86, 0x99, 0x42, 0x8a, 0xb6, 0x7a, 0xf6, 0xdb,
	0x47, 0x47, 0x93, 0x26, 0xe0, 0x14, 0x54, 0x99, 0x14, 0xb0, 0x0a, 0xa6, 0xe8, 0xb9, 0x1b, 0x91,
	0x0a, 0x56, 0x42, 0xbf, 0x37, 0xe8, 0x0d, 0x9f, 0xc6, 0xe4, 0x96, 0xfa, 0x7f, 0xe8, 0x5b, 0xf4,
	0x66, 0xb3, 0x85, 0x56, 0x55, 0x85, 0xc6, 0x99, 0x2c, 0xc9, 0x2e, 0x67, 0xfc, 0xcc, 0x51, 0xbe,
	0xb2, 0x12, 0x82, 0x18, 0x1d, 0x6e, 0x2d, 0xa2, 0xef, 0x0d, 0x7a, 0xc3, 0xc3, 0xf0, 0xa4, 0x45,
	0x60, 0x17, 0x1e, 0x4f, 0x8c, 0x2a, 0x04, 0x4f, 0xd8, 0xa2, 0x86, 0xd8, 0xbf, 0xa5, 0xfe, 0x18,
	0xe9, 0x8e, 0x1a, 0x30, 0xf4, 0x84, 0xe5, 0x29, 0x57, 0xb2, 0xae, 0xfa, 0xfe, 0x03, 0x00, 0xc3,
	0x26, 0xf2, 0x19, 0x1a, 0xec, 0x8d, 0x4c, 0xf3, 0xd1, 0x9a, 0x36, 0x7e, 0xcc, 0xac, 0x08, 0x34,
	0x3a, 0xb0, 0x97, 0xe9, 0x3f, 0x1a, 0xf4, 0x86, 0x47, 0x61, 0x82, 0xf7, 0xdd, 0xbf, 0x39, 0x2e,
	0xde, 0x7c, 0xf3, 0xd4, 0x9d, 0x76, 0xd2, 0xf4, 0x7f, 0x16, 0x75, 0xb9, 0xbf, 0x6a, 0xbf, 0xad,
	0x1d, 0x15, 0xc1, 0x1d, 0x9d, 0x3f, 0x78, 0xaf, 0xc1, 0xfb, 0xac, 0xd6, 0x46, 0x96, 0xa0, 0x34,
	0xb9, 0x76, 0xf2, 0x86, 0xe8, 0x1d, 0x93, 0x26, 0xd7, 0xff, 0xbe, 0xe7, 0x9b, 0xf8, 0xbb, 0x87,
	0xce, 0x33, 0x59, 0xe2, 0x7b, 0x5f, 0x74, 0xfc, 0x62, 0x77, 0xe4, 0xc5, 0x7a, 0xb9, 0x17, 0xbd,
	0x6f, 0x5f, 0xda, 0x4e, 0x2e, 0x17, 0x4c, 0x70, 0x2c, 0x15, 0x27, 0x1c, 0x44, 0xb3, 0x7a, 0xb2,
	0xc9, 0xfc, 0x9f, 0x3f, 0xd9, 0xc7, 0x4e, 0xfd, 0xf0, 0xfc, 0x11, 0xa5, 0x3f, 0xbd, 0xd3, 0x91,
	0x45, 0xd2, 0x5c, 0x63, 0x2b, 0xd7, 0x2a, 0x09, 0xf1, 0xd8, 0x39, 0x7f, 0x39, 0xcf, 0x8c, 0xe6,
	0x7a, 0xd6, 0x79, 0x66, 0x49, 0x38, 0xeb, 0x3c, 0x77, 0xde, 0xb9, 0x2d, 0x44, 0x11, 0xcd, 0x75,
	0x14, 0x75, 0xae, 0x28, 0x4a, 0xc2, 0x28, 0xea, 0x7c, 0xf3, 0x83, 0x26, 0xec, 0xbb, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf0, 0x83, 0x88, 0x3d, 0x10, 0x04, 0x00, 0x00,
}
