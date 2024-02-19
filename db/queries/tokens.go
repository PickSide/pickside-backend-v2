package queries

const DropTokensTable = `
	DROP TABLE IF EXISTS tokens;
`
const InsertNewToken = `
	INSERT 
	INTO tokens (value, associated_user_id) 
	VALUES (?, ?)
`

const BlackListToken = `
	UPDATE tokens 
	SET is_black_listed = 1
	WHERE value = ?
`
