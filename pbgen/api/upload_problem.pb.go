// Code generated by protoc-gen-go.
// source: api/upload_problem.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UploadProblemRequest struct {
	Title         string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	TimeLimit     int32  `protobuf:"varint,2,opt,name=time_limit,json=timeLimit" json:"time_limit,omitempty"`
	CaseTimeLimit int32  `protobuf:"varint,3,opt,name=case_time_limit,json=caseTimeLimit" json:"case_time_limit,omitempty"`
	Description   string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Input         string `protobuf:"bytes,5,opt,name=input" json:"input,omitempty"`
	Output        string `protobuf:"bytes,6,opt,name=output" json:"output,omitempty"`
	SampleInput   string `protobuf:"bytes,7,opt,name=sample_input,json=sampleInput" json:"sample_input,omitempty"`
	SampleOutput  string `protobuf:"bytes,8,opt,name=sample_output,json=sampleOutput" json:"sample_output,omitempty"`
	Source        string `protobuf:"bytes,9,opt,name=source" json:"source,omitempty"`
	Hint          string `protobuf:"bytes,10,opt,name=hint" json:"hint,omitempty"`
}

func (m *UploadProblemRequest) Reset()                    { *m = UploadProblemRequest{} }
func (m *UploadProblemRequest) String() string            { return proto.CompactTextString(m) }
func (*UploadProblemRequest) ProtoMessage()               {}
func (*UploadProblemRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *UploadProblemRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UploadProblemRequest) GetTimeLimit() int32 {
	if m != nil {
		return m.TimeLimit
	}
	return 0
}

func (m *UploadProblemRequest) GetCaseTimeLimit() int32 {
	if m != nil {
		return m.CaseTimeLimit
	}
	return 0
}

func (m *UploadProblemRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UploadProblemRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *UploadProblemRequest) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *UploadProblemRequest) GetSampleInput() string {
	if m != nil {
		return m.SampleInput
	}
	return ""
}

func (m *UploadProblemRequest) GetSampleOutput() string {
	if m != nil {
		return m.SampleOutput
	}
	return ""
}

func (m *UploadProblemRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *UploadProblemRequest) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

type UploadProblemResponse struct {
	LocalPid           int64  `protobuf:"varint,1,opt,name=local_pid,json=localPid" json:"local_pid,omitempty"`
	CheckTitle         string `protobuf:"bytes,2,opt,name=check_title,json=checkTitle" json:"check_title,omitempty"`
	CheckTimeLimit     string `protobuf:"bytes,3,opt,name=check_time_limit,json=checkTimeLimit" json:"check_time_limit,omitempty"`
	CheckCaseTimeLimit string `protobuf:"bytes,4,opt,name=check_case_time_limit,json=checkCaseTimeLimit" json:"check_case_time_limit,omitempty"`
	CheckDescription   string `protobuf:"bytes,5,opt,name=check_description,json=checkDescription" json:"check_description,omitempty"`
	CheckInput         string `protobuf:"bytes,6,opt,name=check_input,json=checkInput" json:"check_input,omitempty"`
	CheckOutput        string `protobuf:"bytes,7,opt,name=check_output,json=checkOutput" json:"check_output,omitempty"`
	CheckSampleInput   string `protobuf:"bytes,8,opt,name=check_sample_input,json=checkSampleInput" json:"check_sample_input,omitempty"`
	CheckSampleOutput  string `protobuf:"bytes,9,opt,name=check_sample_output,json=checkSampleOutput" json:"check_sample_output,omitempty"`
	CheckSource        string `protobuf:"bytes,10,opt,name=check_source,json=checkSource" json:"check_source,omitempty"`
	CheckInt           string `protobuf:"bytes,11,opt,name=check_int,json=checkInt" json:"check_int,omitempty"`
	Error              *Error `protobuf:"bytes,12,opt,name=error" json:"error,omitempty"`
}

