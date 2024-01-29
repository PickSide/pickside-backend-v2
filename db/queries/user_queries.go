package queries

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

const INSERT_USER = "INSERT INTO users (account_type, avatar, bio, city, email, email_verified, full_name, password, username) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

const SELECT_ALL_FROM_USERS = "SELECT * FROM users"
