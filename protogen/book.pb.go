// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: protogen/book.proto

package protogen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BookStatus int32

const (
	BookStatus_BOOK_STATUS_CREATED BookStatus = 0
	BookStatus_BOOK_STATUS_ACTIVE  BookStatus = 1
	BookStatus_BOOK_STATUS_OFFLINE BookStatus = 2
)

// Enum value maps for BookStatus.
var (
	BookStatus_name = map[int32]string{
		0: "BOOK_STATUS_CREATED",
		1: "BOOK_STATUS_ACTIVE",
		2: "BOOK_STATUS_OFFLINE",
	}
	BookStatus_value = map[string]int32{
		"BOOK_STATUS_CREATED": 0,
		"BOOK_STATUS_ACTIVE":  1,
		"BOOK_STATUS_OFFLINE": 2,
	}
)

func (x BookStatus) Enum() *BookStatus {
	p := new(BookStatus)
	*p = x
	return p
}

func (x BookStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BookStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protogen_book_proto_enumTypes[0].Descriptor()
}

func (BookStatus) Type() protoreflect.EnumType {
	return &file_protogen_book_proto_enumTypes[0]
}

func (x BookStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BookStatus.Descriptor instead.
func (BookStatus) EnumDescriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{0}
}

type BookCategory int32

const (
	BookCategory_BOOK_CATEGORY_UNKNOWN BookCategory = 0
	BookCategory_BOOK_CATEGORY_JAVA    BookCategory = 1
	BookCategory_BOOK_CATEGORY_GO      BookCategory = 2
	BookCategory_BOOK_CATEGORY_MATH    BookCategory = 3
)

// Enum value maps for BookCategory.
var (
	BookCategory_name = map[int32]string{
		0: "BOOK_CATEGORY_UNKNOWN",
		1: "BOOK_CATEGORY_JAVA",
		2: "BOOK_CATEGORY_GO",
		3: "BOOK_CATEGORY_MATH",
	}
	BookCategory_value = map[string]int32{
		"BOOK_CATEGORY_UNKNOWN": 0,
		"BOOK_CATEGORY_JAVA":    1,
		"BOOK_CATEGORY_GO":      2,
		"BOOK_CATEGORY_MATH":    3,
	}
)

func (x BookCategory) Enum() *BookCategory {
	p := new(BookCategory)
	*p = x
	return p
}

func (x BookCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BookCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_protogen_book_proto_enumTypes[1].Descriptor()
}

func (BookCategory) Type() protoreflect.EnumType {
	return &file_protogen_book_proto_enumTypes[1]
}

func (x BookCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BookCategory.Descriptor instead.
func (BookCategory) EnumDescriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{1}
}

type BookCreateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string       `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount      uint32       `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Price       float32      `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Category    BookCategory `protobuf:"varint,5,opt,name=category,proto3,enum=protogen.BookCategory" json:"category,omitempty"`
	Author      *Author      `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *BookCreateInput) Reset() {
	*x = BookCreateInput{}
	mi := &file_protogen_book_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookCreateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookCreateInput) ProtoMessage() {}

func (x *BookCreateInput) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookCreateInput.ProtoReflect.Descriptor instead.
func (*BookCreateInput) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{0}
}

func (x *BookCreateInput) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookCreateInput) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BookCreateInput) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BookCreateInput) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BookCreateInput) GetCategory() BookCategory {
	if x != nil {
		return x.Category
	}
	return BookCategory_BOOK_CATEGORY_UNKNOWN
}

