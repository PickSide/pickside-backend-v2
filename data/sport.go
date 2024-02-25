package data

import (
	"database/sql"
	"pickside/service/db"
)

type Sport struct {
	ID               uint64 `json:"id"`
	Name             string `json:"name"`
	FeatureAvailable bool   `json:"featureAvailable"`
	GameModes        string `json:"gameModes"`
}

func AllSports() (*[]Sport, error) {
	dbInstance := db.GetDB()

	sports, err := getAllSports(dbInstance)

	return sports, err
}

func getAllSports(dbInstance *sql.DB) (*[]Sport, error) {
	rows, err := dbInstance.Query("SELECT * FROM sports")
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
			&sport.GameModes,
			&sport.FeatureAvailable,
		)
		if err != nil {
			return nil, err
		}

		sports = append(sports, sport)
	}

	err = rows.Err()

	return &sports, err
}

func getSportById(dbInstance *sql.DB, sportId uint64) (*Sport, error) {
	var sport Sport

	err := dbInstance.QueryRow(`SELECT * FROM sports WHERE id = ?`, sportId).Scan(
		&sport.ID,
		&sport.Name,
		&sport.GameModes,
		&sport.FeatureAvailable,
	)
	if err != nil {
		return nil, err
	}

	return &sport, err
}
