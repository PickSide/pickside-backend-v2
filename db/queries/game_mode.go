package queries

const CreateGameModesTable = `
    CREATE TABLE IF NOT EXISTS game_modes (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL
    );
`
const SelectNameFromGameModes = `
    SELECT name 
    FROM game_modes
`

const InsertIntoGameMode = `
    INSERT INTO game_modes (name)
    VALUES (?)
`
const DropGameModesTable = `
    DROP TABLE IF EXISTS game_modes
`
