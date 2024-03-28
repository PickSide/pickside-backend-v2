package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pickside/service/db"
	qs "pickside/service/pkg"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.OpenConnectionToDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	server := qs.NewServer()

	g := gin.Default()

	g.Use(cors.New(buildCors()))

	g.GET("/ws", func(c *gin.Context) {
		handleWS(c.Writer, c.Request, server)
	})

	if err := g.Run(fmt.Sprintf(":%s", os.Getenv("QS_PORT"))); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func handleWS(w http.ResponseWriter, r *http.Request, server *qs.Server) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		return
	}

	server.HandleWS(conn)
}

func buildCors() cors.Config {
	c := cors.DefaultConfig()
	c.AllowOrigins = []string{"*"}
	c.AllowMethods = []string{"GET", "POST"}
	c.AllowHeaders = []string{"Origin", "X-Request-Id", "Content-Type"}
	c.AllowWebSockets = true
	c.MaxAge = 12 * time.Hour

	return c
}