func (m *UploadProblemResponse) Reset()                    { *m = UploadProblemResponse{} }
func (m *UploadProblemResponse) String() string            { return proto.CompactTextString(m) }
func (*UploadProblemResponse) ProtoMessage()               {}
func (*UploadProblemResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *UploadProblemResponse) GetLocalPid() int64 {
	if m != nil {
		return m.LocalPid
	}
	return 0
}

func (m *UploadProblemResponse) GetCheckTitle() string {
	if m != nil {
		return m.CheckTitle
	}
	return ""
}

func (m *UploadProblemResponse) GetCheckTimeLimit() string {
	if m != nil {
		return m.CheckTimeLimit
	}
	return ""
}

func (m *UploadProblemResponse) GetCheckCaseTimeLimit() string {
	if m != nil {
		return m.CheckCaseTimeLimit
	}
	return ""
}

func (m *UploadProblemResponse) GetCheckDescription() string {
	if m != nil {
		return m.CheckDescription
	}
	return ""
}

func (m *UploadProblemResponse) GetCheckInput() string {
	if m != nil {
		return m.CheckInput
	}
	return ""
}

func (m *UploadProblemResponse) GetCheckOutput() string {
	if m != nil {
		return m.CheckOutput
	}
	return ""
}

func (m *UploadProblemResponse) GetCheckSampleInput() string {
	if m != nil {
		return m.CheckSampleInput
	}
	return ""
}

func (m *UploadProblemResponse) GetCheckSampleOutput() string {
	if m != nil {
		return m.CheckSampleOutput
	}
	return ""
}

func (m *UploadProblemResponse) GetCheckSource() string {
	if m != nil {
		return m.CheckSource
	}
	return ""
}

func (m *UploadProblemResponse) GetCheckInt() string {
	if m != nil {
		return m.CheckInt
	}
	return ""
}

func (m *UploadProblemResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadProblemRequest)(nil), "api.UploadProblemRequest")
	proto.RegisterType((*UploadProblemResponse)(nil), "api.UploadProblemResponse")
}

