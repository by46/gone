// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im/im.foo.proto

package im

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Foo struct {
	Id                           *int32   `protobuf:"varint,1,req,name=Id" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_foo_72be337b0923d521, []int{0}
}

var extRange_Foo = []proto.ExtensionRange{
	{Start: 100, End: 199},
}

func (*Foo) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Foo
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (dst *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(dst, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type Info struct {
	// Types that are valid to be assigned to TestOneof:
	//	*Info_Name
	//	*Info_Marriage
	//	*Info_Address
	TestOneof            isInfo_TestOneof `protobuf_oneof:"test_oneof"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_foo_72be337b0923d521, []int{1}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

type isInfo_TestOneof interface {
	isInfo_TestOneof()
}

type Info_Name struct {
	Name string `protobuf:"bytes,4,opt,name=Name,oneof"`
}
type Info_Marriage struct {
	Marriage bool `protobuf:"varint,5,opt,name=Marriage,oneof"`
}
type Info_Address struct {
	Address string `protobuf:"bytes,6,opt,name=Address,oneof"`
}

func (*Info_Name) isInfo_TestOneof()     {}
func (*Info_Marriage) isInfo_TestOneof() {}
func (*Info_Address) isInfo_TestOneof()  {}

func (m *Info) GetTestOneof() isInfo_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (m *Info) GetName() string {
	if x, ok := m.GetTestOneof().(*Info_Name); ok {
		return x.Name
	}
	return ""
}

func (m *Info) GetMarriage() bool {
	if x, ok := m.GetTestOneof().(*Info_Marriage); ok {
		return x.Marriage
	}
	return false
}

func (m *Info) GetAddress() string {
	if x, ok := m.GetTestOneof().(*Info_Address); ok {
		return x.Address
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Info) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Info_OneofMarshaler, _Info_OneofUnmarshaler, _Info_OneofSizer, []interface{}{
		(*Info_Name)(nil),
		(*Info_Marriage)(nil),
		(*Info_Address)(nil),
	}
}

func _Info_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Info)
	// test_oneof
	switch x := m.TestOneof.(type) {
	case *Info_Name:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *Info_Marriage:
		t := uint64(0)
		if x.Marriage {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Info_Address:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Address)
	case nil:
	default:
		return fmt.Errorf("Info.TestOneof has unexpected type %T", x)
	}
	return nil
}

func _Info_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Info)
	switch tag {
	case 4: // test_oneof.Name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.TestOneof = &Info_Name{x}
		return true, err
	case 5: // test_oneof.Marriage
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TestOneof = &Info_Marriage{x != 0}
		return true, err
	case 6: // test_oneof.Address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.TestOneof = &Info_Address{x}
		return true, err
	default:
		return false, nil
	}
}

