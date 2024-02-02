package data

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"me/pickside/db"
	"me/pickside/db/queries"
	"me/pickside/types"
	"net/http"
	"time"
)

type UserSettings struct {
	ID                    uint64 `json:"id"`
	PreferredSport        string `json:"preferred_sport"`
	PreferredLocale       string `json:"preferred_locale"`
	PreferredTheme        string `json:"preferred_theme"`
	PreferredRegion       string `json:"preferred_region"`
	AllowLocationTracking bool   `json:"allow_location_tracking"`
	ShowAge               bool   `json:"show_age"`
	ShowEmail             bool   `json:"show_email"`
	ShowPhone             bool   `json:"show_phone"`
	ShowGroups            bool   `json:"show_groups"`
	User                  User   `json:"user" gorm:"foreignKey:UserID"`
	UserID                User   `json:"user_id"`
}

type User struct {
	ID                  uint64            `json:"id"`
	AccountType         types.AccountType `json:"account_type"`
	Avatar              string
	Bio                 string
	City                string
	Email               string
	EmailVerified       bool       `json:"email_verified"`
	FullName            string     `json:"full_name"`
	IsInactive          bool       `json:"is_inactive"`
	InactiveDate        *time.Time `json:"inactive_date"`
	JoinDate            time.Time  `json:"join_date"`
	LocaleRegion        string     `json:"locale_region"`
	MatchOrganizedCount int        `json:"match_organized_count"`
	MatchPlayedCount    int        `json:"match_played_count"`
	Password            string
	Permissions         string
	Phone               string
	Reliability         int
	Role                types.Role
	Sexe                types.Sexe
	Timezone            string
	Username            string
}

func GetByID(userID uuid.UUID) error {
	return nil
}

func Authenticate(username string, password string) (*User, error, int) {
	dbInstance := db.GetDB()

	var user User

	row, err := dbInstance.Query(queries.SelectPasswordOnlyWhereUsernameEquals, username)
	if err != nil {
		return nil, nil, http.StatusInternalServerError
	}

	for row.Next() {
		var user User
		err := row.Scan(&user.Password)
		if err != nil {
			return nil, nil, http.StatusInternalServerError
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return nil, errors.New("username or password is incorrect"), http.StatusUnauthorized
		}
	}

	row, err = dbInstance.Query(queries.SelectAllColumnsExceptPasswordWhereUsernameEquals, username)
	if err != nil {
		return nil, nil, http.StatusInternalServerError
	}

	for row.Next() {
		fields := []interface{}{
			&user.ID,
			&user.AccountType,
			&user.Avatar,
			&user.Bio,
			&user.City,
			&user.Email,
			&user.EmailVerified,
			&user.FullName,
			&user.IsInactive,
			&user.InactiveDate,
			&user.JoinDate,
			&user.LocaleRegion,
			&user.MatchOrganizedCount,
			&user.MatchPlayedCount,
			&user.Permissions,
			&user.Phone,
			&user.Reliability,
			&user.Role,
			&user.Sexe,
			&user.Timezone,
			&user.Username,
		}

		err := row.Scan(fields...)
		if err != nil {
			return nil, errors.New("error while scanning users"), http.StatusInternalServerError
		}
	}

	return &user, nil, http.StatusOK
}

//func GetMe() ([]User, error) {
//	dbInstance := db.GetDB()
//
//	rows, err := dbInstance.Query(queries.SelectAllFromUsers)
//
//	var users []User
//
//	columns, err := rows.Columns()
//
//	values := make([]interface{}, len(columns))
//	for i := range values {
//		values[i] = new(interface{})
//	}
//
//	for rows.Next() {
//		var user User
//		err = rows.Scan(values...)
//		//if err != nil {
//		//	log.Fatal("(GetUsers) rows.Scan", err)
//		//}
//		users = append(users, user)
//	}
//
//	//for _, user := range users {
//	//	fmt.Printf("ID: %d\n", user.ID)
//	//	fmt.Printf("Full Name: %s\n", user.FullName)
//	//	fmt.Printf("Email: %s\n", user.Email)
//	//	fmt.Printf("Password: %s\n", user.Password)
//	//}
//
//	return users, err
//}
