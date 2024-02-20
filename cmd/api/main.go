package main

import (
	"log"
	"os"
	"pickside/service/db"
	"pickside/service/handlers"
	"pickside/service/middlewares"
	"time"

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

	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://localhost:3000", "https://pickside.net"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "X-Request-Id", "Content-Type"},
		AllowWebSockets:  true,
		ExposeHeaders:    []string{"X-Request-Id", "Content-Type", "Content-Length", "Content-Encoding"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v2 := g.Group("/api/v2", middlewares.FromValidDomain())

	//activities
	v2.GET("/activities", handlers.HandleGetAllActivities)
	v2.GET("/activities/:activityId/participants", handlers.HandleGetParticipants)
	v2.POST("/activities", middlewares.WithToken(), handlers.HandleCreateActivity)
	v2.PUT("/activities/registration", middlewares.WithToken(), handlers.HandleParticipantsRegistration)

	// groups
	v2.GET("/groups/users/:organizerId", handlers.HandleGetAllGroupsByOrganizer)
	v2.GET("/groups/:id", handlers.HandleGetGroups)
	v2.POST("/groups", handlers.HandleCreateGroup)

	// locales
	v2.GET("/locales", handlers.HandleGetAllLocales)

	// sports
	v2.GET("/sports", handlers.HandleGetAllSports)

	v2.GET("/me", middlewares.WithToken(), handlers.HandleMe)
	v2.POST("/logout", handlers.HandleLogout)
	v2.POST("/login", handlers.HandleLogin)
	v2.POST("/google-login", handlers.HandleLoginWithGoogle)

	// user
	v2.POST("/users", handlers.HandleCreateUser)
	v2.POST("/users/:userId/activity", handlers.HandleCreateUser)
	v2.PUT("/users/:userId/settings", middlewares.WithToken(), handlers.HandleUpdateSettings)
	v2.PUT("/users/:userId/activities/:activityId/favorites", middlewares.WithToken(), handlers.HandleUpdateFavorites)

	err := g.Run(os.Getenv("LISTEN_PORT"))

	if err != nil {
		log.Fatal(err)
	}
}
