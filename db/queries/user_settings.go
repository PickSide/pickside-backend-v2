package queries

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
const DropUserSettingsTable = `
	DROP TABLE IF EXISTS user_settings;
`
