package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleTestHello(g *gin.Context) {
	g.JSON(http.StatusOK, "Hello")
}
