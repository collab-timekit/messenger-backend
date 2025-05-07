package out

import "messenger/internal/domain"

type MessageRepository interface {
	FindByChatID(chatID string, limit, offset int) ([]domain.Message, error)
	Save(message *domain.Message) error
}