// proto file compile
// protoc -I ./protos/Qdaptor_grpc ./protos/Qdaptor_grpc/Qdaptor_grpc.proto --go_out=./protos --go-grpc_out=./protos

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: Qdaptor_grpc.proto

package Qdaptor_grpc

import (
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

type TransactionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId  string `protobuf:"bytes,1,opt,name=callId,proto3" json:"callId,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TransactionMessage) Reset() {
	*x = TransactionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qdaptor_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionMessage) ProtoMessage() {}

func (x *TransactionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Qdaptor_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionMessage.ProtoReflect.Descriptor instead.
func (*TransactionMessage) Descriptor() ([]byte, []int) {
	return file_Qdaptor_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionMessage) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

func (x *TransactionMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_Qdaptor_grpc_proto protoreflect.FileDescriptor

var file_Qdaptor_grpc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x51, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x51, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x22, 0x46, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc3, 0x01, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x10, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x2e, 0x51, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x20, 0x2e, 0x51, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x43, 0x61, 0x6c, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x51, 0x64, 0x61,
	0x70, 0x74, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x20, 0x2e, 0x51,
	0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x51, 0x64, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Qdaptor_grpc_proto_rawDescOnce sync.Once
	file_Qdaptor_grpc_proto_rawDescData = file_Qdaptor_grpc_proto_rawDesc
)

func file_Qdaptor_grpc_proto_rawDescGZIP() []byte {
	file_Qdaptor_grpc_proto_rawDescOnce.Do(func() {
		file_Qdaptor_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qdaptor_grpc_proto_rawDescData)
	})
	return file_Qdaptor_grpc_proto_rawDescData
}

var file_Qdaptor_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Qdaptor_grpc_proto_goTypes = []interface{}{
	(*TransactionMessage)(nil), // 0: Qdaptor_grpc.TransactionMessage
}
var file_Qdaptor_grpc_proto_depIdxs = []int32{
	0, // 0: Qdaptor_grpc.Transaction.HelloTransaction:input_type -> Qdaptor_grpc.TransactionMessage
	0, // 1: Qdaptor_grpc.Transaction.RefCallTransaction:input_type -> Qdaptor_grpc.TransactionMessage
	0, // 2: Qdaptor_grpc.Transaction.HelloTransaction:output_type -> Qdaptor_grpc.TransactionMessage
	0, // 3: Qdaptor_grpc.Transaction.RefCallTransaction:output_type -> Qdaptor_grpc.TransactionMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Qdaptor_grpc_proto_init() }
func file_Qdaptor_grpc_proto_init() {
	if File_Qdaptor_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qdaptor_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionMessage); i {
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
			RawDescriptor: file_Qdaptor_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Qdaptor_grpc_proto_goTypes,
		DependencyIndexes: file_Qdaptor_grpc_proto_depIdxs,
		MessageInfos:      file_Qdaptor_grpc_proto_msgTypes,
	}.Build()
	File_Qdaptor_grpc_proto = out.File
	file_Qdaptor_grpc_proto_rawDesc = nil
	file_Qdaptor_grpc_proto_goTypes = nil
	file_Qdaptor_grpc_proto_depIdxs = nil
}