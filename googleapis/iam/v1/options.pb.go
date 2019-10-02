// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/v1/options.proto

package iam

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
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

// Encapsulates settings provided to GetIamPolicy.
type GetPolicyOptions struct {
	// Optional. The policy format version to be returned.
	//
	// Valid values are 0, 1, and 3. Requests specifying an invalid value will be
	// rejected.
	//
	// Requests for policies with any conditional bindings must specify version 3.
	// Policies without any conditional bindings may specify any valid value or
	// leave the field unset.
	RequestedPolicyVersion int32    `protobuf:"varint,1,opt,name=requested_policy_version,json=requestedPolicyVersion,proto3" json:"requested_policy_version,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GetPolicyOptions) Reset()         { *m = GetPolicyOptions{} }
func (m *GetPolicyOptions) String() string { return proto.CompactTextString(m) }
func (*GetPolicyOptions) ProtoMessage()    {}
func (*GetPolicyOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa09e909092bd1, []int{0}
}

func (m *GetPolicyOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyOptions.Unmarshal(m, b)
}
func (m *GetPolicyOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyOptions.Marshal(b, m, deterministic)
}
func (m *GetPolicyOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyOptions.Merge(m, src)
}
func (m *GetPolicyOptions) XXX_Size() int {
	return xxx_messageInfo_GetPolicyOptions.Size(m)
}
func (m *GetPolicyOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyOptions proto.InternalMessageInfo

func (m *GetPolicyOptions) GetRequestedPolicyVersion() int32 {
	if m != nil {
		return m.RequestedPolicyVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*GetPolicyOptions)(nil), "google.iam.v1.GetPolicyOptions")
}

func init() { proto.RegisterFile("google/iam/v1/options.proto", fileDescriptor_19aa09e909092bd1) }

var fileDescriptor_19aa09e909092bd1 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0x4c, 0xcc, 0xd5, 0x2f, 0x33, 0xd4, 0xcf, 0x2f, 0x28, 0xc9, 0xcc, 0xcf,
	0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0x48, 0xea, 0x65, 0x26, 0xe6, 0xea,
	0x95, 0x19, 0x4a, 0xc9, 0x40, 0xd5, 0x26, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24,
	0x22, 0x29, 0x56, 0xf2, 0xe1, 0x12, 0x70, 0x4f, 0x2d, 0x09, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0xf4,
	0x87, 0x18, 0x23, 0x64, 0xc1, 0x25, 0x51, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x92, 0x9a, 0x12,
	0x5f, 0x00, 0x96, 0x8a, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0d, 0x12, 0x83, 0xcb, 0x43, 0x74, 0x86, 0x41, 0x64, 0x9d, 0x5a, 0x18, 0xb9, 0x04, 0x93,
	0xf3, 0x73, 0xf5, 0x50, 0x5c, 0xe0, 0xc4, 0x03, 0x35, 0x38, 0x00, 0x64, 0x63, 0x00, 0x63, 0x94,
	0x01, 0x54, 0x3a, 0x3d, 0x3f, 0x27, 0x31, 0x2f, 0x5d, 0x2f, 0xbf, 0x28, 0x5d, 0x3f, 0x3d, 0x35,
	0x0f, 0xec, 0x1e, 0x7d, 0x88, 0x54, 0x62, 0x41, 0x66, 0x31, 0xd4, 0x73, 0xd6, 0x99, 0x89, 0xb9,
	0x3f, 0x18, 0x19, 0x57, 0x31, 0x09, 0xbb, 0x43, 0x74, 0x39, 0xe7, 0xe4, 0x97, 0xa6, 0xe8, 0x79,
	0x26, 0xe6, 0xea, 0x85, 0x19, 0x9e, 0x82, 0x89, 0xc6, 0x80, 0x45, 0x63, 0x3c, 0x13, 0x73, 0x63,
	0xc2, 0x0c, 0x93, 0xd8, 0xc0, 0x66, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x96, 0x0c,
	0x8b, 0x27, 0x01, 0x00, 0x00,
}