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
	RunId     int64           `protobuf:"varint,1,opt,name=run_id,json=runId" json:"run_id,omitempty"`
	OjName    string          `protobuf:"bytes,2,opt,name=oj_name,json=ojName" json:"oj_name,omitempty"`
	OjPid     string          `protobuf:"bytes,3,opt,name=oj_pid,json=ojPid" json:"oj_pid,omitempty"`
	Code      string          `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
	IsLocal   bool            `protobuf:"varint,5,opt,name=is_local,json=isLocal" json:"is_local,omitempty"`
	IsSpj     bool            `protobuf:"varint,6,opt,name=is_spj,json=isSpj" json:"is_spj,omitempty"`
	SpjCode   string          `protobuf:"bytes,7,opt,name=spj_code,json=spjCode" json:"spj_code,omitempty"`
	Language  *SubmitLanguage `protobuf:"bytes,8,opt,name=language" json:"language,omitempty"`
	Testcases []*TestCase     `protobuf:"bytes,9,rep,name=testcases" json:"testcases,omitempty"`
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
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xb7, 0x4d, 0xe2, 0x4c, 0x2b, 0x24, 0x2c, 0x3e, 0x0c, 0x08, 0x51, 0x7a, 0xea, 0xa9,
	0xa0, 0x85, 0x7f, 0xb0, 0x17, 0x56, 0x5a, 0x56, 0x90, 0x22, 0xae, 0x96, 0xb7, 0xf1, 0xb6, 0xb6,
	0x92, 0xd8, 0x74, 0x6c, 0x84, 0xc4, 0xaf, 0xe0, 0xca, 0xaf, 0x45, 0x9e, 0xa4, 0x0d, 0x7b, 0xf3,
	0x7b, 0xf3, 0x3c, 0x6f, 0x9e, 0x3d, 0xf0, 0xb8, 0xd5, 0x88, 0x6a, 0xaf, 0xf1, 0x5d, 0xfb, 0x63,
	0xe3, 0x8f, 0x2e, 0x38, 0xce, 0x4e, 0xd4, 0xea, 0x37, 0x3c, 0xda, 0xc6, 0xbb, 0xd6, 0x84, 0x1b,
	0xd5, 0xed, 0xa3, 0xda, 0x6b, 0xfe, 0x0c, 0x72, 0x8c, 0xf7, 0xf7, 0xe6, 0x97, 0x98, 0x2c, 0x27,
	0xeb, 0xb2, 0x1a, 0x10, 0x7f, 0x09, 0x6c, 0xe7, 0x5a, 0x6f, 0x1a, 0x7d, 0x14, 0x17, 0x54, 0x39,
	0x63, 0xce, 0x61, 0xd6, 0xa8, 0x6e, 0x2f, 0xa6, 0xc4, 0xd3, 0x99, 0xbf, 0x85, 0x85, 0xf3, 0xc1,
	0xb8, 0x4e, 0xfe, 0x54, 0x4d, 0xd4, 0x62, 0x46, 0xb5, 0x79, 0xcf, 0x7d, 0x4f, 0xd4, 0xea, 0xcf,
	0x04, 0xd8, 0x37, 0x8d, 0xe1, 0x4a, 0xa1, 0xe6, 0xcf, 0xa1, 0xd8, 0x29, 0xd4, 0xd2, 0xd4, 0x64,
	0x3c, 0xad, 0xf2, 0x04, 0xaf, 0x6b, 0xfe, 0x1a, 0xc0, 0x74, 0x3e, 0x06, 0x79, 0x50, 0x78, 0x20,
	0xeb, 0x45, 0x55, 0x12, 0xf3, 0x49, 0xe1, 0x81, 0xbf, 0x81, 0xb9, 0x8b, 0xe1, 0x5c, 0x9f, 0x52,
	0x1d, 0x7a, 0x8a, 0x04, 0x4f, 0x20, 0x23, 0x35, 0x4d, 0xb0, 0xa8, 0x7a, 0x90, 0x62, 0xf6, 0x1a,
	0x91, 0x11, 0x3d, 0xa0, 0xd5, 0x2d, 0x14, 0x5b, 0x6f, 0xaf, 0x5c, 0xad, 0xf9, 0x2b, 0x28, 0x1b,
	0xb7, 0x53, 0x8d, 0xf4, 0xe7, 0x99, 0x18, 0x11, 0x5f, 0x4c, 0x9d, 0x22, 0xff, 0x37, 0x0f, 0x9d,
	0x13, 0xb7, 0x73, 0xb5, 0x3e, 0x3d, 0x43, 0x3a, 0xaf, 0xfe, 0x5e, 0x00, 0xeb, 0x5f, 0xf8, 0xf3,
	0x57, 0xfe, 0x14, 0xf2, 0x63, 0xec, 0xc6, 0x88, 0xd9, 0x31, 0x76, 0xd7, 0x75, 0x8a, 0xee, 0xac,
	0xec, 0x54, 0xab, 0x87, 0x97, 0xcd, 0x9d, 0xbd, 0x55, 0xad, 0x4e, 0x7a, 0x67, 0xc9, 0xbe, 0x6f,
	0x99, 0x39, 0x3b, 0x78, 0x93, 0xcf, 0x6c, 0xf4, 0xe1, 0x2f, 0x80, 0x19, 0x94, 0x34, 0x1e, 0x25,
	0x62, 0x55, 0x61, 0xf0, 0x26, 0xc1, 0xd4, 0xc5, 0xa0, 0x44, 0x6f, 0x45, 0x4e, 0x85, 0xcc, 0xe0,
	0xd6, 0xdb, 0x74, 0x03, 0xbd, 0x95, 0xd4, 0xa9, 0xa0, 0x4e, 0x05, 0x0e, 0xc9, 0x3f, 0x02, 0x6b,
	0x86, 0x7d, 0x10, 0x6c, 0x39, 0x59, 0xcf, 0x2f, 0xc5, 0xe6, 0xb4, 0x32, 0x9b, 0x87, 0xfb, 0x52,
	0x9d, 0x95, 0xfc, 0x3d, 0x94, 0x41, 0x63, 0x48, 0xdf, 0x86, 0xa2, 0x5c, 0x4e, 0xd7, 0xf3, 0x4b,
	0x3e, 0x5e, 0x3b, 0x7d, 0x74, 0x35, 0x8a, 0xee, 0x72, 0x5a, 0xc7, 0x0f, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa2, 0x13, 0xca, 0x38, 0xa3, 0x02, 0x00, 0x00,
}
