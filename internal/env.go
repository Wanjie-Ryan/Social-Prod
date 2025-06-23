package env

import "os"

// makes reading env variables more convinient by giving you a default if the variable isn't set.
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}
