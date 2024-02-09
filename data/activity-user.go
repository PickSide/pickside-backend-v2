package data

import (
	"pickside/service/db"
	"pickside/service/db/queries"
)

type ActivityUser struct {
	ActivityID uint64 `json:"activityId"`
	UserID     string `json:"userId"`
}

func UpdateParticipants(activityId uint64, userId uint64) bool {
	dbInstance := db.GetDB()

	rows, err := dbInstance.Query(queries.IsUserRegistered, activityId, userId)
	if err != nil {
		return false
	}
	defer rows.Close()

	if rows.Next() {
		_, err := dbInstance.Query(queries.RemoveUserFromActivity, activityId, userId)
		if err != nil {
			return false
		}
	} else {
		_, err := dbInstance.Query(queries.RegisterUserToActivity, activityId, userId)
		if err != nil {
			return false
		}
	}

	return true
}

func GetParticipants(activityId uint64) (*[]ActivityUser, error) {
	dbInstance := db.GetDB()

	rows, err := dbInstance.Query(queries.SelectAllParticipants, activityId)
	if err != nil {
		return nil, err
	}

	var registrations []ActivityUser

	for rows.Next() {
		var registration ActivityUser

		err := rows.Scan(
			&registration.ActivityID,
			&registration.UserID,
		)
		if err != nil {
			return nil, err
		}

		registrations = append(registrations, registration)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &registrations, nil
}
