// Code generated by protoc-gen-go.
// source: messages/mq.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	messages/mq.proto

It has these top-level messages:
	SubmitLanguage
	TestCase
	SpjCode
	SubmitMQ
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SubmitLanguage struct {
	Suffix      string `protobuf:"bytes,1,opt,name=suffix" json:"suffix,omitempty"`
	Compiler    string `protobuf:"bytes,2,opt,name=compiler" json:"compiler,omitempty"`
	Lang        string `protobuf:"bytes,3,opt,name=lang" json:"lang,omitempty"`
	OptionValue string `protobuf:"bytes,4,opt,name=option_value,json=optionValue" json:"option_value,omitempty"`
}

func (m *SubmitLanguage) Reset()                    { *m = SubmitLanguage{} }
func (m *SubmitLanguage) String() string            { return proto.CompactTextString(m) }
func (*SubmitLanguage) ProtoMessage()               {}
func (*SubmitLanguage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SubmitLanguage) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

func (m *SubmitLanguage) GetCompiler() string {
	if m != nil {
		return m.Compiler
	}
	return ""
}

func (m *SubmitLanguage) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *SubmitLanguage) GetOptionValue() string {
	if m != nil {
		return m.OptionValue
	}
	return ""
}

type TestCase struct {
	CaseId     int64  `protobuf:"varint,1,opt,name=case_id,json=caseId" json:"case_id,omitempty"`
	InputHash  []byte `protobuf:"bytes,2,opt,name=input_hash,json=inputHash,proto3" json:"input_hash,omitempty"`
	OutputHash []byte `protobuf:"bytes,3,opt,name=output_hash,json=outputHash,proto3" json:"output_hash,omitempty"`
	Input      []byte `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	Output     []byte `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
}

func (m *TestCase) Reset()                    { *m = TestCase{} }
func (m *TestCase) String() string            { return proto.CompactTextString(m) }
func (*TestCase) ProtoMessage()               {}
func (*TestCase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestCase) GetCaseId() int64 {
	if m != nil {
		return m.CaseId
	}
	return 0
}

func (m *TestCase) GetInputHash() []byte {
	if m != nil {
		return m.InputHash
	}
	return nil
}

func (m *TestCase) GetOutputHash() []byte {
	if m != nil {
		return m.OutputHash
	}
	return nil
}

func (m *TestCase) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *TestCase) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

type SpjCode struct {
	LocalPid int64  `protobuf:"varint,1,opt,name=local_pid,json=localPid" json:"local_pid,omitempty"`
	Hash     []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
}

func (m *SpjCode) Reset()                    { *m = SpjCode{} }
func (m *SpjCode) String() string            { return proto.CompactTextString(m) }
func (*SpjCode) ProtoMessage()               {}
func (*SpjCode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SpjCode) GetLocalPid() int64 {
	if m != nil {
		return m.LocalPid
	}
	return 0
}

