// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/payments_account.proto

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

// A payments account, which can be used to set up billing for an Ads customer.
type PaymentsAccount struct {
	// The resource name of the payments account.
	// PaymentsAccount resource names have the form:
	//
	// `customers/{customer_id}/paymentsAccounts/{payments_account_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// A 16 digit ID used to identify a payments account.
	PaymentsAccountId *wrappers.StringValue `protobuf:"bytes,2,opt,name=payments_account_id,json=paymentsAccountId,proto3" json:"payments_account_id,omitempty"`
	// The name of the payments account.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The currency code of the payments account.
	// A subset of the currency codes derived from the ISO 4217 standard is
	// supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,4,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// A 12 digit ID used to identify the payments profile associated with the
	// payments account.
	PaymentsProfileId *wrappers.StringValue `protobuf:"bytes,5,opt,name=payments_profile_id,json=paymentsProfileId,proto3" json:"payments_profile_id,omitempty"`
	// A secondary payments profile ID present in uncommon situations, e.g.
	// when a sequential liability agreement has been arranged.
	SecondaryPaymentsProfileId *wrappers.StringValue `protobuf:"bytes,6,opt,name=secondary_payments_profile_id,json=secondaryPaymentsProfileId,proto3" json:"secondary_payments_profile_id,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}              `json:"-"`
	XXX_unrecognized           []byte                `json:"-"`
	XXX_sizecache              int32                 `json:"-"`
}

func (m *PaymentsAccount) Reset()         { *m = PaymentsAccount{} }
func (m *PaymentsAccount) String() string { return proto.CompactTextString(m) }
func (*PaymentsAccount) ProtoMessage()    {}
func (*PaymentsAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_37c21bb407f87d61, []int{0}
}

func (m *PaymentsAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentsAccount.Unmarshal(m, b)
}
func (m *PaymentsAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentsAccount.Marshal(b, m, deterministic)
}
func (m *PaymentsAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentsAccount.Merge(m, src)
}
func (m *PaymentsAccount) XXX_Size() int {
	return xxx_messageInfo_PaymentsAccount.Size(m)
}
func (m *PaymentsAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentsAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentsAccount proto.InternalMessageInfo

func (m *PaymentsAccount) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *PaymentsAccount) GetPaymentsAccountId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccountId
	}
	return nil
}

func (m *PaymentsAccount) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PaymentsAccount) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *PaymentsAccount) GetPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsProfileId
	}
	return nil
}

func (m *PaymentsAccount) GetSecondaryPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.SecondaryPaymentsProfileId
	}
	return nil
}

func init() {
	proto.RegisterType((*PaymentsAccount)(nil), "google.ads.googleads.v3.resources.PaymentsAccount")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/payments_account.proto", fileDescriptor_37c21bb407f87d61)
}

var fileDescriptor_37c21bb407f87d61 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6b, 0xd4, 0x40,
	0x18, 0x65, 0xd3, 0xb5, 0xe0, 0xd8, 0x22, 0x46, 0x0f, 0xeb, 0x52, 0xa5, 0x55, 0x0a, 0x7b, 0x9a,
	0x11, 0x73, 0x29, 0xe3, 0x29, 0xf5, 0x50, 0x2a, 0x22, 0x61, 0x85, 0x3d, 0xc8, 0x42, 0x98, 0xce,
	0x4c, 0x43, 0x60, 0x33, 0xdf, 0x30, 0x93, 0x54, 0x96, 0xd2, 0x83, 0x3f, 0xc1, 0xbf, 0xe0, 0xd1,
	0x9f, 0xe2, 0x4f, 0xe9, 0xaf, 0x90, 0x64, 0x32, 0x53, 0x37, 0x88, 0x6e, 0x6f, 0x2f, 0xf9, 0xde,
	0x7b, 0xdf, 0x7b, 0x99, 0x09, 0x3a, 0x29, 0x00, 0x8a, 0x95, 0x24, 0x4c, 0x58, 0xe2, 0x60, 0x8b,
	0xae, 0x12, 0x62, 0xa4, 0x85, 0xc6, 0x70, 0x69, 0x89, 0x66, 0xeb, 0x4a, 0xaa, 0xda, 0xe6, 0x8c,
	0x73, 0x68, 0x54, 0x8d, 0xb5, 0x81, 0x1a, 0xe2, 0x23, 0x47, 0xc7, 0x4c, 0x58, 0x1c, 0x94, 0xf8,
	0x2a, 0xc1, 0x41, 0x39, 0x7d, 0xee, 0xcd, 0x75, 0x19, 0xfc, 0x9c, 0x7a, 0xfa, 0xb2, 0x1f, 0x75,
	0x4f, 0x17, 0xcd, 0x25, 0xf9, 0x6a, 0x98, 0xd6, 0xd2, 0xd8, 0x7e, 0x7e, 0xf0, 0x87, 0x94, 0x29,
	0x05, 0x35, 0xab, 0x4b, 0x50, 0xfd, 0xf4, 0xd5, 0xf7, 0x31, 0x7a, 0x9c, 0xf5, 0xb1, 0x52, 0x97,
	0x2a, 0x7e, 0x8d, 0xf6, 0xfd, 0x8e, 0x5c, 0xb1, 0x4a, 0x4e, 0x46, 0x87, 0xa3, 0xd9, 0xc3, 0xf9,
	0x9e, 0x7f, 0xf9, 0x89, 0x55, 0x32, 0xfe, 0x88, 0x9e, 0x0e, 0xeb, 0xe4, 0xa5, 0x98, 0x44, 0x87,
	0xa3, 0xd9, 0xa3, 0xb7, 0x07, 0x7d, 0x0f, 0xec, 0x43, 0xe1, 0xcf, 0xb5, 0x29, 0x55, 0xb1, 0x60,
	0xab, 0x46, 0xce, 0x9f, 0xe8, 0xcd, 0x85, 0xe7, 0x22, 0x7e, 0x83, 0xc6, 0xdd, 0xa6, 0x9d, 0x2d,
	0xe4, 0x1d, 0x33, 0x4e, 0xd1, 0x3e, 0x6f, 0x8c, 0x91, 0x8a, 0xaf, 0x73, 0x0e, 0x42, 0x4e, 0xc6,
	0x5b, 0x48, 0xf7, 0xbc, 0xe4, 0x3d, 0x88, 0xcd, 0x0a, 0xda, 0xc0, 0x65, 0xb9, 0x92, 0x6d, 0x85,
	0x07, 0xf7, 0xa9, 0x90, 0x39, 0xdd, 0xb9, 0x88, 0x73, 0xf4, 0xc2, 0x4a, 0x0e, 0x4a, 0x30, 0xb3,
	0xce, 0xff, 0xe6, 0xbb, 0xbb, 0x85, 0xef, 0x34, 0x58, 0x64, 0xc3, 0x05, 0xb4, 0xb8, 0x4d, 0x05,
	0x9a, 0xdd, 0x5d, 0x90, 0x1e, 0xe9, 0xd2, 0x62, 0x0e, 0x15, 0x19, 0x9e, 0xe2, 0x09, 0x6f, 0x6c,
	0x0d, 0x95, 0x34, 0x96, 0x5c, 0x7b, 0x78, 0x43, 0x06, 0x9f, 0xde, 0x92, 0xeb, 0xe1, 0x29, 0xde,
	0x9c, 0x7e, 0x8b, 0xd0, 0x31, 0x87, 0x0a, 0xff, 0xf7, 0x5a, 0x9e, 0x3e, 0x1b, 0x2c, 0xcd, 0xda,
	0x4e, 0xd9, 0xe8, 0xcb, 0x87, 0x5e, 0x5a, 0xc0, 0x8a, 0xa9, 0x02, 0x83, 0x29, 0x48, 0x21, 0x55,
	0xd7, 0x98, 0xdc, 0xc5, 0xfe, 0xc7, 0xaf, 0xf2, 0x2e, 0xa0, 0x1f, 0xd1, 0xce, 0x59, 0x9a, 0xfe,
	0x8c, 0x8e, 0xce, 0x9c, 0x65, 0x2a, 0x2c, 0x76, 0xb0, 0x45, 0x8b, 0x04, 0xcf, 0x3d, 0xf3, 0x97,
	0xe7, 0x2c, 0x53, 0x61, 0x97, 0x81, 0xb3, 0x5c, 0x24, 0xcb, 0xc0, 0xb9, 0x8d, 0x8e, 0xdd, 0x80,
	0xd2, 0x54, 0x58, 0x4a, 0x03, 0x8b, 0xd2, 0x45, 0x42, 0x69, 0xe0, 0x5d, 0xec, 0x76, 0x61, 0x93,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x8d, 0xc1, 0xcd, 0xd6, 0x03, 0x00, 0x00,
}