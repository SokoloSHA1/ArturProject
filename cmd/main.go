package main

import (
	arturproject "github.com/SokoloSHA/ArturProject"
	"github.com/SokoloSHA/ArturProject/pkg/handlers"
	"github.com/SokoloSHA/ArturProject/pkg/repository"
	"github.com/SokoloSHA/ArturProject/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	srv := new(arturproject.Server)

	repos := repository.NewRepository()
	service := services.NewService(repos)
	handlers := handlers.NewHandler(service)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
