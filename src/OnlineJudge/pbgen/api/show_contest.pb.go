// Code generated by protoc-gen-go.
// source: api/show_contest.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ShowContestRequest struct {
	ContestId int64 `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
}

func (m *ShowContestRequest) Reset()                    { *m = ShowContestRequest{} }
func (m *ShowContestRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowContestRequest) ProtoMessage()               {}
func (*ShowContestRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *ShowContestRequest) GetContestId() int64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

type Contest struct {
	ContestId     int64  `protobuf:"varint,1,opt,name=contest_id,json=contestId" json:"contest_id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Description   string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	IsVirtual     bool   `protobuf:"varint,4,opt,name=is_virtual,json=isVirtual" json:"is_virtual,omitempty"`
	ContestType   string `protobuf:"bytes,5,opt,name=contest_type,json=contestType" json:"contest_type,omitempty"`
	StartTime     string `protobuf:"bytes,6,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime       string `protobuf:"bytes,7,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	LockBoardTime string `protobuf:"bytes,8,opt,name=lock_board_time,json=lockBoardTime" json:"lock_board_time,omitempty"`
	Access        string `protobuf:"bytes,9,opt,name=access" json:"access,omitempty"`
	Status        string `protobuf:"bytes,10,opt,name=status" json:"status,omitempty"`
}

func (m *Contest) Reset()                    { *m = Contest{} }
func (m *Contest) String() string            { return proto.CompactTextString(m) }
func (*Contest) ProtoMessage()               {}
func (*Contest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *Contest) GetContestId() int64 {
	if m != nil {
		return m.ContestId
	}
	return 0
}

func (m *Contest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Contest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Contest) GetIsVirtual() bool {
	if m != nil {
		return m.IsVirtual
	}
	return false
}

func (m *Contest) GetContestType() string {
	if m != nil {
		return m.ContestType
	}
	return ""
}

func (m *Contest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Contest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Contest) GetLockBoardTime() string {
	if m != nil {
		return m.LockBoardTime
	}
	return ""
}

func (m *Contest) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *Contest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ShowContestResponse struct {
	Contest *Contest `protobuf:"bytes,1,opt,name=contest" json:"contest,omitempty"`
	Error   *Error   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ShowContestResponse) Reset()                    { *m = ShowContestResponse{} }
func (m *ShowContestResponse) String() string            { return proto.CompactTextString(m) }
func (*ShowContestResponse) ProtoMessage()               {}
func (*ShowContestResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *ShowContestResponse) GetContest() *Contest {
	if m != nil {
		return m.Contest
	}
	return nil
}

func (m *ShowContestResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowContestRequest)(nil), "api.ShowContestRequest")
	proto.RegisterType((*Contest)(nil), "api.Contest")
	proto.RegisterType((*ShowContestResponse)(nil), "api.ShowContestResponse")
}

func init() { proto.RegisterFile("api/show_contest.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0xd6, 0x76, 0xbb, 0xd3, 0x8a, 0x12, 0xa5, 0x44, 0x41, 0x58, 0x7b, 0x90, 0x9e,
	0x2a, 0xb4, 0x6f, 0xa0, 0x78, 0xf0, 0xba, 0x16, 0xaf, 0x4b, 0xba, 0x3b, 0xd0, 0x60, 0x77, 0x13,
	0x33, 0x53, 0x4b, 0xdf, 0xcb, 0x07, 0x94, 0x24, 0x5b, 0xd0, 0x93, 0xc7, 0xf9, 0xfe, 0x6f, 0x66,
	0xc2, 0x04, 0xa6, 0xca, 0xea, 0x47, 0xda, 0x9a, 0x43, 0x59, 0x99, 0x96, 0x91, 0x78, 0x61, 0x9d,
	0x61, 0x23, 0xfa, 0xca, 0xea, 0xdb, 0x4b, 0x1f, 0x56, 0xa6, 0x69, 0x4c, 0x1b, 0xf1, 0x6c, 0x05,
	0xe2, 0x6d, 0x6b, 0x0e, 0xcf, 0xd1, 0x2d, 0xf0, 0x73, 0x8f, 0xc4, 0xe2, 0x0e, 0xa0, 0xeb, 0x2e,
	0x75, 0x2d, 0x93, 0x3c, 0x99, 0xf7, 0x8b, 0xac, 0x23, 0xaf, 0xf5, 0xec, 0xbb, 0x07, 0x69, 0xd7,
	0xf1, 0x8f, 0x2a, 0xae, 0x61, 0xc0, 0x9a, 0x77, 0x28, 0x7b, 0x79, 0x32, 0xcf, 0x8a, 0x58, 0x88,
	0x1c, 0xc6, 0x35, 0x52, 0xe5, 0xb4, 0x65, 0x6d, 0x5a, 0xd9, 0x0f, 0xd9, 0x6f, 0xe4, 0xc7, 0x6a,
	0x2a, 0xbf, 0xb4, 0xe3, 0xbd, 0xda, 0xc9, 0xb3, 0x3c, 0x99, 0x8f, 0x8a, 0x4c, 0xd3, 0x7b, 0x04,
	0xe2, 0x1e, 0x26, 0xa7, 0xad, 0x7c, 0xb4, 0x28, 0x07, 0x71, 0x42, 0xc7, 0xd6, 0x47, 0x8b, 0x7e,
	0x02, 0xb1, 0x72, 0x5c, 0xb2, 0x6e, 0x50, 0x0e, 0x83, 0x90, 0x05, 0xb2, 0xd6, 0x0d, 0x8a, 0x1b,
	0x18, 0x61, 0x5b, 0xc7, 0x30, 0x0d, 0x61, 0x8a, 0x6d, 0x1d, 0xa2, 0x07, 0xb8, 0xd8, 0x99, 0xea,
	0xa3, 0xdc, 0x18, 0xe5, 0x3a, 0x63, 0x14, 0x8c, 0x73, 0x8f, 0x9f, 0x3c, 0x0d, 0xde, 0x14, 0x86,
	0xaa, 0xaa, 0x90, 0x48, 0x66, 0x21, 0xee, 0x2a, 0xcf, 0x89, 0x15, 0xef, 0x49, 0x42, 0xe4, 0xb1,
	0x9a, 0x95, 0x70, 0xf5, 0xe7, 0xd6, 0x64, 0x4d, 0x4b, 0x7e, 0x5d, 0xda, 0xbd, 0x3b, 0x9c, 0x6f,
	0xbc, 0x9c, 0x2c, 0x94, 0xd5, 0x8b, 0x93, 0x76, 0x0a, 0x45, 0x0e, 0x03, 0x74, 0xce, 0xb8, 0x70,
	0xca, 0xf1, 0x12, 0x82, 0xf5, 0xe2, 0x49, 0x11, 0x83, 0xcd, 0x30, 0xfc, 0xe9, 0xea, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xeb, 0x82, 0x34, 0x5b, 0x04, 0x02, 0x00, 0x00,
}
