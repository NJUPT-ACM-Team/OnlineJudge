// Code generated by protoc-gen-go.
// source: api/webapi.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WebGetRequest struct {
	ListProblemsRequest    *ListProblemsRequest    `protobuf:"bytes,1,opt,name=list_problems_request,json=listProblemsRequest" json:"list_problems_request,omitempty"`
	ListSubmissionsRequest *ListSubmissionsRequest `protobuf:"bytes,2,opt,name=list_submissions_request,json=listSubmissionsRequest" json:"list_submissions_request,omitempty"`
	ListContestsRequest    *ListContestsRequest    `protobuf:"bytes,3,opt,name=list_contests_request,json=listContestsRequest" json:"list_contests_request,omitempty"`
	ShowProblemRequest     *ShowProblemRequest     `protobuf:"bytes,4,opt,name=show_problem_request,json=showProblemRequest" json:"show_problem_request,omitempty"`
	AboutRequest           *AboutRequest           `protobuf:"bytes,5,opt,name=about_request,json=aboutRequest" json:"about_request,omitempty"`
	ShowContest            *ShowContestRequest     `protobuf:"bytes,6,opt,name=show_contest,json=showContest" json:"show_contest,omitempty"`
}

func (m *WebGetRequest) Reset()                    { *m = WebGetRequest{} }
func (m *WebGetRequest) String() string            { return proto.CompactTextString(m) }
func (*WebGetRequest) ProtoMessage()               {}
func (*WebGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *WebGetRequest) GetListProblemsRequest() *ListProblemsRequest {
	if m != nil {
		return m.ListProblemsRequest
	}
	return nil
}

func (m *WebGetRequest) GetListSubmissionsRequest() *ListSubmissionsRequest {
	if m != nil {
		return m.ListSubmissionsRequest
	}
	return nil
}

func (m *WebGetRequest) GetListContestsRequest() *ListContestsRequest {
	if m != nil {
		return m.ListContestsRequest
	}
	return nil
}

func (m *WebGetRequest) GetShowProblemRequest() *ShowProblemRequest {
	if m != nil {
		return m.ShowProblemRequest
	}
	return nil
}

func (m *WebGetRequest) GetAboutRequest() *AboutRequest {
	if m != nil {
		return m.AboutRequest
	}
	return nil
}

func (m *WebGetRequest) GetShowContest() *ShowContestRequest {
	if m != nil {
		return m.ShowContest
	}
	return nil
}

// WebRequest are all POST requests
type WebPostRequest struct {
	CsrfToken        string            `protobuf:"bytes,1,opt,name=csrf_token,json=csrfToken" json:"csrf_token,omitempty"`
	Captcha          string            `protobuf:"bytes,2,opt,name=captcha" json:"captcha,omitempty"`
	LoginInitRequest *LoginInitRequest `protobuf:"bytes,3,opt,name=login_init_request,json=loginInitRequest" json:"login_init_request,omitempty"`
	LoginAuthRequest *LoginAuthRequest `protobuf:"bytes,4,opt,name=login_auth_request,json=loginAuthRequest" json:"login_auth_request,omitempty"`
	RegisterRequest  *RegisterRequest  `protobuf:"bytes,5,opt,name=register_request,json=registerRequest" json:"register_request,omitempty"`
	SubmitRequest    *SubmitRequest    `protobuf:"bytes,6,opt,name=submit_request,json=submitRequest" json:"submit_request,omitempty"`
	LogoutRequest    *LogoutRequest    `protobuf:"bytes,7,opt,name=logout_request,json=logoutRequest" json:"logout_request,omitempty"`
}

func (m *WebPostRequest) Reset()                    { *m = WebPostRequest{} }
func (m *WebPostRequest) String() string            { return proto.CompactTextString(m) }
func (*WebPostRequest) ProtoMessage()               {}
func (*WebPostRequest) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *WebPostRequest) GetCsrfToken() string {
	if m != nil {
		return m.CsrfToken
	}
	return ""
}

