package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var allowedDomains = []string{"localhost:8080", "127.0.0.1:8080", "pickside.net"}

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
		log.Println(isValidDomain)

		if !isValidDomain {
			g.String(http.StatusForbidden, "invalid domain")
			g.Abort()
			return
		}
	}
}
