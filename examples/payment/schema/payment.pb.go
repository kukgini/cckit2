// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment.proto

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	payment.proto

It has these top-level messages:
	Payment
	PaymentList
	PaymentEvent
*/
package schema

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

type Payment struct {
	Type   string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Amount int32  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *Payment) Reset()                    { *m = Payment{} }
func (m *Payment) String() string            { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()               {}
func (*Payment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Payment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Payment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Payment) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type PaymentList struct {
	Items []*Payment `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
}

func (m *PaymentList) Reset()                    { *m = PaymentList{} }
func (m *PaymentList) String() string            { return proto.CompactTextString(m) }
func (*PaymentList) ProtoMessage()               {}
func (*PaymentList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PaymentList) GetItems() []*Payment {
	if m != nil {
		return m.Items
	}
	return nil
}

type PaymentEvent struct {
	Type   string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Amount int32  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *PaymentEvent) Reset()                    { *m = PaymentEvent{} }
func (m *PaymentEvent) String() string            { return proto.CompactTextString(m) }
func (*PaymentEvent) ProtoMessage()               {}
func (*PaymentEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PaymentEvent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PaymentEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PaymentEvent) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*Payment)(nil), "schema.Payment")
	proto.RegisterType((*PaymentList)(nil), "schema.PaymentList")
	proto.RegisterType((*PaymentEvent)(nil), "schema.PaymentEvent")
}

func init() { proto.RegisterFile("payment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xac, 0xcc,
	0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x4e, 0xce, 0x48, 0xcd,
	0x4d, 0x54, 0x72, 0xe5, 0x62, 0x0f, 0x80, 0x48, 0x08, 0x09, 0x71, 0xb1, 0x94, 0x54, 0x16, 0xa4,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12,
	0x4c, 0x60, 0x11, 0xa6, 0xcc, 0x14, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xdc, 0xfc, 0xd2, 0xbc, 0x12,
	0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x28, 0x4f, 0xc9, 0x84, 0x8b, 0x1b, 0x6a, 0x8c, 0x4f,
	0x66, 0x71, 0x89, 0x90, 0x2a, 0x17, 0xab, 0x67, 0x49, 0x6a, 0x6e, 0xb1, 0x04, 0xa3, 0x02, 0xb3,
	0x06, 0xb7, 0x11, 0xbf, 0x1e, 0xc4, 0x36, 0x3d, 0xa8, 0x9a, 0x20, 0x88, 0xac, 0x92, 0x17, 0x17,
	0x0f, 0x54, 0xc4, 0xb5, 0x8c, 0x42, 0x17, 0x24, 0xb1, 0x81, 0xfd, 0x65, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x27, 0xee, 0x9d, 0x78, 0xe8, 0x00, 0x00, 0x00,
}
