// Code generated by protoc-gen-go.
// source: api/common.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/common.proto
	api/show_problem.proto
	api/submit.proto

It has these top-level messages:
	Error
	ShowProblemRequest
	ShowProblemResponse
	SubmitRequest
	SubmitResponse
*/
package api

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

type Error struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "api.Error")
}

func init() { proto.RegisterFile("api/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0xc8, 0xd4,
	0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c,
	0xc8, 0x54, 0xd2, 0xe5, 0x62, 0x75, 0x2d, 0x2a, 0xca, 0x2f, 0x12, 0x12, 0xe2, 0x62, 0x49, 0xce,
	0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x73,
	0x8b, 0xd3, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x40, 0xcc, 0x24, 0x36, 0xb0, 0x56, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x3f, 0x3e, 0xa6, 0x4e, 0x00, 0x00, 0x00,
}
