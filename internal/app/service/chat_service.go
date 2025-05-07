package service

import (
	"messenger/internal/app/port/in"
	"messenger/internal/app/port/out"
	"messenger/internal/domain"
	"time"

	"github.com/google/uuid"
)

// ChatService implements ChatUseCase
var _ in.ChatUseCase = (*ChatService)(nil)

type ChatService struct {
	repo out.ChatRepository
}

func NewChatService(repo out.ChatRepository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) Create(title string) (*domain.Chat, error) {
	c := &domain.Chat{
		ID:        domain.ChatID(uuid.New().String()),
		Title:     title,
		CreatedAt: time.Now(),
	}
	if err := s.repo.Save(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *ChatService) GetByID(id string) (*domain.Chat, error) {
	return s.repo.FindByID(id)
}

func (s *ChatService) List() ([]*domain.Chat, error) {
	return s.repo.FindAll()
}