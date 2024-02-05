package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestAccessToken(g *gin.Context) {
	AT, err := g.Cookie("accessToken")
	if err != nil {
		g.JSON(http.StatusForbidden, "Forbidden, lost access")
	}
	g.JSON(http.StatusOK, gin.H{
		"result":      "Good, access token was refreshed",
		"accessToken": AT,
	})
}
