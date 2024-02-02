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

	g.Use()

	g.GET("/me", middlewares.WithToken(), handlers.HandleMe)
	g.POST("/login", handlers.HandleLogin)
	g.POST("/users", handlers.HandleCreateMe)

	err := g.Run(os.Getenv("LISTEN_PORT"))

	if err != nil {
		log.Fatal(err)
	}
}
