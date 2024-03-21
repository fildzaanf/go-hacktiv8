package main

import (
	"final-project/app/configs"
	"final-project/app/databases/postgresql"
	"final-project/app/routes"

	"final-project/docs"
	"final-project/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @description This is a simple service for media social
// @termsOfService http://swagger.io/terms/
// @host localhost:8000
// @BasePath /

func main() {
	godotenv.Load()
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load configuration: %v", err)
	}

	docs.SwaggerInfo.Title = "Swagger MyGram API"
	docs.SwaggerInfo.Description = "This is a simple service for media social"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := postgresql.ConnectPostgreSQL()
	r := gin.Default()

	middlewares.Logger(r)

	routes.SetupRoutes(r, db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
