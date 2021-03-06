// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im/im.book.proto

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

type Corpus int32

const (
	Corpus_UNIVERSAL Corpus = 0
	Corpus_WEB       Corpus = 1
	Corpus_IMAGES    Corpus = 2
	Corpus_LOCAL     Corpus = 3
	Corpus_NEW       Corpus = 4
	Corpus_PRODUCTS  Corpus = 5
	Corpus_VEDIO     Corpus = 6
)

var Corpus_name = map[int32]string{
	0: "UNIVERSAL",
	1: "WEB",
	2: "IMAGES",
	3: "LOCAL",
	4: "NEW",
	5: "PRODUCTS",
	6: "VEDIO",
}
var Corpus_value = map[string]int32{
	"UNIVERSAL": 0,
	"WEB":       1,
	"IMAGES":    2,
	"LOCAL":     3,
	"NEW":       4,
	"PRODUCTS":  5,
	"VEDIO":     6,
}

func (x Corpus) Enum() *Corpus {
	p := new(Corpus)
	*p = x
	return p
}
func (x Corpus) String() string {
	return proto.EnumName(Corpus_name, int32(x))
}
func (x *Corpus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Corpus_value, data, "Corpus")
	if err != nil {
		return err
	}
	*x = Corpus(value)
	return nil
}
func (Corpus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_im_book_fcdccfbaee9d0105, []int{0}
}

type EnumAllowAlias int32

const (
	EnumAllowAlias_UNKNOWN EnumAllowAlias = 0
	EnumAllowAlias_STARTED EnumAllowAlias = 1
	EnumAllowAlias_RUNNING EnumAllowAlias = 1
)

var EnumAllowAlias_name = map[int32]string{
	0: "UNKNOWN",
	1: "STARTED",
	// Duplicate value: 1: "RUNNING",
}
var EnumAllowAlias_value = map[string]int32{
	"UNKNOWN": 0,
	"STARTED": 1,
	"RUNNING": 1,
}

func (x EnumAllowAlias) Enum() *EnumAllowAlias {
	p := new(EnumAllowAlias)
	*p = x
	return p
}
func (x EnumAllowAlias) String() string {
	return proto.EnumName(EnumAllowAlias_name, int32(x))
}
func (x *EnumAllowAlias) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumAllowAlias_value, data, "EnumAllowAlias")
	if err != nil {
		return err
	}
	*x = EnumAllowAlias(value)
	return nil
}
func (EnumAllowAlias) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_im_book_fcdccfbaee9d0105, []int{1}
}

type Book struct {
	Id                   *int32   `protobuf:"varint,1,req,name=Id" json:"Id,omitempty"`
	Title                *string  `protobuf:"bytes,2,req,name=Title" json:"Title,omitempty"`
	SBN                  *string  `protobuf:"bytes,3,opt,name=SBN" json:"SBN,omitempty"`
	Author               *Author  `protobuf:"bytes,4,req,name=Author" json:"Author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_book_fcdccfbaee9d0105, []int{0}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (dst *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(dst, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Book) GetSBN() string {
	if m != nil && m.SBN != nil {
		return *m.SBN
	}
	return ""
}

func (m *Book) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

type Person struct {
	Id                   *int32   `protobuf:"varint,1,req,name=Id" json:"Id,omitempty"`
	Name                 *string  `protobuf:"bytes,2,req,name=Name" json:"Name,omitempty"`
	Address              *string  `protobuf:"bytes,3,opt,name=Address,def=Chongqing" json:"Address,omitempty"`
	Corpus               *Corpus  `protobuf:"varint,4,opt,name=corpus,enum=im.Corpus,def=0" json:"corpus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_im_book_fcdccfbaee9d0105, []int{1}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

const Default_Person_Address string = "Chongqing"
const Default_Person_Corpus Corpus = Corpus_UNIVERSAL

func (m *Person) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return Default_Person_Address
}

func (m *Person) GetCorpus() Corpus {
	if m != nil && m.Corpus != nil {
		return *m.Corpus
	}
	return Default_Person_Corpus
}

func init() {
	proto.RegisterType((*Book)(nil), "im.Book")
	proto.RegisterType((*Person)(nil), "im.Person")
	proto.RegisterEnum("im.Corpus", Corpus_name, Corpus_value)
	proto.RegisterEnum("im.EnumAllowAlias", EnumAllowAlias_name, EnumAllowAlias_value)
}

func init() { proto.RegisterFile("im/im.book.proto", fileDescriptor_im_book_fcdccfbaee9d0105) }

