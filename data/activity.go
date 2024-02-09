package data

import (
	"pickside/service/db"
	"pickside/service/db/queries"
	"time"
)

type Activity struct {
	ID          uint64    `json:"id"`
	Address     string    `json:"address"`
	Date        string    `json:"date"`
	Description string    `json:"description"`
	IsPrivate   bool      `json:"isPrivate"`
	MaxPlayers  int       `json:"maxPlayers"`
	Price       float64   `json:"price"`
	Rules       string    `json:"rules"`
	Time        string    `json:"time"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	OrganizerID uint64    `json:"organizerId"`
	SportID     uint64    `json:"sportId"`
}

func AllActivities() (*[]Activity, error) {
	dbInstance := db.GetDB()

	rows, err := dbInstance.Query(queries.SelectAllFromActivities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []Activity

	for rows.Next() {
		var activity Activity

		err := rows.Scan(
			&activity.ID,
			&activity.Address,
			&activity.Date,
			&activity.Description,
			&activity.IsPrivate,
			&activity.MaxPlayers,
			&activity.Price,
			&activity.Rules,
			&activity.Time,
			&activity.Title,
			&activity.CreatedAt,
			&activity.UpdatedAt,
			&activity.OrganizerID,
			&activity.SportID,
		)
		if err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &activities, nil
}

func InsertActivity(address string, date string, maxPlayers int, description string, organizerId int64, price float32, rules string, time string, title string, isPrivate bool, sportId int64) (*Activity, error) {
	dbInstance := db.GetDB()

	result, err := dbInstance.Exec(
		queries.InsertActivity,
		address,
		date,
		description,
		isPrivate,
		maxPlayers,
		price,
		rules,
		organizerId,
		time,
		title,
		sportId,
	)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var insertedActivity Activity

	err = dbInstance.QueryRow(
		queries.SelectActivityById,
		lastInsertID,
	).Scan(
		&insertedActivity.ID,
		&insertedActivity.Address,
		&insertedActivity.Date,
		&insertedActivity.Description,
		&insertedActivity.IsPrivate,
		&insertedActivity.MaxPlayers,
		&insertedActivity.Price,
		&insertedActivity.Rules,
		&insertedActivity.Time,
		&insertedActivity.Title,
		&insertedActivity.CreatedAt,
		&insertedActivity.UpdatedAt,
		&insertedActivity.OrganizerID,
		&insertedActivity.SportID,
	)
	if err != nil {
		return nil, err
	}

	return &insertedActivity, err
}
