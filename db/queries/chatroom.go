package queries

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
const SelectAllFromChatroomById = `
    SELECT * 
    FROM messages
    WHERE chatroom_id = ?
`
const DropChatroomTable = `
    DROP TABLE IF EXISTS chatrooms
`
