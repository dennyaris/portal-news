package config

import "os"

type Config struct {
	Port     string
	AppEnv   string
	DBDriver string
	DSN      string
}

func Load() *Config {
	return &Config{
		Port:     getenv("PORT", "8080"),
		AppEnv:   getenv("APP_ENV", "development"),
		DBDriver: getenv("DB_DRIVER", "mysql"),
		DSN:      getenv("DB_DSN", ""),
	}
}

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" { return v }
	return d
}
