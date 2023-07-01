// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/evmutil/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryParamsRequest defines the request type for querying x/evmutil parameters.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb6e1bbb66ce140, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse defines the response type for querying x/evmutil parameters.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb6e1bbb66ce140, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryDeployedCosmosCoinContractsRequest defines the request type for Query/DeployedCosmosCoinContracts method.
type QueryDeployedCosmosCoinContractsRequest struct {
	// optional query param to only return specific denoms in the list
	// denoms that do not have deployed contracts will be omitted from the result
	// must request fewer than 100 denoms at a time.
	CosmosDenoms []string `protobuf:"bytes,1,rep,name=cosmos_denoms,json=cosmosDenoms,proto3" json:"cosmos_denoms,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDeployedCosmosCoinContractsRequest) Reset() {
	*m = QueryDeployedCosmosCoinContractsRequest{}
}
func (m *QueryDeployedCosmosCoinContractsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDeployedCosmosCoinContractsRequest) ProtoMessage()    {}
func (*QueryDeployedCosmosCoinContractsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb6e1bbb66ce140, []int{2}
}
func (m *QueryDeployedCosmosCoinContractsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDeployedCosmosCoinContractsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDeployedCosmosCoinContractsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDeployedCosmosCoinContractsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDeployedCosmosCoinContractsRequest.Merge(m, src)
}
func (m *QueryDeployedCosmosCoinContractsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDeployedCosmosCoinContractsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDeployedCosmosCoinContractsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDeployedCosmosCoinContractsRequest proto.InternalMessageInfo

// QueryDeployedCosmosCoinContractsResponse defines the response type for the Query/DeployedCosmosCoinContracts method.
type QueryDeployedCosmosCoinContractsResponse struct {
	// deployed_cosmos_coin_contracts is a list of cosmos-sdk coin denom and its deployed contract address
	DeployedCosmosCoinContracts []DeployedCosmosCoinContract `protobuf:"bytes,1,rep,name=deployed_cosmos_coin_contracts,json=deployedCosmosCoinContracts,proto3" json:"deployed_cosmos_coin_contracts"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDeployedCosmosCoinContractsResponse) Reset() {
	*m = QueryDeployedCosmosCoinContractsResponse{}
}
func (m *QueryDeployedCosmosCoinContractsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDeployedCosmosCoinContractsResponse) ProtoMessage()    {}
func (*QueryDeployedCosmosCoinContractsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb6e1bbb66ce140, []int{3}
}
func (m *QueryDeployedCosmosCoinContractsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDeployedCosmosCoinContractsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDeployedCosmosCoinContractsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDeployedCosmosCoinContractsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDeployedCosmosCoinContractsResponse.Merge(m, src)
}
func (m *QueryDeployedCosmosCoinContractsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDeployedCosmosCoinContractsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDeployedCosmosCoinContractsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDeployedCosmosCoinContractsResponse proto.InternalMessageInfo

func (m *QueryDeployedCosmosCoinContractsResponse) GetDeployedCosmosCoinContracts() []DeployedCosmosCoinContract {
	if m != nil {
		return m.DeployedCosmosCoinContracts
	}
	return nil
}

func (m *QueryDeployedCosmosCoinContractsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// DeployedCosmosCoinContract defines a deployed token contract to the evm representing a native cosmos-sdk coin
type DeployedCosmosCoinContract struct {
	CosmosDenom string              `protobuf:"bytes,1,opt,name=cosmos_denom,json=cosmosDenom,proto3" json:"cosmos_denom,omitempty"`
	Address     *InternalEVMAddress `protobuf:"bytes,2,opt,name=address,proto3,customtype=InternalEVMAddress" json:"address,omitempty"`
}

func (m *DeployedCosmosCoinContract) Reset()         { *m = DeployedCosmosCoinContract{} }
func (m *DeployedCosmosCoinContract) String() string { return proto.CompactTextString(m) }
func (*DeployedCosmosCoinContract) ProtoMessage()    {}
func (*DeployedCosmosCoinContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb6e1bbb66ce140, []int{4}
}
func (m *DeployedCosmosCoinContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeployedCosmosCoinContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeployedCosmosCoinContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeployedCosmosCoinContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployedCosmosCoinContract.Merge(m, src)
}
func (m *DeployedCosmosCoinContract) XXX_Size() int {
	return m.Size()
}
func (m *DeployedCosmosCoinContract) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployedCosmosCoinContract.DiscardUnknown(m)
}

var xxx_messageInfo_DeployedCosmosCoinContract proto.InternalMessageInfo

func (m *DeployedCosmosCoinContract) GetCosmosDenom() string {
	if m != nil {
		return m.CosmosDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "fury.evmutil.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "fury.evmutil.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryDeployedCosmosCoinContractsRequest)(nil), "fury.evmutil.v1beta1.QueryDeployedCosmosCoinContractsRequest")
	proto.RegisterType((*QueryDeployedCosmosCoinContractsResponse)(nil), "fury.evmutil.v1beta1.QueryDeployedCosmosCoinContractsResponse")
	proto.RegisterType((*DeployedCosmosCoinContract)(nil), "fury.evmutil.v1beta1.DeployedCosmosCoinContract")
}

