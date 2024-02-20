package data

import (
	"database/sql"
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
	Avatar                string            `json:"avatar,omitempty"`
	Bio                   string            `json:"bio,omitempty"`
	City                  string            `json:"city,omitempty"`
	Email                 string            `json:"email,omitempty"`
	EmailVerified         bool              `json:"emailVerified"`
	Favorites             string            `json:"favorites,omitempty"`
	FullName              string            `json:"fullName,omitempty"`
	InactiveDate          *time.Time        `json:"inactiveDate,omitempty"`
	IsInactive            bool              `json:"isInactive"`
	JoinDate              time.Time         `json:"joinDate,omitempty"`
	LocaleRegion          string            `json:"localeRegion,omitempty"`
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

	user, err := getUserDetails(dbInstance, "username", username)

	return user, nil
}

func MatchUserByEmail(email string, locale string, fullName string, picture string, verifiedEmail bool, accountType types.AccountType) (*User, error) {
	dbInstance := db.GetDB()

	randomStr, err := util.GenerateRandomString(5)
	if err != nil {
		return nil, err
	}

	hashedRandomPwd, err := bcrypt.GenerateFromPassword([]byte(randomStr), 10)
	if err != nil {
		return nil, err
	}

	user, err := getUserDetails(dbInstance, "email", email)
	if err != nil && err == sql.ErrNoRows {
		user, err = createUser(dbInstance, CreateUserStruct{
			AccountType:   accountType,
			AgreedToTerms: true,
			Avatar:        picture,
			Email:         email,
			EmailVerified: verifiedEmail,
			FullName:      fullName,
			Password:      hashedRandomPwd,
		})
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func MatchById(userId uint64) (*User, error) {
	dbInstance := db.GetDB()

	user, err := getUserDetails(dbInstance, "id", userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUser(email string, locale string, fullName string, password string, verifiedEmail bool, accountType types.AccountType) (*User, error) {
	dbInstance := db.GetDB()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user, err := createUser(dbInstance, CreateUserStruct{
		AccountType:   types.DEFAULT,
		AgreedToTerms: true,
		Email:         email,
		EmailVerified: verifiedEmail,
		FullName:      fullName,
		Password:      hashedPwd,
	})
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

func getPasswordFromDB(dbInstance *sql.DB, username string) (string, error) {
	var hashedPwd string
	err := dbInstance.QueryRow(queries.SelectPasswordOnly, username).Scan(&hashedPwd)
	return hashedPwd, err
}

func comparePasswords(hashedPwd, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
}

func getUserDetails(dbInstance *sql.DB, byQuery string, value any) (*User, error) {
	var user User

	selectClause := fmt.Sprintf(`
		SELECT 
			id, account_type, agreed_to_terms, allow_location_tracking, avatar, bio, city, email, 
			email_verified, full_name, favorites, is_inactive, inactive_date, join_date, 
			locale_region, match_organized_count, match_played_count, password, 
			permissions, phone, preferred_locale, preferred_region, preferred_sport, preferred_theme, 
			reliability, role, sexe, show_age, show_email, show_groups, show_phone, timezone, username
		FROM users
		WHERE %s = ?
	`, byQuery)

	log.Println(selectClause)

	err := dbInstance.QueryRow(selectClause, value).Scan(
		&user.ID, &user.AccountType, &user.AgreedToTerms, &user.AllowLocationTracking, &user.Avatar, &user.Bio, &user.City, &user.Email,
		&user.EmailVerified, &user.FullName, &user.Favorites, &user.IsInactive, &user.InactiveDate, &user.JoinDate,
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
	FullName      string
	Password      []byte
	Phone         string
	Picture       string
	Username      string
}

func createUser(dbInstance *sql.DB, fields CreateUserStruct) (*User, error) {
	result, err := dbInstance.Exec(queries.InsertUser,
		fields.AccountType,
		fields.AgreedToTerms,
		fields.Avatar,
		fields.Email,
		fields.EmailVerified,
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

	return &insertedUser, err
}
