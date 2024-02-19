package queries

const SelectGroupById = `
    SELECT * 
    FROM groups
    WHERE id = ?
`
const SelectAllGroupsByOrganizer = `
    SELECT  
        g.id,
        g.cover_photo,
        g.description,
        g.name,
        g.requires_approval,
        g.visibility,
        g.created_at,
        g.updated_at,
        g.organizer_id,
        g.sport_id
    FROM groups g
    JOIN group_users gu ON g.id = gu.group_id
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
        cover_photo, description, 
        name, requires_approval, visibility, 
        organizer_id, sport_id
    )
    VALUES (?, ?, ?, ?, ?, ?, ?)
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
