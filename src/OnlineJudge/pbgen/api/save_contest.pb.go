// Code generated by protoc-gen-go.
// source: api/save_contest.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SaveContestRequest struct {
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
	Problems         []*SaveContestRequest_ContestProblem `protobuf:"bytes,12,rep,name=problems" json:"problems,omitempty"`
}

func (m *SaveContestRequest) Reset()                    { *m = SaveContestRequest{} }
func (m *SaveContestRequest) String() string            { return proto.CompactTextString(m) }
func (*SaveContestRequest) ProtoMessage()               {}
func (*SaveContestRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *SaveContestRequest) GetContestId() int64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *SaveContestRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SaveContestRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SaveContestRequest) GetIsVirtual() bool {
	if m != nil {
		return m.IsVirtual
	}
	return false
}

func (m *SaveContestRequest) GetContestType() string {
	if m != nil {
		return m.ContestType
	}
	return ""
}

func (m *SaveContestRequest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *SaveContestRequest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *SaveContestRequest) GetLockBoardTime() string {
	if m != nil {
		return m.LockBoardTime
	}
	return ""
}

func (m *SaveContestRequest) GetHideOthersStatus() bool {
	if m != nil {
		return m.HideOthersStatus
	}
	return false
}

func (m *SaveContestRequest) GetIsHidden() bool {
	if m != nil {
		return m.IsHidden
	}
	return false
}

func (m *SaveContestRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SaveContestRequest) GetProblems() []*SaveContestRequest_ContestProblem {
	if m != nil {
		return m.Problems
	}
	return nil
}

type SaveContestRequest_ContestProblem struct {
	Alias      string `protobuf:"bytes,1,opt,name=alias" json:"alias,omitempty"`
	ProblemSid string `protobuf:"bytes,2,opt,name=problem_sid,json=problemSid" json:"problem_sid,omitempty"`
	Id         int32  `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	MetaPid    int64  `protobuf:"varint,4,opt,name=meta_pid,json=metaPid" json:"meta_pid,omitempty"`
}

func (m *SaveContestRequest_ContestProblem) Reset()         { *m = SaveContestRequest_ContestProblem{} }
func (m *SaveContestRequest_ContestProblem) String() string { return proto.CompactTextString(m) }
func (*SaveContestRequest_ContestProblem) ProtoMessage()    {}
func (*SaveContestRequest_ContestProblem) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{0, 0}
}

func (m *SaveContestRequest_ContestProblem) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SaveContestRequest_ContestProblem) GetProblemSid() string {
	if m != nil {
		return m.ProblemSid
	}
	return ""
}

func (m *SaveContestRequest_ContestProblem) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SaveContestRequest_ContestProblem) GetMetaPid() int64 {
	if m != nil {
		return m.MetaPid
	}
	return 0
}

type SaveContestResponse struct {
	ContestId int64  `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	Error     *Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SaveContestResponse) Reset()                    { *m = SaveContestResponse{} }
func (m *SaveContestResponse) String() string            { return proto.CompactTextString(m) }
func (*SaveContestResponse) ProtoMessage()               {}
func (*SaveContestResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *SaveContestResponse) GetContestId() int64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *SaveContestResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*SaveContestRequest)(nil), "api.SaveContestRequest")
	proto.RegisterType((*SaveContestRequest_ContestProblem)(nil), "api.SaveContestRequest.ContestProblem")
	proto.RegisterType((*SaveContestResponse)(nil), "api.SaveContestResponse")
}

