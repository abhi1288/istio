// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/mcp/v1alpha1/metadata.proto

package v1alpha1 // import "istio.io/api/config/mcp/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Metadata information that all resources within the Mesh Configuration Protocol must have.
type Metadata struct {
	// The name of the resource. It is unique within the context of a resource type and the origin server of
	// the resource. The resource type is identified by the TypeUrl of the resource field of the Envelope
	// message.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_metadata_a8a46deddbb2f3bb, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "istio.config.mcp.v1alpha1.Metadata")
}

func init() {
	proto.RegisterFile("config/mcp/v1alpha1/metadata.proto", fileDescriptor_metadata_a8a46deddbb2f3bb)
}

var fileDescriptor_metadata_a8a46deddbb2f3bb = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0xcf, 0x4d, 0x2e, 0xd0, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4,
	0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0xcc, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x83, 0xa8, 0xd4, 0xcb, 0x4d, 0x2e, 0xd0, 0x83, 0xa9, 0x54,
	0x92, 0xe3, 0xe2, 0xf0, 0x85, 0x2a, 0x16, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0x94, 0xa2, 0x14, 0x20, 0x9a, 0x33, 0xf3, 0xf5,
	0x13, 0x0b, 0x32, 0xf5, 0xb1, 0xd8, 0x96, 0xc4, 0x06, 0xb6, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x48, 0x8a, 0xd1, 0x7e, 0x8b, 0x00, 0x00, 0x00,
}
