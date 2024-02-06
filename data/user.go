package data

import (
	"me/pickside/db"
	"me/pickside/db/queries"
	"me/pickside/types"
	"time"

	"golang.org/x/crypto/bcrypt"
)

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
	UserSettings        UserSettings `json:"settings" gorm:"foreignKey:UserID"`
}

func MatchUser(username string, password string) (*User, error) {
	dbInstance := db.GetDB()

	row, err := dbInstance.Query(queries.SelectPasswordOnlyWhereUsernameEquals, username)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var hashedPwd string
		err := row.Scan(&hashedPwd)
		if err != nil {
			return nil, err
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
		if err != nil {
			return nil, err
		}
	}

	var user User

	err = dbInstance.QueryRow(queries.SelectAllColumnsExceptPasswordWhereUsernameEquals, username).Scan(
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
	)

	return &user, err
}

func GetMe(id uint64) (*User, error) {
	dbInstance := db.GetDB()

	var user User

	err := dbInstance.QueryRow(queries.SelectAllColumnsExceptPasswordWhereIDEquals, id).Scan(
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
	)

	return &user, err
}

func InsertUser(fullName string, username string, email string, pwd string, phone string, sexe string, agreedToTerms bool) error {
	dbInstance := db.GetDB()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return err
	}

	_, err = dbInstance.Query(queries.InsertUser, fullName, username, email, hashedPwd, phone, sexe, agreedToTerms)

	return err
}
