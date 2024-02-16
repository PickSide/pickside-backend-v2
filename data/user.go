package data

import (
	"database/sql"
	"pickside/service/db"
	"pickside/service/db/queries"
	"pickside/service/types"
	"pickside/service/util"
	"strconv"
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
	Favorites           string     `json:"favorites"`
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

	hashedPwd, err := getPasswordFromDB(dbInstance, username)
	if err != nil {
		return nil, err
	}

	err = comparePasswords(hashedPwd, password)
	if err != nil {
		return nil, err
	}

	user, err := getUserDetailsByUsername(dbInstance, username)

	return user, err
}

func MatchUserByEmail(email string, locale string, name string, picture string, verifiedEmail bool) (*User, error) {
	dbInstance := db.GetDB()

	var user User

	err := dbInstance.QueryRow(queries.SelectUserByEmail, email).Scan(
		&user.ID, &user.AccountType, &user.Avatar, &user.Bio, &user.City, &user.Email,
		&user.EmailVerified, &user.FullName, &user.Favorites, &user.IsInactive, &user.InactiveDate, &user.JoinDate, &user.LocaleRegion,
		&user.MatchOrganizedCount, &user.MatchPlayedCount, &user.Permissions, &user.Phone, &user.Reliability, &user.Role,
		&user.Sexe, &user.Timezone, &user.Username, &user.AgreedToTerms,
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
			types.GOOGLE, "", "", "", email, true,
			name, "", false, nil, time.Now(), "", 0, 0, "",
			strings.Join([]string{types.ACTIVITIES_VIEW, types.NOTIFICATIONS_RECEIVE, types.GOOGLE_LOCATION_SEARCH, types.MAP_VIEW}, ","),
			"", 0.00, types.USER, "male", location.String(), username, true,
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
			&insertedUser.ID, &insertedUser.AccountType, &insertedUser.Avatar, &insertedUser.Bio, &insertedUser.City, &insertedUser.Email,
			&insertedUser.EmailVerified, &insertedUser.FullName, &insertedUser.Favorites, &insertedUser.IsInactive, &insertedUser.InactiveDate, &insertedUser.JoinDate,
			&insertedUser.LocaleRegion, &insertedUser.MatchOrganizedCount, &insertedUser.MatchPlayedCount, &insertedUser.Permissions, &insertedUser.Phone,
			&insertedUser.Reliability, &insertedUser.Role, &insertedUser.Sexe, &insertedUser.Timezone, &insertedUser.Username, &insertedUser.AgreedToTerms,
		)

		user = insertedUser
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetMe(id uint64) (*User, error) {
	dbInstance := db.GetDB()

	user, err := getUserDetailsById(dbInstance, id)

	return user, err
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

func GetFavorites(userId uint64) (*[]Activity, error) {
	dbInstance := db.GetDB()

	favorites, err := getUserFavorites(dbInstance, userId)
	if err != nil {
		return nil, err
	}

	activities, err := getActivitiesByIds(dbInstance, strings.Join(favorites, ","))
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func UpdateFavorites(userId uint64, activityId uint64) (*sql.Result, error) {
	dbInstance := db.GetDB()

	activityIdStr := strconv.FormatUint(activityId, 10)

	favorites, err := getUserFavorites(dbInstance, userId)
	if err != nil {
		return nil, err
	}

	idx := isFavorite(favorites, activityIdStr)

	if idx != -1 {
		favorites = append(favorites[:idx], favorites[idx+1:]...)
	} else {
		if len(favorites) == 0 || favorites[0] != "" {
			favorites = append(favorites, activityIdStr)
		} else {
			favorites[0] = activityIdStr
		}
	}

	tx, err := dbInstance.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	result, err := dbInstance.Exec(queries.UpdateFavorites, strings.Join(favorites, ","), userId)

	return &result, err
}

func getPasswordFromDB(dbInstance *sql.DB, username string) (string, error) {
	var hashedPwd string
	err := dbInstance.QueryRow(queries.SelectPasswordOnly, username).Scan(&hashedPwd)
	return hashedPwd, err
}

func comparePasswords(hashedPwd, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
}

func getUserDetailsByUsername(dbInstance *sql.DB, username string) (*User, error) {
	var user User
	err := dbInstance.QueryRow(queries.SelectByUsername, username).Scan(
		&user.ID, &user.AccountType, &user.Avatar, &user.Bio, &user.City, &user.Email,
		&user.EmailVerified, &user.FullName, &user.Favorites, &user.IsInactive, &user.InactiveDate,
		&user.JoinDate, &user.LocaleRegion, &user.MatchOrganizedCount,
		&user.MatchPlayedCount, &user.Permissions, &user.Phone, &user.Reliability,
		&user.Role, &user.Sexe, &user.Timezone, &user.Username, &user.AgreedToTerms,
	)

	return &user, err
}

func getUserDetailsById(dbInstance *sql.DB, userId uint64) (*User, error) {
	var user User
	err := dbInstance.QueryRow(queries.SelectUserById, userId).Scan(
		&user.ID, &user.AccountType, &user.Avatar, &user.Bio, &user.City, &user.Email,
		&user.EmailVerified, &user.FullName, &user.Favorites, &user.IsInactive, &user.InactiveDate,
		&user.JoinDate, &user.LocaleRegion, &user.MatchOrganizedCount,
		&user.MatchPlayedCount, &user.Permissions, &user.Phone, &user.Reliability,
		&user.Role, &user.Sexe, &user.Timezone, &user.Username, &user.AgreedToTerms,
	)

	return &user, err
}

func getUserFavorites(dbInstance *sql.DB, userId uint64) ([]string, error) {
	var favorites string
	err := dbInstance.QueryRow(queries.SelectFavorites, userId).Scan(&favorites)

	return strings.Split(favorites, ","), err
}

func isFavorite(array []string, value string) int {
	for i, unit := range array {
		if unit == value {
			return i
		}
	}

	return -1
}

func getActivitiesByIds(dbInstance *sql.DB, activityIds string) (*[]Activity, error) {
	activitiesId := strings.Split(activityIds, ",")

	var activities []Activity

	for _, activityId := range activitiesId {
		var activity Activity

		err := dbInstance.QueryRow(queries.SelectActivityById, activityId).Scan(
			&activity.ID, &activity.Address, &activity.Date, &activity.Description, &activity.IsPrivate,
			&activity.MaxPlayers, &activity.Price, &activity.Rules, &activity.Time, &activity.Title,
			&activity.CreatedAt, &activity.UpdatedAt, &activity.OrganizerID, &activity.SportID,
		)
		if err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}

	return &activities, nil
}
