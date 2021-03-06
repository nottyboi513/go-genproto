// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/media_file.proto

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

// A media file.
type MediaFile struct {
	// Immutable. The resource name of the media file.
	// Media file resource names have the form:
	//
	// `customers/{customer_id}/mediaFiles/{media_file_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the media file.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Immutable. Type of the media file.
	Type enums.MediaTypeEnum_MediaType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.MediaTypeEnum_MediaType" json:"type,omitempty"`
	// Output only. The mime type of the media file.
	MimeType enums.MimeTypeEnum_MimeType `protobuf:"varint,6,opt,name=mime_type,json=mimeType,proto3,enum=google.ads.googleads.v2.enums.MimeTypeEnum_MimeType" json:"mime_type,omitempty"`
	// Immutable. The URL of where the original media file was downloaded from (or a file
	// name). Only used for media of type AUDIO and IMAGE.
	SourceUrl *wrappers.StringValue `protobuf:"bytes,7,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// Immutable. The name of the media file. The name can be used by clients to help
	// identify previously uploaded media.
	Name *wrappers.StringValue `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The size of the media file in bytes.
	FileSize *wrappers.Int64Value `protobuf:"bytes,9,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// The specific type of the media file.
	//
	// Types that are valid to be assigned to Mediatype:
	//	*MediaFile_Image
	//	*MediaFile_MediaBundle
	//	*MediaFile_Audio
	//	*MediaFile_Video
	Mediatype            isMediaFile_Mediatype `protobuf_oneof:"mediatype"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MediaFile) Reset()         { *m = MediaFile{} }
func (m *MediaFile) String() string { return proto.CompactTextString(m) }
func (*MediaFile) ProtoMessage()    {}
func (*MediaFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_625a11b7599747b8, []int{0}
}

func (m *MediaFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaFile.Unmarshal(m, b)
}
func (m *MediaFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaFile.Marshal(b, m, deterministic)
}
func (m *MediaFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaFile.Merge(m, src)
}
func (m *MediaFile) XXX_Size() int {
	return xxx_messageInfo_MediaFile.Size(m)
}
func (m *MediaFile) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaFile.DiscardUnknown(m)
}

var xxx_messageInfo_MediaFile proto.InternalMessageInfo

func (m *MediaFile) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MediaFile) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MediaFile) GetType() enums.MediaTypeEnum_MediaType {
	if m != nil {
		return m.Type
	}
	return enums.MediaTypeEnum_UNSPECIFIED
}

func (m *MediaFile) GetMimeType() enums.MimeTypeEnum_MimeType {
	if m != nil {
		return m.MimeType
	}
	return enums.MimeTypeEnum_UNSPECIFIED
}

func (m *MediaFile) GetSourceUrl() *wrappers.StringValue {
	if m != nil {
		return m.SourceUrl
	}
	return nil
}

func (m *MediaFile) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MediaFile) GetFileSize() *wrappers.Int64Value {
	if m != nil {
		return m.FileSize
	}
	return nil
}

type isMediaFile_Mediatype interface {
	isMediaFile_Mediatype()
}

type MediaFile_Image struct {
	Image *MediaImage `protobuf:"bytes,3,opt,name=image,proto3,oneof"`
}

type MediaFile_MediaBundle struct {
	MediaBundle *MediaBundle `protobuf:"bytes,4,opt,name=media_bundle,json=mediaBundle,proto3,oneof"`
}

type MediaFile_Audio struct {
	Audio *MediaAudio `protobuf:"bytes,10,opt,name=audio,proto3,oneof"`
}

type MediaFile_Video struct {
	Video *MediaVideo `protobuf:"bytes,11,opt,name=video,proto3,oneof"`
}

func (*MediaFile_Image) isMediaFile_Mediatype() {}

func (*MediaFile_MediaBundle) isMediaFile_Mediatype() {}

func (*MediaFile_Audio) isMediaFile_Mediatype() {}

func (*MediaFile_Video) isMediaFile_Mediatype() {}

func (m *MediaFile) GetMediatype() isMediaFile_Mediatype {
	if m != nil {
		return m.Mediatype
	}
	return nil
}

func (m *MediaFile) GetImage() *MediaImage {
	if x, ok := m.GetMediatype().(*MediaFile_Image); ok {
		return x.Image
	}
	return nil
}

func (m *MediaFile) GetMediaBundle() *MediaBundle {
	if x, ok := m.GetMediatype().(*MediaFile_MediaBundle); ok {
		return x.MediaBundle
	}
	return nil
}

