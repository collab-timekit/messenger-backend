package in

import (
	"messenger/internal/domain"
	"github.com/google/uuid"
)

// ConversationMemberUseCase defines the methods for managing conversation members.
type ConversationMemberUseCase interface {
	AddMember(conversationID, userID uuid.UUID, role string) error
	RemoveMember(conversationID, userID uuid.UUID) error
	GetMembers(conversationID uuid.UUID) ([]domain.ConversationMember, error)
	IsMember(conversationID, userID uuid.UUID) (bool, error)
}