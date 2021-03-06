// Code generated by protoc-gen-go. DO NOT EDIT.
// source: redis/sumpb/sumpb.proto

package sumpb // import "redis/sumpb"

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

type Users struct {
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	MSSV                 string   `protobuf:"bytes,4,opt,name=MSSV,proto3" json:"MSSV,omitempty"`
	SubjectID            int32    `protobuf:"varint,5,opt,name=SubjectID,proto3" json:"SubjectID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_sumpb_3ed1d666869c20e7, []int{0}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Users) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Users) GetMSSV() string {
	if m != nil {
		return m.MSSV
	}
	return ""
}

func (m *Users) GetSubjectID() int32 {
	if m != nil {
		return m.SubjectID
	}
	return 0
}

type UsersRequest struct {
	User                 *Users   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersRequest) Reset()         { *m = UsersRequest{} }
func (m *UsersRequest) String() string { return proto.CompactTextString(m) }
func (*UsersRequest) ProtoMessage()    {}
func (*UsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sumpb_3ed1d666869c20e7, []int{1}
}
func (m *UsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersRequest.Unmarshal(m, b)
}
func (m *UsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersRequest.Marshal(b, m, deterministic)
}
func (dst *UsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersRequest.Merge(dst, src)
}
func (m *UsersRequest) XXX_Size() int {
	return xxx_messageInfo_UsersRequest.Size(m)
}
func (m *UsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsersRequest proto.InternalMessageInfo

func (m *UsersRequest) GetUser() *Users {
	if m != nil {
		return m.User
	}
	return nil
}

type UsersResponse struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	MSSV                 string   `protobuf:"bytes,4,opt,name=MSSV,proto3" json:"MSSV,omitempty"`
	SubjectID            int32    `protobuf:"varint,5,opt,name=SubjectID,proto3" json:"SubjectID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersResponse) Reset()         { *m = UsersResponse{} }
func (m *UsersResponse) String() string { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()    {}
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sumpb_3ed1d666869c20e7, []int{2}
}
func (m *UsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersResponse.Unmarshal(m, b)
}
func (m *UsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersResponse.Marshal(b, m, deterministic)
}
func (dst *UsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersResponse.Merge(dst, src)
}
func (m *UsersResponse) XXX_Size() int {
	return xxx_messageInfo_UsersResponse.Size(m)
}
func (m *UsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersResponse proto.InternalMessageInfo

func (m *UsersResponse) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *UsersResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UsersResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UsersResponse) GetMSSV() string {
	if m != nil {
		return m.MSSV
	}
	return ""
}

func (m *UsersResponse) GetSubjectID() int32 {
	if m != nil {
		return m.SubjectID
	}
	return 0
}

// get user by id
type GetUserById struct {
	UserID               int32    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserById) Reset()         { *m = GetUserById{} }
func (m *GetUserById) String() string { return proto.CompactTextString(m) }
func (*GetUserById) ProtoMessage()    {}
func (*GetUserById) Descriptor() ([]byte, []int) {
	return fileDescriptor_sumpb_3ed1d666869c20e7, []int{3}
}
func (m *GetUserById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserById.Unmarshal(m, b)
}
func (m *GetUserById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserById.Marshal(b, m, deterministic)
}
func (dst *GetUserById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserById.Merge(dst, src)
}
func (m *GetUserById) XXX_Size() int {
	return xxx_messageInfo_GetUserById.Size(m)
}
func (m *GetUserById) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserById.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserById proto.InternalMessageInfo

func (m *GetUserById) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type GetUserByIDRequest struct {
	UserID               *GetUserById `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetUserByIDRequest) Reset()         { *m = GetUserByIDRequest{} }
func (m *GetUserByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByIDRequest) ProtoMessage()    {}
func (*GetUserByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sumpb_3ed1d666869c20e7, []int{4}
}
func (m *GetUserByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByIDRequest.Unmarshal(m, b)
}
func (m *GetUserByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByIDRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByIDRequest.Merge(dst, src)
}
func (m *GetUserByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByIDRequest.Size(m)
}
func (m *GetUserByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByIDRequest proto.InternalMessageInfo

func (m *GetUserByIDRequest) GetUserID() *GetUserById {
	if m != nil {
		return m.UserID
	}
	return nil
}

func init() {
	proto.RegisterType((*Users)(nil), "sumpb.Users")
	proto.RegisterType((*UsersRequest)(nil), "sumpb.UsersRequest")
	proto.RegisterType((*UsersResponse)(nil), "sumpb.UsersResponse")
	proto.RegisterType((*GetUserById)(nil), "sumpb.GetUserById")
	proto.RegisterType((*GetUserByIDRequest)(nil), "sumpb.GetUserByIDRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CreateServiceClient is the client API for CreateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateServiceClient interface {
	Create(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*UsersResponse, error)
}

type createServiceClient struct {
	cc *grpc.ClientConn
}

func NewCreateServiceClient(cc *grpc.ClientConn) CreateServiceClient {
	return &createServiceClient{cc}
}

func (c *createServiceClient) Create(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/sumpb.CreateService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createServiceClient) GetUserByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/sumpb.CreateService/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateServiceServer is the server API for CreateService service.
type CreateServiceServer interface {
	Create(context.Context, *UsersRequest) (*UsersResponse, error)
	GetUserByID(context.Context, *GetUserByIDRequest) (*UsersResponse, error)
}

func RegisterCreateServiceServer(s *grpc.Server, srv CreateServiceServer) {
	s.RegisterService(&_CreateService_serviceDesc, srv)
}

func _CreateService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sumpb.CreateService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateServiceServer).Create(ctx, req.(*UsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sumpb.CreateService/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateServiceServer).GetUserByID(ctx, req.(*GetUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sumpb.CreateService",
	HandlerType: (*CreateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CreateService_Create_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _CreateService_GetUserByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redis/sumpb/sumpb.proto",
}

func init() { proto.RegisterFile("redis/sumpb/sumpb.proto", fileDescriptor_sumpb_3ed1d666869c20e7) }

var fileDescriptor_sumpb_3ed1d666869c20e7 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x75, 0x6b, 0x53, 0xec, 0xa4, 0xf1, 0x30, 0x8a, 0xc6, 0xe2, 0x21, 0x04, 0x84, 0xe0, 0xa1,
	0x4a, 0xc4, 0x7b, 0xa9, 0x8b, 0x12, 0x50, 0x0f, 0x09, 0x7a, 0xf0, 0x96, 0xb4, 0x73, 0x88, 0x50,
	0x13, 0x77, 0x37, 0x82, 0x7f, 0x20, 0xf8, 0xd3, 0x92, 0xdd, 0xb4, 0x49, 0xf0, 0xe0, 0xcd, 0xcb,
	0xb2, 0xf3, 0xe6, 0xcd, 0x9b, 0x37, 0xc3, 0xc0, 0xb1, 0xa0, 0x55, 0x2e, 0x2f, 0x64, 0xb5, 0x2e,
	0x33, 0xf3, 0xce, 0x4a, 0x51, 0xa8, 0x02, 0x2d, 0x1d, 0xf8, 0x12, 0xac, 0x27, 0x49, 0x42, 0xe2,
	0x29, 0x8c, 0x6f, 0x73, 0x21, 0xd5, 0x63, 0xba, 0x26, 0x77, 0xe0, 0xb1, 0x60, 0x1c, 0xb7, 0x00,
	0x4e, 0x61, 0xef, 0x3e, 0x6d, 0x92, 0xbb, 0x3a, 0xb9, 0x8d, 0x11, 0x61, 0xf8, 0x90, 0x24, 0xcf,
	0xee, 0x50, 0xe3, 0xfa, 0x5f, 0xab, 0x25, 0x55, 0xf6, 0x4a, 0x4b, 0x15, 0x71, 0xd7, 0xf2, 0x58,
	0x60, 0xc5, 0x2d, 0xe0, 0x5f, 0xc2, 0x44, 0x37, 0x8d, 0xe9, 0xbd, 0x22, 0xa9, 0xd0, 0x83, 0x61,
	0x25, 0x49, 0xb8, 0xcc, 0x63, 0x81, 0x1d, 0x4e, 0x66, 0xc6, 0xa7, 0xa1, 0xe8, 0x8c, 0xff, 0xcd,
	0xc0, 0x69, 0x4a, 0x64, 0x59, 0xbc, 0x49, 0xc2, 0x7d, 0x18, 0x44, 0x5c, 0x57, 0x58, 0xf1, 0x20,
	0xe2, 0xff, 0xea, 0xff, 0x0c, 0xec, 0x3b, 0x52, 0xb5, 0x9f, 0xc5, 0x67, 0xb4, 0xc2, 0x23, 0x18,
	0xd5, 0x26, 0xb7, 0x76, 0x9a, 0xc8, 0x9f, 0x03, 0xb6, 0x34, 0xbe, 0x19, 0xf6, 0xbc, 0xc7, 0xb6,
	0x43, 0x6c, 0xc6, 0xed, 0x28, 0x6e, 0x14, 0xc2, 0x2f, 0x06, 0xce, 0x8d, 0xa0, 0x54, 0x51, 0x42,
	0xe2, 0x23, 0x5f, 0x12, 0x5e, 0xc3, 0xc8, 0x00, 0x78, 0xd0, 0x5b, 0x93, 0x11, 0x9f, 0x1e, 0xf6,
	0x41, 0xb3, 0x2b, 0x7f, 0x07, 0xe7, 0x5d, 0xc7, 0x1c, 0x4f, 0x7e, 0xf5, 0xe4, 0x7f, 0x28, 0x2c,
	0x9c, 0x17, 0xbb, 0x73, 0x4a, 0xd9, 0x48, 0x5f, 0xd1, 0xd5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x32, 0x1e, 0x2a, 0xef, 0x60, 0x02, 0x00, 0x00,
}
