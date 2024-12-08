package service

import (
	"context"
	protogen "grpc-boot-starter/apis/protogen/book/v1"
	"grpc-boot-starter/core/config"
	"grpc-boot-starter/core/security"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/migration"
	"grpc-boot-starter/persistence/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	bookRepository *repository.BookRepository
	bookService    *BookService
	ctx            context.Context
)

func Before() {
	config.LoadConf(config.GetBasePath(), "test")
	migration.Migrate()
	//
	db.ConnectDB()
	//
	ctx = security.SaveCurrentUser(context.Background(), &security.Principle{ID: "1"})
	//
	bookRepository = repository.NewBookRepository()
	bookService = NewBookService(bookRepository)
}

func TestMain(t *testing.M) {
	Before()
	t.Run()
	After()
}

func After() {
	db.Close()
}

func TestCreateBook(t *testing.T) {
	in := &protogen.CreateBookRequest{
		Title:       "Book A",
		Description: "This is Book A",
		Amount:      100,
		Price:       10.0,
		Category:    protogen.BookCategory_BOOK_CATEGORY_JAVA,
		Author: &protogen.Author{
			Name: "Author J",
		},
	}
	//
	book := bookService.CreateBook(ctx, in)
	assert.NotNil(t, book)
	//
	assert.True(t, book.Id > 0)
	assert.Equal(t, in.Title, book.Title)
}

func TestGetBook(t *testing.T) {
	defer func() {
		r := recover()
		assert.Nil(t, r)
	}()
	in := &protogen.GetBookRequest{
		Id: 4,
	}
	//
	book := bookService.GetBook(ctx, in)
	assert.NotNil(t, book)
	//
	assert.True(t, book.Id > 0)
	assert.Equal(t, in.Id, book.Id)
}

func TestGetBookNonExist(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	in := &protogen.GetBookRequest{
		Id: 400000000,
	}
	//
	book := bookService.GetBook(ctx, in)
	assert.Nil(t, book)
}

func TestUpdateBook(t *testing.T) {
	in := &protogen.UpdateBookRequest{
		Id:          4,
		Title:       "Book A",
		Description: "This is Book A",
		Amount:      110,
		Price:       11.0,
		Category:    protogen.BookCategory_BOOK_CATEGORY_JAVA,
		Author: &protogen.Author{
			Name: "Author J",
		},
	}
	//
	book := bookService.UpdateBook(ctx, in)
	assert.NotNil(t, book)
	//
	assert.True(t, book.Id > 0)
	assert.Equal(t, in.Title, book.Title)
}

func TestUpdateBookNonExists(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	in := &protogen.UpdateBookRequest{
		Id:          999999999,
		Title:       "Book A",
		Description: "This is Book A",
		Amount:      110,
		Price:       11.0,
		Category:    protogen.BookCategory_BOOK_CATEGORY_JAVA,
		Author: &protogen.Author{
			Name: "Author J",
		},
	}
	//
	book := bookService.UpdateBook(ctx, in)
	assert.Nil(t, book)
}
