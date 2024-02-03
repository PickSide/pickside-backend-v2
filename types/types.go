package types

import "github.com/golang-jwt/jwt"

type AccountType string
type Role string
type Sexe string
type Theme string

const (
	GOOGLE   AccountType = "google"
	FACEBOOK AccountType = "facebook"
	APPLE    AccountType = "apple"
	DEFAULT  AccountType = "default"
	GUEST    AccountType = "guest"

	ACTIVITIES_VIEW        string = "activities-view"
	ACTIVITIES_CREATE      string = "activities-create"
	ACTIVITIES_DELETE      string = "activities-delete"
	ACTIVITIES_REGISTER    string = "activities-register"
	GROUP_CREATE           string = "group-create"
	GROUP_DELETE           string = "group-delete"
	GROUP_SEARCH           string = "group-search"
	USERS_VIEW_ALL         string = "see-all-users"
	USERS_VIEW_DETAIL      string = "see-detail-users"
	SEND_MESSAGES          string = "send-messages"
	NOTIFICATIONS_RECEIVE  string = "notifications-receive"
	GOOGLE_LOCATION_SEARCH string = "google-location-search"
	MAP_VIEW               string = "map-view"

	ADMIN Role = "admin"
	USER  Role = "user"

	LIGHT Theme = "light"
	DARK  Theme = "dark"

	MALE   Sexe = "male"
	FEMALE Sexe = "female"
)

type UserMetadata struct {
	AccountType `json:"accountType"`
}

type JWTClaims struct {
	jwt.Claims
	//UserMetadata  UserMetadata `json:"user_metadata"`
	ID            float64 `json:"sub"`
	Email         string  `json:"email"`
	Username      string  `json:"username"`
	EmailVerified bool    `json:"email_verified"`
}
