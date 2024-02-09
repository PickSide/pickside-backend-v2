package queries

const CreateNotificationTable = `
    CREATE TABLE IF NOT EXISTS notifications (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        expires DATETIME,
		is_read BOOL DEFAULT 0,
		message VARCHAR(255),
		receiver_id BIGINT UNSIGNED,
		sender_id BIGINT UNSIGNED,
		type ENUM('system', 'global', 'like', 'group-invite', 'message-reminder', 'friend-invite'),
        INDEX receiver_id_idx (receiver_id),
        INDEX sender_id_idx (sender_id)
    );
`
const SelectAllFromNotification = `
    SELECT *
    FROM notifications
`

const InsertIntoNotification = `
    INSERT INTO notifications (expires, , flag_code)
    VALUES (?, ?, ?)
`
const DropNotificationTable = `
    DROP TABLE IF EXISTS notifications
`
