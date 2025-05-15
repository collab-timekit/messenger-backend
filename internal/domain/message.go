package domain

import (
	"time"
	"github.com/google/uuid"
)

type Message struct {
	ID             uuid.UUID
	ConversationID uuid.UUID
	SenderID       uuid.UUID
	Content        string
	CreatedAt      time.Time
	Edited         bool
}