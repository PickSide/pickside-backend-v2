package data

import (
	"pickside/service/db"
	"time"
)

type NotificationType string

type Notification struct {
	ID        uint64    `json:"id"`
	Expires   time.Time `json:"expires"`
	IsRead    bool      `json:"isRead"`
	Content   string    `json:"content"`
	Recipient uint64    `json:"recipients"`
	Title     string    `json:"title,omitempty"`
}

func GetUserNotifications(userId uint64) (*[]Notification, error) {
	dbInstance := db.GetDB()

	rows, err := dbInstance.Query("SELECT * FROM notifications WHERE recipient_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []Notification

	for rows.Next() {
		var notification Notification

		err := rows.Scan(
			&notification.ID,
			&notification.Expires,
			&notification.IsRead,
			&notification.Content,
			&notification.Recipient,
			&notification.Title,
		)
		if err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	return &notifications, nil
}

func GetGlobalNotifications() (*[]GlobalNotification, error) {
	dbInstance := db.GetDB()

	rows, err := dbInstance.Query("SELECT * FROM global_notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var globalNotifications []GlobalNotification

	for rows.Next() {
		var globalNotification GlobalNotification

		err := rows.Scan(
			&globalNotification.ID,
			&globalNotification.Expires,
			&globalNotification.Content,
			&globalNotification.Title,
		)
		if err != nil {
			return nil, err
		}

		globalNotifications = append(globalNotifications, globalNotification)
	}

	return &globalNotifications, nil
}
