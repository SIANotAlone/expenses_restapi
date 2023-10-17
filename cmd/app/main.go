package main

import (
	expanses_rest_api "expanses_rest_api/intenal"
	"expanses_rest_api/intenal/handler"
	"expanses_rest_api/intenal/repository"
	"expanses_rest_api/intenal/service"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		ForceColors:     true,
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true})
	fmt.Println("Application starting...")

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("eror loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	srv := new(expanses_rest_api.Server)
	fmt.Println(fmt.Sprintf("[info] Server running on port: %s", viper.GetString("port")))
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
