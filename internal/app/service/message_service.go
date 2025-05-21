package service

import (
	"messenger/internal/app/port/in"
	"messenger/internal/app/port/out"
	"messenger/internal/domain"
	"time"

	"github.com/google/uuid"
)

// MessageService provides methods for working with messages.
type MessageService struct {
	messageRepo out.MessageRepository
}

// NewMessageService creates a new instance of MessageService.
func NewMessageService(messageRepo out.MessageRepository) in.MessageUseCase {
	return &MessageService{messageRepo: messageRepo}
}

// CreateMessage sends a new message in a conversation.
func (s *MessageService) CreateMessage(conversationID uuid.UUID, senderID uuid.UUID, content string) error {
	message := &domain.Message{
		ID:             uuid.New(),
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        content,
		CreatedAt:      time.Now(),
		Edited:         false,
	}

	return s.messageRepo.SendMessage(message)
}

// GetMessagesForConversation retrieves all messages for a given conversation, with pagination.
func (s *MessageService) GetMessagesForConversation(conversationID uuid.UUID, limit int, before *time.Time) ([]domain.Message, error) {
	return s.messageRepo.GetMessages(conversationID, limit, before)
}

// EditMessage allows editing an existing message.
func (s *MessageService) EditMessage(messageID uuid.UUID, newContent string) error {
	return s.messageRepo.EditMessage(messageID, newContent)
}

// DeleteMessage deletes a message from a conversation.
func (s *MessageService) DeleteMessage(messageID uuid.UUID) error {
	return s.messageRepo.DeleteMessage(messageID)
}