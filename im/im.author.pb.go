// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im/im.author.proto

package im

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

type Author struct {
	Name                 *string  `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Ages                 []int32  `protobuf:"varint,2,rep,name=Ages" json:"Ages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_author_42e472f026ad8aef, []int{0}
}
func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (dst *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(dst, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Author) GetAges() []int32 {
	if m != nil {
		return m.Ages
	}
	return nil
}

func init() {
	proto.RegisterType((*Author)(nil), "im.Author")
}

func init() { proto.RegisterFile("im/im.author.proto", fileDescriptor_im_author_42e472f026ad8aef) }

var fileDescriptor_im_author_42e472f026ad8aef = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcc, 0xd5, 0xcf,
	0xcc, 0xd5, 0x4b, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0xca, 0xcc, 0x55, 0x32, 0xe0, 0x62, 0x73, 0x04, 0x8b, 0x09, 0x09, 0x71, 0xb1, 0xf8, 0x25, 0xe6,
	0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x69, 0x70, 0x06, 0x81, 0xd9, 0x20, 0x31, 0xc7, 0xf4, 0xd4, 0x62,
	0x09, 0x26, 0x05, 0x66, 0x0d, 0xd6, 0x20, 0x30, 0x1b, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x9d,
	0xbd, 0x90, 0x4a, 0x00, 0x00, 0x00,
}
