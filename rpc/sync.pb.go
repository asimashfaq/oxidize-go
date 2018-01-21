// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/sync.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import entities "github.com/tclchiam/block_n_go/encoding"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetBestHeaderRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetBestHeaderRequest) Reset()                    { *m = GetBestHeaderRequest{} }
func (m *GetBestHeaderRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBestHeaderRequest) ProtoMessage()               {}
func (*GetBestHeaderRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type GetBestHeaderResponse struct {
	Header           *entities.BlockHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *GetBestHeaderResponse) Reset()                    { *m = GetBestHeaderResponse{} }
func (m *GetBestHeaderResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBestHeaderResponse) ProtoMessage()               {}
func (*GetBestHeaderResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetBestHeaderResponse) GetHeader() *entities.BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type GetHeadersRequest struct {
	LatestHash       []byte  `protobuf:"bytes,1,req,name=latestHash" json:"latestHash,omitempty"`
	LatestIndex      *uint64 `protobuf:"varint,2,req,name=latestIndex" json:"latestIndex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetHeadersRequest) Reset()                    { *m = GetHeadersRequest{} }
func (m *GetHeadersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetHeadersRequest) ProtoMessage()               {}
func (*GetHeadersRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetHeadersRequest) GetLatestHash() []byte {
	if m != nil {
		return m.LatestHash
	}
	return nil
}

func (m *GetHeadersRequest) GetLatestIndex() uint64 {
	if m != nil && m.LatestIndex != nil {
		return *m.LatestIndex
	}
	return 0
}

type GetHeadersResponse struct {
	HeaderCount      *uint32                 `protobuf:"varint,1,req,name=headerCount" json:"headerCount,omitempty"`
	Headers          []*entities.BlockHeader `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *GetHeadersResponse) Reset()                    { *m = GetHeadersResponse{} }
func (m *GetHeadersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetHeadersResponse) ProtoMessage()               {}
func (*GetHeadersResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetHeadersResponse) GetHeaderCount() uint32 {
	if m != nil && m.HeaderCount != nil {
		return *m.HeaderCount
	}
	return 0
}

func (m *GetHeadersResponse) GetHeaders() []*entities.BlockHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBestHeaderRequest)(nil), "rpc.GetBestHeaderRequest")
	proto.RegisterType((*GetBestHeaderResponse)(nil), "rpc.GetBestHeaderResponse")
	proto.RegisterType((*GetHeadersRequest)(nil), "rpc.GetHeadersRequest")
	proto.RegisterType((*GetHeadersResponse)(nil), "rpc.GetHeadersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SyncService service

type SyncServiceClient interface {
	GetBestHeader(ctx context.Context, in *GetBestHeaderRequest, opts ...grpc.CallOption) (*GetBestHeaderResponse, error)
	GetHeaders(ctx context.Context, in *GetHeadersRequest, opts ...grpc.CallOption) (*GetHeadersResponse, error)
}

type syncServiceClient struct {
	cc *grpc.ClientConn
}

func NewSyncServiceClient(cc *grpc.ClientConn) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) GetBestHeader(ctx context.Context, in *GetBestHeaderRequest, opts ...grpc.CallOption) (*GetBestHeaderResponse, error) {
	out := new(GetBestHeaderResponse)
	err := grpc.Invoke(ctx, "/rpc.SyncService/GetBestHeader", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetHeaders(ctx context.Context, in *GetHeadersRequest, opts ...grpc.CallOption) (*GetHeadersResponse, error) {
	out := new(GetHeadersResponse)
	err := grpc.Invoke(ctx, "/rpc.SyncService/GetHeaders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SyncService service

type SyncServiceServer interface {
	GetBestHeader(context.Context, *GetBestHeaderRequest) (*GetBestHeaderResponse, error)
	GetHeaders(context.Context, *GetHeadersRequest) (*GetHeadersResponse, error)
}

func RegisterSyncServiceServer(s *grpc.Server, srv SyncServiceServer) {
	s.RegisterService(&_SyncService_serviceDesc, srv)
}

func _SyncService_GetBestHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBestHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetBestHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.SyncService/GetBestHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetBestHeader(ctx, req.(*GetBestHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.SyncService/GetHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetHeaders(ctx, req.(*GetHeadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBestHeader",
			Handler:    _SyncService_GetBestHeader_Handler,
		},
		{
			MethodName: "GetHeaders",
			Handler:    _SyncService_GetHeaders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/sync.proto",
}

func init() { proto.RegisterFile("rpc/sync.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0xd5, 0x94, 0x3f, 0xd2, 0x85, 0x22, 0x71, 0xa2, 0x6d, 0xc8, 0x80, 0xa2, 0x4c, 0x5d,
	0x48, 0xa4, 0xee, 0x2c, 0x41, 0x6a, 0x61, 0x4d, 0xc5, 0xc2, 0x16, 0x39, 0xa7, 0x36, 0xa2, 0xb2,
	0x8d, 0xcf, 0x45, 0xf4, 0x3d, 0x78, 0x60, 0xd4, 0xd8, 0x85, 0x14, 0xca, 0x78, 0xdf, 0xfd, 0xfc,
	0xf9, 0x7c, 0x86, 0x4b, 0xa3, 0x45, 0xce, 0x5b, 0x29, 0x32, 0x6d, 0x94, 0x55, 0xd8, 0x37, 0x5a,
	0xc4, 0x63, 0x92, 0x42, 0xd5, 0x8d, 0x5c, 0xe6, 0x24, 0x6d, 0x63, 0x1b, 0x62, 0xd7, 0x4d, 0x47,
	0x70, 0x3d, 0x27, 0x5b, 0x10, 0xdb, 0x47, 0xaa, 0x6a, 0x32, 0x25, 0xbd, 0x6d, 0x88, 0x6d, 0x3a,
	0x83, 0xe1, 0x2f, 0xce, 0x5a, 0x49, 0x26, 0xbc, 0x83, 0xb3, 0x55, 0x4b, 0xa2, 0x5e, 0x12, 0x4c,
	0xc2, 0xe9, 0x30, 0xfb, 0x36, 0x16, 0x6b, 0x25, 0x5e, 0x7d, 0xdc, 0x87, 0xd2, 0x67, 0xb8, 0x9a,
	0x93, 0x77, 0xb0, 0x97, 0xe3, 0x2d, 0xc0, 0xba, 0xb2, 0x3b, 0x77, 0xc5, 0xab, 0xd6, 0x73, 0x51,
	0x76, 0x08, 0x26, 0x10, 0xba, 0xea, 0x49, 0xd6, 0xf4, 0x11, 0x05, 0x49, 0x30, 0x39, 0x29, 0xbb,
	0x28, 0x5d, 0x02, 0x76, 0xb5, 0x7e, 0xb6, 0x04, 0x42, 0x77, 0xed, 0x83, 0xda, 0x48, 0xdb, 0x8a,
	0x07, 0x65, 0x17, 0x61, 0x0e, 0xe7, 0xae, 0xe4, 0x28, 0x48, 0xfa, 0xff, 0x8f, 0xbf, 0x4f, 0x4d,
	0x3f, 0x7b, 0x10, 0x2e, 0xb6, 0x52, 0x2c, 0xc8, 0xbc, 0x37, 0x82, 0x70, 0x06, 0x83, 0x83, 0xbd,
	0xe0, 0x4d, 0x66, 0xb4, 0xc8, 0x8e, 0xed, 0x30, 0x8e, 0x8f, 0xb5, 0xfc, 0xa8, 0xf7, 0x00, 0x3f,
	0x0f, 0xc0, 0xd1, 0x3e, 0x79, 0xb8, 0xa8, 0x78, 0xfc, 0x87, 0xbb, 0xe3, 0xc5, 0xe9, 0xcb, 0xee,
	0x5b, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x01, 0xf2, 0x2c, 0x8f, 0xec, 0x01, 0x00, 0x00,
}
