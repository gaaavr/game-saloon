package main

import (
	"log"
	"saloon"
	"saloon/pkg/handler"
	"saloon/pkg/repository"
	"saloon/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handler := handler.NewHandler(services)
	router := handler.Routing()
	srv := new(saloon.Server)
	if err := srv.Run("8080", router.HandleRequest); err != nil {
		log.Fatalf("ошибка при запуске сервера:%s", err.Error())
	}
}
