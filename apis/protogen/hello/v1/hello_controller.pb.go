// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: hello/v1/hello_controller.proto

package hellov1

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

var File_hello_v1_hello_controller_proto protoreflect.FileDescriptor

var file_hello_v1_hello_controller_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x54, 0x0a, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x97, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x30, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x6f, 0x6f, 0x74, 0x2d, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x14, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_hello_v1_hello_controller_proto_goTypes = []any{
	(*SayHelloRequest)(nil),  // 0: hello.v1.SayHelloRequest
	(*SayHelloResponse)(nil), // 1: hello.v1.SayHelloResponse
}
var file_hello_v1_hello_controller_proto_depIdxs = []int32{
	0, // 0: hello.v1.HelloController.SayHello:input_type -> hello.v1.SayHelloRequest
	1, // 1: hello.v1.HelloController.SayHello:output_type -> hello.v1.SayHelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_v1_hello_controller_proto_init() }
func file_hello_v1_hello_controller_proto_init() {
	if File_hello_v1_hello_controller_proto != nil {
		return
	}
	file_hello_v1_hello_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_v1_hello_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_v1_hello_controller_proto_goTypes,
		DependencyIndexes: file_hello_v1_hello_controller_proto_depIdxs,
	}.Build()
	File_hello_v1_hello_controller_proto = out.File
	file_hello_v1_hello_controller_proto_rawDesc = nil
	file_hello_v1_hello_controller_proto_goTypes = nil
	file_hello_v1_hello_controller_proto_depIdxs = nil
}
