package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"me/pickside/db"
	"me/pickside/handlers"
	"me/pickside/middlewares"
	"os"
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

	v1.GET("/me", middlewares.WithToken(), handlers.HandleMe)
	v1.GET("/logout", handlers.HandleLogout)
	v1.POST("/login", handlers.HandleLogin)
	v1.POST("/users", handlers.HandleCreateMe)

	err := g.Run(os.Getenv("LISTEN_PORT"))

	if err != nil {
		log.Fatal(err)
	}
}
