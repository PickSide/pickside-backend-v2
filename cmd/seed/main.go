package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"me/pickside/data"
	"me/pickside/db"
	"me/pickside/db/queries"
	"me/pickside/types"
	"strings"
	"time"
)

type UserData struct {
	FullName string
	Email    string
	Username string
}

func main() {
	if err := db.OpenConnectionToDB(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.GetDB().Close()
		if err != nil {
			panic(err)
		}
	}()

	CreateTables()
	PopulateTables()
}

func CreateTables() {
	log.Println("Creating tables...")

	qs := []string{
		queries.CreateUserTables,
		queries.CreateUserSettingsTable,
	}

	for _, q := range qs {
		_, err := db.GetDB().Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

func PopulateTables() {
	log.Println("Seeding table...")
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("123"), 10)
	if err != nil {
		panic(err)
	}

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	usersData := []UserData{
		{FullName: "Tony Hakim", Email: "tonyown10@example.com", Username: "tony"},
		{FullName: "Niloo Khastavan", Email: "niloo@example.com", Username: "niloo"},
		{FullName: "Rafic Haddad", Email: "rafic10@example.com", Username: "rafic"},
		{FullName: "Ali Idrici", Email: "ali10@example.com", Username: "ali"},
		{FullName: "Ian Piluganov", Email: "ian10@example.com", Username: "ian"},
		{FullName: "Kevin Moniz", Email: "kevin10@example.com", Username: "kevin"},
	}

	var users []data.User

	for _, user := range usersData {
		users = append(users, data.User{
			AccountType:         types.DEFAULT,
			Avatar:              "",
			Bio:                 "My bio",
			City:                "Montreal",
			Password:            string(hashPassword),
			EmailVerified:       true,
			FullName:            user.FullName,
			Email:               user.Email,
			Username:            user.Username,
			IsInactive:          false,
			InactiveDate:        nil,
			JoinDate:            time.Now(),
			LocaleRegion:        "montreal",
			MatchOrganizedCount: 0,
			MatchPlayedCount:    0,
			Permissions:         strings.Join([]string{types.ACTIVITIES_VIEW, types.NOTIFICATIONS_RECEIVE, types.GOOGLE_LOCATION_SEARCH, types.MAP_VIEW}, ","),
			Phone:               "514-123-45679",
			Reliability:         50,
			Role:                types.USER,
			Sexe:                types.MALE,
			Timezone:            location.String(),
		})
	}

	for i, user := range users {
		_, err := db.GetDB().Exec(queries.InsertUser,
			user.AccountType,
			user.Avatar,
			user.Bio,
			user.City,
			user.Email,
			user.EmailVerified,
			user.FullName,
			user.IsInactive,
			user.InactiveDate,
			user.JoinDate,
			user.LocaleRegion,
			user.MatchOrganizedCount,
			user.MatchPlayedCount,
			user.Password,
			user.Permissions,
			user.Phone,
			user.Reliability,
			user.Role,
			user.Sexe,
			user.Timezone,
			user.Username,
		)
		if err != nil {
			panic(err)
		}

		_, err = db.GetDB().Exec(queries.InsertUserSetting,
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
}
