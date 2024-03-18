package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var allowedDomains = []string{"localhost:3000", "localhost:8080", "127.0.0.1:3000", "127.0.0.1:8080", "https://pickside.net"}

func FromValidDomain() gin.HandlerFunc {
	return func(g *gin.Context) {
		incomingDomain := g.Request.Host
		isValidDomain := false

		for _, allowedDomain := range allowedDomains {
			if incomingDomain == allowedDomain {
				isValidDomain = true
				break
			}
		}

		if !isValidDomain {
			g.String(http.StatusForbidden, "invalid domain")
			g.Abort()
			return
		}
	}
}
