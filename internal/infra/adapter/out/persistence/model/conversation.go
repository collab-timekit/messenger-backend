package model

import (
	"time"

	"github.com/google/uuid"
)

// ConversationModel represents a conversation with its metadata, members, and messages.
type ConversationModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type      string    `gorm:"not null"`
	Name      *string
	CreatedAt time.Time
	Members   []ConversationMemberModel
	Messages  []MessageModel
}