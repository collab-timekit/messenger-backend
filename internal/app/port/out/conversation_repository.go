package out

import (
	"messenger/internal/domain"
	"github.com/google/uuid"
)

// ConversationRepository defines the methods for managing conversation members.
type ConversationRepository interface {
	GetAllByUserID(userID uuid.UUID) ([]domain.Conversation, error)
	Create(conv *domain.Conversation) error
	GetByID(id uuid.UUID) (*domain.Conversation, error)
	FindPrivateConversationIDBetweenUsers(user1, user2 uuid.UUID) (*uuid.UUID, error)
}