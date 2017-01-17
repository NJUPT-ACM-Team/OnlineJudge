// Code generated by protoc-gen-go.
// source: api/edit_problem.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EditProblemRequest struct {
	Sid string `protobuf:"bytes,1,opt,name=sid" json:"sid,omitempty"`
}

func (m *EditProblemRequest) Reset()                    { *m = EditProblemRequest{} }
func (m *EditProblemRequest) String() string            { return proto.CompactTextString(m) }
func (*EditProblemRequest) ProtoMessage()               {}
func (*EditProblemRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EditProblemRequest) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

type EditProblemResponse struct {
	Sid     string   `protobuf:"bytes,1,opt,name=sid" json:"sid,omitempty"`
	Problem *Problem `protobuf:"bytes,2,opt,name=problem" json:"problem,omitempty"`
	Error   *Error   `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *EditProblemResponse) Reset()                    { *m = EditProblemResponse{} }
func (m *EditProblemResponse) String() string            { return proto.CompactTextString(m) }
func (*EditProblemResponse) ProtoMessage()               {}
func (*EditProblemResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *EditProblemResponse) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *EditProblemResponse) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

func (m *EditProblemResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*EditProblemRequest)(nil), "api.EditProblemRequest")
	proto.RegisterType((*EditProblemResponse)(nil), "api.EditProblemResponse")
}

func init() { proto.RegisterFile("api/edit_problem.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2c, 0xc8, 0xd4,
	0x4f, 0x4d, 0xc9, 0x2c, 0x89, 0x2f, 0x28, 0xca, 0x4f, 0xca, 0x49, 0xcd, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x12, 0x00, 0x49, 0x26, 0xe7, 0xe7, 0xe6, 0xe6,
	0xe7, 0x41, 0x84, 0xa5, 0xc0, 0xca, 0x8b, 0x33, 0xf2, 0xcb, 0x51, 0x95, 0x2b, 0xa9, 0x71, 0x09,
	0xb9, 0xa6, 0x64, 0x96, 0x04, 0x40, 0x04, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04,
	0xb8, 0x98, 0x8b, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0xa5, 0x42,
	0x2e, 0x61, 0x14, 0x75, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x98, 0x0a, 0x85, 0xd4, 0xb8, 0xd8,
	0xa1, 0x36, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0xf1, 0xe8, 0x25, 0x16, 0x64, 0xea, 0xc1,
	0x34, 0xc2, 0x24, 0x85, 0x14, 0xb8, 0x58, 0x53, 0x8b, 0x8a, 0xf2, 0x8b, 0x24, 0x98, 0xc1, 0xaa,
	0xb8, 0xc0, 0xaa, 0x5c, 0x41, 0x22, 0x41, 0x10, 0x89, 0x24, 0x36, 0xb0, 0x0b, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd5, 0x18, 0x48, 0x6c, 0xea, 0x00, 0x00, 0x00,
}