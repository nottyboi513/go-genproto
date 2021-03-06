// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_client.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A link between the given customer and a client customer. CustomerClients only
// exist for manager customers. All direct and indirect client customers are
// included, as well as the manager itself.
type CustomerClient struct {
	// Output only. The resource name of the customer client.
	// CustomerClient resource names have the form:
	// `customers/{customer_id}/customerClients/{client_customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The resource name of the client-customer which is linked to
	// the given customer. Read only.
	ClientCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=client_customer,json=clientCustomer,proto3" json:"client_customer,omitempty"`
	// Output only. Specifies whether this is a
	// [hidden account](https://support.google.com/google-ads/answer/7519830).
	// Read only.
	Hidden *wrappers.BoolValue `protobuf:"bytes,4,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// Output only. Distance between given customer and client. For self link, the level value
	// will be 0. Read only.
	Level                *wrappers.Int64Value `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CustomerClient) Reset()         { *m = CustomerClient{} }
func (m *CustomerClient) String() string { return proto.CompactTextString(m) }
func (*CustomerClient) ProtoMessage()    {}
func (*CustomerClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1fcc371bcda99b7, []int{0}
}

func (m *CustomerClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClient.Unmarshal(m, b)
}
func (m *CustomerClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClient.Marshal(b, m, deterministic)
}
func (m *CustomerClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClient.Merge(m, src)
}
func (m *CustomerClient) XXX_Size() int {
	return xxx_messageInfo_CustomerClient.Size(m)
}
func (m *CustomerClient) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClient.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClient proto.InternalMessageInfo

func (m *CustomerClient) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerClient) GetClientCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ClientCustomer
	}
	return nil
}

func (m *CustomerClient) GetHidden() *wrappers.BoolValue {
	if m != nil {
		return m.Hidden
	}
	return nil
}

func (m *CustomerClient) GetLevel() *wrappers.Int64Value {
	if m != nil {
		return m.Level
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomerClient)(nil), "google.ads.googleads.v1.resources.CustomerClient")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_client.proto", fileDescriptor_a1fcc371bcda99b7)
}