func init() { proto.RegisterFile("api/save_contest.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0x13, 0x3f,
	0x10, 0xc5, 0x95, 0xa4, 0x69, 0x76, 0x67, 0xfb, 0xef, 0xbf, 0x32, 0x08, 0x99, 0x20, 0xc4, 0xd2,
	0x43, 0x95, 0x03, 0x0a, 0x52, 0xf9, 0x06, 0x45, 0x48, 0x70, 0xa2, 0xda, 0x54, 0xbd, 0x5a, 0x4e,
	0x3c, 0x52, 0x46, 0xec, 0xae, 0x8d, 0xc7, 0x09, 0xea, 0x97, 0xe2, 0x33, 0x22, 0xcf, 0x2e, 0x15,
	0x15, 0x07, 0x8e, 0xef, 0xf7, 0x9e, 0x67, 0xbc, 0xfb, 0x0c, 0x2f, 0x6c, 0xa0, 0xf7, 0x6c, 0x8f,
	0x68, 0x76, 0xbe, 0x4f, 0xc8, 0x69, 0x1d, 0xa2, 0x4f, 0x5e, 0xcd, 0x6c, 0xa0, 0xe5, 0x45, 0x36,
	0x77, 0xbe, 0xeb, 0x7c, 0x3f, 0xe0, 0xcb, 0x9f, 0x27, 0xa0, 0x36, 0xf6, 0x88, 0x1f, 0x87, 0x70,
	0x83, 0xdf, 0x0f, 0xc8, 0x49, 0xbd, 0x06, 0x18, 0x8f, 0x1b, 0x72, 0x7a, 0x52, 0x4f, 0x56, 0xb3,
	0xa6, 0x1c, 0xc9, 0x17, 0xa7, 0x9e, 0xc3, 0x3c, 0x51, 0x6a, 0x51, 0x4f, 0xeb, 0xc9, 0xaa, 0x6c,
	0x06, 0xa1, 0x6a, 0xa8, 0x1c, 0xf2, 0x2e, 0x52, 0x48, 0xe4, 0x7b, 0x3d, 0x13, 0xef, 0x4f, 0x94,
	0xc7, 0x12, 0x9b, 0x23, 0xc5, 0x74, 0xb0, 0xad, 0x3e, 0xa9, 0x27, 0xab, 0xa2, 0x29, 0x89, 0xef,
	0x07, 0xa0, 0xde, 0xc2, 0xd9, 0xef, 0xad, 0xe9, 0x21, 0xa0, 0x9e, 0x0f, 0x13, 0x46, 0x76, 0xf7,
	0x10, 0x30, 0x4f, 0xe0, 0x64, 0x63, 0x32, 0x89, 0x3a, 0xd4, 0xa7, 0x12, 0x28, 0x85, 0xdc, 0x51,
	0x87, 0xea, 0x25, 0x14, 0xd8, 0xbb, 0xc1, 0x5c, 0x88, 0xb9, 0xc0, 0xde, 0x89, 0x75, 0x05, 0xff,
	0xb7, 0x7e, 0xf7, 0xcd, 0x6c, 0xbd, 0x8d, 0x63, 0xa2, 0x90, 0xc4, 0x7f, 0x19, 0xdf, 0x64, 0x2a,
	0xb9, 0x77, 0xa0, 0xf6, 0xe4, 0xd0, 0xf8, 0xb4, 0xc7, 0xc8, 0x86, 0x93, 0x4d, 0x07, 0xd6, 0xa5,
	0xdc, 0xf5, 0x22, 0x3b, 0x5f, 0xc5, 0xd8, 0x08, 0x57, 0xaf, 0xa0, 0x24, 0x36, 0x7b, 0x72, 0x0e,
	0x7b, 0x0d, 0x12, 0x2a, 0x88, 0x3f, 0x8b, 0x56, 0x4b, 0x28, 0x82, 0x65, 0xfe, 0xe1, 0xa3, 0xd3,
	0x95, 0xec, 0x7a, 0xd4, 0xea, 0x06, 0x8a, 0x10, 0xfd, 0xb6, 0xc5, 0x8e, 0xf5, 0x59, 0x3d, 0x5b,
	0x55, 0xd7, 0x57, 0x6b, 0x1b, 0x68, 0xfd, 0x77, 0x19, 0xeb, 0x51, 0xde, 0x0e, 0xf1, 0xe6, 0xf1,
	0xdc, 0x32, 0xc2, 0xf9, 0x53, 0x2f, 0x17, 0x63, 0x5b, 0xb2, 0x2c, 0x95, 0x95, 0xcd, 0x20, 0xd4,
	0x1b, 0xa8, 0xc6, 0x33, 0x86, 0xc9, 0x8d, 0xa5, 0xc1, 0x88, 0x36, 0xe4, 0xd4, 0x39, 0x4c, 0xc9,
	0x49, 0x61, 0xf3, 0x66, 0x4a, 0x2e, 0xff, 0xc6, 0x0e, 0x93, 0x35, 0x81, 0x9c, 0xb4, 0x34, 0x6b,
	0x16, 0x59, 0xdf, 0x92, 0xbb, 0xbc, 0x87, 0x67, 0x4f, 0xae, 0xc8, 0xc1, 0xf7, 0x8c, 0xff, 0x7a,
	0x30, 0x35, 0xcc, 0x31, 0x46, 0x1f, 0x65, 0x77, 0x75, 0x0d, 0xf2, 0xa9, 0x9f, 0x32, 0x69, 0x06,
	0x63, 0x7b, 0x2a, 0xef, 0xf1, 0xc3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0x6a, 0xb9, 0xc2,
	0xc0, 0x02, 0x00, 0x00,
}
