package queries

const CreateSportTable = `
    CREATE TABLE IF NOT EXISTS sports (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255),
        featureAvailable BOOL
    )
`
const SelectAllFromSports = `
    SELECT * 
    FROM sports
`

const InsertIntoSport = `
    INSERT INTO sports (name, featureAvailable)
    VALUES (?, ?)
`
const DropSportTable = `
    DROP TABLE IF EXISTS sports
`
