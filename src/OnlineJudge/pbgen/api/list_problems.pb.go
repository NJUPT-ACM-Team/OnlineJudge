// Code generated by protoc-gen-go.
// source: api/list_problems.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListProblemsRequest_Element int32

const (
	ListProblemsRequest_PID     ListProblemsRequest_Element = 0
	ListProblemsRequest_TITLE   ListProblemsRequest_Element = 1
	ListProblemsRequest_AC_RATE ListProblemsRequest_Element = 2
)

var ListProblemsRequest_Element_name = map[int32]string{
	0: "PID",
	1: "TITLE",
	2: "AC_RATE",
}
var ListProblemsRequest_Element_value = map[string]int32{
	"PID":     0,
	"TITLE":   1,
	"AC_RATE": 2,
}

func (x ListProblemsRequest_Element) String() string {
	return proto.EnumName(ListProblemsRequest_Element_name, int32(x))
}
func (ListProblemsRequest_Element) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor11, []int{0, 0}
}

type ListProblemsRequest_ProblemStatus int32

const (
	ListProblemsRequest_ALL       ListProblemsRequest_ProblemStatus = 0
	ListProblemsRequest_ACCEPTED  ListProblemsRequest_ProblemStatus = 1
	ListProblemsRequest_UNSOLVED  ListProblemsRequest_ProblemStatus = 2
	ListProblemsRequest_ATTEMPTED ListProblemsRequest_ProblemStatus = 3
)

var ListProblemsRequest_ProblemStatus_name = map[int32]string{
	0: "ALL",
	1: "ACCEPTED",
	2: "UNSOLVED",
	3: "ATTEMPTED",
}
var ListProblemsRequest_ProblemStatus_value = map[string]int32{
	"ALL":       0,
	"ACCEPTED":  1,
	"UNSOLVED":  2,
	"ATTEMPTED": 3,
}

func (x ListProblemsRequest_ProblemStatus) String() string {
	return proto.EnumName(ListProblemsRequest_ProblemStatus_name, int32(x))
}
func (ListProblemsRequest_ProblemStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor11, []int{0, 1}
}

type ListProblemsRequest struct {
	Offset        int32                             `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	PerPage       int32                             `protobuf:"varint,2,opt,name=per_page,json=perPage" json:"per_page,omitempty"`
	CurrentPage   int32                             `protobuf:"varint,3,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	OrderBy       ListProblemsRequest_Element       `protobuf:"varint,4,opt,name=order_by,json=orderBy,enum=api.ListProblemsRequest_Element" json:"order_by,omitempty"`
	IsDesc        bool                              `protobuf:"varint,5,opt,name=is_desc,json=isDesc" json:"is_desc,omitempty"`
	FilterOj      string                            `protobuf:"bytes,6,opt,name=filter_oj,json=filterOj" json:"filter_oj,omitempty"`
	FilterPStatus ListProblemsRequest_ProblemStatus `protobuf:"varint,7,opt,name=filter_p_status,json=filterPStatus,enum=api.ListProblemsRequest_ProblemStatus" json:"filter_p_status,omitempty"`
}

func (m *ListProblemsRequest) Reset()                    { *m = ListProblemsRequest{} }
func (m *ListProblemsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListProblemsRequest) ProtoMessage()               {}
func (*ListProblemsRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *ListProblemsRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListProblemsRequest) GetPerPage() int32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *ListProblemsRequest) GetCurrentPage() int32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *ListProblemsRequest) GetOrderBy() ListProblemsRequest_Element {
	if m != nil {
		return m.OrderBy
	}
	return ListProblemsRequest_PID
}

func (m *ListProblemsRequest) GetIsDesc() bool {
	if m != nil {
		return m.IsDesc
	}
	return false
}

func (m *ListProblemsRequest) GetFilterOj() string {
	if m != nil {
		return m.FilterOj
	}
	return ""
}

func (m *ListProblemsRequest) GetFilterPStatus() ListProblemsRequest_ProblemStatus {
	if m != nil {
		return m.FilterPStatus
	}
	return ListProblemsRequest_ALL
}

type ListProblemsResponse struct {
	Lines       []*ListProblemsResponse_PerLine `protobuf:"bytes,1,rep,name=lines" json:"lines,omitempty"`
	TotalLines  int32                           `protobuf:"varint,2,opt,name=total_lines,json=totalLines" json:"total_lines,omitempty"`
	TotalPages  int32                           `protobuf:"varint,3,opt,name=total_pages,json=totalPages" json:"total_pages,omitempty"`
	CurrentPage int32                           `protobuf:"varint,4,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	Error       *Error                          `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
}

func (m *ListProblemsResponse) Reset()                    { *m = ListProblemsResponse{} }
func (m *ListProblemsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListProblemsResponse) ProtoMessage()               {}
func (*ListProblemsResponse) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *ListProblemsResponse) GetLines() []*ListProblemsResponse_PerLine {
	if m != nil {
		return m.Lines
	}
	return nil
}