var fileDescriptor_im_book_fcdccfbaee9d0105 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8d, 0xcf, 0x8f, 0xa2, 0x40,
	0x10, 0x85, 0xed, 0xe6, 0x87, 0x4b, 0xb9, 0x6b, 0x3a, 0x95, 0x3d, 0x10, 0x4f, 0xc4, 0xbd, 0x10,
	0x93, 0x65, 0x93, 0x3d, 0xba, 0xa7, 0x16, 0x89, 0x21, 0xeb, 0x34, 0xa6, 0x01, 0xcd, 0x1c, 0x75,
	0x30, 0x4a, 0x04, 0xda, 0x01, 0xcd, 0x5c, 0xe7, 0x4f, 0x9f, 0x20, 0x9a, 0x39, 0xcc, 0xad, 0xde,
	0xfb, 0x5e, 0xea, 0x03, 0x96, 0x97, 0x7f, 0xf2, 0xd2, 0xdb, 0x29, 0x75, 0xf2, 0xce, 0xb5, 0xba,
	0x28, 0xa4, 0x79, 0x39, 0xc2, 0xae, 0xdd, 0x5e, 0x2f, 0x47, 0x55, 0x77, 0xfd, 0x78, 0x07, 0xfa,
	0x4c, 0xa9, 0x13, 0x0e, 0x81, 0x86, 0x99, 0x4d, 0x1c, 0xea, 0x1a, 0x92, 0x86, 0x19, 0xfe, 0x04,
	0x23, 0xc9, 0x2f, 0xc5, 0xde, 0xa6, 0x0e, 0x75, 0x2d, 0xd9, 0x05, 0x64, 0xa0, 0xc5, 0x33, 0x61,
	0x6b, 0x0e, 0x71, 0x2d, 0xd9, 0x9e, 0x38, 0x06, 0x93, 0xdf, 0xfe, 0xd9, 0xba, 0x43, 0xdd, 0xc1,
	0x5f, 0xf0, 0xf2, 0xd2, 0xeb, 0x1a, 0x79, 0x27, 0xe3, 0x77, 0x02, 0xe6, 0x6a, 0x5f, 0x37, 0xaa,
	0xfa, 0xa2, 0x41, 0xd0, 0xc5, 0xb6, 0x7c, 0x58, 0x6e, 0x37, 0xfe, 0x82, 0x3e, 0xcf, 0xb2, 0x7a,
	0xdf, 0x34, 0x9d, 0x68, 0x6a, 0xf9, 0x47, 0x55, 0x1d, 0x5e, 0xf3, 0xea, 0x20, 0x1f, 0x04, 0x7f,
	0x83, 0xf9, 0xa2, 0xea, 0xf3, 0xb5, 0xb1, 0x75, 0x87, 0xb8, 0xc3, 0xce, 0xeb, 0xdf, 0x9a, 0xa9,
	0x95, 0x8a, 0x70, 0x1d, 0xc8, 0x98, 0x2f, 0xe5, 0x7d, 0x34, 0x79, 0x06, 0xb3, 0x83, 0xf8, 0x03,
	0x3e, 0x31, 0xeb, 0x61, 0x1f, 0xb4, 0x4d, 0x30, 0x63, 0x04, 0x01, 0xcc, 0xf0, 0x89, 0x2f, 0x82,
	0x98, 0x51, 0xb4, 0xc0, 0x58, 0x46, 0x3e, 0x5f, 0x32, 0xad, 0xe5, 0x22, 0xd8, 0x30, 0x1d, 0xbf,
	0xc3, 0xb7, 0x95, 0x8c, 0xe6, 0xa9, 0x9f, 0xc4, 0xcc, 0x68, 0x17, 0xeb, 0x60, 0x1e, 0x46, 0xcc,
	0x9c, 0xfc, 0x83, 0x61, 0x50, 0x5d, 0x4b, 0x5e, 0x14, 0xea, 0x8d, 0x17, 0xf9, 0xb6, 0xc1, 0x01,
	0xf4, 0x53, 0xf1, 0x5f, 0x44, 0x1b, 0xc1, 0x7a, 0x6d, 0x88, 0x13, 0x2e, 0x93, 0x60, 0xce, 0x48,
	0x1b, 0x64, 0x2a, 0x44, 0x28, 0x16, 0x8c, 0x8c, 0x28, 0x23, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xd8, 0x05, 0x2d, 0xfc, 0xa9, 0x01, 0x00, 0x00,
}
