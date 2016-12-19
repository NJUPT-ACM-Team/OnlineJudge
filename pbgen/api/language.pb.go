// Code generated by protoc-gen-go.
// source: api/language.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Language struct {
	Compiler string `protobuf:"bytes,1,opt,name=compiler" json:"compiler,omitempty"`
	Language string `protobuf:"bytes,2,opt,name=language" json:"language,omitempty"`
	// lang_id
	LanguageId int64  `protobuf:"varint,3,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	OjName     string `protobuf:"bytes,4,opt,name=oj_name,json=ojName" json:"oj_name,omitempty"`
}

func (m *Language) Reset()                    { *m = Language{} }
func (m *Language) String() string            { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()               {}
func (*Language) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Language) GetCompiler() string {
	if m != nil {
		return m.Compiler
	}
	return ""
}

func (m *Language) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Language) GetLanguageId() int64 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *Language) GetOjName() string {
	if m != nil {
		return m.OjName
	}
	return ""
}

func init() {
	proto.RegisterType((*Language)(nil), "api.Language")
}

func init() { proto.RegisterFile("api/language.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2c, 0xc8, 0xd4,
	0xcf, 0x49, 0xcc, 0x4b, 0x2f, 0x4d, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4e, 0x2c, 0xc8, 0x54, 0xaa, 0xe1, 0xe2, 0xf0, 0x81, 0x0a, 0x0b, 0x49, 0x71, 0x71, 0x24, 0xe7,
	0xe7, 0x16, 0x64, 0xe6, 0xa4, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x20,
	0x39, 0x98, 0x76, 0x09, 0x26, 0x88, 0x1c, 0x8c, 0x2f, 0x24, 0xcf, 0xc5, 0x0d, 0x63, 0xc7, 0x67,
	0xa6, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0x30, 0x07, 0x71, 0xc1, 0x84, 0x3c, 0x53, 0x84, 0xc4, 0xb9,
	0xd8, 0xf3, 0xb3, 0xe2, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x58, 0xc0, 0x7a, 0xd9, 0xf2, 0xb3, 0xfc,
	0x12, 0x73, 0x53, 0x93, 0xd8, 0xc0, 0x2e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x84,
	0xb4, 0x81, 0x9f, 0x00, 0x00, 0x00,
}
