package queries

const CreateGroupUsersTable = `
    CREATE TABLE IF NOT EXISTS group_users (
        group_id BIGINT UNSIGNED,
        user_id BIGINT UNSIGNED,
        PRIMARY KEY (group_id, user_id),
        INDEX group_id_idx (group_id),
        INDEX user_id_idx (user_id)
    )
`
const SelectAllFromGroupUsersByGroupId = `
    SELECT * 
    FROM group_users
    WHERE group_id = ?
`
const SelectAllFromGroupUsersByUserId = `
    SELECT * 
    FROM group_users
    WHERE user_id = ?
`
const InsertIntoGroupUsers = `
    INSERT INTO group_users (group_id, user_id)
    VALUES (?, ?)
`
const DeleteFromGroupUsersByGroupId = `
    DELETE FROM group_users
    WHERE group_id = ?
`
const DeleteFromGroupUsersByUserId = `
    DELETE FROM group_users
    WHERE user_id = ?
`
const DropGroupUsersTable = `
    DROP TABLE IF EXISTS group_users
`
