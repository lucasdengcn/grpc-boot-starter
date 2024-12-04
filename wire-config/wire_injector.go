package main

import (
	"grpc-boot-starter/infra/db"
	"grpc-boot-starter/persistence/repository"
	"grpc-boot-starter/services"

	"github.com/google/wire"
)

// ProviderSet
var dbSet = wire.NewSet(db.GetDBCon)

// bookServiceSet
var bookServiceSet = wire.NewSet(repository.NewBookRepository, services.NewBookServiceServerImpl)

func InitializeBookService() *services.BookServiceServerImpl {
	wire.Build(bookServiceSet)
	return &services.BookServiceServerImpl{}
}

// helloServiceSet
var helloServiceSet = wire.NewSet(services.NewHelloServiceServerImpl)

func InitializeHelloService() *services.HelloServiceServerImpl {
	wire.Build(helloServiceSet)
	return &services.HelloServiceServerImpl{}
}
