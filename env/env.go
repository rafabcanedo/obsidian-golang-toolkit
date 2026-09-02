package env

import "os"

// GetString retrieves the value of the environment variable named by key.
// If the variable is not set or its value is empty, it returns defaultValue.
func GetString(key, defaultValue string) string {
	value, exist := os.LookupEnv(key)

	if !exist || value == "" {
		return defaultValue
	}

	return value
}
