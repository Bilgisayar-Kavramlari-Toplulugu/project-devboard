package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL             string
	ServerAddress           string
	JWTSecret               string
	AccessTokenExpireHours  int
	RefreshTokenExpireHours int
	AccessTokenCookieName   string
	RefreshTokenCookieName  string
	CookieDomain            string
	CookieSecure            bool
	CookieSameSite          string
	CORSAllowedOrigins      []string
	// SMTP Settings
	SMTPHost                   string
	SMTPPort                   int
	SMTPEmail                  string
	SMTPPassword               string
	PasswordResetExpireMinutes int
	PasswordResetBaseURL       string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error reading it, falling back to OS environment variables")
	}
	return &Config{
		DatabaseURL:                getEnv("DATABASE_URL", ""),
		ServerAddress:              getEnv("SERVER_ADDRESS", ":8080"),
		JWTSecret:                  getEnv("JWT_SECRET", ""),
		AccessTokenExpireHours:     getEnvInt("ACCESS_TOKEN_EXPIRE_HOURS", 720),
		RefreshTokenExpireHours:    getEnvInt("REFRESH_TOKEN_EXPIRE_HOURS", 4320),
		AccessTokenCookieName:      getEnv("ACCESS_TOKEN_COOKIE_NAME", "access_token"),
		RefreshTokenCookieName:     getEnv("REFRESH_TOKEN_COOKIE_NAME", "refresh_token"),
		CookieDomain:               getEnv("COOKIE_DOMAIN", ""),
		CookieSecure:               getEnvBool("COOKIE_SECURE", true),
		CookieSameSite:             getEnv("COOKIE_SAME_SITE", "None"),
		CORSAllowedOrigins:         getEnvList("CORS_ALLOWED_ORIGINS", []string{"http://localhost", "http://127.0.0.1"}),
		SMTPHost:                   getEnv("SMTP_HOST", ""),
		SMTPPort:                   587,
		SMTPEmail:                  getEnv("SMTP_EMAIL", ""),
		SMTPPassword:               getEnv("SMTP_PASSWORD", ""),
		PasswordResetExpireMinutes: 60, // 1 hour
		PasswordResetBaseURL:       getEnv("PASSWORD_RESET_BASE_URL", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid integer value for %s: %q. Using default %d", key, value, defaultValue)
		return defaultValue
	}

	return parsed
}

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		log.Printf("Invalid boolean value for %s: %q. Using default %t", key, value, defaultValue)
		return defaultValue
	}

	return parsed
}

func getEnvList(key string, defaultValue []string) []string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return defaultValue
	}

	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			result = append(result, item)
		}
	}

	if len(result) == 0 {
		return defaultValue
	}

	return result
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
