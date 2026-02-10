// Package config handles application configuration
package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Config holds all application configuration
type Config struct {
	Database DatabaseConfig
}

// DatabaseConfig holds database connection settings
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "inventory_user"),
			Password: getEnv("DB_PASSWORD", "inventory_pass"),
			DBName:   getEnv("DB_NAME", "restaurant_inventory"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
}

// ConnectDB establishes a connection to PostgreSQL
func (c *Config) ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.User,
		c.Database.Password, c.Database.DBName, c.Database.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	return db, nil
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
