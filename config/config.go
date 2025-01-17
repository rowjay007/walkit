package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	BaseURL         string
	JWTSecret       string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file in development
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	return &Config{
		BaseURL:   getEnv("POCKET_BASE_URL", "http://127.0.0.1:8090/api"),
		JWTSecret: getEnv("JWT_SECRET", "your_jwt_secret_key"),
	}
}

// getEnv retrieves the value of the environment variable or returns the default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