var fileDescriptor_a1fcc371bcda99b7 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0xd9, 0x1d, 0x5b, 0x30, 0x6a, 0x85, 0xf1, 0xb2, 0xae, 0x45, 0x5b, 0xa1, 0x58, 0x2f,
	0x09, 0xa3, 0xb2, 0x85, 0x78, 0xca, 0xf4, 0x50, 0xec, 0x41, 0xca, 0x2a, 0x7b, 0x90, 0x85, 0x25,
	0x3b, 0x49, 0xa7, 0x81, 0x4c, 0x32, 0x24, 0xb3, 0xe3, 0x41, 0x0a, 0x5e, 0xfd, 0x1a, 0x1e, 0xfd,
	0x28, 0x7e, 0x8a, 0x9e, 0xfb, 0x11, 0x3c, 0xc9, 0x4e, 0xfe, 0x74, 0xca, 0x82, 0xf6, 0xf6, 0x0c,
	0xcf, 0xef, 0x79, 0xe6, 0x7d, 0xc3, 0x0b, 0x8e, 0x4a, 0xad, 0x4b, 0xc9, 0x11, 0x65, 0x16, 0x39,
	0xb9, 0x56, 0x6d, 0x86, 0x0c, 0xb7, 0x7a, 0x65, 0x0a, 0x6e, 0x51, 0xb1, 0xb2, 0x8d, 0xae, 0xb8,
	0x59, 0x14, 0x52, 0x70, 0xd5, 0xc0, 0xda, 0xe8, 0x46, 0xa7, 0xfb, 0x8e, 0x86, 0x94, 0x59, 0x18,
	0x83, 0xb0, 0xcd, 0x60, 0x0c, 0x8e, 0x5f, 0x84, 0xee, 0x5a, 0xa0, 0x73, 0xc1, 0x25, 0x5b, 0x2c,
	0xf9, 0x05, 0x6d, 0x85, 0x36, 0xae, 0x63, 0xfc, 0xb4, 0x07, 0x84, 0x98, 0xb7, 0x9e, 0x7b, 0xab,
	0xfb, 0x5a, 0xae, 0xce, 0xd1, 0x57, 0x43, 0xeb, 0x9a, 0x1b, 0xeb, 0xfd, 0xdd, 0x5e, 0x94, 0x2a,
	0xa5, 0x1b, 0xda, 0x08, 0xad, 0xbc, 0xfb, 0xf2, 0x47, 0x02, 0x76, 0x8e, 0xfd, 0xd8, 0xc7, 0xdd,
	0xd4, 0xe9, 0x67, 0xf0, 0x28, 0xfc, 0x62, 0xa1, 0x68, 0xc5, 0x47, 0x83, 0xbd, 0xc1, 0xe1, 0xfd,
	0x1c, 0x5d, 0x91, 0xe4, 0x0f, 0x79, 0x0d, 0x5e, 0xdd, 0xec, 0xe0, 0x55, 0x2d, 0x2c, 0x2c, 0x74,
	0x85, 0x6e, 0xf7, 0x4c, 0x1f, 0x86, 0x96, 0x8f, 0xb4, 0xe2, 0xe9, 0x29, 0x78, 0xec, 0x5e, 0x65,
	0x11, 0x5e, 0x69, 0x94, 0xec, 0x0d, 0x0e, 0x1f, 0xbc, 0xd9, 0xf5, 0x35, 0x30, 0x2c, 0x00, 0x3f,
	0x35, 0x46, 0xa8, 0x72, 0x46, 0xe5, 0x8a, 0xe7, 0xc9, 0x15, 0x49, 0xa6, 0x3b, 0x2e, 0x19, 0xfa,
	0xd3, 0x23, 0xb0, 0x7d, 0x21, 0x18, 0xe3, 0x6a, 0x74, 0xaf, 0xab, 0x18, 0x6f, 0x54, 0xe4, 0x5a,
	0xcb, 0x5e, 0x81, 0xc7, 0xd3, 0x09, 0xd8, 0x92, 0xbc, 0xe5, 0x72, 0xb4, 0xd5, 0xe5, 0x9e, 0x6d,
	0xe4, 0x3e, 0xa8, 0x66, 0xf2, 0xae, 0x17, 0x74, 0x38, 0x66, 0xd7, 0x84, 0xde, 0x79, 0xf1, 0x74,
	0x12, 0x36, 0xb4, 0xe8, 0x5b, 0x90, 0x97, 0xf1, 0x38, 0x1c, 0xd4, 0xb3, 0xfc, 0xb5, 0x5c, 0xe6,
	0xdf, 0x87, 0xe0, 0xa0, 0xd0, 0x15, 0xfc, 0xef, 0xbd, 0xe4, 0x4f, 0x6e, 0xff, 0xf1, 0x6c, 0x3d,
	0xfe, 0xd9, 0xe0, 0xcb, 0xa9, 0x4f, 0x96, 0x5a, 0x52, 0x55, 0x42, 0x6d, 0x4a, 0x54, 0x72, 0xd5,
	0x2d, 0x87, 0x6e, 0x46, 0xfe, 0xc7, 0x05, 0xbf, 0x8f, 0xea, 0xe7, 0x30, 0x39, 0x21, 0xe4, 0xd7,
	0x70, 0xff, 0xc4, 0x55, 0x12, 0x66, 0xa1, 0x93, 0x6b, 0x35, 0xcb, 0xe0, 0x34, 0x90, 0xbf, 0x03,
	0x33, 0x27, 0xcc, 0xce, 0x23, 0x33, 0x9f, 0x65, 0xf3, 0xc8, 0x5c, 0x0f, 0x0f, 0x9c, 0x81, 0x31,
	0x61, 0x16, 0xe3, 0x48, 0x61, 0x3c, 0xcb, 0x30, 0x8e, 0xdc, 0x72, 0xbb, 0x1b, 0xf6, 0xed, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x84, 0x7e, 0x61, 0x6d, 0x03, 0x00, 0x00,
}
