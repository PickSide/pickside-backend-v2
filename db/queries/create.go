package queries

const CreateActivityTable = `
    CREATE TABLE IF NOT EXISTS activities (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        address VARCHAR(255),
        date DATE,
        description VARCHAR(255),
        is_private BOOL,
        max_players INT,
        price FLOAT,
        rules VARCHAR(255),
        time VARCHAR(255),
        title VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        organizer_id BIGINT UNSIGNED,
        sport_id BIGINT UNSIGNED,
        INDEX organizer_id_idx (organizer_id),
        INDEX sport_id_idx (sport_id)
    );
`
const CreateChatroomTable = `
    CREATE TABLE IF NOT EXISTS chatrooms (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255),
        number_of_messages INT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        last_message_id BIGINT UNSIGNED,
        INDEX last_message_id_idx (last_message_id)
    )
`
const CreateGameModesTable = `
    CREATE TABLE IF NOT EXISTS game_modes (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL
    );
`
const CreateGroupTable = `
    CREATE TABLE IF NOT EXISTS groups (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        cover_photo VARCHAR(255),
        description VARCHAR(255),
        name VARCHAR(255),
        requires_approval BOOL,
        visibility ENUM('public', 'private'),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        organizer_id BIGINT UNSIGNED,
        sport_id BIGINT UNSIGNED,
        INDEX organizer_id_idx (organizer_id),
        INDEX sport_id_idx (sport_id)
    )
`
const CreateLocaleTable = `
    CREATE TABLE IF NOT EXISTS locales (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255),
        value VARCHAR(255),
        flag_code VARCHAR(255)
    )
`
const CreateMessageTable = `
    CREATE TABLE IF NOT EXISTS messages (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        content TEXT,
        delivered BOOL,
        sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        chatroom_id BIGINT UNSIGNED,
        sender_id BIGINT UNSIGNED,
        INDEX chatroom_id_idx (chatroom_id),
        INDEX sender_id_idx (sender_id)
    )
`
const CreateNotificationTable = `
    CREATE TABLE IF NOT EXISTS notifications (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        expires DATETIME,
		is_read BOOL,
		message VARCHAR(255),
		receiver_id BIGINT UNSIGNED,
		sender_id BIGINT UNSIGNED,
		type ENUM('system', 'global', 'like', 'group-invite', 'message-reminder', 'friend-invite'),
        INDEX receiver_id_idx (receiver_id),
        INDEX sender_id_idx (sender_id)
    );
`
const CreateSportTable = `
    CREATE TABLE IF NOT EXISTS sports (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255),
        feature_available BOOL
    )
`
const CreateTokensTable = `
	CREATE TABLE IF NOT EXISTS tokens (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		value VARCHAR(255),
		is_black_listed BOOL,
		associated_user_id BIGINT,
		INDEX associated_user_id_idx (associated_user_id)
	);
`
const CreateUserTables = `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		account_type ENUM('google', 'facebook', 'apple', 'default', 'guest'),
		avatar VARCHAR(255) NULL,
		bio VARCHAR(255) NULL,
		city VARCHAR(255) NULL,
		email VARCHAR(255),
		email_verified BOOL,
		full_name VARCHAR(255) NULL,
		favorites VARCHAR(255),
		is_inactive BOOL,
		inactive_date TIMESTAMP,
		join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		locale_region VARCHAR(255) NULL,
		match_organized_count INT,
		match_played_count INT,
		password VARCHAR(255),
		permissions VARCHAR(255),
		phone VARCHAR(255) NULL,
		reliability INT,
		role ENUM('admin', 'user'),
		sexe ENUM('male', 'female') DEFAULT 'male',
		timezone VARCHAR(255) NULL,
		username VARCHAR(255),
		agreed_to_terms BOOL,
        allow_location_tracking BOOL,
		preferred_locale VARCHAR(255),
		preferred_region VARCHAR(255),
		preferred_sport VARCHAR(255),
		preferred_theme VARCHAR(255),
		show_age BOOL,
		show_email BOOL,
		show_groups BOOL,
		show_phone BOOL
	);
`

/** LINK TABLE **/

const CreateActivityUserTable = `
    CREATE TABLE IF NOT EXISTS activity_users (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		activity_id BIGINT UNSIGNED,
		user_id BIGINT UNSIGNED,
		join_date DATETIME DEFAULT CURRENT_TIMESTAMP,
		INDEX activity_id_idx (activity_id),
		INDEX user_id_idx (user_id)
    )
`
const CreateChatroomParticipantsTable = `
    CREATE TABLE IF NOT EXISTS chatroom_users (
        chatroom_id BIGINT UNSIGNED,
        user_id BIGINT UNSIGNED,
        PRIMARY KEY (chatroom_id, user_id),
        INDEX chatroom_id_idx (chatroom_id),
        INDEX user_id_idx (user_id)
    )
`
const CreateGroupUsersTable = `
    CREATE TABLE IF NOT EXISTS group_users (
        group_id BIGINT UNSIGNED,
        user_id BIGINT UNSIGNED,
        PRIMARY KEY (group_id, user_id),
        INDEX group_id_idx (group_id),
        INDEX user_id_idx (user_id)
    )
`
const CreateSportGameModesTable = `
    CREATE TABLE IF NOT EXISTS sport_game_modes (
        game_mode_id BIGINT UNSIGNED,
        sport_id BIGINT UNSIGNED,
        PRIMARY KEY (sport_id, game_mode_id),
        INDEX sport_id_idx (sport_id),
        INDEX game_mode_id_idx (game_mode_id)
    );
`