func (m *SpjCode) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *SpjCode) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type SubmitMQ struct {
	RunId       int64           `protobuf:"varint,1,opt,name=run_id,json=runId" json:"run_id,omitempty"`
	OjName      string          `protobuf:"bytes,2,opt,name=oj_name,json=ojName" json:"oj_name,omitempty"`
	OjPid       string          `protobuf:"bytes,3,opt,name=oj_pid,json=ojPid" json:"oj_pid,omitempty"`
	Code        string          `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
	TimeLimit   int32           `protobuf:"varint,5,opt,name=time_limit,json=timeLimit" json:"time_limit,omitempty"`
	MemoryLimit int32           `protobuf:"varint,6,opt,name=memory_limit,json=memoryLimit" json:"memory_limit,omitempty"`
	IsLocal     bool            `protobuf:"varint,7,opt,name=is_local,json=isLocal" json:"is_local,omitempty"`
	IsSpj       bool            `protobuf:"varint,8,opt,name=is_spj,json=isSpj" json:"is_spj,omitempty"`
	SpjCode     string          `protobuf:"bytes,9,opt,name=spj_code,json=spjCode" json:"spj_code,omitempty"`
	SubmitTime  string          `protobuf:"bytes,10,opt,name=submit_time,json=submitTime" json:"submit_time,omitempty"`
	Language    *SubmitLanguage `protobuf:"bytes,11,opt,name=language" json:"language,omitempty"`
	Testcases   []*TestCase     `protobuf:"bytes,12,rep,name=testcases" json:"testcases,omitempty"`
}

func (m *SubmitMQ) Reset()                    { *m = SubmitMQ{} }
func (m *SubmitMQ) String() string            { return proto.CompactTextString(m) }
func (*SubmitMQ) ProtoMessage()               {}
func (*SubmitMQ) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SubmitMQ) GetRunId() int64 {
	if m != nil {
		return m.RunId
	}
	return 0
}

func (m *SubmitMQ) GetOjName() string {
	if m != nil {
		return m.OjName
	}
	return ""
}

func (m *SubmitMQ) GetOjPid() string {
	if m != nil {
		return m.OjPid
	}
	return ""
}

func (m *SubmitMQ) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SubmitMQ) GetTimeLimit() int32 {
	if m != nil {
		return m.TimeLimit
	}
	return 0
}

func (m *SubmitMQ) GetMemoryLimit() int32 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *SubmitMQ) GetIsLocal() bool {
	if m != nil {
		return m.IsLocal
	}
	return false
}

func (m *SubmitMQ) GetIsSpj() bool {
	if m != nil {
		return m.IsSpj
	}
	return false
}

func (m *SubmitMQ) GetSpjCode() string {
	if m != nil {
		return m.SpjCode
	}
	return ""
}

func (m *SubmitMQ) GetSubmitTime() string {
	if m != nil {
		return m.SubmitTime
	}
	return ""
}

func (m *SubmitMQ) GetLanguage() *SubmitLanguage {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *SubmitMQ) GetTestcases() []*TestCase {
	if m != nil {
		return m.Testcases
	}
	return nil
}

func init() {
	proto.RegisterType((*SubmitLanguage)(nil), "messages.SubmitLanguage")
	proto.RegisterType((*TestCase)(nil), "messages.TestCase")
	proto.RegisterType((*SpjCode)(nil), "messages.SpjCode")
	proto.RegisterType((*SubmitMQ)(nil), "messages.SubmitMQ")
}

func init() { proto.RegisterFile("messages/mq.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x92, 0xdd, 0x8e, 0xd3, 0x30,
	0x10, 0x85, 0xd5, 0x6d, 0xf3, 0x37, 0xa9, 0x90, 0xb0, 0xf8, 0x31, 0x20, 0xb4, 0xa5, 0x57, 0xbd,
	0x2a, 0x68, 0xe1, 0x0d, 0xf6, 0x86, 0x95, 0xca, 0x0a, 0xd2, 0x15, 0xb7, 0x96, 0xb7, 0xf1, 0xb6,
	0x8e, 0xe2, 0xd8, 0x64, 0x1c, 0x04, 0xe2, 0x29, 0x78, 0x60, 0x24, 0xe4, 0x71, 0xda, 0xc2, 0x9d,
	0xe7, 0x9b, 0x93, 0xce, 0x39, 0xd3, 0x81, 0xc7, 0x46, 0x21, 0xca, 0xbd, 0xc2, 0xb7, 0xe6, 0xdb,
	0xda, 0xf5, 0xd6, 0x5b, 0x96, 0x1f, 0xd1, 0xf2, 0x17, 0x3c, 0xda, 0x0e, 0xf7, 0x46, 0xfb, 0x8d,
	0xec, 0xf6, 0x83, 0xdc, 0x2b, 0xf6, 0x0c, 0x52, 0x1c, 0x1e, 0x1e, 0xf4, 0x0f, 0x3e, 0x59, 0x4c,
	0x56, 0x45, 0x35, 0x56, 0xec, 0x25, 0xe4, 0x3b, 0x6b, 0x9c, 0x6e, 0x55, 0xcf, 0x2f, 0xa8, 0x73,
	0xaa, 0x19, 0x83, 0x59, 0x2b, 0xbb, 0x3d, 0x9f, 0x12, 0xa7, 0x37, 0x7b, 0x03, 0x73, 0xeb, 0xbc,
	0xb6, 0x9d, 0xf8, 0x2e, 0xdb, 0x41, 0xf1, 0x19, 0xf5, 0xca, 0xc8, 0xbe, 0x06, 0xb4, 0xfc, 0x3d,
	0x81, 0xfc, 0x4e, 0xa1, 0xbf, 0x96, 0xa8, 0xd8, 0x73, 0xc8, 0x76, 0x12, 0x95, 0xd0, 0x35, 0x0d,
	0x9e, 0x56, 0x69, 0x28, 0x6f, 0x6a, 0xf6, 0x1a, 0x40, 0x77, 0x6e, 0xf0, 0xe2, 0x20, 0xf1, 0x40,
	0xa3, 0xe7, 0x55, 0x41, 0xe4, 0xa3, 0xc4, 0x03, 0xbb, 0x84, 0xd2, 0x0e, 0xfe, 0xd4, 0x9f, 0x52,
	0x1f, 0x22, 0x22, 0xc1, 0x13, 0x48, 0x48, 0x4d, 0x0e, 0xe6, 0x55, 0x2c, 0x42, 0xcc, 0xa8, 0xe1,
	0x09, 0xe1, 0xb1, 0x5a, 0xde, 0x42, 0xb6, 0x75, 0xcd, 0xb5, 0xad, 0x15, 0x7b, 0x05, 0x45, 0x6b,
	0x77, 0xb2, 0x15, 0xee, 0xe4, 0x29, 0x27, 0xf0, 0x59, 0xd7, 0x21, 0xf2, 0x3f, 0x7e, 0xe8, 0x1d,
	0xd8, 0xce, 0xd6, 0xea, 0xb8, 0x86, 0xf0, 0x5e, 0xfe, 0xb9, 0x80, 0x3c, 0x6e, 0xf8, 0xd3, 0x17,
	0xf6, 0x14, 0xd2, 0x7e, 0xe8, 0xce, 0x11, 0x93, 0x7e, 0xe8, 0x6e, 0xea, 0x10, 0xdd, 0x36, 0xa2,
	0x93, 0x46, 0x8d, 0x9b, 0x4d, 0x6d, 0x73, 0x2b, 0x8d, 0x0a, 0x7a, 0xdb, 0xd0, 0xf8, 0xf8, 0x93,
	0x89, 0x6d, 0xc6, 0xd9, 0x34, 0x67, 0x76, 0x9e, 0x13, 0xb6, 0xe4, 0xb5, 0x51, 0xa2, 0xd5, 0x46,
	0xc7, 0x4c, 0x49, 0x55, 0x04, 0xb2, 0x09, 0x20, 0xfc, 0x1b, 0x46, 0x19, 0xdb, 0xff, 0x1c, 0x05,
	0x29, 0x09, 0xca, 0xc8, 0xa2, 0xe4, 0x05, 0xe4, 0x1a, 0x05, 0x05, 0xe4, 0xd9, 0x62, 0xb2, 0xca,
	0xab, 0x4c, 0xe3, 0x26, 0x94, 0xc1, 0x87, 0x46, 0x81, 0xae, 0xe1, 0x39, 0x35, 0x12, 0x8d, 0x5b,
	0xd7, 0x84, 0x2f, 0xd0, 0x35, 0x82, 0xbc, 0x14, 0xe4, 0x25, 0xc3, 0x71, 0x77, 0x97, 0x50, 0x22,
	0xa5, 0x16, 0xc1, 0x03, 0x07, 0xea, 0x42, 0x44, 0x77, 0xda, 0x28, 0xf6, 0x01, 0xf2, 0x76, 0x3c,
	0x39, 0x5e, 0x2e, 0x26, 0xab, 0xf2, 0x8a, 0xaf, 0x8f, 0x57, 0xb9, 0xfe, 0xff, 0x24, 0xab, 0x93,
	0x92, 0xbd, 0x83, 0xc2, 0x2b, 0xf4, 0xe1, 0x32, 0x90, 0xcf, 0x17, 0xd3, 0x55, 0x79, 0xc5, 0xce,
	0x9f, 0x1d, 0x6f, 0xa9, 0x3a, 0x8b, 0xee, 0x53, 0xba, 0xf8, 0xf7, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xdf, 0x03, 0x3f, 0xb4, 0x06, 0x03, 0x00, 0x00,
}
