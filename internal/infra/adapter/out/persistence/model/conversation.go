package model

import (
	"time"

	"github.com/google/uuid"
)

// ConversationModel represents a conversation entity in the persistence layer.
type ConversationModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type      string    `gorm:"not null"`
	Name      *string
	CreatedAt time.Time

	Members []ConversationMemberModel `gorm:"foreignKey:ConversationID;constraint:OnDelete:CASCADE;"`
	Messages []MessageModel `gorm:"foreignKey:ConversationID;constraint:OnDelete:CASCADE;"`
}

// TableName returns the name of the table in the database.
func (ConversationModel) TableName() string {
	return "conversations"
}