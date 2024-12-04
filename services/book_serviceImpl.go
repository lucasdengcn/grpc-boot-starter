package services

import (
	"context"
	"errors"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/persistence/entity"
	"grpc-boot-starter/persistence/repository"
	"grpc-boot-starter/protogen"
	"sync"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

var (
	instanceBookService *BookServiceServerImpl
	onceBookService     sync.Once
)

func NewBookServiceServerImpl(bookRepository *repository.BookRepository) *BookServiceServerImpl {
	onceBookService.Do(func() {
		instanceBookService = &BookServiceServerImpl{
			bookRepository: bookRepository,
		}
	})
	return instanceBookService
}

type BookServiceServerImpl struct {
	protogen.UnimplementedBookServiceServer
	bookRepository *repository.BookRepository
}

func (s *BookServiceServerImpl) mapToBookInfo(book *entity.Book) *protogen.BookInfo {
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
func (s *BookServiceServerImpl) GetBook(ctx context.Context, in *protogen.BookGetInput) (*protogen.BookInfo, error) {
	book, err := s.bookRepository.FindBook(ctx, in.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logging.Error(ctx).Err(err).Msgf("FindBook Error. id:%v", in.Id)
		return nil, err
	}
	return s.mapToBookInfo(book), nil
}

// CreateBook to create book via input
func (s *BookServiceServerImpl) CreateBook(ctx context.Context, in *protogen.BookCreateInput) (*protogen.BookInfo, error) {
	book := &entity.Book{
		Title:       in.Title,
		Description: in.Description,
		Author:      in.Author.Name,
		Price:       in.Price,
		Category:    int32(in.Category),
		Amount:      in.Amount,
		Status:      int32(protogen.BookStatus_BOOK_STATUS_CREATED),
	}
	book, err := s.bookRepository.Create(ctx, book)
	if err != nil {
		logging.Error(ctx).Err(err).Msgf("CreateBook Error. %v", in)
		return nil, err
	}
	return s.mapToBookInfo(book), nil
}

func (s *BookServiceServerImpl) UpdateBook(ctx context.Context, in *protogen.BookUpdateInput) (*protogen.BookInfo, error) {
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
		return nil, err
	}
	if !ok {
		logging.Error(ctx).Msgf("UpdateBook Failed. %v", in)
		return nil, nil
	}
	return s.mapToBookInfo(book), nil
}

// DeleteBook to delete book via id
func (s *BookServiceServerImpl) DeleteBook(ctx context.Context, in *protogen.BookDeleteInput) (*protogen.BookDeleteResponse, error) {
	ok, err := s.bookRepository.Delete(ctx, in.Id)
	if err != nil {
		logging.Error(ctx).Err(err).Msgf("DeleteBook Error. id:%v", in.Id)
		return nil, err
	}
	return &protogen.BookDeleteResponse{
		Id:      in.Id,
		Success: ok,
	}, nil
}

// QueryBooks to find books via input criteria
func (s *BookServiceServerImpl) QueryBooks(ctx context.Context, in *protogen.BookQueryInput) (*protogen.BookInfoListResponse, error) {
	// prepare query criteria
	var status int32
	var category int32
	var err error
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
		return nil, err
	}
	// processing pagination
	totalPages := count / in.PageSize
	if count%totalPages > 0 {
		totalPages += 1
	}
	if bookList == nil {
		return &protogen.BookInfoListResponse{
			Books:       nil,
			TotalItems:  count,
			TotalPages:  totalPages,
			PageSize:    in.PageSize,
			PageIndex:   in.PageIndex,
			HasNext:     false,
			HasPrevious: totalPages > 0,
		}, nil
	}
	// prepare and return
	var output = &protogen.BookInfoListResponse{
		TotalItems:  count,
		TotalPages:  totalPages,
		PageSize:    in.PageSize,
		PageIndex:   in.PageIndex,
		HasNext:     totalPages > in.PageIndex,
		HasPrevious: in.PageIndex > 1,
	}
	for _, book := range bookList {
		output.Books = append(output.Books, s.mapToBookInfo(book))
	}
	return output, nil
}