package handlers

import (
	"net/http"
	"pickside/service/data"
	"pickside/service/util"

	"github.com/gin-gonic/gin"
)

func HandleMeSettings(g *gin.Context) {
	refreshToken, err := g.Cookie("refreshToken")
	if err != nil {
		g.JSON(http.StatusUnauthorized, err)
	}

	parsedToken, err := util.ExtractClaims(refreshToken)
	if err != nil {
		g.JSON(http.StatusUnauthorized, err)
	}

	settings, err := data.GetMeSettings(uint64(parsedToken.ID))
	if err != nil {
		g.JSON(http.StatusNotFound, err)
	}

	g.JSON(http.StatusOK, gin.H{"result": settings})
}
