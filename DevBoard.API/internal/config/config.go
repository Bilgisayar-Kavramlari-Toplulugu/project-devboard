package config

import (
	"os"
)

type Config struct {
	DatabaseURL             string
	ServerAddress           string
	JWTSecret               string
	AccessTokenExpireHours  int
	RefreshTokenExpireHours int
	// SMTP Settings
	SMTPHost                   string
	SMTPPort                   int
	SMTPEmail                  string
	SMTPPassword               string
	PasswordResetExpireMinutes int
	PasswordResetBaseURL       string
}

func Load() *Config {
	return &Config{
		DatabaseURL:                getEnv("DATABASE_URL", "host=localhost port=15432 user=postgres password=v7YPHK29g5ZTU1 dbname=devboard sslmode=disable"),
		ServerAddress:              getEnv("SERVER_ADDRESS", ":8080"),
		JWTSecret:                  getEnv("JWT_SECRET", "your-secret-key"),
		AccessTokenExpireHours:     720,  // 1 month
		RefreshTokenExpireHours:    4320, // 6 months
		SMTPHost:                   getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:                   587,
		SMTPEmail:                  getEnv("SMTP_EMAIL", ""),
		SMTPPassword:               getEnv("SMTP_PASSWORD", ""),
		PasswordResetExpireMinutes: 60, // 1 hour
		PasswordResetBaseURL:       getEnv("PASSWORD_RESET_BASE_URL", "https://yourapp.com/reset-password"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Weekday - Custom type to hold value for weekday ranging from 1-7
type DBSchema string

// Declare related constants for each weekday starting with index 1
const (
	Core DBSchema = "core"
	App  DBSchema = "app"
)

// String - Creating common behavior - give the type a String function
func (s DBSchema) String() string {
	return string(s)
}
