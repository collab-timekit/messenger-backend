package in

import (
	"messenger/internal/domain"
	"github.com/google/uuid"
)

// ConversationUseCase defines the methods for managing conversations and their messages.
type ConversationUseCase interface {
	GetUserConversations(userID uuid.UUID) ([]domain.Conversation, error)
	CreateConversation(conv *domain.Conversation) error
	GetOrCreatePrivateConversation(userID, recipientID uuid.UUID) (uuid.UUID, error)
	GetConversationByID(conversationID uuid.UUID, userID uuid.UUID) (*domain.Conversation, error)
}