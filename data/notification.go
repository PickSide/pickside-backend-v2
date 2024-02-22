package data

type Notification struct {
	ID         uint64 `json:"id"`
	Expires    string `json:"expires"`
	IsRead     bool   `json:"isRead"`
	Message    string `json:"message"`
	ReceiverID uint64 `json:"receiverId"`
	SenderID   uint64 `json:"senderId"`
	Type       string `json:"type"`
}