func (m *ListProblemsResponse) GetTotalLines() int32 {
	if m != nil {
		return m.TotalLines
	}
	return 0
}

func (m *ListProblemsResponse) GetTotalPages() int32 {
	if m != nil {
		return m.TotalPages
	}
	return 0
}

func (m *ListProblemsResponse) GetCurrentPage() int32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *ListProblemsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ListProblemsResponse_PerLine struct {
	Sid             string `protobuf:"bytes,1,opt,name=sid" json:"sid,omitempty"`
	Oj              string `protobuf:"bytes,2,opt,name=oj" json:"oj,omitempty"`
	Pid             string `protobuf:"bytes,3,opt,name=pid" json:"pid,omitempty"`
	Title           string `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Source          string `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	AcSubmission    int32  `protobuf:"varint,6,opt,name=ac_submission,json=acSubmission" json:"ac_submission,omitempty"`
	TotalSubmission int32  `protobuf:"varint,7,opt,name=total_submission,json=totalSubmission" json:"total_submission,omitempty"`
}

func (m *ListProblemsResponse_PerLine) Reset()         { *m = ListProblemsResponse_PerLine{} }
func (m *ListProblemsResponse_PerLine) String() string { return proto.CompactTextString(m) }
func (*ListProblemsResponse_PerLine) ProtoMessage()    {}
func (*ListProblemsResponse_PerLine) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{1, 0}
}

func (m *ListProblemsResponse_PerLine) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *ListProblemsResponse_PerLine) GetOj() string {
	if m != nil {
		return m.Oj
	}
	return ""
}

func (m *ListProblemsResponse_PerLine) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *ListProblemsResponse_PerLine) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ListProblemsResponse_PerLine) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ListProblemsResponse_PerLine) GetAcSubmission() int32 {
	if m != nil {
		return m.AcSubmission
	}
	return 0
}

func (m *ListProblemsResponse_PerLine) GetTotalSubmission() int32 {
	if m != nil {
		return m.TotalSubmission
	}
	return 0
}

func init() {
	proto.RegisterType((*ListProblemsRequest)(nil), "api.ListProblemsRequest")
	proto.RegisterType((*ListProblemsResponse)(nil), "api.ListProblemsResponse")
	proto.RegisterType((*ListProblemsResponse_PerLine)(nil), "api.ListProblemsResponse.PerLine")
	proto.RegisterEnum("api.ListProblemsRequest_Element", ListProblemsRequest_Element_name, ListProblemsRequest_Element_value)
	proto.RegisterEnum("api.ListProblemsRequest_ProblemStatus", ListProblemsRequest_ProblemStatus_name, ListProblemsRequest_ProblemStatus_value)
}

