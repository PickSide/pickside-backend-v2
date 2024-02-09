package queries

const CreateActivityTable = `
    CREATE TABLE IF NOT EXISTS activities (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        address VARCHAR(255),
        date DATE,
        description VARCHAR(255),
        is_private BOOL DEFAULT 0,
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

const SelectAllFromActivities = `
    SELECT * 
    FROM activities
`
const SelectActivityById = `
    SELECT *
    FROM activities
    WHERE id = ?
`

/**
 * @argument INSERT, address, date, description, is_private, max_players, price, rules, organizer_id, time, title, sport_id
 */

const InsertActivity = `
    INSERT 
    INTO activities (
        address, date, description, is_private, max_players, price, rules, organizer_id, time, title, sport_id
    )
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

const DropActivityTable = `
    DROP TABLE IF EXISTS activities
`
