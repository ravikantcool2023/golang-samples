// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shakesapp.proto

package shakesapp

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ShakespeareResponse struct {
	// match_count is the number of matching lines.
	MatchCount           int64    `protobuf:"varint,1,opt,name=match_count,json=matchCount,proto3" json:"match_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShakespeareResponse) Reset()         { *m = ShakespeareResponse{} }
func (m *ShakespeareResponse) String() string { return proto.CompactTextString(m) }
func (*ShakespeareResponse) ProtoMessage()    {}
func (*ShakespeareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b49d6006aafbed, []int{0}
}

func (m *ShakespeareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShakespeareResponse.Unmarshal(m, b)
}
func (m *ShakespeareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShakespeareResponse.Marshal(b, m, deterministic)
}
func (m *ShakespeareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShakespeareResponse.Merge(m, src)
}
func (m *ShakespeareResponse) XXX_Size() int {
	return xxx_messageInfo_ShakespeareResponse.Size(m)
}
func (m *ShakespeareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShakespeareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShakespeareResponse proto.InternalMessageInfo

func (m *ShakespeareResponse) GetMatchCount() int64 {
	if m != nil {
		return m.MatchCount
	}
	return 0
}

type ShakespeareRequest struct {
	// query is a substring query.
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShakespeareRequest) Reset()         { *m = ShakespeareRequest{} }
func (m *ShakespeareRequest) String() string { return proto.CompactTextString(m) }
func (*ShakespeareRequest) ProtoMessage()    {}
func (*ShakespeareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80b49d6006aafbed, []int{1}
}

func (m *ShakespeareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShakespeareRequest.Unmarshal(m, b)
}
func (m *ShakespeareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShakespeareRequest.Marshal(b, m, deterministic)
}
func (m *ShakespeareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShakespeareRequest.Merge(m, src)
}
func (m *ShakespeareRequest) XXX_Size() int {
	return xxx_messageInfo_ShakespeareRequest.Size(m)
}
func (m *ShakespeareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShakespeareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShakespeareRequest proto.InternalMessageInfo

func (m *ShakespeareRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func init() {
	proto.RegisterType((*ShakespeareResponse)(nil), "shakesapp.ShakespeareResponse")
	proto.RegisterType((*ShakespeareRequest)(nil), "shakesapp.ShakespeareRequest")
}

func init() { proto.RegisterFile("shakesapp.proto", fileDescriptor_80b49d6006aafbed) }

var fileDescriptor_80b49d6006aafbed = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xce, 0x48, 0xcc,
	0x4e, 0x2d, 0x4e, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x99, 0x71, 0x09, 0x07, 0x83, 0x39, 0x05, 0xa9, 0x89, 0x45, 0xa9, 0x41, 0xa9, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0x42, 0xf2, 0x5c, 0xdc, 0xb9, 0x89, 0x25, 0xc9, 0x19, 0xf1, 0xc9, 0xf9, 0xa5,
	0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x5c, 0x60, 0x21, 0x67, 0x90, 0x88, 0x92,
	0x16, 0x97, 0x10, 0x8a, 0xbe, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0xd6, 0xc2, 0xd2,
	0xd4, 0xa2, 0x4a, 0xb0, 0x06, 0xce, 0x20, 0x08, 0xc7, 0x28, 0x0d, 0x45, 0x6d, 0x70, 0x6a, 0x51,
	0x59, 0x66, 0x72, 0xaa, 0x50, 0x00, 0x17, 0xaf, 0x7b, 0x6a, 0x89, 0x2f, 0xdc, 0x48, 0x21, 0x59,
	0x3d, 0x84, 0x3b, 0x31, 0xcd, 0x96, 0x92, 0xc3, 0x25, 0x0d, 0x71, 0xb2, 0x12, 0x43, 0x12, 0x1b,
	0xd8, 0x77, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0xbf, 0x3e, 0x5a, 0xf0, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShakespeareServiceClient is the client API for ShakespeareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShakespeareServiceClient interface {
	// Accepts a query string and returns the number of lines containing that.
	GetMatchCount(ctx context.Context, in *ShakespeareRequest, opts ...grpc.CallOption) (*ShakespeareResponse, error)
}

type shakespeareServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShakespeareServiceClient(cc grpc.ClientConnInterface) ShakespeareServiceClient {
	return &shakespeareServiceClient{cc}
}

func (c *shakespeareServiceClient) GetMatchCount(ctx context.Context, in *ShakespeareRequest, opts ...grpc.CallOption) (*ShakespeareResponse, error) {
	out := new(ShakespeareResponse)
	err := c.cc.Invoke(ctx, "/shakesapp.ShakespeareService/GetMatchCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShakespeareServiceServer is the server API for ShakespeareService service.
type ShakespeareServiceServer interface {
	// Accepts a query string and returns the number of lines containing that.
	GetMatchCount(context.Context, *ShakespeareRequest) (*ShakespeareResponse, error)
}

// UnimplementedShakespeareServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShakespeareServiceServer struct {
}

func (*UnimplementedShakespeareServiceServer) GetMatchCount(ctx context.Context, req *ShakespeareRequest) (*ShakespeareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchCount not implemented")
}

func RegisterShakespeareServiceServer(s *grpc.Server, srv ShakespeareServiceServer) {
	s.RegisterService(&_ShakespeareService_serviceDesc, srv)
}

func _ShakespeareService_GetMatchCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShakespeareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShakespeareServiceServer).GetMatchCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shakesapp.ShakespeareService/GetMatchCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShakespeareServiceServer).GetMatchCount(ctx, req.(*ShakespeareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShakespeareService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shakesapp.ShakespeareService",
	HandlerType: (*ShakespeareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMatchCount",
			Handler:    _ShakespeareService_GetMatchCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shakesapp.proto",
}