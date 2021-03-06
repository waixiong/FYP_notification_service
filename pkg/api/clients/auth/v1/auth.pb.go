// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package authproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SignInMethod int32

const (
	SignInMethod_EMAIL SignInMethod = 0
	SignInMethod_PHONE SignInMethod = 1
)

var SignInMethod_name = map[int32]string{
	0: "EMAIL",
	1: "PHONE",
}

var SignInMethod_value = map[string]int32{
	"EMAIL": 0,
	"PHONE": 1,
}

func (x SignInMethod) String() string {
	return proto.EnumName(SignInMethod_name, int32(x))
}

func (SignInMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

type VerifyState int32

const (
	VerifyState_SUCCESS VerifyState = 0
	VerifyState_REFRESH VerifyState = 1
	VerifyState_ILLEGAL VerifyState = 2
)

var VerifyState_name = map[int32]string{
	0: "SUCCESS",
	1: "REFRESH",
	2: "ILLEGAL",
}

var VerifyState_value = map[string]int32{
	"SUCCESS": 0,
	"REFRESH": 1,
	"ILLEGAL": 2,
}

func (x VerifyState) String() string {
	return proto.EnumName(VerifyState_name, int32(x))
}

func (VerifyState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

type ThirdPartyAuth_Provider int32

const (
	ThirdPartyAuth_GOOGLE   ThirdPartyAuth_Provider = 0
	ThirdPartyAuth_FACEBOOK ThirdPartyAuth_Provider = 1
)

var ThirdPartyAuth_Provider_name = map[int32]string{
	0: "GOOGLE",
	1: "FACEBOOK",
}

var ThirdPartyAuth_Provider_value = map[string]int32{
	"GOOGLE":   0,
	"FACEBOOK": 1,
}

func (x ThirdPartyAuth_Provider) String() string {
	return proto.EnumName(ThirdPartyAuth_Provider_name, int32(x))
}

func (ThirdPartyAuth_Provider) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0, 0}
}

// The request message sign in detail.
type ThirdPartyAuth struct {
	Provider             ThirdPartyAuth_Provider `protobuf:"varint,1,opt,name=provider,proto3,enum=authproto.ThirdPartyAuth_Provider" json:"provider,omitempty"`
	IdToken              string                  `protobuf:"bytes,2,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ThirdPartyAuth) Reset()         { *m = ThirdPartyAuth{} }
func (m *ThirdPartyAuth) String() string { return proto.CompactTextString(m) }
func (*ThirdPartyAuth) ProtoMessage()    {}
func (*ThirdPartyAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *ThirdPartyAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThirdPartyAuth.Unmarshal(m, b)
}
func (m *ThirdPartyAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThirdPartyAuth.Marshal(b, m, deterministic)
}
func (m *ThirdPartyAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThirdPartyAuth.Merge(m, src)
}
func (m *ThirdPartyAuth) XXX_Size() int {
	return xxx_messageInfo_ThirdPartyAuth.Size(m)
}
func (m *ThirdPartyAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_ThirdPartyAuth.DiscardUnknown(m)
}

var xxx_messageInfo_ThirdPartyAuth proto.InternalMessageInfo

func (m *ThirdPartyAuth) GetProvider() ThirdPartyAuth_Provider {
	if m != nil {
		return m.Provider
	}
	return ThirdPartyAuth_GOOGLE
}

func (m *ThirdPartyAuth) GetIdToken() string {
	if m != nil {
		return m.IdToken
	}
	return ""
}

// The request message sign in detail.
type GetItQECAuth struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetItQECAuth) Reset()         { *m = GetItQECAuth{} }
func (m *GetItQECAuth) String() string { return proto.CompactTextString(m) }
func (*GetItQECAuth) ProtoMessage()    {}
func (*GetItQECAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *GetItQECAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItQECAuth.Unmarshal(m, b)
}
func (m *GetItQECAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItQECAuth.Marshal(b, m, deterministic)
}
func (m *GetItQECAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItQECAuth.Merge(m, src)
}
func (m *GetItQECAuth) XXX_Size() int {
	return xxx_messageInfo_GetItQECAuth.Size(m)
}
func (m *GetItQECAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItQECAuth.DiscardUnknown(m)
}

var xxx_messageInfo_GetItQECAuth proto.InternalMessageInfo

func (m *GetItQECAuth) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetItQECAuth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The request message sign up detail.
type GetItQECRegister struct {
	Method               SignInMethod `protobuf:"varint,1,opt,name=method,proto3,enum=authproto.SignInMethod" json:"method,omitempty"`
	Username             string       `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string       `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetItQECRegister) Reset()         { *m = GetItQECRegister{} }
func (m *GetItQECRegister) String() string { return proto.CompactTextString(m) }
func (*GetItQECRegister) ProtoMessage()    {}
func (*GetItQECRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *GetItQECRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetItQECRegister.Unmarshal(m, b)
}
func (m *GetItQECRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetItQECRegister.Marshal(b, m, deterministic)
}
func (m *GetItQECRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetItQECRegister.Merge(m, src)
}
func (m *GetItQECRegister) XXX_Size() int {
	return xxx_messageInfo_GetItQECRegister.Size(m)
}
func (m *GetItQECRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_GetItQECRegister.DiscardUnknown(m)
}

var xxx_messageInfo_GetItQECRegister proto.InternalMessageInfo

func (m *GetItQECRegister) GetMethod() SignInMethod {
	if m != nil {
		return m.Method
	}
	return SignInMethod_EMAIL
}

func (m *GetItQECRegister) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetItQECRegister) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The response message containing result string
type AuthResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AuthResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type Token struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type Verification struct {
	UserId               string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Result               VerifyState `protobuf:"varint,2,opt,name=result,proto3,enum=authproto.VerifyState" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Verification) Reset()         { *m = Verification{} }
