package domain

import (
	"time"
	"github.com/google/uuid"
)

// ConversationMember represents a member of a conversation with their role and join time.
type ConversationMember struct {
	ConversationID uuid.UUID `json:"conversation_id"`
	UserID         uuid.UUID `json:"user_id"`
	Role           string    `json:"role"`
	JoinedAt       time.Time `json:"joined_at"`
}