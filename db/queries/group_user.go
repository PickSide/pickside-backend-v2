package queries

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
