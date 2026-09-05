package config

import (
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBHost     string
	JWTSecret  string
}

func Load() *Config {
	return &Config{
		DBUser:     getEnv("DB_USER", "lyola"),
		DBPassword: getEnv("DB_PASSWORD", "13lyola"),
		DBName:     getEnv("DB_NAME", "lyoladb"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		JWTSecret:  getEnv("JWT_SECRET", "L7pxkiaY/eZHEXHcvgHcZx93G8hqwvNDJamBIKyn1E4="),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
