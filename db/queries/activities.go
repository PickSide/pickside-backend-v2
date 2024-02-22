package queries

const SelectAllActivities = `
    SELECT *
    FROM activities
`
const SelectActivityById = `
    SELECT *
    FROM activities
    WHERE id = ?
`
const InsertActivity = `
    INSERT 
    INTO activities (address, date, description, is_private, max_players, price, rules, organizer_id, time, title, sport_id)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`
const SelectAllParticipants = `
    SELECT u.id, u.avatar, u.bio, u.city, u.email, u.full_name, u.join_date, u.locale_region, u.match_organized_count, u.match_played_count, u.sexe, u.username
    FROM activity_users
    LEFT JOIN users AS u ON activity_users.user_id = u.id
    WHERE activity_users.activity_id = ?;
`
const IsRegistered = `
    SELECT * 
    FROM activity_users
    WHERE activity_id = ? 
    AND user_id = ?
`
const RegisterTo = `
    INSERT INTO activity_users (activity_id, user_id)
    VALUES (?, ?)
`
const RemoveFromActivity = `
    DELETE FROM activity_users
    WHERE activity_id = ? 
    AND user_id = ?
`
