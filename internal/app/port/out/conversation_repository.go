package out

import (
	"messenger/internal/domain"
	"github.com/google/uuid"
)

type ConversationRepository interface {
	// GetAllByUserID retrieves all conversations by the user ID.
	GetAllByUserID(userID uuid.UUID) ([]domain.Conversation, error)

	// Create creates a new conversation.
	Create(conv *domain.Conversation) error
	// GetByID retrieves a conversation by its ID.
	GetByID(id uuid.UUID) (*domain.Conversation, error)
}