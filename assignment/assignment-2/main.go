package main

import (
	"assignment-2/app/configs"
	"assignment-2/app/databases/postgresql"
	"assignment-2/app/routes"
	"assignment-2/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load configuration: %v", err)
	}

	db := postgresql.ConnectPostgreSQL()

	r := gin.Default()

	middlewares.Logger(r)
  
    routes.SetupRoutes(r, db)

	host := config.SERVER.SERVER_HOST
	port := config.SERVER.SERVER_PORT
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "8000"
	}
	address := host + ":" + port

	log.Printf("server is running on address %s...", address)
	if err := r.Run(address); err != nil {
		logrus.Fatalf("error starting server: %v", err)
	}
}
