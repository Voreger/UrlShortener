package app

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	StorageType string
	APIPort     string
	PostgresDSN string
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		StorageType: getEnv("STORAGE", "memory"),
		APIPort:     getEnv("API_PORT", "8080"),
		PostgresDSN: getEnv("POSTGRES_DB_STRING", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
