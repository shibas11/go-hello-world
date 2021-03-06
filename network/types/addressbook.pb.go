// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=types.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("types.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "types.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "types.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "types.AddressBook")
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor_1eb1a68c9dd6d429) }

var fileDescriptor_1eb1a68c9dd6d429 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x80, 0x4d, 0x9a, 0x04, 0x3b, 0xd1, 0x52, 0x07, 0xd1, 0xd8, 0x8b, 0xa1, 0x20, 0x04, 0x0a,
	0x5b, 0xac, 0x5e, 0x3d, 0x58, 0x28, 0x28, 0x5a, 0x5b, 0x96, 0x8a, 0x47, 0x49, 0xc8, 0x58, 0x43,
	0x93, 0xec, 0x92, 0x4d, 0x0f, 0x7d, 0x30, 0xdf, 0x4f, 0xb2, 0x9b, 0x8a, 0x82, 0xb7, 0xf9, 0xf9,
	0x98, 0xf9, 0x66, 0xe0, 0x24, 0x4e, 0xd3, 0x8a, 0x94, 0x4a, 0x84, 0xd8, 0x30, 0x59, 0x89, 0x5a,
	0xa0, 0x5b, 0xef, 0x24, 0xa9, 0xc1, 0xe5, 0x5a, 0x88, 0x75, 0x4e, 0x63, 0x5d, 0x4c, 0xb6, 0x1f,
	0xe3, 0x3a, 0x2b, 0x48, 0xd5, 0x71, 0x21, 0x0d, 0x37, 0xfc, 0xb2, 0xc1, 0x5b, 0x52, 0xa5, 0x44,
	0x89, 0x08, 0x4e, 0x19, 0x17, 0x14, 0x58, 0xa1, 0x15, 0x75, 0xb9, 0x8e, 0xb1, 0x07, 0x76, 0x96,
	0x06, 0x76, 0x68, 0x45, 0x2e, 0xb7, 0xb3, 0x14, 0x4f, 0xc1, 0xa5, 0x22, 0xce, 0xf2, 0xa0, 0xa3,
	0x21, 0x93, 0xe0, 0x35, 0x78, 0xf2, 0x53, 0x94, 0xa4, 0x02, 0x27, 0xec, 0x44, 0xfe, 0xe4, 0x82,
	0xe9, 0xed, 0xcc, 0x0c, 0x66, 0xcb, 0xa6, 0xf7, 0xb2, 0x2d, 0x12, 0xaa, 0x78, 0x0b, 0xe2, 0x1d,
	0x1c, 0xe5, 0xb1, 0xaa, 0xdf, 0xb7, 0x32, 0x8d, 0x6b, 0x4a, 0x03, 0x37, 0xb4, 0x22, 0x7f, 0x32,
	0x60, 0xc6, 0x97, 0xed, 0x7d, 0xd9, 0x6a, 0xef, 0xcb, 0xfd, 0x86, 0x7f, 0x35, 0xf8, 0x80, 0x83,
	0xff, 0x6b, 0x2a, 0x9e, 0x81, 0x57, 0xea, 0xa8, 0x95, 0x6f, 0x33, 0x1c, 0x81, 0xd3, 0x98, 0xe8,
	0x03, 0x7a, 0x93, 0xf3, 0x7f, 0xb4, 0x56, 0x3b, 0x49, 0x5c, 0x43, 0xc3, 0x11, 0x74, 0x7f, 0x4a,
	0x08, 0xe0, 0xcd, 0x17, 0xd3, 0xc7, 0xe7, 0x59, 0xff, 0x00, 0x0f, 0xc1, 0x79, 0x58, 0xcc, 0x67,
	0x7d, 0xab, 0x89, 0xde, 0x16, 0xfc, 0xa9, 0x6f, 0x0f, 0x6f, 0xc1, 0xbf, 0x37, 0x4f, 0x9f, 0x0a,
	0xb1, 0xc1, 0x2b, 0xf0, 0x24, 0x09, 0x99, 0x37, 0xdf, 0x6b, 0x3e, 0x70, 0xfc, 0x67, 0x15, 0x6f,
	0x9b, 0x89, 0xa7, 0xef, 0xba, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x43, 0xee, 0x68, 0xb1,
	0x01, 0x00, 0x00,
}
