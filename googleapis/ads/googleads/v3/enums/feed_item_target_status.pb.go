// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/feed_item_target_status.proto

package enums

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

// Possible statuses of a feed item target.
type FeedItemTargetStatusEnum_FeedItemTargetStatus int32

const (
	// Not specified.
	FeedItemTargetStatusEnum_UNSPECIFIED FeedItemTargetStatusEnum_FeedItemTargetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemTargetStatusEnum_UNKNOWN FeedItemTargetStatusEnum_FeedItemTargetStatus = 1
	// Feed item target is enabled.
	FeedItemTargetStatusEnum_ENABLED FeedItemTargetStatusEnum_FeedItemTargetStatus = 2
	// Feed item target has been removed.
	FeedItemTargetStatusEnum_REMOVED FeedItemTargetStatusEnum_FeedItemTargetStatus = 3
)

var FeedItemTargetStatusEnum_FeedItemTargetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var FeedItemTargetStatusEnum_FeedItemTargetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x FeedItemTargetStatusEnum_FeedItemTargetStatus) String() string {
	return proto.EnumName(FeedItemTargetStatusEnum_FeedItemTargetStatus_name, int32(x))
}

func (FeedItemTargetStatusEnum_FeedItemTargetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b220975402148c96, []int{0, 0}
}

// Container for enum describing possible statuses of a feed item target.
type FeedItemTargetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemTargetStatusEnum) Reset()         { *m = FeedItemTargetStatusEnum{} }
func (m *FeedItemTargetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemTargetStatusEnum) ProtoMessage()    {}
func (*FeedItemTargetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b220975402148c96, []int{0}
}

func (m *FeedItemTargetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemTargetStatusEnum.Unmarshal(m, b)
}
func (m *FeedItemTargetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemTargetStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemTargetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemTargetStatusEnum.Merge(m, src)
}
func (m *FeedItemTargetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemTargetStatusEnum.Size(m)
}
func (m *FeedItemTargetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemTargetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemTargetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.FeedItemTargetStatusEnum_FeedItemTargetStatus", FeedItemTargetStatusEnum_FeedItemTargetStatus_name, FeedItemTargetStatusEnum_FeedItemTargetStatus_value)
	proto.RegisterType((*FeedItemTargetStatusEnum)(nil), "google.ads.googleads.v3.enums.FeedItemTargetStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/feed_item_target_status.proto", fileDescriptor_b220975402148c96)
}

var fileDescriptor_b220975402148c96 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x4e, 0x2a, 0x31,
	0x14, 0x7d, 0x0c, 0xc9, 0x33, 0x29, 0x0b, 0x27, 0xc4, 0x85, 0x1a, 0x59, 0xc0, 0x07, 0xb4, 0x8b,
	0xee, 0xca, 0xaa, 0x23, 0x85, 0x10, 0xb5, 0x10, 0x91, 0x31, 0x31, 0x93, 0x90, 0x6a, 0x6b, 0x33,
	0x86, 0x69, 0x09, 0x2d, 0x7c, 0x90, 0x4b, 0x3f, 0xc5, 0x4f, 0x71, 0xe1, 0x37, 0x98, 0x76, 0x80,
	0x15, 0xba, 0x69, 0xce, 0xed, 0xb9, 0xe7, 0xdc, 0x73, 0x2f, 0xe8, 0x6b, 0x6b, 0xf5, 0x52, 0x21,
	0x21, 0x1d, 0xaa, 0x61, 0x40, 0x5b, 0x8c, 0x94, 0xd9, 0x54, 0x0e, 0xbd, 0x2a, 0x25, 0x17, 0xa5,
	0x57, 0xd5, 0xc2, 0x8b, 0xb5, 0x56, 0x7e, 0xe1, 0xbc, 0xf0, 0x1b, 0x07, 0x57, 0x6b, 0xeb, 0x6d,
	0xbb, 0x53, 0x2b, 0xa0, 0x90, 0x0e, 0x1e, 0xc4, 0x70, 0x8b, 0x61, 0x14, 0x5f, 0x5e, 0xed, 0xbd,
	0x57, 0x25, 0x12, 0xc6, 0x58, 0x2f, 0x7c, 0x69, 0xcd, 0x4e, 0xdc, 0x7b, 0x03, 0xe7, 0x43, 0xa5,
	0xe4, 0xd8, 0xab, 0xea, 0x21, 0x7a, 0xcf, 0xa2, 0x35, 0x33, 0x9b, 0xaa, 0xc7, 0xc1, 0xd9, 0x31,
	0xae, 0x7d, 0x0a, 0x5a, 0x73, 0x3e, 0x9b, 0xb2, 0xeb, 0xf1, 0x70, 0xcc, 0x06, 0xe9, 0xbf, 0x76,
	0x0b, 0x9c, 0xcc, 0xf9, 0x0d, 0x9f, 0x3c, 0xf2, 0xb4, 0x11, 0x0a, 0xc6, 0x69, 0x76, 0xcb, 0x06,
	0x69, 0x12, 0x8a, 0x7b, 0x76, 0x37, 0xc9, 0xd9, 0x20, 0x6d, 0x66, 0xdf, 0x0d, 0xd0, 0x7d, 0xb1,
	0x15, 0xfc, 0x33, 0x6f, 0x76, 0x71, 0x6c, 0xe6, 0x34, 0x84, 0x9d, 0x36, 0x9e, 0xb2, 0x9d, 0x56,
	0xdb, 0xa5, 0x30, 0x1a, 0xda, 0xb5, 0x46, 0x5a, 0x99, 0xb8, 0xca, 0xfe, 0x70, 0xab, 0xd2, 0xfd,
	0x72, 0xc7, 0x7e, 0x7c, 0xdf, 0x93, 0xe6, 0x88, 0xd2, 0x8f, 0xa4, 0x33, 0xaa, 0xad, 0xa8, 0x74,
	0xb0, 0x86, 0x01, 0xe5, 0x18, 0x86, 0xdd, 0xdd, 0xe7, 0x9e, 0x2f, 0xa8, 0x74, 0xc5, 0x81, 0x2f,
	0x72, 0x5c, 0x44, 0xfe, 0x2b, 0xe9, 0xd6, 0x9f, 0x84, 0x50, 0xe9, 0x08, 0x39, 0x74, 0x10, 0x92,
	0x63, 0x42, 0x62, 0xcf, 0xf3, 0xff, 0x18, 0x0c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x8a,
	0x68, 0x18, 0xdf, 0x01, 0x00, 0x00,
}