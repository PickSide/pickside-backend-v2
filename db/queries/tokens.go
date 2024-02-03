package queries

const CreateTokensTable = `
	CREATE TABLE IF NOT EXISTS tokens (
		id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		value VARCHAR(255),
		is_black_listed BOOL DEFAULT 0,
		associated_user_id BIGINT,
		INDEX associated_user_id_idx (associated_user_id)
	);
`
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
