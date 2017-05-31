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
	FilterUsername   string `protobuf:"bytes,3,opt,name=filter_username,json=filterUsername" json:"filter_username,omitempty"`
	FilterOj         string `protobuf:"bytes,4,opt,name=filter_oj,json=filterOj" json:"filter_oj,omitempty"`
	FilterPid        string `protobuf:"bytes,5,opt,name=filter_pid,json=filterPid" json:"filter_pid,omitempty"`
	FilterStatusCode string `protobuf:"bytes,6,opt,name=filter_status_code,json=filterStatusCode" json:"filter_status_code,omitempty"`
	FilterLanguage   string `protobuf:"bytes,7,opt,name=filter_language,json=filterLanguage" json:"filter_language,omitempty"`
	FilterCompiler   string `protobuf:"bytes,8,opt,name=filter_compiler,json=filterCompiler" json:"filter_compiler,omitempty"`
	IsDesc           bool   `protobuf:"varint,9,opt,name=is_desc,json=isDesc" json:"is_desc,omitempty"`
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
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x95, 0xa6, 0x93, 0x9f, 0x9b, 0xfe, 0x44, 0x96, 0x68, 0x4d, 0x00, 0x11, 0x2a, 0x24,
	0x52, 0x09, 0x05, 0xa9, 0xec, 0xd8, 0x16, 0x16, 0x95, 0x22, 0x11, 0x4d, 0xe9, 0x7a, 0x34, 0x9d,
	0xb9, 0x0d, 0x8e, 0x32, 0xb6, 0xf1, 0xf5, 0x2c, 0x78, 0x19, 0x5e, 0x85, 0x97, 0xe0, 0x81, 0x90,
	0xaf, 0xdd, 0x34, 0xa5, 0x62, 0x17, 0x7f, 0xe7, 0x8c, 0xe7, 0xce, 0x39, 0x37, 0x30, 0x29, 0xad,
	0xfa, 0xb0, 0x51, 0xe4, 0x0b, 0x6a, 0x6f, 0x1b, 0x45, 0xa4, 0x8c, 0xa6, 0xb9, 0x75, 0xc6, 0x1b,
	0xd1, 0x2d, 0xad, 0x9a, 0x8c, 0x83, 0xa1, 0x32, 0x4d, 0x63, 0x74, 0xc4, 0x13, 0xc1, 0x8f, 0x94,
	0x7a, 0xd5, 0x96, 0x2b, 0x8c, 0xec, 0xec, 0xcf, 0x1e, 0x9c, 0x2c, 0x14, 0xf9, 0xeb, 0x87, 0x4b,
	0x72, 0xfc, 0xd1, 0x22, 0x79, 0xf1, 0x1c, 0x06, 0x16, 0x5d, 0x61, 0xcb, 0x15, 0xca, 0xce, 0xb4,
	0x33, 0xcb, 0xf2, 0xbe, 0x45, 0xb7, 0x2c, 0x57, 0x28, 0xde, 0xc0, 0x41, 0xd5, 0x3a, 0x87, 0xda,
	0x47, 0x79, 0x8f, 0xe5, 0x51, 0x62, 0x6c, 0x79, 0x07, 0xc7, 0x77, 0x6a, 0xe3, 0xd1, 0x15, 0x2d,
	0xa1, 0xd3, 0x65, 0x83, 0xb2, 0x3b, 0xed, 0xcc, 0x86, 0xf9, 0x51, 0xc4, 0x37, 0x89, 0x8a, 0x17,
	0x30, 0x4c, 0x46, 0xb3, 0x96, 0xfb, 0x6c, 0x19, 0x44, 0xf0, 0x75, 0x2d, 0x5e, 0x01, 0x24, 0xd1,
	0xaa, 0x5a, 0x66, 0xac, 0x26, 0xfb, 0x52, 0xd5, 0xe2, 0x3d, 0x88, 0x24, 0x93, 0x2f, 0x7d, 0x4b,
	0x45, 0x65, 0x6a, 0x94, 0x3d, 0xb6, 0x8d, 0xa3, 0x72, 0xcd, 0xc2, 0xa5, 0xa9, 0x77, 0x47, 0xba,
	0x0f, 0x41, 0xf6, 0x77, 0x47, 0x5a, 0x24, 0xba, 0x63, 0xac, 0x4c, 0x63, 0xd5, 0x06, 0x9d, 0x1c,
	0xec, 0x1a, 0x2f, 0x13, 0x15, 0xa7, 0xd0, 0x57, 0x54, 0xd4, 0x48, 0x95, 0x1c, 0x4e, 0x3b, 0xb3,
	0x41, 0xde, 0x53, 0xf4, 0x19, 0xa9, 0x3a, 0xfb, 0x95, 0xc1, 0xe9, 0x93, 0x58, 0xc9, 0x1a, 0x4d,
	0x28, 0x3e, 0x41, 0xb6, 0x51, 0x1a, 0x49, 0x76, 0xa6, 0xdd, 0xd9, 0xe8, 0xe2, 0xed, 0xbc, 0xb4,
	0x6a, 0xfe, 0x1f, 0xf3, 0x7c, 0x89, 0x6e, 0xa1, 0x34, 0xe6, 0xf1, 0x11, 0xf1, 0x1a, 0x46, 0xde,
	0xf8, 0x72, 0x53, 0xc4, 0x1b, 0x62, 0xee, 0xc0, 0x68, 0xf1, 0xd8, 0x10, 0x7a, 0x21, 0x8e, 0xfc,
	0xde, 0x10, 0x6a, 0xa1, 0x27, 0xd5, 0xed, 0x3f, 0xad, 0x6e, 0x0a, 0x19, 0x3a, 0x67, 0x1c, 0xe7,
	0x3d, 0xba, 0x00, 0x1e, 0xf0, 0x4b, 0x20, 0x79, 0x14, 0x26, 0xbf, 0xbb, 0xd0, 0x4f, 0x93, 0x89,
	0x31, 0x74, 0x49, 0xd5, 0xbc, 0x21, 0xc3, 0x3c, 0xfc, 0x14, 0xcf, 0xa0, 0xe7, 0x5a, 0x5d, 0xa8,
	0x9a, 0xe7, 0xeb, 0xe6, 0x99, 0x6b, 0xf5, 0x55, 0x2d, 0x26, 0x30, 0xf8, 0x67, 0x15, 0xb6, 0x67,
	0x71, 0x02, 0xbd, 0xd8, 0x60, 0xda, 0x80, 0x74, 0x0a, 0x9f, 0xb3, 0xdb, 0x6c, 0x5c, 0x00, 0xa0,
	0x87, 0x4e, 0x4f, 0xa1, 0x5f, 0x61, 0xa1, 0xf4, 0x9d, 0x49, 0xb5, 0xf7, 0x2a, 0xbc, 0xd2, 0x77,
	0x46, 0x9c, 0xc3, 0xe0, 0x51, 0xcb, 0xa3, 0x8b, 0xc3, 0x18, 0x74, 0x82, 0xf9, 0x56, 0x0e, 0x1b,
	0xe8, 0x55, 0x83, 0x61, 0x51, 0x6b, 0x2e, 0x3a, 0xcb, 0x07, 0x01, 0xdc, 0x10, 0xd6, 0x61, 0x82,
	0x06, 0x1b, 0xe3, 0x7e, 0x46, 0x79, 0x18, 0x03, 0x8d, 0x88, 0x0d, 0x2f, 0x61, 0xe8, 0x91, 0x7c,
	0x55, 0x12, 0x92, 0x04, 0x96, 0x1f, 0x80, 0x38, 0x87, 0xf1, 0xf6, 0x50, 0xd8, 0x92, 0xc2, 0x1d,
	0x23, 0x36, 0x1d, 0x6f, 0xf9, 0x92, 0x71, 0x78, 0x53, 0xf8, 0xc8, 0x62, 0x83, 0x7a, 0xe5, 0xbf,
	0xcb, 0x83, 0xf8, 0xa6, 0x80, 0x16, 0x4c, 0x38, 0x8c, 0xb0, 0x22, 0xbe, 0x08, 0xd3, 0xc9, 0xc3,
	0x14, 0x06, 0xa3, 0x6f, 0xaa, 0x41, 0x21, 0x60, 0x9f, 0x63, 0x3a, 0x62, 0x85, 0x7f, 0x87, 0x32,
	0x14, 0x15, 0x64, 0xd7, 0xf2, 0x98, 0x37, 0x34, 0x53, 0x74, 0x6d, 0xd7, 0xb7, 0x3d, 0xfe, 0xfb,
	0x7f, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x37, 0xb7, 0x97, 0x47, 0x04, 0x00, 0x00,
}