func (m *Verification) String() string { return proto.CompactTextString(m) }
func (*Verification) ProtoMessage()    {}
func (*Verification) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *Verification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Verification.Unmarshal(m, b)
}
func (m *Verification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Verification.Marshal(b, m, deterministic)
}
func (m *Verification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Verification.Merge(m, src)
}
func (m *Verification) XXX_Size() int {
	return xxx_messageInfo_Verification.Size(m)
}
func (m *Verification) XXX_DiscardUnknown() {
	xxx_messageInfo_Verification.DiscardUnknown(m)
}

var xxx_messageInfo_Verification proto.InternalMessageInfo

func (m *Verification) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Verification) GetResult() VerifyState {
	if m != nil {
		return m.Result
	}
	return VerifyState_SUCCESS
}

type SignInManagement struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInManagement) Reset()         { *m = SignInManagement{} }
func (m *SignInManagement) String() string { return proto.CompactTextString(m) }
func (*SignInManagement) ProtoMessage()    {}
func (*SignInManagement) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *SignInManagement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInManagement.Unmarshal(m, b)
}
func (m *SignInManagement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInManagement.Marshal(b, m, deterministic)
}
func (m *SignInManagement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInManagement.Merge(m, src)
}
func (m *SignInManagement) XXX_Size() int {
	return xxx_messageInfo_SignInManagement.Size(m)
}
func (m *SignInManagement) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInManagement.DiscardUnknown(m)
}

var xxx_messageInfo_SignInManagement proto.InternalMessageInfo

