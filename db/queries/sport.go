package queries

const SelectAllFromSports = `
    SELECT
        sgm.sport_id,
        s.name AS name,
        s.feature_available AS feature_available,
        gm.name AS game_mode
    FROM
        sport_game_modes sgm
    JOIN
        sports s ON sgm.sport_id = s.id
    JOIN
        game_modes gm ON sgm.game_mode_id = gm.id
`

const InsertIntoSport = `
    INSERT INTO sports (name, feature_available)
    VALUES (?, ?)
`
const DropSportTable = `
    DROP TABLE IF EXISTS sports
`
