package out

import "messenger/internal/domain"

type ChatRepository interface {
	Save(c *domain.Chat) error
	FindAll() ([]*domain.Chat, error)
    FindByID(id string) (*domain.Chat, error)
}