func (m *SignInManagement) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SignInManagement) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func init() {
	proto.RegisterEnum("authproto.SignInMethod", SignInMethod_name, SignInMethod_value)
	proto.RegisterEnum("authproto.VerifyState", VerifyState_name, VerifyState_value)
	proto.RegisterEnum("authproto.ThirdPartyAuth_Provider", ThirdPartyAuth_Provider_name, ThirdPartyAuth_Provider_value)
	proto.RegisterType((*ThirdPartyAuth)(nil), "authproto.ThirdPartyAuth")
	proto.RegisterType((*GetItQECAuth)(nil), "authproto.GetItQECAuth")
	proto.RegisterType((*GetItQECRegister)(nil), "authproto.GetItQECRegister")
	proto.RegisterType((*AuthResponse)(nil), "authproto.AuthResponse")
	proto.RegisterType((*Token)(nil), "authproto.Token")
	proto.RegisterType((*Verification)(nil), "authproto.Verification")
	proto.RegisterType((*SignInManagement)(nil), "authproto.SignInManagement")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 917 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x73, 0xdb, 0x44,
	0x14, 0xc7, 0x23, 0xa7, 0x75, 0xec, 0x8d, 0xc9, 0x88, 0x1d, 0x9a, 0xa4, 0x4e, 0x99, 0x0a, 0xb5,
	0xcc, 0x74, 0x3c, 0x58, 0x2a, 0x26, 0xa7, 0x30, 0x03, 0xd8, 0x19, 0x25, 0xf5, 0xd4, 0xc5, 0x46,
	0x6e, 0xcb, 0x81, 0x43, 0x46, 0x91, 0x5e, 0xa4, 0xa5, 0x96, 0x56, 0xdd, 0x5d, 0x25, 0x64, 0xb8,
	0xf1, 0x07, 0x30, 0x03, 0xdc, 0xf8, 0x13, 0xf8, 0x77, 0xb8, 0x71, 0xe6, 0x5f, 0xe0, 0xce, 0xec,
	0x6a, 0xed, 0xc8, 0x8d, 0x03, 0xc3, 0x8f, 0x93, 0xf6, 0x69, 0xdf, 0x7e, 0xbe, 0x6f, 0xdf, 0xbe,
	0xb7, 0x8b, 0x50, 0x50, 0x88, 0xc4, 0xc9, 0x19, 0x15, 0x14, 0x37, 0xe5, 0x58, 0x0d, 0xdb, 0xf7,
	0x62, 0x4a, 0xe3, 0x19, 0xb8, 0x41, 0x4e, 0xdc, 0x20, 0xcb, 0xa8, 0x08, 0x04, 0xa1, 0x19, 0x2f,
	0x1d, 0xdb, 0x1f, 0xa8, 0x4f, 0xd8, 0x8d, 0x21, 0xeb, 0xf2, 0x8b, 0x20, 0x8e, 0x81, 0xb9, 0x34,
	0x57, 0x1e, 0x2b, 0xbc, 0xf7, 0x34, 0x4b, 0x59, 0xa7, 0xc5, 0x99, 0x0b, 0x69, 0x2e, 0x2e, 0xcb,
	0x49, 0xfb, 0x07, 0x03, 0x6d, 0x3d, 0x4f, 0x08, 0x8b, 0x26, 0x01, 0x13, 0x97, 0xfd, 0x42, 0x24,
	0xf8, 0x13, 0xd4, 0xc8, 0x19, 0x3d, 0x27, 0x11, 0xb0, 0x5d, 0xc3, 0x32, 0x1e, 0x6d, 0xf5, 0x6c,
	0x67, 0x11, 0x99, 0xb3, 0xec, 0xec, 0x4c, 0xb4, 0xa7, 0xbf, 0x58, 0x83, 0xef, 0xa2, 0x06, 0x89,
	0x4e, 0x04, 0x7d, 0x05, 0xd9, 0x6e, 0xcd, 0x32, 0x1e, 0x35, 0xfd, 0x0d, 0x12, 0x3d, 0x97, 0xa6,
	0xfd, 0x10, 0x35, 0xe6, 0x0b, 0x30, 0x42, 0xf5, 0xe3, 0xf1, 0xf8, 0x78, 0xe4, 0x99, 0x6b, 0xb8,
	0x85, 0x1a, 0x47, 0xfd, 0x43, 0x6f, 0x30, 0x1e, 0x3f, 0x35, 0x0d, 0xfb, 0x08, 0xb5, 0x8e, 0x41,
	0x0c, 0xc5, 0x17, 0xde, 0xa1, 0x0a, 0xa8, 0x8d, 0x1a, 0x05, 0x07, 0x96, 0x05, 0x29, 0xa8, 0x80,
	0x9a, 0xfe, 0xc2, 0x96, 0x73, 0x79, 0xc0, 0xf9, 0x05, 0x65, 0x91, 0x16, 0x5b, 0xd8, 0xf6, 0xb7,
	0xc8, 0x9c, 0x73, 0x7c, 0x88, 0x09, 0x17, 0xc0, 0xb0, 0x8b, 0xea, 0x29, 0x88, 0x84, 0x46, 0x7a,
	0x6b, 0x3b, 0x95, 0xad, 0x4d, 0x49, 0x9c, 0x0d, 0xb3, 0x67, 0x6a, 0xda, 0xd7, 0x6e, 0x4b, 0xe2,
	0xb5, 0xbf, 0x10, 0x5f, 0x7f, 0x43, 0xfc, 0x25, 0x6a, 0xc9, 0xe0, 0x7d, 0xe0, 0x39, 0xcd, 0x38,
	0xe0, 0xf7, 0x50, 0x2b, 0x08, 0x43, 0xe0, 0x5c, 0x67, 0xa6, 0xdc, 0xc8, 0x66, 0xf9, 0x4f, 0x65,
	0x07, 0x3f, 0x40, 0x6f, 0x31, 0x38, 0x63, 0xc0, 0x93, 0xa5, 0xec, 0xb5, 0xf4, 0xcf, 0x32, 0x85,
	0x63, 0x74, 0xbb, 0xf4, 0xfe, 0xbf, 0x80, 0x5f, 0xa2, 0xd6, 0x4b, 0x60, 0xe4, 0x8c, 0x84, 0xaa,
	0x6a, 0xf0, 0x0e, 0xda, 0x90, 0x1b, 0x3c, 0x21, 0x91, 0x46, 0xd6, 0xa5, 0x39, 0x8c, 0xb0, 0x83,
	0xea, 0x0c, 0x78, 0x31, 0x13, 0x0a, 0xb3, 0xd5, 0xdb, 0xae, 0xa4, 0x4e, 0x11, 0x2e, 0xa7, 0x22,
	0x10, 0xe0, 0x6b, 0x2f, 0xfb, 0x53, 0x64, 0xea, 0x8c, 0x06, 0x59, 0x10, 0x43, 0x0a, 0x99, 0xb8,
	0x19, 0x8e, 0xd1, 0x2d, 0x39, 0xd2, 0x11, 0xaa, 0x71, 0xe7, 0x21, 0x6a, 0x55, 0x8f, 0x04, 0x37,
	0xd1, 0x6d, 0xef, 0x59, 0x7f, 0x38, 0x32, 0xd7, 0xe4, 0x70, 0xf2, 0x64, 0xfc, 0xb9, 0x67, 0x1a,
	0x9d, 0x7d, 0xb4, 0x59, 0x51, 0xc7, 0x9b, 0x68, 0x63, 0xfa, 0xe2, 0xf0, 0xd0, 0x9b, 0x4e, 0xcd,
	0x35, 0x69, 0xf8, 0xde, 0x91, 0xef, 0x4d, 0x9f, 0x98, 0x86, 0x34, 0x86, 0xa3, 0x91, 0x77, 0xdc,
	0x1f, 0x99, 0xb5, 0xde, 0x6f, 0x75, 0x74, 0x47, 0x9e, 0x0f, 0x64, 0x42, 0x6f, 0x7c, 0x0a, 0xec,
	0x9c, 0x84, 0x80, 0xcf, 0x90, 0x29, 0x16, 0x35, 0x5e, 0xea, 0xe3, 0xbb, 0x37, 0x36, 0x40, 0xdb,
	0xac, 0x4e, 0xa9, 0x84, 0xbe, 0xff, 0xdd, 0xaf, 0xbf, 0xff, 0x54, 0xbb, 0x6f, 0xb7, 0xcb, 0xee,
	0x2d, 0x44, 0xe2, 0x2a, 0x60, 0x2e, 0xd7, 0x70, 0x12, 0x67, 0x24, 0x3b, 0x30, 0x3a, 0x78, 0x82,
	0xea, 0xbc, 0xa4, 0x57, 0x6b, 0xb0, 0x5a, 0xf8, 0x2b, 0xd8, 0x7b, 0x8a, 0x7d, 0xc7, 0x36, 0xaf,
	0xd8, 0x57, 0xc4, 0xaf, 0x4a, 0xe2, 0x8b, 0x1c, 0xef, 0xad, 0x20, 0xce, 0x5b, 0xa0, 0x5d, 0x95,
	0xab, 0x96, 0xe8, 0x4d, 0xf0, 0x22, 0x97, 0xf0, 0xa7, 0x68, 0x43, 0x1a, 0xe3, 0x42, 0xe0, 0x6b,
	0x61, 0xad, 0x08, 0xf4, 0x9e, 0x62, 0x6d, 0xdb, 0x6f, 0x2f, 0xb3, 0x68, 0x21, 0x34, 0x4c, 0xd7,
	0xe0, 0xbf, 0x85, 0xe9, 0xe5, 0x12, 0x76, 0x80, 0x36, 0xcf, 0x55, 0x01, 0x94, 0x45, 0x7f, 0x1d,
	0xb8, 0xf3, 0x66, 0xa1, 0xea, 0x13, 0xb7, 0xd7, 0x70, 0x82, 0xde, 0x61, 0xc0, 0x41, 0x4c, 0x74,
	0xdb, 0xfa, 0xf0, 0xba, 0x00, 0x2e, 0xfe, 0x06, 0xb2, 0x94, 0xb5, 0x15, 0xc7, 0x3d, 0xbf, 0x04,
	0x5c, 0x56, 0xe2, 0x64, 0x94, 0x27, 0xb2, 0x17, 0x2b, 0x4a, 0xff, 0x44, 0xe2, 0x81, 0x92, 0x78,
	0xd7, 0xde, 0x5d, 0x29, 0xc1, 0x41, 0x09, 0x0c, 0x50, 0x33, 0x88, 0x22, 0x5d, 0xb0, 0x7b, 0xd7,
	0xaf, 0xb5, 0x45, 0x13, 0xb6, 0xb7, 0x9d, 0xf2, 0x45, 0x70, 0xe6, 0x2f, 0x82, 0xe3, 0xc9, 0x17,
	0xc1, 0x5e, 0xc3, 0x1e, 0x6a, 0x45, 0x30, 0x03, 0x01, 0xff, 0x09, 0x33, 0xf8, 0xa5, 0xf6, 0x63,
	0xff, 0x0f, 0x03, 0x7f, 0x6f, 0x94, 0x77, 0xa0, 0xc5, 0xcb, 0xd6, 0xb2, 0x2f, 0xd1, 0xfd, 0x18,
	0x04, 0x11, 0xaf, 0x21, 0xb4, 0x82, 0xa5, 0xe6, 0x9b, 0xbb, 0xe0, 0x41, 0x22, 0x44, 0xce, 0x0f,
	0x5c, 0x37, 0x26, 0x22, 0x29, 0x4e, 0x9d, 0x90, 0xa6, 0x6e, 0x90, 0x72, 0xfa, 0x8a, 0xce, 0xdc,
	0x98, 0x76, 0x63, 0x96, 0x87, 0x5d, 0xe9, 0xd2, 0x65, 0xc0, 0x45, 0x37, 0x25, 0x21, 0xa3, 0x7a,
	0x69, 0x57, 0x14, 0x82, 0x32, 0x12, 0xcc, 0xda, 0x38, 0x85, 0x88, 0x14, 0xe9, 0x67, 0x7a, 0x9d,
	0x64, 0xf4, 0xd6, 0x3f, 0x74, 0x1e, 0x77, 0x8c, 0x5a, 0xcf, 0x0c, 0xf2, 0x7c, 0xa6, 0x45, 0xdd,
	0xaf, 0x39, 0xcd, 0x0e, 0xae, 0xfd, 0xf1, 0x3f, 0x46, 0xeb, 0xfb, 0x8f, 0xf7, 0xf1, 0x3e, 0xea,
	0xf8, 0x20, 0x0a, 0x96, 0x41, 0x64, 0x5d, 0x24, 0x90, 0x59, 0x22, 0x01, 0x8b, 0x01, 0xa7, 0x05,
	0x0b, 0xc1, 0x8a, 0x28, 0x70, 0x2b, 0xa3, 0xc2, 0x82, 0x6f, 0x08, 0x17, 0x0e, 0xae, 0xa3, 0x5b,
	0x3f, 0xd7, 0x8c, 0x0d, 0xb4, 0x1d, 0xd2, 0xd4, 0x99, 0xef, 0xf2, 0x2a, 0x7d, 0x83, 0xd5, 0xb7,
	0xcd, 0xc4, 0x38, 0xad, 0xab, 0xf9, 0x8f, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xd0, 0xc1,
	0x48, 0x0a, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	// reuqest response
	ThirdPartySignIn(ctx context.Context, in *ThirdPartyAuth, opts ...grpc.CallOption) (*Token, error)
	// reuqest response
	SignIn(ctx context.Context, in *GetItQECAuth, opts ...grpc.CallOption) (*Token, error)
	// reuqest response
	SignUp(ctx context.Context, in *GetItQECRegister, opts ...grpc.CallOption) (*AuthResponse, error)
	SignOut(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	Refresh(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	VerifyToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Verification, error)
	ResetPasswordRequest(ctx context.Context, in *Token, opts ...grpc.CallOption) (*AuthResponse, error)
	ResetPassword(ctx context.Context, in *Token, opts ...grpc.CallOption) (*AuthResponse, error)
	AddSignIn(ctx context.Context, in *SignInManagement, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteSignIn(ctx context.Context, in *SignInManagement, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) ThirdPartySignIn(ctx context.Context, in *ThirdPartyAuth, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/thirdPartySignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) SignIn(ctx context.Context, in *GetItQECAuth, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/signIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) SignUp(ctx context.Context, in *GetItQECRegister, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/signUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) SignOut(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/signOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) Refresh(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) VerifyToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Verification, error) {
	out := new(Verification)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/verifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ResetPasswordRequest(ctx context.Context, in *Token, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/resetPasswordRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ResetPassword(ctx context.Context, in *Token, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/resetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) AddSignIn(ctx context.Context, in *SignInManagement, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/addSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) DeleteSignIn(ctx context.Context, in *SignInManagement, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/authproto.AuthenticationService/deleteSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	// reuqest response
	ThirdPartySignIn(context.Context, *ThirdPartyAuth) (*Token, error)
	// reuqest response
	SignIn(context.Context, *GetItQECAuth) (*Token, error)
	// reuqest response
	SignUp(context.Context, *GetItQECRegister) (*AuthResponse, error)
	SignOut(context.Context, *Token) (*Token, error)
	Refresh(context.Context, *Token) (*Token, error)
	VerifyToken(context.Context, *Token) (*Verification, error)
	ResetPasswordRequest(context.Context, *Token) (*AuthResponse, error)
	ResetPassword(context.Context, *Token) (*AuthResponse, error)
	AddSignIn(context.Context, *SignInManagement) (*empty.Empty, error)
	DeleteSignIn(context.Context, *SignInManagement) (*empty.Empty, error)
}

// UnimplementedAuthenticationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (*UnimplementedAuthenticationServiceServer) ThirdPartySignIn(ctx context.Context, req *ThirdPartyAuth) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThirdPartySignIn not implemented")
}
func (*UnimplementedAuthenticationServiceServer) SignIn(ctx context.Context, req *GetItQECAuth) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedAuthenticationServiceServer) SignUp(ctx context.Context, req *GetItQECRegister) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAuthenticationServiceServer) SignOut(ctx context.Context, req *Token) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (*UnimplementedAuthenticationServiceServer) Refresh(ctx context.Context, req *Token) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (*UnimplementedAuthenticationServiceServer) VerifyToken(ctx context.Context, req *Token) (*Verification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (*UnimplementedAuthenticationServiceServer) ResetPasswordRequest(ctx context.Context, req *Token) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPasswordRequest not implemented")
}
func (*UnimplementedAuthenticationServiceServer) ResetPassword(ctx context.Context, req *Token) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (*UnimplementedAuthenticationServiceServer) AddSignIn(ctx context.Context, req *SignInManagement) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSignIn not implemented")
}
func (*UnimplementedAuthenticationServiceServer) DeleteSignIn(ctx context.Context, req *SignInManagement) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSignIn not implemented")
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_ThirdPartySignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThirdPartyAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ThirdPartySignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/ThirdPartySignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ThirdPartySignIn(ctx, req.(*ThirdPartyAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItQECAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).SignIn(ctx, req.(*GetItQECAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItQECRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).SignUp(ctx, req.(*GetItQECRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).SignOut(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Refresh(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).VerifyToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ResetPasswordRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ResetPasswordRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/ResetPasswordRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ResetPasswordRequest(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ResetPassword(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_AddSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInManagement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).AddSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/AddSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).AddSignIn(ctx, req.(*SignInManagement))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_DeleteSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInManagement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).DeleteSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authproto.AuthenticationService/DeleteSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).DeleteSignIn(ctx, req.(*SignInManagement))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authproto.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "thirdPartySignIn",
			Handler:    _AuthenticationService_ThirdPartySignIn_Handler,
		},
		{
			MethodName: "signIn",
			Handler:    _AuthenticationService_SignIn_Handler,
		},
		{
			MethodName: "signUp",
			Handler:    _AuthenticationService_SignUp_Handler,
		},
		{
			MethodName: "signOut",
			Handler:    _AuthenticationService_SignOut_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _AuthenticationService_Refresh_Handler,
		},
		{
			MethodName: "verifyToken",
			Handler:    _AuthenticationService_VerifyToken_Handler,
		},
		{
			MethodName: "resetPasswordRequest",
			Handler:    _AuthenticationService_ResetPasswordRequest_Handler,
		},
		{
			MethodName: "resetPassword",
			Handler:    _AuthenticationService_ResetPassword_Handler,
		},
		{
			MethodName: "addSignIn",
			Handler:    _AuthenticationService_AddSignIn_Handler,
		},
		{
			MethodName: "deleteSignIn",
			Handler:    _AuthenticationService_DeleteSignIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
