// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: modules/comment/pb/rpc.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_modules_comment_pb_rpc_proto protoreflect.FileDescriptor

var file_modules_comment_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x20, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x3d, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x34, 0x0a,
	0x07, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x54, 0x48, 0x55, 0x2d, 0x4c, 0x53, 0x41, 0x4c, 0x41, 0x42, 0x2f, 0x4e, 0x54,
	0x48, 0x55, 0x2d, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2d, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_modules_comment_pb_rpc_proto_goTypes = []interface{}{
	(*HealthzRequest)(nil),  // 0: pb.HealthzRequest
	(*HealthzResponse)(nil), // 1: pb.HealthzResponse
}
var file_modules_comment_pb_rpc_proto_depIdxs = []int32{
	0, // 0: pb.Video.Healthz:input_type -> pb.HealthzRequest
	1, // 1: pb.Video.Healthz:output_type -> pb.HealthzResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_comment_pb_rpc_proto_init() }
func file_modules_comment_pb_rpc_proto_init() {
	if File_modules_comment_pb_rpc_proto != nil {
		return
	}
	file_modules_comment_pb_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_modules_comment_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_comment_pb_rpc_proto_goTypes,
		DependencyIndexes: file_modules_comment_pb_rpc_proto_depIdxs,
	}.Build()
	File_modules_comment_pb_rpc_proto = out.File
	file_modules_comment_pb_rpc_proto_rawDesc = nil
	file_modules_comment_pb_rpc_proto_goTypes = nil
	file_modules_comment_pb_rpc_proto_depIdxs = nil
}