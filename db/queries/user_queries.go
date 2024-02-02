package queries

const CreateUserTables = `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		account_type VARCHAR(255),
		avatar VARCHAR(255),
		bio VARCHAR(255),
		city VARCHAR(255),
		email VARCHAR(255),
		email_verified BOOLEAN,
		full_name VARCHAR(255),
		is_inactive BOOLEAN,
		inactive_date DATETIME NULL,
		join_date DATETIME,
		locale_region VARCHAR(255),
		match_organized_count INT,
		match_played_count INT,
		password VARCHAR(255),
		permissions VARCHAR(255),
		phone VARCHAR(255),
		reliability FLOAT,
		role VARCHAR(255),
		sexe VARCHAR(255),
		timezone VARCHAR(255),
		username VARCHAR(255)
	);

`

const CreateUserSettingsTable = `
	CREATE TABLE IF NOT EXISTS user_settings (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		preferred_sport VARCHAR(255),
		preferred_locale VARCHAR(255),
		preferred_theme VARCHAR(255),
		preferred_region VARCHAR(255),
		allow_location_tracking BOOLEAN,
		show_age BOOLEAN,
		show_email BOOLEAN,
		show_phone BOOLEAN,
		show_groups BOOLEAN,
		user_id BIGINT,
		INDEX user_id_idx (user_id)
	);
`

const DropUserTables = `
	DROP TABLE IF EXISTS users;
`

const DropUserSettingsTable = `
	DROP TABLE IF EXISTS user_settings;
`

const InsertUser = `
	INSERT INTO users (
		account_type, avatar, bio, city, email, email_verified, full_name, is_inactive,
		inactive_date, join_date, locale_region, match_organized_count, match_played_count,
		password, permissions, phone, reliability, role, sexe, timezone, username
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
const SelectPasswordOnlyWhereUsernameEquals = "SELECT password FROM users WHERE username = ?"
const SelectAllColumns = "SELECT * FROM users"
const SelectAllColumnsExceptPasswordWhereUsernameEquals = `
	SELECT 
    	id,
		account_type,
		avatar,
		bio,
		city,
		email,
		email_verified,
		full_name,
 		is_inactive,
		inactive_date,
		join_date,
		locale_region,
		match_organized_count,
		match_played_count,
		permissions,
		phone,
		reliability,
		role,
		sexe,
		timezone,
		username  
	FROM users
	WHERE username = ?`
const SelectAllColumnsExceptPassword = `
	SELECT 
    	id,
		account_type,
		avatar,
		bio,
		city,
		email,
		email_verified,
		full_name,
 		is_inactive,
		inactive_date,
		join_date,
		locale_region,
		match_organized_count,
		match_played_count,
		permissions,
		phone,
		reliability,
		role,
		sexe,
		timezone,
		username  
	FROM users`
