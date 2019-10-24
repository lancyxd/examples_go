// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Phone
	Person
	ContactBook
*/
package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PhoneType int32

const (
	PhoneType_HOME PhoneType = 0
	PhoneType_WORK PhoneType = 1
)

var PhoneType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
}
var PhoneType_value = map[string]int32{
	"HOME": 0,
	"WORK": 1,
}

func (x PhoneType) Enum() *PhoneType {
	p := new(PhoneType)
	*p = x
	return p
}
func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}
func (x *PhoneType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhoneType_value, data, "PhoneType")
	if err != nil {
		return err
	}
	*x = PhoneType(value)
	return nil
}
func (PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Phone struct {
	Type             *PhoneType `protobuf:"varint,1,opt,name=type,enum=test.PhoneType" json:"type,omitempty"`
	Number           *string    `protobuf:"bytes,2,opt,name=number" json:"number,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Phone) Reset()                    { *m = Phone{} }
func (m *Phone) String() string            { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()               {}
func (*Phone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Phone) GetType() PhoneType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PhoneType_HOME
}

func (m *Phone) GetNumber() string {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return ""
}

type Person struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Phones           []*Phone `protobuf:"bytes,3,rep,name=phones" json:"phones,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

func (m *Person) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

type ContactBook struct {
	Persons          []*Person `protobuf:"bytes,1,rep,name=persons" json:"persons,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ContactBook) Reset()                    { *m = ContactBook{} }
func (m *ContactBook) String() string            { return proto.CompactTextString(m) }
func (*ContactBook) ProtoMessage()               {}
func (*ContactBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ContactBook) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

func init() {
	proto.RegisterType((*Phone)(nil), "test.Phone")
	proto.RegisterType((*Person)(nil), "test.Person")
	proto.RegisterType((*ContactBook)(nil), "test.ContactBook")
	proto.RegisterEnum("test.PhoneType", PhoneType_name, PhoneType_value)
}

var fileDescriptor0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xcc, 0xb8, 0x58, 0x03, 0x32,
	0xf2, 0xf3, 0x52, 0x85, 0x64, 0xb9, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35,
	0xf8, 0x8c, 0xf8, 0xf5, 0xc0, 0x2a, 0xc1, 0x52, 0x21, 0x40, 0x61, 0x21, 0x3e, 0x2e, 0xb6, 0xbc,
	0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x09, 0x26, 0xa0, 0x02, 0x4e, 0x25, 0x7b, 0x2e, 0xb6, 0x80, 0xd4,
	0xa2, 0xe2, 0xfc, 0x3c, 0x21, 0x2e, 0x2e, 0xa6, 0xcc, 0x14, 0xb0, 0x36, 0x56, 0x21, 0x1e, 0x2e,
	0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x88, 0x1a, 0x21, 0x69, 0x2e, 0xb6, 0x02, 0x90, 0x01, 0xc5, 0x12,
	0xcc, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xdc, 0x48, 0x86, 0x2a, 0xe9, 0x70, 0x71, 0x3b, 0xe7, 0xe7,
	0x95, 0x24, 0x26, 0x97, 0x38, 0xe5, 0xe7, 0x67, 0x03, 0xad, 0x67, 0x2f, 0x00, 0x9b, 0x57, 0x0c,
	0x34, 0x0a, 0xa4, 0x98, 0x07, 0xaa, 0x18, 0x2c, 0xa8, 0x25, 0xcf, 0xc5, 0x89, 0x70, 0x0b, 0x07,
	0x17, 0x8b, 0x87, 0xbf, 0xaf, 0xab, 0x00, 0x03, 0x88, 0x15, 0xee, 0x1f, 0xe4, 0x2d, 0xc0, 0x08,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xf1, 0xb5, 0xbf, 0xda, 0x00, 0x00, 0x00,
}
