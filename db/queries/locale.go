package queries

const CreateLocaleTable = `
    CREATE TABLE IF NOT EXISTS locales (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255),
        flag_code VARCHAR(255)
    )
`
const SelectAllFromLocales = `
    SELECT * 
    FROM locales
`

const InsertIntoLocale = `
    INSERT INTO locales (name, flag_code)
    VALUES (?, ?)
`
const DropLocaleTable = `
    DROP TABLE IF EXISTS locales
`
