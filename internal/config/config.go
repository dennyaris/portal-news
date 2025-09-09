package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	AppEnv   string
	DBDriver string
	DSN      string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		Port:     getenv("PORT", "9898"),
		AppEnv:   getenv("APP_ENV", "development"),
		DBDriver: getenv("DB_DRIVER", "postgres"),
		DSN:      getenv("DB_DSN", ""),
	}
}

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
