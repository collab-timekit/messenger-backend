package model

import (
	"time"

	"github.com/google/uuid"
)

// MessageModel represents a message entity in the system.
type MessageModel struct {
    ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
    ConversationID uuid.UUID
    SenderID       uuid.UUID
    Content        string
    CreatedAt      time.Time
    Edited         bool
}

// TableName specifies the database table name for the MessageModel.
func (MessageModel) TableName() string {
    return "messages"
}