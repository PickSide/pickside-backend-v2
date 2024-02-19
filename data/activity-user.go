package data

import (
	"database/sql"
	"log"
	"pickside/service/db"
	"pickside/service/db/queries"
	"pickside/service/types"
	"time"
)

type Participant struct {
	ID                  uint64     `json:"id"`
	Avatar              string     `json:"avatar"`
	Bio                 string     `json:"bio"`
	City                string     `json:"city"`
	Email               string     `json:"email"`
	FullName            string     `json:"fullName"`
	JoinDate            time.Time  `json:"joinDate"`
	LocaleRegion        string     `json:"localeRegion"`
	MatchOrganizedCount int        `json:"matchOrganizedCount"`
	MatchPlayedCount    int        `json:"matchPlayedCount"`
	Sexe                types.Sexe `json:"sexe"`
	Username            string     `json:"username"`
}

func UpdateParticipants(activityId uint64, userId uint64) (*[]Participant, bool) {
	dbInstance := db.GetDB()

	rows, err := dbInstance.Query(queries.IsRegistered, activityId, userId)
	if err != nil {
		return nil, false
	}
	defer rows.Close()

	if rows.Next() {
		_, err := dbInstance.Exec(queries.RemoveFromActivity, activityId, userId)
		if err != nil {
			return nil, false
		}
	} else {
		_, err := dbInstance.Exec(queries.RegisterTo, activityId, userId)
		if err != nil {
			return nil, false
		}
	}

	participants, err := getParticipants(dbInstance, activityId)
	if err != nil {
		log.Println("Error fetching participants:", err)
		return nil, false
	}

	return &participants, true
}

func GetParticipants(activityId uint64) (*[]Participant, error) {
	dbInstance := db.GetDB()

	participants, err := getParticipants(dbInstance, activityId)
	if err != nil {
		return nil, err
	}

	return &participants, nil
}

func getParticipants(dbInstance *sql.DB, activityId uint64) ([]Participant, error) {
	rows, err := dbInstance.Query(` 
		SELECT u.id, u.avatar, u.bio, u.city, u.email, u.full_name, u.join_date, u.locale_region, u.match_organized_count, u.match_played_count, u.sexe, u.username
    	FROM activity_users
    	LEFT JOIN users AS u ON activity_users.user_id = u.id
    	WHERE activity_users.activity_id = ?;`,
		activityId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []Participant

	for rows.Next() {
		var participant Participant

		err := rows.Scan(
			&participant.ID, &participant.Avatar, &participant.Bio, &participant.City, &participant.Email,
			&participant.FullName, &participant.JoinDate, &participant.LocaleRegion,
			&participant.MatchOrganizedCount, &participant.MatchPlayedCount, &participant.Sexe, &participant.Username,
		)
		if err != nil {
			return nil, err
		}

		participants = append(participants, participant)
	}

	return participants, nil
}
