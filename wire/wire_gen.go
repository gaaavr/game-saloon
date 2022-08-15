// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"saloon/pkg/cache"
	"saloon/pkg/handler"
	"saloon/pkg/repository"
	"saloon/pkg/service"
)

// Injectors from wire.go:

func InitLayers(config repository.Config) (*handler.Handler, error) {
	db, err := repository.NewPostgresDB(config)
	if err != nil {
		return nil, err
	}
	cacheCache, err := cache.NewCache(db)
	if err != nil {
		return nil, err
	}
	repositoryRepository := repository.NewRepository(db)
	serviceService := service.NewService(cacheCache, repositoryRepository)
	handlerHandler := handler.NewHandler(serviceService)
	return handlerHandler, nil
}