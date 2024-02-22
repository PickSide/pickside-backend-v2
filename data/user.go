package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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
	ID                    uint64            `json:"id"`
	AccountType           types.AccountType `json:"accountType,omitempty"`
	AgreedToTerms         bool              `json:"agreedToTerms"`
	AllowLocationTracking bool              `json:"allowLocationTracking"`
	Avatar                *string           `json:"avatar,omitempty"`
	Bio                   *string           `json:"bio,omitempty"`
	City                  *string           `json:"city,omitempty"`
	Email                 string            `json:"email"`
	EmailVerified         bool              `json:"emailVerified"`
	ExternalID            *string           `json:"externalId"`
	Favorites             *string           `json:"favorites,omitempty"`
	FullName              string            `json:"fullName,omitempty"`
	InactiveDate          *time.Time        `json:"inactiveDate,omitempty"`
	IsInactive            bool              `json:"isInactive"`
	JoinDate              time.Time         `json:"joinDate,omitempty"`
	LocaleRegion          *string           `json:"localeRegion,omitempty"`
	MatchOrganizedCount   int               `json:"matchOrganizedCount"`
	MatchPlayedCount      int               `json:"matchPlayedCount"`
	Password              string            `json:"-"`
	Permissions           string            `json:"permissions,omitempty"`
	Phone                 string            `json:"phone,omitempty"`
	PreferredLocale       string            `json:"preferredLocale,omitempty"`
	PreferredRegion       string            `json:"preferredRegion,omitempty"`
	PreferredSport        string            `json:"preferredSport,omitempty"`
	PreferredTheme        string            `json:"preferredTheme,omitempty"`
	Reliability           int               `json:"reliability"`
	Role                  types.Role        `json:"role,omitempty"`
	Sexe                  types.Sexe        `json:"sexe,omitempty"`
	ShowAge               bool              `json:"showAge"`
	ShowEmail             bool              `json:"showEmail"`
	ShowGroups            bool              `json:"showGroups"`
	ShowPhone             bool              `json:"showPhone"`
	Timezone              string            `json:"timezone,omitempty"`
	Username              string            `json:"username,omitempty"`
}

type MatchUserStruct struct {
	Username string
	Password string
}

func MatchUsername(username string, password string) (*User, error) {
	dbInstance := db.GetDB()

	match := passwordMatch(dbInstance, "username", username, password)
	log.Println("MatchUsername - match", match)
	if !match {
		return nil, errors.New("Username or email is incorrect")
	}

	user, err := getUserDetails(dbInstance, "username", username)

	return user, err
}

type MatchEmailStruct struct {
	Email    string
	Password string
}

func MatchEmail(email string, password string) (*User, error) {
	dbInstance := db.GetDB()

	var user *User
	var err error

	if passwordMatch(dbInstance, "email", email, password) {
		user, err = getUserDetails(dbInstance, "email", email)
	}

	return user, err
}

func MatchExternalCreds() {

}

