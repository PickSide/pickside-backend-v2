package main

import (
	"log"
	"pickside/service/data"
	"pickside/service/db"
	"pickside/service/db/queries"
	"pickside/service/types"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
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

	PopulateTables()
}

func PopulateTables() {
	log.Println("Seeding table...")

	localesData := []data.Locale{
		{Name: "english", Value: "en", FlagCode: "us"},
		{Name: "français", Value: "fr", FlagCode: "fr"},
	}

	for _, locale := range localesData {
		_, err := db.GetDB().Query(queries.InsertIntoLocale, locale.Name, locale.Value, locale.FlagCode)
		if err != nil {
			panic(err)
		}
	}

	sportsData := []data.Sport{
		{Name: "soccer", GameModes: "5 aside,7 aside,8 aside,11 aside", FeatureAvailable: true},
		{Name: "basketball", GameModes: "1 on 1, 3 on 3, 5 on 5", FeatureAvailable: false},
		{Name: "tennis", GameModes: "", FeatureAvailable: false},
		{Name: "american_football", GameModes: "", FeatureAvailable: false},
	}

	for _, sport := range sportsData {
		_, err := db.GetDB().Query(queries.InsertIntoSport, sport.Name, sport.GameModes, sport.FeatureAvailable)
		if err != nil {
			panic(err)
		}
	}

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
			AccountType:           types.DEFAULT,
			AgreedToTerms:         true,
			AllowLocationTracking: false,
			Email:                 user.Email,
			EmailVerified:         true,
			FullName:              user.FullName,
			InactiveDate:          nil,
			IsInactive:            false,
			JoinDate:              time.Now(),
			MatchOrganizedCount:   0,
			MatchPlayedCount:      0,
			Password:              string(hashPassword),
			Permissions:           strings.Join(types.DEFAULT_PERMISSIONS[:], ","),
			Phone:                 "514-123-45679",
			PreferredLocale:       "en",
			PreferredRegion:       "montreal",
			PreferredSport:        "soccer",
			PreferredTheme:        "light",
			Reliability:           50,
			Role:                  types.USER,
			Sexe:                  types.MALE,
			ShowAge:               true,
			ShowEmail:             true,
			ShowGroups:            false,
			ShowPhone:             false,
			Timezone:              location.String(),
			Username:              user.Username,
		})
	}

	for _, user := range users {
		_, err := db.GetDB().Exec(
			`INSERT INTO users (
				account_type, agreed_to_terms, allow_location_tracking, avatar, bio, city, email, email_verified, external_id, favorites,
				full_name, inactive_date, is_inactive, join_date, locale_region, match_organized_count, match_played_count,
				password, permissions, phone, preferred_locale, preferred_region, preferred_sport, preferred_theme,
				reliability, role, sexe, show_age, show_email, show_groups, show_phone, timezone, username
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			user.AccountType, true, user.AllowLocationTracking, user.Avatar, user.Bio, user.City, user.Email, user.EmailVerified, user.ExternalID, user.Favorites,
			user.FullName, user.InactiveDate, user.IsInactive, user.JoinDate, user.LocaleRegion, user.MatchOrganizedCount, user.MatchPlayedCount,
			user.Password, user.Permissions, user.Phone, user.PreferredLocale, user.PreferredRegion, user.PreferredSport, user.PreferredTheme,
			user.Reliability, user.Role, user.Sexe, user.ShowAge, user.ShowEmail, user.ShowGroups, user.ShowPhone, user.Timezone, user.Username,
		)
		if err != nil {
			panic(err)
		}
	}

	activitiesData := []data.Activity{
		{Address: "123 rue du 33", Date: time.Now().Format("2006-01-02"), Description: "unknown description", GameMode: "5 aside", IsPrivate: false, Lat: 9.19756, Lng: 29.67629, MaxPlayers: 11, Price: 0, Rules: "No tackles", OrganizerID: 1, Time: time.Now().Format("15:04:05"), Title: "Activity A", SportID: 1},
		{Address: "123 rue du 33", Date: time.Now().Format("2006-01-02"), Description: "unknown description", GameMode: "5 aside", IsPrivate: false, Lat: 10.92810, Lng: -8.07624, MaxPlayers: 11, Price: 0, Rules: "No tackles", OrganizerID: 1, Time: time.Now().Format("15:04:05"), Title: "Activity A", SportID: 1},
		{Address: "123 rue du 33", Date: time.Now().Format("2006-01-02"), Description: "unknown description", GameMode: "5 aside", IsPrivate: false, Lat: 4.70038, Lng: -77.28465, MaxPlayers: 11, Price: 0, Rules: "No tackles", OrganizerID: 1, Time: time.Now().Format("15:04:05"), Title: "Activity A", SportID: 1},
		{Address: "123 rue du 33", Date: time.Now().Format("2006-01-02"), Description: "unknown description", GameMode: "5 aside", IsPrivate: false, Lat: 19.64322, Lng: -89.74694, MaxPlayers: 11, Price: 0, Rules: "No tackles", OrganizerID: 1, Time: time.Now().Format("15:04:05"), Title: "Activity A", SportID: 1},
		{Address: "123 rue du 33", Date: time.Now().Format("2006-01-02"), Description: "unknown description", GameMode: "5 aside", IsPrivate: false, Lat: -8.90166, Lng: -74.75712, MaxPlayers: 11, Price: 0, Rules: "No tackles", OrganizerID: 1, Time: time.Now().Format("15:04:05"), Title: "Activity A", SportID: 1},
	}

	for _, activity := range activitiesData {
		_, err := db.GetDB().Exec(queries.InsertActivity,
			activity.Address,
			activity.Date,
			activity.Description,
			activity.GameMode,
			activity.Images,
			activity.IsPrivate,
			activity.Lat,
			activity.Lng,
			activity.MaxPlayers,
			activity.OrganizerID,
			activity.Price,
			activity.Rules,
			activity.SportID,
			activity.Time,
			activity.Title,
		)
		if err != nil {
			panic(err)
		}
	}

	log.Println("Done")
}
