package data

import (
	"me/pickside/db"
	"me/pickside/db/queries"
)

type Token struct {
	ID            uint64 `json:"id"`
	Value         string
	IsBlackListed bool   `json:"is_black_listed"`
	AssociatedId  uint64 `json:"associated_id"`
}

func InsertNewToken(value string, userID uint64) error {
	dbInstance := db.GetDB()

	_, err := dbInstance.Exec(queries.InsertNewToken, value, userID)
	if err != nil {
		return err
	}

	return nil
}

func BlackListToken(value string) error {
	dbInstance := db.GetDB()

	_, err := dbInstance.Exec(queries.BlackListToken, value)
	if err != nil {
		return err
	}

	return nil
}
