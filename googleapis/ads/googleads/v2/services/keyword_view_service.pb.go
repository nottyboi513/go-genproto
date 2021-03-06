// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/keyword_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for [KeywordViewService.GetKeywordView][google.ads.googleads.v2.services.KeywordViewService.GetKeywordView].
type GetKeywordViewRequest struct {
	// Required. The resource name of the keyword view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordViewRequest) Reset()         { *m = GetKeywordViewRequest{} }
func (m *GetKeywordViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordViewRequest) ProtoMessage()    {}
func (*GetKeywordViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5de781ce320be574, []int{0}
}

func (m *GetKeywordViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordViewRequest.Unmarshal(m, b)
}
func (m *GetKeywordViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordViewRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordViewRequest.Merge(m, src)
}
func (m *GetKeywordViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordViewRequest.Size(m)
}
func (m *GetKeywordViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordViewRequest proto.InternalMessageInfo

func (m *GetKeywordViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordViewRequest)(nil), "google.ads.googleads.v2.services.GetKeywordViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/keyword_view_service.proto", fileDescriptor_5de781ce320be574)
}

var fileDescriptor_5de781ce320be574 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0x0b, 0xd3, 0x40,
	0x14, 0x27, 0x11, 0x04, 0x83, 0x3a, 0x04, 0xa4, 0x35, 0x0a, 0x96, 0x52, 0x44, 0x8a, 0xdc, 0x49,
	0x14, 0x84, 0x2b, 0x0e, 0xd7, 0xc1, 0x0a, 0x82, 0x54, 0x85, 0x0c, 0x12, 0x08, 0xd7, 0xe4, 0x19,
	0x4f, 0x93, 0x5c, 0xbd, 0x4b, 0x53, 0x44, 0x5c, 0xfc, 0x0a, 0x7e, 0x03, 0x47, 0xbf, 0x85, 0x6b,
	0x57, 0x37, 0x27, 0x07, 0x27, 0x77, 0x17, 0x07, 0x91, 0xf4, 0x72, 0x69, 0xaa, 0x86, 0x6e, 0x3f,
	0xf2, 0xfb, 0x73, 0xef, 0xfd, 0x5e, 0x9c, 0x59, 0x2a, 0x44, 0x9a, 0x01, 0x66, 0x89, 0xc2, 0x1a,
	0xd6, 0xa8, 0xf2, 0xb1, 0x02, 0x59, 0xf1, 0x18, 0x14, 0x7e, 0x05, 0x6f, 0xb6, 0x42, 0x26, 0x51,
	0xc5, 0x61, 0x1b, 0x35, 0x5f, 0xd1, 0x5a, 0x8a, 0x52, 0xb8, 0x23, 0xed, 0x40, 0x2c, 0x51, 0xa8,
	0x35, 0xa3, 0xca, 0x47, 0xc6, 0xec, 0xdd, 0xe9, 0x8b, 0x97, 0xa0, 0xc4, 0x46, 0xfe, 0x9d, 0xaf,
	0x73, 0xbd, 0xab, 0xc6, 0xb5, 0xe6, 0x98, 0x15, 0x85, 0x28, 0x59, 0xc9, 0x45, 0xa1, 0x1a, 0x76,
	0xd0, 0x61, 0xe3, 0x8c, 0x43, 0x51, 0x36, 0xc4, 0xb5, 0x0e, 0xf1, 0x9c, 0x43, 0x96, 0x44, 0x2b,
	0x78, 0xc1, 0x2a, 0x2e, 0x64, 0x23, 0xb8, 0xdc, 0x11, 0x98, 0x01, 0x34, 0x35, 0x7e, 0xe9, 0x5c,
	0x5a, 0x40, 0xf9, 0x50, 0xcf, 0x12, 0x70, 0xd8, 0x3e, 0x81, 0xd7, 0x1b, 0x50, 0xa5, 0xfb, 0xd8,
	0xb9, 0x60, 0xa4, 0x51, 0xc1, 0x72, 0x18, 0x5a, 0x23, 0xeb, 0xc6, 0xb9, 0xf9, 0xcd, 0x6f, 0xd4,
	0xfe, 0x45, 0xaf, 0x3b, 0x93, 0xc3, 0xde, 0x0d, 0x5a, 0x73, 0x85, 0x62, 0x91, 0xe3, 0x6e, 0xd6,
	0x79, 0x13, 0xf1, 0x88, 0xe5, 0xe0, 0xff, 0xb4, 0x1c, 0xb7, 0xc3, 0x3e, 0xd5, 0x65, 0xb9, 0x9f,
	0x2d, 0xe7, 0xe2, 0xf1, 0x0c, 0xee, 0x5d, 0x74, 0xaa, 0x61, 0xf4, 0xdf, 0xa9, 0x3d, 0xd4, 0x6b,
	0x6c, 0x8b, 0x47, 0x1d, 0xdb, 0xf8, 0xfe, 0x57, 0x7a, 0xbc, 0xe6, 0xfb, 0x2f, 0xdf, 0x3f, 0xd8,
	0xb7, 0x5c, 0x54, 0xdf, 0xea, 0xed, 0x11, 0x73, 0x2f, 0xde, 0xa8, 0x52, 0xe4, 0x20, 0x15, 0x9e,
	0x9a, 0xe3, 0xd5, 0x19, 0x0a, 0x4f, 0xdf, 0x79, 0x57, 0x76, 0x74, 0xd8, 0xd7, 0xc8, 0xfc, 0xb7,
	0xe5, 0x4c, 0x62, 0x91, 0x9f, 0xdc, 0x69, 0x3e, 0xf8, 0xb7, 0x9d, 0x65, 0x7d, 0xa5, 0xa5, 0xf5,
	0xec, 0x41, 0x63, 0x4e, 0x45, 0xc6, 0x8a, 0x14, 0x09, 0x99, 0xe2, 0x14, 0x8a, 0xfd, 0x0d, 0xf1,
	0xe1, 0xb9, 0xfe, 0xdf, 0x79, 0x66, 0xc0, 0x47, 0xfb, 0xcc, 0x82, 0xd2, 0x4f, 0xf6, 0x68, 0xa1,
	0x03, 0x69, 0xa2, 0x90, 0x86, 0x35, 0x0a, 0x7c, 0xd4, 0x3c, 0xac, 0x76, 0x46, 0x12, 0xd2, 0x44,
	0x85, 0xad, 0x24, 0x0c, 0xfc, 0xd0, 0x48, 0x7e, 0xd8, 0x13, 0xfd, 0x9d, 0x10, 0x9a, 0x28, 0x42,
	0x5a, 0x11, 0x21, 0x81, 0x4f, 0x88, 0x91, 0xad, 0xce, 0xee, 0xe7, 0xbc, 0xfd, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x51, 0xa0, 0xa7, 0x61, 0x75, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordViewServiceClient is the client API for KeywordViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordViewServiceClient interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error)
}

type keywordViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordViewServiceClient(cc grpc.ClientConnInterface) KeywordViewServiceClient {
	return &keywordViewServiceClient{cc}
}

func (c *keywordViewServiceClient) GetKeywordView(ctx context.Context, in *GetKeywordViewRequest, opts ...grpc.CallOption) (*resources.KeywordView, error) {
	out := new(resources.KeywordView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordViewService/GetKeywordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordViewServiceServer is the server API for KeywordViewService service.
type KeywordViewServiceServer interface {
	// Returns the requested keyword view in full detail.
	GetKeywordView(context.Context, *GetKeywordViewRequest) (*resources.KeywordView, error)
}

// UnimplementedKeywordViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordViewServiceServer struct {
}

func (*UnimplementedKeywordViewServiceServer) GetKeywordView(ctx context.Context, req *GetKeywordViewRequest) (*resources.KeywordView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeywordView not implemented")
}

func RegisterKeywordViewServiceServer(s *grpc.Server, srv KeywordViewServiceServer) {
	s.RegisterService(&_KeywordViewService_serviceDesc, srv)
}

func _KeywordViewService_GetKeywordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordViewService/GetKeywordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordViewServiceServer).GetKeywordView(ctx, req.(*GetKeywordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.KeywordViewService",
	HandlerType: (*KeywordViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordView",
			Handler:    _KeywordViewService_GetKeywordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/keyword_view_service.proto",
}
