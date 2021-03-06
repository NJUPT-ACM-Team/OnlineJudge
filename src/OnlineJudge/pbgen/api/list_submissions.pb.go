// Code generated by protoc-gen-go.
// source: api/list_submissions.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListSubmissionsRequest struct {
	PerPage          int32  `protobuf:"varint,1,opt,name=per_page,json=perPage" json:"per_page,omitempty"`
	CurrentPage      int32  `protobuf:"varint,2,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	FilterRunId      int64  `protobuf:"varint,3,opt,name=filter_run_id,json=filterRunId" json:"filter_run_id,omitempty"`
	FilterUsername   string `protobuf:"bytes,4,opt,name=filter_username,json=filterUsername" json:"filter_username,omitempty"`
	FilterOj         string `protobuf:"bytes,5,opt,name=filter_oj,json=filterOj" json:"filter_oj,omitempty"`
	FilterPid        string `protobuf:"bytes,6,opt,name=filter_pid,json=filterPid" json:"filter_pid,omitempty"`
	FilterStatusCode string `protobuf:"bytes,7,opt,name=filter_status_code,json=filterStatusCode" json:"filter_status_code,omitempty"`
	FilterLanguage   string `protobuf:"bytes,8,opt,name=filter_language,json=filterLanguage" json:"filter_language,omitempty"`
	FilterCompiler   string `protobuf:"bytes,9,opt,name=filter_compiler,json=filterCompiler" json:"filter_compiler,omitempty"`
	IsDesc           bool   `protobuf:"varint,10,opt,name=is_desc,json=isDesc" json:"is_desc,omitempty"`
}

func (m *ListSubmissionsRequest) Reset()                    { *m = ListSubmissionsRequest{} }
func (m *ListSubmissionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSubmissionsRequest) ProtoMessage()               {}
func (*ListSubmissionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *ListSubmissionsRequest) GetPerPage() int32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *ListSubmissionsRequest) GetCurrentPage() int32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *ListSubmissionsRequest) GetFilterRunId() int64 {
	if m != nil {
		return m.FilterRunId
	}
	return 0
}

func (m *ListSubmissionsRequest) GetFilterUsername() string {
	if m != nil {
		return m.FilterUsername
	}
	return ""
}

func (m *ListSubmissionsRequest) GetFilterOj() string {
	if m != nil {
		return m.FilterOj
	}
	return ""
}

func (m *ListSubmissionsRequest) GetFilterPid() string {
	if m != nil {
		return m.FilterPid
	}
	return ""
}

func (m *ListSubmissionsRequest) GetFilterStatusCode() string {
	if m != nil {
		return m.FilterStatusCode
	}
	return ""
}

func (m *ListSubmissionsRequest) GetFilterLanguage() string {
	if m != nil {
		return m.FilterLanguage
	}
	return ""
}

func (m *ListSubmissionsRequest) GetFilterCompiler() string {
	if m != nil {
		return m.FilterCompiler
	}
	return ""
}

func (m *ListSubmissionsRequest) GetIsDesc() bool {
	if m != nil {
		return m.IsDesc
	}
	return false
}

type ListSubmissionsResponse struct {
	Lines       []*ListSubmissionsResponse_PerLine `protobuf:"bytes,1,rep,name=lines" json:"lines,omitempty"`
	TotalLines  int32                              `protobuf:"varint,2,opt,name=total_lines,json=totalLines" json:"total_lines,omitempty"`
	TotalPages  int32                              `protobuf:"varint,3,opt,name=total_pages,json=totalPages" json:"total_pages,omitempty"`
	CurrentPage int32                              `protobuf:"varint,4,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	Error       *Error                             `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
}

func (m *ListSubmissionsResponse) Reset()                    { *m = ListSubmissionsResponse{} }
func (m *ListSubmissionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSubmissionsResponse) ProtoMessage()               {}
func (*ListSubmissionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *ListSubmissionsResponse) GetLines() []*ListSubmissionsResponse_PerLine {
	if m != nil {
		return m.Lines
	}
	return nil
}

func (m *ListSubmissionsResponse) GetTotalLines() int32 {
	if m != nil {
		return m.TotalLines
	}
	return 0
}

