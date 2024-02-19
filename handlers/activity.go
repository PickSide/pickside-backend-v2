package handlers

import (
	"net/http"
	"pickside/service/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetAllActivities(g *gin.Context) {
	results, err := data.AllActivities()
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}

type CreateActivityRequest struct {
	Address     string  `json:"address" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	MaxPlayers  int     `json:"maxPlayers" binding:"required"`
	Description string  `json:"description"`
	OrganizerID int64   `json:"organizerId" binding:"required"`
	Price       float32 `json:"price" binding:"omitempty"`
	Rules       string  `json:"rules"`
	Time        string  `json:"time" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	IsPrivate   bool    `json:"isPrivate"`
	SportID     uint64  `json:"sportId" binding:"required"`
}

func HandleCreateActivity(g *gin.Context) {
	var req CreateActivityRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := data.InsertActivity(
		req.Address,
		req.Date,
		req.MaxPlayers,
		req.Description,
		req.OrganizerID,
		req.Price,
		req.Rules,
		req.Time,
		req.Title,
		req.IsPrivate,
		req.SportID,
	)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"result": result})
	return
}

type RegisterToActivityRequest struct {
	ActivityId uint64 `json:"activityId" binding:"required"`
	UserId     uint64 `json:"userId" binding:"required"`
}

func HandleParticipantsRegistration(g *gin.Context) {
	var req RegisterToActivityRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	participants, ok := data.UpdateParticipants(req.ActivityId, req.UserId)
	if !ok {
		g.JSON(http.StatusInternalServerError, gin.H{"updated": false})
		return
	}

	g.JSON(http.StatusOK, gin.H{"result": participants})
	return
}

func HandleGetParticipants(g *gin.Context) {
	activityIdString := g.Params.ByName("activityId")

	activityId, err := strconv.ParseUint(activityIdString, 10, 64)
	if err != nil {
		g.JSON(http.StatusNotFound, err.Error())
		return
	}

	results, err := data.GetParticipants(activityId)
	if err != nil {
		g.JSON(http.StatusNotFound, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"results": results})
	return
}
