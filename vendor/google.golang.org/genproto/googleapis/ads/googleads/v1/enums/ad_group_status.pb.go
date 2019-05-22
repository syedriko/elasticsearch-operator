// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/ad_group_status.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The possible statuses of an ad group.
type AdGroupStatusEnum_AdGroupStatus int32

const (
	// The status has not been specified.
	AdGroupStatusEnum_UNSPECIFIED AdGroupStatusEnum_AdGroupStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupStatusEnum_UNKNOWN AdGroupStatusEnum_AdGroupStatus = 1
	// The ad group is enabled.
	AdGroupStatusEnum_ENABLED AdGroupStatusEnum_AdGroupStatus = 2
	// The ad group is paused.
	AdGroupStatusEnum_PAUSED AdGroupStatusEnum_AdGroupStatus = 3
	// The ad group is removed.
	AdGroupStatusEnum_REMOVED AdGroupStatusEnum_AdGroupStatus = 4
)

var AdGroupStatusEnum_AdGroupStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}
var AdGroupStatusEnum_AdGroupStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x AdGroupStatusEnum_AdGroupStatus) String() string {
	return proto.EnumName(AdGroupStatusEnum_AdGroupStatus_name, int32(x))
}
func (AdGroupStatusEnum_AdGroupStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_status_554f816bf5eafd8a, []int{0, 0}
}

// Container for enum describing possible statuses of an ad group.
type AdGroupStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupStatusEnum) Reset()         { *m = AdGroupStatusEnum{} }
func (m *AdGroupStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupStatusEnum) ProtoMessage()    {}
func (*AdGroupStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_status_554f816bf5eafd8a, []int{0}
}
func (m *AdGroupStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupStatusEnum.Marshal(b, m, deterministic)
}
func (dst *AdGroupStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupStatusEnum.Merge(dst, src)
}
func (m *AdGroupStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupStatusEnum.Size(m)
}
func (m *AdGroupStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupStatusEnum)(nil), "google.ads.googleads.v1.enums.AdGroupStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.AdGroupStatusEnum_AdGroupStatus", AdGroupStatusEnum_AdGroupStatus_name, AdGroupStatusEnum_AdGroupStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/ad_group_status.proto", fileDescriptor_ad_group_status_554f816bf5eafd8a)
}

var fileDescriptor_ad_group_status_554f816bf5eafd8a = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x77, 0x9d, 0x4c, 0xc8, 0x10, 0x6b, 0x8f, 0xe2, 0x0e, 0xdb, 0x03, 0x24, 0x94, 0xdd, 0xe2,
	0x29, 0xb5, 0x71, 0x0c, 0xb5, 0x2b, 0x96, 0x55, 0x90, 0xc2, 0x88, 0x66, 0xc4, 0xc2, 0x9a, 0x94,
	0xa6, 0xdd, 0x03, 0x79, 0xf4, 0x51, 0x7c, 0x11, 0xc1, 0xa7, 0x90, 0xa4, 0x6b, 0x61, 0x07, 0xbd,
	0x94, 0x5f, 0xbf, 0xdf, 0x9f, 0xfc, 0xbe, 0x0f, 0xcc, 0x85, 0x52, 0x62, 0xb7, 0x45, 0x8c, 0x6b,
	0xd4, 0x42, 0x83, 0xf6, 0x3e, 0xda, 0xca, 0xa6, 0xd0, 0x88, 0xf1, 0x8d, 0xa8, 0x54, 0x53, 0x6e,
	0x74, 0xcd, 0xea, 0x46, 0xc3, 0xb2, 0x52, 0xb5, 0xf2, 0x26, 0xad, 0x12, 0x32, 0xae, 0x61, 0x6f,
	0x82, 0x7b, 0x1f, 0x5a, 0xd3, 0xd5, 0x75, 0x97, 0x59, 0xe6, 0x88, 0x49, 0xa9, 0x6a, 0x56, 0xe7,
	0x4a, 0x1e, 0xcc, 0xb3, 0x77, 0x70, 0x49, 0xf8, 0xc2, 0x84, 0x26, 0x36, 0x93, 0xca, 0xa6, 0x98,
	0x25, 0xe0, 0xfc, 0x68, 0xe8, 0x5d, 0x80, 0xf1, 0x3a, 0x4a, 0x62, 0x7a, 0xbb, 0xbc, 0x5b, 0xd2,
	0xd0, 0x3d, 0xf1, 0xc6, 0xe0, 0x6c, 0x1d, 0xdd, 0x47, 0xab, 0xe7, 0xc8, 0x1d, 0x98, 0x1f, 0x1a,
	0x91, 0xe0, 0x81, 0x86, 0xae, 0xe3, 0x01, 0x30, 0x8a, 0xc9, 0x3a, 0xa1, 0xa1, 0x3b, 0x34, 0xc4,
	0x13, 0x7d, 0x5c, 0xa5, 0x34, 0x74, 0x4f, 0x83, 0xef, 0x01, 0x98, 0xbe, 0xa9, 0x02, 0xfe, 0xdb,
	0x36, 0xf0, 0x8e, 0x1e, 0x8e, 0x4d, 0xc7, 0x78, 0xf0, 0x12, 0x1c, 0x4c, 0x42, 0xed, 0x98, 0x14,
	0x50, 0x55, 0x02, 0x89, 0xad, 0xb4, 0x1b, 0x74, 0x77, 0x2a, 0x73, 0xfd, 0xc7, 0xd9, 0x6e, 0xec,
	0xf7, 0xc3, 0x19, 0x2e, 0x08, 0xf9, 0x74, 0x26, 0x8b, 0x36, 0x8a, 0x70, 0x0d, 0x5b, 0x68, 0x50,
	0xea, 0x43, 0xb3, 0xb9, 0xfe, 0xea, 0xf8, 0x8c, 0x70, 0x9d, 0xf5, 0x7c, 0x96, 0xfa, 0x99, 0xe5,
	0x7f, 0x9c, 0x69, 0x3b, 0xc4, 0x98, 0x70, 0x8d, 0x71, 0xaf, 0xc0, 0x38, 0xf5, 0x31, 0xb6, 0x9a,
	0xd7, 0x91, 0x2d, 0x36, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x23, 0x2a, 0x6d, 0x08, 0xce, 0x01,
	0x00, 0x00,
}