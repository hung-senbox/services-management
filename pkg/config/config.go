package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Server  ServerConfig
	MongoDB MongoDBConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Host string
	Port string
}

// MongoDBConfig holds MongoDB configuration
type MongoDBConfig struct {
	Host     string
	Port     string
	User     string // Optional - leave empty for no auth
	Password string // Optional - leave empty for no auth
	DBName   string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if exists (errors ignored)
	_ = godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "8080"),
		},
		MongoDB: MongoDBConfig{
			Host:     getEnv("MONGO_HOST", "localhost"),
			Port:     getEnv("MONGO_PORT", "27017"),
			User:     getEnv("MONGO_USER", ""),     // Empty = no authentication
			Password: getEnv("MONGO_PASSWORD", ""), // Empty = no authentication
			DBName:   getEnv("MONGO_DB_NAME", "services_management"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