func (m *WebPostRequest) GetCaptcha() string {
	if m != nil {
		return m.Captcha
	}
	return ""
}

func (m *WebPostRequest) GetLoginInitRequest() *LoginInitRequest {
	if m != nil {
		return m.LoginInitRequest
	}
	return nil
}

func (m *WebPostRequest) GetLoginAuthRequest() *LoginAuthRequest {
	if m != nil {
		return m.LoginAuthRequest
	}
	return nil
}

func (m *WebPostRequest) GetRegisterRequest() *RegisterRequest {
	if m != nil {
		return m.RegisterRequest
	}
	return nil
}

func (m *WebPostRequest) GetSubmitRequest() *SubmitRequest {
	if m != nil {
		return m.SubmitRequest
	}
	return nil
}

func (m *WebPostRequest) GetLogoutRequest() *LogoutRequest {
	if m != nil {
		return m.LogoutRequest
	}
	return nil
}

type WebResponse struct {
	SetCsrfToken            string                   `protobuf:"bytes,1,opt,name=set_csrf_token,json=setCsrfToken" json:"set_csrf_token,omitempty"`
	CaptchaUrl              string                   `protobuf:"bytes,2,opt,name=captcha_url,json=captchaUrl" json:"captcha_url,omitempty"`
	ListProblemsResponse    *ListProblemsResponse    `protobuf:"bytes,3,opt,name=list_problems_response,json=listProblemsResponse" json:"list_problems_response,omitempty"`
	ListSubmissionsResponse *ListSubmissionsResponse `protobuf:"bytes,4,opt,name=list_submissions_response,json=listSubmissionsResponse" json:"list_submissions_response,omitempty"`
	ListContestsResponse    *ListContestsResponse    `protobuf:"bytes,5,opt,name=list_contests_response,json=listContestsResponse" json:"list_contests_response,omitempty"`
	LoginInitResponse       *LoginInitResponse       `protobuf:"bytes,6,opt,name=login_init_response,json=loginInitResponse" json:"login_init_response,omitempty"`
	LoginAuthResponse       *LoginAuthResponse       `protobuf:"bytes,7,opt,name=login_auth_response,json=loginAuthResponse" json:"login_auth_response,omitempty"`
	RegisterResponse        *RegisterResponse        `protobuf:"bytes,8,opt,name=register_response,json=registerResponse" json:"register_response,omitempty"`
	ShowProblemResponse     *ShowProblemResponse     `protobuf:"bytes,9,opt,name=show_problem_response,json=showProblemResponse" json:"show_problem_response,omitempty"`
	SubmitResponse          *SubmitResponse          `protobuf:"bytes,10,opt,name=submit_response,json=submitResponse" json:"submit_response,omitempty"`
	AboutResponse           *AboutResponse           `protobuf:"bytes,11,opt,name=about_response,json=aboutResponse" json:"about_response,omitempty"`
	LogoutResponse          *LogoutResponse          `protobuf:"bytes,12,opt,name=logout_response,json=logoutResponse" json:"logout_response,omitempty"`
	ShowContestResponse     *ShowContestResponse     `protobuf:"bytes,13,opt,name=show_contest_response,json=showContestResponse" json:"show_contest_response,omitempty"`
	Error                   *Error                   `protobuf:"bytes,14,opt,name=error" json:"error,omitempty"`
}

func (m *WebResponse) Reset()                    { *m = WebResponse{} }
func (m *WebResponse) String() string            { return proto.CompactTextString(m) }
func (*WebResponse) ProtoMessage()               {}
func (*WebResponse) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

func (m *WebResponse) GetSetCsrfToken() string {
	if m != nil {
		return m.SetCsrfToken
	}
	return ""
}

func (m *WebResponse) GetCaptchaUrl() string {
	if m != nil {
		return m.CaptchaUrl
	}
	return ""
}

