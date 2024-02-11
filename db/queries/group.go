package queries

const CreateGroupTable = `
    CREATE TABLE IF NOT EXISTS groups (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        cover_photo VARCHAR(255),
        description VARCHAR(255),
        name VARCHAR(255),
        requires_approval BOOL DEFAULT 0,
        visibility ENUM('public', 'private'),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        organizer_id BIGINT UNSIGNED,
        sport_id BIGINT UNSIGNED,
        INDEX organizer_id_idx (organizer_id),
        INDEX sport_id_idx (sport_id)
    )
`
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
