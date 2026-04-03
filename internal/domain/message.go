package domain

import "time"

type Message struct {
	ID         int       `json:"id"`
	FromUserID int       `json:"fromUserId"`
	ToUserID   int       `json:"toUserId"`
	Content    string    `json:"content"`
	SentAt     time.Time `json:"sentAt"`
}