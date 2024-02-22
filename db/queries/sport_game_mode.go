package queries

const SelectAllFromSportGameModes = `
    SELECT sport_id, game_mode_id
    FROM sport_game_modes
`

const InsertIntoSportGameMode = `
    INSERT INTO sport_game_modes (game_mode_id, sport_id)
    VALUES (?, ?)
`
const DropSportGameModesTable = `
    DROP TABLE IF EXISTS sport_game_modes
`
