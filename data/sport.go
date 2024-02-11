package data

import (
	"pickside/service/db"
	"pickside/service/db/queries"
)

type Sport struct {
	ID               uint64 `json:"id"`
	Name             string `json:"name"`
	FeatureAvailable bool   `json:"featureAvailable"`
	GameMode         string `json:"gameMode"`
}

func AllSports() (*[]Sport, error) {
	dbInstance := db.GetDB()

	rows, err := dbInstance.Query(queries.SelectAllFromSports)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sports []Sport

	for rows.Next() {
		var sport Sport

		err := rows.Scan(
			&sport.ID,
			&sport.Name,
			&sport.FeatureAvailable,
			&sport.GameMode,
		)
		if err != nil {
			return nil, err
		}

		sports = append(sports, sport)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &sports, nil
}
