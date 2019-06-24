package utils

import "os"

// GetEnv : wrapper for get os env with default
func GetEnv(key, deflt string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = deflt
	}
	return value
}
