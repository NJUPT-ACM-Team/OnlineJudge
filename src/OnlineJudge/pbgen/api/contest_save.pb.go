// Code generated by protoc-gen-go.
// source: api/contest_save.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContestSaveRequest struct {
	ContestId        int64                                `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	Title            string                               `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Description      string                               `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	IsVirtual        bool                                 `protobuf:"varint,4,opt,name=is_virtual,json=isVirtual" json:"is_virtual,omitempty"`
	ContestType      string                               `protobuf:"bytes,5,opt,name=contest_type,json=contestType" json:"contest_type,omitempty"`
	StartTime        string                               `protobuf:"bytes,6,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime          string                               `protobuf:"bytes,7,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	LockBoardTime    string                               `protobuf:"bytes,8,opt,name=lock_board_time,json=lockBoardTime" json:"lock_board_time,omitempty"`
	HideOthersStatus bool                                 `protobuf:"varint,9,opt,name=hide_others_status,json=hideOthersStatus" json:"hide_others_status,omitempty"`
	IsHidden         bool                                 `protobuf:"varint,10,opt,name=is_hidden,json=isHidden" json:"is_hidden,omitempty"`
	Password         string                               `protobuf:"bytes,11,opt,name=password" json:"password,omitempty"`
	Problems         []*ContestSaveRequest_ContestProblem `protobuf:"bytes,12,rep,name=problems" json:"problems,omitempty"`
}

func (m *ContestSaveRequest) Reset()                    { *m = ContestSaveRequest{} }
func (m *ContestSaveRequest) String() string            { return proto.CompactTextString(m) }
func (*ContestSaveRequest) ProtoMessage()               {}
func (*ContestSaveRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *ContestSaveRequest) GetContestId() int64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *ContestSaveRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ContestSaveRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContestSaveRequest) GetIsVirtual() bool {
	if m != nil {
		return m.IsVirtual
	}
	return false
}

func (m *ContestSaveRequest) GetContestType() string {
	if m != nil {
		return m.ContestType
	}
	return ""
}

func (m *ContestSaveRequest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *ContestSaveRequest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *ContestSaveRequest) GetLockBoardTime() string {
	if m != nil {
		return m.LockBoardTime
	}
	return ""
}

func (m *ContestSaveRequest) GetHideOthersStatus() bool {
	if m != nil {
		return m.HideOthersStatus
	}
	return false
}

func (m *ContestSaveRequest) GetIsHidden() bool {
	if m != nil {
		return m.IsHidden
	}
	return false
}

func (m *ContestSaveRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ContestSaveRequest) GetProblems() []*ContestSaveRequest_ContestProblem {
	if m != nil {
		return m.Problems
	}
	return nil
}

type ContestSaveRequest_ContestProblem struct {
	Alias      string `protobuf:"bytes,1,opt,name=alias" json:"alias,omitempty"`
	ProblemSid string `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
	Id         int32  `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	MetaPid    int64  `protobuf:"varint,4,opt,name=meta_pid,json=metaPid" json:"meta_pid,omitempty"`
}

func (m *ContestSaveRequest_ContestProblem) Reset()         { *m = ContestSaveRequest_ContestProblem{} }
func (m *ContestSaveRequest_ContestProblem) String() string { return proto.CompactTextString(m) }
func (*ContestSaveRequest_ContestProblem) ProtoMessage()    {}
func (*ContestSaveRequest_ContestProblem) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 0}
}

func (m *ContestSaveRequest_ContestProblem) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *ContestSaveRequest_ContestProblem) GetProblemSid() string {
	if m != nil {
		return m.ProblemSid
	}
	return ""
}

func (m *ContestSaveRequest_ContestProblem) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContestSaveRequest_ContestProblem) GetMetaPid() int64 {
	if m != nil {
		return m.MetaPid
	}
	return 0
}

