package queries

const SelectGroupById = `
    SELECT 
        id,
        description,
        name,
        organizer_id,
        requires_approval,
        sport_id,
        visibility
    FROM groups
    WHERE id = ?
`
const SelectAllGroupsByOrganizer = `
    SELECT  
        g.id,
        g.description,
        g.name,
        g.organizer_id,
        g.requires_approval,
        g.sport_id,
        g.visibility
    FROM groups g
    JOIN users_groups gu ON g.id = gu.group_id
    WHERE g.organizer_id = ?
`
const SelectAllFromGroupsById = `
    SELECT * 
    FROM groups
    WHERE user_id = ?
`
const SelectAllGroupsBySportId = `
    SELECT * 
    FROM groups
    WHERE sport_id = ?
`
const InsertIntoGroup = `
    INSERT INTO groups (
        description, 
        name, 
        organizer_id, 
        requires_approval, 
        sport_id,
        visibility 
    )
    VALUES (?, ?, ?, ?, ?, ?)
`
const DeleteFromGroupById = `
    DELETE FROM groups
    WHERE id = ?
`
const UpdateGroupParticipants = `
    UPDATE groups
    SET participants = ?
    WHERE id = ?
`
const DropGroupTable = `
    DROP TABLE IF EXISTS groups
`
