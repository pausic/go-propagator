package config

import "os"

type Config struct {
	Addr       string
	ConfigPath string
}

func Load() *Config {
	return &Config{
		Addr:       envOrDefault("ADDR", "localhost:8080"),
		ConfigPath: envOrDefault("CONFIG_PATH", "config.yaml"),
	}

}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
