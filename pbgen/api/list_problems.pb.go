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
	return fileDescriptor4, []int{0, 0}
}

type ListProblemsRequest_Filter_ProblemStatus int32

const (
	ListProblemsRequest_Filter_ALL       ListProblemsRequest_Filter_ProblemStatus = 0
	ListProblemsRequest_Filter_ACCEPTED  ListProblemsRequest_Filter_ProblemStatus = 1
	ListProblemsRequest_Filter_UNSOLVED  ListProblemsRequest_Filter_ProblemStatus = 2
	ListProblemsRequest_Filter_ATTEMPTED ListProblemsRequest_Filter_ProblemStatus = 3
)

var ListProblemsRequest_Filter_ProblemStatus_name = map[int32]string{
	0: "ALL",
	1: "ACCEPTED",
	2: "UNSOLVED",
	3: "ATTEMPTED",
}
var ListProblemsRequest_Filter_ProblemStatus_value = map[string]int32{
	"ALL":       0,
	"ACCEPTED":  1,
	"UNSOLVED":  2,
	"ATTEMPTED": 3,
}

func (x ListProblemsRequest_Filter_ProblemStatus) String() string {
	return proto.EnumName(ListProblemsRequest_Filter_ProblemStatus_name, int32(x))
}
func (ListProblemsRequest_Filter_ProblemStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 0, 0}
}

type ListProblemsRequest struct {
	Offset      int32                       `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	PerPage     int32                       `protobuf:"varint,2,opt,name=per_page,json=perPage" json:"per_page,omitempty"`
	CurrentPage int32                       `protobuf:"varint,3,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	OrderBy     ListProblemsRequest_Element `protobuf:"varint,4,opt,name=order_by,json=orderBy,enum=api.ListProblemsRequest_Element" json:"order_by,omitempty"`
	IsDesc      bool                        `protobuf:"varint,5,opt,name=is_desc,json=isDesc" json:"is_desc,omitempty"`
	Filter      *ListProblemsRequest_Filter `protobuf:"bytes,6,opt,name=filter" json:"filter,omitempty"`
	NeedOjsList bool                        `protobuf:"varint,7,opt,name=need_ojs_list,json=needOjsList" json:"need_ojs_list,omitempty"`
}

func (m *ListProblemsRequest) Reset()                    { *m = ListProblemsRequest{} }
func (m *ListProblemsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListProblemsRequest) ProtoMessage()               {}
func (*ListProblemsRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

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

func (m *ListProblemsRequest) GetFilter() *ListProblemsRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListProblemsRequest) GetNeedOjsList() bool {
	if m != nil {
		return m.NeedOjsList
	}
	return false
}

type ListProblemsRequest_Filter struct {
	Oj      string                                   `protobuf:"bytes,1,opt,name=oj" json:"oj,omitempty"`
	PStatus ListProblemsRequest_Filter_ProblemStatus `protobuf:"varint,2,opt,name=p_status,json=pStatus,enum=api.ListProblemsRequest_Filter_ProblemStatus" json:"p_status,omitempty"`
}

func (m *ListProblemsRequest_Filter) Reset()                    { *m = ListProblemsRequest_Filter{} }
func (m *ListProblemsRequest_Filter) String() string            { return proto.CompactTextString(m) }
func (*ListProblemsRequest_Filter) ProtoMessage()               {}
func (*ListProblemsRequest_Filter) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

func (m *ListProblemsRequest_Filter) GetOj() string {
	if m != nil {
		return m.Oj
	}
	return ""
}

func (m *ListProblemsRequest_Filter) GetPStatus() ListProblemsRequest_Filter_ProblemStatus {
	if m != nil {
		return m.PStatus
	}
	return ListProblemsRequest_Filter_ALL
}

