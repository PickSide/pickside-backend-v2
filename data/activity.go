package data

import (
	"database/sql"
	"pickside/service/db"
	"pickside/service/db/queries"
	"time"
)

type Activity struct {
	ID           uint64        `json:"id"`
	Address      string        `json:"address"`
	Date         string        `json:"date"`
	Description  string        `json:"description"`
	IsPrivate    bool          `json:"isPrivate"`
	MaxPlayers   int           `json:"maxPlayers"`
	Price        float64       `json:"price"`
	Rules        string        `json:"rules"`
	Time         string        `json:"time"`
	Title        string        `json:"title"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	OrganizerID  uint64        `json:"organizerId"`
	SportID      uint64        `json:"sportId"`
	Participants []Participant `json:"participants"`
}

func AllActivities() (*[]Activity, error) {
	dbInstance := db.GetDB()

	activities, err := getActivities(dbInstance)
	if err != nil {
		return nil, err
	}

	for i, activity := range activities {
		participants, err := getParticipants(dbInstance, activity.ID)
		if err != nil {
			return nil, err
		}
		if len(participants) != 0 {
			activities[i].Participants = participants
		}
	}

	return &activities, nil
}

func InsertActivity(address string, date string, maxPlayers int, description string, organizerId int64, price float32, rules string, time string, title string, isPrivate bool, sportId uint64) (*Activity, error) {
	dbInstance := db.GetDB()

	tx, err := dbInstance.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	result, err := tx.Exec(
		queries.InsertActivity, address, date,
		description, isPrivate, maxPlayers, price,
		rules, organizerId, time, title, sportId,
	)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var insertedActivity Activity

	err = tx.QueryRow(
		queries.SelectActivityById,
		lastInsertID,
	).Scan(
		&insertedActivity.ID, &insertedActivity.Address, &insertedActivity.Date, &insertedActivity.Description,
		&insertedActivity.IsPrivate, &insertedActivity.MaxPlayers, &insertedActivity.Price,
		&insertedActivity.Rules, &insertedActivity.Time, &insertedActivity.Title, &insertedActivity.CreatedAt,
		&insertedActivity.UpdatedAt, &insertedActivity.OrganizerID, &insertedActivity.SportID,
	)
	if err != nil {
		return nil, err
	}

	return &insertedActivity, nil
}

func getActivities(dbInstance *sql.DB) ([]Activity, error) {
	rows, err := dbInstance.Query(queries.SelectAllActivities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []Activity

	for rows.Next() {
		var activity Activity

		err := rows.Scan(
			&activity.ID, &activity.Address, &activity.Date, &activity.Description, &activity.IsPrivate,
			&activity.MaxPlayers, &activity.Price, &activity.Rules, &activity.Time, &activity.Title,
			&activity.CreatedAt, &activity.UpdatedAt, &activity.OrganizerID, &activity.SportID,
		)
		if err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}

	return activities, err
}
