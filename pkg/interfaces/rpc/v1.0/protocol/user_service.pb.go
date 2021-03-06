// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_service.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RegisterUserRequestType struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserRequestType) Reset()         { *m = RegisterUserRequestType{} }
func (m *RegisterUserRequestType) String() string { return proto.CompactTextString(m) }
func (*RegisterUserRequestType) ProtoMessage()    {}
func (*RegisterUserRequestType) Descriptor() ([]byte, []int) {
	return fileDescriptor_292f630cd9eb4c90, []int{0}
}

func (m *RegisterUserRequestType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserRequestType.Unmarshal(m, b)
}
func (m *RegisterUserRequestType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserRequestType.Marshal(b, m, deterministic)
}
func (m *RegisterUserRequestType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserRequestType.Merge(m, src)
}
func (m *RegisterUserRequestType) XXX_Size() int {
	return xxx_messageInfo_RegisterUserRequestType.Size(m)
}
func (m *RegisterUserRequestType) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserRequestType.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserRequestType proto.InternalMessageInfo

func (m *RegisterUserRequestType) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RegisterUserResponseType struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserResponseType) Reset()         { *m = RegisterUserResponseType{} }
func (m *RegisterUserResponseType) String() string { return proto.CompactTextString(m) }
func (*RegisterUserResponseType) ProtoMessage()    {}
func (*RegisterUserResponseType) Descriptor() ([]byte, []int) {
	return fileDescriptor_292f630cd9eb4c90, []int{1}
}

func (m *RegisterUserResponseType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserResponseType.Unmarshal(m, b)
}
func (m *RegisterUserResponseType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserResponseType.Marshal(b, m, deterministic)
}
func (m *RegisterUserResponseType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserResponseType.Merge(m, src)
}
func (m *RegisterUserResponseType) XXX_Size() int {
	return xxx_messageInfo_RegisterUserResponseType.Size(m)
}
func (m *RegisterUserResponseType) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserResponseType.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserResponseType proto.InternalMessageInfo

func (m *RegisterUserResponseType) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterUserRequestType)(nil), "protocol.RegisterUserRequestType")
	proto.RegisterType((*RegisterUserResponseType)(nil), "protocol.RegisterUserResponseType")
}

func init() { proto.RegisterFile("user_service.proto", fileDescriptor_292f630cd9eb4c90) }

var fileDescriptor_292f630cd9eb4c90 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2d, 0x4e, 0x2d,
	0x8a, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0xfa, 0x5c, 0xe2, 0x41, 0xa9, 0xe9, 0x99, 0xc5, 0x25, 0xa9,
	0x45, 0xa1, 0xc5, 0xa9, 0x45, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x21, 0x95, 0x05, 0xa9,
	0x42, 0x22, 0x5c, 0xac, 0xa9, 0xb9, 0x89, 0x99, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x10, 0x8e, 0x92, 0x09, 0x97, 0x04, 0xaa, 0x86, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xb0, 0x0e,
	0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0xa8, 0x1e, 0x18, 0xd7, 0x28, 0x85,
	0x8b, 0x1b, 0xa4, 0x3a, 0x18, 0xe2, 0x0a, 0xa1, 0x50, 0x2e, 0x1e, 0x64, 0x43, 0x84, 0x14, 0xf5,
	0x60, 0x0e, 0xd2, 0xc3, 0xe1, 0x1a, 0x29, 0x25, 0x5c, 0x4a, 0x10, 0xf6, 0x27, 0xb1, 0x81, 0x95,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x56, 0x68, 0x9d, 0xf3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequestType, opts ...grpc.CallOption) (*RegisterUserResponseType, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequestType, opts ...grpc.CallOption) (*RegisterUserResponseType, error) {
	out := new(RegisterUserResponseType)
	err := c.cc.Invoke(ctx, "/protocol.UserService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequestType) (*RegisterUserResponseType, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) RegisterUser(ctx context.Context, req *RegisterUserRequestType) (*RegisterUserResponseType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequestType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.UserService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*RegisterUserRequestType))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