type ContestSaveResponse struct {
	ContestId int64  `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	Error     *Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ContestSaveResponse) Reset()                    { *m = ContestSaveResponse{} }
func (m *ContestSaveResponse) String() string            { return proto.CompactTextString(m) }
func (*ContestSaveResponse) ProtoMessage()               {}
func (*ContestSaveResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ContestSaveResponse) GetContestId() int64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *ContestSaveResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ContestSaveRequest)(nil), "api.ContestSaveRequest")
	proto.RegisterType((*ContestSaveRequest_ContestProblem)(nil), "api.ContestSaveRequest.ContestProblem")
	proto.RegisterType((*ContestSaveResponse)(nil), "api.ContestSaveResponse")
}

func init() { proto.RegisterFile("api/contest_save.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0x95, 0xb8, 0x69, 0xec, 0x71, 0x29, 0xd5, 0x82, 0xd0, 0x12, 0x84, 0x30, 0x3d, 0x54,
	0x39, 0xa0, 0x20, 0x95, 0x7f, 0x50, 0x84, 0x04, 0x27, 0x2a, 0xa7, 0xea, 0xd5, 0xda, 0x64, 0x47,
	0xca, 0x08, 0xdb, 0xbb, 0xec, 0x6c, 0x82, 0xfa, 0xa7, 0xf8, 0x8d, 0xc8, 0xb3, 0x4e, 0x45, 0xc5,
	0x81, 0xe3, 0xfb, 0xde, 0xdb, 0x99, 0xb5, 0xdf, 0xc2, 0x2b, 0xe3, 0xe9, 0xe3, 0xd6, 0xf5, 0x11,
	0x39, 0x36, 0x6c, 0x0e, 0xb8, 0xf2, 0xc1, 0x45, 0xa7, 0x32, 0xe3, 0x69, 0x71, 0x91, 0xcc, 0xae,
	0x73, 0x7d, 0xc2, 0x97, 0xbf, 0x4f, 0x40, 0x7d, 0x4e, 0xe9, 0xb5, 0x39, 0x60, 0x8d, 0x3f, 0xf7,
	0xc8, 0x51, 0xbd, 0x05, 0x38, 0xce, 0x20, 0xab, 0x27, 0xd5, 0x64, 0x99, 0xd5, 0xc5, 0x48, 0xbe,
	0x59, 0xf5, 0x12, 0x66, 0x91, 0x62, 0x8b, 0x7a, 0x5a, 0x4d, 0x96, 0x45, 0x9d, 0x84, 0xaa, 0xa0,
	0xb4, 0xc8, 0xdb, 0x40, 0x3e, 0x92, 0xeb, 0x75, 0x26, 0xde, 0xdf, 0x68, 0x18, 0x4b, 0xdc, 0x1c,
	0x28, 0xc4, 0xbd, 0x69, 0xf5, 0x49, 0x35, 0x59, 0xe6, 0x75, 0x41, 0x7c, 0x9f, 0x80, 0x7a, 0x0f,
	0x67, 0xc7, 0xad, 0xf1, 0xc1, 0xa3, 0x9e, 0xa5, 0x09, 0x23, 0xbb, 0x7b, 0xf0, 0x38, 0x4c, 0xe0,
	0x68, 0x42, 0x6c, 0x22, 0x75, 0xa8, 0x4f, 0x25, 0x50, 0x08, 0xb9, 0xa3, 0x0e, 0xd5, 0x6b, 0xc8,
	0xb1, 0xb7, 0xc9, 0x9c, 0x8b, 0x39, 0xc7, 0xde, 0x8a, 0x75, 0x05, 0xcf, 0x5b, 0xb7, 0xfd, 0xd1,
	0x6c, 0x9c, 0x09, 0x63, 0x22, 0x97, 0xc4, 0xb3, 0x01, 0xdf, 0x0c, 0x54, 0x72, 0x1f, 0x40, 0xed,
	0xc8, 0x62, 0xe3, 0xe2, 0x0e, 0x03, 0x37, 0x1c, 0x4d, 0xdc, 0xb3, 0x2e, 0xe4, 0xae, 0x17, 0x83,
	0xf3, 0x5d, 0x8c, 0xb5, 0x70, 0xf5, 0x06, 0x0a, 0xe2, 0x66, 0x47, 0xd6, 0x62, 0xaf, 0x41, 0x42,
	0x39, 0xf1, 0x57, 0xd1, 0x6a, 0x01, 0xb9, 0x37, 0xcc, 0xbf, 0x5c, 0xb0, 0xba, 0x94, 0x5d, 0x8f,
	0x5a, 0xdd, 0x40, 0xee, 0x83, 0xdb, 0xb4, 0xd8, 0xb1, 0x3e, 0xab, 0xb2, 0x65, 0x79, 0x7d, 0xb5,
	0x32, 0x9e, 0x56, 0xff, 0x96, 0x71, 0x44, 0xb7, 0x29, 0x5e, 0x3f, 0x9e, 0x5b, 0x04, 0x38, 0x7f,
	0xea, 0x0d, 0xc5, 0x98, 0x96, 0x0c, 0x4b, 0x65, 0x45, 0x9d, 0x84, 0x7a, 0x07, 0xe5, 0x78, 0xa6,
	0x61, 0xb2, 0x63, 0x69, 0x30, 0xa2, 0x35, 0x59, 0x75, 0x0e, 0x53, 0xb2, 0x52, 0xd8, 0xac, 0x9e,
	0x92, 0x1d, 0x7e, 0x63, 0x87, 0xd1, 0x34, 0x9e, 0xac, 0xb4, 0x94, 0xd5, 0xf3, 0x41, 0xdf, 0x92,
	0xbd, 0xbc, 0x87, 0x17, 0x4f, 0xae, 0xc8, 0xde, 0xf5, 0x8c, 0xff, 0x7b, 0x30, 0x15, 0xcc, 0x30,
	0x04, 0x17, 0x64, 0x77, 0x79, 0x0d, 0xf2, 0xa9, 0x5f, 0x06, 0x52, 0x27, 0x63, 0x73, 0x2a, 0xef,
	0xf1, 0xd3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x3f, 0xf8, 0x14, 0xc0, 0x02, 0x00, 0x00,
}