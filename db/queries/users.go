package queries

const InsertUserSeed = `
	INSERT INTO users (
		account_type, avatar, bio, city, email, email_verified, full_name, favorites, is_inactive,
		inactive_date, join_date, locale_region, match_organized_count, match_played_count,
		password, permissions, phone, reliability, role, sexe, timezone, username, agreed_to_terms
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`
const InsertUser = `
	INSERT INTO users (
		account_type, 
		agreed_to_terms,
		avatar,
		email, 
		email_verified, 
		full_name,
		password, 
		permissions, 
		phone, 
		role, 
		username
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`
const InsertUserSetting = `
	INSERT INTO user_settings (
		allow_location_tracking,
		preferred_locale,
		preferred_region,
		preferred_sport,
		preferred_theme,
		show_age,
		show_email,
		show_groups,
		show_phone,
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
const SelectUserSetting = `
	SELECT 
		allow_location_tracking,
		preferred_locale,
		preferred_region,
		preferred_sport,
		preferred_theme,
		show_age,
		show_email,
		show_groups,
		show_phone
	FROM user_settings
	WHERE user_id = ?
`
const UpdateSettings = `
	UPDATE user_settings
	SET
		allow_location_tracking = ?,
		preferred_locale = ?,
		preferred_region = ?,
		preferred_sport = ?,
		preferred_theme = ?,
		show_age = ?,
		show_email = ?,
		show_groups = ?,
		show_phone = ?
	WHERE user_id = ?
`
