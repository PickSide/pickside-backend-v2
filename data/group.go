package data

import (
	"database/sql"
	"pickside/service/db"
	"pickside/service/db/queries"
)

type Group struct {
	ID               uint64 `json:"id"`
	CoverPhoto       string `json:"coverPhoto"`
	Description      string `json:"description"`
	Name             string `json:"name"`
	RequiresApproval bool   `json:"requiresApproval"`
	Visibility       string `json:"visibility"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
	OrganizerID      uint64 `json:"organizerId"`
	SportID          uint64 `json:"sportId"`
}

type GroupUser struct {
	GroupID uint64
	UserID  uint64
}

func AllGroupsByOrganizer(organizerID uint64) (*[]Group, error) {
	dbInstance := db.GetDB()

	// Fetch all groups directly using JOIN
	rows, err := dbInstance.Query(queries.SelectAllGroupsByOrganizer, organizerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Group

	for rows.Next() {
		var group Group

		err := rows.Scan(
			&group.ID,
			&group.CoverPhoto,
			&group.Description,
			&group.Name,
			&group.RequiresApproval,
			&group.Visibility,
			&group.CreatedAt,
			&group.UpdatedAt,
			&group.OrganizerID,
			&group.SportID,
		)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &groups, nil
}

func InsertGroup(coverPhoto string, description string, name string, requiresApproval bool, visibility string, organizerId uint64, sportId uint64) (*Group, error) {
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
		queries.InsertIntoGroup,
		coverPhoto,
		description,
		name,
		requiresApproval,
		visibility,
		organizerId,
		sportId,
	)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var insertedGroup Group

	err = tx.QueryRow(
		queries.SelectGroupById,
		lastInsertID,
	).Scan(
		&insertedGroup.ID,
		&insertedGroup.CoverPhoto,
		&insertedGroup.Description,
		&insertedGroup.Name,
		&insertedGroup.RequiresApproval,
		&insertedGroup.Visibility,
		&insertedGroup.CreatedAt,
		&insertedGroup.UpdatedAt,
		&insertedGroup.OrganizerID,
		&insertedGroup.SportID,
	)
	if err != nil {
		return nil, err
	}

	err = InsertGroupUsers(tx, insertedGroup.ID, insertedGroup.OrganizerID)
	if err != nil {
		return nil, err
	}

	return &insertedGroup, nil
}

func InsertGroupUsers(tx *sql.Tx, groupId uint64, userId uint64) error {
	_, err := tx.Exec(queries.InsertIntoGroupUsers,
		groupId,
		userId,
	)

	return err
}
