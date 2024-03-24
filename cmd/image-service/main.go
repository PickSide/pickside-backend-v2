package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"pickside/service/db"
	"pickside/service/middlewares"
	s3client "pickside/service/s3"
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
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://localhost:3000", "https://pickside.net", "https://pickside.net/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "X-Request-Id", "Content-Type"},
		AllowWebSockets:  true,
		ExposeHeaders:    []string{"X-Request-Id", "Content-Type", "Content-Length", "Content-Encoding"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v2 := g.Group("/api/v2/image-service", middlewares.FromValidDomain())

	v2.GET("/download", handleDownloadFile)
	v2.POST("/:path/:keyId", middlewares.ValidateS3Path(), handleUploadFile)

	err := g.Run(os.Getenv("LISTEN_PORT"))

	if err != nil {
		log.Fatal(err)
	}
}

type FileRequest struct {
	Path string `json:"path"`
}

func handleDownloadFile(g *gin.Context) {
	var req FileRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body, contentType, err := s3client.DownloadFile(req.Path)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer body.Close()

	_, filename := filepath.Split(req.Path)
	contentDisposition := fmt.Sprintf("inline; filename=\"%s\"", filename)

	g.Header("Content-Disposition", contentDisposition)
	g.Header("Content-Type", *contentType)
	g.DataFromReader(http.StatusOK, -1, *contentType, body, nil)
}

func handleUploadFile(g *gin.Context) {
	dir := g.Param("path")
	keyId := g.Param("keyId")
	file, err := g.FormFile("file")
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer openedFile.Close()

	path, err := s3client.UploadContentToDirectory(dir, keyId, file.Filename, openedFile)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "image uploaded successfully",
		"path":    &path,
	})
}
