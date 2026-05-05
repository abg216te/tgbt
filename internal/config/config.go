package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// DB
	DBUrl string
	// Telegram Bot Token
	Token string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
		// Не возвращаем ошибку - пытаемся использовать системные переменные
	}

	dbUrl := getEnv("DB_URL", "")
	if dbUrl == "" {
		return nil, fmt.Errorf("DB_URL environment variable is not set or empty")
	}

	token := getEnv("TOKEN", "")
	if token == "" {
		return nil, fmt.Errorf("TOKEN environment variable is not set or empty")
	}

	return &Config{
		DBUrl: dbUrl,
		Token: token,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
