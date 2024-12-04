package services

import (
	"context"
	"grpc-boot-starter/config"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/migration"
	"grpc-boot-starter/persistence/repository"
	"grpc-boot-starter/protogen"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	bookRepository        *repository.BookRepository
	bookServiceServerImpl *BookServiceServerImpl
)

func Before() {
	config.LoadConf(config.GetBasePath(), "test")
	migration.Migrate()
	//
	db.ConnectDB()
	//
	bookRepository = repository.NewBookRepository()
	bookServiceServerImpl = NewBookServiceServerImpl(bookRepository)
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
	in := &protogen.BookCreateInput{
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
	book, err := bookServiceServerImpl.CreateBook(context.Background(), in)
	assert.NoError(t, err)
	assert.NotNil(t, book)
	//
	assert.True(t, book.Id > 0)
	assert.Equal(t, in.Title, book.Title)
}

func TestGetBook(t *testing.T) {
	in := &protogen.BookGetInput{
		Id: 4,
	}
	//
	book, err := bookServiceServerImpl.GetBook(context.Background(), in)
	assert.NoError(t, err)
	assert.NotNil(t, book)
	//
	assert.True(t, book.Id > 0)
	assert.Equal(t, in.Id, book.Id)
}

func TestGetBookNonExist(t *testing.T) {
	in := &protogen.BookGetInput{
		Id: 400000000,
	}
	//
	book, err := bookServiceServerImpl.GetBook(context.Background(), in)
	assert.Nil(t, err)
	assert.Nil(t, book)
}

func TestUpdateBook(t *testing.T) {
	in := &protogen.BookUpdateInput{
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
	book, err := bookServiceServerImpl.UpdateBook(context.Background(), in)
	assert.NoError(t, err)
	assert.NotNil(t, book)
	//
	assert.True(t, book.Id > 0)
	assert.Equal(t, in.Title, book.Title)
}

func TestUpdateBookNonExists(t *testing.T) {
	in := &protogen.BookUpdateInput{
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
	book, err := bookServiceServerImpl.UpdateBook(context.Background(), in)
	assert.NoError(t, err)
	assert.Nil(t, book)
}
