package queries

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
const SelectAllParticipants = `
    SELECT activity_id, user_id 
    FROM activity_users
    WHERE activity_id = ?
`
const IsUserRegistered = `
    SELECT * 
    FROM activity_users
    WHERE activity_id = ? 
    AND user_id = ?
`
const RegisterUserToActivity = `
    INSERT INTO activity_users (activity_id, user_id)
    VALUES (?, ?)
`
const RemoveUserFromActivity = `
    DELETE FROM activity_users
    WHERE activity_id = ? 
    AND user_id = ?
`
const DropActivityUserTable = `
    DROP TABLE IF EXISTS activity_users
`
