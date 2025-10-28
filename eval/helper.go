package eval

import (
	"os"
	"strconv"
)

func GetEnvStr(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetEnvInt(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if ret, err := strconv.ParseInt(value, 10, 64); err == nil {
			return ret
		}
	}
	return defaultValue
}

func GetEnvFloat(key string, defaultValue float64) float64 {
	if value := os.Getenv(key); value != "" {
		if ret, err := strconv.ParseFloat(value, 32); err == nil {
			return ret
		}
	}
	return defaultValue
}

func GetEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if ret, err := strconv.ParseBool(value); err == nil {
			return ret
		}
	}
	return defaultValue
}
