package handlers

import (
	"net/http"
	"pickside/service/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetAllGroupsByOrganizer(g *gin.Context) {
	organizerIdString := g.Params.ByName("organizerId")

	organizerId, err := strconv.ParseUint(organizerIdString, 10, 64)
	if err != nil {
		g.JSON(http.StatusNotFound, err.Error())
		return
	}

	results, err := data.AllGroupsByOrganizer(organizerId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}

func HandleGetGroups(g *gin.Context) {
	g.JSON(http.StatusNotImplemented, gin.H{
		"message": "function was not implemented yet",
	})
}

type CreateGroupRequest struct {
	CoverPhoto       string `json:"coverPhoto" binding:"omitempty"`
	Description      string `json:"description" binding:"required"`
	Name             string `json:"name" binding:"required"`
	RequiresApproval bool   `json:"requiresApproval"`
	Visibility       string `json:"visibility" binding:"required"`
	OrganizerId      uint64 `json:"organizerId" binding:"required"`
	SportId          uint64 `json:"sportId" binding:"required"`
}

func HandleCreateGroup(g *gin.Context) {
	var req CreateGroupRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := data.InsertGroup(
		req.CoverPhoto,
		req.Description,
		req.Name,
		req.RequiresApproval,
		req.Visibility,
		req.OrganizerId,
		req.SportId,
	)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"result": results})
	return
}
