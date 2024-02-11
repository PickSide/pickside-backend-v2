package main

import (
	"log"
	"os"
	"pickside/service/db"
	"pickside/service/handlers"
	"pickside/service/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	g := gin.Default()

	if err := db.OpenConnectionToDB(); err != nil {
		log.Fatal(err)
	}

	g.Use(cors.Default())
	v1 := g.Group("/api/v1", middlewares.FromValidDomain())

	//activities
	v1.GET("/activities", handlers.HandleGetAllActivities)
	v1.GET("/activities/:activityId/participants", handlers.HandleGetParticipants)
	v1.POST("/activities", middlewares.WithToken(), handlers.HandleCreateActivity)
	v1.PUT("/activities/registration", middlewares.WithToken(), handlers.HandleParticipantsRegistration)

	// groups
	v1.GET("/groups/users/:organizerId", handlers.HandleGetAllGroupsByOrganizer)
	v1.GET("/groups/:id", handlers.HandleGetGroups)
	v1.POST("/groups", handlers.HandleCreateGroup)

	// locales
	v1.GET("/locales", handlers.HandleGetAllLocales)

	// sports
	v1.GET("/sports", handlers.HandleGetAllSports)

	// user
	v1.GET("/me", middlewares.WithToken(), handlers.HandleMe)
	v1.GET("/me-settings", middlewares.WithToken(), handlers.HandleMeSettings)
	v1.GET("/logout", handlers.HandleLogout)
	v1.POST("/login", handlers.HandleLogin)
	v1.POST("/users", handlers.HandleCreateMe)

	err := g.Run(os.Getenv("LISTEN_PORT"))

	if err != nil {
		log.Fatal(err)
	}
}
