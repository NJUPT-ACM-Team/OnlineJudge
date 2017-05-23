// Code generated by protoc-gen-go.
// source: api/ojinfo.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OJInfo struct {
	OjId       int64  `protobuf:"varint,1,opt,name=oj_id,json=ojId" json:"oj_id,omitempty"`
	OjName     string `protobuf:"bytes,2,opt,name=oj_name,json=ojName" json:"oj_name,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Int64Io    string `protobuf:"bytes,4,opt,name=int64io" json:"int64io,omitempty"`
	Javaclass  string `protobuf:"bytes,5,opt,name=javaclass" json:"javaclass,omitempty"`
	Status     string `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	StatusInfo string `protobuf:"bytes,7,opt,name=status_info,json=statusInfo" json:"status_info,omitempty"`
	Lastcheck  string `protobuf:"bytes,8,opt,name=lastcheck" json:"lastcheck,omitempty"`
	ProblemNum int32  `protobuf:"varint,9,opt,name=problem_num,json=problemNum" json:"problem_num,omitempty"`
}

func (m *OJInfo) Reset()                    { *m = OJInfo{} }
func (m *OJInfo) String() string            { return proto.CompactTextString(m) }
func (*OJInfo) ProtoMessage()               {}
func (*OJInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *OJInfo) GetOjId() int64 {
	if m != nil {
		return m.OjId
	}
	return 0
}

func (m *OJInfo) GetOjName() string {
	if m != nil {
		return m.OjName
	}
	return ""
}

func (m *OJInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *OJInfo) GetInt64Io() string {
	if m != nil {
		return m.Int64Io
	}
	return ""
}

func (m *OJInfo) GetJavaclass() string {
	if m != nil {
		return m.Javaclass
	}
	return ""
}

func (m *OJInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OJInfo) GetStatusInfo() string {
	if m != nil {
		return m.StatusInfo
	}
	return ""
}

func (m *OJInfo) GetLastcheck() string {
	if m != nil {
		return m.Lastcheck
	}
	return ""
}

func (m *OJInfo) GetProblemNum() int32 {
	if m != nil {
		return m.ProblemNum
	}
	return 0
}

func init() {
	proto.RegisterType((*OJInfo)(nil), "api.OJInfo")
}

func init() { proto.RegisterFile("api/ojinfo.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0xa6, 0x71, 0xc8, 0xb1, 0x20, 0x23, 0xc1, 0x0d, 0x95, 0x88, 0x98, 0x32, 0xc1,
	0x00, 0xe2, 0x1d, 0xca, 0x50, 0xa4, 0xbc, 0x40, 0x74, 0x4d, 0x5d, 0x61, 0x13, 0xfb, 0xa2, 0xd8,
	0xe9, 0xcc, 0xa3, 0xa3, 0xd8, 0x01, 0x36, 0xff, 0xdf, 0x27, 0xdf, 0xfd, 0x3a, 0xb8, 0xa1, 0x51,
	0x3f, 0xb3, 0xd1, 0xee, 0xcc, 0x4f, 0xe3, 0xc4, 0x81, 0x65, 0x4e, 0xa3, 0x7e, 0xfc, 0xde, 0x80,
	0xf8, 0x78, 0xdf, 0xbb, 0x33, 0xcb, 0x5b, 0x28, 0xd8, 0x74, 0xfa, 0x84, 0x59, 0x9d, 0x35, 0x79,
	0xbb, 0x65, 0xb3, 0x3f, 0xc9, 0x7b, 0x28, 0xd9, 0x74, 0x8e, 0xac, 0xc2, 0x4d, 0x9d, 0x35, 0x55,
	0x2b, 0xd8, 0x1c, 0xc8, 0x2a, 0x89, 0x50, 0x5e, 0xd4, 0xe4, 0x35, 0x3b, 0xcc, 0xa3, 0xf8, 0x8d,
	0x8b, 0xd1, 0x2e, 0xbc, 0xbd, 0x6a, 0xc6, 0x6d, 0x32, 0x6b, 0x94, 0x3b, 0xa8, 0x0c, 0x5d, 0xa8,
	0x1f, 0xc8, 0x7b, 0x2c, 0xa2, 0xfb, 0x07, 0xf2, 0x0e, 0x84, 0x0f, 0x14, 0x66, 0x8f, 0x22, 0x6d,
	0x4a, 0x49, 0x3e, 0xc0, 0x75, 0x7a, 0x75, 0x4b, 0x79, 0x2c, 0xa3, 0x84, 0x84, 0x62, 0xf1, 0x1d,
	0x54, 0x03, 0xf9, 0xd0, 0x7f, 0xaa, 0xfe, 0x0b, 0xaf, 0xd2, 0xd8, 0x3f, 0xb0, 0x7c, 0x1f, 0x27,
	0x3e, 0x0e, 0xca, 0x76, 0x6e, 0xb6, 0x58, 0xd5, 0x59, 0x53, 0xb4, 0xb0, 0xa2, 0xc3, 0x6c, 0x8f,
	0x22, 0x9e, 0xe3, 0xe5, 0x27, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x74, 0x42, 0x93, 0x22, 0x01, 0x00,
	0x00,
}
