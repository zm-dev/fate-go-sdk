// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package ZMDev_Fate_pb is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	UserID
	TicketID
	LoginCheckRes
	Unused
	Credential
	AccessToken
	Certificate
*/
package ZMDev_Fate_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CertificateType int32

const (
	CertificateType_StudentNum CertificateType = 0
)

var CertificateType_name = map[int32]string{
	0: "StudentNum",
}
var CertificateType_value = map[string]int32{
	"StudentNum": 0,
}

func (x CertificateType) String() string {
	return proto.EnumName(CertificateType_name, int32(x))
}
func (CertificateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UserID struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *UserID) Reset()                    { *m = UserID{} }
func (m *UserID) String() string            { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()               {}
func (*UserID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TicketID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *TicketID) Reset()                    { *m = TicketID{} }
func (m *TicketID) String() string            { return proto.CompactTextString(m) }
func (*TicketID) ProtoMessage()               {}
func (*TicketID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TicketID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type LoginCheckRes struct {
	IsLogin bool  `protobuf:"varint,1,opt,name=is_login,json=isLogin" json:"is_login,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *LoginCheckRes) Reset()                    { *m = LoginCheckRes{} }
func (m *LoginCheckRes) String() string            { return proto.CompactTextString(m) }
func (*LoginCheckRes) ProtoMessage()               {}
func (*LoginCheckRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginCheckRes) GetIsLogin() bool {
	if m != nil {
		return m.IsLogin
	}
	return false
}

func (m *LoginCheckRes) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Unused struct {
}

func (m *Unused) Reset()                    { *m = Unused{} }
func (m *Unused) String() string            { return proto.CompactTextString(m) }
func (*Unused) ProtoMessage()               {}
func (*Unused) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Credential struct {
	AppId     int32  `protobuf:"varint,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	AppSecret string `protobuf:"bytes,2,opt,name=app_secret,json=appSecret" json:"app_secret,omitempty"`
}

func (m *Credential) Reset()                    { *m = Credential{} }
func (m *Credential) String() string            { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()               {}
func (*Credential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Credential) GetAppId() int32 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *Credential) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

type AccessToken struct {
	Token     string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	ExpiredAt int64  `protobuf:"varint,2,opt,name=expired_at,json=expiredAt" json:"expired_at,omitempty"`
}

func (m *AccessToken) Reset()                    { *m = AccessToken{} }
func (m *AccessToken) String() string            { return proto.CompactTextString(m) }
func (*AccessToken) ProtoMessage()               {}
func (*AccessToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AccessToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccessToken) GetExpiredAt() int64 {
	if m != nil {
		return m.ExpiredAt
	}
	return 0
}

type Certificate struct {
	Id              int64           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId          int64           `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Account         string          `protobuf:"bytes,3,opt,name=account" json:"account,omitempty"`
	CertificateType CertificateType `protobuf:"varint,4,opt,name=certificate_type,json=certificateType,enum=ZMDev.Fate.pb.CertificateType" json:"certificate_type,omitempty"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Certificate) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Certificate) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Certificate) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Certificate) GetCertificateType() CertificateType {
	if m != nil {
		return m.CertificateType
	}
	return CertificateType_StudentNum
}

