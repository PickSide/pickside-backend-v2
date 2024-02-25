package data

import (
	"database/sql"
	"pickside/service/db"
	"pickside/service/db/queries"
)

type Group struct {
	ID               uint64 `json:"id"`
	Description      string `json:"description"`
	Name             string `json:"name"`
	RequiresApproval bool   `json:"requiresApproval"`
	Visibility       string `json:"visibility"`
	CreatedAt        string `json:"createdAt,omitempty"`
	UpdatedAt        string `json:"updatedAt,omitempty"`
	OrganizerID      uint64 `json:"organizerId"`
	SportID          uint64 `json:"sportId"`
}

type GroupUser struct {
	GroupID uint64
	UserID  uint64
}

func AllGroupsByOrganizer(organizerId uint64) (*[]Group, error) {
	dbInstance := db.GetDB()

	rows, err := dbInstance.Query(
		`
			SELECT id, description, name, organizer_id, requires_approval, sport_id, visibility 
			FROM groups 
			WHERE organizer_id = ?
		`, organizerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Group

	for rows.Next() {
		var group Group

		err := rows.Scan(
			&group.ID,
			&group.Description,
			&group.Name,
			&group.OrganizerID,
			&group.RequiresApproval,
			&group.SportID,
			&group.Visibility,
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

func InsertGroup(group Group) (*Group, error) {
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
		`
			INSERT INTO groups (
				description, 
				name, 
				organizer_id, 
				requires_approval, 
				sport_id,
				visibility 
			)
			VALUES (?, ?, ?, ?, ?, ?)
		`,
		group.Description,
		group.Name,
		group.OrganizerID,
		group.RequiresApproval,
		group.SportID,
		group.Visibility,
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
		` 
			SELECT id, description, name, organizer_id, requires_approval, sport_id, visibility
    		FROM groups
    		WHERE id = ?
		`,
		lastInsertID,
	).Scan(
		&insertedGroup.ID,
		&insertedGroup.Description,
		&insertedGroup.Name,
		&insertedGroup.OrganizerID,
		&insertedGroup.RequiresApproval,
		&insertedGroup.SportID,
		&insertedGroup.Visibility,
	)
	if err != nil {
		return nil, err
	}

	return &insertedGroup, nil
}

func InsertGroupUsers(groupId uint64, userId uint64) error {
	dbInstance := db.GetDB()

	tx, err := dbInstance.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	_, err = tx.Exec(queries.InsertIntoGroupUsers,
		groupId,
		userId,
	)
	if err != nil {
		return err
	}

	return err
}

func DeleteGroup(groupId uint64, organizerId uint64) (*sql.Result, error) {
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

	_, err = tx.Exec("DELETE FROM users_groups WHERE group_id = ?", groupId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	results, err := tx.Exec("DELETE FROM groups WHERE id = ?", groupId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &results, nil
}

func DeleteUsersGroups(groupId uint64, user_id uint64) (*sql.Result, error) {
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

	result, err := tx.Exec(`DELETE FROM groups_users WHERE group_id = ? AND organizer_id = ?`,
		groupId,
		user_id,
	)

	return &result, err
}
