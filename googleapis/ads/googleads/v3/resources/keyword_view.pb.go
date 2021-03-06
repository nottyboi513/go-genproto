// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/keyword_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// A keyword view.
type KeywordView struct {
	// Output only. The resource name of the keyword view.
	// Keyword view resource names have the form:
	//
	// `customers/{customer_id}/keywordViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordView) Reset()         { *m = KeywordView{} }
func (m *KeywordView) String() string { return proto.CompactTextString(m) }
func (*KeywordView) ProtoMessage()    {}
func (*KeywordView) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b06c74ce279f21, []int{0}
}

func (m *KeywordView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordView.Unmarshal(m, b)
}
func (m *KeywordView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordView.Marshal(b, m, deterministic)
}
func (m *KeywordView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordView.Merge(m, src)
}
func (m *KeywordView) XXX_Size() int {
	return xxx_messageInfo_KeywordView.Size(m)
}
func (m *KeywordView) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordView.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordView proto.InternalMessageInfo

func (m *KeywordView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*KeywordView)(nil), "google.ads.googleads.v3.resources.KeywordView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/keyword_view.proto", fileDescriptor_b4b06c74ce279f21)
}

var fileDescriptor_b4b06c74ce279f21 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4a, 0xf3, 0x40,
	0x1c, 0xc5, 0x49, 0x0b, 0x1f, 0x7c, 0x51, 0x41, 0xba, 0xd2, 0x22, 0x68, 0xc5, 0x8a, 0x0b, 0x99,
	0x11, 0xe2, 0x6a, 0x5c, 0x4d, 0x37, 0x05, 0x05, 0xa9, 0x5d, 0x04, 0xd1, 0x40, 0x99, 0x66, 0xc6,
	0x38, 0xd8, 0xe4, 0x5f, 0x66, 0xd2, 0x14, 0x29, 0xbd, 0x8c, 0x4b, 0xef, 0xe0, 0x05, 0x3c, 0x85,
	0xeb, 0x1e, 0x41, 0x10, 0x24, 0x9d, 0xce, 0x34, 0x2b, 0x75, 0xf7, 0xe0, 0xff, 0x7b, 0x2f, 0x2f,
	0x6f, 0xfc, 0xf3, 0x04, 0x20, 0x19, 0x09, 0xcc, 0xb8, 0xc6, 0x46, 0x96, 0xaa, 0x08, 0xb0, 0x12,
	0x1a, 0x26, 0x2a, 0x16, 0x1a, 0x3f, 0x89, 0xe7, 0x29, 0x28, 0x3e, 0x28, 0xa4, 0x98, 0xa2, 0xb1,
	0x82, 0x1c, 0x1a, 0x2d, 0x83, 0x22, 0xc6, 0x35, 0x72, 0x2e, 0x54, 0x04, 0xc8, 0xb9, 0x9a, 0xfb,
	0x36, 0x78, 0x2c, 0xf1, 0x83, 0x14, 0x23, 0x3e, 0x18, 0x8a, 0x47, 0x56, 0x48, 0x50, 0x26, 0xa3,
	0xb9, 0x5b, 0x01, 0xac, 0x6d, 0x75, 0xda, 0xab, 0x9c, 0x58, 0x96, 0x41, 0xce, 0x72, 0x09, 0x99,
	0x36, 0xd7, 0xc3, 0x37, 0xcf, 0xdf, 0xb8, 0x32, 0x9d, 0x42, 0x29, 0xa6, 0x8d, 0x1b, 0x7f, 0xcb,
	0xfa, 0x07, 0x19, 0x4b, 0xc5, 0x8e, 0x77, 0xe0, 0x9d, 0xfc, 0xef, 0x9c, 0x7e, 0xd0, 0xfa, 0x27,
	0x3d, 0xf6, 0x8f, 0xd6, 0x05, 0x57, 0x6a, 0x2c, 0x35, 0x8a, 0x21, 0xc5, 0x95, 0x90, 0xfe, 0xa6,
	0x8d, 0xb8, 0x66, 0xa9, 0x20, 0xf7, 0x0b, 0x7a, 0xfb, 0x37, 0x63, 0xe3, 0x2c, 0x9e, 0xe8, 0x1c,
	0x52, 0xa1, 0x34, 0x9e, 0x59, 0x39, 0xb7, 0x9b, 0x95, 0x84, 0xc6, 0xb3, 0xea, 0x82, 0xf3, 0xce,
	0x97, 0xe7, 0xb7, 0x63, 0x48, 0xd1, 0xaf, 0x1b, 0x76, 0xb6, 0x2b, 0x1f, 0xea, 0x95, 0xff, 0xde,
	0xf3, 0xee, 0x2e, 0x57, 0xb6, 0x04, 0x46, 0x2c, 0x4b, 0x10, 0xa8, 0x04, 0x27, 0x22, 0x5b, 0x2e,
	0x83, 0xd7, 0x35, 0x7f, 0x78, 0xcf, 0x0b, 0xa7, 0x5e, 0x6a, 0xf5, 0x2e, 0xa5, 0xaf, 0xb5, 0x56,
	0xd7, 0x44, 0x52, 0xae, 0x91, 0x91, 0xa5, 0x0a, 0x03, 0xd4, 0xb7, 0xe4, 0xbb, 0x65, 0x22, 0xca,
	0x75, 0xe4, 0x98, 0x28, 0x0c, 0x22, 0xc7, 0x2c, 0x6a, 0x6d, 0x73, 0x20, 0x84, 0x72, 0x4d, 0x88,
	0xa3, 0x08, 0x09, 0x03, 0x42, 0x1c, 0x37, 0xfc, 0xb7, 0x2c, 0x1b, 0x7c, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xd0, 0x7f, 0xf2, 0x77, 0x7b, 0x02, 0x00, 0x00,
}
