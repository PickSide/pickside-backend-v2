package queries

const CreateUserTables = `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		account_type ENUM('google', 'facebook', 'apple', 'default', 'guest'),
		avatar VARCHAR(255),
		bio VARCHAR(255),
		city VARCHAR(255),
		email VARCHAR(255),
		email_verified BOOL DEFAULT 0,
		full_name VARCHAR(255),
		favorites VARCHAR(255),
		is_inactive BOOL DEFAULT 0,
		inactive_date TIMESTAMP,
		join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		locale_region VARCHAR(255),
		match_organized_count INT,
		match_played_count INT,
		password VARCHAR(255),
		permissions VARCHAR(255),
		phone VARCHAR(255),
		reliability FLOAT,
		role ENUM('admin', 'user'),
		sexe VARCHAR(255),
		timezone VARCHAR(255),
		username VARCHAR(255),
		agreed_to_terms BOOL DEFAULT 1
	);

`
const DropUserTables = `
	DROP TABLE IF EXISTS users;
`
const InsertUser = `
	INSERT INTO users (
		account_type, avatar, bio, city, email, email_verified, full_name, favorites, is_inactive,
		inactive_date, join_date, locale_region, match_organized_count, match_played_count,
		password, permissions, phone, reliability, role, sexe, timezone, username, agreed_to_terms
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`
const InsertUserSetting = `
	INSERT INTO user_settings (
		preferred_sport,
		preferred_locale,
		preferred_theme,
		preferred_region,
		allow_location_tracking,
		show_age,
		show_email,
		show_phone,
		show_groups,
		user_id
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`
const SelectUserByEmail = `
	SELECT 
		id, account_type, avatar, bio, city, email,
		email_verified,	full_name, favorites, is_inactive, inactive_date,
		join_date, locale_region, match_organized_count, 
		match_played_count, permissions, phone, reliability,
		role, sexe, timezone, username, agreed_to_terms
	FROM users 
	WHERE email = ?
`
const SelectUserById = `
	SELECT 
		id, account_type, avatar, bio, city, email,
		email_verified,	full_name, favorites, is_inactive, inactive_date,
		join_date, locale_region, match_organized_count, 
		match_played_count, permissions, phone, reliability,
		role, sexe, timezone, username, agreed_to_terms
	FROM users 
	WHERE id = ?
`
const SelectPasswordOnly = `
	SELECT password 
	FROM users 
	WHERE username = ?
`
const SelectByUsername = `
	SELECT 
    	id, account_type, avatar, bio, city, email,
		email_verified,	full_name, favorites, is_inactive, inactive_date,
		join_date, locale_region, match_organized_count, 
		match_played_count, permissions, phone, reliability,
		role, sexe, timezone, username, agreed_to_terms
	FROM users
	WHERE username = ?
`
const SelectFavorites = `
	SELECT favorites
	FROM users
	WHERE id = ?
`
const UpdateFavorites = `
	UPDATE users
	SET favorites = ?
	WHERE id = ?
`
