package queries

const UpdateFavorites = `
	UPDATE users
	SET favorites = ?
	WHERE id = ?
`
