// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/lightclient/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type QueryGetLightClientRequest struct {
	RollappId string `protobuf:"bytes,1,opt,name=rollapp_id,json=rollappId,proto3" json:"rollapp_id,omitempty"`
}

func (m *QueryGetLightClientRequest) Reset()         { *m = QueryGetLightClientRequest{} }
func (m *QueryGetLightClientRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetLightClientRequest) ProtoMessage()    {}
func (*QueryGetLightClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a51f5810cc34625d, []int{0}
}
func (m *QueryGetLightClientRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetLightClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetLightClientRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetLightClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetLightClientRequest.Merge(m, src)
}
func (m *QueryGetLightClientRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetLightClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetLightClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetLightClientRequest proto.InternalMessageInfo

func (m *QueryGetLightClientRequest) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

type QueryGetLightClientResponse struct {
	// empty if doesn't exist
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *QueryGetLightClientResponse) Reset()         { *m = QueryGetLightClientResponse{} }
func (m *QueryGetLightClientResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetLightClientResponse) ProtoMessage()    {}
func (*QueryGetLightClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a51f5810cc34625d, []int{1}
}
func (m *QueryGetLightClientResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetLightClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetLightClientResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetLightClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetLightClientResponse.Merge(m, src)
}
func (m *QueryGetLightClientResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetLightClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetLightClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetLightClientResponse proto.InternalMessageInfo

func (m *QueryGetLightClientResponse) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type QueryExpectedClientStateRequest struct {
}

func (m *QueryExpectedClientStateRequest) Reset()         { *m = QueryExpectedClientStateRequest{} }
func (m *QueryExpectedClientStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryExpectedClientStateRequest) ProtoMessage()    {}
func (*QueryExpectedClientStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a51f5810cc34625d, []int{2}
}
func (m *QueryExpectedClientStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryExpectedClientStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryExpectedClientStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryExpectedClientStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryExpectedClientStateRequest.Merge(m, src)
}
func (m *QueryExpectedClientStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryExpectedClientStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryExpectedClientStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryExpectedClientStateRequest proto.InternalMessageInfo

type QueryExpectedClientStateResponse struct {
	// client state
	ClientState *types.Any `protobuf:"bytes,2,opt,name=client_state,json=clientState,proto3" json:"client_state,omitempty" yaml:"client_state"`
}

func (m *QueryExpectedClientStateResponse) Reset()         { *m = QueryExpectedClientStateResponse{} }
func (m *QueryExpectedClientStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryExpectedClientStateResponse) ProtoMessage()    {}
func (*QueryExpectedClientStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a51f5810cc34625d, []int{3}
}
func (m *QueryExpectedClientStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryExpectedClientStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryExpectedClientStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryExpectedClientStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryExpectedClientStateResponse.Merge(m, src)
}
func (m *QueryExpectedClientStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryExpectedClientStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryExpectedClientStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryExpectedClientStateResponse proto.InternalMessageInfo

func (m *QueryExpectedClientStateResponse) GetClientState() *types.Any {
	if m != nil {
		return m.ClientState
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetLightClientRequest)(nil), "dymensionxyz.dymension.lightclient.QueryGetLightClientRequest")
	proto.RegisterType((*QueryGetLightClientResponse)(nil), "dymensionxyz.dymension.lightclient.QueryGetLightClientResponse")
	proto.RegisterType((*QueryExpectedClientStateRequest)(nil), "dymensionxyz.dymension.lightclient.QueryExpectedClientStateRequest")
	proto.RegisterType((*QueryExpectedClientStateResponse)(nil), "dymensionxyz.dymension.lightclient.QueryExpectedClientStateResponse")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/lightclient/query.proto", fileDescriptor_a51f5810cc34625d)
}

var fileDescriptor_a51f5810cc34625d = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xab, 0xa8, 0xac, 0xd2, 0x87, 0x73, 0xf4, 0x73, 0x32, 0xd3, 0x33,
	0x4a, 0x92, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0xf4, 0x0b, 0x4b, 0x53, 0x8b, 0x2a, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0x94, 0x90, 0xd5, 0x23, 0x34, 0xeb, 0x21, 0xa9, 0x97, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x2b, 0xd7, 0x07, 0xb1, 0x20, 0x3a, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73,
	0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3,
	0x8a, 0xa1, 0xb2, 0x92, 0x50, 0x59, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0x3f, 0x31, 0x0f, 0x6a, 0xa5,
	0x92, 0x23, 0x97, 0x54, 0x20, 0xc8, 0x05, 0xee, 0xa9, 0x25, 0x3e, 0x20, 0x5b, 0x9c, 0xc1, 0xb6,
	0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x72, 0x71, 0x15, 0xe5, 0xe7, 0xe4, 0x24,
	0x16, 0x14, 0xc4, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x42, 0x45, 0x3c,
	0x53, 0xac, 0x58, 0x5e, 0x2c, 0x90, 0x67, 0x50, 0xb2, 0xe2, 0x92, 0xc6, 0x6a, 0x44, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x34, 0x17, 0x27, 0xc4, 0xe9, 0x20, 0x23, 0x98, 0xc0, 0x46, 0x70,
	0x40, 0x04, 0x3c, 0x53, 0x94, 0x14, 0xb9, 0xe4, 0xc1, 0x7a, 0x5d, 0x2b, 0x0a, 0x52, 0x93, 0x4b,
	0x52, 0x53, 0x20, 0x7a, 0x83, 0x4b, 0x12, 0x4b, 0x52, 0xa1, 0x6e, 0x50, 0x2a, 0xe1, 0x52, 0xc0,
	0xad, 0x04, 0x6a, 0x47, 0x00, 0x17, 0x0f, 0xd4, 0x8e, 0x62, 0x90, 0x38, 0xd8, 0x1a, 0x6e, 0x23,
	0x11, 0x3d, 0x88, 0xbf, 0xf5, 0x60, 0xfe, 0xd6, 0x73, 0xcc, 0xab, 0x74, 0x12, 0xff, 0x74, 0x4f,
	0x5e, 0xb8, 0x32, 0x31, 0x37, 0xc7, 0x4a, 0x09, 0x59, 0x8f, 0x52, 0x10, 0x77, 0x32, 0xc2, 0x64,
	0xa3, 0x43, 0xcc, 0x5c, 0xac, 0x60, 0x6b, 0x85, 0xae, 0x30, 0x72, 0x71, 0x23, 0xf9, 0x4b, 0xc8,
	0x4e, 0x8f, 0x70, 0x2c, 0xe9, 0xe1, 0x0e, 0x53, 0x29, 0x7b, 0xb2, 0xf5, 0x43, 0x3c, 0xab, 0xe4,
	0xd2, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x3b, 0x21, 0x1b, 0x7d, 0x22, 0x92, 0x17, 0x32, 0xbb, 0x1a,
	0x11, 0x97, 0xb5, 0x42, 0x0f, 0x19, 0xb9, 0x84, 0xb1, 0x04, 0xa9, 0x90, 0x33, 0xd1, 0xce, 0xc3,
	0x1d, 0x67, 0x52, 0x2e, 0x94, 0x19, 0x02, 0xf5, 0xa8, 0x3d, 0xd8, 0xa3, 0x96, 0x42, 0xe6, 0xc4,
	0x78, 0x34, 0x15, 0x6a, 0x10, 0x84, 0x0b, 0x8e, 0x52, 0xa7, 0xa0, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xc5, 0x65, 0x78, 0x99, 0xb1, 0x7e, 0x05, 0x8a, 0x0d, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c,
	0xe0, 0xc4, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x06, 0x98, 0x0c, 0x48, 0xdc, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	LightClient(ctx context.Context, in *QueryGetLightClientRequest, opts ...grpc.CallOption) (*QueryGetLightClientResponse, error)
	ExpectedClientState(ctx context.Context, in *QueryExpectedClientStateRequest, opts ...grpc.CallOption) (*QueryExpectedClientStateResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) LightClient(ctx context.Context, in *QueryGetLightClientRequest, opts ...grpc.CallOption) (*QueryGetLightClientResponse, error) {
	out := new(QueryGetLightClientResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.lightclient.Query/LightClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ExpectedClientState(ctx context.Context, in *QueryExpectedClientStateRequest, opts ...grpc.CallOption) (*QueryExpectedClientStateResponse, error) {
	out := new(QueryExpectedClientStateResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.lightclient.Query/ExpectedClientState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	LightClient(context.Context, *QueryGetLightClientRequest) (*QueryGetLightClientResponse, error)
	ExpectedClientState(context.Context, *QueryExpectedClientStateRequest) (*QueryExpectedClientStateResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) LightClient(ctx context.Context, req *QueryGetLightClientRequest) (*QueryGetLightClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LightClient not implemented")
}
func (*UnimplementedQueryServer) ExpectedClientState(ctx context.Context, req *QueryExpectedClientStateRequest) (*QueryExpectedClientStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpectedClientState not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_LightClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetLightClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LightClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.lightclient.Query/LightClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LightClient(ctx, req.(*QueryGetLightClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ExpectedClientState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExpectedClientStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ExpectedClientState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.lightclient.Query/ExpectedClientState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ExpectedClientState(ctx, req.(*QueryExpectedClientStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dymensionxyz.dymension.lightclient.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LightClient",
			Handler:    _Query_LightClient_Handler,
		},
		{
			MethodName: "ExpectedClientState",
			Handler:    _Query_ExpectedClientState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dymensionxyz/dymension/lightclient/query.proto",
}

func (m *QueryGetLightClientRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetLightClientRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetLightClientRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetLightClientResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetLightClientResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetLightClientResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *QueryExpectedClientStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryExpectedClientStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryExpectedClientStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryExpectedClientStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryExpectedClientStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryExpectedClientStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientState != nil {
		{
			size, err := m.ClientState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetLightClientRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetLightClientResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryExpectedClientStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryExpectedClientStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClientState != nil {
		l = m.ClientState.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetLightClientRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetLightClientRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetLightClientRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetLightClientResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetLightClientResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetLightClientResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryExpectedClientStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryExpectedClientStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryExpectedClientStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryExpectedClientStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryExpectedClientStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryExpectedClientStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClientState == nil {
				m.ClientState = &types.Any{}
			}
			if err := m.ClientState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)