func (x *BookCreateInput) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type BookUpdateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount      uint32       `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Price       float32      `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Category    BookCategory `protobuf:"varint,6,opt,name=category,proto3,enum=protogen.BookCategory" json:"category,omitempty"`
	Author      *Author      `protobuf:"bytes,7,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *BookUpdateInput) Reset() {
	*x = BookUpdateInput{}
	mi := &file_protogen_book_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookUpdateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookUpdateInput) ProtoMessage() {}

func (x *BookUpdateInput) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookUpdateInput.ProtoReflect.Descriptor instead.
func (*BookUpdateInput) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{1}
}

func (x *BookUpdateInput) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookUpdateInput) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookUpdateInput) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BookUpdateInput) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BookUpdateInput) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BookUpdateInput) GetCategory() BookCategory {
	if x != nil {
		return x.Category
	}
	return BookCategory_BOOK_CATEGORY_UNKNOWN
}

func (x *BookUpdateInput) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type BookStatusUpdateInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status BookStatus `protobuf:"varint,2,opt,name=status,proto3,enum=protogen.BookStatus" json:"status,omitempty"`
}

func (x *BookStatusUpdateInput) Reset() {
	*x = BookStatusUpdateInput{}
	mi := &file_protogen_book_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookStatusUpdateInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookStatusUpdateInput) ProtoMessage() {}

func (x *BookStatusUpdateInput) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookStatusUpdateInput.ProtoReflect.Descriptor instead.
func (*BookStatusUpdateInput) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{2}
}

func (x *BookStatusUpdateInput) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookStatusUpdateInput) GetStatus() BookStatus {
	if x != nil {
		return x.Status
	}
	return BookStatus_BOOK_STATUS_CREATED
}

type BookDeleteInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BookDeleteInput) Reset() {
	*x = BookDeleteInput{}
	mi := &file_protogen_book_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookDeleteInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookDeleteInput) ProtoMessage() {}

func (x *BookDeleteInput) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookDeleteInput.ProtoReflect.Descriptor instead.
func (*BookDeleteInput) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{3}
}

func (x *BookDeleteInput) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BookDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BookDeleteResponse) Reset() {
	*x = BookDeleteResponse{}
	mi := &file_protogen_book_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookDeleteResponse) ProtoMessage() {}

func (x *BookDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookDeleteResponse.ProtoReflect.Descriptor instead.
func (*BookDeleteResponse) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{4}
}

func (x *BookDeleteResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookDeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type BookGetInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BookGetInput) Reset() {
	*x = BookGetInput{}
	mi := &file_protogen_book_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookGetInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookGetInput) ProtoMessage() {}

func (x *BookGetInput) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookGetInput.ProtoReflect.Descriptor instead.
func (*BookGetInput) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{5}
}

func (x *BookGetInput) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BookQueryInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *BookCategory `protobuf:"varint,1,opt,name=category,proto3,enum=protogen.BookCategory,oneof" json:"category,omitempty"`
	Status   *BookStatus   `protobuf:"varint,2,opt,name=status,proto3,enum=protogen.BookStatus,oneof" json:"status,omitempty"`
	PageSize  uint32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageIndex uint32  `protobuf:"varint,4,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	CursorId  *uint32 `protobuf:"varint,5,opt,name=cursor_id,json=cursorId,proto3,oneof" json:"cursor_id,omitempty"`
}

func (x *BookQueryInput) Reset() {
	*x = BookQueryInput{}
	mi := &file_protogen_book_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookQueryInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookQueryInput) ProtoMessage() {}

func (x *BookQueryInput) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookQueryInput.ProtoReflect.Descriptor instead.
func (*BookQueryInput) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{6}
}

func (x *BookQueryInput) GetCategory() BookCategory {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return BookCategory_BOOK_CATEGORY_UNKNOWN
}

func (x *BookQueryInput) GetStatus() BookStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return BookStatus_BOOK_STATUS_CREATED
}

func (x *BookQueryInput) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BookQueryInput) GetPageIndex() uint32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *BookQueryInput) GetCursorId() uint32 {
	if x != nil && x.CursorId != nil {
		return *x.CursorId
	}
	return 0
}

type BookInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Author      *Author                `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Amount      uint32                 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Price       float32                `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	Category    BookCategory           `protobuf:"varint,7,opt,name=category,proto3,enum=protogen.BookCategory" json:"category,omitempty"`
	Status      BookStatus             `protobuf:"varint,8,opt,name=status,proto3,enum=protogen.BookStatus" json:"status,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime  *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	Deleted     bool                   `protobuf:"varint,12,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *BookInfo) Reset() {
	*x = BookInfo{}
	mi := &file_protogen_book_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookInfo) ProtoMessage() {}

func (x *BookInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookInfo.ProtoReflect.Descriptor instead.
func (*BookInfo) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{7}
}

func (x *BookInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BookInfo) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *BookInfo) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BookInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BookInfo) GetCategory() BookCategory {
	if x != nil {
		return x.Category
	}
	return BookCategory_BOOK_CATEGORY_UNKNOWN
}

func (x *BookInfo) GetStatus() BookStatus {
	if x != nil {
		return x.Status
	}
	return BookStatus_BOOK_STATUS_CREATED
}

