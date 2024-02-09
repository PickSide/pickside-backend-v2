package queries

const CreateChatroomParticipantsTable = `
    CREATE TABLE IF NOT EXISTS chatroom_participants (
        chatroom_id BIGINT UNSIGNED,
        user_id BIGINT UNSIGNED,
        PRIMARY KEY (chatroom_id, user_id),
        INDEX chatroom_id_idx (chatroom_id),
        INDEX user_id_idx (user_id)
    )
`
const SelectAllFromChatroomParticipantsByChatroomId = `
    SELECT * 
    FROM chatroom_participants
    WHERE group_id = ?
`
const SelectAllFromChatroomParticipantsByUserId = `
    SELECT * 
    FROM chatroom_participants
    WHERE chatroom_id = ?
`
const InsertIntoChatroomParticipants = `
    INSERT INTO chatroom_participants (chatroom_id, user_id)
    VALUES (?, ?)
`
const DeleteFromChatroomParticipantsByGroupId = `
    DELETE FROM chatroom_participants
    WHERE group_id = ?
`
const DeleteFromChatroomParticipantsByUserId = `
    DELETE FROM chatroom_participants
    WHERE user_id = ?
`
const DropChatroomParticipantsTable = `
    DROP TABLE IF EXISTS chatroom_participants
`
