package data

import (
	"database/sql"
	"pickside/service/db"
	"pickside/service/db/queries"
	"strings"
	"time"
)

type Activity struct {
	ID           uint64        `json:"id"`
	Address      string        `json:"address"`
	Date         string        `json:"date"`
	Description  string        `json:"description"`
	GameMode     string        `json:"gameMode"`
	Images       string        `json:"images"`
	IsPrivate    bool          `json:"isPrivate"`
	Lat          float64       `json:"lat"`
	Lng          float64       `json:"lng"`
	MaxPlayers   int           `json:"maxPlayers"`
	OrganizerID  uint64        `json:"organizerId"`
	Price        float64       `json:"price"`
	Rules        string        `json:"rules"`
	SportID      uint64        `json:"sportId"`
	Time         string        `json:"time"`
	Title        string        `json:"title"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	Organizer    User          `json:"organizer"`
	Participants []Participant `json:"participants"`
	Sport        Sport         `json:"sport"`
}

func AllActivities() (*[]Activity, error) {
	dbInstance := db.GetDB()

	activities, err := getAllActivities(dbInstance)
	if err != nil {
		return nil, err
	}

	for i, activity := range *activities {
		participants, err := getParticipants(dbInstance, activity.ID)
		if err != nil {
			return nil, err
		}

		if len(participants) != 0 {
			(*activities)[i].Participants = participants
		}

		organizer, err := getUserDetails(dbInstance, "id", activity.OrganizerID)
		if err != nil {
			return nil, err
		}

		(*activities)[i].Organizer = *organizer

		sport, err := getSportById(dbInstance, activity.SportID)
		if err != nil {
			return nil, err
		}

		(*activities)[i].Sport = *sport
	}

	return activities, nil
}

func CreateActivity(address string, date string, description string, gameMode string, images []string, isPrivate bool, lat float32, lng float32, maxPlayers int, organizerId uint64, price float32, rules string, sportId uint64, time string, title string) (*Activity, error) {
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

	result, err := tx.Exec(queries.InsertActivity,
		address,
		date,
		description,
		gameMode,
		strings.Join(images, ","),
		isPrivate,
		lat,
		lng,
		maxPlayers,
		organizerId,
		price,
		rules,
		sportId,
		time,
		title,
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
		&insertedActivity.ID,
		&insertedActivity.Address,
		&insertedActivity.Date,
		&insertedActivity.Description,
		&insertedActivity.GameMode,
		&insertedActivity.Images,
		&insertedActivity.IsPrivate,
		&insertedActivity.Lat,
		&insertedActivity.Lng,
		&insertedActivity.MaxPlayers,
		&insertedActivity.OrganizerID,
		&insertedActivity.Price,
		&insertedActivity.Rules,
		&insertedActivity.SportID,
		&insertedActivity.Time,
		&insertedActivity.Title,
		&insertedActivity.CreatedAt,
		&insertedActivity.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &insertedActivity, nil
}

func getAllActivities(dbInstance *sql.DB) (*[]Activity, error) {
	rows, err := dbInstance.Query(queries.SelectAllActivities)
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
			&activity.GameMode,
			&activity.Images,
			&activity.IsPrivate,
			&activity.Lat,
			&activity.Lng,
			&activity.MaxPlayers,
			&activity.OrganizerID,
			&activity.Price,
			&activity.Rules,
			&activity.SportID,
			&activity.Time,
			&activity.Title,
			&activity.CreatedAt,
			&activity.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}

	return &activities, err
}

func getActivityById(dbInstance *sql.DB, activityId string) (*Activity, error) {
	var activity Activity

	err := dbInstance.QueryRow(queries.SelectActivityById, activityId).Scan(
		&activity.Address,
		&activity.Date,
		&activity.Description,
		&activity.GameMode,
		&activity.ID,
		&activity.Images,
		&activity.IsPrivate,
		&activity.Lat,
		&activity.Lng,
		&activity.MaxPlayers,
		&activity.OrganizerID,
		&activity.Price,
		&activity.Rules,
		&activity.SportID,
		&activity.Time,
		&activity.Title,
	)

	return &activity, err
}
