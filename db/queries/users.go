package queries

const InsertUserSeed = `
	INSERT INTO users (
		account_type, agreed_to_terms, allow_location_tracking, avatar, bio, city, email, email_verified, external_id, favorites,
		full_name, inactive_date, is_inactive, join_date, locale_region, match_organized_count, match_played_count,
		password, permissions, phone, preferred_locale, preferred_region, preferred_sport, preferred_theme,
		reliability, role, sexe, show_age, show_email, show_groups, show_phone, timezone, username
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
const SelectUserByEmail = `
	SELECT 
		id, account_type, agreed_to_terms, allow_location_tracking, avatar, bio, city, email, 
		email_verified, full_name, favorites, is_inactive, inactive_date, join_date, 
		locale_region, match_organized_count, match_played_count, password, 
		permissions, phone, preferred_locale, preferred_region, preferred_sport, preferred_theme, 
		reliability, role, sexe, show_age, show_email, show_groups, show_phone, timezone, username
	FROM users 
	WHERE email = ?
`
const SelectUserById = `
	SELECT 
		id, account_type, agreed_to_terms, allow_location_tracking, avatar, bio, city, email, 
		email_verified, full_name, favorites, is_inactive, inactive_date, join_date, 
		locale_region, match_organized_count, match_played_count, password, 
		permissions, phone, preferred_locale, preferred_region, preferred_sport, preferred_theme, 
		reliability, role, sexe, show_age, show_email, show_groups, show_phone, timezone, username
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
		id, account_type, agreed_to_terms, allow_location_tracking, avatar, bio, city, email, 
		email_verified, full_name, favorites, is_inactive, inactive_date, join_date, 
		locale_region, match_organized_count, match_played_count, password, 
		permissions, phone, preferred_locale, preferred_region, preferred_sport, preferred_theme, 
		reliability, role, sexe, show_age, show_email, show_groups, show_phone, timezone, username
	FROM users
	WHERE username = ?
`
const SelectFavorites = `
	SELECT favorites
	FROM users
	WHERE id = ?
`
