// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: message.proto

package Interceptor

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RequestArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arg1 int32 `protobuf:"varint,1,opt,name=arg1,proto3" json:"arg1,omitempty"`
	Arg2 int32 `protobuf:"varint,2,opt,name=arg2,proto3" json:"arg2,omitempty"`
}

func (x *RequestArgs) Reset() {
	*x = RequestArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestArgs) ProtoMessage() {}

func (x *RequestArgs) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestArgs.ProtoReflect.Descriptor instead.
func (*RequestArgs) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *RequestArgs) GetArg1() int32 {
	if x != nil {
		return x.Arg1
	}
	return 0
}

func (x *RequestArgs) GetArg2() int32 {
	if x != nil {
		return x.Arg2
	}
	return 0
}

type ResponseArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ResponseArgs) Reset() {
	*x = ResponseArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseArgs) ProtoMessage() {}

func (x *ResponseArgs) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseArgs.ProtoReflect.Descriptor instead.
func (*ResponseArgs) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseArgs) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseArgs) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x0b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61, 0x72, 0x67, 0x31, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61,
	0x72, 0x67, 0x32, 0x22, 0x34, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41,
	0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x51, 0x0a, 0x0b, 0x4d, 0x61, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a,
	0x19, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x72, 0x67, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_proto_goTypes = []interface{}{
	(*RequestArgs)(nil),  // 0: Interceptor.RequestArgs
	(*ResponseArgs)(nil), // 1: Interceptor.ResponseArgs
}
var file_message_proto_depIdxs = []int32{
	0, // 0: Interceptor.MathService.AddMethod:input_type -> Interceptor.RequestArgs
	1, // 1: Interceptor.MathService.AddMethod:output_type -> Interceptor.ResponseArgs
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestArgs); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseArgs); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathServiceClient interface {
	AddMethod(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseArgs, error)
}

type mathServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMathServiceClient(cc grpc.ClientConnInterface) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) AddMethod(ctx context.Context, in *RequestArgs, opts ...grpc.CallOption) (*ResponseArgs, error) {
	out := new(ResponseArgs)
	err := c.cc.Invoke(ctx, "/Interceptor.MathService/AddMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServiceServer is the server API for MathService service.
type MathServiceServer interface {
	AddMethod(context.Context, *RequestArgs) (*ResponseArgs, error)
}

// UnimplementedMathServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMathServiceServer struct {
}

func (*UnimplementedMathServiceServer) AddMethod(context.Context, *RequestArgs) (*ResponseArgs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMethod not implemented")
}

func RegisterMathServiceServer(s *grpc.Server, srv MathServiceServer) {
	s.RegisterService(&_MathService_serviceDesc, srv)
}

func _MathService_AddMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).AddMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Interceptor.MathService/AddMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).AddMethod(ctx, req.(*RequestArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Interceptor.MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMethod",
			Handler:    _MathService_AddMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
