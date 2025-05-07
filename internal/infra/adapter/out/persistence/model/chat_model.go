package model

import (
	"time"
)

type ChatModel struct {
	ID        string    `gorm:"primaryKey;column:id"`
	Title     string    `gorm:"column:title"`
	IsPublic  bool      `gorm:"column:is_public;default:true"`
	CreatedAt time.Time `gorm:"column:created_at"`

	Messages []MessageModel `gorm:"foreignKey:ChatID"`
}

func (ChatModel) TableName() string {
	return "chats"
}