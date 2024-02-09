package queries

const CreateGroupMembersTable = `
    CREATE TABLE IF NOT EXISTS group_members (
        group_id BIGINT UNSIGNED,
        user_id BIGINT UNSIGNED,
        PRIMARY KEY (group_id, user_id),
        INDEX group_id_idx (group_id),
        INDEX user_id_idx (user_id)
    )
`
const SelectAllFromGroupMembersByGroupId = `
    SELECT * 
    FROM group_members
    WHERE group_id = ?
`
const SelectAllFromGroupMembersByUserId = `
    SELECT * 
    FROM group_members
    WHERE user_id = ?
`
const InsertIntoGroupMembers = `
    INSERT INTO group_members (group_id, user_id)
    VALUES (?, ?)
`
const DeleteFromGroupMembersByGroupId = `
    DELETE FROM group_members
    WHERE group_id = ?
`
const DeleteFromGroupMembersByUserId = `
    DELETE FROM group_members
    WHERE user_id = ?
`
const DropGroupMembersTable = `
    DROP TABLE IF EXISTS group_members
`