func (m *WebResponse) GetListProblemsResponse() *ListProblemsResponse {
	if m != nil {
		return m.ListProblemsResponse
	}
	return nil
}

func (m *WebResponse) GetListSubmissionsResponse() *ListSubmissionsResponse {
	if m != nil {
		return m.ListSubmissionsResponse
	}
	return nil
}

func (m *WebResponse) GetListContestsResponse() *ListContestsResponse {
	if m != nil {
		return m.ListContestsResponse
	}
	return nil
}

func (m *WebResponse) GetLoginInitResponse() *LoginInitResponse {
	if m != nil {
		return m.LoginInitResponse
	}
	return nil
}

func (m *WebResponse) GetLoginAuthResponse() *LoginAuthResponse {
	if m != nil {
		return m.LoginAuthResponse
	}
	return nil
}

func (m *WebResponse) GetRegisterResponse() *RegisterResponse {
	if m != nil {
		return m.RegisterResponse
	}
	return nil
}

func (m *WebResponse) GetShowProblemResponse() *ShowProblemResponse {
	if m != nil {
		return m.ShowProblemResponse
	}
	return nil
}

func (m *WebResponse) GetSubmitResponse() *SubmitResponse {
	if m != nil {
		return m.SubmitResponse
	}
	return nil
}

func (m *WebResponse) GetAboutResponse() *AboutResponse {
	if m != nil {
		return m.AboutResponse
	}
	return nil
}

func (m *WebResponse) GetLogoutResponse() *LogoutResponse {
	if m != nil {
		return m.LogoutResponse
	}
	return nil
}

func (m *WebResponse) GetShowContestResponse() *ShowContestResponse {
	if m != nil {
		return m.ShowContestResponse
	}
	return nil
}

func (m *WebResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*WebGetRequest)(nil), "api.WebGetRequest")
	proto.RegisterType((*WebPostRequest)(nil), "api.WebPostRequest")
	proto.RegisterType((*WebResponse)(nil), "api.WebResponse")
}

