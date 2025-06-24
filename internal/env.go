package env

import (
	"os"
	"strconv"
)

// makes reading env variables more convinient by giving you a default if the variable isn't set.
// the function below gets the env variable as an integer
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

// the function below gets the env variable as an integer
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	//  convert the key from string to int
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valAsInt
}
