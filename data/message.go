package data

type Message struct {
	ID         uint64 `json:"id"`
	Content    string `json:"content"`
	Delivered  bool   `json:"delivered"`
	SentAt     string `json:"sentAt"`
	ChatroomID uint64 `json:"chatroomId"`
	SenderID   uint64 `json:"senderId"`
}
