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
	ContestId  int64  `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	ProblemSid string `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	LanguageId int64  `protobuf:"varint,4,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	IsShared   bool   `protobuf:"varint,5,opt,name=is_shared,json=isShared" json:"is_shared,omitempty"`
}

func (m *SubmitRequest) Reset()                    { *m = SubmitRequest{} }
func (m *SubmitRequest) String() string            { return proto.CompactTextString(m) }
func (*SubmitRequest) ProtoMessage()               {}
func (*SubmitRequest) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

func (m *SubmitRequest) GetContestId() int64 {
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
func (*SubmitResponse) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{1} }

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

func init() { proto.RegisterFile("api/submit.proto", fileDescriptor20) }

var fileDescriptor20 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xdd, 0x2e, 0xdb, 0x59, 0x14, 0x09, 0x08, 0x45, 0x11, 0xcb, 0x9e, 0x7a, 0xaa,
	0xa0, 0xcf, 0xe0, 0xa1, 0xd7, 0xf4, 0x01, 0x4a, 0xda, 0x84, 0x75, 0x60, 0x9b, 0x89, 0x99, 0xe4,
	0x81, 0x7c, 0x53, 0x69, 0x76, 0x8b, 0xb7, 0xe4, 0xfb, 0x67, 0x7e, 0xbe, 0x81, 0x47, 0xed, 0xf1,
	0x9d, 0xd3, 0xb4, 0x60, 0xec, 0x7c, 0xa0, 0x48, 0xb2, 0xd0, 0x1e, 0x9f, 0x33, 0x9e, 0x69, 0x59,
	0xc8, 0x5d, 0xf1, 0xe9, 0x57, 0xc0, 0xfd, 0x90, 0xe7, 0x94, 0xfd, 0x49, 0x96, 0xa3, 0x7c, 0x05,
	0x98, 0xc9, 0x45, 0xcb, 0x71, 0x44, 0x53, 0x8b, 0x46, 0xb4, 0x85, 0xaa, 0x6e, 0xa4, 0x37, 0xf2,
	0x0d, 0x8e, 0x3e, 0xd0, 0x74, 0xb1, 0xcb, 0xc8, 0x68, 0xea, 0xbb, 0x46, 0xb4, 0x95, 0x82, 0x1b,
	0x1a, 0xd0, 0x48, 0x09, 0xbb, 0x99, 0x8c, 0xad, 0x8b, 0x9c, 0xe4, 0xf7, 0xba, 0x74, 0xd1, 0xee,
	0x9c, 0xf4, 0xd9, 0xae, 0xa5, 0xbb, 0x5c, 0x0a, 0x1b, 0xea, 0x8d, 0x7c, 0x81, 0x0a, 0x79, 0xe4,
	0x6f, 0x1d, 0xac, 0xa9, 0xcb, 0x46, 0xb4, 0x07, 0x75, 0x40, 0x1e, 0xf2, 0xff, 0xd4, 0xc3, 0xc3,
	0xa6, 0xc8, 0x9e, 0x1c, 0x5b, 0xf9, 0x04, 0xfb, 0x90, 0xdc, 0xbf, 0x5f, 0x19, 0x92, 0xeb, 0x8d,
	0x6c, 0xa0, 0xb4, 0x21, 0x50, 0xc8, 0x56, 0xc7, 0x0f, 0xe8, 0xb4, 0xc7, 0xee, 0x6b, 0x25, 0xea,
	0x1a, 0x4c, 0xfb, 0x7c, 0xf5, 0xe7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x50, 0xdf, 0xe6,
	0x20, 0x01, 0x00, 0x00,
}
