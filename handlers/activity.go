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
	Address     string   `json:"address" binding:"required"`
	Date        string   `json:"date" binding:"required"`
	Description string   `json:"description,omitempty"`
	GameMode    string   `json:"gameMode" binding:"required"`
	Images      []string `json:"images,omitempty"`
	IsPrivate   bool     `json:"isPrivate"`
	Lat         float32  `json:"lat,omitempty"`
	Lng         float32  `json:"lng,omitempty"`
	MaxPlayers  int      `json:"maxPlayers" binding:"required"`
	OrganizerID uint64   `json:"organizerId" binding:"required"`
	Price       float32  `json:"price,omitempty"`
	Rules       string   `json:"rules,omitempty"`
	SportID     uint64   `json:"sportId" binding:"required"`
	Time        string   `json:"time" binding:"required"`
	Title       string   `json:"title" binding:"required"`
}

func HandleCreateActivity(g *gin.Context) {
	var req CreateActivityRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := data.CreateActivity(
		req.Address,
		req.Date,
		req.Description,
		req.GameMode,
		req.Images,
		req.IsPrivate,
		req.Lat,
		req.Lng,
		req.MaxPlayers,
		req.OrganizerID,
		req.Price,
		req.Rules,
		req.SportID,
		req.Time,
		req.Title,
	)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"message": "created successfully",
		"result":  result,
	})
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
