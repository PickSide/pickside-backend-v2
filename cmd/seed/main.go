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

	CreateTables()
	PopulateTables()
}

func CreateTables() {
	log.Println("Creating tables...")

	qs := []string{
		queries.CreateActivityTable,
		queries.CreateActivityUserTable,
		queries.CreateChatroomParticipantsTable,
		queries.CreateChatroomTable,
		queries.CreateGameModesTable,
		queries.CreateGroupUsersTable,
		queries.CreateGroupTable,
		queries.CreateLocaleTable,
		queries.CreateMessageTable,
		queries.CreateNotificationTable,
		queries.CreateSportGameModesTable,
		queries.CreateSportTable,
		queries.CreateTokensTable,
		queries.CreateUserSettingsTable,
		queries.CreateUserTables,
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

	localesData := []data.Locale{
		{Name: "english", FlagCode: "en"},
		{Name: "français", FlagCode: "fr"},
	}

	for _, locale := range localesData {
		_, err := db.GetDB().Query(queries.InsertIntoLocale, locale.Name, locale.FlagCode)
		if err != nil {
			panic(err)
		}
	}

	sportsData := []data.Sport{
		{Name: "soccer", FeatureAvailable: true},
		{Name: "basketball", FeatureAvailable: false},
		{Name: "tennis", FeatureAvailable: false},
		{Name: "american_football", FeatureAvailable: false},
	}

	for _, sport := range sportsData {
		_, err := db.GetDB().Query(queries.InsertIntoSport, sport.Name, sport.FeatureAvailable)
		if err != nil {
			panic(err)
		}
	}

	gameModesData := []data.GameMode{
		{Name: "1v1"},
		{Name: "5v5"},
		{Name: "7v7"},
		{Name: "8v8"},
		{Name: "11v11"},
	}

	for _, gameMode := range gameModesData {
		_, err := db.GetDB().Query(queries.InsertIntoGameMode, gameMode.Name)
		if err != nil {
			panic(err)
		}
	}

	sportGameModesData := []struct {
		GameModeID uint64
		SportID    uint64
	}{
		{GameModeID: 1, SportID: 1}, // soccer
		{GameModeID: 2, SportID: 1}, // soccer
		{GameModeID: 3, SportID: 1}, // soccer
		{GameModeID: 4, SportID: 1}, // soccer

		{GameModeID: 1, SportID: 2}, // bball
		{GameModeID: 2, SportID: 2}, // bball
	}

	for _, sportGameMode := range sportGameModesData {
		_, err := db.GetDB().Query(queries.InsertIntoSportGameMode, sportGameMode.GameModeID, sportGameMode.SportID)
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
			i,
		)
		if err != nil {
			panic(err)
		}
	}
	activitiesData := []data.Activity{
		{Address: "123 rue du 33", Date: time.Now().Format("2006-01-02"), Description: "unknown description", IsPrivate: false, MaxPlayers: 11, Price: 0, Rules: "No tackles", OrganizerID: 1, Time: time.Now().Format("15:04:05"), Title: "Activity A", SportID: 1},
		{Address: "123 rue du 34", Date: time.Now().Format("2006-01-02"), Description: "unknown description", IsPrivate: false, MaxPlayers: 22, Price: 5, Rules: "No tackles", OrganizerID: 2, Time: time.Now().Format("15:04:05"), Title: "Activity B", SportID: 1},
		{Address: "123 rue du 35", Date: time.Now().Format("2006-01-02"), Description: "unknown description", IsPrivate: false, MaxPlayers: 22, Price: 10, Rules: "No tackles", OrganizerID: 3, Time: time.Now().Format("15:04:05"), Title: "Activity C", SportID: 1},
		{Address: "123 rue du 36", Date: time.Now().Format("2006-01-02"), Description: "unknown description", IsPrivate: false, MaxPlayers: 22, Price: 0, Rules: "No tackles", OrganizerID: 4, Time: time.Now().Format("15:04:05"), Title: "Activity D", SportID: 1},
		{Address: "123 rue du 37", Date: time.Now().Format("2006-01-02"), Description: "unknown description", IsPrivate: false, MaxPlayers: 22, Price: 0, Rules: "No tackles", OrganizerID: 5, Time: time.Now().Format("15:04:05"), Title: "Activity E", SportID: 1},
		{Address: "123 rue du 38", Date: time.Now().Format("2006-01-02"), Description: "unknown description", IsPrivate: false, MaxPlayers: 22, Price: 0, Rules: "No tackles", OrganizerID: 1, Time: time.Now().Format("15:04:05"), Title: "Activity F", SportID: 1},
	}

	for _, activity := range activitiesData {
		_, err := db.GetDB().Query(queries.InsertActivity,
			activity.Address,
			activity.Date,
			activity.Description,
			activity.IsPrivate,
			activity.MaxPlayers,
			activity.Price,
			activity.Rules,
			activity.OrganizerID,
			activity.Time,
			activity.Title,
			activity.SportID,
		)
		if err != nil {
			panic(err)
		}
	}

	log.Println("Done")
}