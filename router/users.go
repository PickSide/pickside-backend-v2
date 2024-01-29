package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"me/pickside/db/queries"
	"net/http"
	"time"
)

type AccountType string
type Permission string
type Role string
type Sexe string
type Theme string

const (
	GOOGLE   AccountType = "google"
	FACEBOOK AccountType = "facebook"
	APPLE    AccountType = "apple"
	DEFAULT  AccountType = "default"
	GUEST    AccountType = "guest"

	ACTIVITIES_VIEW       Permission = "activities-view"
	ACTIVITIES_CREATE     Permission = "activities-create"
	ACTIVITIES_DELETE     Permission = "activities-delete"
	ACTIVITIES_REGISTER   Permission = "activities-register"
	GROUP_CREATE          Permission = "group-create"
	GROUP_DELETE          Permission = "group-delete"
	GROUP_SEARCH          Permission = "group-search"
	USERS_VIEW_ALL        Permission = "see-all-users"
	USERS_VIEW_DETAIL     Permission = "see-detail-users"
	SEND_MESSAGES         Permission = "send-messages"
	NOTIFICATIONS_RECEIVE Permission = "notifications-receive"
	GOOGLE_SEARCH         Permission = "google-search"
	MAP_VIEW              Permission = "map-view"

	ADMIN Role = "admin"
	USER  Role = "user"

	LIGHT Theme = "light"
	DARK  Theme = "dark"
)

type UserSettings struct {
	ID                    uint16 `json:"id"`
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
	ID                  uint16      `json:"id"`
	AccountType         AccountType `json:"account_type"`
	Avatar              string      `json:"avatar" default:"default_avatar.jpg"`
	Bio                 string      `json:"bio" default:"My default bio"`
	City                string      `json:"city" default:"Unknown"`
	Email               string      `json:"email"`
	EmailVerified       bool        `json:"email_verified" default:"false"`
	FullName            string      `json:"full_name" default:"John Doe"`
	IsInactive          bool        `json:"is_inactive" default:"false"`
	InactiveDate        time.Time   `json:"inactive_date" default:"2000-01-01T00:00:00Z"`
	JoinDate            time.Time   `json:"join_date"`
	LocaleRegion        string      `json:"locale_region" default:"en_US"`
	MatchOrganizedCount int         `json:"match_organized_count" default:"0"`
	MatchPlayedCount    int         `json:"match_played_count" default:"0"`
	Password            string      `json:"password" default:"default_password"`
	Permissions         string      `json:"permissions" default:"read"`
	Phone               string      `json:"phone" default:""`
	Reliability         int         `json:"reliability" default:"50"`
	Role                Role        `json:"role" default:"user"`
	Sexe                string      `json:"sexe" default:"unknown"`
	Timezone            string      `json:"timezone" default:"0"`
	Username            string      `json:"username" default:"guest"`
}

func GetUsers(c *gin.Context, db *sql.DB) {
	rows, err := db.Query(queries.SELECT_ALL_FROM_USERS)
	if err != nil {
		return
	}
	defer rows.Close()

	var users []User

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	for rows.Next() {
		var user User
		err := rows.Scan(values...)
		if err != nil {
			log.Fatal("(GetUsers) rows.Scan", err)
		}
		users = append(users, user)
	}

	for _, user := range users {
		fmt.Printf("ID: %d\n", user.ID)
		fmt.Printf("Full Name: %s\n", user.FullName)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Printf("Password: %s\n", user.Password)
	}
	c.JSON(http.StatusOK, users)
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context, db *sql.DB) {
	var user_req UserRequest

	if err := c.ShouldBindJSON(&user_req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(user_req.Username)
	log.Println(user_req.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user_req.Password), 10)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(queries.INSERT_USER, "tonya", "tonyown11@gmail.com", user_req.Username, hashedPassword)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	c.JSON(http.StatusCreated, user_req)

}
