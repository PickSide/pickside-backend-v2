package data

import (
	"database/sql"
	"pickside/service/db"
	"pickside/service/db/queries"
	"pickside/service/types"
	"pickside/service/util"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                  uint64            `json:"id"`
	AccountType         types.AccountType `json:"accountType"`
	Avatar              string            `json:"avatar" binding:"omitempty"`
	Bio                 string            `json:"bio" binding:"omitempty"`
	City                string            `json:"city" binding:"omitempty"`
	Email               string
	EmailVerified       bool       `json:"emailVerified"`
	FullName            string     `json:"fullName"`
	IsInactive          bool       `json:"isInactive"`
	InactiveDate        *time.Time `json:"inactiveDate"`
	JoinDate            time.Time  `json:"joinDate"`
	LocaleRegion        string     `json:"localeRegion"`
	MatchOrganizedCount int        `json:"matchOrganizedCount"`
	MatchPlayedCount    int        `json:"matchPlayedCount"`
	Password            string
	Permissions         string
	Phone               string
	Reliability         int
	Role                types.Role
	Sexe                types.Sexe
	Timezone            string
	Username            string
	AgreedToTerms       bool         `json:"agreedToTerms"`
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

func MatchUserByEmail(email string, locale string, name string, picture string, verifiedEmail bool) (*User, error) {
	dbInstance := db.GetDB()

	var user User

	err := dbInstance.QueryRow(queries.SelectUserByEmail, email).Scan(
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
		&user.AgreedToTerms,
	)

	if err == sql.ErrNoRows {
		username, err := util.GenerateRandomUsername(10)
		if err != nil {
			return nil, err
		}

		location, err := time.LoadLocation("America/New_York")
		if err != nil {
			panic(err)
		}

		result, err := dbInstance.Exec(queries.InsertUser,
			types.GOOGLE,
			"",
			"",
			"",
			email,
			true,
			name,
			false,      //is inactive
			nil,        // inactive date
			time.Now(), // join date
			"",         // locale region
			0,
			0,
			"",
			strings.Join([]string{types.ACTIVITIES_VIEW, types.NOTIFICATIONS_RECEIVE, types.GOOGLE_LOCATION_SEARCH, types.MAP_VIEW}, ","),
			"",
			0.00,
			types.USER,
			"male",
			location.String(),
			username,
			true,
		)
		if err != nil {
			return nil, err
		}

		lastInsertID, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		var insertedUser User

		err = dbInstance.QueryRow(
			queries.SelectUserById,
			lastInsertID,
		).Scan(
			&insertedUser.ID,
			&insertedUser.AccountType,
			&insertedUser.Avatar,
			&insertedUser.Bio,
			&insertedUser.City,
			&insertedUser.Email,
			&insertedUser.EmailVerified,
			&insertedUser.FullName,
			&insertedUser.IsInactive,
			&insertedUser.InactiveDate,
			&insertedUser.JoinDate,
			&insertedUser.LocaleRegion,
			&insertedUser.MatchOrganizedCount,
			&insertedUser.MatchPlayedCount,
			&insertedUser.Permissions,
			&insertedUser.Phone,
			&insertedUser.Reliability,
			&insertedUser.Role,
			&insertedUser.Sexe,
			&insertedUser.Timezone,
			&insertedUser.Username,
			&insertedUser.AgreedToTerms,
		)

		user = insertedUser
	} else if err != nil {
		return nil, err
	}

	return &user, nil
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

func InsertUser(fullName string, username string, email string, pwd string, phone string, sexe string, agreedToTerms bool, verifiedEmail bool) error {
	dbInstance := db.GetDB()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return err
	}

	_, err = dbInstance.Query(queries.InsertUser,
		fullName,
		username,
		email,
		hashedPwd,
		phone,
		sexe,
		false,
		agreedToTerms,
	)

	return err
}
