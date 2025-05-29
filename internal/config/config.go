package config

import (
	"fmt"
	"os"

	"github.com/GuiFernandess7/echoAuthBoilerplate/pkg/database"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	LogLevel   string
	DB         database.DBConfig
}

func LoadConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	return &AppConfig{
		ServerPort: os.Getenv("SERVER_PORT"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
		DB: database.DBConfig{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMode"),
		},
	}, nil
}
