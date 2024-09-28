package env

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	// Загружаем переменные окружения из .env файла (если он есть)
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}

func GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}
