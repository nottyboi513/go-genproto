// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/maps/routes/v1/waypoint.proto

package routes

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// Encapsulates a waypoint. Waypoints mark both the beginning and end of a
// route, and include intermediate stops along the route.
type Waypoint struct {
	// Different ways to represent a location.
	//
	// Types that are valid to be assigned to LocationType:
	//	*Waypoint_Location
	//	*Waypoint_PlaceId
	LocationType isWaypoint_LocationType `protobuf_oneof:"location_type"`
	// Marks this waypoint as a milestone, as opposed to a stopping point. For
	// each non-via waypoint in the request, the response appends an entry to the
	// `legs` array to provide the details for stopovers on that leg of the
	// trip. Set this value to true when you want the route to pass through this
	// waypoint without stopping over. Via waypoints don't cause an entry to be
	// added to the `legs` array, but they do route the journey through the
	// waypoint. You can only set this value on waypoints that are intermediates.
	// If you set this field on terminal waypoints, then the request fails.
	Via bool `protobuf:"varint,3,opt,name=via,proto3" json:"via,omitempty"`
	// Indicates that the waypoint is meant for vehicles to stop at, where the
	// intention is to either pickup or drop-off. When you set this value, the
	// calculated route won't include non-`via` waypoints on roads that are
	// unsuitable for pickup and drop-off. This option works only for `DRIVE` and
	// `TWO_WHEELER` travel modes, and when the `location_type` is `location`.
	VehicleStopover bool `protobuf:"varint,4,opt,name=vehicle_stopover,json=vehicleStopover,proto3" json:"vehicle_stopover,omitempty"`
	// Indicates that the location of this waypoint is meant to have a preference
	// for the vehicle to stop at a particular side of road. When you set this
	// value, the route will pass through the location so that the vehicle can
	// stop at the side of road that the location is biased towards from the
	// center of the road. This option works only for 'DRIVE' and 'TWO_WHEELER'
	// travel modes, and when the 'location_type' is set to 'location'.
	SideOfRoad           bool     `protobuf:"varint,5,opt,name=side_of_road,json=sideOfRoad,proto3" json:"side_of_road,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Waypoint) Reset()         { *m = Waypoint{} }
func (m *Waypoint) String() string { return proto.CompactTextString(m) }
func (*Waypoint) ProtoMessage()    {}
func (*Waypoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a74b9fab6220a4a, []int{0}
}

func (m *Waypoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Waypoint.Unmarshal(m, b)
}
func (m *Waypoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Waypoint.Marshal(b, m, deterministic)
}
func (m *Waypoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Waypoint.Merge(m, src)
}
func (m *Waypoint) XXX_Size() int {
	return xxx_messageInfo_Waypoint.Size(m)
}
func (m *Waypoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Waypoint.DiscardUnknown(m)
}

var xxx_messageInfo_Waypoint proto.InternalMessageInfo

type isWaypoint_LocationType interface {
	isWaypoint_LocationType()
}

type Waypoint_Location struct {
	Location *Location `protobuf:"bytes,1,opt,name=location,proto3,oneof"`
}

type Waypoint_PlaceId struct {
	PlaceId string `protobuf:"bytes,2,opt,name=place_id,json=placeId,proto3,oneof"`
}

func (*Waypoint_Location) isWaypoint_LocationType() {}

func (*Waypoint_PlaceId) isWaypoint_LocationType() {}

func (m *Waypoint) GetLocationType() isWaypoint_LocationType {
	if m != nil {
		return m.LocationType
	}
	return nil
}

func (m *Waypoint) GetLocation() *Location {
	if x, ok := m.GetLocationType().(*Waypoint_Location); ok {
		return x.Location
	}
	return nil
}

func (m *Waypoint) GetPlaceId() string {
	if x, ok := m.GetLocationType().(*Waypoint_PlaceId); ok {
		return x.PlaceId
	}
	return ""
}

func (m *Waypoint) GetVia() bool {
	if m != nil {
		return m.Via
	}
	return false
}

func (m *Waypoint) GetVehicleStopover() bool {
	if m != nil {
		return m.VehicleStopover
	}
	return false
}

func (m *Waypoint) GetSideOfRoad() bool {
	if m != nil {
		return m.SideOfRoad
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Waypoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Waypoint_Location)(nil),
		(*Waypoint_PlaceId)(nil),
	}
}

// Encapsulates a location (a geographic point, and an optional heading).
type Location struct {
	// The waypoint's geographic coordinates.
	LatLng *latlng.LatLng `protobuf:"bytes,1,opt,name=lat_lng,json=latLng,proto3" json:"lat_lng,omitempty"`
	// The compass heading associated with the direction of the flow of traffic.
	// This value is used to specify the side of the road to use for pickup and
	// drop-off. Heading values can be from 0 to 360, where 0 specifies a heading
	// of due North, 90 specifies a heading of due East, etc. You can use this
	// field only for `DRIVE` and `TWO_WHEELER` travel modes.
	Heading              *wrappers.Int32Value `protobuf:"bytes,2,opt,name=heading,proto3" json:"heading,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a74b9fab6220a4a, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatLng() *latlng.LatLng {
	if m != nil {
		return m.LatLng
	}
	return nil
}

func (m *Location) GetHeading() *wrappers.Int32Value {
	if m != nil {
		return m.Heading
	}
	return nil
}

func init() {
	proto.RegisterType((*Waypoint)(nil), "google.maps.routes.v1.Waypoint")
	proto.RegisterType((*Location)(nil), "google.maps.routes.v1.Location")
}

func init() {
	proto.RegisterFile("google/maps/routes/v1/waypoint.proto", fileDescriptor_5a74b9fab6220a4a)
}