func (m *MediaFile) GetAudio() *MediaAudio {
	if x, ok := m.GetMediatype().(*MediaFile_Audio); ok {
		return x.Audio
	}
	return nil
}

func (m *MediaFile) GetVideo() *MediaVideo {
	if x, ok := m.GetMediatype().(*MediaFile_Video); ok {
		return x.Video
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MediaFile) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MediaFile_Image)(nil),
		(*MediaFile_MediaBundle)(nil),
		(*MediaFile_Audio)(nil),
		(*MediaFile_Video)(nil),
	}
}

// Encapsulates an Image.
type MediaImage struct {
	// Immutable. Raw image data.
	Data                 *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaImage) Reset()         { *m = MediaImage{} }
func (m *MediaImage) String() string { return proto.CompactTextString(m) }
func (*MediaImage) ProtoMessage()    {}
func (*MediaImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_625a11b7599747b8, []int{1}
}

func (m *MediaImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaImage.Unmarshal(m, b)
}
func (m *MediaImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaImage.Marshal(b, m, deterministic)
}
func (m *MediaImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaImage.Merge(m, src)
}
func (m *MediaImage) XXX_Size() int {
	return xxx_messageInfo_MediaImage.Size(m)
}
func (m *MediaImage) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaImage.DiscardUnknown(m)
}

var xxx_messageInfo_MediaImage proto.InternalMessageInfo

func (m *MediaImage) GetData() *wrappers.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

// Represents a ZIP archive media the content of which contains HTML5 assets.
type MediaBundle struct {
	// Immutable. Raw zipped data.
	Data                 *wrappers.BytesValue `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaBundle) Reset()         { *m = MediaBundle{} }
func (m *MediaBundle) String() string { return proto.CompactTextString(m) }
func (*MediaBundle) ProtoMessage()    {}
func (*MediaBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_625a11b7599747b8, []int{2}
}

func (m *MediaBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaBundle.Unmarshal(m, b)
}
func (m *MediaBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaBundle.Marshal(b, m, deterministic)
}
func (m *MediaBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaBundle.Merge(m, src)
}
func (m *MediaBundle) XXX_Size() int {
	return xxx_messageInfo_MediaBundle.Size(m)
}
func (m *MediaBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaBundle.DiscardUnknown(m)
}

var xxx_messageInfo_MediaBundle proto.InternalMessageInfo

func (m *MediaBundle) GetData() *wrappers.BytesValue {
	if m != nil {
		return m.Data
	}
	return nil
}

// Encapsulates an Audio.
type MediaAudio struct {
	// Output only. The duration of the Audio in milliseconds.
	AdDurationMillis     *wrappers.Int64Value `protobuf:"bytes,1,opt,name=ad_duration_millis,json=adDurationMillis,proto3" json:"ad_duration_millis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MediaAudio) Reset()         { *m = MediaAudio{} }
func (m *MediaAudio) String() string { return proto.CompactTextString(m) }
func (*MediaAudio) ProtoMessage()    {}
func (*MediaAudio) Descriptor() ([]byte, []int) {
	return fileDescriptor_625a11b7599747b8, []int{3}
}

func (m *MediaAudio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaAudio.Unmarshal(m, b)
}
func (m *MediaAudio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaAudio.Marshal(b, m, deterministic)
}
func (m *MediaAudio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaAudio.Merge(m, src)
}
func (m *MediaAudio) XXX_Size() int {
	return xxx_messageInfo_MediaAudio.Size(m)
}
func (m *MediaAudio) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaAudio.DiscardUnknown(m)
}

var xxx_messageInfo_MediaAudio proto.InternalMessageInfo

func (m *MediaAudio) GetAdDurationMillis() *wrappers.Int64Value {
	if m != nil {
		return m.AdDurationMillis
	}
	return nil
}

