package controller

import (
	"context"
	pbbook "grpc-boot-starter/apis/protogen/book/v1"
	"grpc-boot-starter/core/exception"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/service"

	"github.com/bufbuild/protovalidate-go"
)

func NewBookControllerImpl(bookService *service.BookService) *BookControllerImpl {
	return &BookControllerImpl{
		bookService: bookService,
	}
}

type BookControllerImpl struct {
	pbbook.UnimplementedBookControllerServiceServer
	bookService *service.BookService
}

// GetBook get book detail via id
// to handle panic properly, return named values.
func (s *BookControllerImpl) GetBook(ctx context.Context, in *pbbook.GetBookRequest) (resp *pbbook.GetBookResponse, err error) {
	// validate input
	// verify ACL if pass then continue
	// call service
	bookInfo := s.bookService.GetBook(ctx, in)
	resp = &pbbook.GetBookResponse{
		Book: bookInfo,
	}
	return
}

// CreateBook create a book via input detail
// to handle panic properly, return named values.
func (s *BookControllerImpl) CreateBook(ctx context.Context, in *pbbook.CreateBookRequest) (resp *pbbook.CreateBookResponse, err error) {
	// validate input
	if err0 := protovalidate.Validate(in); err0 != nil {
		logging.Error(ctx).Err(err0).Msgf("BookCreateInput invalid. %v", in)
		return nil, exception.NewValidationErrorOnFailed(ctx, "BOOK_CREATE_INPUTS_400", err0)
	}
	// verify ACL if pass then continue
	// start Tx
	ctx = db.BeginTx(ctx)
	defer func() {
		r := recover()
		err = db.RecoverErrorHandle(ctx, r)
	}()
	// call service A, B, C etc.
	bookInfo := s.bookService.CreateBook(ctx, in)
	resp = &pbbook.CreateBookResponse{
		Book:    bookInfo,
		Success: true,
	}
	return
}

// UpdateBook update a book via input detail, start db tx
// to handle panic properly, return named values.
func (s *BookControllerImpl) UpdateBook(ctx context.Context, in *pbbook.UpdateBookRequest) (resp *pbbook.UpdateBookResponse, err error) {
	// validate input
	// verify ACL if pass then continue
	// start Tx
	ctx = db.BeginTx(ctx)
	defer func() {
		r := recover()
		err = db.RecoverErrorHandle(ctx, r)
	}()
	// call service A, B, C etc.
	bookInfo := s.bookService.UpdateBook(ctx, in)
	resp = &pbbook.UpdateBookResponse{
		Book:    bookInfo,
		Success: true,
	}
	return
}

// DeleteBook delete a book via input, start db tx
// to handle panic properly, return named values.
func (s *BookControllerImpl) DeleteBook(ctx context.Context, in *pbbook.DeleteBookRequest) (resp *pbbook.DeleteBookResponse, err error) {
	// validate input
	// verify ACL if pass then continue
	// start Tx
	ctx = db.BeginTx(ctx)
	defer func() {
		r := recover()
		err = db.RecoverErrorHandle(ctx, r)
	}()
	// call service A, B, C etc.
	ok := s.bookService.DeleteBook(ctx, in)
	resp = &pbbook.DeleteBookResponse{
		Id:      in.Id,
		Success: ok,
	}
	return
}

// QueryBooks query books via input criteria.
// to handle panic properly, return named values.
func (s *BookControllerImpl) QueryBooks(ctx context.Context, in *pbbook.QueryBooksRequest) (resp *pbbook.QueryBooksResponse, err error) {
	// validate input
	// verify ACL if pass then continue
	resp = s.bookService.QueryBooks(ctx, in)
	return
}
