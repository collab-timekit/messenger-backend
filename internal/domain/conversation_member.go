package domain

import (
	"time"
	"github.com/google/uuid"
)

type ConversationMember struct {
	ConversationID uuid.UUID
	UserID         uuid.UUID
	Role           string
	JoinedAt       time.Time
}