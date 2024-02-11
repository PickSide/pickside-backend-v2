package queries

const CreateSportGameModesTable = `
    CREATE TABLE IF NOT EXISTS sport_game_modes (
        game_mode_id BIGINT UNSIGNED,
        sport_id BIGINT UNSIGNED,
        PRIMARY KEY (sport_id, game_mode_id),
        INDEX sport_id_idx (sport_id),
        INDEX game_mode_id_idx (game_mode_id)
    );
`
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