func init() { proto.RegisterFile("api/upload_problem.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x93, 0x6f, 0x6f, 0x94, 0x40,
	0x10, 0xc6, 0xc3, 0x51, 0x10, 0x86, 0xab, 0xb6, 0x6b, 0xdb, 0x6c, 0x6c, 0x8c, 0xb4, 0x26, 0x86,
	0x44, 0x83, 0x51, 0x3f, 0x82, 0xfa, 0xc2, 0xc4, 0xa4, 0x0d, 0xad, 0xaf, 0x09, 0x85, 0x4d, 0x6e,
	0x23, 0xb0, 0x2b, 0x2c, 0x9f, 0xcd, 0x0f, 0xe4, 0x17, 0x31, 0x3b, 0xb3, 0x77, 0x59, 0xee, 0xdd,
	0xcd, 0x33, 0xcf, 0xfc, 0xd9, 0xdf, 0x1c, 0xc0, 0x1b, 0x2d, 0x3f, 0x2e, 0xba, 0x57, 0x4d, 0x57,
	0xeb, 0x49, 0x3d, 0xf5, 0x62, 0x28, 0xf5, 0xa4, 0x8c, 0x62, 0x61, 0xa3, 0xe5, 0xab, 0x33, 0x9b,
	0x6e, 0xd5, 0x30, 0xa8, 0x91, 0xe4, 0xdb, 0xbf, 0x1b, 0xb8, 0xf8, 0x85, 0xfe, 0x7b, 0xb2, 0x57,
	0xe2, 0xcf, 0x22, 0x66, 0xc3, 0x2e, 0x20, 0x32, 0xd2, 0xf4, 0x82, 0x07, 0x79, 0x50, 0xa4, 0x15,
	0x05, 0xec, 0x35, 0x80, 0x91, 0x83, 0xa8, 0x7b, 0x39, 0x48, 0xc3, 0x37, 0x79, 0x50, 0x44, 0x55,
	0x6a, 0x95, 0x9f, 0x56, 0x60, 0xef, 0xe0, 0x45, 0xdb, 0xcc, 0xa2, 0xf6, 0x3c, 0x21, 0x7a, 0x4e,
	0xad, 0xfc, 0x78, 0xf0, 0xe5, 0x90, 0x75, 0x62, 0x6e, 0x27, 0xa9, 0x8d, 0x54, 0x23, 0x3f, 0xc1,
	0x11, 0xbe, 0x64, 0xc7, 0xcb, 0x51, 0x2f, 0x86, 0x47, 0x34, 0x1e, 0x03, 0x76, 0x05, 0xb1, 0x5a,
	0x8c, 0x95, 0x63, 0x94, 0x5d, 0xc4, 0x6e, 0x60, 0x3b, 0x37, 0x83, 0xee, 0x45, 0x4d, 0x45, 0xcf,
	0xa8, 0x21, 0x69, 0x3f, 0xb0, 0xf4, 0x2d, 0x9c, 0x3a, 0x8b, 0xeb, 0x90, 0xa0, 0xc7, 0xd5, 0xdd,
	0x51, 0x9f, 0x2b, 0x88, 0x67, 0xb5, 0x4c, 0xad, 0xe0, 0x29, 0xf5, 0xa7, 0x88, 0x31, 0x38, 0xd9,
	0xc9, 0xd1, 0x70, 0x40, 0x15, 0x7f, 0xdf, 0xfe, 0x0b, 0xe1, 0xf2, 0x88, 0xdc, 0xac, 0xd5, 0x38,
	0x0b, 0x76, 0x0d, 0x69, 0xaf, 0xda, 0xa6, 0xaf, 0xb5, 0xec, 0x10, 0x5f, 0x58, 0x25, 0x28, 0xdc,
	0xcb, 0x8e, 0xbd, 0x81, 0xac, 0xdd, 0x89, 0xf6, 0x77, 0x4d, 0x74, 0x37, 0xd8, 0x11, 0x50, 0x7a,
	0x44, 0xc4, 0x05, 0x9c, 0xed, 0x0d, 0x2b, 0x88, 0x69, 0xf5, 0xdc, 0xb9, 0xf6, 0x14, 0x3f, 0xc1,
	0x25, 0x39, 0x8f, 0x99, 0x13, 0x4f, 0x86, 0xc9, 0xaf, 0x2b, 0xf0, 0xef, 0xe1, 0x9c, 0x4a, 0x7c,
	0xfc, 0x84, 0x98, 0xa6, 0x7e, 0xf3, 0x6e, 0x70, 0x58, 0x95, 0xa0, 0xc6, 0xde, 0xaa, 0xc4, 0xf4,
	0x06, 0xb6, 0x64, 0x70, 0x48, 0x1d, 0x76, 0xd4, 0x1c, 0xd1, 0x0f, 0x40, 0x6b, 0xd4, 0xab, 0xfb,
	0x24, 0xde, 0xc4, 0x07, 0xef, 0x48, 0x25, 0xbc, 0x5c, 0xb9, 0x5d, 0x5f, 0x3a, 0xc6, 0xb9, 0x67,
	0xbf, 0x3b, 0xdc, 0xdd, 0xf9, 0xe9, 0x6a, 0xe0, 0x2d, 0xf0, 0x40, 0xa7, 0xbb, 0x86, 0x74, 0xff,
	0x08, 0xc3, 0x33, 0xcc, 0x27, 0xee, 0x09, 0xf6, 0x7f, 0x18, 0x89, 0x69, 0x52, 0x13, 0xdf, 0xe6,
	0x41, 0x91, 0x7d, 0x86, 0xb2, 0xd1, 0xb2, 0xfc, 0x6e, 0x95, 0x8a, 0x12, 0x4f, 0x31, 0x7e, 0x26,
	0x5f, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xce, 0x6b, 0x49, 0x1f, 0x59, 0x03, 0x00, 0x00,
}