type ListProblemsResponse struct {
	Lines       []*ListProblemsResponse_PerLine `protobuf:"bytes,1,rep,name=lines" json:"lines,omitempty"`
	TotalLines  int32                           `protobuf:"varint,2,opt,name=total_lines,json=totalLines" json:"total_lines,omitempty"`
	TotalPages  int32                           `protobuf:"varint,3,opt,name=total_pages,json=totalPages" json:"total_pages,omitempty"`
	CurrentPage int32                           `protobuf:"varint,4,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	OjsList     []string                        `protobuf:"bytes,5,rep,name=ojs_list,json=ojsList" json:"ojs_list,omitempty"`
	Error       *Error                          `protobuf:"bytes,6,opt,name=error" json:"error,omitempty"`
}

func (m *ListProblemsResponse) Reset()                    { *m = ListProblemsResponse{} }
func (m *ListProblemsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListProblemsResponse) ProtoMessage()               {}
func (*ListProblemsResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

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

func (m *ListProblemsResponse) GetOjsList() []string {
	if m != nil {
		return m.OjsList
	}
	return nil
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

func (m *ListProblemsResponse_PerLine) Reset()                    { *m = ListProblemsResponse_PerLine{} }
func (m *ListProblemsResponse_PerLine) String() string            { return proto.CompactTextString(m) }
func (*ListProblemsResponse_PerLine) ProtoMessage()               {}
func (*ListProblemsResponse_PerLine) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1, 0} }

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
	proto.RegisterType((*ListProblemsRequest_Filter)(nil), "api.ListProblemsRequest.Filter")
	proto.RegisterType((*ListProblemsResponse)(nil), "api.ListProblemsResponse")
	proto.RegisterType((*ListProblemsResponse_PerLine)(nil), "api.ListProblemsResponse.PerLine")
	proto.RegisterEnum("api.ListProblemsRequest_Element", ListProblemsRequest_Element_name, ListProblemsRequest_Element_value)
	proto.RegisterEnum("api.ListProblemsRequest_Filter_ProblemStatus", ListProblemsRequest_Filter_ProblemStatus_name, ListProblemsRequest_Filter_ProblemStatus_value)
}

func init() { proto.RegisterFile("api/list_problems.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0x9a, 0xa5, 0x69, 0x5e, 0xd7, 0x11, 0x99, 0x89, 0x65, 0xbb, 0x2c, 0x2b, 0x97, 0x80,
	0x44, 0x90, 0xca, 0x61, 0x07, 0x4e, 0xa5, 0x0d, 0x62, 0x52, 0x60, 0x95, 0x17, 0xb8, 0x46, 0x69,
	0xea, 0x4d, 0xae, 0xd2, 0xd8, 0xd8, 0xee, 0x61, 0xbf, 0x88, 0x23, 0x3f, 0x80, 0xbf, 0xc5, 0x0f,
	0x40, 0xb6, 0xc3, 0xa0, 0xda, 0x10, 0x37, 0xbf, 0xef, 0xfb, 0x9e, 0xf3, 0xe5, 0xbd, 0xcf, 0x70,
	0x5c, 0x71, 0xfa, 0xba, 0xa1, 0x52, 0x95, 0x5c, 0xb0, 0x65, 0x43, 0x36, 0x32, 0xe5, 0x82, 0x29,
	0x86, 0xdc, 0x8a, 0xd3, 0xd3, 0x50, 0xb3, 0x35, 0xdb, 0x6c, 0x58, 0x6b, 0xe1, 0xf1, 0x4f, 0x17,
	0x9e, 0xe6, 0x54, 0xaa, 0x45, 0xa7, 0xc6, 0xe4, 0xeb, 0x96, 0x48, 0x85, 0x9e, 0x41, 0x9f, 0xdd,
	0xdc, 0x48, 0xa2, 0x22, 0x27, 0x76, 0x12, 0x0f, 0x77, 0x15, 0x3a, 0x81, 0x01, 0x27, 0xa2, 0xe4,
	0xd5, 0x2d, 0x89, 0x7a, 0x86, 0xf1, 0x39, 0x11, 0x8b, 0xea, 0x96, 0xa0, 0x73, 0x38, 0xa8, 0xb7,
	0x42, 0x90, 0x56, 0x59, 0xda, 0x35, 0xf4, 0xb0, 0xc3, 0x8c, 0xe4, 0x2d, 0x0c, 0x98, 0x58, 0x11,
	0x51, 0x2e, 0xef, 0xa2, 0xfd, 0xd8, 0x49, 0x0e, 0x27, 0x71, 0x5a, 0x71, 0x9a, 0x3e, 0xe2, 0x20,
	0xcd, 0x1a, 0xb2, 0x21, 0xad, 0xc2, 0xbe, 0xe9, 0x78, 0x77, 0x87, 0x8e, 0xc1, 0xa7, 0xb2, 0x5c,
	0x11, 0x59, 0x47, 0x5e, 0xec, 0x24, 0x03, 0xdc, 0xa7, 0x72, 0x4e, 0x64, 0x8d, 0x2e, 0xa0, 0x7f,
	0x43, 0x1b, 0x45, 0x44, 0xd4, 0x8f, 0x9d, 0x64, 0x38, 0x39, 0xfb, 0xe7, 0x9d, 0xef, 0x8d, 0x0c,
	0x77, 0x72, 0x34, 0x86, 0x51, 0x4b, 0xc8, 0xaa, 0x64, 0x6b, 0x59, 0xea, 0x99, 0x45, 0xbe, 0xb9,
	0x77, 0xa8, 0xc1, 0xab, 0xb5, 0xd4, 0x37, 0x9c, 0x7e, 0x73, 0xa0, 0x6f, 0xdb, 0xd0, 0x21, 0xf4,
	0xd8, 0xda, 0xcc, 0x23, 0xc0, 0x3d, 0xb6, 0x46, 0x1f, 0x60, 0xc0, 0x4b, 0xa9, 0x2a, 0xb5, 0x95,
	0x66, 0x16, 0x87, 0x93, 0x57, 0xff, 0xf9, 0x72, 0xda, 0xc1, 0xd7, 0xa6, 0x09, 0xfb, 0xdc, 0x1e,
	0xc6, 0x33, 0x18, 0xed, 0x30, 0xc8, 0x07, 0x77, 0x9a, 0xe7, 0xe1, 0x1e, 0x3a, 0x80, 0xc1, 0x74,
	0x36, 0xcb, 0x16, 0x45, 0x36, 0x0f, 0x1d, 0x5d, 0x7d, 0xfe, 0x74, 0x7d, 0x95, 0x7f, 0xc9, 0xe6,
	0x61, 0x0f, 0x8d, 0x20, 0x98, 0x16, 0x45, 0xf6, 0xd1, 0x90, 0xee, 0xf8, 0x25, 0xf8, 0xdd, 0xcc,
	0x74, 0xfb, 0xe2, 0x72, 0x1e, 0xee, 0xa1, 0x00, 0xbc, 0xe2, 0xb2, 0xc8, 0xb3, 0xd0, 0x41, 0x43,
	0xf0, 0xa7, 0xb3, 0x12, 0x4f, 0x8b, 0x2c, 0xec, 0x8d, 0xbf, 0xbb, 0x70, 0xb4, 0x6b, 0x53, 0x72,
	0xd6, 0x4a, 0x82, 0x2e, 0xc0, 0x6b, 0x68, 0x4b, 0x64, 0xe4, 0xc4, 0x6e, 0x32, 0x9c, 0x9c, 0x3f,
	0xf2, 0x43, 0x56, 0x99, 0x2e, 0x88, 0xc8, 0x69, 0x4b, 0xb0, 0xd5, 0xa3, 0x33, 0x18, 0x2a, 0xa6,
	0xaa, 0xa6, 0xb4, 0xed, 0x36, 0x1b, 0x60, 0xa0, 0x7c, 0x57, 0xa0, 0xc3, 0x21, 0xbb, 0x74, 0x58,
	0x81, 0xce, 0x86, 0x7c, 0x90, 0x9f, 0xfd, 0x87, 0xf9, 0x39, 0x81, 0xc1, 0xfd, 0xae, 0xbc, 0xd8,
	0x4d, 0x02, 0xec, 0x33, 0xbb, 0x27, 0x14, 0x83, 0x47, 0x84, 0x60, 0xbf, 0x33, 0x00, 0xc6, 0x78,
	0xa6, 0x11, 0x6c, 0x89, 0xd3, 0x1f, 0x0e, 0xf8, 0x9d, 0x69, 0x14, 0x82, 0x2b, 0xe9, 0xaa, 0xdb,
	0xa5, 0x3e, 0x76, 0xcb, 0xed, 0xdd, 0x2f, 0x37, 0x04, 0x97, 0xd3, 0x95, 0xb1, 0x19, 0x60, 0x7d,
	0x44, 0x47, 0xe0, 0x29, 0xaa, 0x1a, 0x6b, 0x2c, 0xc0, 0xb6, 0xd0, 0x0f, 0x45, 0xb2, 0xad, 0xa8,
	0x89, 0x09, 0x65, 0x80, 0xbb, 0x0a, 0x3d, 0x87, 0x51, 0x55, 0x97, 0x72, 0xbb, 0xdc, 0x50, 0x29,
	0x29, 0x6b, 0x8d, 0x2f, 0x0f, 0x1f, 0x54, 0xf5, 0xf5, 0x3d, 0x86, 0x5e, 0x40, 0x68, 0x67, 0xf2,
	0x97, 0xce, 0x37, 0xba, 0x27, 0x06, 0xff, 0x23, 0x5d, 0xf6, 0xcd, 0x7b, 0x7d, 0xf3, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0xb4, 0xeb, 0x0a, 0xda, 0xe1, 0x03, 0x00, 0x00,
}
