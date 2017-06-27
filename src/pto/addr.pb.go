// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addr.proto

/*
Package tutorial is a generated protocol buffer package.

It is generated from these files:
	addr.proto

It has these top-level messages:
	Person
	AddressBook
*/
package tutorial

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
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

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
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("addr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x74, 0xd3, 0x74, 0x69, 0x5f, 0xa0, 0x84, 0x87, 0x48, 0x28, 0x1e, 0x42, 0x4e, 0x01, 0x61,
	0x0f, 0x55, 0xf0, 0x6c, 0xa1, 0xa0, 0x68, 0x4d, 0x59, 0x14, 0xcf, 0x29, 0xfb, 0xc0, 0x60, 0x92,
	0xb7, 0x6c, 0xd2, 0x43, 0xff, 0xdd, 0x83, 0x64, 0x1b, 0x45, 0xc4, 0xdb, 0xbc, 0x99, 0x61, 0x76,
	0x76, 0x00, 0x4a, 0x63, 0x9c, 0xb2, 0x8e, 0x7b, 0xc6, 0x59, 0x7f, 0xe8, 0xd9, 0x55, 0x65, 0x9d,
	0x7d, 0x0a, 0x90, 0x3b, 0x72, 0x1d, 0xb7, 0x88, 0x10, 0xb6, 0x65, 0x43, 0x89, 0x48, 0x45, 0x3e,
	0xd7, 0x1e, 0xe3, 0x02, 0x82, 0xca, 0x24, 0x41, 0x2a, 0xf2, 0xa9, 0x0e, 0x2a, 0x83, 0xe7, 0x30,
	0xa5, 0xa6, 0xac, 0xea, 0x64, 0xe2, 0x4d, 0xa7, 0x03, 0x6f, 0x40, 0xda, 0x77, 0x6e, 0xa9, 0x4b,
	0xc2, 0x74, 0x92, 0x47, 0xab, 0x4b, 0xf5, 0x9d, 0xaf, 0x4e, 0xd9, 0x6a, 0x37, 0xc8, 0xcf, 0x87,
	0x66, 0x4f, 0x4e, 0x8f, 0xde, 0xe5, 0x2b, 0x44, 0xbf, 0x68, 0xbc, 0x00, 0xd9, 0x7a, 0x34, 0x16,
	0x18, 0x2f, 0x54, 0x10, 0xf6, 0x47, 0x4b, 0xbe, 0xc4, 0x62, 0xb5, 0xfc, 0x3f, 0xfa, 0xe5, 0x68,
	0x49, 0x7b, 0x5f, 0x76, 0x05, 0xf3, 0x1f, 0x0a, 0x01, 0xe4, 0xb6, 0x58, 0x3f, 0x3c, 0x6d, 0xe2,
	0x33, 0x9c, 0x41, 0x78, 0x5f, 0x6c, 0x37, 0xb1, 0x18, 0xd0, 0x5b, 0xa1, 0x1f, 0xe3, 0x20, 0xbb,
	0x85, 0xe8, 0xce, 0x18, 0x47, 0x5d, 0xb7, 0x66, 0xfe, 0xc0, 0x1c, 0xa4, 0x25, 0xb6, 0xf5, 0x30,
	0xc2, 0xf0, 0x91, 0xf8, 0xef, 0x6b, 0x7a, 0xd4, 0xf7, 0xd2, 0x0f, 0x79, 0xfd, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0x89, 0xd3, 0x48, 0x56, 0x01, 0x00, 0x00,
}
