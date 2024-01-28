package queries

type AccountType string
type Permission string
type Role string

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
)

const CREATE_USER_TABLE = `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		account_type VARCHAR(255),
		avatar VARCHAR(255),
		bio VARCHAR(255),
		city VARCHAR(255),
		email VARCHAR(255),
		email_verified BOOLEAN,
		full_name VARCHAR(255),
		is_inactive BOOLEAN,
		inactive_date DATE,
		join_date DATE,
		locale_region VARCHAR(255),
		match_organized_count INT,
		match_played_count INT,
		password VARCHAR(255),
		permissions VARCHAR(255),
		phone VARCHAR(255),
		user_settings_id VARCHAR(255),
		reliability FLOAT,
		role VARCHAR(255),
		sexe VARCHAR(255),
		timezone VARCHAR(255),
		username VARCHAR(255)
	)
`
const DROP_USER_TABLE = "DROP TABLE IF EXISTS users"

const INSERT_USER = "INSERT INTO users (full_name, email, password) VALUES (?, ?, ?, ?)"

const SELECT_ALL_FROM_USERS = "SELECT * FROM users"
