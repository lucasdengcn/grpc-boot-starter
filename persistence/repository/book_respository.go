package repository

import (
	"context"
	"grpc-boot-starter/core/logging"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/persistence/entity"
	"sync"
	"time"
)

var (
	instanceBookRepo *BookRepository
	onceBookRepo     sync.Once
)

func NewBookRepository() *BookRepository {
	onceBookRepo.Do(func() {
		instanceBookRepo = &BookRepository{}
	})
	return instanceBookRepo
}

type BookRepository struct {
}

// Create create book
func (b *BookRepository) Create(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	result := db.GetTx(ctx).Create(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

// Update update book all fields
func (b *BookRepository) Update(ctx context.Context, book *entity.Book) (bool, error) {
	result := db.GetTx(ctx).Model(&book).UpdateColumns(book)
	if result.Error != nil {
		return false, result.Error
	}
	logging.Debug(ctx).Msgf("Update book RowsAffected: %v", result.RowsAffected)
	return result.RowsAffected > 0, nil
}

// UpdateStatus update book status field
func (b *BookRepository) UpdateStatus(ctx context.Context, book *entity.Book) (bool, error) {
	result := db.GetTx(ctx).Model(book).Update("status", book.Status)
	if result.Error != nil {
		return false, result.Error
	}
	logging.Debug(ctx).Msgf("Update book status RowsAffected: %v", result.RowsAffected)
	return result.RowsAffected > 0, nil
}

// Delete delete book via id
func (b *BookRepository) Delete(ctx context.Context, id uint32) (bool, error) {
	// update raw sql
	result := db.GetTx(ctx).Exec("update books set deleted = ?, deleted_at = ? where id = ?", true, time.Now(), id)
	//
	logging.Debug(ctx).Msgf("Delete book RowsAffected: %v", result.RowsAffected)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// FindBook find a book via id
func (b *BookRepository) FindBook(ctx context.Context, id uint32) (*entity.Book, error) {
	book := &entity.Book{}
	result := db.GetTx(ctx).Take(book, id)
	if result.Error != nil {
		// errors.Is(result.Error, gorm.ErrRecordNotFound)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return book, nil
}

// FindBooks find books via status, category fields.
func (b *BookRepository) FindBooks(ctx context.Context, status int32, category int32, cursorId uint32) ([]*entity.Book, error) {
	books := []*entity.Book{}
	if status >= 0 && category >= 0 {
		result := db.GetTx(ctx).Where("status = ? and category = ?", status, category).Find(&books)
		if result.Error != nil {
			return nil, result.Error
		}
	} else if status >= 0 {
		result := db.GetTx(ctx).Where("status = ? ", status).Find(&books)
		if result.Error != nil {
			return nil, result.Error
		}
	} else if category >= 0 {
		result := db.GetTx(ctx).Where("category = ?", category).Find(&books)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return books, nil
}

// CountBooks count books via status, category fields.
func (b *BookRepository) CountBooks(ctx context.Context, status int32, category int32, cursorId uint32) (uint32, error) {
	var total uint32 = 0
	if status >= 0 && category >= 0 {
		result := db.GetTx(ctx).Raw("select count(*) from books where status = ? and category = ?", status, category).Scan(&total)
		if result.Error != nil {
			return 0, result.Error
		}
	} else if status >= 0 {
		result := db.GetTx(ctx).Raw("select count(*) from books where status = ?", status).Scan(&total)
		if result.Error != nil {
			return 0, result.Error
		}
	} else if category >= 0 {
		result := db.GetTx(ctx).Raw("select count(*) from books where category = ?", category).Scan(&total)
		if result.Error != nil {
			return 0, result.Error
		}
	}
	return total, nil
}