func (m *ListSubmissionsResponse) GetTotalPages() int32 {
	if m != nil {
		return m.TotalPages
	}
	return 0
}

func (m *ListSubmissionsResponse) GetCurrentPage() int32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *ListSubmissionsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ListSubmissionsResponse_PerLine struct {
	Sid             string    `protobuf:"bytes,1,opt,name=sid" json:"sid,omitempty"`
	RunId           int64     `protobuf:"varint,2,opt,name=run_id,json=runId" json:"run_id,omitempty"`
	Username        string    `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	Status          string    `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	StatusCode      string    `protobuf:"bytes,5,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	CeInfo          string    `protobuf:"bytes,6,opt,name=ce_info,json=ceInfo" json:"ce_info,omitempty"`
	Language        *Language `protobuf:"bytes,7,opt,name=language" json:"language,omitempty"`
	TimeUsed        int32     `protobuf:"varint,8,opt,name=time_used,json=timeUsed" json:"time_used,omitempty"`
	MemoryUsed      int32     `protobuf:"varint,9,opt,name=memory_used,json=memoryUsed" json:"memory_used,omitempty"`
	Testcases       int32     `protobuf:"varint,10,opt,name=testcases" json:"testcases,omitempty"`
	TestcasesPassed int32     `protobuf:"varint,11,opt,name=testcases_passed,json=testcasesPassed" json:"testcases_passed,omitempty"`
	CodeLength      int32     `protobuf:"varint,12,opt,name=code_length,json=codeLength" json:"code_length,omitempty"`
	SubmitTime      string    `protobuf:"bytes,13,opt,name=submit_time,json=submitTime" json:"submit_time,omitempty"`
	Code            string    `protobuf:"bytes,14,opt,name=code" json:"code,omitempty"`
	IsSpj           bool      `protobuf:"varint,15,opt,name=is_spj,json=isSpj" json:"is_spj,omitempty"`
}

func (m *ListSubmissionsResponse_PerLine) Reset()         { *m = ListSubmissionsResponse_PerLine{} }
func (m *ListSubmissionsResponse_PerLine) String() string { return proto.CompactTextString(m) }
func (*ListSubmissionsResponse_PerLine) ProtoMessage()    {}
func (*ListSubmissionsResponse_PerLine) Descriptor() ([]byte, []int) {
	return fileDescriptor12, []int{1, 0}
}

func (m *ListSubmissionsResponse_PerLine) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *ListSubmissionsResponse_PerLine) GetRunId() int64 {
	if m != nil {
		return m.RunId
	}
	return 0
}

func (m *ListSubmissionsResponse_PerLine) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ListSubmissionsResponse_PerLine) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ListSubmissionsResponse_PerLine) GetStatusCode() string {
	if m != nil {
		return m.StatusCode
	}
	return ""
}

func (m *ListSubmissionsResponse_PerLine) GetCeInfo() string {
	if m != nil {
		return m.CeInfo
	}
	return ""
}

func (m *ListSubmissionsResponse_PerLine) GetLanguage() *Language {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *ListSubmissionsResponse_PerLine) GetTimeUsed() int32 {
	if m != nil {
		return m.TimeUsed
	}
	return 0
}

func (m *ListSubmissionsResponse_PerLine) GetMemoryUsed() int32 {
	if m != nil {
		return m.MemoryUsed
	}
	return 0
}

func (m *ListSubmissionsResponse_PerLine) GetTestcases() int32 {
	if m != nil {
		return m.Testcases
	}
	return 0
}

func (m *ListSubmissionsResponse_PerLine) GetTestcasesPassed() int32 {
	if m != nil {
		return m.TestcasesPassed
	}
	return 0
}

func (m *ListSubmissionsResponse_PerLine) GetCodeLength() int32 {
	if m != nil {
		return m.CodeLength
	}
	return 0
}

func (m *ListSubmissionsResponse_PerLine) GetSubmitTime() string {
	if m != nil {
		return m.SubmitTime
	}
	return ""
}

func (m *ListSubmissionsResponse_PerLine) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ListSubmissionsResponse_PerLine) GetIsSpj() bool {
	if m != nil {
		return m.IsSpj
	}
	return false
}

func init() {
	proto.RegisterType((*ListSubmissionsRequest)(nil), "api.ListSubmissionsRequest")
	proto.RegisterType((*ListSubmissionsResponse)(nil), "api.ListSubmissionsResponse")
	proto.RegisterType((*ListSubmissionsResponse_PerLine)(nil), "api.ListSubmissionsResponse.PerLine")
}

func init() { proto.RegisterFile("api/list_submissions.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x94, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x95, 0xba, 0x4e, 0x9c, 0x49, 0xff, 0x44, 0x2b, 0xd1, 0x9a, 0x00, 0x22, 0x54, 0x48,
	0xa4, 0x12, 0x0a, 0x52, 0xb9, 0x71, 0x2d, 0x1c, 0x2a, 0x45, 0x22, 0x72, 0xe9, 0xd9, 0x72, 0xed,
	0x69, 0xd8, 0x28, 0xde, 0x5d, 0x76, 0xd6, 0x07, 0xbe, 0x0c, 0x5f, 0x85, 0xaf, 0xc5, 0x11, 0xed,
	0xec, 0x36, 0x0d, 0x8a, 0xb8, 0xc5, 0xbf, 0xf7, 0xbc, 0x9e, 0x9d, 0xf7, 0x14, 0x98, 0x54, 0x46,
	0x7e, 0xd8, 0x48, 0x72, 0x25, 0x75, 0xf7, 0xad, 0x24, 0x92, 0x5a, 0xd1, 0xdc, 0x58, 0xed, 0xb4,
	0x48, 0x2a, 0x23, 0x27, 0x63, 0x6f, 0xa8, 0x75, 0xdb, 0x6a, 0x15, 0xf0, 0x44, 0xf0, 0x2b, 0x95,
	0x5a, 0x75, 0xd5, 0x0a, 0x03, 0xbb, 0xf8, 0x73, 0x00, 0x67, 0x0b, 0x49, 0xee, 0xf6, 0xe9, 0x90,
	0x02, 0x7f, 0x74, 0x48, 0x4e, 0x3c, 0x87, 0xcc, 0xa0, 0x2d, 0x4d, 0xb5, 0xc2, 0xbc, 0x37, 0xed,
	0xcd, 0xd2, 0x62, 0x60, 0xd0, 0x2e, 0xab, 0x15, 0x8a, 0x37, 0x70, 0x54, 0x77, 0xd6, 0xa2, 0x72,
	0x41, 0x3e, 0x60, 0x79, 0x14, 0x19, 0x5b, 0x2e, 0xe0, 0xf8, 0x41, 0x6e, 0x1c, 0xda, 0xd2, 0x76,
	0xaa, 0x94, 0x4d, 0x9e, 0x4c, 0x7b, 0xb3, 0xa4, 0x18, 0x05, 0x58, 0x74, 0xea, 0xa6, 0x11, 0xef,
	0xe0, 0x34, 0x7a, 0x3a, 0x42, 0xab, 0xaa, 0x16, 0xf3, 0xc3, 0x69, 0x6f, 0x36, 0x2c, 0x4e, 0x02,
	0xbe, 0x8b, 0x54, 0xbc, 0x80, 0x61, 0x34, 0xea, 0x75, 0x9e, 0xb2, 0x25, 0x0b, 0xe0, 0xeb, 0x5a,
	0xbc, 0x02, 0x88, 0xa2, 0x91, 0x4d, 0xde, 0x67, 0x35, 0xda, 0x97, 0xb2, 0x11, 0xef, 0x41, 0x44,
	0x99, 0x5c, 0xe5, 0x3a, 0x2a, 0x6b, 0xdd, 0x60, 0x3e, 0x60, 0xdb, 0x38, 0x28, 0xb7, 0x2c, 0x5c,
	0xeb, 0x06, 0x77, 0x46, 0x7a, 0x5c, 0x54, 0x9e, 0xed, 0x8e, 0xb4, 0x88, 0x74, 0xc7, 0x58, 0xeb,
	0xd6, 0xc8, 0x0d, 0xda, 0x7c, 0xb8, 0x6b, 0xbc, 0x8e, 0x54, 0x9c, 0xc3, 0x40, 0x52, 0xd9, 0x20,
	0xd5, 0x39, 0x4c, 0x7b, 0xb3, 0xac, 0xe8, 0x4b, 0xfa, 0x8c, 0x54, 0x5f, 0xfc, 0x4a, 0xe1, 0x7c,
	0x6f, 0xf5, 0x64, 0xb4, 0x22, 0x14, 0x9f, 0x20, 0xdd, 0x48, 0x85, 0x94, 0xf7, 0xa6, 0xc9, 0x6c,
	0x74, 0xf5, 0x76, 0x5e, 0x19, 0x39, 0xff, 0x8f, 0x79, 0xbe, 0x44, 0xbb, 0x90, 0x0a, 0x8b, 0xf0,
	0x8a, 0x78, 0x0d, 0x23, 0xa7, 0x5d, 0xb5, 0x29, 0xc3, 0x09, 0x21, 0x1b, 0x60, 0xb4, 0xf8, 0xd7,
	0xe0, 0xb3, 0x23, 0x0e, 0xe6, 0xd1, 0xe0, 0xa3, 0xa3, 0xbd, 0x78, 0x0f, 0xf7, 0xe3, 0x9d, 0x42,
	0x8a, 0xd6, 0x6a, 0xcb, 0x69, 0x8c, 0xae, 0x80, 0x07, 0xfc, 0xe2, 0x49, 0x11, 0x84, 0xc9, 0xef,
	0x04, 0x06, 0x71, 0x32, 0x31, 0x86, 0x84, 0x64, 0xc3, 0x2d, 0x1a, 0x16, 0xfe, 0xa7, 0x78, 0x06,
	0xfd, 0xd8, 0x8b, 0x03, 0xee, 0x45, 0x6a, 0xb9, 0x11, 0x13, 0xc8, 0xb6, 0x55, 0x48, 0x42, 0xce,
	0x8f, 0xcf, 0xe2, 0x0c, 0xfa, 0x21, 0xc1, 0x58, 0x92, 0xf8, 0xe4, 0xaf, 0xb3, 0x9b, 0x6c, 0xa8,
	0x07, 0xd0, 0x53, 0xa6, 0xe7, 0x30, 0xa8, 0xb1, 0x94, 0xea, 0x41, 0xc7, 0x76, 0xf4, 0x6b, 0xbc,
	0x51, 0x0f, 0x5a, 0x5c, 0x42, 0xb6, 0x4d, 0x79, 0xc0, 0xf7, 0x38, 0x0e, 0x8b, 0x8e, 0xb0, 0xd8,
	0xca, 0xbe, 0x81, 0x4e, 0xb6, 0xe8, 0x8b, 0xda, 0x70, 0x23, 0xd2, 0x22, 0xf3, 0xe0, 0x8e, 0xb0,
	0xf1, 0x13, 0xb4, 0xd8, 0x6a, 0xfb, 0x33, 0xc8, 0xc3, 0xb0, 0xd0, 0x80, 0xd8, 0xf0, 0x12, 0x86,
	0x0e, 0xc9, 0xd5, 0x15, 0x21, 0x71, 0x0b, 0xd2, 0xe2, 0x09, 0x88, 0x4b, 0x18, 0x6f, 0x1f, 0x4a,
	0x53, 0x91, 0x3f, 0x63, 0xc4, 0xa6, 0xd3, 0x2d, 0x5f, 0x32, 0xf6, 0x5f, 0xf2, 0x97, 0x2c, 0x37,
	0xa8, 0x56, 0xee, 0x7b, 0x7e, 0x14, 0xbe, 0xe4, 0xd1, 0x82, 0x09, 0x2f, 0xc3, 0x57, 0xc4, 0x95,
	0x7e, 0xba, 0xfc, 0x38, 0x2e, 0x83, 0xd1, 0x37, 0xd9, 0xa2, 0x10, 0x70, 0xc8, 0x6b, 0x3a, 0x61,
	0x85, 0x7f, 0xfb, 0x30, 0x24, 0x95, 0x64, 0xd6, 0xf9, 0x29, 0x37, 0x34, 0x95, 0x74, 0x6b, 0xd6,
	0xf7, 0x7d, 0xfe, 0x8b, 0xf8, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x37, 0xe4, 0xe6, 0xcd, 0x6b,
	0x04, 0x00, 0x00,
}
