package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func OpenConnectionToDB() error {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	once.Do(func() {
		db, err = sql.Open("mysql", os.Getenv("DSN"))
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}
		log.Println("Successfully connected to db")
	})
	return nil
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error closing the database:", err)
		}
	}
	log.Println("Disconnected from db")
}

func GetDB() *sql.DB {
	return db
}
