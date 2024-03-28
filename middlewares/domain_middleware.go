package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var allowedDomains = []string{
	"localhost:8080",
	"127.0.0.1:8080",
	"localhost:8081",
	"127.0.0.1:8081",
	"localhost:8082",
	"127.0.0.1:8082",
	"localhost:8084",
	"127.0.0.1:8084",
	"pickside-service.com",
}

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
