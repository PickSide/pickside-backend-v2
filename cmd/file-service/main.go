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
		"path":    fmt.Sprintf("%s/%s", os.Getenv("CDN"), *path),
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.OpenConnectionToDB(); err != nil {
		log.Fatal(err)
	}

	g := gin.Default()
	g.Use(cors.New(buildCors()))

	fs := g.Group("/image-service", middlewares.FromValidDomain())

	fs.GET("/download", handleDownloadFile)
	fs.POST("/:path/:keyId", middlewares.ValidateS3Path(), handleUploadFile)

	if err := g.Run(fmt.Sprintf(":%s", os.Getenv("FS_PORT"))); err != nil {
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
