package env

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envPath := filepath.Join("./", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("⚠️ No .env file found, using system environment variables")
	}
}

func GetStringEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultValue
}

func GetIntEnv(key string, defaultValue int) int {

	if val, ok := os.LookupEnv(key); ok {
		intVal, err := strconv.Atoi(val)

		if err != nil {
			return defaultValue
		}
		return intVal
	}

	return defaultValue
}