func init() { proto.RegisterFile("fury/evmutil/v1beta1/query.proto", fileDescriptor_acb6e1bbb66ce140) }

var fileDescriptor_acb6e1bbb66ce140 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xb7, 0x03, 0x04, 0x72, 0x29, 0xcb, 0x11, 0xa1, 0x2a, 0x8d, 0x9c, 0x62, 0x10, 0x0d, 0x48,
	0xf8, 0xda, 0x80, 0x18, 0x2a, 0x40, 0x22, 0x29, 0xa0, 0x0e, 0x48, 0xd4, 0x03, 0x03, 0x4b, 0x74,
	0xb6, 0x0f, 0x73, 0x22, 0xb9, 0x73, 0x7c, 0xe7, 0x42, 0xc4, 0x06, 0x0b, 0x23, 0x12, 0x5f, 0x20,
	0x1f, 0xa7, 0x63, 0x25, 0x16, 0xd4, 0xa1, 0x42, 0x09, 0x03, 0x62, 0xe2, 0x23, 0xa0, 0xdc, 0x5d,
	0xda, 0x20, 0x9c, 0x04, 0xb1, 0x9d, 0x5e, 0x7e, 0xef, 0xfd, 0xfe, 0xbc, 0x17, 0x83, 0xf5, 0x97,
	0x59, 0x3a, 0x40, 0x64, 0xbf, 0x97, 0x49, 0xda, 0x45, 0xfb, 0x5b, 0x01, 0x91, 0x78, 0x0b, 0xf5,
	0x33, 0x92, 0x0e, 0xbc, 0x24, 0xe5, 0x92, 0xc3, 0xca, 0x04, 0xe1, 0x19, 0x84, 0x67, 0x10, 0xd5,
	0x9b, 0x21, 0x17, 0x3d, 0x2e, 0x50, 0x80, 0x05, 0xd1, 0xf0, 0x93, 0xe6, 0x04, 0xc7, 0x94, 0x61,
	0x49, 0x39, 0xd3, 0x13, 0xaa, 0x6e, 0x2e, 0x47, 0x4c, 0x18, 0x11, 0x54, 0x18, 0x4c, 0x25, 0xe6,
	0x31, 0x57, 0x4f, 0x34, 0x79, 0x99, 0x6a, 0x2d, 0xe6, 0x3c, 0xee, 0x12, 0x84, 0x13, 0x8a, 0x30,
	0x63, 0x5c, 0xaa, 0xb1, 0xa6, 0xc7, 0xad, 0x00, 0xb8, 0x37, 0x61, 0x7e, 0x86, 0x53, 0xdc, 0x13,
	0x3e, 0xe9, 0x67, 0x44, 0x48, 0x77, 0x0f, 0x5c, 0xfa, 0xa3, 0x2a, 0x12, 0xce, 0x04, 0x81, 0xdb,
	0xa0, 0x98, 0xa8, 0xca, 0xaa, 0xbd, 0x6e, 0x37, 0xca, 0xcd, 0x9a, 0x97, 0xe7, 0xcb, 0xd3, 0x5d,
	0xad, 0xb3, 0x07, 0xc7, 0x75, 0xcb, 0x37, 0x1d, 0xee, 0xd0, 0x06, 0x1b, 0x6a, 0xe6, 0x0e, 0x49,
	0xba, 0x7c, 0x40, 0xa2, 0xb6, 0x32, 0xdf, 0xe6, 0x94, 0xb5, 0x39, 0x93, 0x29, 0x0e, 0xe5, 0x94,
	0x1e, 0x5e, 0x05, 0x17, 0x75, 0x34, 0x9d, 0x88, 0x30, 0xae, 0xe8, 0xce, 0x34, 0x4a, 0xfe, 0x8a,
	0x2e, 0xee, 0xa8, 0x1a, 0x7c, 0x0c, 0xc0, 0x69, 0x4a, 0xab, 0x05, 0x25, 0xe8, 0xba, 0xa7, 0x21,
	0xde, 0x24, 0x52, 0x4f, 0x6f, 0xe0, 0x54, 0x55, 0x4c, 0x0c, 0x81, 0x3f, 0xd3, 0xb9, 0x7d, 0xe1,
	0xe3, 0xb0, 0x6e, 0xfd, 0x18, 0xd6, 0x2d, 0xf7, 0x97, 0x0d, 0x1a, 0xcb, 0x25, 0x9a, 0x2c, 0xde,
	0x01, 0x27, 0x32, 0xb0, 0x8e, 0x11, 0x1b, 0x72, 0xca, 0x3a, 0xe1, 0x14, 0xa9, 0x44, 0x97, 0x9b,
	0x9b, 0xf9, 0x19, 0xcd, 0xa7, 0x30, 0xb9, 0xad, 0x45, 0xf3, 0x45, 0xc0, 0x27, 0x39, 0xde, 0x37,
	0x96, 0x7a, 0xd7, 0xca, 0x67, 0xcd, 0xbb, 0x7d, 0x50, 0x9d, 0xaf, 0x04, 0x5e, 0x01, 0x2b, 0xb3,
	0x7b, 0x50, 0x5b, 0x2f, 0xf9, 0xe5, 0x99, 0x35, 0xc0, 0x4d, 0x70, 0x1e, 0x47, 0x51, 0x4a, 0x84,
	0x50, 0x32, 0x4a, 0xad, 0xcb, 0x47, 0xc7, 0x75, 0xb8, 0xcb, 0x24, 0x49, 0x19, 0xee, 0x3e, 0x7a,
	0xfe, 0xf4, 0xa1, 0xfe, 0xd5, 0x9f, 0xc2, 0x9a, 0x3f, 0x0b, 0xe0, 0x9c, 0x4a, 0x19, 0x7e, 0xb0,
	0x41, 0x51, 0xdf, 0x0a, 0x6c, 0xe4, 0xa7, 0xf4, 0xf7, 0x69, 0x56, 0x6f, 0xfc, 0x03, 0x52, 0x1b,
	0x75, 0xaf, 0xbd, 0xff, 0xf2, 0xfd, 0x73, 0xc1, 0x81, 0x35, 0x94, 0xfb, 0xe7, 0xd1, 0x87, 0x09,
	0x8f, 0x6c, 0xb0, 0xb6, 0x60, 0xe1, 0xf0, 0xfe, 0x02, 0xc2, 0xe5, 0xb7, 0x5c, 0x7d, 0xf0, 0xbf,
	0xed, 0xc6, 0xc4, 0x3d, 0x65, 0xe2, 0x2e, 0xbc, 0x93, 0x6f, 0x62, 0xf1, 0x0d, 0xb6, 0x76, 0x0f,
	0x46, 0x8e, 0x7d, 0x38, 0x72, 0xec, 0x6f, 0x23, 0xc7, 0xfe, 0x34, 0x76, 0xac, 0xc3, 0xb1, 0x63,
	0x7d, 0x1d, 0x3b, 0xd6, 0x0b, 0x14, 0x53, 0xf9, 0x2a, 0x0b, 0xbc, 0x90, 0xf7, 0x10, 0x65, 0x61,
	0x16, 0x64, 0xe2, 0x16, 0x23, 0xf2, 0x0d, 0x4f, 0x5f, 0x6b, 0xa6, 0xb7, 0x27, 0x5c, 0x72, 0x90,
	0x10, 0x11, 0x14, 0xd5, 0x07, 0xe3, 0xf6, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x76, 0xf6,
	0x39, 0xee, 0x04, 0x00, 0x00,
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
	// Params queries all parameters of the evmutil module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// DeployedCosmosCoinContracts queries a list cosmos coin denom and their deployed erc20 address
	DeployedCosmosCoinContracts(ctx context.Context, in *QueryDeployedCosmosCoinContractsRequest, opts ...grpc.CallOption) (*QueryDeployedCosmosCoinContractsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/fury.evmutil.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DeployedCosmosCoinContracts(ctx context.Context, in *QueryDeployedCosmosCoinContractsRequest, opts ...grpc.CallOption) (*QueryDeployedCosmosCoinContractsResponse, error) {
	out := new(QueryDeployedCosmosCoinContractsResponse)
	err := c.cc.Invoke(ctx, "/fury.evmutil.v1beta1.Query/DeployedCosmosCoinContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries all parameters of the evmutil module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// DeployedCosmosCoinContracts queries a list cosmos coin denom and their deployed erc20 address
	DeployedCosmosCoinContracts(context.Context, *QueryDeployedCosmosCoinContractsRequest) (*QueryDeployedCosmosCoinContractsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DeployedCosmosCoinContracts(ctx context.Context, req *QueryDeployedCosmosCoinContractsRequest) (*QueryDeployedCosmosCoinContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployedCosmosCoinContracts not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fury.evmutil.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DeployedCosmosCoinContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeployedCosmosCoinContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DeployedCosmosCoinContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fury.evmutil.v1beta1.Query/DeployedCosmosCoinContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DeployedCosmosCoinContracts(ctx, req.(*QueryDeployedCosmosCoinContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fury.evmutil.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DeployedCosmosCoinContracts",
			Handler:    _Query_DeployedCosmosCoinContracts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fury/evmutil/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryDeployedCosmosCoinContractsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDeployedCosmosCoinContractsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDeployedCosmosCoinContractsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CosmosDenoms) > 0 {
		for iNdEx := len(m.CosmosDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CosmosDenoms[iNdEx])
			copy(dAtA[i:], m.CosmosDenoms[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.CosmosDenoms[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDeployedCosmosCoinContractsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDeployedCosmosCoinContractsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDeployedCosmosCoinContractsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DeployedCosmosCoinContracts) > 0 {
		for iNdEx := len(m.DeployedCosmosCoinContracts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DeployedCosmosCoinContracts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DeployedCosmosCoinContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeployedCosmosCoinContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeployedCosmosCoinContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Address != nil {
		{
			size := m.Address.Size()
			i -= size
			if _, err := m.Address.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CosmosDenom) > 0 {
		i -= len(m.CosmosDenom)
		copy(dAtA[i:], m.CosmosDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CosmosDenom)))
		i--
		dAtA[i] = 0xa
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryDeployedCosmosCoinContractsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CosmosDenoms) > 0 {
		for _, s := range m.CosmosDenoms {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDeployedCosmosCoinContractsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DeployedCosmosCoinContracts) > 0 {
		for _, e := range m.DeployedCosmosCoinContracts {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *DeployedCosmosCoinContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CosmosDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Address != nil {
		l = m.Address.Size()
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDeployedCosmosCoinContractsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDeployedCosmosCoinContractsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDeployedCosmosCoinContractsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosDenoms", wireType)
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
			m.CosmosDenoms = append(m.CosmosDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDeployedCosmosCoinContractsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDeployedCosmosCoinContractsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDeployedCosmosCoinContractsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeployedCosmosCoinContracts", wireType)
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
			m.DeployedCosmosCoinContracts = append(m.DeployedCosmosCoinContracts, DeployedCosmosCoinContract{})
			if err := m.DeployedCosmosCoinContracts[len(m.DeployedCosmosCoinContracts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *DeployedCosmosCoinContract) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DeployedCosmosCoinContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeployedCosmosCoinContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosDenom", wireType)
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
			m.CosmosDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			var v InternalEVMAddress
			m.Address = &v
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
