package domain

import "time"

type Message struct {
	ID        string
	ChatID    string
	SenderID  string
	Content   string
	CreatedAt time.Time
}