package services

import (
	"grpc-boot-starter/persistence/repository"
	"grpc-boot-starter/protogen"
)

func NewBookServiceServerImpl(bookRepository *repository.BookRepository) *BookServiceServerImpl {
	return &BookServiceServerImpl{
		BookRepository: bookRepository,
	}
}

type BookServiceServerImpl struct {
	protogen.UnimplementedBookServiceServer
	BookRepository *repository.BookRepository
}
