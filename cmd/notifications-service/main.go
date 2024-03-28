package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pickside/service/db"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type NotificationRequest struct {
	Type       string          `json:"type,required"`
	Content    string          `json:"content"`
	Recipients []string        `json:"recipients,omitempty"`
	Title      string          `json:"title,omitempty"`
	Priority   string          `json:"priority,omitempty"`
	MetaData   json.RawMessage `json:"metaData,omitempty"`
}

func pushGlobalNotification(g *gin.Context) {
	log.Printf("http://localhost:%s", os.Getenv("QS_PORT"))
	g.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("http://localhost:%s/publish", os.Getenv("QS_PORT")))
}

func handleIncomingNotification(g *gin.Context) {
	var req NotificationRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		spew.Dump(req)
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var query string
	var queryParams []interface{}

	query = `INSERT INTO notifications (type, content, recipients, title, priority, metaData, expires) VALUES (?, ?, ?, ?, ?, ?, ?)`
	queryParams = append(queryParams, req.Type, req.Content, strings.Join(req.Recipients, ","), req.Title, req.Priority, req.MetaData, time.Now().AddDate(0, 0, 7))

	// if _, err := db.GetDB().Exec(query, queryParams...); err != nil {
	// 	g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
	// 	return
	// }

	spew.Dump(query)

	g.JSON(http.StatusOK, gin.H{"status": "success"})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	if err := db.OpenConnectionToDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	g := gin.Default()
	g.Use(cors.New(buildCors()))

	ns := g.Group("/notifications")

	ns.POST("/", pushGlobalNotification)

	if err := g.Run(fmt.Sprintf(":%s", os.Getenv("NS_PORT"))); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func buildCors() cors.Config {
	c := cors.DefaultConfig()
	c.AllowOrigins = []string{"http://127.0.0.1:3000", "http://localhost:3000", "https://pickside.net"}
	c.AllowMethods = []string{"GET", "POST"}
	c.AllowHeaders = []string{"Origin", "X-Request-Id", "Content-Type"}
	c.AllowWebSockets = false
	c.AllowCredentials = true
	c.MaxAge = 12 * time.Hour

	return c
}
