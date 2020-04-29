// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/hotel_group_view.proto

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

// A hotel group view.
type HotelGroupView struct {
	// Output only. The resource name of the hotel group view.
	// Hotel Group view resource names have the form:
	//
	// `customers/{customer_id}/hotelGroupViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelGroupView) Reset()         { *m = HotelGroupView{} }
func (m *HotelGroupView) String() string { return proto.CompactTextString(m) }
func (*HotelGroupView) ProtoMessage()    {}
func (*HotelGroupView) Descriptor() ([]byte, []int) {
	return fileDescriptor_e471d0eff5bb8f4a, []int{0}
}

func (m *HotelGroupView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelGroupView.Unmarshal(m, b)
}
func (m *HotelGroupView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelGroupView.Marshal(b, m, deterministic)
}
func (m *HotelGroupView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelGroupView.Merge(m, src)
}
func (m *HotelGroupView) XXX_Size() int {
	return xxx_messageInfo_HotelGroupView.Size(m)
}
func (m *HotelGroupView) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelGroupView.DiscardUnknown(m)
}

var xxx_messageInfo_HotelGroupView proto.InternalMessageInfo

func (m *HotelGroupView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*HotelGroupView)(nil), "google.ads.googleads.v3.resources.HotelGroupView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/hotel_group_view.proto", fileDescriptor_e471d0eff5bb8f4a)
}

var fileDescriptor_e471d0eff5bb8f4a = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x18, 0x85, 0x49, 0x0a, 0x17, 0x6e, 0xb8, 0xd7, 0x45, 0xdd, 0x68, 0x11, 0xb4, 0x42, 0x51, 0x37,
	0x33, 0x8b, 0x2c, 0x94, 0x71, 0x35, 0xdd, 0x54, 0x5c, 0x48, 0x29, 0x92, 0x85, 0x04, 0xc2, 0x34,
	0x19, 0xd3, 0x81, 0x24, 0x7f, 0x98, 0x49, 0xd3, 0x45, 0x29, 0xf8, 0x2c, 0x2e, 0x7d, 0x14, 0xc1,
	0x77, 0x70, 0xdd, 0x47, 0x70, 0x25, 0xe9, 0x74, 0xa6, 0xad, 0x0b, 0x75, 0x77, 0xe0, 0xff, 0xce,
	0x99, 0xc3, 0x19, 0xef, 0x2a, 0x05, 0x48, 0x33, 0x8e, 0x59, 0xa2, 0xb0, 0x96, 0x8d, 0xaa, 0x7d,
	0x2c, 0xb9, 0x82, 0xa9, 0x8c, 0xb9, 0xc2, 0x13, 0xa8, 0x78, 0x16, 0xa5, 0x12, 0xa6, 0x65, 0x54,
	0x0b, 0x3e, 0x43, 0xa5, 0x84, 0x0a, 0xda, 0x5d, 0x8d, 0x23, 0x96, 0x28, 0x64, 0x9d, 0xa8, 0xf6,
	0x91, 0x75, 0x76, 0x8e, 0x4d, 0x78, 0x29, 0xf0, 0xa3, 0xe0, 0x59, 0x12, 0x8d, 0xf9, 0x84, 0xd5,
	0x02, 0xa4, 0xce, 0xe8, 0x1c, 0x6e, 0x01, 0xc6, 0xb6, 0x3e, 0x1d, 0x6d, 0x9d, 0x58, 0x51, 0x40,
	0xc5, 0x2a, 0x01, 0x85, 0xd2, 0xd7, 0xd3, 0x37, 0xc7, 0xdb, 0xbb, 0x69, 0x7a, 0x0d, 0x9a, 0x5a,
	0x81, 0xe0, 0xb3, 0xf6, 0xbd, 0xf7, 0xdf, 0x44, 0x44, 0x05, 0xcb, 0xf9, 0x81, 0x73, 0xe2, 0x9c,
	0xff, 0xed, 0xe3, 0x77, 0xda, 0xfa, 0xa0, 0x17, 0xde, 0xd9, 0xa6, 0xe3, 0x5a, 0x95, 0x42, 0xa1,
	0x18, 0x72, 0xbc, 0x9b, 0x33, 0xfa, 0x67, 0x52, 0xee, 0x58, 0xce, 0x09, 0x5f, 0xd2, 0xf1, 0xaf,
	0xbd, 0xed, 0xcb, 0x78, 0xaa, 0x2a, 0xc8, 0xb9, 0x54, 0x78, 0x6e, 0xe4, 0x42, 0x0f, 0x68, 0x21,
	0x85, 0xe7, 0x5f, 0x17, 0x5d, 0xf4, 0x9f, 0x5c, 0xaf, 0x17, 0x43, 0x8e, 0x7e, 0xdc, 0xb4, 0xbf,
	0xbf, 0xfb, 0xe4, 0xb0, 0x99, 0x63, 0xe8, 0x3c, 0xdc, 0xae, 0x9d, 0x29, 0x64, 0xac, 0x48, 0x11,
	0xc8, 0x14, 0xa7, 0xbc, 0x58, 0x8d, 0x85, 0x37, 0x9d, 0xbf, 0xf9, 0xe6, 0x6b, 0xab, 0x9e, 0xdd,
	0xd6, 0x80, 0xd2, 0x17, 0xb7, 0x3b, 0xd0, 0x91, 0x34, 0x51, 0x48, 0xcb, 0x46, 0x05, 0x3e, 0x1a,
	0x19, 0xf2, 0xd5, 0x30, 0x21, 0x4d, 0x54, 0x68, 0x99, 0x30, 0xf0, 0x43, 0xcb, 0x2c, 0xdd, 0x9e,
	0x3e, 0x10, 0x42, 0x13, 0x45, 0x88, 0xa5, 0x08, 0x09, 0x7c, 0x42, 0x2c, 0x37, 0xfe, 0xb3, 0x2a,
	0xeb, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xae, 0xf1, 0x86, 0x0a, 0x92, 0x02, 0x00, 0x00,
}