// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/age_range_view_service.proto

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

// Request message for [AgeRangeViewService.GetAgeRangeView][google.ads.googleads.v2.services.AgeRangeViewService.GetAgeRangeView].
type GetAgeRangeViewRequest struct {
	// Required. The resource name of the age range view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgeRangeViewRequest) Reset()         { *m = GetAgeRangeViewRequest{} }
func (m *GetAgeRangeViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetAgeRangeViewRequest) ProtoMessage()    {}
func (*GetAgeRangeViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_057905eda0e1111f, []int{0}
}

func (m *GetAgeRangeViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgeRangeViewRequest.Unmarshal(m, b)
}
func (m *GetAgeRangeViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgeRangeViewRequest.Marshal(b, m, deterministic)
}
func (m *GetAgeRangeViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgeRangeViewRequest.Merge(m, src)
}
func (m *GetAgeRangeViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetAgeRangeViewRequest.Size(m)
}
func (m *GetAgeRangeViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgeRangeViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgeRangeViewRequest proto.InternalMessageInfo

func (m *GetAgeRangeViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAgeRangeViewRequest)(nil), "google.ads.googleads.v2.services.GetAgeRangeViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/age_range_view_service.proto", fileDescriptor_057905eda0e1111f)
}

var fileDescriptor_057905eda0e1111f = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0x8b, 0xd4, 0x40,
	0x1c, 0x25, 0x11, 0x04, 0x83, 0x22, 0x44, 0xd0, 0x35, 0x0a, 0x2e, 0xc7, 0x89, 0x72, 0xe0, 0x0c,
	0x46, 0x10, 0x19, 0xb9, 0x62, 0xb6, 0x89, 0x95, 0x1c, 0x2b, 0xa4, 0x90, 0x40, 0x98, 0x4b, 0x7e,
	0x8e, 0x03, 0xc9, 0xcc, 0x3a, 0x93, 0xcd, 0x15, 0x62, 0xe3, 0x57, 0xf0, 0x1b, 0x58, 0xfa, 0x3d,
	0x6c, 0xae, 0xb5, 0xb3, 0xb2, 0xb0, 0xf2, 0x23, 0x5c, 0xa3, 0xe4, 0x26, 0x93, 0xcb, 0x7a, 0x1b,
	0xb6, 0x7b, 0xcc, 0x7b, 0xef, 0xf7, 0xe7, 0xfd, 0x26, 0x38, 0xe4, 0x4a, 0xf1, 0x0a, 0x30, 0x2b,
	0x0d, 0xb6, 0xb0, 0x43, 0x6d, 0x8c, 0x0d, 0xe8, 0x56, 0x14, 0x60, 0x30, 0xe3, 0x90, 0x6b, 0x26,
	0x39, 0xe4, 0xad, 0x80, 0x93, 0xbc, 0x7f, 0x47, 0x2b, 0xad, 0x1a, 0x15, 0xce, 0xad, 0x07, 0xb1,
	0xd2, 0xa0, 0xc1, 0x8e, 0xda, 0x18, 0x39, 0x7b, 0xf4, 0x7c, 0xaa, 0x81, 0x06, 0xa3, 0xd6, 0xfa,
	0x72, 0x07, 0x5b, 0x39, 0xba, 0xef, 0x7c, 0x2b, 0x81, 0x99, 0x94, 0xaa, 0x61, 0x8d, 0x50, 0xd2,
	0xf4, 0xec, 0x9d, 0x11, 0x5b, 0x54, 0x02, 0x64, 0xd3, 0x13, 0x0f, 0x46, 0xc4, 0x3b, 0x01, 0x55,
	0x99, 0x1f, 0xc3, 0x7b, 0xd6, 0x0a, 0xa5, 0x7b, 0xc1, 0xdd, 0x91, 0xc0, 0x8d, 0x60, 0xa9, 0xbd,
	0x2a, 0xb8, 0x9d, 0x40, 0x43, 0x39, 0x2c, 0xbb, 0x61, 0x52, 0x01, 0x27, 0x4b, 0xf8, 0xb0, 0x06,
	0xd3, 0x84, 0xcb, 0xe0, 0x86, 0xd3, 0xe6, 0x92, 0xd5, 0x30, 0xf3, 0xe6, 0xde, 0xe3, 0x6b, 0x8b,
	0x27, 0xbf, 0xa8, 0x7f, 0x46, 0x1f, 0x05, 0x0f, 0x2f, 0x56, 0xef, 0xd1, 0x4a, 0x18, 0x54, 0xa8,
	0x1a, 0x6f, 0x14, 0xbb, 0xee, 0x6a, 0xbc, 0x66, 0x35, 0xc4, 0x67, 0x5e, 0x70, 0x6b, 0x4c, 0xbf,
	0xb1, 0x89, 0x85, 0xdf, 0xbd, 0xe0, 0xe6, 0x7f, 0x63, 0x84, 0x2f, 0xd0, 0xae, 0x9c, 0xd1, 0xf6,
	0xc9, 0x23, 0x3c, 0xe9, 0x1c, 0xf2, 0x47, 0x63, 0xdf, 0x5e, 0xf2, 0x93, 0x6e, 0xee, 0xfa, 0xf9,
	0xc7, 0xef, 0x2f, 0xfe, 0xd3, 0x10, 0x77, 0x37, 0xfb, 0xb8, 0xc1, 0x1c, 0x16, 0x6b, 0xd3, 0xa8,
	0x1a, 0xb4, 0xc1, 0x07, 0xdd, 0x11, 0x87, 0x22, 0x06, 0x1f, 0x7c, 0x8a, 0xee, 0x9d, 0xd2, 0xd9,
	0x54, 0x2e, 0x8b, 0xbf, 0x5e, 0xb0, 0x5f, 0xa8, 0x7a, 0xe7, 0x5a, 0x8b, 0xd9, 0x96, 0x88, 0x8e,
	0xba, 0x6b, 0x1d, 0x79, 0x6f, 0x5f, 0xf5, 0x6e, 0xae, 0x2a, 0x26, 0x39, 0x52, 0x9a, 0x63, 0x0e,
	0xf2, 0xfc, 0x96, 0xf8, 0xa2, 0xdf, 0xf4, 0xd7, 0x7e, 0xe9, 0xc0, 0x57, 0xff, 0x4a, 0x42, 0xe9,
	0x37, 0x7f, 0x9e, 0xd8, 0x82, 0xb4, 0x34, 0xc8, 0xc2, 0x0e, 0xa5, 0x31, 0xea, 0x1b, 0x9b, 0x53,
	0x27, 0xc9, 0x68, 0x69, 0xb2, 0x41, 0x92, 0xa5, 0x71, 0xe6, 0x24, 0x7f, 0xfc, 0x7d, 0xfb, 0x4e,
	0x08, 0x2d, 0x0d, 0x21, 0x83, 0x88, 0x90, 0x34, 0x26, 0xc4, 0xc9, 0x8e, 0xaf, 0x9e, 0xcf, 0xf9,
	0xec, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xee, 0xc4, 0x25, 0x81, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgeRangeViewServiceClient is the client API for AgeRangeViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgeRangeViewServiceClient interface {
	// Returns the requested age range view in full detail.
	GetAgeRangeView(ctx context.Context, in *GetAgeRangeViewRequest, opts ...grpc.CallOption) (*resources.AgeRangeView, error)
}

type ageRangeViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgeRangeViewServiceClient(cc grpc.ClientConnInterface) AgeRangeViewServiceClient {
	return &ageRangeViewServiceClient{cc}
}

func (c *ageRangeViewServiceClient) GetAgeRangeView(ctx context.Context, in *GetAgeRangeViewRequest, opts ...grpc.CallOption) (*resources.AgeRangeView, error) {
	out := new(resources.AgeRangeView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AgeRangeViewService/GetAgeRangeView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgeRangeViewServiceServer is the server API for AgeRangeViewService service.
type AgeRangeViewServiceServer interface {
	// Returns the requested age range view in full detail.
	GetAgeRangeView(context.Context, *GetAgeRangeViewRequest) (*resources.AgeRangeView, error)
}

// UnimplementedAgeRangeViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgeRangeViewServiceServer struct {
}

func (*UnimplementedAgeRangeViewServiceServer) GetAgeRangeView(ctx context.Context, req *GetAgeRangeViewRequest) (*resources.AgeRangeView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgeRangeView not implemented")
}

func RegisterAgeRangeViewServiceServer(s *grpc.Server, srv AgeRangeViewServiceServer) {
	s.RegisterService(&_AgeRangeViewService_serviceDesc, srv)
}

func _AgeRangeViewService_GetAgeRangeView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgeRangeViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgeRangeViewServiceServer).GetAgeRangeView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AgeRangeViewService/GetAgeRangeView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgeRangeViewServiceServer).GetAgeRangeView(ctx, req.(*GetAgeRangeViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgeRangeViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AgeRangeViewService",
	HandlerType: (*AgeRangeViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgeRangeView",
			Handler:    _AgeRangeViewService_GetAgeRangeView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/age_range_view_service.proto",
}
