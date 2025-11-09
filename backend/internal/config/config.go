package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	// Application
	AppEnv     string
	AppName    string
	Port       string
	APIVersion string

	// Database
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBSSLMode         string
	DBMaxConnections  int
	DBMaxIdle         int
	DBMaxLifetime     time.Duration

	// Redis
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	// JWT
	JWTSecret            string
	JWTExpiration        time.Duration
	JWTRefreshSecret     string
	JWTRefreshExpiration time.Duration

	// Security
	BcryptCost         int
	RateLimitRequests  int
	RateLimitDuration  time.Duration

	// CORS
	CORSAllowedOrigins string
	CORSAllowedMethods string
	CORSAllowedHeaders string

	// MinIO
	MinIOEndpoint  string
	MinIOAccessKey string
	MinIOSecretKey string
	MinIOUseSSL    bool
	MinIOBucket    string

	// SMTP
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	SMTPFrom     string

	// Payment
	StripeSecretKey      string
	StripeWebhookSecret  string
	StripePublishableKey string

	// Frontend
	FrontendURL string

	// Logging
	LogLevel  string
	LogFormat string
}

func Load() *Config {
	return &Config{
		// Application
		AppEnv:     getEnv("APP_ENV", "development"),
		AppName:    getEnv("APP_NAME", "Gophiway"),
		Port:       getEnv("APP_PORT", "8080"),
		APIVersion: getEnv("API_VERSION", "v1"),

		// Database
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "5432"),
		DBUser:            getEnv("DB_USER", "gophiway"),
		DBPassword:        getEnv("DB_PASSWORD", "gophiway_dev_password"),
		DBName:            getEnv("DB_NAME", "gophiway_dev"),
		DBSSLMode:         getEnv("DB_SSL_MODE", "disable"),
		DBMaxConnections:  getEnvAsInt("DB_MAX_CONNECTIONS", 100),
		DBMaxIdle:         getEnvAsInt("DB_MAX_IDLE_CONNECTIONS", 10),
		DBMaxLifetime:     time.Duration(getEnvAsInt("DB_MAX_LIFETIME", 3600)) * time.Second,

		// Redis
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),

		// JWT
		JWTSecret:            getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
		JWTExpiration:        parseDuration(getEnv("JWT_EXPIRATION", "15m")),
		JWTRefreshSecret:     getEnv("JWT_REFRESH_SECRET", "your-super-secret-refresh-key"),
		JWTRefreshExpiration: parseDuration(getEnv("JWT_REFRESH_EXPIRATION", "7d")),

		// Security
		BcryptCost:        getEnvAsInt("BCRYPT_COST", 12),
		RateLimitRequests: getEnvAsInt("RATE_LIMIT_REQUESTS", 100),
		RateLimitDuration: parseDuration(getEnv("RATE_LIMIT_DURATION", "1m")),

		// CORS
		CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173"),
		CORSAllowedMethods: getEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,DELETE,OPTIONS"),
		CORSAllowedHeaders: getEnv("CORS_ALLOWED_HEADERS", "Content-Type,Authorization"),

		// MinIO
		MinIOEndpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
		MinIOAccessKey: getEnv("MINIO_ACCESS_KEY", "gophiway"),
		MinIOSecretKey: getEnv("MINIO_SECRET_KEY", "gophiway_minio_password"),
		MinIOUseSSL:    getEnvAsBool("MINIO_USE_SSL", false),
		MinIOBucket:    getEnv("MINIO_BUCKET", "products"),

		// SMTP
		SMTPHost:     getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:     getEnv("SMTP_PORT", "587"),
		SMTPUser:     getEnv("SMTP_USER", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
		SMTPFrom:     getEnv("SMTP_FROM", "noreply@gophiway.com"),

		// Payment
		StripeSecretKey:      getEnv("STRIPE_SECRET_KEY", ""),
		StripeWebhookSecret:  getEnv("STRIPE_WEBHOOK_SECRET", ""),
		StripePublishableKey: getEnv("STRIPE_PUBLISHABLE_KEY", ""),

		// Frontend
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:5173"),

		// Logging
		LogLevel:  getEnv("LOG_LEVEL", "debug"),
		LogFormat: getEnv("LOG_FORMAT", "json"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func parseDuration(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		return 15 * time.Minute
	}
	return duration
}
