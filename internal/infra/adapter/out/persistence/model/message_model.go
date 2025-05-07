package model

import "time"

type MessageModel struct {
	ID        string    `gorm:"primaryKey;column:id"`
	ChatID    string    `gorm:"column:chat_id"`
	SenderID  string    `gorm:"column:sender_id"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (MessageModel) TableName() string {
	return "messages"
}