package data

import (
	"database/sql"
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
	ID                  uint64            `json:"id"`
	AccountType         types.AccountType `json:"accountType,omitempty"`
	Avatar              *string           `json:"avatar,omitempty"`
	Bio                 *string           `json:"bio,omitempty"`
	City                *string           `json:"city,omitempty"`
	Email               string            `json:"email,omitempty"`
	EmailVerified       bool              `json:"emailVerified,omitempty"`
	FullName            string            `json:"fullName,omitempty"`
	Favorites           *string           `json:"favorites,omitempty"`
	IsInactive          bool              `json:"isInactive,omitempty"`
	InactiveDate        *time.Time        `json:"inactiveDate,omitempty"`
	JoinDate            time.Time         `json:"joinDate,omitempty"`
	LocaleRegion        *string           `json:"localeRegion,omitempty"`
	MatchOrganizedCount int               `json:"matchOrganizedCount,omitempty"`
	MatchPlayedCount    int               `json:"matchPlayedCount,omitempty"`
	Password            string            `json:"-"`
	Permissions         string            `json:"permissions,omitempty"`
	Phone               *string           `json:"phone,omitempty"`
	Reliability         *int              `json:"reliability,omitempty"`
	Role                types.Role        `json:"role,omitempty"`
	Sexe                types.Sexe        `json:"sexe,omitempty"`
	Timezone            *string           `json:"timezone,omitempty"`
	Username            string            `json:"username,omitempty"`
	AgreedToTerms       bool              `json:"agreedToTerms,omitempty"`
	Settings            *UserSettings     `json:"settings,omitempty"`
}

type UserSettings struct {
	ID                    uint64 `json:"-"`
	AllowLocationTracking bool   `json:"allowLocationTracking,omitempty"`
	PreferredLocale       string `json:"preferredLocale,omitempty"`
	PreferredRegion       string `json:"preferredRegion,omitempty"`
	PreferredSport        string `json:"preferredSport,omitempty"`
	PreferredTheme        string `json:"preferredTheme,omitempty"`
	ShowAge               bool   `json:"showAge,omitempty"`
	ShowEmail             bool   `json:"showEmail,omitempty"`
	ShowGroups            bool   `json:"showGroups,omitempty"`
	ShowPhone             bool   `json:"showPhone,omitempty"`
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

	user, err := getUserDetails(dbInstance, queries.SelectByUsername, username)

	settings, err := getUserSettings(dbInstance, user.ID)
	if err != nil && err == sql.ErrNoRows {
		settings, err = createUserSettings(dbInstance, user.ID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	user.Settings = settings

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

	user, err := getUserDetails(dbInstance, queries.SelectUserByEmail, email)
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

	settings, err := getUserSettings(dbInstance, user.ID)
	if err != nil && err == sql.ErrNoRows {
		settings, err = createUserSettings(dbInstance, user.ID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	user.Settings = settings

	return user, nil
}

func MatchById(userId uint64) (*User, error) {
	dbInstance := db.GetDB()

	user, err := getUserDetails(dbInstance, queries.SelectUserById, userId)
	if err != nil {
		return nil, err
	}

	settings, err := getUserSettings(dbInstance, user.ID)
	if err != nil && err == sql.ErrNoRows {
		settings, err = createUserSettings(dbInstance, user.ID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	user.Settings = settings

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

	settings, err := getUserSettings(dbInstance, user.ID)
	if err != nil && err == sql.ErrNoRows {
		settings, err = createUserSettings(dbInstance, user.ID)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	user.Settings = settings

	return user, nil
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

func UpdateFavorites(userId uint64, activityId uint64) (*string, error) {
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

	updated := strings.Join(favorites, ",")

	_, err = dbInstance.Exec(queries.UpdateFavorites, updated, userId)

	return &updated, err
}

func GetUserSettings(userId uint64) (*UserSettings, error) {
	dbInstance := db.GetDB()

	settings, err := getUserSettings(dbInstance, userId)

	return settings, err
}

func UpdateUserSettings(userId uint64, newSettings UserSettings) (*UserSettings, error) {
	dbInstance := db.GetDB()

	settings, err := getUserSettings(dbInstance, userId)
	if err != nil {
		return nil, err
	}

	appendSettings(settings, newSettings)

	err = updateUserSettings(dbInstance, userId, settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
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
	err := dbInstance.QueryRow(byQuery, value).Scan(
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

func getUserSettings(dbInstance *sql.DB, userId uint64) (*UserSettings, error) {
	var settings UserSettings

	err := dbInstance.QueryRow(queries.SelectUserSetting, userId).Scan(
		&settings.AllowLocationTracking,
		&settings.PreferredLocale,
		&settings.PreferredRegion,
		&settings.PreferredSport,
		&settings.PreferredTheme,
		&settings.ShowAge,
		&settings.ShowEmail,
		&settings.ShowGroups,
		&settings.ShowPhone,
	)

	return &settings, err
}

func updateUserSettings(dbInstance *sql.DB, userId uint64, settings *UserSettings) error {
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

	_, err = dbInstance.Exec(queries.UpdateSettings,
		settings.AllowLocationTracking,
		settings.PreferredLocale,
		settings.PreferredRegion,
		settings.PreferredSport,
		settings.PreferredTheme,
		settings.ShowAge,
		settings.ShowEmail,
		settings.ShowGroups,
		settings.ShowPhone,
		userId,
	)

	return err
}

func createUserSettings(dbInstance *sql.DB, userId uint64) (*UserSettings, error) {
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

	settings := UserSettings{
		AllowLocationTracking: false,
		PreferredLocale:       "",
		PreferredRegion:       "",
		PreferredSport:        "",
		PreferredTheme:        "",
		ShowAge:               true,
		ShowEmail:             true,
		ShowGroups:            false,
		ShowPhone:             false,
	}

	_, err = dbInstance.Exec(queries.InsertUserSetting,
		settings.AllowLocationTracking,
		settings.PreferredLocale,
		settings.PreferredRegion,
		settings.PreferredSport,
		settings.PreferredTheme,
		settings.ShowAge,
		settings.ShowEmail,
		settings.ShowGroups,
		settings.ShowPhone,
		userId,
	)

	return &settings, err
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
	log.Println(fields)
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

func appendSettings(currentSettings *UserSettings, updated UserSettings) {
	if updated.AllowLocationTracking {
		currentSettings.AllowLocationTracking = updated.AllowLocationTracking
	}
	if updated.PreferredLocale != "" {
		currentSettings.PreferredLocale = updated.PreferredLocale
	}
	if updated.PreferredRegion != "" {
		currentSettings.PreferredRegion = updated.PreferredRegion
	}
	if updated.PreferredSport != "" {
		currentSettings.PreferredSport = updated.PreferredSport
	}
	if updated.PreferredTheme != "" {
		currentSettings.PreferredTheme = updated.PreferredTheme
	}
	if updated.ShowAge {
		currentSettings.ShowAge = updated.ShowAge
	}
	if updated.ShowEmail {
		currentSettings.ShowEmail = updated.ShowEmail
	}
	if updated.ShowGroups {
		currentSettings.ShowGroups = updated.ShowGroups
	}
	if updated.ShowPhone {
		currentSettings.ShowPhone = updated.ShowPhone
	}
}
