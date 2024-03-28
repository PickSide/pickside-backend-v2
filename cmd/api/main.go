package main

import (
	"fmt"
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

	if err := db.OpenConnectionToDB(); err != nil {
		log.Fatal(err)
	}

	g := gin.Default()
	g.Use(cors.New(buildCors()))

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
	v2.DELETE("/groups/:groupId/users/:organizerId", handlers.HandleDeleteGroup)

	// locales
	v2.GET("/locales", handlers.HandleGetAllLocales)

	// sports
	v2.GET("/sports", handlers.HandleGetAllSports)

	v2.GET("/me", middlewares.WithToken(), handlers.HandleMe)
	v2.POST("/logout", handlers.HandleLogout)
	v2.POST("/login", handlers.HandleLogin)
	v2.POST("/google-login", handlers.HandleLoginWithGoogle)

	// user
	v2.GET("/users", middlewares.WithToken(), handlers.HandleGetAllUsers)
	v2.GET("/users/:userId/activities/favorites", middlewares.WithToken(), handlers.HandleGetFavorites)
	v2.POST("/users", handlers.HandleCreateUser)
	v2.PUT("/users/:userId/settings", middlewares.WithToken(), handlers.HandleUpdateUser)
	v2.PUT("/users/:userId/activities/:activityId/favorites", middlewares.WithToken(), handlers.HandleUpdateFavorites)

	if err := g.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT"))); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func buildCors() cors.Config {
	c := cors.DefaultConfig()
	c.AllowOrigins = []string{"http://127.0.0.1:3000", "http://localhost:3000", "https://pickside.net"}
	c.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	c.AllowHeaders = []string{"Origin", "X-Request-Id", "Content-Type"}
	c.AllowWebSockets = false
	c.AllowCredentials = true
	c.MaxAge = 12 * time.Hour

	return c
}
