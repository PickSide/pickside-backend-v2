package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"me/pickside/db/queries"
	"strings"
	"time"
)

type UserData struct {
	FullName string
	Email    string
	Username string
}

func DropTables(db *sql.DB) {
	log.Println("Dropping tables...")
	_, err := db.Exec(queries.DropUserTables)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(queries.DropUserSettingsTable)
	if err != nil {
		panic(err)
	}
}

func CreateTables(db *sql.DB) {
	log.Println("Creating tables...")
	_, err := db.Exec(queries.CreateUserTables)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(queries.CreateUserSettingsTable)
	if err != nil {
		panic(err)
	}
}

func Seed(c *gin.Context, db *sql.DB) {
	DropTables(db)
	CreateTables(db)

	log.Println("Seeding table...")
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("123"), 10)
	if err != nil {
		panic(err)
	}

	location, err := time.LoadLocation("America/New_York")
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
		users = append(users, User{
			AccountType:         DEFAULT,
			Avatar:              "",
			Bio:                 "My bio",
			City:                "Montreal",
			Password:            string(hashPassword),
			EmailVerified:       true,
			FullName:            user.FullName,
			Email:               user.Email,
			Username:            user.Username,
			IsInactive:          false,
			InactiveDate:        time.Now(),
			JoinDate:            time.Now(),
			LocaleRegion:        "montreal",
			MatchOrganizedCount: 0,
			MatchPlayedCount:    0,
			Permissions:         []string{ACTIVITIES_VIEW, NOTIFICATIONS_RECEIVE, GOOGLE_LOCATION_SEARCH, MAP_VIEW},
			Phone:               "514-123-45679",
			Reliability:         50,
			Role:                USER,
			Sexe:                MALE,
			Timezone:            location,
		})
	}

	for i, user := range users {
		_, err := db.Exec(queries.InsertUser,
			user.AccountType,
			user.Avatar,
			user.Bio,
			user.City,
			user.Email,
			user.EmailVerified,
			user.FullName,
			user.Password,
			user.Username,
			user.IsInactive,
			user.InactiveDate,
			user.JoinDate,
			user.LocaleRegion,
			user.MatchOrganizedCount,
			user.MatchPlayedCount,
			strings.Join([]string(user.Permissions), ","),
			user.Phone,
			user.Reliability,
			user.Role,
			user.Sexe,
			user.Timezone.String(),
		)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(queries.InsertUserSetting,
			"soccer",   // preferred_sport
			"en",       // preferred_locale
			"light",    // preferred_theme
			"montreal", // preferred_region
			false,      // allow_location_tracking
			true,       // show_age
			true,       // show_email
			false,      // show_phone,
			false,      // show_groups,
			i,          // user_id
		)
		if err != nil {
			panic(err)
		}
	}

	log.Println("Done")

	return
}
