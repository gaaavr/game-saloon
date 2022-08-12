package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"saloon"
	"saloon/pkg/cache"
	"saloon/pkg/handler"
	"saloon/pkg/repository"
	"saloon/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("ошибка при чтении конфига:%s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка загрузки переменных окружения:%s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных:%s", err.Error())
	}
	cache := cache.NewCache()
	err = cache.RestoreCache(db)
	if err != nil {
		log.Fatalf("Ошибка кеширования данных из базы:%s", err.Error())
	}
	fmt.Println(cache)
	repo := repository.NewRepository(db)
	services := service.NewService(cache, repo)
	handler := handler.NewHandler(services)
	router := handler.Routing()
	srv := new(saloon.Server)
	if err := srv.Run(viper.GetString("port"), router.HandleRequest); err != nil {
		log.Fatalf("ошибка при запуске сервера:%s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
