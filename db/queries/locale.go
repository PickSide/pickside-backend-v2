package queries

const SelectAllFromLocales = `
    SELECT * 
    FROM locales
`

const InsertIntoLocale = `
    INSERT INTO locales (name, value, flag_code)
    VALUES (?, ?, ?)
`
const DropLocaleTable = `
    DROP TABLE IF EXISTS locales
`
