package model

import (
	"time"

	"github.com/google/uuid"
)

// MessageModel represents the structure of a message in the persistence layer.
type MessageModel struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	ConversationID uuid.UUID `gorm:"type:uuid;not null;index"`
	SenderID       uuid.UUID `gorm:"type:uuid;not null"`
	Content        string
	CreatedAt      time.Time
	Edited         bool
}

// TableName returns the name of the table in the database.
func (MessageModel) TableName() string {
    return "messages"
}