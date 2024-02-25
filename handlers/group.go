package handlers

import (
	"errors"
	"net/http"
	"pickside/service/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetGroups(g *gin.Context) {
	g.JSON(http.StatusNotImplemented, gin.H{
		"message": "function was not implemented yet",
	})
}

func HandleGetAllGroupsByOrganizer(g *gin.Context) {
	organizerId, err := strconv.ParseUint(g.Params.ByName("organizerId"), 10, 64)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	results, err := data.AllGroupsByOrganizer(organizerId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}

type CreateGroupRequest struct {
	Description      string   `json:"description,omitempty"`
	Members          []uint64 `json:"members" binding:"required"`
	Name             string   `json:"name" binding:"required"`
	OrganizerId      uint64   `json:"organizerId" binding:"required"`
	RequiresApproval bool     `json:"requiresApproval"`
	SportId          uint64   `json:"sportId" binding:"required"`
	Visibility       string   `json:"visibility" binding:"required"`
}

func (c *CreateGroupRequest) ValidateFields() error {
	if c.Visibility != "private" && c.Visibility != "public" {
		return errors.New("bad payload")
	}
	return nil
}

func HandleCreateGroup(g *gin.Context) {
	var req CreateGroupRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := req.ValidateFields()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := data.InsertGroup(data.Group{
		Description:      req.Description,
		Name:             req.Name,
		RequiresApproval: req.RequiresApproval,
		Visibility:       req.Visibility,
		OrganizerID:      req.OrganizerId,
		SportID:          req.SportId,
	})
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, member := range req.Members {
		data.InsertGroupUsers(results.ID, member)
	}

	g.JSON(http.StatusCreated, gin.H{
		"message": "group created",
		"result":  results,
	})
	return
}

type DeleteGroupRequest struct {
	OrganizerID uint64 `json:"organizerId,omitempty"`
}

func HandleDeleteGroup(g *gin.Context) {
	groupId, err := strconv.ParseUint(g.Params.ByName("groupId"), 10, 64)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	organizerId, err := strconv.ParseUint(g.Params.ByName("organizerId"), 10, 64)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	result, err := data.DeleteGroup(groupId, organizerId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	rowsAffected, err := (*result).RowsAffected()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if rowsAffected == 0 {
		g.JSON(http.StatusNotFound, gin.H{"message": "group not found"})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "successfully removed",
	})
	return
}
