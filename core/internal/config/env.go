package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl     string
}

var Envs = initConfig()

func initConfig() Config {
	if err := LoadEnv(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		DBUrl:     getEnv("DB_STRING"),
	}

	return config
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Fatal("Could get ENV variable ", key)
	return ""
}

func LoadEnv() error {
	// Start from current directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Keep going up until we find .env or hit root
	for {
		if _, err := os.Stat(filepath.Join(dir, ".env")); err == nil {
			return godotenv.Load(filepath.Join(dir, ".env"))
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return errors.New(".env file not found")
		}
		dir = parent
	}
}
