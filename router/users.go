package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"me/pickside/db/queries"
	"net/http"
)

type UserSettings struct {
	ID                    uint16
	PreferredSport        string
	PreferredLocale       string
	PreferredTheme        string
	PreferredRegion       string
	AllowLocationTracking bool
	ShowAge               bool
	ShowEmail             bool
	ShowPhone             bool
	ShowGroups            bool
	User                  User `gorm:"foreignKey:UserID"`
	UserID                User
}

type User struct {
	ID                  uint16 `json:"id"`
	AccountType         string
	Avatar              string
	Bio                 string
	City                string
	Email               string `json:"email"`
	EmailVerified       bool
	FullName            string `json:"full_name"`
	IsInactive          bool
	InactiveDate        int
	JoinDate            int64
	LocaleRegion        string
	MatchOrganizedCount int
	MatchPlayedCount    int
	Password            string
	Permissions         string
	Phone               string
	UserSettingsID      int
	Reliability         int
	Role                string
	Sexe                string
	Timezone            int64
	Username            string
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

	log.Println(users)
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

	rows, err := db.Query(queries.INSERT_USER, "tony", "tonyown10@gmail.com", user_req.Username, hashedPassword)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	c.JSON(http.StatusCreated, user_req)

}
