package model

import (
	"time"

	"github.com/google/uuid"
)

// ConversationMemberModel represents a member of a conversation with their role and join time.
type ConversationMemberModel struct {
	ConversationID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Role           string
	JoinedAt       time.Time
}

// TableName returns the name of the table in the database.
func (ConversationMemberModel) TableName() string {
    return "conversation_members"
}