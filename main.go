package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"me/pickside/router"
	"os"
)

var db *sql.DB

func main() {
	OpenConnectionToDB()

	r := gin.Default()

	r.GET("/seed", func(c *gin.Context) {
		router.Seed(c, db)
	})
	r.GET("/users", func(c *gin.Context) {
		router.GetUsers(c, db)
	})
	r.POST("/users", func(c *gin.Context) {
		router.CreateUser(c, db)
	})

	err := r.Run()

	if err != nil {
		return
	}
}

func OpenConnectionToDB() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err = sql.Open("mysql", os.Getenv("DSN"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("successfully connected to db")
}
