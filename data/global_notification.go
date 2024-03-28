package data

import "time"

type GlobalNotification struct {
	ID      uint64    `json:"id"`
	Expires time.Time `json:"expires"`
	Content string    `json:"content"`
	Title   string    `json:"title,omitempty"`
}
