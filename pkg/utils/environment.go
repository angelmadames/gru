package utils

import (
	"github.com/charmbracelet/log"
	"os"
)

func GetEnvironment(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatal(
			"A required env var is not set.",
			"environmentVariable",
			key,
		)
	}

	return value
}
