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
	ContestId    uint64 `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	ProblemSid   string `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
	Code         string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	LanguageCode int32  `protobuf:"varint,4,opt,name=language_code,json=languageCode" json:"language_code,omitempty"`
	IsShared     bool   `protobuf:"varint,5,opt,name=is_shared,json=isShared" json:"is_shared,omitempty"`
	IpAddr       string `protobuf:"bytes,6,opt,name=ip_addr,json=ipAddr" json:"ip_addr,omitempty"`
}

func (m *SubmitRequest) Reset()                    { *m = SubmitRequest{} }
func (m *SubmitRequest) String() string            { return proto.CompactTextString(m) }
func (*SubmitRequest) ProtoMessage()               {}
func (*SubmitRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

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

func (m *SubmitRequest) GetLanguageCode() int32 {
	if m != nil {
		return m.LanguageCode
	}
	return 0
}

func (m *SubmitRequest) GetIsShared() bool {
	if m != nil {
		return m.IsShared
	}
	return false
}

func (m *SubmitRequest) GetIpAddr() string {
	if m != nil {
		return m.IpAddr
	}
	return ""
}

type SubmitResponse struct {
	RunId int64  `protobuf:"varint,1,opt,name=run_id,json=runId" json:"run_id,omitempty"`
	Error *Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SubmitResponse) Reset()                    { *m = SubmitResponse{} }
func (m *SubmitResponse) String() string            { return proto.CompactTextString(m) }
func (*SubmitResponse) ProtoMessage()               {}
func (*SubmitResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

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

func init() { proto.RegisterFile("api/submit.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x34, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x89, 0xfd, 0xe3, 0x76, 0xd6, 0x15, 0x09, 0x88, 0x45, 0x11, 0xcb, 0x7a, 0xe9, 0xa9,
	0x82, 0x3e, 0x81, 0x88, 0x87, 0x5e, 0xd3, 0x07, 0x28, 0x69, 0x27, 0xac, 0x03, 0xdb, 0x24, 0x4e,
	0xda, 0xd7, 0xf3, 0xd9, 0x64, 0xb3, 0xdb, 0xdb, 0xcc, 0xef, 0xfb, 0x18, 0x7e, 0x0c, 0xdc, 0x69,
	0x4f, 0x6f, 0x61, 0x19, 0x26, 0x9a, 0x1b, 0xcf, 0x6e, 0x76, 0x32, 0xd1, 0x9e, 0x1e, 0x23, 0x1e,
	0xdd, 0x34, 0x39, 0x7b, 0xc6, 0xfb, 0x3f, 0x01, 0xbb, 0x2e, 0xf6, 0x94, 0xf9, 0x5d, 0x4c, 0x98,
	0xe5, 0x33, 0xc0, 0xe8, 0xec, 0x6c, 0xc2, 0xdc, 0x13, 0x96, 0xa2, 0x12, 0x75, 0xaa, 0x8a, 0x0b,
	0x69, 0x51, 0xbe, 0xc0, 0xd6, 0xb3, 0x1b, 0x8e, 0x66, 0xea, 0x03, 0x61, 0x79, 0x55, 0x89, 0xba,
	0x50, 0x70, 0x41, 0x1d, 0xa1, 0x94, 0x90, 0x8e, 0x0e, 0x4d, 0x99, 0xc4, 0x24, 0xce, 0xf2, 0x15,
	0x76, 0x47, 0x6d, 0x0f, 0x8b, 0x3e, 0x98, 0x3e, 0x86, 0x69, 0x25, 0xea, 0x4c, 0xdd, 0xac, 0xf0,
	0xeb, 0x54, 0x7a, 0x82, 0x82, 0x42, 0x1f, 0x7e, 0x34, 0x1b, 0x2c, 0xb3, 0x4a, 0xd4, 0x1b, 0xb5,
	0xa1, 0xd0, 0xc5, 0x5d, 0x3e, 0xc0, 0x35, 0xf9, 0x5e, 0x23, 0x72, 0x99, 0xc7, 0xc3, 0x39, 0xf9,
	0x4f, 0x44, 0xde, 0xb7, 0x70, 0xbb, 0xfa, 0x07, 0xef, 0x6c, 0x30, 0xf2, 0x1e, 0x72, 0x5e, 0xec,
	0x2a, 0x9f, 0xa8, 0x8c, 0x17, 0xdb, 0xa2, 0xac, 0x20, 0x33, 0xcc, 0x8e, 0xa3, 0xf2, 0xf6, 0x1d,
	0x1a, 0xed, 0xa9, 0xf9, 0x3e, 0x11, 0x75, 0x0e, 0x86, 0x3c, 0xbe, 0xe4, 0xe3, 0x3f, 0x00, 0x00,
	0xff, 0xff, 0x39, 0x1f, 0x40, 0x6d, 0x3d, 0x01, 0x00, 0x00,
}
