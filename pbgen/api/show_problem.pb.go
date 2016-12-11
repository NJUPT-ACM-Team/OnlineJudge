// Code generated by protoc-gen-go.
// source: api/show_problem.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ShowProblemRequest struct {
	ContestId  uint64 `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	ProblemSid string `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
}

func (m *ShowProblemRequest) Reset()                    { *m = ShowProblemRequest{} }
func (m *ShowProblemRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowProblemRequest) ProtoMessage()               {}
func (*ShowProblemRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *ShowProblemRequest) GetContestId() uint64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *ShowProblemRequest) GetProblemSid() string {
	if m != nil {
		return m.ProblemSid
	}
	return ""
}

// response
type ShowProblemResponse struct {
	ContestId     uint64      `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	ProblemSid    string      `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
	Title         string      `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	TimeLimit     int32       `protobuf:"varint,4,opt,name=time_limit,json=timeLimit" json:"time_limit,omitempty"`
	CaseTimeLimit int32       `protobuf:"varint,5,opt,name=case_time_limit,json=caseTimeLimit" json:"case_time_limit,omitempty"`
	MemoryLimit   int32       `protobuf:"varint,6,opt,name=memory_limit,json=memoryLimit" json:"memory_limit,omitempty"`
	Description   string      `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	Input         string      `protobuf:"bytes,8,opt,name=input" json:"input,omitempty"`
	Output        string      `protobuf:"bytes,9,opt,name=output" json:"output,omitempty"`
	SampleInput   string      `protobuf:"bytes,10,opt,name=sample_input,json=sampleInput" json:"sample_input,omitempty"`
	SampleOutput  string      `protobuf:"bytes,11,opt,name=sample_output,json=sampleOutput" json:"sample_output,omitempty"`
	Source        string      `protobuf:"bytes,12,opt,name=source" json:"source,omitempty"`
	Hint          string      `protobuf:"bytes,13,opt,name=hint" json:"hint,omitempty"`
	Hide          bool        `protobuf:"varint,14,opt,name=hide" json:"hide,omitempty"`
	Languages     []*Language `protobuf:"bytes,15,rep,name=languages" json:"languages,omitempty"`
	Error         *Error      `protobuf:"bytes,16,opt,name=error" json:"error,omitempty"`
}

func (m *ShowProblemResponse) Reset()                    { *m = ShowProblemResponse{} }
func (m *ShowProblemResponse) String() string            { return proto.CompactTextString(m) }
func (*ShowProblemResponse) ProtoMessage()               {}
func (*ShowProblemResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ShowProblemResponse) GetContestId() uint64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *ShowProblemResponse) GetProblemSid() string {
	if m != nil {
		return m.ProblemSid
	}
	return ""
}

func (m *ShowProblemResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ShowProblemResponse) GetTimeLimit() int32 {
	if m != nil {
		return m.TimeLimit
	}
	return 0
}

func (m *ShowProblemResponse) GetCaseTimeLimit() int32 {
	if m != nil {
		return m.CaseTimeLimit
	}
	return 0
}

func (m *ShowProblemResponse) GetMemoryLimit() int32 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *ShowProblemResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ShowProblemResponse) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ShowProblemResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *ShowProblemResponse) GetSampleInput() string {
	if m != nil {
		return m.SampleInput
	}
	return ""
}

func (m *ShowProblemResponse) GetSampleOutput() string {
	if m != nil {
		return m.SampleOutput
	}
	return ""
}

func (m *ShowProblemResponse) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ShowProblemResponse) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *ShowProblemResponse) GetHide() bool {
	if m != nil {
		return m.Hide
	}
	return false
}

func (m *ShowProblemResponse) GetLanguages() []*Language {
	if m != nil {
		return m.Languages
	}
	return nil
}

func (m *ShowProblemResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowProblemRequest)(nil), "api.ShowProblemRequest")
	proto.RegisterType((*ShowProblemResponse)(nil), "api.ShowProblemResponse")
}

