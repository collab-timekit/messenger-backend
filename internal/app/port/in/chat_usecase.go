package in

import "messenger/internal/domain"

type ChatUseCase interface {
	Create(title string) (*domain.Chat, error)
	GetByID(id string) (*domain.Chat, error)
	List() ([]*domain.Chat, error)
}