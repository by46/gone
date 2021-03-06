// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im/im.foo.proto

package im

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_foo_f37d429e64f8774a, []int{0}
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
	if m != nil {
		return m.Id
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
	return fileDescriptor_im_foo_f37d429e64f8774a, []int{1}
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
	ID                   int32    `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DFIS) Reset()         { *m = DFIS{} }
func (m *DFIS) String() string { return proto.CompactTextString(m) }
func (*DFIS) ProtoMessage()    {}
func (*DFIS) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_foo_f37d429e64f8774a, []int{2}
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
	if m != nil {
		return m.ID
	}
	return 0
}

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_foo_f37d429e64f8774a, []int{3}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_foo_f37d429e64f8774a, []int{4}
}
func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (dst *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(dst, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

var E_NeweggWatch = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51234,
	Name:          "im.newegg_watch",
	Tag:           "bytes,51234,opt,name=newegg_watch,json=neweggWatch",
	Filename:      "im/im.foo.proto",
}

func init() {
	proto.RegisterType((*Foo)(nil), "im.Foo")
	proto.RegisterType((*Info)(nil), "im.Info")
	proto.RegisterType((*DFIS)(nil), "im.DFIS")
	proto.RegisterType((*HelloRequest)(nil), "im.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "im.HelloReply")
	proto.RegisterExtension(E_NeweggWatch)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/im.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/im.Greeter/SayHelloAgain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.Greeter/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "im.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Greeter_SayHelloAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im/im.foo.proto",
}

func init() { proto.RegisterFile("im/im.foo.proto", fileDescriptor_im_foo_f37d429e64f8774a) }

var fileDescriptor_im_foo_f37d429e64f8774a = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x6a, 0xea, 0x40,
	0x14, 0x86, 0x8d, 0x46, 0x8d, 0xe7, 0x7a, 0xbd, 0x97, 0xa1, 0x85, 0x54, 0x0a, 0x0d, 0x59, 0x14,
	0x17, 0x25, 0xd2, 0x76, 0xe7, 0xce, 0x12, 0xac, 0x59, 0xd8, 0x42, 0x5c, 0x74, 0x29, 0xd1, 0x9c,
	0xa4, 0x53, 0x92, 0x9c, 0x74, 0x66, 0x44, 0xdc, 0xf6, 0x09, 0x4a, 0x1f, 0xa1, 0x4f, 0x5a, 0x92,
	0x18, 0x71, 0xd7, 0xdd, 0x9c, 0x9f, 0xef, 0x9b, 0xf3, 0xc3, 0x81, 0x7f, 0x3c, 0x1d, 0xf3, 0xd4,
	0x89, 0x88, 0x9c, 0x5c, 0x90, 0x22, 0xd6, 0xe4, 0xe9, 0xd0, 0x8a, 0x89, 0xe2, 0x04, 0xc7, 0x65,
	0xb2, 0xde, 0x46, 0xe3, 0x10, 0xe5, 0x46, 0xf0, 0x5c, 0x91, 0xa8, 0x28, 0xfb, 0x1c, 0x5a, 0x33,
	0x22, 0x36, 0x80, 0xa6, 0x17, 0x9a, 0x9a, 0xa5, 0x8d, 0xda, 0x7e, 0xd3, 0x0b, 0xed, 0x10, 0x74,
	0x2f, 0x8b, 0x88, 0x9d, 0x81, 0xfe, 0x14, 0xa4, 0x68, 0xea, 0x96, 0x36, 0xea, 0xcd, 0x1b, 0x7e,
	0x39, 0xb1, 0x4b, 0x30, 0x16, 0x81, 0x10, 0x3c, 0x88, 0xd1, 0x6c, 0x5b, 0xda, 0xc8, 0x98, 0x37,
	0xfc, 0x63, 0xc2, 0x86, 0xd0, 0x9d, 0x86, 0xa1, 0x40, 0x29, 0xcd, 0xce, 0x41, 0xab, 0x83, 0x87,
	0x3e, 0x80, 0x42, 0xa9, 0x56, 0x94, 0x21, 0x45, 0xb6, 0x05, 0xba, 0x3b, 0xf3, 0x96, 0xe5, 0x76,
	0xf7, 0xb8, 0xdd, 0x9d, 0x18, 0x5f, 0x1f, 0x17, 0x7a, 0x18, 0x71, 0x69, 0xdb, 0xd0, 0x9f, 0x63,
	0x92, 0x90, 0x8f, 0xef, 0x5b, 0x94, 0x8a, 0x31, 0xd0, 0xb3, 0xa2, 0x4f, 0xc1, 0xf6, 0xfc, 0xf2,
	0x6d, 0x5f, 0x03, 0x1c, 0x98, 0x3c, 0xd9, 0x33, 0x13, 0xba, 0x29, 0x4a, 0x59, 0x54, 0xab, 0xa0,
	0x7a, 0xbc, 0x7b, 0x83, 0xee, 0xa3, 0x40, 0x54, 0x28, 0xd8, 0x0d, 0x18, 0xcb, 0x60, 0x5f, 0x5a,
	0xec, 0xbf, 0xc3, 0x53, 0xe7, 0x74, 0xc9, 0x70, 0x70, 0x92, 0x14, 0x5f, 0xde, 0xc2, 0xdf, 0x9a,
	0x9e, 0xc6, 0x01, 0xcf, 0x7e, 0x57, 0x26, 0x2e, 0xf4, 0x33, 0xdc, 0x61, 0x1c, 0xaf, 0x76, 0x81,
	0xda, 0xbc, 0xb2, 0x2b, 0xa7, 0xba, 0x84, 0x53, 0x5f, 0xc2, 0x59, 0x54, 0xad, 0x9e, 0x73, 0xc5,
	0x29, 0x93, 0xe6, 0xf7, 0x67, 0xab, 0x6c, 0xfb, 0xa7, 0xd2, 0x5e, 0x0a, 0x6b, 0xdd, 0x29, 0xe9,
	0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xec, 0xb4, 0xbe, 0xdc, 0x01, 0x00, 0x00,
}