func init() { proto.RegisterFile("api/show_problem.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x8a, 0xdb, 0x30,
	0x10, 0x87, 0x71, 0x1d, 0xbb, 0xf1, 0x38, 0x6e, 0x82, 0x5a, 0x82, 0x08, 0x94, 0xaa, 0x29, 0x14,
	0x43, 0x21, 0x85, 0xf4, 0x19, 0x7a, 0x08, 0x04, 0x76, 0x71, 0x72, 0x37, 0x8e, 0x2d, 0x12, 0x81,
	0x65, 0x69, 0x25, 0x99, 0xb0, 0xef, 0xbe, 0x87, 0x45, 0x7f, 0x76, 0x37, 0x7b, 0xde, 0x9b, 0xe6,
	0x9b, 0x6f, 0x7e, 0x9a, 0xc3, 0xc0, 0xb2, 0x91, 0xec, 0xaf, 0xbe, 0x88, 0x6b, 0x2d, 0x95, 0x38,
	0xf5, 0x94, 0x6f, 0xa4, 0x12, 0x46, 0xa0, 0xb8, 0x91, 0x6c, 0xb5, 0xb0, 0xcd, 0x56, 0x70, 0x2e,
	0x06, 0x8f, 0x57, 0xc8, 0x92, 0xbe, 0x19, 0xce, 0x63, 0x73, 0xa6, 0x9e, 0xad, 0x8f, 0x80, 0x0e,
	0x17, 0x71, 0xbd, 0xf7, 0xf3, 0x15, 0x7d, 0x18, 0xa9, 0x36, 0xe8, 0x3b, 0x40, 0x2b, 0x06, 0x43,
	0xb5, 0xa9, 0x59, 0x87, 0x23, 0x12, 0x95, 0x93, 0x2a, 0x0b, 0x64, 0xd7, 0xa1, 0x1f, 0x90, 0x87,
	0x0f, 0x6b, 0xcd, 0x3a, 0xfc, 0x89, 0x44, 0x65, 0x56, 0x41, 0x40, 0x07, 0xd6, 0xad, 0x9f, 0x62,
	0xf8, 0xfa, 0x2e, 0x56, 0x4b, 0x31, 0x68, 0xfa, 0xd1, 0x5c, 0xf4, 0x0d, 0x12, 0xc3, 0x4c, 0x4f,
	0x71, 0xec, 0x5a, 0xbe, 0xb0, 0xa9, 0x86, 0x71, 0x5a, 0xf7, 0x8c, 0x33, 0x83, 0x27, 0x24, 0x2a,
	0x93, 0x2a, 0xb3, 0x64, 0x6f, 0x01, 0xfa, 0x0d, 0xf3, 0xb6, 0xd1, 0xb4, 0xbe, 0x71, 0x12, 0xe7,
	0x14, 0x16, 0x1f, 0x5f, 0xbd, 0x9f, 0x30, 0xe3, 0x94, 0x0b, 0xf5, 0x18, 0xa4, 0xd4, 0x49, 0xb9,
	0x67, 0x5e, 0x21, 0x90, 0x77, 0x54, 0xb7, 0x8a, 0x49, 0xc3, 0xc4, 0x80, 0x3f, 0xbb, 0x2d, 0x6e,
	0x91, 0xdd, 0x90, 0x0d, 0x72, 0x34, 0x78, 0xea, 0x37, 0x74, 0x05, 0x5a, 0x42, 0x2a, 0x46, 0x63,
	0x71, 0xe6, 0x70, 0xa8, 0xec, 0x97, 0xba, 0xe1, 0xb2, 0xa7, 0xb5, 0x1f, 0x02, 0x1f, 0xe8, 0xd9,
	0xce, 0x8d, 0xfe, 0x82, 0x22, 0x28, 0x21, 0x21, 0x77, 0x4e, 0x98, 0xbb, 0xf3, 0x39, 0x4b, 0x48,
	0xb5, 0x18, 0x55, 0x4b, 0xf1, 0xcc, 0xe7, 0xfb, 0x0a, 0x21, 0x98, 0x5c, 0xd8, 0x60, 0x70, 0xe1,
	0xa8, 0x7b, 0x7b, 0xd6, 0x51, 0xfc, 0x85, 0x44, 0xe5, 0xb4, 0x72, 0x6f, 0xf4, 0x07, 0xb2, 0x97,
	0xbb, 0xd0, 0x78, 0x4e, 0xe2, 0x32, 0xdf, 0x16, 0x9b, 0x46, 0xb2, 0xcd, 0x3e, 0xd0, 0xea, 0xad,
	0x8f, 0x08, 0x24, 0x54, 0x29, 0xa1, 0xf0, 0x82, 0x44, 0x65, 0xbe, 0x05, 0x27, 0xfe, 0xb7, 0xa4,
	0xf2, 0x8d, 0x53, 0xea, 0x6e, 0xeb, 0xdf, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x87, 0x56,
	0xf3, 0xa0, 0x02, 0x00, 0x00,
}
