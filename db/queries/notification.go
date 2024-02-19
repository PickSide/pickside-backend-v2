package queries

const SelectAllFromNotification = `
    SELECT *
    FROM notifications
`

const InsertIntoNotification = `
    INSERT INTO notifications (expires, , flag_code)
    VALUES (?, ?, ?)
`
const DropNotificationTable = `
    DROP TABLE IF EXISTS notifications
`
