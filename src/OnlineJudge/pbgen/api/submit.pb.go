// Code generated by protoc-gen-go.
// source: api/submit.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SubmitRequest struct {
	ContestId  uint64 `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	ProblemSid string `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	LanguageId int64  `protobuf:"varint,4,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	IsShared   bool   `protobuf:"varint,5,opt,name=is_shared,json=isShared" json:"is_shared,omitempty"`
}

func (m *SubmitRequest) Reset()                    { *m = SubmitRequest{} }
func (m *SubmitRequest) String() string            { return proto.CompactTextString(m) }
func (*SubmitRequest) ProtoMessage()               {}
func (*SubmitRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *SubmitRequest) GetContestId() uint64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *SubmitRequest) GetProblemSid() string {
	if m != nil {
		return m.ProblemSid
	}
	return ""
}

func (m *SubmitRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SubmitRequest) GetLanguageId() int64 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *SubmitRequest) GetIsShared() bool {
	if m != nil {
		return m.IsShared
	}
	return false
}

type SubmitResponse struct {
	RunId int64  `protobuf:"varint,1,opt,name=run_id,json=runId" json:"run_id,omitempty"`
	Error *Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SubmitResponse) Reset()                    { *m = SubmitResponse{} }
func (m *SubmitResponse) String() string            { return proto.CompactTextString(m) }
func (*SubmitResponse) ProtoMessage()               {}
func (*SubmitResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *SubmitResponse) GetRunId() int64 {
	if m != nil {
		return m.RunId
	}
	return 0
}

func (m *SubmitResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*SubmitRequest)(nil), "api.SubmitRequest")
	proto.RegisterType((*SubmitResponse)(nil), "api.SubmitResponse")
}

func init() { proto.RegisterFile("api/submit.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x34, 0x8f, 0xc1, 0x6a, 0xc3, 0x30,
	0x0c, 0x86, 0xf1, 0x92, 0x94, 0x46, 0x65, 0x63, 0x18, 0x06, 0x61, 0x63, 0xcc, 0xf4, 0x94, 0x53,
	0x06, 0xdb, 0x33, 0xec, 0x90, 0xab, 0xf3, 0x00, 0xc1, 0x89, 0x4d, 0x67, 0x68, 0x2c, 0xcf, 0xb2,
	0x1f, 0x68, 0x6f, 0x5a, 0xe2, 0x26, 0x37, 0xe9, 0xfb, 0x25, 0xf1, 0x09, 0x9e, 0x95, 0xb7, 0x9f,
	0x94, 0xa6, 0xc5, 0xc6, 0xce, 0x07, 0x8c, 0xc8, 0x0b, 0xe5, 0xed, 0x6b, 0xc6, 0x33, 0x2e, 0x0b,
	0xba, 0x3b, 0x3e, 0xff, 0x33, 0x78, 0x1c, 0xf2, 0x9c, 0x34, 0x7f, 0xc9, 0x50, 0xe4, 0xef, 0x00,
	0x33, 0xba, 0x68, 0x28, 0x8e, 0x56, 0x37, 0x4c, 0xb0, 0xb6, 0x94, 0xf5, 0x46, 0x7a, 0xcd, 0x3f,
	0xe0, 0xe4, 0x03, 0x4e, 0x57, 0xb3, 0x8c, 0x64, 0x75, 0xf3, 0x20, 0x58, 0x5b, 0x4b, 0xd8, 0xd0,
	0x60, 0x35, 0xe7, 0x50, 0xce, 0xa8, 0x4d, 0x53, 0xe4, 0x24, 0xd7, 0xeb, 0xd2, 0x55, 0xb9, 0x4b,
	0x52, 0x17, 0xb3, 0x1e, 0x2d, 0x05, 0x6b, 0x0b, 0x09, 0x3b, 0xea, 0x35, 0x7f, 0x83, 0xda, 0xd2,
	0x48, 0xbf, 0x2a, 0x18, 0xdd, 0x54, 0x82, 0xb5, 0x47, 0x79, 0xb4, 0x34, 0xe4, 0xfe, 0xdc, 0xc3,
	0xd3, 0xae, 0x48, 0x1e, 0x1d, 0x19, 0xfe, 0x02, 0x87, 0x90, 0xdc, 0xee, 0x57, 0xc8, 0x2a, 0x24,
	0xd7, 0x6b, 0x2e, 0xa0, 0x32, 0x21, 0x60, 0xc8, 0x56, 0xa7, 0x2f, 0xe8, 0x94, 0xb7, 0xdd, 0xcf,
	0x4a, 0xe4, 0x3d, 0x98, 0x0e, 0xf9, 0xeb, 0xef, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0xb4,
	0xb7, 0xfe, 0x20, 0x01, 0x00, 0x00,
}
