//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"saloon"
	"saloon/pkg/repository"
	"saloon/wire"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("ошибка при чтении конфига:%s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка загрузки переменных окружения:%s", err.Error())
	}
	c := repository.Config{
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"),
	}
	handler, err := wire.InitLayers(c)
	if err != nil {
		log.Fatalf("ошибка при инициализации слоёв: %s", err.Error())
	}
	router := handler.Routing()
	srv := new(saloon.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), router.HandleRequest); err != nil {
			log.Fatalf("ошибка при запуске сервера:%s", err.Error())
		}
	}()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit
	fmt.Println("завершаем работу приложения")
	if err = srv.Shutdown(); err != nil {
		log.Fatalf("ошибка при завершении работы сервера: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