var fileDescriptor_5a74b9fab6220a4a = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4b, 0x1c, 0x31,
	0x18, 0xc6, 0xcd, 0x6a, 0xdd, 0x69, 0xac, 0x28, 0x53, 0x84, 0xa9, 0x42, 0xbb, 0x48, 0x0f, 0x5b,
	0x28, 0x09, 0xbb, 0xd2, 0x93, 0xf4, 0xb2, 0x17, 0x15, 0x56, 0x2a, 0x11, 0xb6, 0x50, 0x16, 0x86,
	0xb8, 0x93, 0x8d, 0x81, 0x98, 0x37, 0xcc, 0x64, 0x47, 0xfc, 0x3a, 0xed, 0xad, 0x1f, 0xa5, 0x1f,
	0xa0, 0x9f, 0xa5, 0xc7, 0x92, 0x7f, 0x3d, 0x14, 0x4f, 0x13, 0x9e, 0xe7, 0xf7, 0xe6, 0xc9, 0xf3,
	0x32, 0xf8, 0xbd, 0x04, 0x90, 0x5a, 0xd0, 0x07, 0x6e, 0x3b, 0xda, 0xc2, 0xc6, 0x89, 0x8e, 0xf6,
	0x13, 0xfa, 0xc8, 0x9f, 0x2c, 0x28, 0xe3, 0x88, 0x6d, 0xc1, 0x41, 0x79, 0x14, 0x29, 0xe2, 0x29,
	0x12, 0x29, 0xd2, 0x4f, 0x8e, 0xdf, 0xa6, 0xe1, 0x00, 0xdd, 0x6d, 0xd6, 0xf4, 0xb1, 0xe5, 0xd6,
	0x8a, 0xb6, 0x8b, 0x63, 0xc7, 0x55, 0xf2, 0xdd, 0x93, 0x15, 0x54, 0x73, 0xa7, 0x8d, 0x8c, 0xce,
	0xe9, 0x6f, 0x84, 0x8b, 0xaf, 0x29, 0xa3, 0xfc, 0x8c, 0x0b, 0x0d, 0x2b, 0xee, 0x14, 0x98, 0x0a,
	0x8d, 0xd0, 0x78, 0x6f, 0xfa, 0x8e, 0x3c, 0x1b, 0x48, 0xe6, 0x09, 0xbb, 0xdc, 0x62, 0xff, 0x46,
	0xca, 0x13, 0x5c, 0x58, 0xcd, 0x57, 0xa2, 0x56, 0x4d, 0x35, 0x18, 0xa1, 0xf1, 0xcb, 0xcb, 0x2d,
	0x36, 0x0c, 0xca, 0x55, 0x53, 0x1e, 0xe2, 0xed, 0x5e, 0xf1, 0x6a, 0x7b, 0x84, 0xc6, 0x05, 0xf3,
	0xc7, 0xf2, 0x03, 0x3e, 0xec, 0xc5, 0xbd, 0x5a, 0x69, 0x51, 0x77, 0x0e, 0x2c, 0xf4, 0xa2, 0xad,
	0x76, 0x82, 0x7d, 0x90, 0xf4, 0xdb, 0x24, 0x97, 0x23, 0xfc, 0xaa, 0x53, 0x8d, 0xa8, 0x61, 0x5d,
	0xb7, 0xc0, 0x9b, 0xea, 0x45, 0xc0, 0xb0, 0xd7, 0xbe, 0xac, 0x19, 0xf0, 0x66, 0x76, 0x80, 0xf7,
	0xf3, 0x3b, 0x6a, 0xdf, 0xf2, 0x14, 0x70, 0x91, 0x1f, 0x59, 0x7e, 0xc4, 0x43, 0xcd, 0x5d, 0xad,
	0x8d, 0x4c, 0xb5, 0x5e, 0xe7, 0x5a, 0x1e, 0x25, 0x73, 0xee, 0xe6, 0x46, 0xb2, 0x5d, 0x1d, 0xbe,
	0xe5, 0x27, 0x3c, 0xbc, 0x17, 0xbc, 0x51, 0x46, 0x86, 0x16, 0x7b, 0xd3, 0x93, 0x4c, 0xe7, 0xf5,
	0x92, 0x2b, 0xe3, 0xce, 0xa6, 0x0b, 0xae, 0x37, 0x82, 0x65, 0x76, 0xf6, 0x03, 0xe1, 0x37, 0x2b,
	0x78, 0x78, 0x7e, 0x61, 0xb3, 0xfd, 0xbc, 0xe4, 0x1b, 0x7f, 0xc7, 0x0d, 0xfa, 0x76, 0x9e, 0x38,
	0x09, 0x9a, 0x1b, 0x49, 0xa0, 0x95, 0x54, 0x0a, 0x13, 0x12, 0x68, 0xb4, 0xb8, 0x55, 0xdd, 0x7f,
	0xbf, 0xc3, 0x79, 0x3c, 0xfd, 0x41, 0xe8, 0xfb, 0x60, 0xe7, 0xe2, 0x9a, 0xdd, 0xfe, 0x1c, 0x1c,
	0x5d, 0xc4, 0x7b, 0xae, 0x7d, 0x1e, 0x8b, 0x79, 0x8b, 0xc9, 0xaf, 0xac, 0x2f, 0xbd, 0xbe, 0x8c,
	0xfa, 0x72, 0x31, 0xb9, 0xdb, 0x0d, 0x09, 0x67, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x00, 0x98,
	0xc0, 0x83, 0x6f, 0x02, 0x00, 0x00,
}
