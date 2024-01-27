package db_utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserSettings struct {
	ID                    string
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
	ID                  int
	AccountType         string
	Avatar              string
	Bio                 string
	City                string
	Email               string
	EmailVerified       bool
	FullName            string
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

func Seed(c *gin.Context) {
	var newUser User
	err := c.BindJSON(&newUser)
	if err != nil {
		log.Fatal("(CreateProduct) c.BindJSON", err)
	}

	query := `INSERT INTO products (name, price) VALUES (?, ?)`
	res, err := db.Exec(query, newUser.FullName, newUser.Email)
	if err != nil {
		log.Fatal("(CreateProduct) db.Exec", err)
	}
	newUser.Id, err = res.LastInsertId()
	if err != nil {
		log.Fatal("(CreateProduct) res.LastInsertId", err)
	}

	c.JSON(http.StatusOK, newUser)
}
