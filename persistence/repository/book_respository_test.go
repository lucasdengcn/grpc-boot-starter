package repository

import (
	"context"
	"grpc-boot-starter/config"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/migration"
	"grpc-boot-starter/persistence/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	bookRepository *BookRepository
)

func Before() {
	config.LoadConf(config.GetBasePath(), "test")
	migration.Migrate()
	//
	db.ConnectDB()
	//
	bookRepository = NewBookRepository()
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
	ctx := context.Background()
	book, err := bookRepository.Create(ctx, &entity.Book{
		Title:       "Java Book",
		Description: "Book Description",
		Author:      "Author John",
		Price:       10.0,
		Amount:      100,
		Category:    "JAVA",
		Status:      1,
		Active:      true,
	})
	//
	assert.NoError(t, err)
	assert.True(t, book.ID > 0)
}

func TestUpdateBook(t *testing.T) {
	ctx := context.Background()
	//
	newBook(ctx, t)
}

func newBook(ctx context.Context, t *testing.T) *entity.Book {
	book := &entity.Book{
		Title:       "Java Book II",
		Description: "Book Description updated",
		Author:      "Author John updated",
		Price:       12.0,
		Amount:      110,
		Category:    "JAVA",
		Status:      1,
		Active:      true,
	}
	book, err := bookRepository.Create(ctx, book)
	assert.NoError(t, err)
	assert.NotNil(t, book)
	return book
}

func TestUpdateBookStatus(t *testing.T) {
	ctx := context.Background()
	ok, err := bookRepository.UpdateStatus(ctx, &entity.Book{
		Model: gorm.Model{
			ID: 1,
		},
		Status: 2,
	})
	//
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestDeleteBook(t *testing.T) {
	ctx := context.Background()
	ok, err := bookRepository.Delete(ctx, 2)
	//
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestFindDeletedBook(t *testing.T) {
	ctx := context.Background()
	book, err := bookRepository.FindBook(ctx, 2)
	//
	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	assert.Nil(t, book)
}

func TestFindNonExistBook(t *testing.T) {
	ctx := context.Background()
	book, err := bookRepository.FindBook(ctx, 20000000)
	//
	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	assert.Nil(t, book)
}

func TestFindActiveBook(t *testing.T) {
	ctx := context.Background()
	book := newBook(ctx, t)
	//
	book, err := bookRepository.FindBook(ctx, int(book.ID))
	//
	assert.NoError(t, err)
	assert.NotNil(t, book)
}

func TestFindActiveBooks(t *testing.T) {
	ctx := context.Background()
	//
	books, err := bookRepository.FindBooks(ctx, 1, "JAVA", 0)
	//
	assert.NoError(t, err)
	assert.NotNil(t, books)
	assert.NotEmpty(t, books)
}

func TestFindActiveBooksInvalidArgs(t *testing.T) {
	ctx := context.Background()
	// category not exist
	books, err := bookRepository.FindBooks(ctx, 1, "Java", 0)
	//
	assert.NoError(t, err)
	assert.NotNil(t, books)
	assert.Empty(t, books)
}

func TestCountActiveBooksInvalidArgs(t *testing.T) {
	ctx := context.Background()
	// category not exist
	count, err := bookRepository.CountBooks(ctx, 1, "Java", 0)
	//
	assert.NoError(t, err)
	assert.True(t, count == 0)
}

func TestCountActiveBooksValidArgs(t *testing.T) {
	ctx := context.Background()
	// category exist
	count, err := bookRepository.CountBooks(ctx, 1, "JAVA", 0)
	//
	assert.NoError(t, err)
	assert.True(t, count > 0)
}
