package internal

import (
	"os"
)

const (
	defaultPort = "8080"
)

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return defaultPort
	}

	return port
}
