// Code generated by protoc-gen-go. DO NOT EDIT.
// source: watchtowerrpc/watchtower.proto

package watchtowerrpc // import "github.com/lightningnetwork/lnd/lnrpc/watchtowerrpc"

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

type GetInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_watchtower_c1164e485ef0665e, []int{0}
}
func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(dst, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

type GetInfoResponse struct {
	// / The public key of the watchtower.
	Pubkey []byte `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	// / The listening addresses of the watchtower.
	Listeners []string `protobuf:"bytes,2,rep,name=listeners,proto3" json:"listeners,omitempty"`
	// / The URIs of the watchtower.
	Uris                 []string `protobuf:"bytes,3,rep,name=uris,proto3" json:"uris,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoResponse) Reset()         { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()    {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_watchtower_c1164e485ef0665e, []int{1}
}
func (m *GetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResponse.Unmarshal(m, b)
}
func (m *GetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResponse.Marshal(b, m, deterministic)
}
func (dst *GetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResponse.Merge(dst, src)
}
func (m *GetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfoResponse.Size(m)
}
func (m *GetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResponse proto.InternalMessageInfo

func (m *GetInfoResponse) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *GetInfoResponse) GetListeners() []string {
	if m != nil {
		return m.Listeners
	}
	return nil
}

func (m *GetInfoResponse) GetUris() []string {
	if m != nil {
		return m.Uris
	}
	return nil
}

func init() {
	proto.RegisterType((*GetInfoRequest)(nil), "watchtowerrpc.GetInfoRequest")
	proto.RegisterType((*GetInfoResponse)(nil), "watchtowerrpc.GetInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WatchtowerClient is the client API for Watchtower service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WatchtowerClient interface {
	// * lncli: tower info
	// GetInfo returns general information concerning the companion watchtower
	// including it's public key and URIs where the server is currently
	// listening for clients.
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type watchtowerClient struct {
	cc *grpc.ClientConn
}

func NewWatchtowerClient(cc *grpc.ClientConn) WatchtowerClient {
	return &watchtowerClient{cc}
}

func (c *watchtowerClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/watchtowerrpc.Watchtower/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatchtowerServer is the server API for Watchtower service.
type WatchtowerServer interface {
	// * lncli: tower info
	// GetInfo returns general information concerning the companion watchtower
	// including it's public key and URIs where the server is currently
	// listening for clients.
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
}

func RegisterWatchtowerServer(s *grpc.Server, srv WatchtowerServer) {
	s.RegisterService(&_Watchtower_serviceDesc, srv)
}

func _Watchtower_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watchtowerrpc.Watchtower/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Watchtower_serviceDesc = grpc.ServiceDesc{
	ServiceName: "watchtowerrpc.Watchtower",
	HandlerType: (*WatchtowerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Watchtower_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "watchtowerrpc/watchtower.proto",
}

func init() {
	proto.RegisterFile("watchtowerrpc/watchtower.proto", fileDescriptor_watchtower_c1164e485ef0665e)
}

var fileDescriptor_watchtower_c1164e485ef0665e = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4f, 0x2c, 0x49,
	0xce, 0x28, 0xc9, 0x2f, 0x4f, 0x2d, 0x2a, 0x2a, 0x48, 0xd6, 0x47, 0xf0, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x78, 0x51, 0xe4, 0x95, 0x04, 0xb8, 0xf8, 0xdc, 0x53, 0x4b, 0x3c, 0xf3, 0xd2,
	0xf2, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xa2, 0xb9, 0xf8, 0xe1, 0x22, 0xc5, 0x05,
	0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0x05, 0xa5, 0x49, 0xd9, 0xa9, 0x95, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x50, 0x9e, 0x90, 0x0c, 0x17, 0x67, 0x4e, 0x66, 0x71, 0x49, 0x6a,
	0x5e, 0x6a, 0x51, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x42, 0x40, 0x48, 0x88, 0x8b,
	0xa5, 0xb4, 0x28, 0xb3, 0x58, 0x82, 0x19, 0x2c, 0x01, 0x66, 0x1b, 0x85, 0x71, 0x71, 0x85, 0xc3,
	0xed, 0x17, 0xf2, 0xe0, 0x62, 0x87, 0x5a, 0x25, 0x24, 0xab, 0x87, 0xe2, 0x2e, 0x3d, 0x54, 0x47,
	0x49, 0xc9, 0xe1, 0x92, 0x86, 0xb8, 0xd0, 0xc9, 0x34, 0xca, 0x38, 0x3d, 0xb3, 0x24, 0xa3, 0x34,
	0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x27, 0x33, 0x3d, 0xa3, 0x24, 0x2f, 0x33, 0x2f, 0x3d, 0x2f,
	0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x5b, 0x3f, 0x27, 0x2f, 0x45, 0x3f, 0x27, 0x0f, 0x35, 0x3c, 0x8a,
	0x0a, 0x92, 0x93, 0xd8, 0xc0, 0x61, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x23, 0x0b,
	0x68, 0x35, 0x01, 0x00, 0x00,
}
