// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: book_controller.proto

package protov1

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

var File_book_controller_proto protoreflect.FileDescriptor

var file_book_controller_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x1a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd7, 0x02,
	0x0a, 0x15, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x45, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x46, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x8d, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x27, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x6f, 0x6f, 0x74, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02,
	0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_book_controller_proto_goTypes = []any{
	(*BookCreateInput)(nil),      // 0: proto.v1.BookCreateInput
	(*BookUpdateInput)(nil),      // 1: proto.v1.BookUpdateInput
	(*BookDeleteInput)(nil),      // 2: proto.v1.BookDeleteInput
	(*BookGetInput)(nil),         // 3: proto.v1.BookGetInput
	(*BookQueryInput)(nil),       // 4: proto.v1.BookQueryInput
	(*BookInfo)(nil),             // 5: proto.v1.BookInfo
	(*BookDeleteResponse)(nil),   // 6: proto.v1.BookDeleteResponse
	(*BookInfoListResponse)(nil), // 7: proto.v1.BookInfoListResponse
}
var file_book_controller_proto_depIdxs = []int32{
	0, // 0: proto.v1.BookControllerService.CreateBook:input_type -> proto.v1.BookCreateInput
	1, // 1: proto.v1.BookControllerService.UpdateBook:input_type -> proto.v1.BookUpdateInput
	2, // 2: proto.v1.BookControllerService.DeleteBook:input_type -> proto.v1.BookDeleteInput
	3, // 3: proto.v1.BookControllerService.GetBook:input_type -> proto.v1.BookGetInput
	4, // 4: proto.v1.BookControllerService.QueryBooks:input_type -> proto.v1.BookQueryInput
	5, // 5: proto.v1.BookControllerService.CreateBook:output_type -> proto.v1.BookInfo
	5, // 6: proto.v1.BookControllerService.UpdateBook:output_type -> proto.v1.BookInfo
	6, // 7: proto.v1.BookControllerService.DeleteBook:output_type -> proto.v1.BookDeleteResponse
	5, // 8: proto.v1.BookControllerService.GetBook:output_type -> proto.v1.BookInfo
	7, // 9: proto.v1.BookControllerService.QueryBooks:output_type -> proto.v1.BookInfoListResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_book_controller_proto_init() }
func file_book_controller_proto_init() {
	if File_book_controller_proto != nil {
		return
	}
	file_book_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_book_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_controller_proto_goTypes,
		DependencyIndexes: file_book_controller_proto_depIdxs,
	}.Build()
	File_book_controller_proto = out.File
	file_book_controller_proto_rawDesc = nil
	file_book_controller_proto_goTypes = nil
	file_book_controller_proto_depIdxs = nil
}