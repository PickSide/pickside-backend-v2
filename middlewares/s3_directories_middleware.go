package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var allowedDirectories = []string{"activities", "messages", "users"}

func ValidateS3Path() gin.HandlerFunc {
	return func(g *gin.Context) {
		incomingPath := g.Param("path")
		isValidPath := false

		for _, allowedDirectory := range allowedDirectories {
			if incomingPath == allowedDirectory {
				isValidPath = true
				break
			}
		}

		if !isValidPath {
			g.String(http.StatusForbidden, "invalid s3 path")
			g.Abort()
			return
		}
	}
}
