package controller

import (
	"context"
	protogen "grpc-boot-starter/apis/protov1"
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
	protogen.UnimplementedBookControllerServiceServer
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
		err = db.RecoverErrorHandle(ctx, r)
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
		err = db.RecoverErrorHandle(ctx, r)
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
