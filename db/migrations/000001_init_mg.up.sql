CREATE TABLE IF NOT EXISTS activities (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(255),
    date DATE,
    description VARCHAR(255),
    game_mode VARCHAR(255),
    images VARCHAR(255) NULL,
    is_private BOOLEAN,
    lat DECIMAL(10, 8) NULL,
    lng DECIMAL(11, 8) NULL,
    max_players INT,
    organizer_id BIGINT UNSIGNED,
    price FLOAT,
    rules VARCHAR(255),
    sport_id BIGINT UNSIGNED,
    time VARCHAR(255),
    title VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX organizer_id_idx (organizer_id),
    INDEX sport_id_idx (sport_id)
);

CREATE TABLE IF NOT EXISTS chatrooms (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    number_of_messages INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    last_message_id BIGINT UNSIGNED,
    INDEX last_message_id_idx (last_message_id)
);

CREATE TABLE IF NOT EXISTS groups (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255),
    name VARCHAR(255),
    organizer_id BIGINT UNSIGNED,
    requires_approval BOOLEAN,
    sport_id BIGINT UNSIGNED,
    visibility ENUM('public', 'private'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX organizer_id_idx (organizer_id),
    INDEX sport_id_idx (sport_id)
);

CREATE TABLE IF NOT EXISTS locales (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    value VARCHAR(255),
    flag_code VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS messages (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    content TEXT,
    delivered BOOLEAN,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    chatroom_id BIGINT UNSIGNED,
    sender_id BIGINT UNSIGNED,
    INDEX chatroom_id_idx (chatroom_id),
    INDEX sender_id_idx (sender_id)
);

CREATE TABLE IF NOT EXISTS notifications (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    expires DATETIME,
    is_read BOOLEAN DEFAULT 0,
    content VARCHAR(255),
    recipient_id BIGINT UNSIGNED,
    title VARCHAR(255),
    INDEX recipient_id_idx (recipient_id)
);

CREATE TABLE IF NOT EXISTS global_notifications (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    expires DATETIME,
    content VARCHAR(255),
    title VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS sports (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    game_modes VARCHAR(255),
    feature_available BOOLEAN
);

CREATE TABLE IF NOT EXISTS tokens (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    value VARCHAR(255),
    is_black_listed BOOLEAN,
    associated_user_id BIGINT,
    INDEX associated_user_id_idx (associated_user_id)
);

CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    account_type ENUM(
        'google',
        'facebook',
        'apple',
        'default',
        'guest'
    ),
    avatar VARCHAR(255) DEFAULT '',
    bio VARCHAR(255) DEFAULT '',
    city VARCHAR(255) DEFAULT '',
    email VARCHAR(255) NOT NULL,
    email_verified BOOLEAN,
    external_id VARCHAR(255) NULL,
    full_name VARCHAR(255) DEFAULT '',
    favorites VARCHAR(255) DEFAULT '',
    is_inactive BOOLEAN DEFAULT 0,
    inactive_date TIMESTAMP,
    join_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    locale_region VARCHAR(255),
    match_organized_count INT DEFAULT 0,
    match_played_count INT DEFAULT 0,
    password VARCHAR(255) DEFAULT '',
    permissions VARCHAR(255) DEFAULT '',
    phone VARCHAR(255) DEFAULT '',
    reliability INT DEFAULT 0,
    role ENUM('admin', 'user') DEFAULT 'user',
    sexe ENUM('male', 'female') DEFAULT 'male',
    timezone VARCHAR(255) DEFAULT '',
    username VARCHAR(255) NOT NULL,
    agreed_to_terms BOOLEAN DEFAULT 0,
    allow_location_tracking BOOLEAN DEFAULT 0,
    preferred_locale VARCHAR(255) DEFAULT '',
    preferred_region VARCHAR(255) DEFAULT '',
    preferred_sport VARCHAR(255) DEFAULT '',
    preferred_theme VARCHAR(255) DEFAULT '',
    show_age BOOLEAN DEFAULT 0,
    show_email BOOLEAN DEFAULT 0,
    show_groups BOOLEAN DEFAULT 0,
    show_phone BOOLEAN DEFAULT 0
);

CREATE TABLE IF NOT EXISTS activity_users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    activity_id BIGINT UNSIGNED,
    user_id BIGINT UNSIGNED,
    join_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX activity_id_idx (activity_id),
    INDEX user_id_idx (user_id)
);

CREATE TABLE IF NOT EXISTS chatroom_users (
    chatroom_id BIGINT UNSIGNED,
    user_id BIGINT UNSIGNED,
    PRIMARY KEY (chatroom_id, user_id),
    INDEX user_id_idx (user_id)
);

CREATE TABLE IF NOT EXISTS users_groups (
    group_id BIGINT UNSIGNED,
    user_id BIGINT UNSIGNED,
    PRIMARY KEY (group_id, user_id),
    INDEX user_id_idx (user_id)
);