package queries

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
const SelectAllFromMessageByChatroomId = `
    SELECT * 
    FROM messages
    WHERE chatroom_id = ?
`
const SelectAllFromMessageByChatroomIdAndSenderId = `
    SELECT * 
    FROM messages
    WHERE chatroom_id = ?
    AND sender_id = ?
`
const InsertIntoMessageForChatroomId = `
    INSERT INTO messages (content, delivered, chatroom_id, sender_id)
    VALUES (?, ?, ?, ?)
`
const DropMessageTable = `
    DROP TABLE IF EXISTS messages
`
