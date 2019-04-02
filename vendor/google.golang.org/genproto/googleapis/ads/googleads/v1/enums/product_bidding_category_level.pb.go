// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/product_bidding_category_level.proto

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

// Enum describing the level of the product bidding category.
type ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel int32

const (
	// Not specified.
	ProductBiddingCategoryLevelEnum_UNSPECIFIED ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel = 0
	// Used for return value only. Represents value unknown in this version.
	ProductBiddingCategoryLevelEnum_UNKNOWN ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel = 1
	// Level 1.
	ProductBiddingCategoryLevelEnum_LEVEL1 ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel = 2
	// Level 2.
	ProductBiddingCategoryLevelEnum_LEVEL2 ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel = 3
	// Level 3.
	ProductBiddingCategoryLevelEnum_LEVEL3 ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel = 4
	// Level 4.
	ProductBiddingCategoryLevelEnum_LEVEL4 ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel = 5
	// Level 5.
	ProductBiddingCategoryLevelEnum_LEVEL5 ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel = 6
)

var ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "LEVEL1",
	3: "LEVEL2",
	4: "LEVEL3",
	5: "LEVEL4",
	6: "LEVEL5",
}
var ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"LEVEL1":      2,
	"LEVEL2":      3,
	"LEVEL3":      4,
	"LEVEL4":      5,
	"LEVEL5":      6,
}

func (x ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel) String() string {
	return proto.EnumName(ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel_name, int32(x))
}
func (ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_product_bidding_category_level_18cbb4e8d9f3d6a7, []int{0, 0}
}

// Level of a product bidding category.
type ProductBiddingCategoryLevelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductBiddingCategoryLevelEnum) Reset()         { *m = ProductBiddingCategoryLevelEnum{} }
func (m *ProductBiddingCategoryLevelEnum) String() string { return proto.CompactTextString(m) }
func (*ProductBiddingCategoryLevelEnum) ProtoMessage()    {}
func (*ProductBiddingCategoryLevelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_bidding_category_level_18cbb4e8d9f3d6a7, []int{0}
}
func (m *ProductBiddingCategoryLevelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductBiddingCategoryLevelEnum.Unmarshal(m, b)
}
func (m *ProductBiddingCategoryLevelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductBiddingCategoryLevelEnum.Marshal(b, m, deterministic)
}
func (dst *ProductBiddingCategoryLevelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductBiddingCategoryLevelEnum.Merge(dst, src)
}
func (m *ProductBiddingCategoryLevelEnum) XXX_Size() int {
	return xxx_messageInfo_ProductBiddingCategoryLevelEnum.Size(m)
}
func (m *ProductBiddingCategoryLevelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductBiddingCategoryLevelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProductBiddingCategoryLevelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProductBiddingCategoryLevelEnum)(nil), "google.ads.googleads.v1.enums.ProductBiddingCategoryLevelEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel", ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel_name, ProductBiddingCategoryLevelEnum_ProductBiddingCategoryLevel_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/product_bidding_category_level.proto", fileDescriptor_product_bidding_category_level_18cbb4e8d9f3d6a7)
}

var fileDescriptor_product_bidding_category_level_18cbb4e8d9f3d6a7 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0xed, 0xfe, 0x4e, 0xc8, 0x0e, 0x86, 0x1e, 0xd5, 0xa1, 0xdb, 0x07, 0x48, 0xa9,
	0xd3, 0x4b, 0x3c, 0xb5, 0xb3, 0x8e, 0xe1, 0xa8, 0x05, 0x59, 0x05, 0x29, 0x8c, 0x6c, 0x09, 0xa1,
	0xd0, 0x25, 0xa5, 0xe9, 0x26, 0x7e, 0x15, 0x8f, 0x1e, 0xfd, 0x28, 0x7e, 0x14, 0xef, 0xde, 0xa5,
	0xc9, 0x96, 0x9b, 0xbb, 0x94, 0x1f, 0x7d, 0xdf, 0xf7, 0x79, 0xde, 0xf7, 0x09, 0x88, 0xb8, 0x94,
	0xbc, 0x64, 0x3e, 0xa1, 0xca, 0x37, 0xd8, 0xd2, 0x36, 0xf0, 0x99, 0xd8, 0xac, 0x95, 0x5f, 0xd5,
	0x92, 0x6e, 0x56, 0xcd, 0x62, 0x59, 0x50, 0x5a, 0x08, 0xbe, 0x58, 0x91, 0x86, 0x71, 0x59, 0xbf,
	0x2d, 0x4a, 0xb6, 0x65, 0x25, 0xaa, 0x6a, 0xd9, 0x48, 0xaf, 0x6f, 0x06, 0x11, 0xa1, 0x0a, 0x59,
	0x0d, 0xb4, 0x0d, 0x90, 0xd6, 0x38, 0x3d, 0xdf, 0x5b, 0x54, 0x85, 0x4f, 0x84, 0x90, 0x0d, 0x69,
	0x0a, 0x29, 0x94, 0x19, 0x1e, 0xbe, 0x3b, 0xe0, 0x22, 0x35, 0x2e, 0x91, 0x31, 0x19, 0xef, 0x3c,
	0x66, 0xad, 0x45, 0x2c, 0x36, 0xeb, 0xe1, 0x2b, 0x38, 0x3b, 0xd0, 0xe2, 0x9d, 0x80, 0xde, 0x3c,
	0x79, 0x4a, 0xe3, 0xf1, 0xf4, 0x7e, 0x1a, 0xdf, 0xc1, 0x7f, 0x5e, 0x0f, 0x1c, 0xcf, 0x93, 0x87,
	0xe4, 0xf1, 0x39, 0x81, 0x8e, 0x07, 0x40, 0x77, 0x16, 0x67, 0xf1, 0x2c, 0x80, 0xae, 0xe5, 0x2b,
	0xd8, 0xb1, 0x3c, 0x82, 0xff, 0x2d, 0x5f, 0xc3, 0x23, 0xcb, 0x37, 0xb0, 0x1b, 0xfd, 0x38, 0x60,
	0xb0, 0x92, 0x6b, 0x74, 0xf0, 0xc0, 0xe8, 0xf2, 0xc0, 0x72, 0x69, 0x7b, 0x64, 0xea, 0xbc, 0xec,
	0x72, 0x46, 0x5c, 0x96, 0x44, 0x70, 0x24, 0x6b, 0xee, 0x73, 0x26, 0x74, 0x04, 0xfb, 0xdc, 0xab,
	0x42, 0xfd, 0xf1, 0x0c, 0xb7, 0xfa, 0xfb, 0xe1, 0x76, 0x26, 0x61, 0xf8, 0xe9, 0xf6, 0x27, 0x46,
	0x2a, 0xa4, 0x0a, 0x19, 0x6c, 0x29, 0x0b, 0x50, 0x9b, 0x95, 0xfa, 0xda, 0xd7, 0xf3, 0x90, 0xaa,
	0xdc, 0xd6, 0xf3, 0x2c, 0xc8, 0x75, 0xfd, 0xdb, 0x1d, 0x98, 0x9f, 0x18, 0x87, 0x54, 0x61, 0x6c,
	0x3b, 0x30, 0xce, 0x02, 0x8c, 0x75, 0xcf, 0xb2, 0xab, 0x17, 0x1b, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x0f, 0xc8, 0x40, 0xe4, 0x1e, 0x02, 0x00, 0x00,
}
