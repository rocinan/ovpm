// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package pb

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

type NetworkCreateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	CIDR string `protobuf:"bytes,2,opt,name=CIDR" json:"CIDR,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=Type" json:"Type,omitempty"`
	Via  string `protobuf:"bytes,4,opt,name=Via" json:"Via,omitempty"`
}

func (m *NetworkCreateRequest) Reset()                    { *m = NetworkCreateRequest{} }
func (m *NetworkCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkCreateRequest) ProtoMessage()               {}
func (*NetworkCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *NetworkCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkCreateRequest) GetCIDR() string {
	if m != nil {
		return m.CIDR
	}
	return ""
}

func (m *NetworkCreateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkCreateRequest) GetVia() string {
	if m != nil {
		return m.Via
	}
	return ""
}

type NetworkListRequest struct {
}

func (m *NetworkListRequest) Reset()                    { *m = NetworkListRequest{} }
func (m *NetworkListRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkListRequest) ProtoMessage()               {}
func (*NetworkListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type NetworkDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *NetworkDeleteRequest) Reset()                    { *m = NetworkDeleteRequest{} }
func (m *NetworkDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkDeleteRequest) ProtoMessage()               {}
func (*NetworkDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *NetworkDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NetworkGetAllTypesRequest struct {
}

func (m *NetworkGetAllTypesRequest) Reset()                    { *m = NetworkGetAllTypesRequest{} }
func (m *NetworkGetAllTypesRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkGetAllTypesRequest) ProtoMessage()               {}
func (*NetworkGetAllTypesRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type NetworkAssociateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
}

func (m *NetworkAssociateRequest) Reset()                    { *m = NetworkAssociateRequest{} }
func (m *NetworkAssociateRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkAssociateRequest) ProtoMessage()               {}
func (*NetworkAssociateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *NetworkAssociateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkAssociateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type NetworkDissociateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
}

func (m *NetworkDissociateRequest) Reset()                    { *m = NetworkDissociateRequest{} }
func (m *NetworkDissociateRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkDissociateRequest) ProtoMessage()               {}
func (*NetworkDissociateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *NetworkDissociateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkDissociateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Network struct {
	Name                string   `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	CIDR                string   `protobuf:"bytes,2,opt,name=CIDR" json:"CIDR,omitempty"`
	Type                string   `protobuf:"bytes,3,opt,name=Type" json:"Type,omitempty"`
	CreatedAt           string   `protobuf:"bytes,4,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
	AssociatedUsernames []string `protobuf:"bytes,5,rep,name=AssociatedUsernames" json:"AssociatedUsernames,omitempty"`
	Via                 string   `protobuf:"bytes,6,opt,name=Via" json:"Via,omitempty"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *Network) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Network) GetCIDR() string {
	if m != nil {
		return m.CIDR
	}
	return ""
}

func (m *Network) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Network) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Network) GetAssociatedUsernames() []string {
	if m != nil {
		return m.AssociatedUsernames
	}
	return nil
}

func (m *Network) GetVia() string {
	if m != nil {
		return m.Via
	}
	return ""
}

type NetworkCreateResponse struct {
	Network *Network `protobuf:"bytes,1,opt,name=Network" json:"Network,omitempty"`
}

func (m *NetworkCreateResponse) Reset()                    { *m = NetworkCreateResponse{} }
func (m *NetworkCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkCreateResponse) ProtoMessage()               {}
func (*NetworkCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *NetworkCreateResponse) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type NetworkListResponse struct {
	Networks []*Network `protobuf:"bytes,1,rep,name=Networks" json:"Networks,omitempty"`
}

func (m *NetworkListResponse) Reset()                    { *m = NetworkListResponse{} }
func (m *NetworkListResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkListResponse) ProtoMessage()               {}
func (*NetworkListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *NetworkListResponse) GetNetworks() []*Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

type NetworkDeleteResponse struct {
	Network *Network `protobuf:"bytes,1,opt,name=Network" json:"Network,omitempty"`
}

func (m *NetworkDeleteResponse) Reset()                    { *m = NetworkDeleteResponse{} }
func (m *NetworkDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkDeleteResponse) ProtoMessage()               {}
func (*NetworkDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *NetworkDeleteResponse) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

type NetworkGetAllTypesResponse struct {
	Types []string `protobuf:"bytes,1,rep,name=Types" json:"Types,omitempty"`
}

func (m *NetworkGetAllTypesResponse) Reset()                    { *m = NetworkGetAllTypesResponse{} }
func (m *NetworkGetAllTypesResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkGetAllTypesResponse) ProtoMessage()               {}
func (*NetworkGetAllTypesResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *NetworkGetAllTypesResponse) GetTypes() []string {
	if m != nil {
		return m.Types
	}
	return nil
}

type NetworkAssociateResponse struct {
}

func (m *NetworkAssociateResponse) Reset()                    { *m = NetworkAssociateResponse{} }
func (m *NetworkAssociateResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkAssociateResponse) ProtoMessage()               {}
func (*NetworkAssociateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

type NetworkDissociateResponse struct {
}

func (m *NetworkDissociateResponse) Reset()                    { *m = NetworkDissociateResponse{} }
func (m *NetworkDissociateResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkDissociateResponse) ProtoMessage()               {}
func (*NetworkDissociateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func init() {
	proto.RegisterType((*NetworkCreateRequest)(nil), "pb.NetworkCreateRequest")
	proto.RegisterType((*NetworkListRequest)(nil), "pb.NetworkListRequest")
	proto.RegisterType((*NetworkDeleteRequest)(nil), "pb.NetworkDeleteRequest")
	proto.RegisterType((*NetworkGetAllTypesRequest)(nil), "pb.NetworkGetAllTypesRequest")
	proto.RegisterType((*NetworkAssociateRequest)(nil), "pb.NetworkAssociateRequest")
	proto.RegisterType((*NetworkDissociateRequest)(nil), "pb.NetworkDissociateRequest")
	proto.RegisterType((*Network)(nil), "pb.Network")
	proto.RegisterType((*NetworkCreateResponse)(nil), "pb.NetworkCreateResponse")
	proto.RegisterType((*NetworkListResponse)(nil), "pb.NetworkListResponse")
	proto.RegisterType((*NetworkDeleteResponse)(nil), "pb.NetworkDeleteResponse")
	proto.RegisterType((*NetworkGetAllTypesResponse)(nil), "pb.NetworkGetAllTypesResponse")
	proto.RegisterType((*NetworkAssociateResponse)(nil), "pb.NetworkAssociateResponse")
	proto.RegisterType((*NetworkDissociateResponse)(nil), "pb.NetworkDissociateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkService service

type NetworkServiceClient interface {
	Create(ctx context.Context, in *NetworkCreateRequest, opts ...grpc.CallOption) (*NetworkCreateResponse, error)
	List(ctx context.Context, in *NetworkListRequest, opts ...grpc.CallOption) (*NetworkListResponse, error)
	Delete(ctx context.Context, in *NetworkDeleteRequest, opts ...grpc.CallOption) (*NetworkDeleteResponse, error)
	GetAllTypes(ctx context.Context, in *NetworkGetAllTypesRequest, opts ...grpc.CallOption) (*NetworkGetAllTypesResponse, error)
	Associate(ctx context.Context, in *NetworkAssociateRequest, opts ...grpc.CallOption) (*NetworkAssociateResponse, error)
	Dissociate(ctx context.Context, in *NetworkDissociateRequest, opts ...grpc.CallOption) (*NetworkDissociateResponse, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Create(ctx context.Context, in *NetworkCreateRequest, opts ...grpc.CallOption) (*NetworkCreateResponse, error) {
	out := new(NetworkCreateResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) List(ctx context.Context, in *NetworkListRequest, opts ...grpc.CallOption) (*NetworkListResponse, error) {
	out := new(NetworkListResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Delete(ctx context.Context, in *NetworkDeleteRequest, opts ...grpc.CallOption) (*NetworkDeleteResponse, error) {
	out := new(NetworkDeleteResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) GetAllTypes(ctx context.Context, in *NetworkGetAllTypesRequest, opts ...grpc.CallOption) (*NetworkGetAllTypesResponse, error) {
	out := new(NetworkGetAllTypesResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/GetAllTypes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Associate(ctx context.Context, in *NetworkAssociateRequest, opts ...grpc.CallOption) (*NetworkAssociateResponse, error) {
	out := new(NetworkAssociateResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/Associate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Dissociate(ctx context.Context, in *NetworkDissociateRequest, opts ...grpc.CallOption) (*NetworkDissociateResponse, error) {
	out := new(NetworkDissociateResponse)
	err := grpc.Invoke(ctx, "/pb.NetworkService/Dissociate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkService service

type NetworkServiceServer interface {
	Create(context.Context, *NetworkCreateRequest) (*NetworkCreateResponse, error)
	List(context.Context, *NetworkListRequest) (*NetworkListResponse, error)
	Delete(context.Context, *NetworkDeleteRequest) (*NetworkDeleteResponse, error)
	GetAllTypes(context.Context, *NetworkGetAllTypesRequest) (*NetworkGetAllTypesResponse, error)
	Associate(context.Context, *NetworkAssociateRequest) (*NetworkAssociateResponse, error)
	Dissociate(context.Context, *NetworkDissociateRequest) (*NetworkDissociateResponse, error)
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Create(ctx, req.(*NetworkCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).List(ctx, req.(*NetworkListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Delete(ctx, req.(*NetworkDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_GetAllTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkGetAllTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).GetAllTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/GetAllTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).GetAllTypes(ctx, req.(*NetworkGetAllTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Associate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkAssociateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Associate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/Associate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Associate(ctx, req.(*NetworkAssociateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Dissociate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkDissociateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Dissociate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NetworkService/Dissociate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Dissociate(ctx, req.(*NetworkDissociateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NetworkService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetworkService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetworkService_Delete_Handler,
		},
		{
			MethodName: "GetAllTypes",
			Handler:    _NetworkService_GetAllTypes_Handler,
		},
		{
			MethodName: "Associate",
			Handler:    _NetworkService_Associate_Handler,
		},
		{
			MethodName: "Dissociate",
			Handler:    _NetworkService_Dissociate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network.proto",
}

func init() { proto.RegisterFile("network.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x25, 0x04, 0x32, 0x72, 0xd1, 0xa6, 0xc9, 0xb0, 0x61, 0x0c, 0x4c, 0xc8, 0xd2, 0x34, 0xb4,
	0x07, 0x34, 0xb1, 0xa7, 0xbd, 0x80, 0x10, 0x48, 0x1b, 0xdb, 0xc4, 0x43, 0xfa, 0xf1, 0xce, 0x87,
	0x1f, 0xa2, 0x52, 0x92, 0xc6, 0x69, 0xab, 0xfe, 0xa2, 0xfe, 0x81, 0xfe, 0xc0, 0xca, 0xb1, 0x63,
	0x1c, 0x08, 0xad, 0x2a, 0xde, 0x9c, 0x73, 0xae, 0xef, 0xf1, 0x3d, 0xf7, 0xde, 0xc0, 0xfb, 0x2d,
	0x8b, 0xef, 0x83, 0xe8, 0xaa, 0x1f, 0x46, 0x41, 0x1c, 0xa0, 0x62, 0xb8, 0xa4, 0x6b, 0xa8, 0xcf,
	0x25, 0x38, 0x89, 0xd8, 0x22, 0x66, 0x1e, 0xbb, 0xb9, 0x65, 0x3c, 0x46, 0x08, 0x4a, 0xf3, 0xc5,
	0x35, 0xc3, 0x56, 0xd7, 0xea, 0xb9, 0x5e, 0x72, 0x16, 0xd8, 0x64, 0x36, 0xf5, 0x70, 0x51, 0x62,
	0xe2, 0x2c, 0xb0, 0xf3, 0x87, 0x90, 0x61, 0x5b, 0x62, 0xe2, 0x8c, 0x3e, 0x82, 0x7d, 0xe9, 0x2f,
	0x70, 0x29, 0x81, 0xc4, 0x91, 0xd6, 0x01, 0x29, 0x95, 0xff, 0x3e, 0x8f, 0x95, 0x06, 0xfd, 0xae,
	0xb5, 0xa7, 0x6c, 0xc3, 0x5e, 0xd4, 0xa6, 0x2d, 0x68, 0xaa, 0xd8, 0xdf, 0x2c, 0x1e, 0x6f, 0x36,
	0x42, 0x88, 0xa7, 0x89, 0x66, 0xd0, 0x50, 0xe4, 0x98, 0xf3, 0x60, 0xe5, 0xbf, 0x52, 0x07, 0x81,
	0xca, 0x05, 0x67, 0xd1, 0x56, 0xe0, 0xb2, 0x16, 0xfd, 0x4d, 0xff, 0x02, 0x4e, 0xdf, 0xe4, 0x9f,
	0x9a, 0xeb, 0xd1, 0x82, 0x77, 0x2a, 0xd9, 0x49, 0x7e, 0xb6, 0xc1, 0x95, 0xcd, 0x59, 0x8f, 0x63,
	0xe5, 0xea, 0x0e, 0x40, 0x3f, 0xa0, 0xa6, 0xab, 0x5e, 0xa7, 0xda, 0x1c, 0x97, 0xbb, 0x76, 0xcf,
	0xf5, 0xf2, 0xa8, 0xb4, 0x3f, 0xce, 0xae, 0x3f, 0x43, 0xf8, 0xb4, 0x37, 0x05, 0x3c, 0x0c, 0xb6,
	0x9c, 0xa1, 0xaf, 0xba, 0x82, 0xe4, 0xe5, 0xd5, 0x41, 0xb5, 0x1f, 0x2e, 0xfb, 0x0a, 0xf2, 0x52,
	0x8e, 0x0e, 0xa1, 0x96, 0xe9, 0xaf, 0xba, 0xfd, 0x0d, 0x2a, 0x0a, 0xe6, 0xd8, 0xea, 0xda, 0xfb,
	0xd7, 0x35, 0x69, 0xe8, 0xa7, 0x93, 0xf0, 0x36, 0xfd, 0x01, 0x90, 0xbc, 0xe9, 0x50, 0x49, 0xea,
	0x50, 0x4e, 0x80, 0xe4, 0x0d, 0xae, 0x27, 0x3f, 0x28, 0xd1, 0x9d, 0x36, 0x86, 0x46, 0xde, 0x30,
	0xa6, 0xcd, 0x9c, 0x02, 0x49, 0x0e, 0x9e, 0x6c, 0xf8, 0xa0, 0xd8, 0x33, 0x16, 0xdd, 0xf9, 0x2b,
	0x86, 0x46, 0xe0, 0x48, 0xe3, 0x10, 0x36, 0xde, 0x97, 0xd9, 0x28, 0xd2, 0xcc, 0x61, 0x94, 0x5c,
	0x01, 0xfd, 0x82, 0x92, 0x70, 0x0e, 0x7d, 0x36, 0x82, 0x8c, 0x55, 0x21, 0x8d, 0x03, 0x5c, 0x5f,
	0x1d, 0x81, 0x23, 0x4d, 0xcb, 0x68, 0x67, 0x36, 0x2a, 0xa3, 0x9d, 0x75, 0x98, 0x16, 0xd0, 0x1c,
	0xaa, 0x86, 0x6b, 0xa8, 0x63, 0xc4, 0x1e, 0xee, 0x1a, 0xf9, 0x72, 0x8c, 0xd6, 0xf9, 0xfe, 0x80,
	0xab, 0x1d, 0x45, 0x2d, 0x23, 0x7c, 0x7f, 0x39, 0x49, 0x3b, 0x9f, 0xd4, 0x99, 0xfe, 0x01, 0xec,
	0xfc, 0x47, 0x66, 0xf4, 0xc1, 0x72, 0x92, 0xce, 0x11, 0x36, 0x4d, 0xb6, 0x74, 0x92, 0x9f, 0xde,
	0xcf, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xb3, 0x16, 0xc9, 0x05, 0x05, 0x00, 0x00,
}
