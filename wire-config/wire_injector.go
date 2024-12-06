package server

import (
	"grpc-boot-starter/apis/controller"
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/persistence/repository"
	"grpc-boot-starter/service"

	"github.com/google/wire"
)

// ProviderSet
var dbSet = wire.NewSet(db.GetDBCon)

// bookServiceSet
var bookServiceSet = wire.NewSet(repository.NewBookRepository, service.NewBookService)

func InitializeBookService() *service.BookService {
	wire.Build(bookServiceSet)
	return &service.BookService{}
}

func InitializeBookController() *controller.BookControllerImpl {
	wire.Build(bookServiceSet, controller.NewBookControllerImpl)
	return &controller.BookControllerImpl{}
}

// helloServiceSet

func InitializeHelloController() *controller.HelloControllerImpl {
	wire.Build(controller.NewHelloControllerImpl)
	return &controller.HelloControllerImpl{}
}
