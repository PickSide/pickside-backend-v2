// Command golang-example demonstrates how to connect to PlanetScale from a Go
// application.
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	UserID                int
	User                  User `gorm:"foreignKey:UserID"`
}

// A User contains metadata about a user for sale.
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

func main() {
	// Load environment variables from file.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Connect to PlanetScale database using DSN environment variable.
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	}

	// Create an API handler which serves data from PlanetScale.
	handler := NewHandler(db)

	// Start an HTTP API server.
	const addr = ":8080"
	log.Printf("successfully connected to PlanetScale, starting HTTP server on %q", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}

// A Handler is an HTTP API server handler.
type Handler struct {
	db *gorm.DB
}

// NewHandler creates an http.Handler which wraps a PlanetScale database
// connection.
func NewHandler(db *gorm.DB) http.Handler {
	h := &Handler{db: db}

	r := mux.NewRouter()
	r.HandleFunc("/seed", h.seedDatabase).Methods(http.MethodGet)
	r.HandleFunc("/users", h.getProducts).Methods(http.MethodGet)
	r.HandleFunc("/user-settings", h.getUserSettings).Methods(http.MethodGet)

	return r
}

// seedDatabase is the HTTP handler for GET /seed.
func (h *Handler) seedDatabase(w http.ResponseWriter, r *http.Request) {
	// Perform initial schema migrations.
	if err := h.db.AutoMigrate(&User{}); err != nil {
		http.Error(w, "failed to migrate users table", http.StatusInternalServerError)
		return
	}

	if err := h.db.AutoMigrate(&UserSettings{}); err != nil {
		http.Error(w, "failed to migrate categories table", http.StatusInternalServerError)
		return
	}

	h.db.Create(&User{
		FullName: "Tony Hakim",
		Bio:      "Description 1",
		Email:    "tonyown10@gmail.com",
	})
	h.db.Create(&User{
		FullName: "Niloo Khastavan",
		Bio:      "Description 2",
		Email:    "niloo@gmail.com",
	})

	// Seed categories and users for those categories.
	h.db.Create(&UserSettings{
		PreferredSport:        "Soccer",
		PreferredLocale:       "en",
		PreferredTheme:        "dark",
		PreferredRegion:       "montreal",
		AllowLocationTracking: false,
		ShowAge:               false,
		ShowEmail:             false,
		ShowPhone:             false,
		ShowGroups:            false,
		User:                  User{ID: 1},
	})
	h.db.Create(&UserSettings{
		PreferredSport:        "Soccer",
		PreferredLocale:       "en",
		PreferredTheme:        "dark",
		PreferredRegion:       "montreal",
		AllowLocationTracking: false,
		ShowAge:               false,
		ShowEmail:             false,
		ShowPhone:             false,
		ShowGroups:            false,
		User:                  User{ID: 2},
	})

	io.WriteString(w, "Migrations and Seeding of database complete\n")
}

// getProducts is the HTTP handler for GET /users.
func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	var users []User
	result := h.db.Preload("UserSettings").Find(&users)
	if result.Error != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(users)
}

// getProduct is the HTTP handler for GET /users/{id}.
func (h *Handler) getProduct(w http.ResponseWriter, r *http.Request) {
	var user User
	result := h.db.First(&user, mux.Vars(r)["id"])
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// getCategories is the HTTP handler for GET /categories.
func (h *Handler) getCategories(w http.ResponseWriter, r *http.Request) {
	var categories []UserSettings
	result := h.db.Find(&categories)
	if result.Error != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(categories)
}

// getUserSettings is the HTTP handler for GET /UserSettings/{id}.
func (h *Handler) getUserSettings(w http.ResponseWriter, r *http.Request) {
	var UserSettings UserSettings
	result := h.db.First(&UserSettings, mux.Vars(r)["id"])
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(UserSettings)
}
