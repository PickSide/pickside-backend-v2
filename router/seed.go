package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"me/pickside/db/queries"
)

func Seed(c *gin.Context, db *sql.DB) {
	log.Println("Drop tables...")
	_, err := db.Exec(queries.DROP_USER_TABLE)
	if err != nil {
		return
	}

	log.Println("Creating tables")
	_, err = db.Exec(queries.CREATE_USER_TABLE)

	log.Println("Seeding table...")

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("123"), 10)
	if err != nil {
		panic(err)
	}

	users := []User{
		{FullName: "John Doe", Email: "john.doe@example.com", Password: string(hashPassword)},
		{FullName: "Jane Smith", Email: "jane.smith@example.com", Password: string(hashPassword)},
		// Add more users as needed
	}

	for _, user := range users {
		_, err := db.Exec(queries.INSERT_USER, user.FullName, user.Email)
		if err != nil {
			return
		}
	}

	log.Println("Done")

	return
}