func MatchId(id uint64) (*User, error) {
	dbInstance := db.GetDB()

	user, err := getUserDetails(dbInstance, "id", id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func MatchExternalId(id string) (*User, error) {
	dbInstance := db.GetDB()

	user, err := getUserDetails(dbInstance, "external_id", id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUser(fields CreateUserStruct) (*User, error) {
	dbInstance := db.GetDB()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(fields.Password), 10)
	if err != nil {
		return nil, err
	}

	fields.Password = hashedPwd

	user, err := createUser(dbInstance, fields)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetFavorites(userId uint64) (*[]Activity, error) {
	dbInstance := db.GetDB()

	favorites, err := getUserFavorites(dbInstance, userId)
	if err != nil {
		return nil, err
	}

	favoritesStr := strings.Join(favorites, ",")

	activities, err := getActivitiesByIds(dbInstance, favoritesStr)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func UpdateFavorites(userId uint64, activityId uint64) (*string, error) {
	dbInstance := db.GetDB()

	activityIdStr := strconv.FormatUint(activityId, 10)

	favorites, err := getUserFavorites(dbInstance, userId)
	if err != nil {
		return nil, err
	}
	if favorites == nil {
		return nil, err
	}

	favoritesDRF := favorites

	idx := isFavorite(favoritesDRF, activityIdStr)

	if idx != -1 {
		favoritesDRF = append(favoritesDRF[:idx], favoritesDRF[idx+1:]...)
	} else {
		if len(favoritesDRF) == 0 || favoritesDRF[0] != "" {
			favoritesDRF = append(favoritesDRF, activityIdStr)
		} else {
			favoritesDRF[0] = activityIdStr
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

	updated := strings.Join(favoritesDRF, ",")

	_, err = dbInstance.Exec(queries.UpdateFavorites, updated, userId)

	return &updated, err
}

func UpdateSettings(userId uint64, settings map[string]interface{}) error {
	dbInstance := db.GetDB()

	err := updateUserSettings(dbInstance, userId, &settings)

	return err
}

func passwordMatch(dbInstance *sql.DB, queryBy string, value any, unhashedPassword string) bool {
	var hashedPwd string

	selectClause := fmt.Sprintf(`
		SELECT password 
		FROM users
		WHERE %s = ?
	`, queryBy)

	err := dbInstance.QueryRow(selectClause, value).Scan(&hashedPwd)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(unhashedPassword))
	if err != nil {
		return false
	}

	return true
}

func getUserDetails(dbInstance *sql.DB, queryBy string, value any) (*User, error) {
	var user User

	selectClause := fmt.Sprintf(`
		SELECT 
			id, account_type, agreed_to_terms, allow_location_tracking, avatar, bio, city, email, 
			email_verified, external_id, full_name, favorites, is_inactive, inactive_date, join_date, 
			locale_region, match_organized_count, match_played_count, password, 
			permissions, phone, preferred_locale, preferred_region, preferred_sport, preferred_theme, 
			reliability, role, sexe, show_age, show_email, show_groups, show_phone, timezone, username
		FROM users
		WHERE %s = ?
	`, queryBy)

	err := dbInstance.QueryRow(selectClause, value).Scan(
		&user.ID, &user.AccountType, &user.AgreedToTerms, &user.AllowLocationTracking, &user.Avatar, &user.Bio, &user.City, &user.Email,
		&user.EmailVerified, &user.ExternalID, &user.FullName, &user.Favorites, &user.IsInactive, &user.InactiveDate, &user.JoinDate,
		&user.LocaleRegion, &user.MatchOrganizedCount, &user.MatchPlayedCount, &user.Password,
		&user.Permissions, &user.Phone, &user.PreferredLocale, &user.PreferredRegion, &user.PreferredSport, &user.PreferredTheme,
		&user.Reliability, &user.Role, &user.Sexe, &user.ShowAge, &user.ShowEmail, &user.ShowGroups, &user.ShowPhone, &user.Timezone, &user.Username,
	)

	return &user, err
}

func getUserFavorites(dbInstance *sql.DB, userId uint64) ([]string, error) {
	var favorites *string

	err := dbInstance.QueryRow(queries.SelectFavorites, userId).Scan(&favorites)
	if err != nil {
		return nil, err
	}

	if favorites == nil {
		return []string{}, nil
	}

	splitFavorites := strings.Split(*favorites, ",")

	return splitFavorites, nil
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

func updateUserSettings(dbInstance *sql.DB, userId uint64, settings *map[string]interface{}) error {
	tx, err := dbInstance.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	var setColumns []string
	var setValues []interface{}

	for key, value := range *settings {
		columnName := util.CamelToSnake(key)

		setColumns = append(setColumns, columnName+" = ?")
		setValues = append(setValues, value)
	}

	setClause := strings.Join(setColumns, ", ")

	_, err = tx.Exec(fmt.Sprintf("UPDATE users SET %s WHERE id = ?", setClause), append(setValues, userId)...)

	return err
}

type CreateUserStruct struct {
	AccountType   types.AccountType
	AgreedToTerms bool
	Avatar        string
	Email         string
	EmailVerified bool
	ExternalID    string
	FullName      string
	Locale        string
	Password      []byte
	Phone         string
	Picture       string
	Username      string
}

func createUser(dbInstance *sql.DB, fields CreateUserStruct) (*User, error) {
	result, err := dbInstance.Exec(`
			INSERT INTO users (
				account_type, 
				agreed_to_terms,
				avatar,
				email, 
				email_verified, 
				external_id,
				full_name,
				password, 
				permissions, 
				phone, 
				role, 
				username
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		fields.AccountType,
		fields.AgreedToTerms,
		fields.Avatar,
		fields.Email,
		fields.EmailVerified,
		fields.ExternalID,
		fields.FullName,
		fields.Password,
		strings.Join(types.DEFAULT_PERMISSIONS[:], ","),
		fields.Phone,
		types.USER,
		fields.Username,
	)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	insertedUser, err := getUserDetails(dbInstance, "id", lastInsertID)

	log.Println("createUser - lastInsertID", lastInsertID)
	log.Println("createUser - insertedUser", insertedUser)

	return insertedUser, err
}
