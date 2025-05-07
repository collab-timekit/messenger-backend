package domain

import (
	"time"
)

type ChatID string

type Chat struct {
	ID        ChatID
	Title     string
	CreatedAt time.Time
}