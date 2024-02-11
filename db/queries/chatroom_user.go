package queries

const CreateChatroomParticipantsTable = `
    CREATE TABLE IF NOT EXISTS chatroom_users (
        chatroom_id BIGINT UNSIGNED,
        user_id BIGINT UNSIGNED,
        PRIMARY KEY (chatroom_id, user_id),
        INDEX chatroom_id_idx (chatroom_id),
        INDEX user_id_idx (user_id)
    )
`
const SelectAllFromChatroomParticipantsByChatrooId = `
    SELECT * 
    FROM chatroom_users
    WHERE group_id = ?
`
const SelectAllFromChatroomParticipantsByUserId = `
    SELECT * 
    FROM chatroom_users
    WHERE chatroom_id = ?
`
const InsertIntoChatroomParticipants = `
    INSERT INTO chatroom_users (chatroom_id, user_id)
    VALUES (?, ?)
`
const DeleteFromChatroomParticipantsByGroupId = `
    DELETE FROM chatroom_users
    WHERE group_id = ?
`
const DeleteFromChatroomParticipantsByUserId = `
    DELETE FROM chatroom_users
    WHERE user_id = ?
`
const DropChatroomParticipantsTable = `
    DROP TABLE IF EXISTS chatroom_users
`
