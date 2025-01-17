package config

import (
    "log"
    "os"
    "strings" 

    "github.com/joho/godotenv"
)

type Config struct {
    BaseURL           string
    JWTSecret         string
    Environment       string 
    CORSAllowedOrigins []string  
    Port              string 
}

func LoadConfig() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using default values")
    }

    return &Config{
        BaseURL:         getEnv("POCKET_BASE_URL", "http://127.0.0.1:8090/api"),
        JWTSecret:       getEnv("JWT_SECRET", "your_jwt_secret_key"),
        Environment:     getEnv("APP_ENV", "development"), 
        CORSAllowedOrigins: getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"*"}), 
        Port:            getEnv("PORT", "8080"), 
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
    if value, exists := os.LookupEnv(key); exists {
        return splitCommaSeparated(value)
    }
    return defaultValue
}

func splitCommaSeparated(value string) []string {
    return strings.Split(value, ",")
}
