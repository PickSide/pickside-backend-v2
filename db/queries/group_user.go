package queries

const SelectAllFromGroupUsersByGroupId = `
    SELECT * 
    FROM users_groups
    WHERE group_id = ?
`
const SelectAllFromGroupUsersByUserId = `
    SELECT * 
    FROM users_groups
    WHERE user_id = ?
`
const InsertIntoGroupUsers = `
    INSERT INTO users_groups (group_id, user_id)
    VALUES (?, ?)
`
const DeleteFromGroupUsersByGroupId = `
    DELETE FROM users_groups
    WHERE group_id = ?
`
const DeleteFromGroupUsersByUserId = `
    DELETE FROM users_groups
    WHERE user_id = ?
`
