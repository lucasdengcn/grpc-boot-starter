package controller

import (
	"context"
	"grpc-boot-starter/apis/protogen"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/service"
)

func NewBookControllerImpl(bookService *service.BookService) *BookControllerImpl {
	return &BookControllerImpl{
		bookService: bookService,
	}
}

type BookControllerImpl struct {
	protogen.UnimplementedBookControllerServer
	ControllerBase
	bookService *service.BookService
}

// GetBook get book detail via id
// to handle panic properly, return named values.
func (s *BookControllerImpl) GetBook(ctx context.Context, in *protogen.BookGetInput) (resp *protogen.BookInfo, err error) {
	// validate input
	// verify ACL if pass then continue
	// call service
	resp = s.bookService.GetBook(ctx, in)
	return
}

// CreateBook create a book via input detail
// to handle panic properly, return named values.
func (s *BookControllerImpl) CreateBook(ctx context.Context, in *protogen.BookCreateInput) (bookInfo *protogen.BookInfo, err error) {
	// validate input
	// verify ACL if pass then continue
	// start Tx
	ctx = db.BeginTx(ctx)
	defer func() {
		r := recover()
		err = s.deferTxCallback(ctx, r)
	}()
	// call service A, B, C etc.
	bookInfo = s.bookService.CreateBook(ctx, in)
	return
}

// UpdateBook update a book via input detail, start db tx
// to handle panic properly, return named values.
func (s *BookControllerImpl) UpdateBook(ctx context.Context, in *protogen.BookUpdateInput) (bookInfo *protogen.BookInfo, err error) {
	// validate input
	// verify ACL if pass then continue
	// start Tx
	ctx = db.BeginTx(ctx)
	defer func() {
		r := recover()
		err = s.deferTxCallback(ctx, r)
	}()
	// call service A, B, C etc.
	bookInfo = s.bookService.UpdateBook(ctx, in)
	return
}

// DeleteBook delete a book via input, start db tx
// to handle panic properly, return named values.
func (s *BookControllerImpl) DeleteBook(ctx context.Context, in *protogen.BookDeleteInput) (resp *protogen.BookDeleteResponse, err error) {
	// validate input
	// verify ACL if pass then continue
	// start Tx
	ctx = db.BeginTx(ctx)
	defer func() {
		r := recover()
		err = s.deferTxCallback(ctx, r)
	}()
	// call service A, B, C etc.
	resp = s.bookService.DeleteBook(ctx, in)
	return
}

// QueryBooks query books via input criteria.
// to handle panic properly, return named values.
func (s *BookControllerImpl) QueryBooks(ctx context.Context, in *protogen.BookQueryInput) (resp *protogen.BookInfoListResponse, err error) {
	// validate input
	// verify ACL if pass then continue
	resp = s.bookService.QueryBooks(ctx, in)
	return
}