func init() {
	proto.RegisterType((*UserID)(nil), "ZMDev.Fate.pb.UserID")
	proto.RegisterType((*TicketID)(nil), "ZMDev.Fate.pb.TicketID")
	proto.RegisterType((*LoginCheckRes)(nil), "ZMDev.Fate.pb.LoginCheckRes")
	proto.RegisterType((*Unused)(nil), "ZMDev.Fate.pb.Unused")
	proto.RegisterType((*Credential)(nil), "ZMDev.Fate.pb.Credential")
	proto.RegisterType((*AccessToken)(nil), "ZMDev.Fate.pb.AccessToken")
	proto.RegisterType((*Certificate)(nil), "ZMDev.Fate.pb.Certificate")
	proto.RegisterEnum("ZMDev.Fate.pb.CertificateType", CertificateType_name, CertificateType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccessTokenService service

type AccessTokenServiceClient interface {
	// 获取 access token
	Token(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*AccessToken, error)
}

type accessTokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessTokenServiceClient(cc *grpc.ClientConn) AccessTokenServiceClient {
	return &accessTokenServiceClient{cc}
}

func (c *accessTokenServiceClient) Token(ctx context.Context, in *Credential, opts ...grpc.CallOption) (*AccessToken, error) {
	out := new(AccessToken)
	err := grpc.Invoke(ctx, "/ZMDev.Fate.pb.AccessTokenService/Token", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccessTokenService service

type AccessTokenServiceServer interface {
	// 获取 access token
	Token(context.Context, *Credential) (*AccessToken, error)
}

func RegisterAccessTokenServiceServer(s *grpc.Server, srv AccessTokenServiceServer) {
	s.RegisterService(&_AccessTokenService_serviceDesc, srv)
}

func _AccessTokenService_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessTokenServiceServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ZMDev.Fate.pb.AccessTokenService/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessTokenServiceServer).Token(ctx, req.(*Credential))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessTokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ZMDev.Fate.pb.AccessTokenService",
	HandlerType: (*AccessTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Token",
			Handler:    _AccessTokenService_Token_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for CertificateService service

type CertificateServiceClient interface {
	Update(ctx context.Context, in *Certificate, opts ...grpc.CallOption) (*Unused, error)
}

type certificateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCertificateServiceClient(cc *grpc.ClientConn) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) Update(ctx context.Context, in *Certificate, opts ...grpc.CallOption) (*Unused, error) {
	out := new(Unused)
	err := grpc.Invoke(ctx, "/ZMDev.Fate.pb.CertificateService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateService service

type CertificateServiceServer interface {
	Update(context.Context, *Certificate) (*Unused, error)
}

func RegisterCertificateServiceServer(s *grpc.Server, srv CertificateServiceServer) {
	s.RegisterService(&_CertificateService_serviceDesc, srv)
}

func _CertificateService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Certificate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ZMDev.Fate.pb.CertificateService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).Update(ctx, req.(*Certificate))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ZMDev.Fate.pb.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _CertificateService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for LoginChecker service

type LoginCheckerClient interface {
	Check(ctx context.Context, in *TicketID, opts ...grpc.CallOption) (*LoginCheckRes, error)
	Logout(ctx context.Context, in *TicketID, opts ...grpc.CallOption) (*Unused, error)
}

type loginCheckerClient struct {
	cc *grpc.ClientConn
}

func NewLoginCheckerClient(cc *grpc.ClientConn) LoginCheckerClient {
	return &loginCheckerClient{cc}
}

func (c *loginCheckerClient) Check(ctx context.Context, in *TicketID, opts ...grpc.CallOption) (*LoginCheckRes, error) {
	out := new(LoginCheckRes)
	err := grpc.Invoke(ctx, "/ZMDev.Fate.pb.LoginChecker/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginCheckerClient) Logout(ctx context.Context, in *TicketID, opts ...grpc.CallOption) (*Unused, error) {
	out := new(Unused)
	err := grpc.Invoke(ctx, "/ZMDev.Fate.pb.LoginChecker/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoginChecker service

type LoginCheckerServer interface {
	Check(context.Context, *TicketID) (*LoginCheckRes, error)
	Logout(context.Context, *TicketID) (*Unused, error)
}

func RegisterLoginCheckerServer(s *grpc.Server, srv LoginCheckerServer) {
	s.RegisterService(&_LoginChecker_serviceDesc, srv)
}

func _LoginChecker_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginCheckerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ZMDev.Fate.pb.LoginChecker/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginCheckerServer).Check(ctx, req.(*TicketID))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginChecker_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginCheckerServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ZMDev.Fate.pb.LoginChecker/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginCheckerServer).Logout(ctx, req.(*TicketID))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ZMDev.Fate.pb.LoginChecker",
	HandlerType: (*LoginCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _LoginChecker_Check_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _LoginChecker_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for UserService service

type UserServiceClient interface {
	// todo update password
	UpdatePassword(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Unused, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Unused, error) {
	out := new(Unused)
	err := grpc.Invoke(ctx, "/ZMDev.Fate.pb.UserService/UpdatePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	// todo update password
	UpdatePassword(context.Context, *UserID) (*Unused, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ZMDev.Fate.pb.UserService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ZMDev.Fate.pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xcd, 0xa6, 0x64, 0x93, 0x9d, 0xd0, 0xb4, 0xb2, 0xa8, 0xba, 0x5d, 0x01, 0x2a, 0x7b, 0xaa,
	0x38, 0xec, 0x21, 0xdc, 0x90, 0x90, 0x68, 0x53, 0x21, 0x45, 0x2a, 0x1f, 0xda, 0x24, 0x1c, 0xb8,
	0xac, 0x5c, 0x7b, 0x28, 0x56, 0xca, 0xda, 0xb2, 0xbd, 0x85, 0xfc, 0x07, 0xfe, 0x01, 0x7f, 0x16,
	0xd9, 0xd9, 0x90, 0x64, 0xab, 0xbd, 0xd9, 0xf3, 0xc6, 0xef, 0xcd, 0x3c, 0xcf, 0x40, 0x44, 0x95,
	0xc8, 0x94, 0x96, 0x56, 0x92, 0xc3, 0x6f, 0x1f, 0xaf, 0xf1, 0x21, 0xfb, 0x40, 0x2d, 0x66, 0xea,
	0x36, 0x8d, 0x21, 0x5c, 0x18, 0xd4, 0xd3, 0x6b, 0x32, 0x82, 0xae, 0xe0, 0x71, 0x70, 0x1e, 0x5c,
	0x1c, 0xe4, 0x5d, 0xc1, 0xd3, 0x04, 0x06, 0x73, 0xc1, 0x96, 0x68, 0xf7, 0xb0, 0xc8, 0x63, 0x13,
	0x38, 0xbc, 0x91, 0x77, 0xa2, 0x9c, 0xfc, 0x40, 0xb6, 0xcc, 0xd1, 0x90, 0x33, 0x18, 0x08, 0x53,
	0xdc, 0xbb, 0x98, 0x4f, 0x1b, 0xe4, 0x7d, 0x61, 0x7c, 0x0a, 0x39, 0x85, 0x7e, 0x65, 0x50, 0x17,
	0x82, 0xc7, 0x5d, 0x4f, 0x1e, 0xba, 0xeb, 0x94, 0xa7, 0x03, 0x08, 0x17, 0x65, 0x65, 0x90, 0xa7,
	0x57, 0x00, 0x13, 0x8d, 0x1c, 0x4b, 0x2b, 0xe8, 0x3d, 0x39, 0x81, 0x90, 0x2a, 0x55, 0xd4, 0x82,
	0xbd, 0xbc, 0x47, 0x95, 0x9a, 0x72, 0xf2, 0x02, 0xc0, 0x85, 0x0d, 0x32, 0x8d, 0xd6, 0x53, 0x45,
	0x79, 0x44, 0x95, 0x9a, 0xf9, 0x40, 0x7a, 0x05, 0xc3, 0x4b, 0xc6, 0xd0, 0x98, 0xb9, 0x5c, 0x62,
	0x49, 0x9e, 0x41, 0xcf, 0xba, 0x43, 0x5d, 0xf4, 0xfa, 0xe2, 0x38, 0xf0, 0xb7, 0x12, 0x1a, 0x79,
	0x41, 0x6d, 0x5d, 0x4e, 0x54, 0x47, 0x2e, 0x6d, 0xfa, 0x37, 0x80, 0xe1, 0x04, 0xb5, 0x15, 0xdf,
	0x05, 0xa3, 0x16, 0x9b, 0x96, 0xb4, 0xb6, 0x42, 0x62, 0xe8, 0x53, 0xc6, 0x64, 0x55, 0xda, 0xf8,
	0xc0, 0xeb, 0x6d, 0xae, 0x64, 0x0a, 0xc7, 0x6c, 0xcb, 0x58, 0xd8, 0x95, 0xc2, 0xf8, 0xc9, 0x79,
	0x70, 0x31, 0x1a, 0xbf, 0xcc, 0xf6, 0x7e, 0x22, 0xdb, 0x11, 0x9e, 0xaf, 0x14, 0xe6, 0x47, 0x6c,
	0x3f, 0xf0, 0xfa, 0x15, 0x1c, 0x35, 0x72, 0xc8, 0x08, 0x60, 0x66, 0x2b, 0x67, 0xdc, 0xa7, 0xea,
	0xe7, 0x71, 0x67, 0xfc, 0x15, 0xc8, 0x8e, 0x09, 0x33, 0xd4, 0x0f, 0x82, 0x21, 0x79, 0x0f, 0xbd,
	0xb5, 0x29, 0x67, 0x4d, 0xc9, 0xff, 0xa6, 0x27, 0x49, 0x03, 0xda, 0xa1, 0x49, 0x3b, 0xe3, 0x19,
	0x90, 0x1d, 0xe9, 0x0d, 0xef, 0x3b, 0x08, 0x17, 0x8a, 0x3b, 0xa3, 0x92, 0xf6, 0x5e, 0x92, 0x93,
	0x06, 0x56, 0xff, 0x79, 0x67, 0xfc, 0x27, 0x80, 0xa7, 0xdb, 0x29, 0x42, 0xed, 0xea, 0xf4, 0x47,
	0x72, 0xda, 0x78, 0xb2, 0x99, 0xc3, 0xe4, 0x79, 0x03, 0xd8, 0x1b, 0xc2, 0xb4, 0x43, 0xde, 0x42,
	0x78, 0x23, 0xef, 0x64, 0x65, 0xdb, 0x29, 0x5a, 0xcb, 0xf9, 0x0c, 0x43, 0xb7, 0x09, 0x5b, 0xd3,
	0x46, 0xeb, 0xe6, 0xbe, 0x50, 0x63, 0x7e, 0x49, 0xcd, 0xc9, 0xa3, 0x97, 0x7e, 0x6f, 0x5a, 0x09,
	0x6f, 0x43, 0xbf, 0x70, 0x6f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x65, 0x64, 0x6f, 0x7d,
	0x03, 0x00, 0x00,
}