// Encapsulates a Video.
type MediaVideo struct {
	// Output only. The duration of the Video in milliseconds.
	AdDurationMillis *wrappers.Int64Value `protobuf:"bytes,1,opt,name=ad_duration_millis,json=adDurationMillis,proto3" json:"ad_duration_millis,omitempty"`
	// Immutable. The YouTube video ID (as seen in YouTube URLs).
	YoutubeVideoId *wrappers.StringValue `protobuf:"bytes,2,opt,name=youtube_video_id,json=youtubeVideoId,proto3" json:"youtube_video_id,omitempty"`
	// Output only. The Advertising Digital Identification code for this video, as defined by
	// the American Association of Advertising Agencies, used mainly for
	// television commercials.
	AdvertisingIdCode *wrappers.StringValue `protobuf:"bytes,3,opt,name=advertising_id_code,json=advertisingIdCode,proto3" json:"advertising_id_code,omitempty"`
	// Output only. The Industry Standard Commercial Identifier code for this video, used
	// mainly for television commercials.
	IsciCode             *wrappers.StringValue `protobuf:"bytes,4,opt,name=isci_code,json=isciCode,proto3" json:"isci_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MediaVideo) Reset()         { *m = MediaVideo{} }
func (m *MediaVideo) String() string { return proto.CompactTextString(m) }
func (*MediaVideo) ProtoMessage()    {}
func (*MediaVideo) Descriptor() ([]byte, []int) {
	return fileDescriptor_625a11b7599747b8, []int{4}
}

func (m *MediaVideo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaVideo.Unmarshal(m, b)
}
func (m *MediaVideo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaVideo.Marshal(b, m, deterministic)
}
func (m *MediaVideo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaVideo.Merge(m, src)
}
func (m *MediaVideo) XXX_Size() int {
	return xxx_messageInfo_MediaVideo.Size(m)
}
func (m *MediaVideo) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaVideo.DiscardUnknown(m)
}

var xxx_messageInfo_MediaVideo proto.InternalMessageInfo

func (m *MediaVideo) GetAdDurationMillis() *wrappers.Int64Value {
	if m != nil {
		return m.AdDurationMillis
	}
	return nil
}

func (m *MediaVideo) GetYoutubeVideoId() *wrappers.StringValue {
	if m != nil {
		return m.YoutubeVideoId
	}
	return nil
}

func (m *MediaVideo) GetAdvertisingIdCode() *wrappers.StringValue {
	if m != nil {
		return m.AdvertisingIdCode
	}
	return nil
}

func (m *MediaVideo) GetIsciCode() *wrappers.StringValue {
	if m != nil {
		return m.IsciCode
	}
	return nil
}

func init() {
	proto.RegisterType((*MediaFile)(nil), "google.ads.googleads.v2.resources.MediaFile")
	proto.RegisterType((*MediaImage)(nil), "google.ads.googleads.v2.resources.MediaImage")
	proto.RegisterType((*MediaBundle)(nil), "google.ads.googleads.v2.resources.MediaBundle")
	proto.RegisterType((*MediaAudio)(nil), "google.ads.googleads.v2.resources.MediaAudio")
	proto.RegisterType((*MediaVideo)(nil), "google.ads.googleads.v2.resources.MediaVideo")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/media_file.proto", fileDescriptor_625a11b7599747b8)
}

var fileDescriptor_625a11b7599747b8 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0xc9, 0x57, 0x1b, 0x4f, 0x4a, 0x55, 0x86, 0x1b, 0x53, 0x2a, 0x68, 0x23, 0x2a, 0x55,
	0x88, 0x8e, 0x91, 0x29, 0xbd, 0x30, 0x02, 0x61, 0x17, 0x28, 0x29, 0xea, 0x07, 0x2e, 0x8d, 0x10,
	0x8a, 0x64, 0x4d, 0x32, 0x53, 0x33, 0x92, 0xed, 0x89, 0x3c, 0x76, 0x50, 0x5a, 0xf5, 0x21, 0x78,
	0x05, 0x2e, 0x79, 0x06, 0x9e, 0x80, 0xa7, 0xe8, 0x75, 0x1f, 0x61, 0xf7, 0x66, 0x35, 0xe3, 0xb1,
	0x13, 0xed, 0x6e, 0x37, 0xe9, 0x6a, 0xef, 0xce, 0xe4, 0xfc, 0xff, 0x3f, 0x9f, 0x39, 0x73, 0x66,
	0x02, 0xec, 0x90, 0xf3, 0x30, 0xa2, 0x16, 0x26, 0xc2, 0x2a, 0x42, 0x19, 0x4d, 0x6c, 0x2b, 0xa5,
	0x82, 0xe7, 0xe9, 0x88, 0x0a, 0x2b, 0xa6, 0x84, 0xe1, 0xe0, 0x9a, 0x45, 0x14, 0x8d, 0x53, 0x9e,
	0x71, 0xb8, 0x53, 0x08, 0x11, 0x26, 0x02, 0x55, 0x1e, 0x34, 0xb1, 0x51, 0xe5, 0xd9, 0x44, 0x8f,
	0x61, 0x69, 0x92, 0xc7, 0x25, 0x32, 0x9b, 0x8e, 0x35, 0x72, 0x73, 0x7f, 0x81, 0x9e, 0xc5, 0x74,
	0x5e, 0xfe, 0x69, 0x29, 0x1f, 0x33, 0xeb, 0x9a, 0xd1, 0x88, 0x04, 0x43, 0xfa, 0x27, 0x9e, 0x30,
	0x9e, 0x6a, 0xc1, 0x47, 0x73, 0x82, 0xb2, 0x2a, 0x9d, 0xfa, 0x44, 0xa7, 0xd4, 0x6a, 0x98, 0x5f,
	0x5b, 0x7f, 0xa5, 0x78, 0x3c, 0xa6, 0xa9, 0xd0, 0xf9, 0xad, 0x39, 0x2b, 0x4e, 0x12, 0x9e, 0xe1,
	0x8c, 0xf1, 0x44, 0x67, 0xbb, 0x7f, 0xaf, 0x02, 0xe3, 0x54, 0x56, 0xff, 0x13, 0x8b, 0x28, 0x3c,
	0x07, 0xef, 0x97, 0xf4, 0x20, 0xc1, 0x31, 0x35, 0x6b, 0xdb, 0xb5, 0x3d, 0xc3, 0xfb, 0xfc, 0xde,
	0x6d, 0x3d, 0x73, 0x3f, 0x03, 0xdd, 0x59, 0x77, 0x74, 0x34, 0x66, 0x02, 0x8d, 0x78, 0x6c, 0x55,
	0x08, 0x7f, 0xad, 0x04, 0x9c, 0xe1, 0x98, 0xc2, 0x2f, 0x41, 0x9d, 0x11, 0xb3, 0xbe, 0x5d, 0xdb,
	0xeb, 0xd8, 0x1f, 0x6b, 0x13, 0x2a, 0x2b, 0x45, 0xbd, 0x24, 0x3b, 0x3c, 0xe8, 0xe3, 0x28, 0xa7,
	0x5e, 0xe3, 0xde, 0x6d, 0xf8, 0x75, 0x46, 0xe0, 0x39, 0x68, 0xca, 0xc6, 0x98, 0xad, 0xed, 0xda,
	0xde, 0xba, 0x7d, 0x88, 0x1e, 0x3b, 0x1b, 0xd5, 0x48, 0xa4, 0xbe, 0xfb, 0xdb, 0x74, 0x4c, 0x7f,
	0x4c, 0xf2, 0x78, 0xb6, 0x92, 0xb8, 0x96, 0xaf, 0x40, 0xb0, 0x0f, 0x8c, 0xaa, 0xdd, 0xe6, 0x8a,
	0xa2, 0x1e, 0x2c, 0xa2, 0xb2, 0x98, 0xce, 0xa0, 0x7a, 0x51, 0x94, 0xd8, 0x8e, 0xf5, 0x12, 0x7e,
	0x0f, 0x80, 0xee, 0x54, 0x9e, 0x46, 0xe6, 0xaa, 0xda, 0xe2, 0xd6, 0x2b, 0x5b, 0xbc, 0xcc, 0x52,
	0x96, 0x84, 0xd5, 0x1e, 0x5b, 0xbe, 0x51, 0x98, 0xae, 0xd2, 0x08, 0x7e, 0x0d, 0x9a, 0xaa, 0xc9,
	0xed, 0x65, 0xbd, 0x4a, 0x0e, 0xbf, 0x05, 0x86, 0x1c, 0xde, 0x40, 0xb0, 0x1b, 0x6a, 0x1a, 0x4b,
	0xb6, 0xb6, 0x2d, 0x2d, 0x97, 0xec, 0x86, 0xc2, 0x13, 0xd0, 0x62, 0x31, 0x0e, 0xa9, 0xd9, 0x50,
	0xd6, 0x7d, 0xb4, 0x70, 0xfa, 0x8b, 0xbe, 0xf6, 0xa4, 0x49, 0xd5, 0xf1, 0xf3, 0x7b, 0x7e, 0x81,
	0x80, 0xbf, 0x83, 0xb5, 0x62, 0xf4, 0x87, 0x79, 0x42, 0x22, 0x6a, 0x36, 0x15, 0x12, 0x2d, 0x8b,
	0xf4, 0x94, 0xab, 0x64, 0x76, 0xe2, 0xd9, 0x6f, 0xb2, 0x4a, 0x9c, 0x13, 0xc6, 0x4d, 0xf0, 0xb4,
	0x2a, 0x5d, 0x69, 0x52, 0x5b, 0x96, 0x55, 0x2a, 0x84, 0x64, 0x4d, 0x18, 0xa1, 0xdc, 0xec, 0x3c,
	0x8d, 0xd5, 0x97, 0xa6, 0x6a, 0xc7, 0x0a, 0xe1, 0x5c, 0x3d, 0xb8, 0xfe, 0x32, 0xf7, 0x00, 0x7e,
	0x31, 0xca, 0x45, 0xc6, 0x63, 0x9a, 0x0a, 0xeb, 0xb6, 0x0c, 0xef, 0x8a, 0x87, 0x42, 0xe6, 0x85,
	0x75, 0x3b, 0x7b, 0x87, 0xee, 0xbc, 0x0e, 0x30, 0xd4, 0x52, 0x0e, 0x69, 0xd7, 0x03, 0x60, 0xd6,
	0x71, 0x78, 0x00, 0x9a, 0x04, 0x67, 0x58, 0x5d, 0xc5, 0xd7, 0x9d, 0xb4, 0x37, 0xcd, 0xa8, 0x98,
	0x1f, 0x12, 0xa9, 0xee, 0x1e, 0x81, 0xce, 0x5c, 0x8b, 0xdf, 0x12, 0x32, 0xd0, 0x85, 0xa8, 0xa6,
	0xc2, 0x33, 0x00, 0x31, 0x09, 0x48, 0x9e, 0xaa, 0x07, 0x24, 0x88, 0x59, 0x14, 0x31, 0xf1, 0x28,
	0xf1, 0xe5, 0x01, 0xdc, 0xc0, 0xe4, 0x07, 0x6d, 0x3d, 0x55, 0xce, 0xee, 0x7f, 0x75, 0x8d, 0x57,
	0x7d, 0x7e, 0xd7, 0x78, 0xf8, 0x0b, 0xd8, 0x98, 0xf2, 0x3c, 0xcb, 0x87, 0x34, 0x50, 0x47, 0x17,
	0x54, 0x0f, 0xd1, 0x12, 0x37, 0x6d, 0x5d, 0x5b, 0x55, 0x69, 0x3d, 0x02, 0x7f, 0x05, 0x1f, 0x62,
	0x32, 0xa1, 0x69, 0xc6, 0x04, 0x4b, 0xc2, 0x80, 0x91, 0x60, 0xc4, 0x49, 0x79, 0x85, 0x16, 0xf2,
	0x1a, 0xfe, 0x07, 0x73, 0xee, 0x1e, 0x39, 0xe2, 0x84, 0xc2, 0xef, 0x80, 0xc1, 0xc4, 0x88, 0x15,
	0xa0, 0xe6, 0xb2, 0xa0, 0xb6, 0xf4, 0x48, 0xbf, 0xf7, 0xbc, 0x06, 0x76, 0x47, 0x3c, 0x5e, 0x3c,
	0xcc, 0xde, 0x7a, 0x35, 0x95, 0x17, 0x92, 0x7b, 0x51, 0xfb, 0xe3, 0x44, 0x9b, 0x42, 0x1e, 0xe1,
	0x24, 0x44, 0x3c, 0x0d, 0xad, 0x90, 0x26, 0xea, 0xab, 0xd6, 0x6c, 0xa2, 0xdf, 0xf0, 0x27, 0xfa,
	0x4d, 0x15, 0xfd, 0x53, 0x6f, 0x1c, 0xbb, 0xee, 0xbf, 0xf5, 0x9d, 0xe3, 0x02, 0xe9, 0x12, 0x81,
	0x8a, 0x50, 0x46, 0x7d, 0x1b, 0xf9, 0xa5, 0xf2, 0xff, 0x52, 0x33, 0x70, 0x89, 0x18, 0x54, 0x9a,
	0x41, 0xdf, 0x1e, 0x54, 0x9a, 0x87, 0xfa, 0x6e, 0x91, 0x70, 0x1c, 0x97, 0x08, 0xc7, 0xa9, 0x54,
	0x8e, 0xd3, 0xb7, 0x1d, 0xa7, 0xd2, 0x0d, 0x57, 0x54, 0xb1, 0x5f, 0xbd, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x20, 0xab, 0x79, 0xcd, 0xf0, 0x07, 0x00, 0x00,
}
