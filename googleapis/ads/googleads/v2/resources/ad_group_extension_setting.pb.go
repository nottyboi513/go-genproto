// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/ad_group_extension_setting.proto

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

// An ad group extension setting.
type AdGroupExtensionSetting struct {
	// Immutable. The resource name of the ad group extension setting.
	// AdGroupExtensionSetting resource names have the form:
	//
	// `customers/{customer_id}/adGroupExtensionSettings/{ad_group_id}~{extension_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The extension type of the ad group extension setting.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v2.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// Immutable. The resource name of the ad group. The linked extension feed items will
	// serve under this ad group.
	// AdGroup resource names have the form:
	//
	// `customers/{customer_id}/adGroups/{ad_group_id}`
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The resource names of the extension feed items to serve under the ad group.
	// ExtensionFeedItem resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ExtensionFeedItems []*wrappers.StringValue `protobuf:"bytes,4,rep,name=extension_feed_items,json=extensionFeedItems,proto3" json:"extension_feed_items,omitempty"`
	// The device for which the extensions will serve. Optional.
	Device               enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice `protobuf:"varint,5,opt,name=device,proto3,enum=google.ads.googleads.v2.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *AdGroupExtensionSetting) Reset()         { *m = AdGroupExtensionSetting{} }
func (m *AdGroupExtensionSetting) String() string { return proto.CompactTextString(m) }
func (*AdGroupExtensionSetting) ProtoMessage()    {}
func (*AdGroupExtensionSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_aae21be15f4744e1, []int{0}
}

func (m *AdGroupExtensionSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupExtensionSetting.Unmarshal(m, b)
}
func (m *AdGroupExtensionSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupExtensionSetting.Marshal(b, m, deterministic)
}
func (m *AdGroupExtensionSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupExtensionSetting.Merge(m, src)
}
func (m *AdGroupExtensionSetting) XXX_Size() int {
	return xxx_messageInfo_AdGroupExtensionSetting.Size(m)
}
func (m *AdGroupExtensionSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupExtensionSetting.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupExtensionSetting proto.InternalMessageInfo

func (m *AdGroupExtensionSetting) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupExtensionSetting) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *AdGroupExtensionSetting) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupExtensionSetting) GetExtensionFeedItems() []*wrappers.StringValue {
	if m != nil {
		return m.ExtensionFeedItems
	}
	return nil
}

func (m *AdGroupExtensionSetting) GetDevice() enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice {
	if m != nil {
		return m.Device
	}
	return enums.ExtensionSettingDeviceEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupExtensionSetting)(nil), "google.ads.googleads.v2.resources.AdGroupExtensionSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/ad_group_extension_setting.proto", fileDescriptor_aae21be15f4744e1)
}

var fileDescriptor_aae21be15f4744e1 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0x93, 0x3f, 0xfd, 0x61, 0xa1, 0x3d, 0x58, 0x48, 0x84, 0xaa, 0x82, 0xb4, 0x52, 0xa5,
	0x08, 0xa1, 0x5d, 0x64, 0x2e, 0xc8, 0x20, 0xa4, 0xb5, 0x28, 0x11, 0x3d, 0xa0, 0xc8, 0x45, 0x39,
	0xa0, 0x20, 0x6b, 0x93, 0x9d, 0xb8, 0x46, 0xf1, 0xae, 0xe5, 0x5d, 0x07, 0xaa, 0xaa, 0x20, 0x8e,
	0xbc, 0x06, 0x47, 0x1e, 0x80, 0x87, 0xe0, 0x29, 0x7a, 0xee, 0x23, 0xf4, 0x84, 0x62, 0x7b, 0x9d,
	0x04, 0x30, 0x29, 0xb7, 0xcf, 0x3b, 0xdf, 0x7c, 0x33, 0xf3, 0xcd, 0xae, 0x91, 0x17, 0x4a, 0x19,
	0x4e, 0x81, 0x30, 0xae, 0x48, 0x01, 0xe7, 0x68, 0xe6, 0x90, 0x14, 0x94, 0xcc, 0xd2, 0x31, 0x28,
	0xc2, 0x78, 0x10, 0xa6, 0x32, 0x4b, 0x02, 0xf8, 0xa0, 0x41, 0xa8, 0x48, 0x8a, 0x40, 0x81, 0xd6,
	0x91, 0x08, 0x71, 0x92, 0x4a, 0x2d, 0xed, 0xdd, 0x22, 0x11, 0x33, 0xae, 0x70, 0xa5, 0x81, 0x67,
	0x0e, 0xae, 0x34, 0xb6, 0x9f, 0xd6, 0x95, 0x01, 0x91, 0xc5, 0x8a, 0xfc, 0xa6, 0x1c, 0x70, 0x98,
	0x45, 0x63, 0x28, 0x0a, 0x6c, 0x3b, 0x57, 0xcd, 0xd6, 0x27, 0x89, 0xc9, 0xb9, 0x67, 0x72, 0x92,
	0x88, 0x4c, 0x22, 0x98, 0xf2, 0x60, 0x04, 0xc7, 0x6c, 0x16, 0xc9, 0xb4, 0x24, 0xdc, 0x59, 0x22,
	0x98, 0x46, 0xcb, 0xd0, 0xdd, 0x32, 0x94, 0x7f, 0x8d, 0xb2, 0x09, 0x79, 0x9f, 0xb2, 0x24, 0x81,
	0x54, 0x95, 0xf1, 0x9d, 0xa5, 0x54, 0x26, 0x84, 0xd4, 0x4c, 0x47, 0x52, 0x94, 0xd1, 0xbd, 0xef,
	0x2d, 0x74, 0x9b, 0xf2, 0xde, 0xdc, 0xb2, 0x03, 0xd3, 0xd9, 0x51, 0x31, 0x96, 0xfd, 0x16, 0x6d,
	0x9a, 0x5a, 0x81, 0x60, 0x31, 0xb4, 0xad, 0x8e, 0xd5, 0xbd, 0xee, 0x3d, 0x3e, 0xa7, 0xad, 0x4b,
	0xea, 0xa0, 0x87, 0x0b, 0xfb, 0x4a, 0x94, 0x44, 0x0a, 0x8f, 0x65, 0x4c, 0x6a, 0x04, 0xfd, 0x9b,
	0x46, 0xee, 0x15, 0x8b, 0xc1, 0x7e, 0x87, 0xb6, 0x56, 0xcd, 0x68, 0x37, 0x3a, 0x56, 0x77, 0xcb,
	0x79, 0x86, 0xeb, 0x56, 0x94, 0x3b, 0x88, 0x2b, 0xd9, 0xd7, 0x27, 0x09, 0x1c, 0x88, 0x2c, 0x5e,
	0x3d, 0xf1, 0x9a, 0xe7, 0xb4, 0xe5, 0x6f, 0xc2, 0xf2, 0x99, 0xcd, 0xd0, 0x35, 0x73, 0x33, 0xda,
	0xcd, 0x8e, 0xd5, 0xbd, 0xe1, 0xec, 0x98, 0x2a, 0xc6, 0x37, 0x7c, 0xa4, 0xd3, 0x48, 0x84, 0x03,
	0x36, 0xcd, 0xc0, 0xeb, 0xe6, 0x33, 0xee, 0xa1, 0xce, 0xba, 0x19, 0xfd, 0xff, 0x59, 0x01, 0xec,
	0x4f, 0xe8, 0xd6, 0x62, 0x9c, 0x09, 0x00, 0x0f, 0x22, 0x0d, 0xb1, 0x6a, 0xff, 0xd7, 0x69, 0xae,
	0x2d, 0x47, 0x2e, 0xe9, 0x03, 0x74, 0xbf, 0xb6, 0x56, 0x35, 0xdf, 0x0b, 0x00, 0xfe, 0x52, 0x43,
	0xec, 0xdb, 0xf0, 0xeb, 0x91, 0xb2, 0x8f, 0xd1, 0x46, 0x71, 0x11, 0xdb, 0xad, 0xdc, 0xc7, 0xfe,
	0x55, 0x7d, 0x2c, 0xd7, 0xf3, 0x3c, 0x4f, 0x5e, 0x35, 0x74, 0x25, 0xe4, 0x97, 0xfa, 0xee, 0x67,
	0xeb, 0x82, 0x7e, 0xfc, 0xf7, 0xfd, 0xdb, 0x87, 0xe3, 0x4c, 0x69, 0x19, 0x43, 0xaa, 0xc8, 0xa9,
	0x81, 0x67, 0x84, 0xfd, 0x99, 0xad, 0xc8, 0x69, 0xfd, 0x63, 0x3e, 0xf3, 0xbe, 0x34, 0xd0, 0xfe,
	0x58, 0xc6, 0x78, 0xed, 0x73, 0xf6, 0x76, 0x6a, 0xda, 0xe9, 0xcf, 0x37, 0xd1, 0xb7, 0xde, 0x1c,
	0x96, 0x12, 0xa1, 0x9c, 0x32, 0x11, 0x62, 0x99, 0x86, 0x24, 0x04, 0x91, 0xef, 0x89, 0x2c, 0x06,
	0xfb, 0xcb, 0x4f, 0xe7, 0x49, 0x85, 0xbe, 0x36, 0x9a, 0x3d, 0x4a, 0xbf, 0x35, 0x76, 0x7b, 0x85,
	0x24, 0xe5, 0x0a, 0x17, 0x70, 0x8e, 0x06, 0x0e, 0xf6, 0x0d, 0xf3, 0x87, 0xe1, 0x0c, 0x29, 0x57,
	0xc3, 0x8a, 0x33, 0x1c, 0x38, 0xc3, 0x8a, 0x73, 0xd1, 0xd8, 0x2f, 0x02, 0xae, 0x4b, 0xb9, 0x72,
	0xdd, 0x8a, 0xe5, 0xba, 0x03, 0xc7, 0x75, 0x2b, 0xde, 0x68, 0x23, 0x6f, 0xf6, 0xd1, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc7, 0xcb, 0xff, 0x22, 0x20, 0x05, 0x00, 0x00,
}
