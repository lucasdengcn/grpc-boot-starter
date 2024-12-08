package service

import (
	"context"
	"errors"
	protogen "grpc-boot-starter/apis/protogen/book/v1"
	"grpc-boot-starter/core/exception"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/core/security"
	"grpc-boot-starter/persistence/entity"
	"grpc-boot-starter/persistence/repository"
	"sync"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

var (
	instanceBookService *BookService
	onceBookService     sync.Once
)

// naming convention: err$module$action$code
// value convention: ${module}_${action}_${code}
var (
	errBookGet500    = "BOOK_GET_500"
	errBookCreate500 = "BOOK_CREATE_500"
	errBookDelete500 = "BOOK_DELETE_500"
	errBookQuery500  = "BOOK_QUERY_500"
	errBookUpdate500 = "BOOK_UPDATE_500"
	errBookUpdate404 = "BOOK_UPDATE_404"
)

func NewBookService(bookRepository *repository.BookRepository) *BookService {
	onceBookService.Do(func() {
		instanceBookService = &BookService{
			bookRepository: bookRepository,
		}
	})
	return instanceBookService
}

type BookService struct {
	bookRepository *repository.BookRepository
}

// mapping book entity to book info response
func (s *BookService) mapToBookInfo(book *entity.Book) *protogen.BookInfo {
	return &protogen.BookInfo{
		Id:          uint32(book.ID),
		Title:       book.Title,
		Description: book.Description,
		Author: &protogen.Author{
			Name: book.Author,
		},
		Amount:     book.Amount,
		Price:      book.Price,
		Category:   *protogen.BookCategory.Enum(protogen.BookCategory(book.Category)),
		Status:     *protogen.BookStatus.Enum(protogen.BookStatus(book.Status)),
		CreateTime: timestamppb.New(book.CreatedAt),
		UpdateTime: timestamppb.New(book.UpdatedAt),
		DeleteTime: timestamppb.New(book.DeletedAt.Time),
	}
}

// GetBook to find book's detail
// panic if repository error
func (s *BookService) GetBook(ctx context.Context, in *protogen.GetBookRequest) *protogen.BookInfo {
	//
	principle := security.CurrentUser(ctx)
	logging.Debug(ctx).Msgf("current user: %v", principle.GetID())
	//
	book, err := s.bookRepository.FindBook(ctx, in.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(exception.NewEntityNotFoundError(ctx, in.Id, "Book Not Found"))
		}
		logging.Error(ctx).Err(err).Msgf("FindBook Error. id:%v", in.Id)
		panic(exception.NewRepositoryError(ctx, errBookGet500, err.Error()))
	}
	return s.mapToBookInfo(book)
}

// CreateBook to create book via input
// panic if repository error
func (s *BookService) CreateBook(ctx context.Context, in *protogen.CreateBookRequest) *protogen.BookInfo {
	logging.Info(ctx).Msgf("CreateBook criteria: %v", in)
	book := &entity.Book{
		Title:       in.Title,
		Description: in.Description,
		Author:      in.Author.Name,
		Price:       in.Price,
		Category:    int32(in.Category),
		Amount:      in.Amount,
		Status:      int32(protogen.BookStatus_BOOK_STATUS_CREATED),
	}
	//
	book, err := s.bookRepository.Create(ctx, book)
	if err != nil {
		logging.Error(ctx).Err(err).Msgf("CreateBook Error. %v", in)
		panic(exception.NewRepositoryError(ctx, errBookCreate500, err.Error()))
		// return nil, exception.NewRepositoryError(ctx, errBookCreate500, err.Error())
	}
	return s.mapToBookInfo(book)
}

// UpdateBook to update book via id
// panic if repository error
func (s *BookService) UpdateBook(ctx context.Context, in *protogen.UpdateBookRequest) *protogen.BookInfo {
	logging.Info(ctx).Msgf("UpdateBook criteria: %v", in)
	book := &entity.Book{
		Model: gorm.Model{
			ID: uint(in.Id),
		},
		Title:       in.Title,
		Description: in.Description,
		Author:      in.Author.Name,
		Price:       in.Price,
		Category:    int32(in.Category),
		Amount:      in.Amount,
		Status:      int32(protogen.BookStatus_BOOK_STATUS_ACTIVE),
	}
	ok, err := s.bookRepository.Update(ctx, book)
	if err != nil {
		logging.Error(ctx).Err(err).Msgf("UpdateBook Error. %v", in)
		panic(exception.NewRepositoryError(ctx, errBookUpdate500, err.Error()))
	}
	if !ok {
		logging.Error(ctx).Msgf("UpdateBook Failed. %v", in)
		panic(exception.NewEntityNotFoundError(ctx, errBookUpdate404, "Update book but Got no updates"))
	}
	return s.mapToBookInfo(book)
}

// DeleteBook to delete book via id
// panic if repository error
func (s *BookService) DeleteBook(ctx context.Context, in *protogen.DeleteBookRequest) bool {
	logging.Info(ctx).Msgf("DeleteBook criteria: %v", in)
	ok, err := s.bookRepository.Delete(ctx, in.Id)
	if err != nil {
		logging.Error(ctx).Err(err).Msgf("DeleteBook Error. id:%v", in.Id)
		panic(exception.NewRepositoryError(ctx, errBookDelete500, err.Error()))
	}
	return ok
}

// QueryBooks to find books via input criteria
// panic if repository error
func (s *BookService) QueryBooks(ctx context.Context, in *protogen.QueryBooksRequest) *protogen.QueryBooksResponse {
	// prepare query criteria
	logging.Info(ctx).Msgf("QueryBooks criteria: %v", in)
	var status int32
	var category int32
	if in.Status != nil {
		status = int32(*in.Status)
	} else {
		status = -1
	}
	if in.Category != nil {
		category = int32(*in.Category)
	} else {
		category = -1
	}
	// counting
	count, err := s.bookRepository.CountBooks(ctx, status, category, 0)
	// query items
	var bookList []*entity.Book
	bookList, err = s.bookRepository.FindBooks(ctx, status, category, 0)
	if err != nil {
		logging.Error(ctx).Err(err).Msgf("QueryBooks Error. %v", in)
		panic(exception.NewRepositoryError(ctx, errBookQuery500, err.Error()))
	}
	// processing pagination
	totalPages := count / in.PageSize
	if count%in.PageSize > 0 {
		totalPages += 1
	}
	if bookList == nil {
		return &protogen.QueryBooksResponse{
			Books: nil,
			PageInfo: &protogen.Pagination{
				TotalItems:  count,
				TotalPages:  totalPages,
				PageSize:    in.PageSize,
				PageIndex:   in.PageIndex,
				HasNext:     false,
				HasPrevious: totalPages > 0,
			},
		}
	}
	// prepare and return
	var output = &protogen.QueryBooksResponse{
		PageInfo: &protogen.Pagination{
			TotalItems:  count,
			TotalPages:  totalPages,
			PageSize:    in.PageSize,
			PageIndex:   in.PageIndex,
			HasNext:     totalPages > in.PageIndex,
			HasPrevious: in.PageIndex > 1,
		},
	}
	for _, book := range bookList {
		output.Books = append(output.Books, s.mapToBookInfo(book))
	}
	return output
}
