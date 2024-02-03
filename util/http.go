package util

import (
	"os"
)

func IsSecure() bool {
	env := os.Getenv("GO_ENV")

	switch env {
	case "prod":
		return true
	default:
		return false
	}
}
