package data

import (
	"pickside/service/db"
	"pickside/service/db/queries"
)

type Token struct {
	ID            uint64 `json:"id"`
	Value         string
	IsBlackListed bool   `json:"isBlackListed"`
	AssociatedId  uint64 `json:"associatedId"`
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
