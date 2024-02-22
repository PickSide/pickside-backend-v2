package queries

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
