// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bicasfiattoken/bicasfiattoken/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("bicasfiattoken/bicasfiattoken/tx.proto", fileDescriptor_bcc4a1d06cadfd78)
}

var fileDescriptor_bcc4a1d06cadfd78 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0xca, 0x4c, 0x4e,
	0x2c, 0x4e, 0xcb, 0x4c, 0x2c, 0x29, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x47, 0xe3, 0x96, 0x54, 0xe8,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x69, 0x55, 0x96, 0x16, 0x97, 0x16, 0xe4, 0x97, 0x25, 0xe6,
	0x95, 0xe4, 0xe7, 0xe9, 0xa1, 0xaa, 0x42, 0xe3, 0x1a, 0xb1, 0x72, 0x31, 0xfb, 0x16, 0xa7, 0x3b,
	0x85, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6d, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xb2, 0xb9, 0xfa, 0x4e, 0x9e, 0xce, 0x8e, 0xc1,
	0x6e, 0x99, 0x89, 0x25, 0x21, 0x60, 0xdb, 0x2b, 0x30, 0x9c, 0x53, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0x76, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x07, 0xfd, 0x02, 0xbc, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yusupovanton.bicasfiattoken.bicasfiattoken.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "bicasfiattoken/bicasfiattoken/tx.proto",
}
