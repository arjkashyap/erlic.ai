package env

import (
	"os"
	"strconv"

	"github.com/arjkashyap/erlic.ai/internal/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Error("Error loading .env file")
	}
}

func GetString(key string, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)

	if err != nil {
		return fallback
	}

	return valAsInt
}

func GetBool(key string, fallback bool) bool {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}

	return boolVal
}