func (x *BookInfo) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BookInfo) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *BookInfo) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *BookInfo) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type BookInfoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books       []*BookInfo `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	PageSize    uint32      `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageIndex   uint32      `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	TotalPages  uint32      `protobuf:"varint,4,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	HasNext     bool        `protobuf:"varint,5,opt,name=has_next,json=hasNext,proto3" json:"has_next,omitempty"`
	HasPrevious bool        `protobuf:"varint,6,opt,name=has_previous,json=hasPrevious,proto3" json:"has_previous,omitempty"`
	TotalItems  uint32      `protobuf:"varint,7,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
}

func (x *BookInfoListResponse) Reset() {
	*x = BookInfoListResponse{}
	mi := &file_protogen_book_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookInfoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookInfoListResponse) ProtoMessage() {}

func (x *BookInfoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protogen_book_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookInfoListResponse.ProtoReflect.Descriptor instead.
func (*BookInfoListResponse) Descriptor() ([]byte, []int) {
	return file_protogen_book_proto_rawDescGZIP(), []int{8}
}

func (x *BookInfoListResponse) GetBooks() []*BookInfo {
	if x != nil {
		return x.Books
	}
	return nil
}

func (x *BookInfoListResponse) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BookInfoListResponse) GetPageIndex() uint32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *BookInfoListResponse) GetTotalPages() uint32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *BookInfoListResponse) GetHasNext() bool {
	if x != nil {
		return x.HasNext
	}
	return false
}

func (x *BookInfoListResponse) GetHasPrevious() bool {
	if x != nil {
		return x.HasPrevious
	}
	return false
}

func (x *BookInfoListResponse) GetTotalItems() uint32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

var File_protogen_book_proto protoreflect.FileDescriptor

var file_protogen_book_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22,
	0xe5, 0x01, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x55, 0x0a, 0x15, 0x42, 0x6f, 0x6f, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x21,
	0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3e, 0x0a, 0x12, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x1e, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x80, 0x02, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x48, 0x00,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x09,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x02, 0x52, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x22, 0xdd, 0x03, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x22, 0xfc, 0x01, 0x0a, 0x14, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x2a, 0x56, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x4f, 0x4f, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x4f,
	0x4f, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x4f, 0x4f, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x2a, 0x6f, 0x0a, 0x0c, 0x42,
	0x6f, 0x6f, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x15, 0x42,
	0x4f, 0x4f, 0x4b, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x4f, 0x4f, 0x4b, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4a, 0x41, 0x56, 0x41, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x42, 0x4f, 0x4f, 0x4b, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f,
	0x47, 0x4f, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x4f, 0x4f, 0x4b, 0x5f, 0x43, 0x41, 0x54,
	0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4d, 0x41, 0x54, 0x48, 0x10, 0x03, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protogen_book_proto_rawDescOnce sync.Once
	file_protogen_book_proto_rawDescData = file_protogen_book_proto_rawDesc
)

func file_protogen_book_proto_rawDescGZIP() []byte {
	file_protogen_book_proto_rawDescOnce.Do(func() {
		file_protogen_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_protogen_book_proto_rawDescData)
	})
	return file_protogen_book_proto_rawDescData
}

var file_protogen_book_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protogen_book_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protogen_book_proto_goTypes = []any{
	(BookStatus)(0),               // 0: protogen.BookStatus
	(BookCategory)(0),             // 1: protogen.BookCategory
	(*BookCreateInput)(nil),       // 2: protogen.BookCreateInput
	(*BookUpdateInput)(nil),       // 3: protogen.BookUpdateInput
	(*BookStatusUpdateInput)(nil), // 4: protogen.BookStatusUpdateInput
	(*BookDeleteInput)(nil),       // 5: protogen.BookDeleteInput
	(*BookDeleteResponse)(nil),    // 6: protogen.BookDeleteResponse
	(*BookGetInput)(nil),          // 7: protogen.BookGetInput
	(*BookQueryInput)(nil),        // 8: protogen.BookQueryInput
	(*BookInfo)(nil),              // 9: protogen.BookInfo
	(*BookInfoListResponse)(nil),  // 10: protogen.BookInfoListResponse
	(*Author)(nil),                // 11: protogen.Author
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_protogen_book_proto_depIdxs = []int32{
	1,  // 0: protogen.BookCreateInput.category:type_name -> protogen.BookCategory
	11, // 1: protogen.BookCreateInput.author:type_name -> protogen.Author
	1,  // 2: protogen.BookUpdateInput.category:type_name -> protogen.BookCategory
	11, // 3: protogen.BookUpdateInput.author:type_name -> protogen.Author
	0,  // 4: protogen.BookStatusUpdateInput.status:type_name -> protogen.BookStatus
	1,  // 5: protogen.BookQueryInput.category:type_name -> protogen.BookCategory
	0,  // 6: protogen.BookQueryInput.status:type_name -> protogen.BookStatus
	11, // 7: protogen.BookInfo.author:type_name -> protogen.Author
	1,  // 8: protogen.BookInfo.category:type_name -> protogen.BookCategory
	0,  // 9: protogen.BookInfo.status:type_name -> protogen.BookStatus
	12, // 10: protogen.BookInfo.create_time:type_name -> google.protobuf.Timestamp
	12, // 11: protogen.BookInfo.update_time:type_name -> google.protobuf.Timestamp
	12, // 12: protogen.BookInfo.delete_time:type_name -> google.protobuf.Timestamp
	9,  // 13: protogen.BookInfoListResponse.books:type_name -> protogen.BookInfo
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_protogen_book_proto_init() }
func file_protogen_book_proto_init() {
	if File_protogen_book_proto != nil {
		return
	}
	file_protogen_author_proto_init()
	file_protogen_book_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protogen_book_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protogen_book_proto_goTypes,
		DependencyIndexes: file_protogen_book_proto_depIdxs,
		EnumInfos:         file_protogen_book_proto_enumTypes,
		MessageInfos:      file_protogen_book_proto_msgTypes,
	}.Build()
	File_protogen_book_proto = out.File
	file_protogen_book_proto_rawDesc = nil
	file_protogen_book_proto_goTypes = nil
	file_protogen_book_proto_depIdxs = nil
}
