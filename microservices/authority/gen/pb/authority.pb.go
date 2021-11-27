// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: authority.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdToken      string `protobuf:"bytes,1,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authority_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_authority_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_authority_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type CreateTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *CreateTokenReq) Reset() {
	*x = CreateTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authority_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenReq) ProtoMessage() {}

func (x *CreateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_authority_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenReq.ProtoReflect.Descriptor instead.
func (*CreateTokenReq) Descriptor() ([]byte, []int) {
	return file_authority_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTokenReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type ListPublicKeysReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPublicKeysReq) Reset() {
	*x = ListPublicKeysReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authority_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPublicKeysReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublicKeysReq) ProtoMessage() {}

func (x *ListPublicKeysReq) ProtoReflect() protoreflect.Message {
	mi := &file_authority_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublicKeysReq.ProtoReflect.Descriptor instead.
func (*ListPublicKeysReq) Descriptor() ([]byte, []int) {
	return file_authority_proto_rawDescGZIP(), []int{2}
}

type CreateTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateTokenRes) Reset() {
	*x = CreateTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authority_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRes) ProtoMessage() {}

func (x *CreateTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_authority_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRes.ProtoReflect.Descriptor instead.
func (*CreateTokenRes) Descriptor() ([]byte, []int) {
	return file_authority_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTokenRes) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type ListPublicKeysRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwks string `protobuf:"bytes,1,opt,name=jwks,proto3" json:"jwks,omitempty"`
}

func (x *ListPublicKeysRes) Reset() {
	*x = ListPublicKeysRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authority_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPublicKeysRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublicKeysRes) ProtoMessage() {}

func (x *ListPublicKeysRes) ProtoReflect() protoreflect.Message {
	mi := &file_authority_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublicKeysRes.ProtoReflect.Descriptor instead.
func (*ListPublicKeysRes) Descriptor() ([]byte, []int) {
	return file_authority_proto_rawDescGZIP(), []int{4}
}

func (x *ListPublicKeysRes) GetJwks() string {
	if x != nil {
		return x.Jwks
	}
	return ""
}

var File_authority_proto protoreflect.FileDescriptor

var file_authority_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x22,
	0x38, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x27, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x77,
	0x6b, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authority_proto_rawDescOnce sync.Once
	file_authority_proto_rawDescData = file_authority_proto_rawDesc
)

func file_authority_proto_rawDescGZIP() []byte {
	file_authority_proto_rawDescOnce.Do(func() {
		file_authority_proto_rawDescData = protoimpl.X.CompressGZIP(file_authority_proto_rawDescData)
	})
	return file_authority_proto_rawDescData
}

var file_authority_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_authority_proto_goTypes = []interface{}{
	(*Token)(nil),             // 0: authority.Token
	(*CreateTokenReq)(nil),    // 1: authority.CreateTokenReq
	(*ListPublicKeysReq)(nil), // 2: authority.ListPublicKeysReq
	(*CreateTokenRes)(nil),    // 3: authority.CreateTokenRes
	(*ListPublicKeysRes)(nil), // 4: authority.ListPublicKeysRes
}
var file_authority_proto_depIdxs = []int32{
	0, // 0: authority.CreateTokenRes.token:type_name -> authority.Token
	1, // 1: authority.AuthorityService.CreateToken:input_type -> authority.CreateTokenReq
	2, // 2: authority.AuthorityService.ListPublicKeys:input_type -> authority.ListPublicKeysReq
	3, // 3: authority.AuthorityService.CreateToken:output_type -> authority.CreateTokenRes
	4, // 4: authority.AuthorityService.ListPublicKeys:output_type -> authority.ListPublicKeysRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_authority_proto_init() }
func file_authority_proto_init() {
	if File_authority_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authority_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authority_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authority_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPublicKeysReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authority_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authority_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPublicKeysRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authority_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authority_proto_goTypes,
		DependencyIndexes: file_authority_proto_depIdxs,
		MessageInfos:      file_authority_proto_msgTypes,
	}.Build()
	File_authority_proto = out.File
	file_authority_proto_rawDesc = nil
	file_authority_proto_goTypes = nil
	file_authority_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorityServiceClient is the client API for AuthorityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorityServiceClient interface {
	CreateToken(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*CreateTokenRes, error)
	ListPublicKeys(ctx context.Context, in *ListPublicKeysReq, opts ...grpc.CallOption) (*ListPublicKeysRes, error)
}

type authorityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorityServiceClient(cc grpc.ClientConnInterface) AuthorityServiceClient {
	return &authorityServiceClient{cc}
}

func (c *authorityServiceClient) CreateToken(ctx context.Context, in *CreateTokenReq, opts ...grpc.CallOption) (*CreateTokenRes, error) {
	out := new(CreateTokenRes)
	err := c.cc.Invoke(ctx, "/authority.AuthorityService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityServiceClient) ListPublicKeys(ctx context.Context, in *ListPublicKeysReq, opts ...grpc.CallOption) (*ListPublicKeysRes, error) {
	out := new(ListPublicKeysRes)
	err := c.cc.Invoke(ctx, "/authority.AuthorityService/ListPublicKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorityServiceServer is the server API for AuthorityService service.
type AuthorityServiceServer interface {
	CreateToken(context.Context, *CreateTokenReq) (*CreateTokenRes, error)
	ListPublicKeys(context.Context, *ListPublicKeysReq) (*ListPublicKeysRes, error)
}

// UnimplementedAuthorityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorityServiceServer struct {
}

func (*UnimplementedAuthorityServiceServer) CreateToken(context.Context, *CreateTokenReq) (*CreateTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (*UnimplementedAuthorityServiceServer) ListPublicKeys(context.Context, *ListPublicKeysReq) (*ListPublicKeysRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublicKeys not implemented")
}

func RegisterAuthorityServiceServer(s *grpc.Server, srv AuthorityServiceServer) {
	s.RegisterService(&_AuthorityService_serviceDesc, srv)
}

func _AuthorityService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authority.AuthorityService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).CreateToken(ctx, req.(*CreateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityService_ListPublicKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublicKeysReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).ListPublicKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authority.AuthorityService/ListPublicKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).ListPublicKeys(ctx, req.(*ListPublicKeysReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authority.AuthorityService",
	HandlerType: (*AuthorityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _AuthorityService_CreateToken_Handler,
		},
		{
			MethodName: "ListPublicKeys",
			Handler:    _AuthorityService_ListPublicKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authority.proto",
}