func init() { proto.RegisterFile("api/webapi.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x95, 0xdd, 0x6e, 0x13, 0x3d,
	0x10, 0x86, 0xd5, 0x2f, 0x5f, 0x1b, 0x32, 0xf9, 0x69, 0xea, 0xb4, 0x69, 0x5a, 0x40, 0x54, 0x15,
	0x07, 0x1c, 0x15, 0x09, 0x24, 0x24, 0x10, 0x12, 0x2a, 0x15, 0xa0, 0x4a, 0x95, 0xa8, 0xb6, 0x54,
	0xe5, 0x2c, 0xda, 0x8d, 0x4c, 0x63, 0xe1, 0xac, 0x17, 0xdb, 0x51, 0xef, 0x81, 0x3b, 0xe0, 0x72,
	0xb8, 0x33, 0x64, 0x7b, 0xec, 0xf5, 0x3a, 0x7b, 0xe8, 0x77, 0xf6, 0x7d, 0x3c, 0x99, 0x1f, 0x07,
	0xc6, 0x79, 0xc5, 0x5e, 0x3e, 0xd0, 0x22, 0xaf, 0xd8, 0x59, 0x25, 0x85, 0x16, 0xa4, 0x93, 0x57,
	0xec, 0xd8, 0xca, 0x0b, 0xb1, 0x5a, 0x89, 0xd2, 0xc9, 0xc7, 0x87, 0x46, 0xe1, 0x4c, 0xe9, 0x79,
	0x25, 0x45, 0xc1, 0xe9, 0x4a, 0x61, 0xe0, 0x38, 0x04, 0xd4, 0xba, 0x58, 0x31, 0xa5, 0x98, 0x28,
	0xd5, 0x86, 0x69, 0x21, 0x4a, 0x4d, 0x95, 0xf6, 0x81, 0x5d, 0x1b, 0x10, 0xf7, 0xcc, 0xe3, 0x89,
	0x11, 0x24, 0xbd, 0x67, 0x4a, 0x53, 0x89, 0xda, 0xd4, 0x68, 0x6a, 0x29, 0x1e, 0xfc, 0x95, 0x1b,
	0x3a, 0x52, 0x51, 0xb7, 0x49, 0xdb, 0x24, 0x74, 0x7c, 0x4d, 0x5e, 0x88, 0x75, 0xe3, 0x13, 0x2e,
	0xee, 0x83, 0x72, 0xfa, 0xb7, 0x03, 0xc3, 0x3b, 0x5a, 0x7c, 0xa1, 0x3a, 0xa3, 0xbf, 0xd6, 0x54,
	0x69, 0x72, 0x05, 0x07, 0x8d, 0xdf, 0x39, 0x97, 0x2e, 0x30, 0xdb, 0x3a, 0xd9, 0x7a, 0xd1, 0x7f,
	0x35, 0x3b, 0x33, 0xb5, 0xba, 0x62, 0x4a, 0x5f, 0xe3, 0x07, 0x68, 0xcc, 0x26, 0x7c, 0x53, 0x24,
	0xb7, 0x30, 0x4b, 0x8b, 0x13, 0x80, 0xff, 0x59, 0xe0, 0xe3, 0x00, 0xbc, 0xa9, 0xbf, 0xf1, 0xcc,
	0x29, 0x6f, 0xd5, 0x43, 0x92, 0xbe, 0xae, 0x81, 0xd9, 0x49, 0x92, 0xbc, 0xc0, 0x0f, 0x1a, 0x49,
	0x26, 0x22, 0xb9, 0x84, 0xfd, 0xb8, 0xce, 0x01, 0xf6, 0xbf, 0x85, 0x1d, 0x5a, 0xd8, 0xcd, 0x52,
	0x3c, 0xe0, 0x8f, 0xf3, 0x2c, 0xa2, 0x36, 0x34, 0xf2, 0x06, 0x86, 0xb6, 0xe0, 0x81, 0xb1, 0x6d,
	0x19, 0x7b, 0x96, 0x71, 0x6e, 0x22, 0xde, 0x3d, 0xc8, 0xa3, 0x13, 0x79, 0x07, 0x83, 0xb8, 0xa5,
	0xb3, 0x9d, 0xe4, 0x6a, 0x4c, 0xd9, 0x9b, 0xfb, 0xaa, 0xd6, 0x4e, 0x7f, 0x77, 0x60, 0x74, 0x47,
	0x8b, 0x6b, 0x11, 0xe2, 0xe4, 0x29, 0xc0, 0x42, 0xc9, 0x1f, 0x73, 0x2d, 0x7e, 0xd2, 0xd2, 0x76,
	0xae, 0x97, 0xf5, 0x8c, 0xf2, 0xcd, 0x08, 0x64, 0x06, 0xdd, 0x45, 0x5e, 0xe9, 0xc5, 0x32, 0xb7,
	0x4d, 0xe8, 0x65, 0xfe, 0x48, 0x2e, 0x80, 0xd8, 0xb9, 0x9c, 0xb3, 0x92, 0xe9, 0xa4, 0xaa, 0x07,
	0xae, 0xaa, 0x26, 0x7c, 0x59, 0xb2, 0x90, 0xcb, 0x98, 0x27, 0x4a, 0x0d, 0xc9, 0xd7, 0x7a, 0x99,
	0x54, 0x33, 0x82, 0x9c, 0xaf, 0xf5, 0xb2, 0x09, 0x89, 0x14, 0xf2, 0x01, 0xc6, 0x7e, 0x21, 0x92,
	0x62, 0xee, 0x5b, 0x44, 0x86, 0x41, 0x4f, 0xd8, 0x95, 0x4d, 0x81, 0xbc, 0x85, 0x91, 0xdb, 0x86,
	0x60, 0x77, 0x45, 0x25, 0xae, 0xa8, 0x36, 0xe4, 0xcd, 0x43, 0x15, 0x1f, 0x8d, 0xd5, 0x6d, 0x49,
	0xb0, 0x76, 0x23, 0xeb, 0x95, 0x0d, 0x05, 0x2b, 0x8f, 0x8f, 0xa7, 0x7f, 0xba, 0xd0, 0xbf, 0xa3,
	0x45, 0x46, 0x55, 0x25, 0x4a, 0x45, 0xc9, 0x73, 0x18, 0x29, 0xaa, 0xe7, 0x1b, 0xdd, 0x18, 0x28,
	0xaa, 0x2f, 0x42, 0x43, 0x9e, 0x41, 0x1f, 0x3b, 0x30, 0x5f, 0x4b, 0x8e, 0x4d, 0x01, 0x94, 0x6e,
	0x25, 0x27, 0x5f, 0x61, 0x9a, 0x6e, 0xa5, 0xbb, 0x00, 0x7b, 0x73, 0xd4, 0xb2, 0x96, 0xee, 0x83,
	0x6c, 0x9f, 0xb7, 0xa8, 0xe4, 0x3b, 0x1c, 0xb5, 0x2c, 0x26, 0x32, 0x5d, 0xab, 0x9e, 0xb4, 0x6f,
	0x26, 0x62, 0x0f, 0x79, 0x7b, 0x20, 0xa4, 0x1a, 0xed, 0x26, 0x62, 0xb7, 0x93, 0x54, 0xeb, 0x3d,
	0x8c, 0x53, 0x4d, 0x55, 0xf2, 0x19, 0x26, 0x8d, 0x99, 0x44, 0x9a, 0xeb, 0xe6, 0x34, 0x1d, 0x4a,
	0x44, 0xed, 0xf1, 0x54, 0xaa, 0x39, 0x38, 0x96, 0xc8, 0xe9, 0xa6, 0x1c, 0x37, 0x85, 0x0d, 0x4e,
	0x2c, 0x91, 0x8f, 0xb0, 0x17, 0x4d, 0x26, 0x52, 0x1e, 0x45, 0xd3, 0x5d, 0x8f, 0x26, 0x42, 0xc6,
	0x32, 0x51, 0xcc, 0x03, 0x96, 0x3c, 0x39, 0xc8, 0xe9, 0x45, 0x0f, 0x58, 0xe3, 0xcd, 0x41, 0xd4,
	0x44, 0x6d, 0x8a, 0xe4, 0x3d, 0xec, 0x86, 0x51, 0x47, 0x0e, 0x58, 0xce, 0xa4, 0x31, 0xeb, 0x88,
	0x18, 0xa9, 0xc6, 0xd9, 0x4c, 0xbb, 0x7f, 0xb3, 0xd0, 0xdc, 0x8f, 0xa6, 0x1d, 0x1f, 0x2d, 0xf4,
	0x0e, 0xf3, 0xf8, 0x68, 0x2e, 0x0e, 0x8b, 0x82, 0xde, 0x41, 0x74, 0xb1, 0xdf, 0x14, 0x7f, 0x31,
	0x6f, 0x9c, 0x43, 0x11, 0x70, 0x52, 0x6a, 0xc6, 0x30, 0x29, 0x42, 0x78, 0xfd, 0xe2, 0x22, 0x24,
	0x22, 0x39, 0x81, 0x6d, 0x2a, 0xa5, 0x90, 0xb3, 0x91, 0x75, 0x83, 0x75, 0x7f, 0x32, 0x4a, 0xe6,
	0x02, 0xc5, 0x8e, 0xfd, 0xcf, 0x7b, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x26, 0xf3, 0x35, 0x7a,
	0xf6, 0x07, 0x00, 0x00,
}
