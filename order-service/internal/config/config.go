package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	DBHost      string
	DBPort      int
	DBUser      string
	DBPassword  string
	DBName      string
	Environment string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Port:        getEnv("PORT", "50051"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      dbPort,
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "postgres"),
		DBName:      getEnv("DB_NAME", "orders_service"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
