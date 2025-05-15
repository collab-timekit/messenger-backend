package in

import (
	"messenger/internal/domain"
	"time"
	"github.com/google/uuid"
)

// MessageUseCase defines the contract for message-related operations.
type MessageUseCase interface {
	GetMessagesForConversation(conversationID uuid.UUID, limit int, before *time.Time) ([]domain.Message, error)
	CreateMessage(conversationID uuid.UUID, senderID uuid.UUID, content string) error
	EditMessage(messageID uuid.UUID, newContent string) error
	DeleteMessage(messageID uuid.UUID) error
}