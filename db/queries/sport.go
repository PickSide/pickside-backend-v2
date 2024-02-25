package queries

const SelectAllFromSports = `
    SELECT *
    FROM sports
`

const InsertIntoSport = `
    INSERT INTO sports (name, game_modes, feature_available)
    VALUES (?, ?, ?)
`
const DropSportTable = `
    DROP TABLE IF EXISTS sports
`
