// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: book/v1/book_controller.proto

package bookv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "grpc-boot-starter/api/protogen/google/api"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_book_v1_book_controller_proto protoreflect.FileDescriptor

var file_book_v1_book_controller_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x83, 0x07, 0x0a, 0x0e, 0x42,
	0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0xa9, 0x01,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62, 0x92, 0x41, 0x4b, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x61, 0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x20, 0x76, 0x69, 0x61, 0x20, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x12, 0xae, 0x01, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x67, 0x92, 0x41, 0x4b, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20,
	0x62, 0x6f, 0x6f, 0x6b, 0x20, 0x76, 0x69, 0x61, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x2a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xab, 0x01, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x64, 0x92, 0x41, 0x4b, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x61, 0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61,
	0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x20, 0x76, 0x69, 0x61, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8f, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x92, 0x41, 0x38, 0x0a, 0x0e, 0x42, 0x6f,
	0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0a, 0x47, 0x65,
	0x74, 0x20, 0x61, 0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x11, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20,
	0x62, 0x6f, 0x6f, 0x6b, 0x20, 0x76, 0x69, 0x61, 0x20, 0x69, 0x64, 0x2a, 0x07, 0x67, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xd3, 0x01, 0x0a, 0x0a, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x8b, 0x01, 0x92, 0x41, 0x53, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x20,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x20, 0x76, 0x69, 0x61, 0x20, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x1a, 0x1c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x20, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x20,
	0x76, 0x69, 0x61, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x2a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2f, 0x12, 0x2d, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x7d, 0x2f, 0x7b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x7d,
	0x42, 0x8e, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31,
	0x42, 0x13, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x6f,
	0x6f, 0x74, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x3b,
	0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x42,
	0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x13, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_book_v1_book_controller_proto_goTypes = []any{
	(*CreateBookRequest)(nil),  // 0: book.v1.CreateBookRequest
	(*UpdateBookRequest)(nil),  // 1: book.v1.UpdateBookRequest
	(*DeleteBookRequest)(nil),  // 2: book.v1.DeleteBookRequest
	(*GetBookRequest)(nil),     // 3: book.v1.GetBookRequest
	(*QueryBooksRequest)(nil),  // 4: book.v1.QueryBooksRequest
	(*CreateBookResponse)(nil), // 5: book.v1.CreateBookResponse
	(*UpdateBookResponse)(nil), // 6: book.v1.UpdateBookResponse
	(*DeleteBookResponse)(nil), // 7: book.v1.DeleteBookResponse
	(*GetBookResponse)(nil),    // 8: book.v1.GetBookResponse
	(*QueryBooksResponse)(nil), // 9: book.v1.QueryBooksResponse
}
var file_book_v1_book_controller_proto_depIdxs = []int32{
	0, // 0: book.v1.BookController.CreateBook:input_type -> book.v1.CreateBookRequest
	1, // 1: book.v1.BookController.UpdateBook:input_type -> book.v1.UpdateBookRequest
	2, // 2: book.v1.BookController.DeleteBook:input_type -> book.v1.DeleteBookRequest
	3, // 3: book.v1.BookController.GetBook:input_type -> book.v1.GetBookRequest
	4, // 4: book.v1.BookController.QueryBooks:input_type -> book.v1.QueryBooksRequest
	5, // 5: book.v1.BookController.CreateBook:output_type -> book.v1.CreateBookResponse
	6, // 6: book.v1.BookController.UpdateBook:output_type -> book.v1.UpdateBookResponse
	7, // 7: book.v1.BookController.DeleteBook:output_type -> book.v1.DeleteBookResponse
	8, // 8: book.v1.BookController.GetBook:output_type -> book.v1.GetBookResponse
	9, // 9: book.v1.BookController.QueryBooks:output_type -> book.v1.QueryBooksResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_book_v1_book_controller_proto_init() }
func file_book_v1_book_controller_proto_init() {
	if File_book_v1_book_controller_proto != nil {
		return
	}
	file_book_v1_book_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_book_v1_book_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_v1_book_controller_proto_goTypes,
		DependencyIndexes: file_book_v1_book_controller_proto_depIdxs,
	}.Build()
	File_book_v1_book_controller_proto = out.File
	file_book_v1_book_controller_proto_rawDesc = nil
	file_book_v1_book_controller_proto_goTypes = nil
	file_book_v1_book_controller_proto_depIdxs = nil
}