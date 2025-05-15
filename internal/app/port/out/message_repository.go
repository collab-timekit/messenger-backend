package out

import (
	"messenger/internal/domain"
	"time"
	"github.com/google/uuid"
)

// MessageRepository defines the methods for managing messages in the repository.
type MessageRepository interface {
	SendMessage(msg *domain.Message) error
	GetMessages(conversationID uuid.UUID, limit int, before *time.Time) ([]domain.Message, error)
	EditMessage(messageID uuid.UUID, newContent string) error
	DeleteMessage(messageID uuid.UUID) error
}