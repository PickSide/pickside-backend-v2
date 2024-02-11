package handlers

import (
	"net/http"
	"pickside/service/data"

	"github.com/gin-gonic/gin"
)

func HandleGetAllSports(g *gin.Context) {
	results, err := data.AllSports()
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}
