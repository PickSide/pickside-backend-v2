package data

import "time"

type Chatroom struct {
	ID               uint64 `json:"id"`
	Name             string
	NumberOfMessages int `json:"numberOfMessages"`
	LastMessageID    int `json:"lastMessageId"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	LastMessage      Message `json:"last_message" gorm:"foreignKey:LastMessageID"`
}
