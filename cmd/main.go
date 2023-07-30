package main

import (
	"github.com/JamshedJ/todo"
	"github.com/JamshedJ/todo/configs"
	"github.com/JamshedJ/todo/pkg/handler"
	"github.com/JamshedJ/todo/pkg/repository"
	"github.com/JamshedJ/todo/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := configs.InitConfigs(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occurred while running http server: ", err.Error())
	}
}
