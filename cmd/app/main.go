package main

import (
	expanses_rest_api "expanses_rest_api/intenal"
	"expanses_rest_api/intenal/handler"
	"expanses_rest_api/intenal/repository"
	"expanses_rest_api/intenal/service"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	fmt.Println("Application starting...")

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		UserName: "postgres",
		Password: "12345678",
		DBName:   "expense",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	srv := new(expanses_rest_api.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
