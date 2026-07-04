package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Addr          string
	DBHost        string
	DBPort        string
	DBName        string
	DBUser        string
	DBPassword    string
	AdminUser     string
	AdminPassword string
	TokenSecret   string
	CORSOrigins   []string
	UploadDir     string
	MaxUploadBytes int64
}

func LoadConfig() Config {
	maxUpload := int64(5 << 20)
	if v := os.Getenv("MAX_UPLOAD_MB"); v != "" {
		if mb, err := strconv.ParseInt(v, 10, 64); err == nil && mb > 0 {
			maxUpload = mb << 20
		}
	}
	origins := []string{"*"}
	if v := os.Getenv("CORS_ORIGINS"); v != "" {
		origins = splitCSV(v)
	}
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "oss"
	}
	return Config{
		Addr:          getenv("SERVER_ADDR", ":8080"),
		DBHost:        getenv("DB_HOST", "180.102.24.183"),
		DBPort:        getenv("DB_PORT", "3306"),
		DBName:        getenv("DB_NAME", "wsd_social"),
		DBUser:        getenv("DB_USER", "root"),
		DBPassword:    getenv("DB_PASSWORD", "jny@2216446"),
		AdminUser:     getenv("ADMIN_USER", "admin"),
		AdminPassword: getenv("ADMIN_PASSWORD", "change-me"),
		TokenSecret:   getenv("TOKEN_SECRET", "change-me-token-secret"),
		CORSOrigins:   origins,
		UploadDir:     uploadDir,
		MaxUploadBytes: maxUpload,
	}
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func (c Config) TokenTTL() time.Duration {
	return 24 * time.Hour
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func splitCSV(v string) []string {
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	if len(out) == 0 {
		return []string{"*"}
	}
	return out
}
