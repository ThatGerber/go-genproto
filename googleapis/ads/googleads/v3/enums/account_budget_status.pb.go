// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/account_budget_status.proto

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

// The possible statuses of an AccountBudget.
type AccountBudgetStatusEnum_AccountBudgetStatus int32

const (
	// Not specified.
	AccountBudgetStatusEnum_UNSPECIFIED AccountBudgetStatusEnum_AccountBudgetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	AccountBudgetStatusEnum_UNKNOWN AccountBudgetStatusEnum_AccountBudgetStatus = 1
	// The account budget is pending approval.
	AccountBudgetStatusEnum_PENDING AccountBudgetStatusEnum_AccountBudgetStatus = 2
	// The account budget has been approved.
	AccountBudgetStatusEnum_APPROVED AccountBudgetStatusEnum_AccountBudgetStatus = 3
	// The account budget has been cancelled by the user.
	AccountBudgetStatusEnum_CANCELLED AccountBudgetStatusEnum_AccountBudgetStatus = 4
)

var AccountBudgetStatusEnum_AccountBudgetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "APPROVED",
	4: "CANCELLED",
}

var AccountBudgetStatusEnum_AccountBudgetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PENDING":     2,
	"APPROVED":    3,
	"CANCELLED":   4,
}

func (x AccountBudgetStatusEnum_AccountBudgetStatus) String() string {
	return proto.EnumName(AccountBudgetStatusEnum_AccountBudgetStatus_name, int32(x))
}

func (AccountBudgetStatusEnum_AccountBudgetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b869fdf507ee832, []int{0, 0}
}

// Message describing AccountBudget statuses.
type AccountBudgetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBudgetStatusEnum) Reset()         { *m = AccountBudgetStatusEnum{} }
func (m *AccountBudgetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetStatusEnum) ProtoMessage()    {}
func (*AccountBudgetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b869fdf507ee832, []int{0}
}

func (m *AccountBudgetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetStatusEnum.Unmarshal(m, b)
}
func (m *AccountBudgetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetStatusEnum.Marshal(b, m, deterministic)
}
func (m *AccountBudgetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetStatusEnum.Merge(m, src)
}
func (m *AccountBudgetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetStatusEnum.Size(m)
}
func (m *AccountBudgetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AccountBudgetStatusEnum_AccountBudgetStatus", AccountBudgetStatusEnum_AccountBudgetStatus_name, AccountBudgetStatusEnum_AccountBudgetStatus_value)
	proto.RegisterType((*AccountBudgetStatusEnum)(nil), "google.ads.googleads.v3.enums.AccountBudgetStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/account_budget_status.proto", fileDescriptor_7b869fdf507ee832)
}

var fileDescriptor_7b869fdf507ee832 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xfb, 0x30,
	0x00, 0xfe, 0xad, 0xfb, 0xe1, 0x9f, 0x4c, 0xb1, 0xd4, 0x83, 0x22, 0xee, 0xb0, 0x3d, 0x40, 0x7a,
	0xe8, 0xc9, 0x78, 0x4a, 0xd7, 0x38, 0x86, 0x23, 0x2b, 0x8e, 0x55, 0x90, 0xca, 0xc8, 0x96, 0x12,
	0x06, 0x5b, 0x32, 0x96, 0x74, 0xf8, 0x3c, 0x1e, 0x7d, 0x14, 0x1f, 0x45, 0xf0, 0x1d, 0xa4, 0xc9,
	0xd6, 0xd3, 0xf4, 0x12, 0xbe, 0xe4, 0xfb, 0xc3, 0x97, 0x0f, 0xdc, 0x09, 0xa5, 0xc4, 0xb2, 0x08,
	0x19, 0xd7, 0xa1, 0x83, 0x15, 0xda, 0x46, 0x61, 0x21, 0xcb, 0x95, 0x0e, 0xd9, 0x7c, 0xae, 0x4a,
	0x69, 0xa6, 0xb3, 0x92, 0x8b, 0xc2, 0x4c, 0xb5, 0x61, 0xa6, 0xd4, 0x70, 0xbd, 0x51, 0x46, 0x05,
	0x6d, 0xa7, 0x87, 0x8c, 0x6b, 0x58, 0x5b, 0xe1, 0x36, 0x82, 0xd6, 0x7a, 0x73, 0xbb, 0x4f, 0x5e,
	0x2f, 0x42, 0x26, 0xa5, 0x32, 0xcc, 0x2c, 0x94, 0xdc, 0x99, 0xbb, 0x6f, 0xe0, 0x0a, 0xbb, 0xec,
	0xd8, 0x46, 0x8f, 0x6d, 0x32, 0x91, 0xe5, 0xaa, 0xfb, 0x0a, 0x2e, 0x0f, 0x50, 0xc1, 0x05, 0x68,
	0x4d, 0xe8, 0x38, 0x25, 0xbd, 0xc1, 0xc3, 0x80, 0x24, 0xfe, 0xbf, 0xa0, 0x05, 0x8e, 0x27, 0xf4,
	0x91, 0x8e, 0x9e, 0xa9, 0xdf, 0xa8, 0x2e, 0x29, 0xa1, 0xc9, 0x80, 0xf6, 0x7d, 0x2f, 0x38, 0x03,
	0x27, 0x38, 0x4d, 0x9f, 0x46, 0x19, 0x49, 0xfc, 0x66, 0x70, 0x0e, 0x4e, 0x7b, 0x98, 0xf6, 0xc8,
	0x70, 0x48, 0x12, 0xff, 0x7f, 0xfc, 0xdd, 0x00, 0x9d, 0xb9, 0x5a, 0xc1, 0x3f, 0xdb, 0xc7, 0xd7,
	0x07, 0x2a, 0xa4, 0x55, 0xf3, 0xb4, 0xf1, 0x12, 0xef, 0xac, 0x42, 0x2d, 0x99, 0x14, 0x50, 0x6d,
	0x44, 0x28, 0x0a, 0x69, 0xff, 0xb5, 0xdf, 0x70, 0xbd, 0xd0, 0xbf, 0x4c, 0x7a, 0x6f, 0xcf, 0x77,
	0xaf, 0xd9, 0xc7, 0xf8, 0xc3, 0x6b, 0xf7, 0x5d, 0x14, 0xe6, 0x1a, 0x3a, 0x58, 0xa1, 0x2c, 0x82,
	0xd5, 0x12, 0xfa, 0x73, 0xcf, 0xe7, 0x98, 0xeb, 0xbc, 0xe6, 0xf3, 0x2c, 0xca, 0x2d, 0xff, 0xe5,
	0x75, 0xdc, 0x23, 0x42, 0x98, 0x6b, 0x84, 0x6a, 0x05, 0x42, 0x59, 0x84, 0x90, 0xd5, 0xcc, 0x8e,
	0x6c, 0xb1, 0xe8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x17, 0x24, 0xba, 0xc5, 0xea, 0x01, 0x00, 0x00,
}