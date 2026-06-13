package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds service configuration.
type Config struct {
	Port     int
	DBHost   string
	DBPass   string
	Workers  int
	LogLevel string
}

// LoadConfig reads configuration from environment variables.
// In production, this would be more sophisticated (YAML, flags, etc.).
func LoadConfig() *Config {
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	workers, _ := strconv.Atoi(os.Getenv("APP_WORKERS"))
	return &Config{
		Port:     port,
		DBHost:   os.Getenv("APP_DB_HOST"),
		DBPass:   os.Getenv("APP_DB_PASS"),
		Workers:  workers,
		LogLevel: os.Getenv("APP_LOG_LEVEL"),
	}
}

// Validate checks that the config is usable.
func (c *Config) Validate() error {
	// TODO: implement validation
	return nil
}

// String returns a safe representation of the config for logging.
func (c *Config) String() string {
	// TODO: redact DBPass
	return fmt.Sprintf("Config{Port:%d DBHost:%s DBPass:%s Workers:%d LogLevel:%s}",
		c.Port, c.DBHost, c.DBPass, c.Workers, c.LogLevel)
}

func main() {
	cfg := LoadConfig()
	if err := cfg.Validate(); err != nil {
		fmt.Println("config error:", err)
		os.Exit(1)
	}
	fmt.Println("config loaded:", cfg.String())
}
