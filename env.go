package env

import (
	"os"
	"strconv"
	"time"
)

// Int returns the value of the environment variable with the given name as an integer.
// If the environment variable is not set or cannot be parsed as an integer, the default value is returned.
func Int(name string, defaultValue int) int {
	// load the environment variable
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	// parse the environment variable
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// Float64 returns the value of the environment variable with the given name as a float64.
// If the environment variable is not set or cannot be parsed as a float64, the default value is returned.
func Float64(name string, defaultValue float64) float64 {
	// load the environment variable
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	// parse the environment variable
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return floatValue
}

// Duration returns the value of the environment variable with the given name as a time.Duration.
// If the environment variable is not set or cannot be parsed as a time.Duration, the default value is returned.
func Duration(name string, defaultValue time.Duration) time.Duration {
	// load the environment variable
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	// parse the environment variable
	durationValue, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return durationValue
}

// String returns the value of the environment variable with the given name as a string.
// If the environment variable is not set, the default value is returned.
func String(name string, defaultValue string) string {
	// load the environment variable
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}
