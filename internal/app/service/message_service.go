package service

import (
	"messenger/internal/app/port/in"
	"messenger/internal/app/port/out"
	"messenger/internal/domain"
	"time"

	"github.com/google/uuid"
)

// MessageService implements MessageUseCase
var _ in.MessageUseCase = (*MessageService)(nil)

type MessageService struct {
	repo out.MessageRepository
}

func NewMessageService(repo out.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) GetMessages(chatID string, limit, offset int) ([]domain.Message, error) {
	return s.repo.FindByChatID(chatID, limit, offset)
}

func (s *MessageService) SendMessage(chatID, senderID, content string) (*domain.Message, error) {
	message := &domain.Message{
		ID:        uuid.NewString(),
		ChatID:    chatID,
		SenderID:  senderID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Save(message); err != nil {
		return nil, err
	}
	return message, nil
}