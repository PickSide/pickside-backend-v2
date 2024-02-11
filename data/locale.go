package data

import (
	"pickside/service/db"
	"pickside/service/db/queries"
)

type Locale struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	FlagCode string `json:"flagCode"`
}

func AllLocales() (*[]Locale, error) {
	dbInstance := db.GetDB()

	rows, err := dbInstance.Query(queries.SelectAllFromLocales)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locales []Locale

	for rows.Next() {
		var locale Locale

		err := rows.Scan(
			&locale.ID,
			&locale.Name,
			&locale.FlagCode,
		)
		if err != nil {
			return nil, err
		}

		locales = append(locales, locale)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &locales, nil
}
