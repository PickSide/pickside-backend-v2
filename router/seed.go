package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"me/pickside/db/queries"
)

type UserData struct {
	FullName string
	Email    string
	Username string
}

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

	users_data := []UserData{
		{FullName: "Tony Hakim", Email: "tonyown10@example.com", Username: "tony"},
		{FullName: "Niloo Khastavan", Email: "niloo@example.com", Username: "niloo"},
		{FullName: "Rafic Haddad", Email: "tonyown10@example.com", Username: "rafic"},
		{FullName: "Ali Idrici", Email: "tonyown10@example.com", Username: "ali"},
		{FullName: "Ian Piluganov", Email: "tonyown10@example.com", Username: "ian"},
		{FullName: "Kevin Moniz", Email: "tonyown10@example.com", Username: "kevin"},
	}

	var users []User

	for _, user := range users_data {
		users = append(users, User{AccountType: DEFAULT, Avatar: "", Bio: "My bio", City: "Montreal", Password: string(hashPassword), EmailVerified: true, FullName: user.FullName, Email: user.Email, Username: user.Username})
	}

	for _, user := range users {
		_, err := db.Exec(queries.INSERT_USER,
			user.AccountType,
			user.Avatar,
			user.Bio,
			user.City,
			user.Email,
			user.EmailVerified,
			user.FullName,
			user.Password,
			user.Username,
		)
		if err != nil {
			panic(err)
		}
	}

	log.Println("Done")

	return
}