func _Info_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Info)
	// test_oneof
	switch x := m.TestOneof.(type) {
	case *Info_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *Info_Marriage:
		n += 1 // tag and wire
		n += 1
	case *Info_Address:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Address)))
		n += len(x.Address)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DFIS struct {
	ID                   *int32   `protobuf:"varint,1,req,name=ID" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DFIS) Reset()         { *m = DFIS{} }
func (m *DFIS) String() string { return proto.CompactTextString(m) }
func (*DFIS) ProtoMessage()    {}
func (*DFIS) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_foo_72be337b0923d521, []int{2}
}
func (m *DFIS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DFIS.Unmarshal(m, b)
}
func (m *DFIS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DFIS.Marshal(b, m, deterministic)
}
func (dst *DFIS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DFIS.Merge(dst, src)
}
func (m *DFIS) XXX_Size() int {
	return xxx_messageInfo_DFIS.Size(m)
}
func (m *DFIS) XXX_DiscardUnknown() {
	xxx_messageInfo_DFIS.DiscardUnknown(m)
}

var xxx_messageInfo_DFIS proto.InternalMessageInfo

func (m *DFIS) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

var E_NeweggWatch = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51234,
	Name:          "im.newegg_watch",
	Tag:           "bytes,51234,opt,name=newegg_watch,json=neweggWatch",
	Filename:      "im/im.foo.proto",
}

var E_Name = &proto.ExtensionDesc{
	ExtendedType:  (*Foo)(nil),
	ExtensionType: (*string)(nil),
	Field:         100,
	Name:          "im.Name",
	Tag:           "bytes,100,opt,name=Name",
	Filename:      "im/im.foo.proto",
}

func init() {
	proto.RegisterType((*Foo)(nil), "im.Foo")
	proto.RegisterType((*Info)(nil), "im.Info")
	proto.RegisterType((*DFIS)(nil), "im.DFIS")
	proto.RegisterExtension(E_NeweggWatch)
	proto.RegisterExtension(E_Name)
}

func init() { proto.RegisterFile("im/im.foo.proto", fileDescriptor_im_foo_72be337b0923d521) }

var fileDescriptor_im_foo_72be337b0923d521 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x9b, 0x34, 0xb5, 0x71, 0x2c, 0x2a, 0x8b, 0x87, 0xb5, 0x16, 0x0c, 0x3d, 0x05, 0x0f,
	0x9b, 0x7b, 0x6e, 0x4a, 0x08, 0xcd, 0xa1, 0x0a, 0xf1, 0xe0, 0xb1, 0xc4, 0xee, 0x64, 0x5d, 0x30,
	0x99, 0xb0, 0xbb, 0xd2, 0xbb, 0x4f, 0x20, 0x3e, 0x82, 0x4f, 0xe3, 0x63, 0x49, 0x13, 0xdb, 0xe3,
	0x7c, 0xfc, 0x3f, 0xf3, 0x7f, 0x70, 0xa1, 0x9b, 0x44, 0x37, 0xa2, 0x26, 0x12, 0x9d, 0x21, 0x47,
	0xcc, 0xd7, 0xcd, 0x3c, 0x52, 0x44, 0xea, 0x1d, 0x93, 0x9e, 0xbc, 0x7e, 0xd4, 0x89, 0x44, 0xbb,
	0x35, 0xba, 0x73, 0x64, 0x86, 0xd4, 0x72, 0x01, 0xe3, 0x9c, 0x88, 0x9d, 0x83, 0x5f, 0x48, 0xee,
	0x45, 0x7e, 0x3c, 0x29, 0xfd, 0x42, 0xde, 0x4d, 0x42, 0x79, 0xf9, 0xeb, 0x2d, 0x25, 0x04, 0x45,
	0x5b, 0x13, 0xbb, 0x82, 0xe0, 0xb1, 0x6a, 0x90, 0x07, 0x91, 0x17, 0x9f, 0xae, 0x46, 0x65, 0x7f,
	0xb1, 0x05, 0x84, 0xeb, 0xca, 0x18, 0x5d, 0x29, 0xe4, 0x93, 0xc8, 0x8b, 0xc3, 0xd5, 0xa8, 0x3c,
	0x12, 0x36, 0x87, 0xe9, 0xbd, 0x94, 0x06, 0xad, 0xe5, 0x27, 0xff, 0xb5, 0x03, 0x78, 0x98, 0x01,
	0x38, 0xb4, 0x6e, 0x43, 0x2d, 0x52, 0xbd, 0x8c, 0x20, 0xc8, 0xf2, 0xe2, 0xb9, 0x1f, 0x91, 0x1d,
	0x47, 0x64, 0x69, 0xf8, 0xfd, 0x79, 0x1d, 0xc8, 0x5a, 0xdb, 0x34, 0x83, 0x59, 0x8b, 0x3b, 0x54,
	0x6a, 0xb3, 0xab, 0xdc, 0xf6, 0x8d, 0xdd, 0x8a, 0x41, 0x4c, 0x1c, 0xc4, 0xc4, 0x1a, 0xad, 0xad,
	0x14, 0x3e, 0x75, 0x4e, 0x53, 0x6b, 0xf9, 0xcf, 0xd7, 0x78, 0xff, 0xb3, 0x3c, 0x1b, 0x6a, 0x2f,
	0xfb, 0x56, 0x7a, 0x33, 0x58, 0xb0, 0xa9, 0xd0, 0x8d, 0xc8, 0x89, 0xb8, 0xec, 0x33, 0x3d, 0xfc,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x70, 0x10, 0x7f, 0x40, 0x01, 0x00, 0x00,
}
