package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	POSTGRESQL   PostgreSQLConfig
	SMTP         SMTPConfig
	SERVER       ServerConfig
}

type PostgreSQLConfig struct {
	POSTGRESQL_USER string
	POSTGRESQL_PASS string
	POSTGRESQL_HOST string
	POSTGRESQL_PORT string
	POSTGRESQL_NAME string
}


type SMTPConfig struct {
	SMTP_USER string
	SMTP_PASS string
	SMTP_PORT string
	SMTP_HOST string
}

type ServerConfig struct {
	SERVER_HOST string
	SERVER_PORT string
}


func LoadConfig() (*Configuration, error) {

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load environment variables from .env file: %w", err)
		}
	} else {
		fmt.Println("warning: .env file not found. make sure environment variables are set")
	}

	return &Configuration{
		POSTGRESQL: PostgreSQLConfig{
			POSTGRESQL_USER: os.Getenv("POSTGRESQL_USER"),
			POSTGRESQL_PASS: os.Getenv("POSTGRESQL_PASS"),
			POSTGRESQL_HOST: os.Getenv("POSTGRESQL_HOST"),
			POSTGRESQL_PORT: os.Getenv("POSTGRESQL_PORT"),
			POSTGRESQL_NAME: os.Getenv("POSTGRESQL_NAME"),
		},
		SMTP: SMTPConfig{
			SMTP_USER: os.Getenv("SMTP_USER"),
			SMTP_PASS: os.Getenv("SMTP_PASS"),
			SMTP_PORT: os.Getenv("SMTP_PORT"),
			SMTP_HOST: os.Getenv("SMTP_HOST"),
		},
		SERVER: ServerConfig{
			SERVER_HOST: os.Getenv("SERVER_HOST"),
			SERVER_PORT: os.Getenv("SERVER_PORT"),
		},
	}, nil
}