func init() { proto.RegisterFile("api/list_problems.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x5f, 0x9a, 0xa5, 0x69, 0x5e, 0xd7, 0x2d, 0x32, 0x13, 0x0b, 0xe3, 0x40, 0x56, 0x24, 0x54,
	0x38, 0x14, 0x69, 0x1c, 0x38, 0x70, 0x2a, 0x6d, 0x0e, 0x93, 0xc2, 0x16, 0xb9, 0x81, 0x6b, 0x94,
	0xa6, 0xee, 0xe4, 0x2a, 0x8d, 0x8d, 0x9f, 0x7b, 0xd8, 0xf7, 0xe0, 0xce, 0x07, 0xe1, 0xcb, 0xa1,
	0xd8, 0x11, 0x50, 0x6d, 0xdc, 0xfc, 0xfb, 0xf3, 0xf2, 0x7e, 0xd6, 0xcf, 0x81, 0x8b, 0x52, 0xf2,
	0xf7, 0x35, 0x47, 0x5d, 0x48, 0x25, 0x56, 0x35, 0xdb, 0xe1, 0x54, 0x2a, 0xa1, 0x05, 0x71, 0x4b,
	0xc9, 0x2f, 0xc3, 0x56, 0xad, 0xc4, 0x6e, 0x27, 0x1a, 0x4b, 0x8f, 0x7f, 0xba, 0xf0, 0x2c, 0xe5,
	0xa8, 0xb3, 0xce, 0x4d, 0xd9, 0xf7, 0x3d, 0x43, 0x4d, 0x9e, 0x43, 0x5f, 0x6c, 0x36, 0xc8, 0x74,
	0xe4, 0xc4, 0xce, 0xc4, 0xa3, 0x1d, 0x22, 0x2f, 0x60, 0x20, 0x99, 0x2a, 0x64, 0x79, 0xcf, 0xa2,
	0x9e, 0x51, 0x7c, 0xc9, 0x54, 0x56, 0xde, 0x33, 0x72, 0x05, 0x27, 0xd5, 0x5e, 0x29, 0xd6, 0x68,
	0x2b, 0xbb, 0x46, 0x1e, 0x76, 0x9c, 0xb1, 0x7c, 0x82, 0x81, 0x50, 0x6b, 0xa6, 0x8a, 0xd5, 0x43,
	0x74, 0x1c, 0x3b, 0x93, 0xd3, 0xeb, 0x78, 0x5a, 0x4a, 0x3e, 0x7d, 0x22, 0xc1, 0x34, 0xa9, 0xd9,
	0x8e, 0x35, 0x9a, 0xfa, 0x66, 0xe2, 0xf3, 0x03, 0xb9, 0x00, 0x9f, 0x63, 0xb1, 0x66, 0x58, 0x45,
	0x5e, 0xec, 0x4c, 0x06, 0xb4, 0xcf, 0x71, 0xc1, 0xb0, 0x22, 0x2f, 0x21, 0xd8, 0xf0, 0x5a, 0x33,
	0x55, 0x88, 0x6d, 0xd4, 0x8f, 0x9d, 0x49, 0x40, 0x07, 0x96, 0xb8, 0xdb, 0x92, 0x5b, 0x38, 0xeb,
	0x44, 0x59, 0xa0, 0x2e, 0xf5, 0x1e, 0x23, 0xdf, 0x6c, 0x7e, 0xf3, 0xdf, 0xcd, 0x1d, 0x5e, 0x1a,
	0x37, 0x1d, 0xd9, 0xf1, 0xcc, 0xc2, 0xf1, 0x3b, 0xf0, 0xbb, 0x64, 0xc4, 0x07, 0x37, 0xbb, 0x59,
	0x84, 0x47, 0x24, 0x00, 0x2f, 0xbf, 0xc9, 0xd3, 0x24, 0x74, 0xc8, 0x10, 0xfc, 0xd9, 0xbc, 0xa0,
	0xb3, 0x3c, 0x09, 0x7b, 0xe3, 0x39, 0x8c, 0x0e, 0xbe, 0xd5, 0x4e, 0xcc, 0xd2, 0x34, 0x3c, 0x22,
	0x27, 0x30, 0x98, 0xcd, 0xe7, 0x49, 0x96, 0x27, 0x8b, 0xd0, 0x69, 0xd1, 0xd7, 0xdb, 0xe5, 0x5d,
	0xfa, 0x2d, 0x59, 0x84, 0x3d, 0x32, 0x82, 0x60, 0x96, 0xe7, 0xc9, 0x17, 0x23, 0xba, 0xe3, 0x1f,
	0x2e, 0x9c, 0x1f, 0xa6, 0x44, 0x29, 0x1a, 0x64, 0xe4, 0x23, 0x78, 0x35, 0x6f, 0x18, 0x46, 0x4e,
	0xec, 0x4e, 0x86, 0xd7, 0x57, 0x4f, 0xdc, 0xc7, 0x3a, 0xa7, 0x19, 0x53, 0x29, 0x6f, 0x18, 0xb5,
	0x7e, 0xf2, 0x0a, 0x86, 0x5a, 0xe8, 0xb2, 0x2e, 0xec, 0xb8, 0xad, 0x11, 0x0c, 0x95, 0x1e, 0x1a,
	0xda, 0x1e, 0xb1, 0x2b, 0xd2, 0x1a, 0xda, 0x1a, 0xf1, 0x51, 0xd5, 0xc7, 0x8f, 0xab, 0x8e, 0xc1,
	0x63, 0x4a, 0x09, 0x65, 0xba, 0x1a, 0x5e, 0x83, 0x49, 0x97, 0xb4, 0x0c, 0xb5, 0xc2, 0xe5, 0x2f,
	0x07, 0xfc, 0x2e, 0x19, 0x09, 0xc1, 0x45, 0xbe, 0x36, 0x6f, 0x2d, 0xa0, 0xed, 0x91, 0x9c, 0x42,
	0x4f, 0x6c, 0x4d, 0xb6, 0x80, 0xf6, 0xc4, 0xb6, 0x75, 0x48, 0xbe, 0x36, 0x59, 0x02, 0xda, 0x1e,
	0xc9, 0x39, 0x78, 0x9a, 0xeb, 0xda, 0x6e, 0x0f, 0xa8, 0x05, 0xed, 0xc3, 0x45, 0xb1, 0x57, 0x15,
	0x33, 0x8b, 0x03, 0xda, 0x21, 0xf2, 0x1a, 0x46, 0x65, 0x55, 0xe0, 0x7e, 0xb5, 0xe3, 0x88, 0x5c,
	0x34, 0xe6, 0xa1, 0x78, 0xf4, 0xa4, 0xac, 0x96, 0x7f, 0x38, 0xf2, 0x16, 0x42, 0x7b, 0xf1, 0x7f,
	0x7c, 0xbe, 0xf1, 0x9d, 0x19, 0xfe, 0xaf, 0x75, 0xd5, 0x37, 0xff, 0xcf, 0x87, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x7e, 0xcd, 0xc4, 0xe4, 0x71, 0x03, 0x00, 0x00,
}
