package in

import "messenger/internal/domain"

type MessageUseCase interface {
	GetMessages(chatID string, limit, offset int) ([]domain.Message, error)
	SendMessage(chatID, senderID, content string) (*domain.Message, error)
}