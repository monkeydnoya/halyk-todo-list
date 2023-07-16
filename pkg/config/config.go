package config

import (
	"os"

	"github.com/joho/godotenv"
)

var isLoaded bool = false

func Config(key string) string {
	if !isLoaded {
		godotenv.Load(".env")
		isLoaded = true
	}
	return os.Getenv(key)
}
